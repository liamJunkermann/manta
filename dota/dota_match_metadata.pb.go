// Code generated by protoc-gen-go.
// source: dota_match_metadata.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CDOTAMatchMetadataFile struct {
	Version          *int32              `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	MatchId          *uint64             `protobuf:"varint,2,req,name=match_id" json:"match_id,omitempty"`
	Metadata         *CDOTAMatchMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	PrivateMetadata  []byte              `protobuf:"bytes,4,opt,name=private_metadata" json:"private_metadata,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *CDOTAMatchMetadataFile) Reset()                    { *m = CDOTAMatchMetadataFile{} }
func (m *CDOTAMatchMetadataFile) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadataFile) ProtoMessage()               {}
func (*CDOTAMatchMetadataFile) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *CDOTAMatchMetadataFile) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *CDOTAMatchMetadataFile) GetMatchId() uint64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CDOTAMatchMetadataFile) GetMetadata() *CDOTAMatchMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CDOTAMatchMetadataFile) GetPrivateMetadata() []byte {
	if m != nil {
		return m.PrivateMetadata
	}
	return nil
}

type CDOTAMatchMetadata struct {
	Teams            []*CDOTAMatchMetadata_Team  `protobuf:"bytes,1,rep,name=teams" json:"teams,omitempty"`
	ItemRewards      []*CLobbyTimedRewardDetails `protobuf:"bytes,2,rep,name=item_rewards" json:"item_rewards,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *CDOTAMatchMetadata) Reset()                    { *m = CDOTAMatchMetadata{} }
func (m *CDOTAMatchMetadata) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata) ProtoMessage()               {}
func (*CDOTAMatchMetadata) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *CDOTAMatchMetadata) GetTeams() []*CDOTAMatchMetadata_Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

func (m *CDOTAMatchMetadata) GetItemRewards() []*CLobbyTimedRewardDetails {
	if m != nil {
		return m.ItemRewards
	}
	return nil
}

type CDOTAMatchMetadata_Team struct {
	DotaTeam          *uint32                           `protobuf:"varint,1,opt,name=dota_team" json:"dota_team,omitempty"`
	Players           []*CDOTAMatchMetadata_Team_Player `protobuf:"bytes,2,rep,name=players" json:"players,omitempty"`
	GraphExperience   []float32                         `protobuf:"fixed32,3,rep,name=graph_experience" json:"graph_experience,omitempty"`
	GraphGoldEarned   []float32                         `protobuf:"fixed32,4,rep,name=graph_gold_earned" json:"graph_gold_earned,omitempty"`
	GraphNetWorth     []float32                         `protobuf:"fixed32,5,rep,name=graph_net_worth" json:"graph_net_worth,omitempty"`
	CmFirstPick       *bool                             `protobuf:"varint,6,opt,name=cm_first_pick" json:"cm_first_pick,omitempty"`
	CmCaptainPlayerId *uint32                           `protobuf:"varint,7,opt,name=cm_captain_player_id" json:"cm_captain_player_id,omitempty"`
	CmBans            []uint32                          `protobuf:"varint,8,rep,name=cm_bans" json:"cm_bans,omitempty"`
	CmPicks           []uint32                          `protobuf:"varint,9,rep,name=cm_picks" json:"cm_picks,omitempty"`
	CmPenalty         *uint32                           `protobuf:"varint,10,opt,name=cm_penalty" json:"cm_penalty,omitempty"`
	XXX_unrecognized  []byte                            `json:"-"`
}

func (m *CDOTAMatchMetadata_Team) Reset()                    { *m = CDOTAMatchMetadata_Team{} }
func (m *CDOTAMatchMetadata_Team) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team) ProtoMessage()               {}
func (*CDOTAMatchMetadata_Team) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1, 0} }

func (m *CDOTAMatchMetadata_Team) GetDotaTeam() uint32 {
	if m != nil && m.DotaTeam != nil {
		return *m.DotaTeam
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team) GetPlayers() []*CDOTAMatchMetadata_Team_Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphExperience() []float32 {
	if m != nil {
		return m.GraphExperience
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphGoldEarned() []float32 {
	if m != nil {
		return m.GraphGoldEarned
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphNetWorth() []float32 {
	if m != nil {
		return m.GraphNetWorth
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmFirstPick() bool {
	if m != nil && m.CmFirstPick != nil {
		return *m.CmFirstPick
	}
	return false
}

func (m *CDOTAMatchMetadata_Team) GetCmCaptainPlayerId() uint32 {
	if m != nil && m.CmCaptainPlayerId != nil {
		return *m.CmCaptainPlayerId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team) GetCmBans() []uint32 {
	if m != nil {
		return m.CmBans
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmPicks() []uint32 {
	if m != nil {
		return m.CmPicks
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmPenalty() uint32 {
	if m != nil && m.CmPenalty != nil {
		return *m.CmPenalty
	}
	return 0
}

type CDOTAMatchMetadata_Team_PlayerKill struct {
	VictimSlot       *uint32 `protobuf:"varint,1,opt,name=victim_slot" json:"victim_slot,omitempty"`
	Count            *uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) Reset()         { *m = CDOTAMatchMetadata_Team_PlayerKill{} }
func (m *CDOTAMatchMetadata_Team_PlayerKill) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_PlayerKill) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_PlayerKill) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{1, 0, 0}
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) GetVictimSlot() uint32 {
	if m != nil && m.VictimSlot != nil {
		return *m.VictimSlot
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type CDOTAMatchMetadata_Team_ItemPurchase struct {
	ItemId           *uint32 `protobuf:"varint,1,opt,name=item_id" json:"item_id,omitempty"`
	PurchaseTime     *uint32 `protobuf:"varint,2,opt,name=purchase_time" json:"purchase_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_ItemPurchase) Reset()         { *m = CDOTAMatchMetadata_Team_ItemPurchase{} }
func (m *CDOTAMatchMetadata_Team_ItemPurchase) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_ItemPurchase) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_ItemPurchase) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{1, 0, 1}
}

func (m *CDOTAMatchMetadata_Team_ItemPurchase) GetItemId() uint32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_ItemPurchase) GetPurchaseTime() uint32 {
	if m != nil && m.PurchaseTime != nil {
		return *m.PurchaseTime
	}
	return 0
}

type CDOTAMatchMetadata_Team_Player struct {
	AccountId         *uint32                                 `protobuf:"varint,1,opt,name=account_id" json:"account_id,omitempty"`
	AbilityUpgrades   []uint32                                `protobuf:"varint,2,rep,name=ability_upgrades" json:"ability_upgrades,omitempty"`
	PlayerSlot        *uint32                                 `protobuf:"varint,3,opt,name=player_slot" json:"player_slot,omitempty"`
	EquippedEconItems []*CSOEconItem                          `protobuf:"bytes,4,rep,name=equipped_econ_items" json:"equipped_econ_items,omitempty"`
	Kills             []*CDOTAMatchMetadata_Team_PlayerKill   `protobuf:"bytes,5,rep,name=kills" json:"kills,omitempty"`
	Items             []*CDOTAMatchMetadata_Team_ItemPurchase `protobuf:"bytes,6,rep,name=items" json:"items,omitempty"`
	AvgKillsX16       *uint32                                 `protobuf:"varint,7,opt,name=avg_kills_x16" json:"avg_kills_x16,omitempty"`
	AvgDeathsX16      *uint32                                 `protobuf:"varint,8,opt,name=avg_deaths_x16" json:"avg_deaths_x16,omitempty"`
	AvgAssistsX16     *uint32                                 `protobuf:"varint,9,opt,name=avg_assists_x16" json:"avg_assists_x16,omitempty"`
	AvgGpmX16         *uint32                                 `protobuf:"varint,10,opt,name=avg_gpm_x16" json:"avg_gpm_x16,omitempty"`
	AvgXpmX16         *uint32                                 `protobuf:"varint,11,opt,name=avg_xpm_x16" json:"avg_xpm_x16,omitempty"`
	BestKillsX16      *uint32                                 `protobuf:"varint,12,opt,name=best_kills_x16" json:"best_kills_x16,omitempty"`
	BestAssistsX16    *uint32                                 `protobuf:"varint,13,opt,name=best_assists_x16" json:"best_assists_x16,omitempty"`
	BestGpmX16        *uint32                                 `protobuf:"varint,14,opt,name=best_gpm_x16" json:"best_gpm_x16,omitempty"`
	BestXpmX16        *uint32                                 `protobuf:"varint,15,opt,name=best_xpm_x16" json:"best_xpm_x16,omitempty"`
	WinStreak         *uint32                                 `protobuf:"varint,16,opt,name=win_streak" json:"win_streak,omitempty"`
	BestWinStreak     *uint32                                 `protobuf:"varint,17,opt,name=best_win_streak" json:"best_win_streak,omitempty"`
	FightScore        *float32                                `protobuf:"fixed32,18,opt,name=fight_score" json:"fight_score,omitempty"`
	FarmScore         *float32                                `protobuf:"fixed32,19,opt,name=farm_score" json:"farm_score,omitempty"`
	SupportScore      *float32                                `protobuf:"fixed32,20,opt,name=support_score" json:"support_score,omitempty"`
	PushScore         *float32                                `protobuf:"fixed32,21,opt,name=push_score" json:"push_score,omitempty"`
	XXX_unrecognized  []byte                                  `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_Player) Reset()         { *m = CDOTAMatchMetadata_Team_Player{} }
func (m *CDOTAMatchMetadata_Team_Player) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_Player) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_Player) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{1, 0, 2}
}

func (m *CDOTAMatchMetadata_Team_Player) GetAccountId() uint32 {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAbilityUpgrades() []uint32 {
	if m != nil {
		return m.AbilityUpgrades
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetPlayerSlot() uint32 {
	if m != nil && m.PlayerSlot != nil {
		return *m.PlayerSlot
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetEquippedEconItems() []*CSOEconItem {
	if m != nil {
		return m.EquippedEconItems
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetKills() []*CDOTAMatchMetadata_Team_PlayerKill {
	if m != nil {
		return m.Kills
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetItems() []*CDOTAMatchMetadata_Team_ItemPurchase {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgKillsX16() uint32 {
	if m != nil && m.AvgKillsX16 != nil {
		return *m.AvgKillsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgDeathsX16() uint32 {
	if m != nil && m.AvgDeathsX16 != nil {
		return *m.AvgDeathsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgAssistsX16() uint32 {
	if m != nil && m.AvgAssistsX16 != nil {
		return *m.AvgAssistsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgGpmX16() uint32 {
	if m != nil && m.AvgGpmX16 != nil {
		return *m.AvgGpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgXpmX16() uint32 {
	if m != nil && m.AvgXpmX16 != nil {
		return *m.AvgXpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestKillsX16() uint32 {
	if m != nil && m.BestKillsX16 != nil {
		return *m.BestKillsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestAssistsX16() uint32 {
	if m != nil && m.BestAssistsX16 != nil {
		return *m.BestAssistsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestGpmX16() uint32 {
	if m != nil && m.BestGpmX16 != nil {
		return *m.BestGpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestXpmX16() uint32 {
	if m != nil && m.BestXpmX16 != nil {
		return *m.BestXpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetWinStreak() uint32 {
	if m != nil && m.WinStreak != nil {
		return *m.WinStreak
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestWinStreak() uint32 {
	if m != nil && m.BestWinStreak != nil {
		return *m.BestWinStreak
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetFightScore() float32 {
	if m != nil && m.FightScore != nil {
		return *m.FightScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetFarmScore() float32 {
	if m != nil && m.FarmScore != nil {
		return *m.FarmScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetSupportScore() float32 {
	if m != nil && m.SupportScore != nil {
		return *m.SupportScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetPushScore() float32 {
	if m != nil && m.PushScore != nil {
		return *m.PushScore
	}
	return 0
}

type CDOTAMatchPrivateMetadata struct {
	Kills            []*CDOTAMatchPrivateMetadata_Kill `protobuf:"bytes,1,rep,name=kills" json:"kills,omitempty"`
	TestString       *string                           `protobuf:"bytes,100,opt,name=test_string" json:"test_string,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata) Reset()                    { *m = CDOTAMatchPrivateMetadata{} }
func (m *CDOTAMatchPrivateMetadata) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata) ProtoMessage()               {}
func (*CDOTAMatchPrivateMetadata) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *CDOTAMatchPrivateMetadata) GetKills() []*CDOTAMatchPrivateMetadata_Kill {
	if m != nil {
		return m.Kills
	}
	return nil
}

func (m *CDOTAMatchPrivateMetadata) GetTestString() string {
	if m != nil && m.TestString != nil {
		return *m.TestString
	}
	return ""
}

type CDOTAMatchPrivateMetadata_Kill struct {
	Timestamp        *int32   `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	KillerHeroId     *uint32  `protobuf:"varint,2,opt,name=killer_hero_id" json:"killer_hero_id,omitempty"`
	VictimHeroId     *uint32  `protobuf:"varint,3,opt,name=victim_hero_id" json:"victim_hero_id,omitempty"`
	AssistHeroIds    []uint32 `protobuf:"varint,4,rep,name=assist_hero_ids" json:"assist_hero_ids,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata_Kill) Reset()         { *m = CDOTAMatchPrivateMetadata_Kill{} }
func (m *CDOTAMatchPrivateMetadata_Kill) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata_Kill) ProtoMessage()    {}
func (*CDOTAMatchPrivateMetadata_Kill) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{2, 0}
}

func (m *CDOTAMatchPrivateMetadata_Kill) GetTimestamp() int32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Kill) GetKillerHeroId() uint32 {
	if m != nil && m.KillerHeroId != nil {
		return *m.KillerHeroId
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Kill) GetVictimHeroId() uint32 {
	if m != nil && m.VictimHeroId != nil {
		return *m.VictimHeroId
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Kill) GetAssistHeroIds() []uint32 {
	if m != nil {
		return m.AssistHeroIds
	}
	return nil
}

func init() {
	proto.RegisterType((*CDOTAMatchMetadataFile)(nil), "dota.CDOTAMatchMetadataFile")
	proto.RegisterType((*CDOTAMatchMetadata)(nil), "dota.CDOTAMatchMetadata")
	proto.RegisterType((*CDOTAMatchMetadata_Team)(nil), "dota.CDOTAMatchMetadata.Team")
	proto.RegisterType((*CDOTAMatchMetadata_Team_PlayerKill)(nil), "dota.CDOTAMatchMetadata.Team.PlayerKill")
	proto.RegisterType((*CDOTAMatchMetadata_Team_ItemPurchase)(nil), "dota.CDOTAMatchMetadata.Team.ItemPurchase")
	proto.RegisterType((*CDOTAMatchMetadata_Team_Player)(nil), "dota.CDOTAMatchMetadata.Team.Player")
	proto.RegisterType((*CDOTAMatchPrivateMetadata)(nil), "dota.CDOTAMatchPrivateMetadata")
	proto.RegisterType((*CDOTAMatchPrivateMetadata_Kill)(nil), "dota.CDOTAMatchPrivateMetadata.Kill")
}

var fileDescriptor12 = []byte{
	// 770 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x4f, 0x1b, 0x49,
	0x10, 0x5d, 0x7f, 0x61, 0xbb, 0x6c, 0x03, 0x1e, 0x3e, 0x76, 0xb0, 0xd8, 0x15, 0x42, 0x7b, 0xb0,
	0xd0, 0xca, 0xda, 0x65, 0x37, 0x44, 0x39, 0x26, 0x21, 0x91, 0xa2, 0x04, 0x81, 0x12, 0xee, 0xad,
	0xf6, 0x4c, 0x33, 0xd3, 0x62, 0x3e, 0x3a, 0xdd, 0x6d, 0x83, 0x6f, 0xc9, 0x5f, 0xca, 0x25, 0xd7,
	0x5c, 0xf3, 0xaf, 0x52, 0xfd, 0x31, 0x18, 0x12, 0x82, 0x72, 0x7d, 0xaf, 0xea, 0xd5, 0xeb, 0x57,
	0x35, 0x03, 0x3b, 0x71, 0xa9, 0x29, 0xc9, 0xa9, 0x8e, 0x52, 0x92, 0x33, 0x4d, 0x63, 0xaa, 0xe9,
	0x44, 0xc8, 0x52, 0x97, 0x41, 0xd3, 0x50, 0xa3, 0xad, 0x29, 0x55, 0x8c, 0x24, 0x51, 0xce, 0x94,
	0xa2, 0x09, 0x53, 0x8e, 0x1c, 0xed, 0xda, 0xbe, 0x25, 0x4c, 0xa2, 0x32, 0xcf, 0xcb, 0xc2, 0xb1,
	0xfb, 0x1f, 0x6b, 0xb0, 0xfd, 0xfc, 0xf8, 0xf4, 0xfc, 0xe9, 0x89, 0x11, 0x3e, 0xf1, 0xba, 0x2f,
	0x79, 0xc6, 0x82, 0x35, 0x68, 0xcf, 0x99, 0x54, 0xbc, 0x2c, 0xc2, 0xda, 0x5e, 0x7d, 0xdc, 0x0a,
	0xd6, 0xa1, 0xe3, 0xc6, 0xf3, 0x38, 0xac, 0x23, 0xd2, 0x0c, 0x0e, 0x10, 0xf1, 0x2d, 0x61, 0x63,
	0xaf, 0x36, 0xee, 0x1d, 0x86, 0x13, 0x33, 0x6e, 0xf2, 0xa3, 0x64, 0x10, 0xc2, 0xba, 0x90, 0x7c,
	0x4e, 0x35, 0xbb, 0xb1, 0x1f, 0x36, 0xb1, 0xa7, 0xbf, 0xff, 0xa9, 0x03, 0xc1, 0x3d, 0x0d, 0x7f,
	0x43, 0x4b, 0x33, 0x9a, 0x2b, 0x9c, 0xde, 0x40, 0xe5, 0x3f, 0x7e, 0xa6, 0x3c, 0x39, 0xc7, 0xaa,
	0xe0, 0x7f, 0xe8, 0x73, 0xcd, 0x72, 0x22, 0xd9, 0x15, 0x95, 0xb1, 0x42, 0x83, 0xa6, 0xe9, 0x4f,
	0xdf, 0xf4, 0xa6, 0x9c, 0x4e, 0x17, 0xe7, 0x3c, 0x67, 0xf1, 0x5b, 0xcb, 0x1f, 0x63, 0x2f, 0xcf,
	0xd4, 0xe8, 0x73, 0x1b, 0x9a, 0xb6, 0x7d, 0x08, 0x5d, 0x9b, 0x93, 0x99, 0x88, 0x03, 0x6b, 0xe3,
	0x41, 0xf0, 0x08, 0xda, 0x22, 0xa3, 0x0b, 0x8c, 0xc0, 0x8b, 0xfd, 0xf5, 0xa0, 0x83, 0xc9, 0x99,
	0x2d, 0x36, 0xef, 0x4c, 0x24, 0x15, 0x29, 0x61, 0xd7, 0x82, 0x49, 0xce, 0x8a, 0x88, 0x61, 0x36,
	0x8d, 0x71, 0x3d, 0xd8, 0x81, 0xa1, 0x63, 0x92, 0x32, 0x8b, 0x09, 0xa3, 0xb2, 0x60, 0x31, 0x46,
	0x60, 0xa8, 0xdf, 0x61, 0xcd, 0x51, 0x05, 0xd3, 0xe4, 0xaa, 0x94, 0x3a, 0x0d, 0x5b, 0x96, 0xd8,
	0x82, 0x41, 0x94, 0x93, 0x0b, 0x2e, 0x95, 0x26, 0x82, 0x47, 0x97, 0xe1, 0x0a, 0x7a, 0xeb, 0x04,
	0xbb, 0xb0, 0x89, 0x70, 0x44, 0x05, 0x3e, 0xa3, 0x20, 0xce, 0xa6, 0x59, 0x4b, 0xdb, 0x3a, 0xc7,
	0xcd, 0x21, 0x3b, 0xa5, 0x85, 0x0a, 0x3b, 0xa8, 0x32, 0x30, 0x9b, 0x43, 0xc0, 0xf4, 0xab, 0xb0,
	0x6b, 0x91, 0x00, 0xc0, 0x20, 0xac, 0xa0, 0x99, 0x5e, 0x84, 0x60, 0xda, 0x46, 0xff, 0x00, 0xb8,
	0x37, 0xbc, 0xe6, 0x59, 0x16, 0x6c, 0x40, 0x6f, 0xce, 0x23, 0xcd, 0x73, 0xa2, 0xb2, 0x52, 0xfb,
	0x4c, 0x06, 0xd0, 0x8a, 0xca, 0x59, 0xa1, 0x31, 0x11, 0xd3, 0x71, 0x04, 0xfd, 0x57, 0x18, 0xfa,
	0xd9, 0x4c, 0x46, 0x29, 0x1e, 0x9f, 0x19, 0x6c, 0x97, 0x80, 0x4e, 0x5c, 0x3d, 0xda, 0x17, 0x9e,
	0x24, 0x28, 0xc5, 0x7c, 0xdf, 0x97, 0x26, 0xac, 0xf8, 0xb8, 0xd0, 0x08, 0x8d, 0xac, 0xe6, 0xb2,
	0x0b, 0x23, 0xa4, 0x53, 0x9e, 0x71, 0xbd, 0x20, 0x33, 0x81, 0xb9, 0xc4, 0xcc, 0xad, 0x60, 0x60,
	0x4c, 0xf9, 0xc7, 0x5a, 0x53, 0x0d, 0x5b, 0x3e, 0x81, 0x0d, 0xf6, 0x7e, 0xc6, 0x85, 0x60, 0x98,
	0x6a, 0x54, 0x16, 0xc4, 0x78, 0x50, 0x36, 0xd9, 0xde, 0xe1, 0xd0, 0x2f, 0xed, 0xdd, 0xe9, 0x0b,
	0xa4, 0x8c, 0xdb, 0xe0, 0x31, 0xb4, 0x2e, 0xf1, 0x85, 0xca, 0x46, 0xdc, 0x3b, 0x1c, 0xff, 0xca,
	0x5a, 0x6d, 0x24, 0x4f, 0xa0, 0xe5, 0xa4, 0x57, 0x6c, 0xe3, 0xc1, 0xc3, 0x8d, 0x77, 0x92, 0xc1,
	0x20, 0xe8, 0x3c, 0x21, 0x76, 0x2e, 0xb9, 0xfe, 0xf7, 0xc8, 0x6f, 0x6a, 0x1b, 0x56, 0x0d, 0x1c,
	0x33, 0xaa, 0x53, 0x87, 0x77, 0x2c, 0x8e, 0xf7, 0x60, 0x70, 0xaa, 0x14, 0x57, 0xda, 0x11, 0x5d,
	0x4b, 0x60, 0x00, 0x86, 0x48, 0x44, 0x6e, 0x41, 0xb8, 0x0d, 0x5e, 0x7b, 0xb0, 0x57, 0x49, 0x4f,
	0x19, 0x5e, 0xcd, 0x72, 0x64, 0xbf, 0x0a, 0xd7, 0xe2, 0xb7, 0xb5, 0x07, 0x96, 0xd9, 0x84, 0xbe,
	0x65, 0x2a, 0xf1, 0xd5, 0x3b, 0x68, 0xa5, 0xbe, 0x66, 0x51, 0x5c, 0xdb, 0x15, 0x5e, 0x9e, 0xd2,
	0x92, 0xd1, 0xcb, 0x70, 0xbd, 0x32, 0x6d, 0x2b, 0x6f, 0x11, 0xc3, 0xca, 0xdf, 0x05, 0x4f, 0x52,
	0x4d, 0x54, 0x54, 0x4a, 0x16, 0x06, 0x08, 0xd6, 0x8d, 0xc2, 0x05, 0x95, 0xb9, 0xc7, 0x36, 0x2c,
	0x86, 0x29, 0xa9, 0x99, 0x10, 0x78, 0xff, 0x1e, 0xde, 0xac, 0x4a, 0xc5, 0x4c, 0xa5, 0x1e, 0xdb,
	0x32, 0xd8, 0xfe, 0xd7, 0x1a, 0xec, 0x2c, 0x93, 0x3f, 0x73, 0x7f, 0x96, 0x9b, 0x7f, 0xc7, 0x7f,
	0xd5, 0x8a, 0x6b, 0xf7, 0x7f, 0xb9, 0xdf, 0xd5, 0x4f, 0xaa, 0x8b, 0xd7, 0xc6, 0x3f, 0x7a, 0xe7,
	0x45, 0x12, 0xc6, 0x38, 0xa7, 0x3b, 0x9a, 0x42, 0xd3, 0x92, 0xf8, 0x83, 0x30, 0x07, 0xac, 0x34,
	0xcd, 0x85, 0x3d, 0xd3, 0x96, 0x49, 0xd8, 0x0c, 0xc1, 0x63, 0x4c, 0x99, 0x2c, 0xdd, 0x5f, 0xd1,
	0x27, 0xef, 0xbf, 0x9c, 0x0a, 0x6f, 0xdc, 0x2c, 0xd5, 0x86, 0x5e, 0xe1, 0xee, 0x46, 0x07, 0xcf,
	0x1a, 0x1f, 0x6a, 0xbf, 0x7d, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x4f, 0xab, 0x3a, 0xe0, 0x05,
	0x00, 0x00,
}