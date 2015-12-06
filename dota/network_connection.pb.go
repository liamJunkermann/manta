// Code generated by protoc-gen-go.
// source: network_connection.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import math "math"
import google_protobuf "github.com/dotabuff/manta/dota/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ENetworkDisconnectionReason int32

const (
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_INVALID                     ENetworkDisconnectionReason = 0
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SHUTDOWN                    ENetworkDisconnectionReason = 1
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DISCONNECT_BY_USER          ENetworkDisconnectionReason = 2
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DISCONNECT_BY_SERVER        ENetworkDisconnectionReason = 3
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOST                        ENetworkDisconnectionReason = 4
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_OVERFLOW                    ENetworkDisconnectionReason = 5
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_BANNED                ENetworkDisconnectionReason = 6
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_INUSE                 ENetworkDisconnectionReason = 7
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_TICKET                ENetworkDisconnectionReason = 8
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_LOGON                 ENetworkDisconnectionReason = 9
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_AUTHCANCELLED         ENetworkDisconnectionReason = 10
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_AUTHALREADYUSED       ENetworkDisconnectionReason = 11
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_AUTHINVALID           ENetworkDisconnectionReason = 12
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_VACBANSTATE           ENetworkDisconnectionReason = 13
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_LOGGED_IN_ELSEWHERE   ENetworkDisconnectionReason = 14
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_VAC_CHECK_TIMEDOUT    ENetworkDisconnectionReason = 15
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_DROPPED               ENetworkDisconnectionReason = 16
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_OWNERSHIP             ENetworkDisconnectionReason = 17
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVERINFO_OVERFLOW         ENetworkDisconnectionReason = 18
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_TICKMSG_OVERFLOW            ENetworkDisconnectionReason = 19
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STRINGTABLEMSG_OVERFLOW     ENetworkDisconnectionReason = 20
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DELTAENTMSG_OVERFLOW        ENetworkDisconnectionReason = 21
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_TEMPENTMSG_OVERFLOW         ENetworkDisconnectionReason = 22
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SOUNDSMSG_OVERFLOW          ENetworkDisconnectionReason = 23
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SNAPSHOTOVERFLOW            ENetworkDisconnectionReason = 24
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SNAPSHOTERROR               ENetworkDisconnectionReason = 25
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_RELIABLEOVERFLOW            ENetworkDisconnectionReason = 26
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BADDELTATICK                ENetworkDisconnectionReason = 27
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_NOMORESPLITS                ENetworkDisconnectionReason = 28
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_TIMEDOUT                    ENetworkDisconnectionReason = 29
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DISCONNECTED                ENetworkDisconnectionReason = 30
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LEAVINGSPLIT                ENetworkDisconnectionReason = 31
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DIFFERENTCLASSTABLES        ENetworkDisconnectionReason = 32
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BADRELAYPASSWORD            ENetworkDisconnectionReason = 33
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BADSPECTATORPASSWORD        ENetworkDisconnectionReason = 34
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVRESTRICTED              ENetworkDisconnectionReason = 35
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_NOSPECTATORS                ENetworkDisconnectionReason = 36
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVUNAVAILABLE             ENetworkDisconnectionReason = 37
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVSTOP                    ENetworkDisconnectionReason = 38
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_KICKED                      ENetworkDisconnectionReason = 39
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BANADDED                    ENetworkDisconnectionReason = 40
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_KICKBANADDED                ENetworkDisconnectionReason = 41
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVDIRECT                  ENetworkDisconnectionReason = 42
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_PURESERVER_CLIENTEXTRA      ENetworkDisconnectionReason = 43
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_PURESERVER_MISMATCH         ENetworkDisconnectionReason = 44
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_USERCMD                     ENetworkDisconnectionReason = 45
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REJECTED_BY_GAME            ENetworkDisconnectionReason = 46
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_MESSAGE_PARSE_ERROR         ENetworkDisconnectionReason = 47
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_INVALID_MESSAGE_ERROR       ENetworkDisconnectionReason = 48
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BAD_SERVER_PASSWORD         ENetworkDisconnectionReason = 49
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DIRECT_CONNECT_RESERVATION  ENetworkDisconnectionReason = 50
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CONNECTION_FAILURE          ENetworkDisconnectionReason = 51
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_NO_PEER_GROUP_HANDLERS      ENetworkDisconnectionReason = 52
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_RECONNECTION                ENetworkDisconnectionReason = 53
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOOPSHUTDOWN                ENetworkDisconnectionReason = 54
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOOPDEACTIVATE              ENetworkDisconnectionReason = 55
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HOST_ENDGAME                ENetworkDisconnectionReason = 56
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOOP_LEVELLOAD_ACTIVATE     ENetworkDisconnectionReason = 57
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CREATE_SERVER_FAILED        ENetworkDisconnectionReason = 58
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_EXITING                     ENetworkDisconnectionReason = 59
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REQUEST_HOSTSTATE_IDLE      ENetworkDisconnectionReason = 60
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REQUEST_HOSTSTATE_HLTVRELAY ENetworkDisconnectionReason = 61
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_CONSISTENCY_FAIL     ENetworkDisconnectionReason = 62
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_UNABLE_TO_CRC_MAP    ENetworkDisconnectionReason = 63
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_NO_MAP               ENetworkDisconnectionReason = 64
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_DIFFERENT_MAP        ENetworkDisconnectionReason = 65
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVER_REQUIRES_STEAM       ENetworkDisconnectionReason = 66
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_DENY_MISC             ENetworkDisconnectionReason = 67
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_DENY_BAD_ANTI_CHEAT   ENetworkDisconnectionReason = 68
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVER_SHUTDOWN             ENetworkDisconnectionReason = 69
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SPLITPACKET_SEND_OVERFLOW   ENetworkDisconnectionReason = 70
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REPLAY_INCOMPATIBLE         ENetworkDisconnectionReason = 71
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CONNECT_REQUEST_TIMEDOUT    ENetworkDisconnectionReason = 72
)

var ENetworkDisconnectionReason_name = map[int32]string{
	0:  "NETWORK_DISCONNECT_INVALID",
	1:  "NETWORK_DISCONNECT_SHUTDOWN",
	2:  "NETWORK_DISCONNECT_DISCONNECT_BY_USER",
	3:  "NETWORK_DISCONNECT_DISCONNECT_BY_SERVER",
	4:  "NETWORK_DISCONNECT_LOST",
	5:  "NETWORK_DISCONNECT_OVERFLOW",
	6:  "NETWORK_DISCONNECT_STEAM_BANNED",
	7:  "NETWORK_DISCONNECT_STEAM_INUSE",
	8:  "NETWORK_DISCONNECT_STEAM_TICKET",
	9:  "NETWORK_DISCONNECT_STEAM_LOGON",
	10: "NETWORK_DISCONNECT_STEAM_AUTHCANCELLED",
	11: "NETWORK_DISCONNECT_STEAM_AUTHALREADYUSED",
	12: "NETWORK_DISCONNECT_STEAM_AUTHINVALID",
	13: "NETWORK_DISCONNECT_STEAM_VACBANSTATE",
	14: "NETWORK_DISCONNECT_STEAM_LOGGED_IN_ELSEWHERE",
	15: "NETWORK_DISCONNECT_STEAM_VAC_CHECK_TIMEDOUT",
	16: "NETWORK_DISCONNECT_STEAM_DROPPED",
	17: "NETWORK_DISCONNECT_STEAM_OWNERSHIP",
	18: "NETWORK_DISCONNECT_SERVERINFO_OVERFLOW",
	19: "NETWORK_DISCONNECT_TICKMSG_OVERFLOW",
	20: "NETWORK_DISCONNECT_STRINGTABLEMSG_OVERFLOW",
	21: "NETWORK_DISCONNECT_DELTAENTMSG_OVERFLOW",
	22: "NETWORK_DISCONNECT_TEMPENTMSG_OVERFLOW",
	23: "NETWORK_DISCONNECT_SOUNDSMSG_OVERFLOW",
	24: "NETWORK_DISCONNECT_SNAPSHOTOVERFLOW",
	25: "NETWORK_DISCONNECT_SNAPSHOTERROR",
	26: "NETWORK_DISCONNECT_RELIABLEOVERFLOW",
	27: "NETWORK_DISCONNECT_BADDELTATICK",
	28: "NETWORK_DISCONNECT_NOMORESPLITS",
	29: "NETWORK_DISCONNECT_TIMEDOUT",
	30: "NETWORK_DISCONNECT_DISCONNECTED",
	31: "NETWORK_DISCONNECT_LEAVINGSPLIT",
	32: "NETWORK_DISCONNECT_DIFFERENTCLASSTABLES",
	33: "NETWORK_DISCONNECT_BADRELAYPASSWORD",
	34: "NETWORK_DISCONNECT_BADSPECTATORPASSWORD",
	35: "NETWORK_DISCONNECT_HLTVRESTRICTED",
	36: "NETWORK_DISCONNECT_NOSPECTATORS",
	37: "NETWORK_DISCONNECT_HLTVUNAVAILABLE",
	38: "NETWORK_DISCONNECT_HLTVSTOP",
	39: "NETWORK_DISCONNECT_KICKED",
	40: "NETWORK_DISCONNECT_BANADDED",
	41: "NETWORK_DISCONNECT_KICKBANADDED",
	42: "NETWORK_DISCONNECT_HLTVDIRECT",
	43: "NETWORK_DISCONNECT_PURESERVER_CLIENTEXTRA",
	44: "NETWORK_DISCONNECT_PURESERVER_MISMATCH",
	45: "NETWORK_DISCONNECT_USERCMD",
	46: "NETWORK_DISCONNECT_REJECTED_BY_GAME",
	47: "NETWORK_DISCONNECT_MESSAGE_PARSE_ERROR",
	48: "NETWORK_DISCONNECT_INVALID_MESSAGE_ERROR",
	49: "NETWORK_DISCONNECT_BAD_SERVER_PASSWORD",
	50: "NETWORK_DISCONNECT_DIRECT_CONNECT_RESERVATION",
	51: "NETWORK_DISCONNECT_CONNECTION_FAILURE",
	52: "NETWORK_DISCONNECT_NO_PEER_GROUP_HANDLERS",
	53: "NETWORK_DISCONNECT_RECONNECTION",
	54: "NETWORK_DISCONNECT_LOOPSHUTDOWN",
	55: "NETWORK_DISCONNECT_LOOPDEACTIVATE",
	56: "NETWORK_DISCONNECT_HOST_ENDGAME",
	57: "NETWORK_DISCONNECT_LOOP_LEVELLOAD_ACTIVATE",
	58: "NETWORK_DISCONNECT_CREATE_SERVER_FAILED",
	59: "NETWORK_DISCONNECT_EXITING",
	60: "NETWORK_DISCONNECT_REQUEST_HOSTSTATE_IDLE",
	61: "NETWORK_DISCONNECT_REQUEST_HOSTSTATE_HLTVRELAY",
	62: "NETWORK_DISCONNECT_CLIENT_CONSISTENCY_FAIL",
	63: "NETWORK_DISCONNECT_CLIENT_UNABLE_TO_CRC_MAP",
	64: "NETWORK_DISCONNECT_CLIENT_NO_MAP",
	65: "NETWORK_DISCONNECT_CLIENT_DIFFERENT_MAP",
	66: "NETWORK_DISCONNECT_SERVER_REQUIRES_STEAM",
	67: "NETWORK_DISCONNECT_STEAM_DENY_MISC",
	68: "NETWORK_DISCONNECT_STEAM_DENY_BAD_ANTI_CHEAT",
	69: "NETWORK_DISCONNECT_SERVER_SHUTDOWN",
	70: "NETWORK_DISCONNECT_SPLITPACKET_SEND_OVERFLOW",
	71: "NETWORK_DISCONNECT_REPLAY_INCOMPATIBLE",
	72: "NETWORK_DISCONNECT_CONNECT_REQUEST_TIMEDOUT",
}
var ENetworkDisconnectionReason_value = map[string]int32{
	"NETWORK_DISCONNECT_INVALID":                     0,
	"NETWORK_DISCONNECT_SHUTDOWN":                    1,
	"NETWORK_DISCONNECT_DISCONNECT_BY_USER":          2,
	"NETWORK_DISCONNECT_DISCONNECT_BY_SERVER":        3,
	"NETWORK_DISCONNECT_LOST":                        4,
	"NETWORK_DISCONNECT_OVERFLOW":                    5,
	"NETWORK_DISCONNECT_STEAM_BANNED":                6,
	"NETWORK_DISCONNECT_STEAM_INUSE":                 7,
	"NETWORK_DISCONNECT_STEAM_TICKET":                8,
	"NETWORK_DISCONNECT_STEAM_LOGON":                 9,
	"NETWORK_DISCONNECT_STEAM_AUTHCANCELLED":         10,
	"NETWORK_DISCONNECT_STEAM_AUTHALREADYUSED":       11,
	"NETWORK_DISCONNECT_STEAM_AUTHINVALID":           12,
	"NETWORK_DISCONNECT_STEAM_VACBANSTATE":           13,
	"NETWORK_DISCONNECT_STEAM_LOGGED_IN_ELSEWHERE":   14,
	"NETWORK_DISCONNECT_STEAM_VAC_CHECK_TIMEDOUT":    15,
	"NETWORK_DISCONNECT_STEAM_DROPPED":               16,
	"NETWORK_DISCONNECT_STEAM_OWNERSHIP":             17,
	"NETWORK_DISCONNECT_SERVERINFO_OVERFLOW":         18,
	"NETWORK_DISCONNECT_TICKMSG_OVERFLOW":            19,
	"NETWORK_DISCONNECT_STRINGTABLEMSG_OVERFLOW":     20,
	"NETWORK_DISCONNECT_DELTAENTMSG_OVERFLOW":        21,
	"NETWORK_DISCONNECT_TEMPENTMSG_OVERFLOW":         22,
	"NETWORK_DISCONNECT_SOUNDSMSG_OVERFLOW":          23,
	"NETWORK_DISCONNECT_SNAPSHOTOVERFLOW":            24,
	"NETWORK_DISCONNECT_SNAPSHOTERROR":               25,
	"NETWORK_DISCONNECT_RELIABLEOVERFLOW":            26,
	"NETWORK_DISCONNECT_BADDELTATICK":                27,
	"NETWORK_DISCONNECT_NOMORESPLITS":                28,
	"NETWORK_DISCONNECT_TIMEDOUT":                    29,
	"NETWORK_DISCONNECT_DISCONNECTED":                30,
	"NETWORK_DISCONNECT_LEAVINGSPLIT":                31,
	"NETWORK_DISCONNECT_DIFFERENTCLASSTABLES":        32,
	"NETWORK_DISCONNECT_BADRELAYPASSWORD":            33,
	"NETWORK_DISCONNECT_BADSPECTATORPASSWORD":        34,
	"NETWORK_DISCONNECT_HLTVRESTRICTED":              35,
	"NETWORK_DISCONNECT_NOSPECTATORS":                36,
	"NETWORK_DISCONNECT_HLTVUNAVAILABLE":             37,
	"NETWORK_DISCONNECT_HLTVSTOP":                    38,
	"NETWORK_DISCONNECT_KICKED":                      39,
	"NETWORK_DISCONNECT_BANADDED":                    40,
	"NETWORK_DISCONNECT_KICKBANADDED":                41,
	"NETWORK_DISCONNECT_HLTVDIRECT":                  42,
	"NETWORK_DISCONNECT_PURESERVER_CLIENTEXTRA":      43,
	"NETWORK_DISCONNECT_PURESERVER_MISMATCH":         44,
	"NETWORK_DISCONNECT_USERCMD":                     45,
	"NETWORK_DISCONNECT_REJECTED_BY_GAME":            46,
	"NETWORK_DISCONNECT_MESSAGE_PARSE_ERROR":         47,
	"NETWORK_DISCONNECT_INVALID_MESSAGE_ERROR":       48,
	"NETWORK_DISCONNECT_BAD_SERVER_PASSWORD":         49,
	"NETWORK_DISCONNECT_DIRECT_CONNECT_RESERVATION":  50,
	"NETWORK_DISCONNECT_CONNECTION_FAILURE":          51,
	"NETWORK_DISCONNECT_NO_PEER_GROUP_HANDLERS":      52,
	"NETWORK_DISCONNECT_RECONNECTION":                53,
	"NETWORK_DISCONNECT_LOOPSHUTDOWN":                54,
	"NETWORK_DISCONNECT_LOOPDEACTIVATE":              55,
	"NETWORK_DISCONNECT_HOST_ENDGAME":                56,
	"NETWORK_DISCONNECT_LOOP_LEVELLOAD_ACTIVATE":     57,
	"NETWORK_DISCONNECT_CREATE_SERVER_FAILED":        58,
	"NETWORK_DISCONNECT_EXITING":                     59,
	"NETWORK_DISCONNECT_REQUEST_HOSTSTATE_IDLE":      60,
	"NETWORK_DISCONNECT_REQUEST_HOSTSTATE_HLTVRELAY": 61,
	"NETWORK_DISCONNECT_CLIENT_CONSISTENCY_FAIL":     62,
	"NETWORK_DISCONNECT_CLIENT_UNABLE_TO_CRC_MAP":    63,
	"NETWORK_DISCONNECT_CLIENT_NO_MAP":               64,
	"NETWORK_DISCONNECT_CLIENT_DIFFERENT_MAP":        65,
	"NETWORK_DISCONNECT_SERVER_REQUIRES_STEAM":       66,
	"NETWORK_DISCONNECT_STEAM_DENY_MISC":             67,
	"NETWORK_DISCONNECT_STEAM_DENY_BAD_ANTI_CHEAT":   68,
	"NETWORK_DISCONNECT_SERVER_SHUTDOWN":             69,
	"NETWORK_DISCONNECT_SPLITPACKET_SEND_OVERFLOW":   70,
	"NETWORK_DISCONNECT_REPLAY_INCOMPATIBLE":         71,
	"NETWORK_DISCONNECT_CONNECT_REQUEST_TIMEDOUT":    72,
}

func (x ENetworkDisconnectionReason) Enum() *ENetworkDisconnectionReason {
	p := new(ENetworkDisconnectionReason)
	*p = x
	return p
}
func (x ENetworkDisconnectionReason) String() string {
	return proto.EnumName(ENetworkDisconnectionReason_name, int32(x))
}
func (x *ENetworkDisconnectionReason) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ENetworkDisconnectionReason_value, data, "ENetworkDisconnectionReason")
	if err != nil {
		return err
	}
	*x = ENetworkDisconnectionReason(value)
	return nil
}

var E_NetworkConnectionToken = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50500,
	Name:          "dota.network_connection_token",
	Tag:           "bytes,50500,opt,name=network_connection_token",
}

func init() {
	proto.RegisterEnum("dota.ENetworkDisconnectionReason", ENetworkDisconnectionReason_name, ENetworkDisconnectionReason_value)
	proto.RegisterExtension(E_NetworkConnectionToken)
}
