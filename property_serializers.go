package manta

import (
	"regexp"
	"strconv"
)

// Type for a decoder function
type DecodeFcn func(*Reader, *dt_field) interface{}

// PropertySerializer interface
type PropertySerializer struct {
	Decode          DecodeFcn
	DecodeContainer DecodeFcn
	IsArray         bool
	Length          uint32
	ArraySerializer *PropertySerializer
	Name            string
}

// Contains a list of available property serializers
type PropertySerializerTable struct {
	Serializers map[string]*PropertySerializer
}

// Returns a table containing all know property serializers
func GetDefaultPropertySerializerTable() *PropertySerializerTable {
	return &PropertySerializerTable{Serializers: map[string]*PropertySerializer{}}
}

// Regex for array and vector
var matchArray = regexp.MustCompile(`([^[\]]+)\[(\d+)]`)
var matchVector = regexp.MustCompile(`CUtlVector\<\s(.*)\s>$`)

// Flags used in the quantized decoding
const qf_rounddown int32 = (1 << 0)
const qf_roundup int32 = (1 << 1)
const qf_encode_zero int32 = (1 << 2)
const qf_encode_integers int32 = (1 << 3)

// Initializes a QuantizedFloat encoder
func InitQuantized(f *dt_field) {
	// Make sure all field have their default values
	var BitCount int
	var Low float32
	var High float32
	var Flags int32

	if f.BitCount != nil {
		BitCount = int(*f.BitCount)
	}

	if f.LowValue != nil {
		Low = *f.LowValue
	} else {
		Low = 0.0
	}

	if f.HighValue != nil {
		High = *f.HighValue
	} else {
		High = 1.0
	}

	if f.Flags != nil {
		Flags = *f.Flags
	} else {
		Flags = 0
	}

	// if the bitcount is <= 0 or >= 32, treat this as a noscale with 32 bits
	if BitCount <= 0 || BitCount >= 32 {
		f.Serializer = &PropertySerializer{decodeFloatNoscale, nil, false, 0, nil, "unkown"}
		return
	}

	// Discard zero flag when encoding min / max set to 0
	if (Low == 0.0 && (Flags&qf_rounddown) != 0) || (High == 0.0 && (Flags&qf_roundup) != 0) {
		Flags &= ^qf_encode_zero
	}

	// If min / max is zero when encoding zero, switch to round up / round down instead
	if Low == 0.0 && (Flags&qf_encode_zero) != 0 {
		Flags |= qf_rounddown
		Flags &= ^qf_encode_zero
	}

	if High == 0.0 && (Flags&qf_encode_zero) != 0 {
		Flags |= qf_roundup
		Flags &= ^qf_encode_zero
	}

	// Check if the range spans zero
	if Low > 0.0 || High < 0.0 {
		Flags &= ^qf_encode_zero
	}

	// If we are left with encode zero, only leave integer flag
	if (Flags & qf_encode_zero) != 0 {
		Flags &= ^(qf_roundup | qf_rounddown | qf_encode_zero)
	}

	f.Flags = &Flags
	f.Serializer = &PropertySerializer{decodeQuantized, nil, false, 0, nil, "unkown"}
}

// Fills serializer in dt_field
func (pst *PropertySerializerTable) FillSerializer(field *dt_field) {
	// Handle special decoders that need the complete field data here
	switch field.Type {
	case "CNetworkedQuantizedFloat":
		InitQuantized(field)
		return
	}

	field.Serializer = pst.GetPropertySerializerByName(field.Type)
}

// Returns a serializer by name
func (pst *PropertySerializerTable) GetPropertySerializerByName(name string) *PropertySerializer {
	// Return existing serializer
	if ser := pst.Serializers[name]; ser != nil {
		return ser
	}

	// Set decoder
	var decoder DecodeFcn
	var decoderContainer DecodeFcn

	switch name {
	case "float32":
		decoder = decodeFloat
	case "int8":
		fallthrough
	case "int16":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		decoder = decodeSigned
	case "uint8":
		fallthrough
	case "uint16":
		fallthrough
	case "uint32":
		fallthrough
	case "uint64":
		fallthrough
	case "Color":
		decoder = decodeUnsigned
	case "char":
		fallthrough
	case "CUtlSymbolLarge":
		decoder = decodeString
	case "Vector":
		decoder = decodeFVector
	case "bool":
		decoder = decodeBoolean
	case "CNetworkedQuantizedFloat":
		decoder = decodeQuantized
	case "CRenderComponent":
		fallthrough
	case "CPhysicsComponent":
		fallthrough
	case "CBodyComponent":
		decoder = decodeComponent
	case "CDOTASpectatorGraphManager*":
		fallthrough
	case "CEntityIdentity*":
		decoder = decodeBoolean
	case "QAngle":
		decoder = decodeQAngle
	case "CGameSceneNodeHandle":
		decoder = decodeHandle
	default:
		// check for specific types
		switch {
		case hasPrefix(name, "CHandle"):
			decoder = decodeHandle
		case hasPrefix(name, "CStrongHandle"):
			decoder = decodeUnsigned
		case hasPrefix(name, "CUtlVector< "):
			if match := matchVector.FindStringSubmatch(name); match != nil {
				decoderContainer = decodeVector
				decoder = pst.GetPropertySerializerByName(match[1]).Decode
			} else {
				_panicf("Unable to read vector type for %s", name)
			}
		default:
			//_debugf("No decoder for type %s", name)
		}
	}

	// create a new serializer based on it's name
	if match := matchArray.FindStringSubmatch(name); match != nil {
		typeName := match[1]
		length, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil {
			_panicf("Array length doesn't seem to be a number: %v", match[2])
		}

		serializer, found := pst.Serializers[typeName]
		if !found {
			serializer = pst.GetPropertySerializerByName(typeName)
			pst.Serializers[typeName] = serializer
		}

		ps := &PropertySerializer{
			Decode:          serializer.Decode,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(length),
			ArraySerializer: serializer,
			Name:            typeName,
		}
		pst.Serializers[name] = ps
		return ps
	}

	if match := matchVector.FindStringSubmatch(name); match != nil {
		ps := &PropertySerializer{
			Decode:          decoder,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(128),
			ArraySerializer: &PropertySerializer{},
		}
		pst.Serializers[name] = ps
		return ps
	}

	if name == "C_DOTA_ItemStockInfo[MAX_ITEM_STOCKS]" {
		typeName := "C_DOTA_ItemStockInfo"

		serializer, found := pst.Serializers[typeName]
		if !found {
			serializer = pst.GetPropertySerializerByName(typeName)
			pst.Serializers[typeName] = serializer
		}

		ps := &PropertySerializer{
			Decode:          serializer.Decode,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(8),
			ArraySerializer: serializer,
			Name:            typeName,
		}

		pst.Serializers[name] = ps
		return ps
	}

	if name == "CDOTA_AbilityDraftAbilityState[MAX_ABILITY_DRAFT_ABILITIES]" {
		typeName := "CDOTA_AbilityDraftAbilityState"

		serializer, found := pst.Serializers[typeName]
		if !found {
			serializer = pst.GetPropertySerializerByName(typeName)
			pst.Serializers[typeName] = serializer
		}

		ps := &PropertySerializer{
			Decode:          serializer.Decode,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(48),
			ArraySerializer: serializer,
			Name:            typeName,
		}

		pst.Serializers[name] = ps
		return ps
	}

	// That the type does not indicate an array is somewhat bad for the way we are
	// parsing things at the moment :(
	if name == "m_SpeechBubbles" {
		typeName := "m_SpeechBubbles"

		ps := &PropertySerializer{
			Decode:          decoder,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(5),
			ArraySerializer: nil,
			Name:            typeName,
		}

		pst.Serializers[name] = ps
		return ps
	}

	if name == "DOTA_PlayerChallengeInfo" {
		typeName := "DOTA_PlayerChallengeInfo"

		ps := &PropertySerializer{
			Decode:          decoder,
			DecodeContainer: decoderContainer,
			IsArray:         true,
			Length:          uint32(30),
			ArraySerializer: nil,
			Name:            typeName,
		}

		pst.Serializers[name] = ps
		return ps
	}

	// This function should panic at some point
	return &PropertySerializer{decoder, decoderContainer, false, 0, nil, "unkown"}
}
