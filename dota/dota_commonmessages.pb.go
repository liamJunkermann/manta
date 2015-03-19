// Code generated by protoc-gen-go.
// source: dota_commonmessages.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import math "math"

// discarding unused import google_protobuf "github.com/dotabuff/manta/dota/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type EDOTAChatWheelMessage int32

const (
	EDOTAChatWheelMessage_k_EDOTA_CW_Ok                  EDOTAChatWheelMessage = 0
	EDOTAChatWheelMessage_k_EDOTA_CW_Care                EDOTAChatWheelMessage = 1
	EDOTAChatWheelMessage_k_EDOTA_CW_GetBack             EDOTAChatWheelMessage = 2
	EDOTAChatWheelMessage_k_EDOTA_CW_NeedWards           EDOTAChatWheelMessage = 3
	EDOTAChatWheelMessage_k_EDOTA_CW_Stun                EDOTAChatWheelMessage = 4
	EDOTAChatWheelMessage_k_EDOTA_CW_Help                EDOTAChatWheelMessage = 5
	EDOTAChatWheelMessage_k_EDOTA_CW_Push                EDOTAChatWheelMessage = 6
	EDOTAChatWheelMessage_k_EDOTA_CW_GoodJob             EDOTAChatWheelMessage = 7
	EDOTAChatWheelMessage_k_EDOTA_CW_Missing             EDOTAChatWheelMessage = 8
	EDOTAChatWheelMessage_k_EDOTA_CW_Missing_Top         EDOTAChatWheelMessage = 9
	EDOTAChatWheelMessage_k_EDOTA_CW_Missing_Mid         EDOTAChatWheelMessage = 10
	EDOTAChatWheelMessage_k_EDOTA_CW_Missing_Bottom      EDOTAChatWheelMessage = 11
	EDOTAChatWheelMessage_k_EDOTA_CW_Go                  EDOTAChatWheelMessage = 12
	EDOTAChatWheelMessage_k_EDOTA_CW_Initiate            EDOTAChatWheelMessage = 13
	EDOTAChatWheelMessage_k_EDOTA_CW_Follow              EDOTAChatWheelMessage = 14
	EDOTAChatWheelMessage_k_EDOTA_CW_Group_Up            EDOTAChatWheelMessage = 15
	EDOTAChatWheelMessage_k_EDOTA_CW_Spread_Out          EDOTAChatWheelMessage = 16
	EDOTAChatWheelMessage_k_EDOTA_CW_Split_Farm          EDOTAChatWheelMessage = 17
	EDOTAChatWheelMessage_k_EDOTA_CW_Attack              EDOTAChatWheelMessage = 18
	EDOTAChatWheelMessage_k_EDOTA_CW_BRB                 EDOTAChatWheelMessage = 19
	EDOTAChatWheelMessage_k_EDOTA_CW_Dive                EDOTAChatWheelMessage = 20
	EDOTAChatWheelMessage_k_EDOTA_CW_OMW                 EDOTAChatWheelMessage = 21
	EDOTAChatWheelMessage_k_EDOTA_CW_Get_Ready           EDOTAChatWheelMessage = 22
	EDOTAChatWheelMessage_k_EDOTA_CW_Bait                EDOTAChatWheelMessage = 23
	EDOTAChatWheelMessage_k_EDOTA_CW_Heal                EDOTAChatWheelMessage = 24
	EDOTAChatWheelMessage_k_EDOTA_CW_Mana                EDOTAChatWheelMessage = 25
	EDOTAChatWheelMessage_k_EDOTA_CW_OOM                 EDOTAChatWheelMessage = 26
	EDOTAChatWheelMessage_k_EDOTA_CW_Skill_Cooldown      EDOTAChatWheelMessage = 27
	EDOTAChatWheelMessage_k_EDOTA_CW_Ulti_Ready          EDOTAChatWheelMessage = 28
	EDOTAChatWheelMessage_k_EDOTA_CW_Enemy_Returned      EDOTAChatWheelMessage = 29
	EDOTAChatWheelMessage_k_EDOTA_CW_All_Missing         EDOTAChatWheelMessage = 30
	EDOTAChatWheelMessage_k_EDOTA_CW_Enemy_Incoming      EDOTAChatWheelMessage = 31
	EDOTAChatWheelMessage_k_EDOTA_CW_Invis_Enemy         EDOTAChatWheelMessage = 32
	EDOTAChatWheelMessage_k_EDOTA_CW_Enemy_Had_Rune      EDOTAChatWheelMessage = 33
	EDOTAChatWheelMessage_k_EDOTA_CW_Split_Push          EDOTAChatWheelMessage = 34
	EDOTAChatWheelMessage_k_EDOTA_CW_Coming_To_Gank      EDOTAChatWheelMessage = 35
	EDOTAChatWheelMessage_k_EDOTA_CW_Request_Gank        EDOTAChatWheelMessage = 36
	EDOTAChatWheelMessage_k_EDOTA_CW_Fight_Under_Tower   EDOTAChatWheelMessage = 37
	EDOTAChatWheelMessage_k_EDOTA_CW_Deny_Tower          EDOTAChatWheelMessage = 38
	EDOTAChatWheelMessage_k_EDOTA_CW_Buy_Courier         EDOTAChatWheelMessage = 39
	EDOTAChatWheelMessage_k_EDOTA_CW_Upgrade_Courier     EDOTAChatWheelMessage = 40
	EDOTAChatWheelMessage_k_EDOTA_CW_Need_Detection      EDOTAChatWheelMessage = 41
	EDOTAChatWheelMessage_k_EDOTA_CW_They_Have_Detection EDOTAChatWheelMessage = 42
	EDOTAChatWheelMessage_k_EDOTA_CW_Buy_TP              EDOTAChatWheelMessage = 43
	EDOTAChatWheelMessage_k_EDOTA_CW_Reuse_Courier       EDOTAChatWheelMessage = 44
	EDOTAChatWheelMessage_k_EDOTA_CW_Deward              EDOTAChatWheelMessage = 45
	EDOTAChatWheelMessage_k_EDOTA_CW_Building_Mek        EDOTAChatWheelMessage = 46
	EDOTAChatWheelMessage_k_EDOTA_CW_Building_Pipe       EDOTAChatWheelMessage = 47
	EDOTAChatWheelMessage_k_EDOTA_CW_Stack_And_Pull      EDOTAChatWheelMessage = 48
	EDOTAChatWheelMessage_k_EDOTA_CW_Pull                EDOTAChatWheelMessage = 49
	EDOTAChatWheelMessage_k_EDOTA_CW_Pulling             EDOTAChatWheelMessage = 50
	EDOTAChatWheelMessage_k_EDOTA_CW_Stack               EDOTAChatWheelMessage = 51
	EDOTAChatWheelMessage_k_EDOTA_CW_Jungling            EDOTAChatWheelMessage = 52
	EDOTAChatWheelMessage_k_EDOTA_CW_Roshan              EDOTAChatWheelMessage = 53
	EDOTAChatWheelMessage_k_EDOTA_CW_Affirmative         EDOTAChatWheelMessage = 54
	EDOTAChatWheelMessage_k_EDOTA_CW_Wait                EDOTAChatWheelMessage = 55
	EDOTAChatWheelMessage_k_EDOTA_CW_Pause               EDOTAChatWheelMessage = 56
	EDOTAChatWheelMessage_k_EDOTA_CW_Current_Time        EDOTAChatWheelMessage = 57
	EDOTAChatWheelMessage_k_EDOTA_CW_Check_Runes         EDOTAChatWheelMessage = 58
	EDOTAChatWheelMessage_k_EDOTA_CW_Smoke_Gank          EDOTAChatWheelMessage = 59
	EDOTAChatWheelMessage_k_EDOTA_CW_GLHF                EDOTAChatWheelMessage = 60
	EDOTAChatWheelMessage_k_EDOTA_CW_Nice                EDOTAChatWheelMessage = 61
	EDOTAChatWheelMessage_k_EDOTA_CW_Thanks              EDOTAChatWheelMessage = 62
	EDOTAChatWheelMessage_k_EDOTA_CW_Sorry               EDOTAChatWheelMessage = 63
	EDOTAChatWheelMessage_k_EDOTA_CW_No_Give_Up          EDOTAChatWheelMessage = 64
	EDOTAChatWheelMessage_k_EDOTA_CW_Just_Happened       EDOTAChatWheelMessage = 65
	EDOTAChatWheelMessage_k_EDOTA_CW_Game_Is_Hard        EDOTAChatWheelMessage = 66
	EDOTAChatWheelMessage_k_EDOTA_CW_New_Meta            EDOTAChatWheelMessage = 67
	EDOTAChatWheelMessage_k_EDOTA_CW_My_Bad              EDOTAChatWheelMessage = 68
	EDOTAChatWheelMessage_k_EDOTA_CW_Regret              EDOTAChatWheelMessage = 69
	EDOTAChatWheelMessage_k_EDOTA_CW_Relax               EDOTAChatWheelMessage = 70
	EDOTAChatWheelMessage_k_EDOTA_CW_MissingHero         EDOTAChatWheelMessage = 71
	EDOTAChatWheelMessage_k_EDOTA_CW_ReturnedHero        EDOTAChatWheelMessage = 72
)

var EDOTAChatWheelMessage_name = map[int32]string{
	0:  "k_EDOTA_CW_Ok",
	1:  "k_EDOTA_CW_Care",
	2:  "k_EDOTA_CW_GetBack",
	3:  "k_EDOTA_CW_NeedWards",
	4:  "k_EDOTA_CW_Stun",
	5:  "k_EDOTA_CW_Help",
	6:  "k_EDOTA_CW_Push",
	7:  "k_EDOTA_CW_GoodJob",
	8:  "k_EDOTA_CW_Missing",
	9:  "k_EDOTA_CW_Missing_Top",
	10: "k_EDOTA_CW_Missing_Mid",
	11: "k_EDOTA_CW_Missing_Bottom",
	12: "k_EDOTA_CW_Go",
	13: "k_EDOTA_CW_Initiate",
	14: "k_EDOTA_CW_Follow",
	15: "k_EDOTA_CW_Group_Up",
	16: "k_EDOTA_CW_Spread_Out",
	17: "k_EDOTA_CW_Split_Farm",
	18: "k_EDOTA_CW_Attack",
	19: "k_EDOTA_CW_BRB",
	20: "k_EDOTA_CW_Dive",
	21: "k_EDOTA_CW_OMW",
	22: "k_EDOTA_CW_Get_Ready",
	23: "k_EDOTA_CW_Bait",
	24: "k_EDOTA_CW_Heal",
	25: "k_EDOTA_CW_Mana",
	26: "k_EDOTA_CW_OOM",
	27: "k_EDOTA_CW_Skill_Cooldown",
	28: "k_EDOTA_CW_Ulti_Ready",
	29: "k_EDOTA_CW_Enemy_Returned",
	30: "k_EDOTA_CW_All_Missing",
	31: "k_EDOTA_CW_Enemy_Incoming",
	32: "k_EDOTA_CW_Invis_Enemy",
	33: "k_EDOTA_CW_Enemy_Had_Rune",
	34: "k_EDOTA_CW_Split_Push",
	35: "k_EDOTA_CW_Coming_To_Gank",
	36: "k_EDOTA_CW_Request_Gank",
	37: "k_EDOTA_CW_Fight_Under_Tower",
	38: "k_EDOTA_CW_Deny_Tower",
	39: "k_EDOTA_CW_Buy_Courier",
	40: "k_EDOTA_CW_Upgrade_Courier",
	41: "k_EDOTA_CW_Need_Detection",
	42: "k_EDOTA_CW_They_Have_Detection",
	43: "k_EDOTA_CW_Buy_TP",
	44: "k_EDOTA_CW_Reuse_Courier",
	45: "k_EDOTA_CW_Deward",
	46: "k_EDOTA_CW_Building_Mek",
	47: "k_EDOTA_CW_Building_Pipe",
	48: "k_EDOTA_CW_Stack_And_Pull",
	49: "k_EDOTA_CW_Pull",
	50: "k_EDOTA_CW_Pulling",
	51: "k_EDOTA_CW_Stack",
	52: "k_EDOTA_CW_Jungling",
	53: "k_EDOTA_CW_Roshan",
	54: "k_EDOTA_CW_Affirmative",
	55: "k_EDOTA_CW_Wait",
	56: "k_EDOTA_CW_Pause",
	57: "k_EDOTA_CW_Current_Time",
	58: "k_EDOTA_CW_Check_Runes",
	59: "k_EDOTA_CW_Smoke_Gank",
	60: "k_EDOTA_CW_GLHF",
	61: "k_EDOTA_CW_Nice",
	62: "k_EDOTA_CW_Thanks",
	63: "k_EDOTA_CW_Sorry",
	64: "k_EDOTA_CW_No_Give_Up",
	65: "k_EDOTA_CW_Just_Happened",
	66: "k_EDOTA_CW_Game_Is_Hard",
	67: "k_EDOTA_CW_New_Meta",
	68: "k_EDOTA_CW_My_Bad",
	69: "k_EDOTA_CW_Regret",
	70: "k_EDOTA_CW_Relax",
	71: "k_EDOTA_CW_MissingHero",
	72: "k_EDOTA_CW_ReturnedHero",
}
var EDOTAChatWheelMessage_value = map[string]int32{
	"k_EDOTA_CW_Ok":                  0,
	"k_EDOTA_CW_Care":                1,
	"k_EDOTA_CW_GetBack":             2,
	"k_EDOTA_CW_NeedWards":           3,
	"k_EDOTA_CW_Stun":                4,
	"k_EDOTA_CW_Help":                5,
	"k_EDOTA_CW_Push":                6,
	"k_EDOTA_CW_GoodJob":             7,
	"k_EDOTA_CW_Missing":             8,
	"k_EDOTA_CW_Missing_Top":         9,
	"k_EDOTA_CW_Missing_Mid":         10,
	"k_EDOTA_CW_Missing_Bottom":      11,
	"k_EDOTA_CW_Go":                  12,
	"k_EDOTA_CW_Initiate":            13,
	"k_EDOTA_CW_Follow":              14,
	"k_EDOTA_CW_Group_Up":            15,
	"k_EDOTA_CW_Spread_Out":          16,
	"k_EDOTA_CW_Split_Farm":          17,
	"k_EDOTA_CW_Attack":              18,
	"k_EDOTA_CW_BRB":                 19,
	"k_EDOTA_CW_Dive":                20,
	"k_EDOTA_CW_OMW":                 21,
	"k_EDOTA_CW_Get_Ready":           22,
	"k_EDOTA_CW_Bait":                23,
	"k_EDOTA_CW_Heal":                24,
	"k_EDOTA_CW_Mana":                25,
	"k_EDOTA_CW_OOM":                 26,
	"k_EDOTA_CW_Skill_Cooldown":      27,
	"k_EDOTA_CW_Ulti_Ready":          28,
	"k_EDOTA_CW_Enemy_Returned":      29,
	"k_EDOTA_CW_All_Missing":         30,
	"k_EDOTA_CW_Enemy_Incoming":      31,
	"k_EDOTA_CW_Invis_Enemy":         32,
	"k_EDOTA_CW_Enemy_Had_Rune":      33,
	"k_EDOTA_CW_Split_Push":          34,
	"k_EDOTA_CW_Coming_To_Gank":      35,
	"k_EDOTA_CW_Request_Gank":        36,
	"k_EDOTA_CW_Fight_Under_Tower":   37,
	"k_EDOTA_CW_Deny_Tower":          38,
	"k_EDOTA_CW_Buy_Courier":         39,
	"k_EDOTA_CW_Upgrade_Courier":     40,
	"k_EDOTA_CW_Need_Detection":      41,
	"k_EDOTA_CW_They_Have_Detection": 42,
	"k_EDOTA_CW_Buy_TP":              43,
	"k_EDOTA_CW_Reuse_Courier":       44,
	"k_EDOTA_CW_Deward":              45,
	"k_EDOTA_CW_Building_Mek":        46,
	"k_EDOTA_CW_Building_Pipe":       47,
	"k_EDOTA_CW_Stack_And_Pull":      48,
	"k_EDOTA_CW_Pull":                49,
	"k_EDOTA_CW_Pulling":             50,
	"k_EDOTA_CW_Stack":               51,
	"k_EDOTA_CW_Jungling":            52,
	"k_EDOTA_CW_Roshan":              53,
	"k_EDOTA_CW_Affirmative":         54,
	"k_EDOTA_CW_Wait":                55,
	"k_EDOTA_CW_Pause":               56,
	"k_EDOTA_CW_Current_Time":        57,
	"k_EDOTA_CW_Check_Runes":         58,
	"k_EDOTA_CW_Smoke_Gank":          59,
	"k_EDOTA_CW_GLHF":                60,
	"k_EDOTA_CW_Nice":                61,
	"k_EDOTA_CW_Thanks":              62,
	"k_EDOTA_CW_Sorry":               63,
	"k_EDOTA_CW_No_Give_Up":          64,
	"k_EDOTA_CW_Just_Happened":       65,
	"k_EDOTA_CW_Game_Is_Hard":        66,
	"k_EDOTA_CW_New_Meta":            67,
	"k_EDOTA_CW_My_Bad":              68,
	"k_EDOTA_CW_Regret":              69,
	"k_EDOTA_CW_Relax":               70,
	"k_EDOTA_CW_MissingHero":         71,
	"k_EDOTA_CW_ReturnedHero":        72,
}

func (x EDOTAChatWheelMessage) Enum() *EDOTAChatWheelMessage {
	p := new(EDOTAChatWheelMessage)
	*p = x
	return p
}
func (x EDOTAChatWheelMessage) String() string {
	return proto.EnumName(EDOTAChatWheelMessage_name, int32(x))
}
func (x *EDOTAChatWheelMessage) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EDOTAChatWheelMessage_value, data, "EDOTAChatWheelMessage")
	if err != nil {
		return err
	}
	*x = EDOTAChatWheelMessage(value)
	return nil
}

type EDOTAStatPopupTypes int32

const (
	EDOTAStatPopupTypes_k_EDOTA_SPT_Textline EDOTAStatPopupTypes = 0
	EDOTAStatPopupTypes_k_EDOTA_SPT_Basic    EDOTAStatPopupTypes = 1
	EDOTAStatPopupTypes_k_EDOTA_SPT_Poll     EDOTAStatPopupTypes = 2
	EDOTAStatPopupTypes_k_EDOTA_SPT_Grid     EDOTAStatPopupTypes = 3
)

var EDOTAStatPopupTypes_name = map[int32]string{
	0: "k_EDOTA_SPT_Textline",
	1: "k_EDOTA_SPT_Basic",
	2: "k_EDOTA_SPT_Poll",
	3: "k_EDOTA_SPT_Grid",
}
var EDOTAStatPopupTypes_value = map[string]int32{
	"k_EDOTA_SPT_Textline": 0,
	"k_EDOTA_SPT_Basic":    1,
	"k_EDOTA_SPT_Poll":     2,
	"k_EDOTA_SPT_Grid":     3,
}

func (x EDOTAStatPopupTypes) Enum() *EDOTAStatPopupTypes {
	p := new(EDOTAStatPopupTypes)
	*p = x
	return p
}
func (x EDOTAStatPopupTypes) String() string {
	return proto.EnumName(EDOTAStatPopupTypes_name, int32(x))
}
func (x *EDOTAStatPopupTypes) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EDOTAStatPopupTypes_value, data, "EDOTAStatPopupTypes")
	if err != nil {
		return err
	}
	*x = EDOTAStatPopupTypes(value)
	return nil
}

type CDOTAMsg_LocationPing struct {
	X                *int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                *int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Target           *int32 `protobuf:"varint,3,opt,name=target" json:"target,omitempty"`
	DirectPing       *bool  `protobuf:"varint,4,opt,name=direct_ping" json:"direct_ping,omitempty"`
	Type             *int32 `protobuf:"varint,5,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAMsg_LocationPing) Reset()         { *m = CDOTAMsg_LocationPing{} }
func (m *CDOTAMsg_LocationPing) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_LocationPing) ProtoMessage()    {}

func (m *CDOTAMsg_LocationPing) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_LocationPing) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_LocationPing) GetTarget() int32 {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return 0
}

func (m *CDOTAMsg_LocationPing) GetDirectPing() bool {
	if m != nil && m.DirectPing != nil {
		return *m.DirectPing
	}
	return false
}

func (m *CDOTAMsg_LocationPing) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type CDOTAMsg_ItemAlert struct {
	X                *int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                *int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Itemid           *int32 `protobuf:"varint,3,opt,name=itemid" json:"itemid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAMsg_ItemAlert) Reset()         { *m = CDOTAMsg_ItemAlert{} }
func (m *CDOTAMsg_ItemAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_ItemAlert) ProtoMessage()    {}

func (m *CDOTAMsg_ItemAlert) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_ItemAlert) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_ItemAlert) GetItemid() int32 {
	if m != nil && m.Itemid != nil {
		return *m.Itemid
	}
	return 0
}

type CDOTAMsg_EnemyItemAlert struct {
	PlayerId         *uint32 `protobuf:"varint,1,opt,name=player_id" json:"player_id,omitempty"`
	Itemid           *int32  `protobuf:"varint,2,opt,name=itemid" json:"itemid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMsg_EnemyItemAlert) Reset()         { *m = CDOTAMsg_EnemyItemAlert{} }
func (m *CDOTAMsg_EnemyItemAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_EnemyItemAlert) ProtoMessage()    {}

func (m *CDOTAMsg_EnemyItemAlert) GetPlayerId() uint32 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *CDOTAMsg_EnemyItemAlert) GetItemid() int32 {
	if m != nil && m.Itemid != nil {
		return *m.Itemid
	}
	return 0
}

type CDOTAMsg_MapLine struct {
	X                *int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                *int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Initial          *bool  `protobuf:"varint,3,opt,name=initial" json:"initial,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAMsg_MapLine) Reset()         { *m = CDOTAMsg_MapLine{} }
func (m *CDOTAMsg_MapLine) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_MapLine) ProtoMessage()    {}

func (m *CDOTAMsg_MapLine) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_MapLine) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_MapLine) GetInitial() bool {
	if m != nil && m.Initial != nil {
		return *m.Initial
	}
	return false
}

type CDOTAMsg_WorldLine struct {
	X                *int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                *int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Z                *int32 `protobuf:"varint,3,opt,name=z" json:"z,omitempty"`
	Initial          *bool  `protobuf:"varint,4,opt,name=initial" json:"initial,omitempty"`
	End              *bool  `protobuf:"varint,5,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAMsg_WorldLine) Reset()         { *m = CDOTAMsg_WorldLine{} }
func (m *CDOTAMsg_WorldLine) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_WorldLine) ProtoMessage()    {}

func (m *CDOTAMsg_WorldLine) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_WorldLine) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_WorldLine) GetZ() int32 {
	if m != nil && m.Z != nil {
		return *m.Z
	}
	return 0
}

func (m *CDOTAMsg_WorldLine) GetInitial() bool {
	if m != nil && m.Initial != nil {
		return *m.Initial
	}
	return false
}

func (m *CDOTAMsg_WorldLine) GetEnd() bool {
	if m != nil && m.End != nil {
		return *m.End
	}
	return false
}

type CDOTAMsg_SendStatPopup struct {
	Style            *EDOTAStatPopupTypes `protobuf:"varint,1,opt,name=style,enum=dota.EDOTAStatPopupTypes,def=0" json:"style,omitempty"`
	StatStrings      []string             `protobuf:"bytes,2,rep,name=stat_strings" json:"stat_strings,omitempty"`
	StatImages       []int32              `protobuf:"varint,3,rep,name=stat_images" json:"stat_images,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *CDOTAMsg_SendStatPopup) Reset()         { *m = CDOTAMsg_SendStatPopup{} }
func (m *CDOTAMsg_SendStatPopup) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_SendStatPopup) ProtoMessage()    {}

const Default_CDOTAMsg_SendStatPopup_Style EDOTAStatPopupTypes = EDOTAStatPopupTypes_k_EDOTA_SPT_Textline

func (m *CDOTAMsg_SendStatPopup) GetStyle() EDOTAStatPopupTypes {
	if m != nil && m.Style != nil {
		return *m.Style
	}
	return Default_CDOTAMsg_SendStatPopup_Style
}

func (m *CDOTAMsg_SendStatPopup) GetStatStrings() []string {
	if m != nil {
		return m.StatStrings
	}
	return nil
}

func (m *CDOTAMsg_SendStatPopup) GetStatImages() []int32 {
	if m != nil {
		return m.StatImages
	}
	return nil
}

type CDOTAMsg_CoachHUDPing struct {
	X                *uint32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                *uint32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Tgtpath          *string `protobuf:"bytes,3,opt,name=tgtpath" json:"tgtpath,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMsg_CoachHUDPing) Reset()         { *m = CDOTAMsg_CoachHUDPing{} }
func (m *CDOTAMsg_CoachHUDPing) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_CoachHUDPing) ProtoMessage()    {}

func (m *CDOTAMsg_CoachHUDPing) GetX() uint32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_CoachHUDPing) GetY() uint32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_CoachHUDPing) GetTgtpath() string {
	if m != nil && m.Tgtpath != nil {
		return *m.Tgtpath
	}
	return ""
}

func init() {
	proto.RegisterEnum("dota.EDOTAChatWheelMessage", EDOTAChatWheelMessage_name, EDOTAChatWheelMessage_value)
	proto.RegisterEnum("dota.EDOTAStatPopupTypes", EDOTAStatPopupTypes_name, EDOTAStatPopupTypes_value)
}
