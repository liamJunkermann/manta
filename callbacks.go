package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

func (p *Parser) OnCDemoStop(msg Message, obj *dota.CDemoStop) { p.Stop() }

// OK
func (p *Parser) OnCDemoFileHeader(msg Message, obj *dota.CDemoFileHeader) {}
func (p *Parser) OnCDemoSyncTick(msg Message, obj *dota.CDemoSyncTick)     {}
func (p *Parser) OnCDemoConsoleCmd(msg Message, obj *dota.CDemoConsoleCmd) {}
func (p *Parser) OnCDemoClassInfo(msg Message, obj *dota.CDemoClassInfo) {
	// PP(obj.GetClasses())
}
func (p *Parser) OnCDemoStringTables(msg Message, obj *dota.CDemoStringTables) {}
func (p *Parser) OnCDemoSendTables(msg Message, obj *dota.CDemoSendTables) {
	pbuf := proto.NewBuffer(obj.GetData())
	pbuf.DecodeVarint()
	o := dota.CSVCMsg_FlattenedSerializer{}
	err := pbuf.Unmarshal(&o)
	if err != nil {
		panic(err)
	}
	j, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("send_tables.json", j, 0644)
}

var matchBytes = []byte{0x2a, 0xc8, 0x01, 0x42}

func shiftLower(bit byte, b []byte) byte {
	bit = bit << 7
	for i := len(b) - 1; i >= 0; i-- {
		newByte := b[i] >> 1
		newByte |= bit
		bit = (b[i] & 1) << 7
		b[i] = newByte
	}
	return (bit >> 7)
}

func init() {
	return
	m := &dota.CUserMessageSayText2{
		Chat:        proto.Bool(true),
		Entityindex: proto.Uint32(1),
		Messagename: proto.String("DOTA_Chat_All"),
		Param1:      proto.String("manveru"),
		Param2:      proto.String("one"),
		Param3:      proto.String(""),
		Param4:      proto.String(""),
	}

	out, _ := proto.Marshal(m)
	headerType := proto.EncodeVarint(uint64(dota.EBaseUserMessages_UM_SayText2))
	headerLen := proto.EncodeVarint(uint64(len(out)))
	PP(headerType, headerLen, out)
	msg := append(append(headerType, headerLen...), out...)
	PP(out)
	PP(msg)

	ioutil.WriteFile("expected", msg, 0644)
}

// FIXME
func (p *Parser) OnCDemoPacket(msg Message, obj *dota.CDemoPacket) {
	if msg.Tick != 1445 || msg.Size != 173 {
		return
	}

	data := obj.GetData()
	PP(data)
	shiftLower(0, data)
	shiftLower(0, data)
	PP(msg, data)
	o := &dota.CUserMessageSayText2{}
	for n := 0; n <= len(data); n++ {
		for m := 0; m <= len(data); m++ {
			if n+m >= len(data) || m <= n {
				continue
			}
			err := proto.Unmarshal(data[n:m], o)
			if err != nil {
				// spew.Println(err)
				continue
			}
			PP(n, m, data[n:m], o)
		}
	}

	//orig := make([]byte, len(data))
	//copy(orig[:], data[:])
	//shiftLower(0, data)
	//shiftLower(0, data)
	//PP(msg)
	//mustDecode(orig, data, &dota.CUserMessageSayText2{})

	return

	if len(data) != 51 || msg.Tick != 260 {
		return
	}
	PP(msg, data)
	ioutil.WriteFile("actual", data, 0644)
	return

	defer func() {
		if r := recover(); r != nil {
			spew.Println(r)
		}
	}()

	i := 0
	for i < len(data)*8 {
		br := NewBitReader(data)
		for n := 0; n <= i; n++ {
			br.ReadBoolean()
		}

		iLen := br.ReadVarInt()
		iType := br.ReadVarInt()

		ebum := dota.EBaseUserMessages(iType)
		// edum := dota.EDotaUserMessages(iType)
		// edem := dota.EDotaEntityMessages(iType)
		// svcm := dota.SVC_Messages(iType)

		if ebum == dota.EBaseUserMessages_UM_SayText2 {
			PP(iLen, iType)
			/*
				PP(iLen, i/8, i, iType, ebum, edum, edem, svcm)
				buf := br.ReadBytes(int(iLen))
				PP(buf)
				mustDecode(buf, &dota.CUserMessageSayText2{})
			*/
		}

		i++
	}
	return

	/*
		spew.Printf("%v type: %s(%d) len: %d\n", flags, ebum, iType, iLen)

		switch ebum {
		case dota.EBaseUserMessages_UM_SayText2:
			mustDecode(b, &dota.CUserMessageSayText2{})
		case dota.EBaseUserMessages_UM_SendAudio:
			mustDecode(b, &dota.CUserMessageSendAudio{})
		case dota.EBaseUserMessages_UM_DesiredTimescale:
			mustDecode(b, &dota.CUserMessageDesiredTimescale{})
		case dota.EBaseUserMessages_UM_CurrentTimescale:
			mustDecode(b, &dota.CUserMessageCurrentTimescale{})
		default:
			// PP(iType, raw)
			// tryDecodeBytes(raw)
		}
	*/
}

func mustDecode(o, b []byte, msg proto.Message) {
	err := proto.Unmarshal(b, msg)
	if err != nil {
		spew.Println(err)
	}
	PP(o, b, msg)
}

/*
([]uint8) (len=40 cap=40) {
	00000000  d6 95 20 04 40 04 68 34  10 3d 51 05 7d 0d a1 85  |.. .@.h4.=Q.}...|
	00000010  d1 7d 05 b1 b1 89 1c b4  85 b9 d9 95 c9 d5 a9 0c  |.}..............|
	00000020  bc b9 95 c9 00 e8 00 00                           |........|
}

([]uint8) (len=40 cap=40) {
	00000000  d6 95 20 04 40 04 68 34  10 3d 51 05 7d 0d a1 85  |.. .@.h4.=Q.}...|
	00000010  d1 7d 05 b1 b1 89 1c b4  85 b9 d9 95 c9 d5 a9 0c  |.}..............|
	00000020  d0 dd bd c9 00 e8 00 00                           |........|
}

([]uint8) (len=42 cap=42) {
	00000000  d6 9d 20 04 40 04 68 34  10 3d 51 05 7d 0d a1 85  |.. .@.h4.=Q.}...|
	00000010  d1 7d 05 b1 b1 89 1c b4  85 b9 d9 95 c9 d5 a9 14  |.}..............|
	00000020  d0 a1 c9 95 95 c9 00 e8  00 00                    |..........|
}

([]uint8) (len=41 cap=41) {
	00000000  d6 99 20 04 40 04 68 34  10 3d 51 05 7d 0d a1 85  |.. .@.h4.=Q.}...|
	00000010  d1 7d 05 b1 b1 89 1c b4  85 b9 d9 95 c9 d5 a9 10  |.}..............|
	00000020  98 bd d5 c9 c9 00 e8 00  00                       |.........|
}

([]uint8) (len=41 cap=41) {
	00000000  d6 99 20 04 40 04 68 34  10 3d 51 05 7d 0d a1 85  |.. .@.h4.=Q.}...|
	00000010  d1 7d 05 b1 b1 89 1c b4  85 b9 d9 95 c9 d5 a9 10  |.}..............|
	00000020  98 a5 d9 95 c9 00 e8 00  00                       |.........|
}

([]uint8) (len=64 cap=64) {
	00000000  d6 f5 20 04 40 04 68 34  10 3d 51 05 7d 0d a1 85  |.. .@.h4.=Q.}...|
	00000010  d1 7d 05 b1 b1 89 1c b4  85 b9 d9 95 c9 d5 a9 6c  |.}.............l|
	00000020  bc b9 95 81 a0 d5 b9 91  c9 95 91 81 84 b9 91 81  |................|
	00000030  d0 dd 95 b9 d1 e5 81 98  a5 d9 95 c9 00 e8 00 00  |................|
}

([]uint8) (len=164 cap=164) {
	00000000  d6 81 06 20 04 40 04 68  34 10 3d 51 05 7d 0d a1  |... .@.h4.=Q.}..|
	00000010  85 d1 7d 05 b1 b1 89 1c  b4 85 b9 d9 95 c9 d5 a9  |..}.............|
	00000020  f8 31 bd c9 95 b5 81 a4  c1 cd d5 b5 81 90 bd b1  |.1..............|
	00000030  bd c9 81 cc a5 d1 81 84  b5 95 d1 b1 80 8c bd b9  |................|
	00000040  cd 95 8d d1 95 d1 d5 c9  81 84 91 a5 c1 a5 cd 8d  |................|
	00000050  a5 b9 9d 81 94 b1 a5 d1  b9 80 10 bd b9 95 8d 81  |................|
	00000060  84 81 90 a5 85 b5 81 b0  95 8d d1 d5 cd b9 80 4c  |...............L|
	00000070  95 91 81 cc a5 d1 81 84  b5 95 d1 81 a4 c1 cd d5  |................|
	00000080  b5 81 b4 85 d5 c9 a5 cd  b9 80 34 85 95 8d 95 b9  |..........4.....|
	00000090  85 cd 81 8c bd b9 9d d5  95 81 b0 a5 9d d5 b1 c9  |................|
	000000a0  00 e8 00 00                                       |....|
}

*/

func (p *Parser) OnCDemoSpawnGroups(msg Message, obj *dota.CDemoSpawnGroups) {
	// PP(obj.GetData())
	/*
		pbuf := proto.NewBuffer(obj.GetData())
		l, err := pbuf.DecodeVarint()
		if err != nil {
			panic(err)
		}
		PP(l)
	*/
}
func (p *Parser) OnCDemoUserCmd(msg Message, obj *dota.CDemoUserCmd) {
	// PP(obj.GetData())
}

func tryDecodeBuf(buf func() *proto.Buffer) []string {
	results := []string{}

	for _, msg := range msgs() {
		err := buf().Unmarshal(msg)
		if err != nil {
			return results
		}

		_, err = proto.Marshal(msg)
		if err != nil {
			return results
		}

		PP(msg)
	}

	return results
}

func tryDecodeBytes(buf []byte) {
	for _, msg := range msgs() {
		err := proto.Unmarshal(buf, msg)
		if err == nil {
			PP(msg)
		}
	}
}

func msgs() []proto.Message {
	return []proto.Message{
		&dota.C2S_CONNECT_Message{},
		&dota.CAdditionalEquipSlotClientMsg{},
		&dota.CAdditionalEquipSlot{},
		&dota.CAttribute_ItemDynamicRecipeComponent{},
		&dota.CAttribute_String{},
		&dota.CBidirMsg_RebroadcastGameEvent{},
		&dota.CBidirMsg_RebroadcastSource{},
		&dota.CCLCMsg_BaselineAck{},
		&dota.CCLCMsg_ClientInfo{},
		&dota.CCLCMsg_ClientMessage{},
		&dota.CCLCMsg_CmdKeyValues{},
		&dota.CCLCMsg_FileCRCCheck{},
		&dota.CCLCMsg_ListenEvents{},
		&dota.CCLCMsg_LoadingProgress{},
		&dota.CCLCMsg_Move{},
		&dota.CCLCMsg_RequestPause{},
		&dota.CCLCMsg_RespondCvarValue{},
		&dota.CCLCMsg_ServerPing{},
		&dota.CCLCMsg_ServerStatus{},
		&dota.CCLCMsg_SplitPlayerConnect{},
		&dota.CCLCMsg_SplitPlayerDisconnect{},
		&dota.CCLCMsg_VoiceData{},
		&dota.CCompendiumGameList{},
		&dota.CCompendiumGameTimeline{},
		&dota.CCompendiumTimestampedData{},
		&dota.CDOTABroadcastMsg_LANLobbyReply_CLobbyMember{},
		&dota.CDOTABroadcastMsg_LANLobbyReply{},
		&dota.CDOTABroadcastMsg_LANLobbyRequest{},
		&dota.CDOTABroadcastMsg{},
		&dota.CDOTAClientMsg_AspectRatio{},
		&dota.CDOTAClientMsg_AutoPurchaseItems{},
		&dota.CDOTAClientMsg_BeginLastHitChallenge{},
		&dota.CDOTAClientMsg_BroadcasterUsingAssistedCameraOperator{},
		&dota.CDOTAClientMsg_BroadcasterUsingCameraman{},
		&dota.CDOTAClientMsg_BuyBackStateAlert{},
		&dota.CDOTAClientMsg_CameraZoomAmount{},
		&dota.CDOTAClientMsg_ChatWheel{},
		&dota.CDOTAClientMsg_CoachHUDPing{},
		&dota.CDOTAClientMsg_EnemyItemAlert{},
		&dota.CDOTAClientMsg_FillEmptySlotsWithBots{},
		&dota.CDOTAClientMsg_FreeInventory{},
		&dota.CDOTAClientMsg_ItemAlert{},
		&dota.CDOTAClientMsg_MapLine{},
		&dota.CDOTAClientMsg_MapPing{},
		&dota.CDOTAClientMsg_Pause{},
		&dota.CDOTAClientMsg_PlayerShowCase{},
		&dota.CDOTAClientMsg_RecordVote{},
		&dota.CDOTAClientMsg_RequestGraphUpdate{},
		&dota.CDOTAClientMsg_SearchString{},
		&dota.CDOTAClientMsg_SendStatPopup{},
		&dota.CDOTAClientMsg_SetUnitShareFlag{},
		&dota.CDOTAClientMsg_ShopViewMode{},
		&dota.CDOTAClientMsg_SwapAccept{},
		&dota.CDOTAClientMsg_SwapRequest{},
		&dota.CDOTAClientMsg_TeleportRequiresHalt{},
		&dota.CDOTAClientMsg_TestItems{},
		&dota.CDOTAClientMsg_UnitsAutoAttackAfterSpell{},
		&dota.CDOTAClientMsg_UnitsAutoAttack{},
		&dota.CDOTAClientMsg_UpdateCoachListen{},
		&dota.CDOTAClientMsg_UpdateQuickBuyItem{},
		&dota.CDOTAClientMsg_UpdateQuickBuy{},
		&dota.CDOTAClientMsg_WillPurchaseAlert{},
		&dota.CDOTAClientMsg_WorldLine{},
		&dota.CDOTALobbyMember_CDOTALobbyMemberXPBonus{},
		&dota.CDOTALobbyMember{},
		&dota.CDOTAModifierBuffTableEntry{},
		&dota.CDOTAMsg_CoachHUDPing{},
		&dota.CDOTAMsg_EnemyItemAlert{},
		&dota.CDOTAMsg_ItemAlert{},
		&dota.CDOTAMsg_LocationPing{},
		&dota.CDOTAMsg_MapLine{},
		&dota.CDOTAMsg_SendStatPopup{},
		&dota.CDOTAMsg_WorldLine{},
		&dota.CDOTAResponseQuerySerialized_Fact{},
		&dota.CDOTAResponseQuerySerialized{},
		&dota.CDOTASaveGame_Player{},
		&dota.CDOTASaveGame_SaveInstance_PlayerPositions{},
		&dota.CDOTASaveGame_SaveInstance{},
		&dota.CDOTASaveGame{},
		&dota.CDOTASpeechMatchOnClient{},
		&dota.CDOTAUserMsg_AIDebugLine{},
		&dota.CDOTAUserMsg_AbilityPing{},
		&dota.CDOTAUserMsg_AbilitySteal{},
		&dota.CDOTAUserMsg_AddQuestLogEntry{},
		&dota.CDOTAUserMsg_BoosterStatePlayer{},
		&dota.CDOTAUserMsg_BoosterState{},
		&dota.CDOTAUserMsg_BotChat{},
		&dota.CDOTAUserMsg_BuyBackStateAlert{},
		&dota.CDOTAUserMsg_ChatEvent{},
		&dota.CDOTAUserMsg_ChatWheel{},
		&dota.CDOTAUserMsg_ClientLoadGridNav{},
		&dota.CDOTAUserMsg_CoachHUDPing{},
		&dota.CDOTAUserMsg_CombatHeroPositions{},
		&dota.CDOTAUserMsg_CombatLogShowDeath{},
		&dota.CDOTAUserMsg_CourierKilledAlert{},
		&dota.CDOTAUserMsg_CreateLinearProjectile{},
		&dota.CDOTAUserMsg_CustomHeaderMessage{},
		&dota.CDOTAUserMsg_CustomMsg{},
		&dota.CDOTAUserMsg_DestroyLinearProjectile{},
		&dota.CDOTAUserMsg_DodgeTrackingProjectiles{},
		&dota.CDOTAUserMsg_EnemyItemAlert{},
		&dota.CDOTAUserMsg_GamerulesStateChanged{},
		&dota.CDOTAUserMsg_GlobalLightColor{},
		&dota.CDOTAUserMsg_GlobalLightDirection{},
		&dota.CDOTAUserMsg_HalloweenDrops{},
		&dota.CDOTAUserMsg_HudError{},
		&dota.CDOTAUserMsg_InvalidCommand{},
		&dota.CDOTAUserMsg_ItemAlert{},
		&dota.CDOTAUserMsg_ItemFound{},
		&dota.CDOTAUserMsg_ItemPurchased{},
		&dota.CDOTAUserMsg_LocationPing{},
		&dota.CDOTAUserMsg_MapLine{},
		&dota.CDOTAUserMsg_MiniKillCamInfo_Attacker_Ability{},
		&dota.CDOTAUserMsg_MiniKillCamInfo_Attacker{},
		&dota.CDOTAUserMsg_MiniKillCamInfo{},
		&dota.CDOTAUserMsg_MiniTaunt{},
		&dota.CDOTAUserMsg_MinimapDebugPoint{},
		&dota.CDOTAUserMsg_MinimapEvent{},
		&dota.CDOTAUserMsg_NevermoreRequiem{},
		&dota.CDOTAUserMsg_OverheadEvent{},
		&dota.CDOTAUserMsg_ParticleManager_ChangeControlPointAttachment{},
		&dota.CDOTAUserMsg_ParticleManager_CreateParticle{},
		&dota.CDOTAUserMsg_ParticleManager_DestroyParticleInvolving{},
		&dota.CDOTAUserMsg_ParticleManager_DestroyParticle{},
		&dota.CDOTAUserMsg_ParticleManager_ReleaseParticleIndex{},
		&dota.CDOTAUserMsg_ParticleManager_UpdateParticleEnt{},
		&dota.CDOTAUserMsg_ParticleManager_UpdateParticleFallback{},
		&dota.CDOTAUserMsg_ParticleManager_UpdateParticleFwd{},
		&dota.CDOTAUserMsg_ParticleManager_UpdateParticleOffset{},
		&dota.CDOTAUserMsg_ParticleManager_UpdateParticleOrient{},
		&dota.CDOTAUserMsg_ParticleManager_UpdateParticleSetFrozen{},
		&dota.CDOTAUserMsg_ParticleManager_UpdateParticleShouldDraw{},
		&dota.CDOTAUserMsg_ParticleManager_UpdateParticle{},
		&dota.CDOTAUserMsg_ParticleManager{},
		&dota.CDOTAUserMsg_Ping{},
		&dota.CDOTAUserMsg_PlayerMMR{},
		&dota.CDOTAUserMsg_ReceivedXmasGift{},
		&dota.CDOTAUserMsg_SendFinalGold{},
		&dota.CDOTAUserMsg_SendGenericToolTip{},
		&dota.CDOTAUserMsg_SendRoshanPopup{},
		&dota.CDOTAUserMsg_SendStatPopup{},
		&dota.CDOTAUserMsg_SetNextAutobuyItem{},
		&dota.CDOTAUserMsg_SharedCooldown{},
		&dota.CDOTAUserMsg_ShowGenericPopup{},
		&dota.CDOTAUserMsg_ShowSurvey{},
		&dota.CDOTAUserMsg_SpectatorPlayerClick{},
		&dota.CDOTAUserMsg_SpeechBubble{},
		&dota.CDOTAUserMsg_StatsHeroMinuteDetails{},
		&dota.CDOTAUserMsg_StatsKillDetails{},
		&dota.CDOTAUserMsg_StatsMatchDetails{},
		&dota.CDOTAUserMsg_StatsPlayerKillShare{},
		&dota.CDOTAUserMsg_StatsTeamMinuteDetails{},
		&dota.CDOTAUserMsg_SwapVerify{},
		&dota.CDOTAUserMsg_TE_DotaBloodImpact{},
		&dota.CDOTAUserMsg_TE_ProjectileLoc{},
		&dota.CDOTAUserMsg_TE_Projectile{},
		&dota.CDOTAUserMsg_TE_UnitAnimationEnd{},
		&dota.CDOTAUserMsg_TE_UnitAnimation{},
		&dota.CDOTAUserMsg_TutorialFade{},
		&dota.CDOTAUserMsg_TutorialFinish{},
		&dota.CDOTAUserMsg_TutorialMinimapPosition{},
		&dota.CDOTAUserMsg_TutorialPingMinimap{},
		&dota.CDOTAUserMsg_TutorialRequestExp{},
		&dota.CDOTAUserMsg_TutorialTipInfo{},
		&dota.CDOTAUserMsg_UnitEvent_AddGesture{},
		&dota.CDOTAUserMsg_UnitEvent_BloodImpact{},
		&dota.CDOTAUserMsg_UnitEvent_FadeGesture{},
		&dota.CDOTAUserMsg_UnitEvent_RemoveGesture{},
		&dota.CDOTAUserMsg_UnitEvent_SpeechMute{},
		&dota.CDOTAUserMsg_UnitEvent_Speech{},
		&dota.CDOTAUserMsg_UnitEvent{},
		&dota.CDOTAUserMsg_UpdateSharedContent{},
		&dota.CDOTAUserMsg_VoteEnd{},
		&dota.CDOTAUserMsg_VoteStart{},
		&dota.CDOTAUserMsg_VoteUpdate{},
		&dota.CDOTAUserMsg_WillPurchaseAlert{},
		&dota.CDOTAUserMsg_WorldLine{},
		&dota.CDemoClassInfoClassT{},
		&dota.CDemoClassInfo{},
		&dota.CDemoConsoleCmd{},
		&dota.CDemoCustomDataCallbacks{},
		&dota.CDemoCustomData{},
		&dota.CDemoFileHeader{},
		&dota.CDemoFileInfo{},
		&dota.CDemoFullPacket{},
		&dota.CDemoPacket{},
		&dota.CDemoSaveGame{},
		&dota.CDemoSendTables{},
		&dota.CDemoSpawnGroups{},
		&dota.CDemoStop{},
		&dota.CDemoStringTablesItemsT{},
		&dota.CDemoStringTablesTableT{},
		&dota.CDemoStringTables{},
		&dota.CDemoSyncTick{},
		&dota.CDemoUserCmd{},
		&dota.CEntityMessageDoSpark{},
		&dota.CEntityMessageFixAngle{},
		&dota.CEntityMessagePlayJingle{},
		&dota.CEntityMessagePropagateForce{},
		&dota.CEntityMessageRemoveAllDecals{},
		&dota.CEntityMessageScreenOverlay{},
		&dota.CGCMsgGetIPLocationResponse{},
		&dota.CGCMsgGetIPLocation{},
		&dota.CGCMsgGetSystemStatsResponse{},
		&dota.CGCMsgGetSystemStats{},
		&dota.CGCMsgMemCachedDelete{},
		&dota.CGCMsgMemCachedGetResponse_ValueTag{},
		&dota.CGCMsgMemCachedGetResponse{},
		&dota.CGCMsgMemCachedGet{},
		&dota.CGCMsgMemCachedSet_KeyPair{},
		&dota.CGCMsgMemCachedSet{},
		&dota.CGCMsgMemCachedStatsResponse{},
		&dota.CGCMsgMemCachedStats{},
		&dota.CGCMsgSQLStatsResponse{},
		&dota.CGCMsgSQLStats{},
		&dota.CGCMsgSystemStatsSchema{},
		&dota.CGCStorePurchaseInit_LineItem{},
		&dota.CGCSystemMsg_GetAccountDetails_Response{},
		&dota.CGCSystemMsg_GetAccountDetails{},
		&dota.CGCSystemMsg_GetPurchaseTrust_Request{},
		&dota.CGCSystemMsg_GetPurchaseTrust_Response{},
		&dota.CGCToGCMsgMasterAck_Process{},
		&dota.CGCToGCMsgMasterAck_Response{},
		&dota.CGCToGCMsgMasterAck{},
		&dota.CGCToGCMsgMasterStartupComplete_GCInfo{},
		&dota.CGCToGCMsgMasterStartupComplete{},
		&dota.CGCToGCMsgRoutedReply{},
		&dota.CGCToGCMsgRouted{},
		&dota.CGameInfo_CDotaGameInfo_CHeroSelectEvent{},
		&dota.CGameInfo_CDotaGameInfo_CPlayerInfo{},
		&dota.CGameInfo_CDotaGameInfo{},
		&dota.CGameInfo{},
		&dota.CHTMLHeader{},
		&dota.CHTMLPageSecurityInfo{},
		&dota.CIPLocationInfo{},
		&dota.CLeague{},
		&dota.CLobbyBroadcastChannelInfo{},
		&dota.CLobbyTeamDetails{},
		&dota.CLobbyTimedRewardDetails{},
		&dota.CMatchAdditionalUnitInventory{},
		&dota.CMatchHeroSelectEvent{},
		&dota.CMatchPlayerAbilityUpgrade{},
		&dota.CMsgAMAddFreeLicenseResponse{},
		&dota.CMsgAMAddFreeLicense{},
		&dota.CMsgAMFindAccountsResponse{},
		&dota.CMsgAMFindAccounts{},
		&dota.CMsgAMGetLicensesResponse{},
		&dota.CMsgAMGetLicenses{},
		&dota.CMsgAMGetUserGameStatsResponse_Achievement_Blocks{},
		&dota.CMsgAMGetUserGameStatsResponse_Stats{},
		&dota.CMsgAMGetUserGameStatsResponse{},
		&dota.CMsgAMGetUserGameStats{},
		&dota.CMsgAMGrantGuestPasses2Response{},
		&dota.CMsgAMGrantGuestPasses2{},
		&dota.CMsgAMSendEmailResponse{},
		&dota.CMsgAMSendEmail_PersonaNameReplacementToken{},
		&dota.CMsgAMSendEmail_ReplacementToken{},
		&dota.CMsgAMSendEmail{},
		&dota.CMsgAbandonCurrentGame{},
		&dota.CMsgAddHeader{},
		&dota.CMsgAddItemToSocketData{},
		&dota.CMsgAddItemToSocketResponse{},
		&dota.CMsgAddItemToSocket{},
		&dota.CMsgAddSocketResponse{},
		&dota.CMsgAddSocket{},
		&dota.CMsgAdjustItemEquippedState{},
		&dota.CMsgApplyAutograph{},
		&dota.CMsgApplyEggEssence{},
		&dota.CMsgApplyPennantUpgrade{},
		&dota.CMsgApplyStrangePart{},
		&dota.CMsgApplyTeamToPracticeLobby{},
		&dota.CMsgBalancedShuffleLobby{},
		&dota.CMsgBotGameCreate{},
		&dota.CMsgBrowserCreateResponse{},
		&dota.CMsgBrowserCreate{},
		&dota.CMsgBrowserErrorStrings{},
		&dota.CMsgBrowserPosition{},
		&dota.CMsgBrowserReady{},
		&dota.CMsgBrowserRemove{},
		&dota.CMsgBrowserSize{},
		&dota.CMsgCanGoBackAndForward{},
		&dota.CMsgCancelWatchGame{},
		&dota.CMsgCastMatchVote{},
		&dota.CMsgClearDecalsForSkeletonInstanceEvent{},
		&dota.CMsgClearEntityDecalsEvent{},
		&dota.CMsgClearWorldDecalsEvent{},
		&dota.CMsgClientGameserverPings_PingDetail{},
		&dota.CMsgClientGameserverPings{},
		&dota.CMsgClientHello{},
		&dota.CMsgClientProvideSurveyResult_Response{},
		&dota.CMsgClientProvideSurveyResult{},
		&dota.CMsgClientSuspended{},
		&dota.CMsgClientToGCEmoticonDataRequest{},
		&dota.CMsgClientToGCGetAdditionalEquipsResponse{},
		&dota.CMsgClientToGCGetAdditionalEquips{},
		&dota.CMsgClientToGCGetAllHeroProgressResponse{},
		&dota.CMsgClientToGCGetAllHeroProgress{},
		&dota.CMsgClientToGCGetProfileCard{},
		&dota.CMsgClientToGCGetTrophyListResponse_Trophy{},
		&dota.CMsgClientToGCGetTrophyListResponse{},
		&dota.CMsgClientToGCGetTrophyList{},
		&dota.CMsgClientToGCMarkNotificationListRead{},
		&dota.CMsgClientToGCSetAdditionalEquips{},
		&dota.CMsgClientToGCSetProfileCardSlot{},
		&dota.CMsgClientToGCTrackDialogResult{},
		&dota.CMsgClientWelcome_Location{},
		&dota.CMsgClientWelcome{},
		&dota.CMsgClientsRejoinChatChannels{},
		&dota.CMsgClose{},
		&dota.CMsgConVarValue{},
		&dota.CMsgConnectedPlayers_Player{},
		&dota.CMsgConnectedPlayers{},
		&dota.CMsgConnectionStatus{},
		&dota.CMsgConsumableExhausted{},
		&dota.CMsgCopy{},
		&dota.CMsgCraftStatue{},
		&dota.CMsgCraftingResponse{},
		&dota.CMsgCustomGameCreate{},
		&dota.CMsgDOTAAccountGuildMembershipsSDO_Invitation{},
		&dota.CMsgDOTAAccountGuildMembershipsSDO_Membership{},
		&dota.CMsgDOTAAccountGuildMembershipsSDO{},
		&dota.CMsgDOTAAwardEventPoints_AwardPoints{},
		&dota.CMsgDOTAAwardEventPoints{},
		&dota.CMsgDOTABetaParticipation{},
		&dota.CMsgDOTABroadcastNotification{},
		&dota.CMsgDOTAChatChannelFullUpdate{},
		&dota.CMsgDOTAChatChannelMemberUpdate_JoinedMember{},
		&dota.CMsgDOTAChatChannelMemberUpdate{},
		&dota.CMsgDOTAChatGetUserListResponse_Member{},
		&dota.CMsgDOTAChatGetUserListResponse{},
		&dota.CMsgDOTAChatGetUserList{},
		&dota.CMsgDOTAChatMember{},
		&dota.CMsgDOTAChatMessage{},
		&dota.CMsgDOTAClearNotifySuccessfulReport{},
		&dota.CMsgDOTAClearTournamentGame{},
		&dota.CMsgDOTAClientIgnoredUser{},
		&dota.CMsgDOTACombatLogEntry{},
		&dota.CMsgDOTACompendiumDataRequest{},
		&dota.CMsgDOTACompendiumDataResponse{},
		&dota.CMsgDOTACompendiumData{},
		&dota.CMsgDOTACompendiumSelectionResponse{},
		&dota.CMsgDOTACompendiumSelection{},
		&dota.CMsgDOTAConsumeFantasyTicketFailure{},
		&dota.CMsgDOTAConsumeFantasyTicket{},
		&dota.CMsgDOTACreateFantasyLeagueRequest{},
		&dota.CMsgDOTACreateFantasyLeagueResponse{},
		&dota.CMsgDOTACreateTeamResponse{},
		&dota.CMsgDOTACreateTeam{},
		&dota.CMsgDOTADiretidePrizesRewardedResponse{},
		&dota.CMsgDOTADisbandTeamResponse{},
		&dota.CMsgDOTADisbandTeam{},
		&dota.CMsgDOTAEditFantasyTeamRequest{},
		&dota.CMsgDOTAEditFantasyTeamResponse{},
		&dota.CMsgDOTAEditTeamDetailsResponse{},
		&dota.CMsgDOTAEditTeamDetails{},
		&dota.CMsgDOTAEditTeamLogoResponse{},
		&dota.CMsgDOTAEditTeamLogo{},
		&dota.CMsgDOTAEditTeam{},
		&dota.CMsgDOTAEmoticonAccessSDO{},
		&dota.CMsgDOTAExchangeEventPointsResponse{},
		&dota.CMsgDOTAExchangeEventPoints{},
		&dota.CMsgDOTAFantasyLeagueCreateRequest{},
		&dota.CMsgDOTAFantasyLeagueCreateResponse{},
		&dota.CMsgDOTAFantasyLeagueDraftPlayerRequest{},
		&dota.CMsgDOTAFantasyLeagueDraftPlayerResponse{},
		&dota.CMsgDOTAFantasyLeagueDraftStatusRequest{},
		&dota.CMsgDOTAFantasyLeagueDraftStatus{},
		&dota.CMsgDOTAFantasyLeagueEditInfoRequest{},
		&dota.CMsgDOTAFantasyLeagueEditInfoResponse{},
		&dota.CMsgDOTAFantasyLeagueEditInvitesRequest_InviteChange{},
		&dota.CMsgDOTAFantasyLeagueEditInvitesRequest{},
		&dota.CMsgDOTAFantasyLeagueEditInvitesResponse{},
		&dota.CMsgDOTAFantasyLeagueFindRequest{},
		&dota.CMsgDOTAFantasyLeagueFindResponse{},
		&dota.CMsgDOTAFantasyLeagueInfoRequest{},
		&dota.CMsgDOTAFantasyLeagueInfoResponse{},
		&dota.CMsgDOTAFantasyLeagueInfo_OwnerInfo{},
		&dota.CMsgDOTAFantasyLeagueInfo{},
		&dota.CMsgDOTAFantasyLeagueMatchupsRequest{},
		&dota.CMsgDOTAFantasyLeagueMatchupsResponse_Matchup{},
		&dota.CMsgDOTAFantasyLeagueMatchupsResponse_WeeklyMatchups{},
		&dota.CMsgDOTAFantasyLeagueMatchupsResponse{},
		&dota.CMsgDOTAFantasyLeaveLeagueRequest{},
		&dota.CMsgDOTAFantasyLeaveLeagueResponse{},
		&dota.CMsgDOTAFantasyMatchFinished{},
		&dota.CMsgDOTAFantasyMessageAdd{},
		&dota.CMsgDOTAFantasyMessagesRequest{},
		&dota.CMsgDOTAFantasyMessagesResponse_Message{},
		&dota.CMsgDOTAFantasyMessagesResponse{},
		&dota.CMsgDOTAFantasyPlayerHisoricalStatsRequest{},
		&dota.CMsgDOTAFantasyPlayerHisoricalStatsResponse_PlayerScoreAccumulator{},
		&dota.CMsgDOTAFantasyPlayerHisoricalStatsResponse_PlayerStats{},
		&dota.CMsgDOTAFantasyPlayerHisoricalStatsResponse{},
		&dota.CMsgDOTAFantasyPlayerInfoRequest{},
		&dota.CMsgDOTAFantasyPlayerInfoResponse{},
		&dota.CMsgDOTAFantasyPlayerScoreDetailsRequest{},
		&dota.CMsgDOTAFantasyPlayerScoreDetailsResponse_PlayerMatchData{},
		&dota.CMsgDOTAFantasyPlayerScoreDetailsResponse{},
		&dota.CMsgDOTAFantasyPlayerScoreRequest{},
		&dota.CMsgDOTAFantasyPlayerScoreResponse{},
		&dota.CMsgDOTAFantasyPlayerStandingsRequest{},
		&dota.CMsgDOTAFantasyPlayerStandingsResponse_CMsgPlayerScore{},
		&dota.CMsgDOTAFantasyPlayerStandingsResponse{},
		&dota.CMsgDOTAFantasyPlayerStats{},
		&dota.CMsgDOTAFantasyRemoveOwnerResponse{},
		&dota.CMsgDOTAFantasyRemoveOwner{},
		&dota.CMsgDOTAFantasyScheduledMatchesRequest{},
		&dota.CMsgDOTAFantasyScheduledMatchesResponse_ScheduledMatchDays{},
		&dota.CMsgDOTAFantasyScheduledMatchesResponse{},
		&dota.CMsgDOTAFantasyTeamCreateRequest{},
		&dota.CMsgDOTAFantasyTeamCreateResponse{},
		&dota.CMsgDOTAFantasyTeamInfoRequestByFantasyLeagueID{},
		&dota.CMsgDOTAFantasyTeamInfoRequestByOwnerAccountID{},
		&dota.CMsgDOTAFantasyTeamInfoResponse{},
		&dota.CMsgDOTAFantasyTeamInfo{},
		&dota.CMsgDOTAFantasyTeamRosterAddDropRequest{},
		&dota.CMsgDOTAFantasyTeamRosterAddDropResponse{},
		&dota.CMsgDOTAFantasyTeamRosterRequest{},
		&dota.CMsgDOTAFantasyTeamRosterResponse{},
		&dota.CMsgDOTAFantasyTeamRosterSwapRequest{},
		&dota.CMsgDOTAFantasyTeamRosterSwapResponse{},
		&dota.CMsgDOTAFantasyTeamScoreRequest{},
		&dota.CMsgDOTAFantasyTeamScoreResponse_CMsgPlayerScore{},
		&dota.CMsgDOTAFantasyTeamScoreResponse{},
		&dota.CMsgDOTAFantasyTeamStandingsRequest{},
		&dota.CMsgDOTAFantasyTeamStandingsResponse_CMsgTeamScore{},
		&dota.CMsgDOTAFantasyTeamStandingsResponse{},
		&dota.CMsgDOTAFantasyTeamTradeCancelRequest{},
		&dota.CMsgDOTAFantasyTeamTradeCancelResponse{},
		&dota.CMsgDOTAFantasyTeamTradesRequest{},
		&dota.CMsgDOTAFantasyTeamTradesResponse_Trade{},
		&dota.CMsgDOTAFantasyTeamTradesResponse{},
		&dota.CMsgDOTAFeaturedItems{},
		&dota.CMsgDOTAFriendRecruitInviteAcceptDecline{},
		&dota.CMsgDOTAFriendRecruitsRequest{},
		&dota.CMsgDOTAFriendRecruitsResponse_FriendRecruitStatus{},
		&dota.CMsgDOTAFriendRecruitsResponse{},
		&dota.CMsgDOTAFrostivusTimeElapsed_User{},
		&dota.CMsgDOTAFrostivusTimeElapsed{},
		&dota.CMsgDOTAGCToGCDestroyOpenGuildPartyRequest{},
		&dota.CMsgDOTAGCToGCDestroyOpenGuildPartyResponse{},
		&dota.CMsgDOTAGCToGCUpdateOpenGuildPartyRequest{},
		&dota.CMsgDOTAGCToGCUpdateOpenGuildPartyResponse{},
		&dota.CMsgDOTAGenerateDiretidePrizeListResponse{},
		&dota.CMsgDOTAGenerateDiretidePrizeList{},
		&dota.CMsgDOTAGenericResult{},
		&dota.CMsgDOTAGetEventPointsResponse_Action{},
		&dota.CMsgDOTAGetEventPointsResponse{},
		&dota.CMsgDOTAGetEventPoints{},
		&dota.CMsgDOTAGetPlayerMatchHistoryResponse_Match{},
		&dota.CMsgDOTAGetPlayerMatchHistoryResponse{},
		&dota.CMsgDOTAGetPlayerMatchHistory{},
		&dota.CMsgDOTAGuildAuditSDO_Entry{},
		&dota.CMsgDOTAGuildAuditSDO{},
		&dota.CMsgDOTAGuildCancelInviteRequest{},
		&dota.CMsgDOTAGuildCancelInviteResponse{},
		&dota.CMsgDOTAGuildCreateRequest{},
		&dota.CMsgDOTAGuildCreateResponse{},
		&dota.CMsgDOTAGuildEditLogoRequest{},
		&dota.CMsgDOTAGuildEditLogoResponse{},
		&dota.CMsgDOTAGuildInviteAccountRequest{},
		&dota.CMsgDOTAGuildInviteAccountResponse{},
		&dota.CMsgDOTAGuildInviteData{},
		&dota.CMsgDOTAGuildOpenPartyRefresh_OpenParty{},
		&dota.CMsgDOTAGuildOpenPartyRefresh{},
		&dota.CMsgDOTAGuildSDO_Invitation{},
		&dota.CMsgDOTAGuildSDO_Member{},
		&dota.CMsgDOTAGuildSDO{},
		&dota.CMsgDOTAGuildSetAccountRoleRequest{},
		&dota.CMsgDOTAGuildSetAccountRoleResponse{},
		&dota.CMsgDOTAGuildUpdateDetailsRequest{},
		&dota.CMsgDOTAGuildUpdateDetailsResponse{},
		&dota.CMsgDOTAGuildUpdateMessage{},
		&dota.CMsgDOTAHallOfFameRequest{},
		&dota.CMsgDOTAHallOfFameResponse{},
		&dota.CMsgDOTAHallOfFame_FeaturedFarmer{},
		&dota.CMsgDOTAHallOfFame_FeaturedPlayer{},
		&dota.CMsgDOTAHallOfFame{},
		&dota.CMsgDOTAHalloweenHighScoreRequest{},
		&dota.CMsgDOTAHalloweenHighScoreResponse{},
		&dota.CMsgDOTAHasItemQuery{},
		&dota.CMsgDOTAHasItemResponse{},
		&dota.CMsgDOTAHeroFavoritesAdd{},
		&dota.CMsgDOTAHeroFavoritesRemove{},
		&dota.CMsgDOTAJoinChatChannelResponse{},
		&dota.CMsgDOTAJoinChatChannel{},
		&dota.CMsgDOTAJoinOpenGuildPartyRequest{},
		&dota.CMsgDOTAJoinOpenGuildPartyResponse{},
		&dota.CMsgDOTAKickTeamMemberResponse{},
		&dota.CMsgDOTAKickTeamMember{},
		&dota.CMsgDOTAKickedFromMatchmakingQueue{},
		&dota.CMsgDOTALastHitChallengeHighScorePost{},
		&dota.CMsgDOTALastHitChallengeHighScoreRequest{},
		&dota.CMsgDOTALastHitChallengeHighScoreResponse{},
		&dota.CMsgDOTALeagueScheduleEditResponse{},
		&dota.CMsgDOTALeagueScheduleEdit{},
		&dota.CMsgDOTALeagueScheduleRequest{},
		&dota.CMsgDOTALeagueScheduleResponse{},
		&dota.CMsgDOTALeaguesInMonthRequest{},
		&dota.CMsgDOTALeaguesInMonthResponse{},
		&dota.CMsgDOTALeague{},
		&dota.CMsgDOTALeaveChatChannel{},
		&dota.CMsgDOTALeaveTeamResponse{},
		&dota.CMsgDOTALeaveTeam{},
		&dota.CMsgDOTALiveLeagueGameUpdate{},
		&dota.CMsgDOTALiveScoreboardUpdate_Team_Player{},
		&dota.CMsgDOTALiveScoreboardUpdate_Team{},
		&dota.CMsgDOTALiveScoreboardUpdate{},
		&dota.CMsgDOTAMatchHistoryFilter{},
		&dota.CMsgDOTAMatchVotes_PlayerVote{},
		&dota.CMsgDOTAMatchVotes{},
		&dota.CMsgDOTAMatch_BroadcasterChannel{},
		&dota.CMsgDOTAMatch_BroadcasterInfo{},
		&dota.CMsgDOTAMatch_Player{},
		&dota.CMsgDOTAMatchmakingStatsRequest{},
		&dota.CMsgDOTAMatchmakingStatsResponse{},
		&dota.CMsgDOTAMatch{},
		&dota.CMsgDOTANotifyAccountFlagsChange{},
		&dota.CMsgDOTAOtherJoinedChatChannel{},
		&dota.CMsgDOTAOtherLeftChatChannel{},
		&dota.CMsgDOTAPCBangTimedReward{},
		&dota.CMsgDOTAPartyMemberSetCoach{},
		&dota.CMsgDOTAPartySetOpenGuildRequest{},
		&dota.CMsgDOTAPartySetOpenGuildResponse{},
		&dota.CMsgDOTAPassportPlayerCardChallenge{},
		&dota.CMsgDOTAPassportStampedPlayer{},
		&dota.CMsgDOTAPassportVoteGenericSelection{},
		&dota.CMsgDOTAPassportVoteTeamGuess{},
		&dota.CMsgDOTAPassportVote{},
		&dota.CMsgDOTAPlayerFailedToConnect{},
		&dota.CMsgDOTAPlayerMatchHistory{},
		&dota.CMsgDOTAPopup{},
		&dota.CMsgDOTAProTeamListRequest{},
		&dota.CMsgDOTAProTeamListResponse_TeamEntry{},
		&dota.CMsgDOTAProTeamListResponse{},
		&dota.CMsgDOTAProcessFantasyScheduledEvent{},
		&dota.CMsgDOTAProfileCard_Slot_Item{},
		&dota.CMsgDOTAProfileCard_Slot_Stat{},
		&dota.CMsgDOTAProfileCard_Slot_Trophy{},
		&dota.CMsgDOTAProfileCard_Slot{},
		&dota.CMsgDOTAProfileCard{},
		&dota.CMsgDOTAProfileRequest{},
		&dota.CMsgDOTAProfileResponse_EventTicket{},
		&dota.CMsgDOTAProfileResponse_FeaturedItem{},
		&dota.CMsgDOTAProfileResponse_LeaguePass{},
		&dota.CMsgDOTAProfileResponse_PlayedHero{},
		&dota.CMsgDOTAProfileResponse_ShowcaseHero{},
		&dota.CMsgDOTAProfileResponse{},
		&dota.CMsgDOTARedeemEventPrizeResponse{},
		&dota.CMsgDOTARedeemEventPrize{},
		&dota.CMsgDOTARedeemItemResponse{},
		&dota.CMsgDOTARedeemItem{},
		&dota.CMsgDOTAReportCountsRequest{},
		&dota.CMsgDOTAReportCountsResponse{},
		&dota.CMsgDOTAReportsRemainingRequest{},
		&dota.CMsgDOTAReportsRemainingResponse{},
		&dota.CMsgDOTARequestBatchPlayerResourcesResponse_Result{},
		&dota.CMsgDOTARequestBatchPlayerResourcesResponse{},
		&dota.CMsgDOTARequestBatchPlayerResources{},
		&dota.CMsgDOTARequestChatChannelListResponse_ChatChannel{},
		&dota.CMsgDOTARequestChatChannelListResponse{},
		&dota.CMsgDOTARequestChatChannelList{},
		&dota.CMsgDOTARequestGuildData{},
		&dota.CMsgDOTARequestMatchesResponse_Series{},
		&dota.CMsgDOTARequestMatchesResponse{},
		&dota.CMsgDOTARequestMatches{},
		&dota.CMsgDOTARequestPlayerResourcesResponse{},
		&dota.CMsgDOTARequestPlayerResources{},
		&dota.CMsgDOTARequestSaveGamesResponse{},
		&dota.CMsgDOTARequestSaveGames{},
		&dota.CMsgDOTARequestTeamDataResponse{},
		&dota.CMsgDOTARequestTeamData{},
		&dota.CMsgDOTARewardDiretidePrizes{},
		&dota.CMsgDOTARewardTutorialPrizes{},
		&dota.CMsgDOTASendFriendRecruits{},
		&dota.CMsgDOTASetMatchHistoryAccessResponse{},
		&dota.CMsgDOTASetMatchHistoryAccess{},
		&dota.CMsgDOTASetProfilePrivacyResponse{},
		&dota.CMsgDOTASetProfilePrivacy{},
		&dota.CMsgDOTAStartDailyHeroChallenge{},
		&dota.CMsgDOTAStorePromoPagesRequest{},
		&dota.CMsgDOTAStorePromoPagesResponse_PromoPage{},
		&dota.CMsgDOTAStorePromoPagesResponse{},
		&dota.CMsgDOTASubmitPlayerReportResponse{},
		&dota.CMsgDOTASubmitPlayerReport{},
		&dota.CMsgDOTATeamAdminSDO{},
		&dota.CMsgDOTATeamData{},
		&dota.CMsgDOTATeamIDByNameRequest{},
		&dota.CMsgDOTATeamIDByNameResponse{},
		&dota.CMsgDOTATeamInvite_GCImmediateResponseToInviter{},
		&dota.CMsgDOTATeamInvite_GCRequestToInvitee{},
		&dota.CMsgDOTATeamInvite_GCResponseToInvitee{},
		&dota.CMsgDOTATeamInvite_GCResponseToInviter{},
		&dota.CMsgDOTATeamInvite_InviteeResponseToGC{},
		&dota.CMsgDOTATeamInvite_InviterToGC{},
		&dota.CMsgDOTATeamMemberProfileRequest{},
		&dota.CMsgDOTATeamMemberSDO{},
		&dota.CMsgDOTATeamMember{},
		&dota.CMsgDOTATeamOnProfile{},
		&dota.CMsgDOTATeamProfileRequest{},
		&dota.CMsgDOTATeamProfileResponse{},
		&dota.CMsgDOTATeam{},
		&dota.CMsgDOTATournamentRequest{},
		&dota.CMsgDOTATournamentResponse{},
		&dota.CMsgDOTATournament_Game{},
		&dota.CMsgDOTATournament_Node{},
		&dota.CMsgDOTATournament_Team{},
		&dota.CMsgDOTATournament{},
		&dota.CMsgDOTATransferTeamAdminResponse{},
		&dota.CMsgDOTATransferTeamAdmin{},
		&dota.CMsgDOTAUpdateTI4HeroQuest_Player{},
		&dota.CMsgDOTAUpdateTI4HeroQuest{},
		&dota.CMsgDOTAWelcome{},
		&dota.CMsgDPPartnerMicroTxnsResponse{},
		&dota.CMsgDPPartnerMicroTxns_PartnerInfo{},
		&dota.CMsgDPPartnerMicroTxns_PartnerMicroTxn{},
		&dota.CMsgDPPartnerMicroTxns{},
		&dota.CMsgDevNewItemRequestResponse{},
		&dota.CMsgDevNewItemRequest{},
		&dota.CMsgDismissLootGreevilResponse{},
		&dota.CMsgDismissLootGreevil{},
		&dota.CMsgEconPlayerStrangeCountAdjustment_CStrangeCountAdjustment{},
		&dota.CMsgEconPlayerStrangeCountAdjustment{},
		&dota.CMsgEffectData{},
		&dota.CMsgEventGameCreate{},
		&dota.CMsgExecuteJavaScript{},
		&dota.CMsgExtractGems{},
		&dota.CMsgFantasyLeagueScoring{},
		&dota.CMsgFileLoadDialogResponse{},
		&dota.CMsgFileLoadDialog{},
		&dota.CMsgFindSourceTVGames{},
		&dota.CMsgFind{},
		&dota.CMsgFinishedRequest{},
		&dota.CMsgFlipLobbyTeams{},
		&dota.CMsgFriendPracticeLobbyListRequest{},
		&dota.CMsgFriendPracticeLobbyListResponse{},
		&dota.CMsgFulfillDynamicRecipeComponent{},
		&dota.CMsgGCBanStatusRequest{},
		&dota.CMsgGCBanStatusResponse{},
		&dota.CMsgGCBannedWordListRequest{},
		&dota.CMsgGCBannedWordListResponse{},
		&dota.CMsgGCBannedWord{},
		&dota.CMsgGCCheckFriendship_Response{},
		&dota.CMsgGCCheckFriendship{},
		&dota.CMsgGCClientDisplayNotification{},
		&dota.CMsgGCClientMarketDataEntry{},
		&dota.CMsgGCClientMarketDataRequest{},
		&dota.CMsgGCClientMarketData{},
		&dota.CMsgGCClientPing{},
		&dota.CMsgGCClientVersionUpdated{},
		&dota.CMsgGCCollectItem{},
		&dota.CMsgGCEconSQLWorkItemEmbeddedRollbackData{},
		&dota.CMsgGCError{},
		&dota.CMsgGCGetCommandListResponse{},
		&dota.CMsgGCGetCommandList{},
		&dota.CMsgGCGetEmailTemplateResponse{},
		&dota.CMsgGCGetEmailTemplate{},
		&dota.CMsgGCGetHeroStandingsResponse_Hero{},
		&dota.CMsgGCGetHeroStandingsResponse{},
		&dota.CMsgGCGetHeroStandings{},
		&dota.CMsgGCGetPartnerAccountLink_Response{},
		&dota.CMsgGCGetPartnerAccountLink{},
		&dota.CMsgGCGetPersonaNames_Response_PersonaName{},
		&dota.CMsgGCGetPersonaNames_Response{},
		&dota.CMsgGCGetPersonaNames{},
		&dota.CMsgGCHAccountVacStatusChange{},
		&dota.CMsgGCHUpdateSession_ExtraField{},
		&dota.CMsgGCHUpdateSession{},
		&dota.CMsgGCIncrementKillCountResponse{},
		&dota.CMsgGCItemEditorReleaseReservationResponse{},
		&dota.CMsgGCItemEditorReleaseReservation{},
		&dota.CMsgGCItemEditorReservationsRequest{},
		&dota.CMsgGCItemEditorReservationsResponse{},
		&dota.CMsgGCItemEditorReservation{},
		&dota.CMsgGCItemEditorReserveItemDefResponse{},
		&dota.CMsgGCItemEditorReserveItemDef{},
		&dota.CMsgGCItemPreviewItemBoughtNotification{},
		&dota.CMsgGCLeagueAdminState_PrivateLeagueKeys{},
		&dota.CMsgGCLeagueAdminState{},
		&dota.CMsgGCLobbyUpdateBroadcastChannelInfo{},
		&dota.CMsgGCMatchDetailsRequest{},
		&dota.CMsgGCMatchDetailsResponse{},
		&dota.CMsgGCMsgMasterSetClientMsgRouting_Entry{},
		&dota.CMsgGCMsgMasterSetClientMsgRouting_Response{},
		&dota.CMsgGCMsgMasterSetClientMsgRouting{},
		&dota.CMsgGCMsgMasterSetDirectory_Response{},
		&dota.CMsgGCMsgMasterSetDirectory_SubGC{},
		&dota.CMsgGCMsgMasterSetDirectory{},
		&dota.CMsgGCMsgMasterSetWebAPIRouting_Entry{},
		&dota.CMsgGCMsgMasterSetWebAPIRouting_Response{},
		&dota.CMsgGCMsgMasterSetWebAPIRouting{},
		&dota.CMsgGCMsgSetOptions_MessageRange{},
		&dota.CMsgGCMsgSetOptions{},
		&dota.CMsgGCMsgWebAPIJobRequestForwardResponse{},
		&dota.CMsgGCMultiplexMessage{},
		&dota.CMsgGCNameItemNotification{},
		&dota.CMsgGCNewBloomModeStateResponse{},
		&dota.CMsgGCNewBloomModeState{},
		&dota.CMsgGCNotificationsMarkReadRequest{},
		&dota.CMsgGCNotificationsRequest{},
		&dota.CMsgGCNotificationsResponse_Notification{},
		&dota.CMsgGCNotificationsResponse{},
		&dota.CMsgGCPartnerBalanceRequest{},
		&dota.CMsgGCPartnerBalanceResponse{},
		&dota.CMsgGCPartnerRechargeRedirectURLRequest{},
		&dota.CMsgGCPartnerRechargeRedirectURLResponse{},
		&dota.CMsgGCPlayerInfoRequest{},
		&dota.CMsgGCPlayerInfoSubmitResponse{},
		&dota.CMsgGCPlayerInfoSubmit{},
		&dota.CMsgGCPlayerInfo_PlayerInfo{},
		&dota.CMsgGCPlayerInfo{},
		&dota.CMsgGCReportAbuseResponse{},
		&dota.CMsgGCReportAbuse{},
		&dota.CMsgGCRequestStoreSalesDataResponse_Price{},
		&dota.CMsgGCRequestStoreSalesDataResponse{},
		&dota.CMsgGCRequestStoreSalesDataUpToDateResponse{},
		&dota.CMsgGCRequestStoreSalesData{},
		&dota.CMsgGCRequestSubGCSessionInfoResponse{},
		&dota.CMsgGCRequestSubGCSessionInfo{},
		&dota.CMsgGCRoutingInfo{},
		&dota.CMsgGCServerVersionUpdated{},
		&dota.CMsgGCShowItemsPickedUp{},
		&dota.CMsgGCSteamProfileRequestResponse{},
		&dota.CMsgGCSteamProfileRequest{},
		&dota.CMsgGCStorePurchaseCancelResponse{},
		&dota.CMsgGCStorePurchaseCancel{},
		&dota.CMsgGCStorePurchaseFinalizeResponse{},
		&dota.CMsgGCStorePurchaseFinalize{},
		&dota.CMsgGCStorePurchaseInitResponse{},
		&dota.CMsgGCStorePurchaseInit{},
		&dota.CMsgGCToClientEmoticonData{},
		&dota.CMsgGCToClientTournamentItemDrop{},
		&dota.CMsgGCToClientTrophyAwarded{},
		&dota.CMsgGCToGCApplyLocalizationDiffResponse{},
		&dota.CMsgGCToGCApplyLocalizationDiff{},
		&dota.CMsgGCToGCBannedWordListBroadcast{},
		&dota.CMsgGCToGCBannedWordListUpdated{},
		&dota.CMsgGCToGCBroadcastConsoleCommand{},
		&dota.CMsgGCToGCCanUseDropRateBonus{},
		&dota.CMsgGCToGCCheckAccountTradeStatusResponse{},
		&dota.CMsgGCToGCCheckAccountTradeStatus{},
		&dota.CMsgGCToGCDirtyMultipleSDOCache{},
		&dota.CMsgGCToGCDirtySDOCache{},
		&dota.CMsgGCToGCGetUserPCBangNoResponse{},
		&dota.CMsgGCToGCGetUserPCBangNo{},
		&dota.CMsgGCToGCGetUserServerMembersResponse{},
		&dota.CMsgGCToGCGetUserServerMembers{},
		&dota.CMsgGCToGCGetUserSessionServerResponse{},
		&dota.CMsgGCToGCGetUserSessionServer{},
		&dota.CMsgGCToGCGrantAccountRolledItems_Item_AdditionalAuditEntry{},
		&dota.CMsgGCToGCGrantAccountRolledItems_Item_DynamicAttribute{},
		&dota.CMsgGCToGCGrantAccountRolledItems_Item{},
		&dota.CMsgGCToGCGrantAccountRolledItems{},
		&dota.CMsgGCToGCGrantSelfMadeItemToAccount{},
		&dota.CMsgGCToGCLoadSessionSOCacheResponse{},
		&dota.CMsgGCToGCLoadSessionSOCache{},
		&dota.CMsgGCToGCPingRequest{},
		&dota.CMsgGCToGCPingResponse{},
		&dota.CMsgGCToGCRefreshSOCache{},
		&dota.CMsgGCToGCSOCacheSubscribe_CMsgHaveVersions{},
		&dota.CMsgGCToGCSOCacheSubscribe{},
		&dota.CMsgGCToGCSOCacheUnsubscribe{},
		&dota.CMsgGCToGCUpdateSQLKeyValue{},
		&dota.CMsgGCToGCUpdateSessionStats{},
		&dota.CMsgGCToGCWebAPIAccountChanged{},
		&dota.CMsgGCToRelayConnectResponse{},
		&dota.CMsgGCToRelayConnect{},
		&dota.CMsgGCToServerConsoleCommand{},
		&dota.CMsgGCToServerPingRequest{},
		&dota.CMsgGCToServerPingResponse{},
		&dota.CMsgGCUpdateSubGCSessionInfo_CMsgUpdate{},
		&dota.CMsgGCUpdateSubGCSessionInfo{},
		&dota.CMsgGCWatchDownloadedReplay{},
		&dota.CMsgGameChatLog_CChatLine{},
		&dota.CMsgGameChatLog{},
		&dota.CMsgGameMatchSignOutPermissionRequest{},
		&dota.CMsgGameMatchSignOutPermissionResponse{},
		&dota.CMsgGameMatchSignOut_CAdditionalSignoutMsg{},
		&dota.CMsgGameMatchSignOut_CTeam_CPlayer{},
		&dota.CMsgGameMatchSignOut_CTeam{},
		&dota.CMsgGameMatchSignOut{},
		&dota.CMsgGameMatchSignoutResponse_CAdditionalSignoutMsg{},
		&dota.CMsgGameMatchSignoutResponse{},
		&dota.CMsgGameServerGetLoadGameResult{},
		&dota.CMsgGameServerGetLoadGame{},
		&dota.CMsgGameServerInfo{},
		&dota.CMsgGameServerSaveGameResult{},
		&dota.CMsgGameServerUploadSaveGame{},
		&dota.CMsgGetZoomResponse{},
		&dota.CMsgGetZoom{},
		&dota.CMsgGoBack{},
		&dota.CMsgGoForward{},
		&dota.CMsgGuildmatePracticeLobbyListRequest{},
		&dota.CMsgGuildmatePracticeLobbyListResponse{},
		&dota.CMsgHeroPickStatPlayer{},
		&dota.CMsgHeroPickStatsRequest{},
		&dota.CMsgHeroPickStatsResponse{},
		&dota.CMsgHidePopup{},
		&dota.CMsgHideToolTip{},
		&dota.CMsgHorizontalScrollBarSizeResponse{},
		&dota.CMsgHorizontalScrollBarSize{},
		&dota.CMsgHttpRequest_QueryParam{},
		&dota.CMsgHttpRequest_RequestHeader{},
		&dota.CMsgHttpRequest{},
		&dota.CMsgHttpResponse_ResponseHeader{},
		&dota.CMsgHttpResponse{},
		&dota.CMsgIPCAddress{},
		&dota.CMsgInitialQuestionnaireResponse{},
		&dota.CMsgInvitationCreated{},
		&dota.CMsgInviteToParty{},
		&dota.CMsgItemAcknowledged{},
		&dota.CMsgJSAlert{},
		&dota.CMsgJSConfirm{},
		&dota.CMsgJSDialogResponse{},
		&dota.CMsgJoinableCustomGameModesRequest{},
		&dota.CMsgJoinableCustomGameModesResponseEntry{},
		&dota.CMsgJoinableCustomGameModesResponse{},
		&dota.CMsgJoinableCustomLobbiesRequest{},
		&dota.CMsgJoinableCustomLobbiesResponseEntry{},
		&dota.CMsgJoinableCustomLobbiesResponse{},
		&dota.CMsgKeyChar{},
		&dota.CMsgKeyDown{},
		&dota.CMsgKeyUp{},
		&dota.CMsgKickFromParty{},
		&dota.CMsgLANServerAvailable{},
		&dota.CMsgLeagueAdminList{},
		&dota.CMsgLeagueScheduleBlockTeamInfo{},
		&dota.CMsgLeagueScheduleBlock{},
		&dota.CMsgLeaveParty{},
		&dota.CMsgLeaverDetectedResponse{},
		&dota.CMsgLeaverDetected{},
		&dota.CMsgLeaverState{},
		&dota.CMsgLinkAtPositionResponse{},
		&dota.CMsgLinkAtPosition{},
		&dota.CMsgLoadingResource{},
		&dota.CMsgLookupMultipleAccountNamesResponse_Account{},
		&dota.CMsgLookupMultipleAccountNamesResponse{},
		&dota.CMsgLookupMultipleAccountNames{},
		&dota.CMsgMakeOffering{},
		&dota.CMsgMatchVoteResponse{},
		&dota.CMsgMatchmakingGroupServerSample{},
		&dota.CMsgMouseDblClick{},
		&dota.CMsgMouseDown{},
		&dota.CMsgMouseMove{},
		&dota.CMsgMouseUp{},
		&dota.CMsgMouseWheel{},
		&dota.CMsgNeedsPaintResponse{},
		&dota.CMsgNeedsPaint{},
		&dota.CMsgNexonPartnerUpdate{},
		&dota.CMsgNotificationOfSuspiciousActivity_MultipleGameInstances{},
		&dota.CMsgNotificationOfSuspiciousActivity{},
		&dota.CMsgNotifyWatchdog{},
		&dota.CMsgOpenNewTabResponse{},
		&dota.CMsgOpenNewTab{},
		&dota.CMsgOpenSteamURL{},
		&dota.CMsgPackageLicense{},
		&dota.CMsgPartyInviteResponse{},
		&dota.CMsgPartyLeaderWatchGamePrompt{},
		&dota.CMsgPassportDataRequest{},
		&dota.CMsgPassportDataResponse{},
		&dota.CMsgPaste{},
		&dota.CMsgPerfectWorldUserLookupRequest{},
		&dota.CMsgPerfectWorldUserLookupResponse{},
		&dota.CMsgPlaceDecalEvent{},
		&dota.CMsgPopupHTMLWindowResponse{},
		&dota.CMsgPopupHTMLWindow{},
		&dota.CMsgPostURL{},
		&dota.CMsgPracticeLobbyCreate_SaveGame{},
		&dota.CMsgPracticeLobbyCreate{},
		&dota.CMsgPracticeLobbyJoinBroadcastChannel{},
		&dota.CMsgPracticeLobbyJoinResponse{},
		&dota.CMsgPracticeLobbyJoin{},
		&dota.CMsgPracticeLobbyKick{},
		&dota.CMsgPracticeLobbyLaunch{},
		&dota.CMsgPracticeLobbyLeave{},
		&dota.CMsgPracticeLobbyListResponseEntry_CLobbyMember{},
		&dota.CMsgPracticeLobbyListResponseEntry{},
		&dota.CMsgPracticeLobbyListResponse{},
		&dota.CMsgPracticeLobbyList{},
		&dota.CMsgPracticeLobbySetCoach{},
		&dota.CMsgPracticeLobbySetDetails{},
		&dota.CMsgPracticeLobbySetTeamSlot{},
		&dota.CMsgPracticeLobbyToggleBroadcastChannelCameramanStatus{},
		&dota.CMsgPresentedClientTerminateDlg{},
		&dota.CMsgProtoBufHeader{},
		&dota.CMsgQAngle{},
		&dota.CMsgQuickJoinCustomLobbyResponse{},
		&dota.CMsgQuickJoinCustomLobby{},
		&dota.CMsgReadyUpStatus{},
		&dota.CMsgReadyUp{},
		&dota.CMsgRecipeComponent{},
		&dota.CMsgRedeemCodeResponse{},
		&dota.CMsgRedeemCode{},
		&dota.CMsgRefreshPartnerAccountLink{},
		&dota.CMsgReload{},
		&dota.CMsgReplayUploadedToYouTube{},
		&dota.CMsgReplicateConVars{},
		&dota.CMsgRequestCrateItemsResponse{},
		&dota.CMsgRequestCrateItems{},
		&dota.CMsgRequestInternationalTicket{},
		&dota.CMsgRequestInventoryRefresh{},
		&dota.CMsgRequestItemPurgatory_FinalizePurchaseResponse{},
		&dota.CMsgRequestItemPurgatory_FinalizePurchase{},
		&dota.CMsgRequestItemPurgatory_RefundPurchaseResponse{},
		&dota.CMsgRequestItemPurgatory_RefundPurchase{},
		&dota.CMsgRequestLeagueInfo{},
		&dota.CMsgRequestLeaguePrizePoolResponse{},
		&dota.CMsgRequestLeaguePrizePool{},
		&dota.CMsgRequestOfferingsResponse_NewYearsOffering{},
		&dota.CMsgRequestOfferingsResponse{},
		&dota.CMsgRequestOfferings{},
		&dota.CMsgRequestWeekendTourneySchedule{},
		&dota.CMsgResetMapLocationsResponse{},
		&dota.CMsgResetMapLocations{},
		&dota.CMsgResetStrangeGemCount{},
		&dota.CMsgResizeResponse{},
		&dota.CMsgResourceResponse{},
		&dota.CMsgResponseLeagueInfo{},
		&dota.CMsgResponseTeamFanfare{},
		&dota.CMsgRetrieveMatchVote{},
		&dota.CMsgSDONoMemcached{},
		&dota.CMsgSHA1Digest{},
		&dota.CMsgSOCacheHaveVersion{},
		&dota.CMsgSOCacheSubscribedUpToDate{},
		&dota.CMsgSOCacheSubscribed_SubscribedType{},
		&dota.CMsgSOCacheSubscribed{},
		&dota.CMsgSOCacheSubscriptionCheck{},
		&dota.CMsgSOCacheSubscriptionRefresh{},
		&dota.CMsgSOCacheUnsubscribed{},
		&dota.CMsgSOCacheVersion{},
		&dota.CMsgSOIDOwner{},
		&dota.CMsgSOMultipleObjects_SingleObject{},
		&dota.CMsgSOMultipleObjects{},
		&dota.CMsgSOSingleObject{},
		&dota.CMsgSQLAddDropRateBonus{},
		&dota.CMsgSQLUpgradeBattleBooster{},
		&dota.CMsgSavePageToJPEG{},
		&dota.CMsgSearchResults{},
		&dota.CMsgSelectItemPresetForClassReply{},
		&dota.CMsgSelectItemPresetForClass{},
		&dota.CMsgSerializedSOCache_Cache_Version{},
		&dota.CMsgSerializedSOCache_Cache{},
		&dota.CMsgSerializedSOCache_TypeCache{},
		&dota.CMsgSerializedSOCache{},
		&dota.CMsgServerAvailable{},
		&dota.CMsgServerGCUpdateSpectatorCount{},
		&dota.CMsgServerGetEventPointsResponse_Points{},
		&dota.CMsgServerGetEventPointsResponse{},
		&dota.CMsgServerGetEventPoints{},
		&dota.CMsgServerGrantSurveyPermissionResponse{},
		&dota.CMsgServerGrantSurveyPermission_Survey{},
		&dota.CMsgServerGrantSurveyPermission{},
		&dota.CMsgServerPeer{},
		&dota.CMsgServerToGCGetAdditionalEquipsResponse_CUserEquips{},
		&dota.CMsgServerToGCGetAdditionalEquipsResponse{},
		&dota.CMsgServerToGCGetAdditionalEquips{},
		&dota.CMsgServerToGCGetProfileCardResponse{},
		&dota.CMsgServerToGCGetProfileCard{},
		&dota.CMsgServerToGCMatchConnectionStats_Player{},
		&dota.CMsgServerToGCMatchConnectionStats{},
		&dota.CMsgServerToGCRequestStatus_Response{},
		&dota.CMsgServerToGCRequestStatus{},
		&dota.CMsgServerToGCVictoryPredictions_Record{},
		&dota.CMsgServerToGCVictoryPredictions{},
		&dota.CMsgServerUseItem{},
		&dota.CMsgSetCookie{},
		&dota.CMsgSetCursor{},
		&dota.CMsgSetFeaturedItems{},
		&dota.CMsgSetFocus{},
		&dota.CMsgSetHTMLTitle{},
		&dota.CMsgSetHorizontalScroll{},
		&dota.CMsgSetItemPositions_ItemPosition{},
		&dota.CMsgSetItemPositions{},
		&dota.CMsgSetMapLocationStateResponse{},
		&dota.CMsgSetMapLocationState{},
		&dota.CMsgSetPresetItemPosition{},
		&dota.CMsgSetShowcaseHero{},
		&dota.CMsgSetTargetFrameRate{},
		&dota.CMsgSetVerticalScroll{},
		&dota.CMsgSetZoomLevel{},
		&dota.CMsgShowPopup{},
		&dota.CMsgShowToolTip{},
		&dota.CMsgSignOutBotInfo_CMsgBotSlotDifficulty{},
		&dota.CMsgSignOutBotInfo{},
		&dota.CMsgSignOutDraftInfo{},
		&dota.CMsgSizePopup{},
		&dota.CMsgSockAddrList{},
		&dota.CMsgSortItems{},
		&dota.CMsgSosSetLibraryStackField{},
		&dota.CMsgSosSetSoundEventParam{},
		&dota.CMsgSosStartSoundEvent{},
		&dota.CMsgSosStopSoundEventHash{},
		&dota.CMsgSosStopSoundEvent{},
		&dota.CMsgSource1LegacyGameEventList{},
		&dota.CMsgSource1LegacyGameEvent{},
		&dota.CMsgSource1LegacyListenEvents{},
		&dota.CMsgSourceTVGamesResponse{},
		&dota.CMsgSpawnLootGreevil{},
		&dota.CMsgSpectateFriendGameResponse{},
		&dota.CMsgSpectateFriendGame{},
		&dota.CMsgStartFindingMatch{},
		&dota.CMsgStartRequestResponse{},
		&dota.CMsgStartRequest{},
		&dota.CMsgStatusText{},
		&dota.CMsgStopFindingMatch{},
		&dota.CMsgStopFind{},
		&dota.CMsgStopLoad{},
		&dota.CMsgStoreGetUserDataResponse{},
		&dota.CMsgStoreGetUserData{},
		&dota.CMsgSuspiciousActivity{},
		&dota.CMsgSystemBroadcast{},
		&dota.CMsgTEArmorRicochet{},
		&dota.CMsgTEBSPDecal{},
		&dota.CMsgTEBaseBeam{},
		&dota.CMsgTEBeamEntPoint{},
		&dota.CMsgTEBeamEnts{},
		&dota.CMsgTEBeamPoints{},
		&dota.CMsgTEBeamRing{},
		&dota.CMsgTEBloodStream{},
		&dota.CMsgTEBreakModel{},
		&dota.CMsgTEBubbleTrail{},
		&dota.CMsgTEBubbles{},
		&dota.CMsgTEDecal{},
		&dota.CMsgTEDust{},
		&dota.CMsgTEEffectDispatch{},
		&dota.CMsgTEEnergySplash{},
		&dota.CMsgTEExplosion{},
		&dota.CMsgTEFizz{},
		&dota.CMsgTEGlowSprite{},
		&dota.CMsgTEImpact{},
		&dota.CMsgTELargeFunnel{},
		&dota.CMsgTEMuzzleFlash{},
		&dota.CMsgTEPhysicsProp{},
		&dota.CMsgTEPlayerDecal{},
		&dota.CMsgTEProjectedDecal{},
		&dota.CMsgTEShatterSurface{},
		&dota.CMsgTESmoke{},
		&dota.CMsgTESparks{},
		&dota.CMsgTEWorldDecal{},
		&dota.CMsgTeamFanfare{},
		&dota.CMsgTournamentItemEventResponse{},
		&dota.CMsgTournamentItemEvent{},
		&dota.CMsgURLChanged{},
		&dota.CMsgUpdateItemSchema{},
		&dota.CMsgUpdateToolTip{},
		&dota.CMsgUpgradeLeagueItemResponse{},
		&dota.CMsgUpgradeLeagueItem{},
		&dota.CMsgUseItem{},
		&dota.CMsgVDebugGameSessionIDEvent{},
		&dota.CMsgVector2D{},
		&dota.CMsgVector{},
		&dota.CMsgVerticalScrollBarSizeResponse{},
		&dota.CMsgVerticalScrollBarSize{},
		&dota.CMsgViewSource{},
		&dota.CMsgWatchGameResponse{},
		&dota.CMsgWatchGame{},
		&dota.CMsgWebAPIKey{},
		&dota.CMsgWebAPIRequest{},
		&dota.CMsgWeekendTourneySchedule_Division{},
		&dota.CMsgWeekendTourneySchedule{},
		&dota.CMsgZoomToElementAtPositionResponse{},
		&dota.CMsgZoomToElementAtPosition{},
		&dota.CMsg_CVars_CVar{},
		&dota.CMsg_CVars{},
		&dota.CNETMsg_Disconnect{},
		&dota.CNETMsg_File{},
		&dota.CNETMsg_NOP{},
		&dota.CNETMsg_SetConVar{},
		&dota.CNETMsg_SignonState{},
		&dota.CNETMsg_SpawnGroup_ForceBlockingLoad{},
		&dota.CNETMsg_SpawnGroup_LoadCompleted{},
		&dota.CNETMsg_SpawnGroup_Load{},
		&dota.CNETMsg_SpawnGroup_ManifestUpdate{},
		&dota.CNETMsg_SpawnGroup_SetCreationTick{},
		&dota.CNETMsg_SpawnGroup_Unload{},
		&dota.CNETMsg_SplitScreenUser{},
		&dota.CNETMsg_StringCmd{},
		&dota.CNETMsg_Tick{},
		&dota.CP2P_TextMessage{},
		&dota.CP2P_Voice{},
		&dota.CProtoItemSocket_AbilityEffect{},
		&dota.CProtoItemSocket_AssetModifier{},
		&dota.CProtoItemSocket_Autograph{},
		&dota.CProtoItemSocket_Color{},
		&dota.CProtoItemSocket_Effect{},
		&dota.CProtoItemSocket_Empty{},
		&dota.CProtoItemSocket_Spectator{},
		&dota.CProtoItemSocket_StaticVisuals{},
		&dota.CProtoItemSocket_Strange{},
		&dota.CProtoItemSocket{},
		&dota.CSODOTAGameAccountClient{},
		&dota.CSODOTAGameHeroFavorites{},
		&dota.CSODOTALobby_CExtraMsg{},
		&dota.CSODOTALobby{},
		&dota.CSODOTAMapLocationState{},
		&dota.CSODOTAPartyInvite_PartyMember{},
		&dota.CSODOTAPartyInvite{},
		&dota.CSODOTAPartyMember{},
		&dota.CSODOTAParty{},
		&dota.CSOEconClaimCode{},
		&dota.CSOEconGameAccountClient{},
		&dota.CSOEconItemAttribute{},
		&dota.CSOEconItemDropRateBonus{},
		&dota.CSOEconItemEquipped{},
		&dota.CSOEconItemEventTicket{},
		&dota.CSOEconItemLeagueViewPass{},
		&dota.CSOEconItemPresetInstance{},
		&dota.CSOEconItemTournamentPassport{},
		&dota.CSOEconItem{},
		&dota.CSOItemCriteriaCondition{},
		&dota.CSOItemCriteria{},
		&dota.CSOItemRecipe{},
		&dota.CSOLobbyInvite{},
		&dota.CSOPartyInvite{},
		&dota.CSOSelectedItemPreset{},
		&dota.CSVCMsgList_GameEventsEventT{},
		&dota.CSVCMsgList_GameEvents{},
		&dota.CSVCMsgList_UserMessagesUsermsgT{},
		&dota.CSVCMsgList_UserMessages{},
		&dota.CSVCMsg_BSPDecal{},
		&dota.CSVCMsg_ClassInfoClassT{},
		&dota.CSVCMsg_ClassInfo{},
		&dota.CSVCMsg_ClearAllStringTables{},
		&dota.CSVCMsg_CmdKeyValues{},
		&dota.CSVCMsg_CreateStringTable{},
		&dota.CSVCMsg_CrosshairAngle{},
		&dota.CSVCMsg_FixAngle{},
		&dota.CSVCMsg_FlattenedSerializer{},
		&dota.CSVCMsg_FullFrameSplit{},
		&dota.CSVCMsg_GameEventKeyT{},
		&dota.CSVCMsg_GameEventListDescriptorT{},
		&dota.CSVCMsg_GameEventListKeyT{},
		&dota.CSVCMsg_GameEventList{},
		&dota.CSVCMsg_GameEvent{},
		&dota.CSVCMsg_GameSessionConfiguration{},
		&dota.CSVCMsg_GetCvarValue{},
		&dota.CSVCMsg_Menu{},
		&dota.CSVCMsg_PacketEntities{},
		&dota.CSVCMsg_PacketReliable{},
		&dota.CSVCMsg_PeerList{},
		&dota.CSVCMsg_Prefetch{},
		&dota.CSVCMsg_Print{},
		&dota.CSVCMsg_SendTableSendpropT{},
		&dota.CSVCMsg_SendTable{},
		&dota.CSVCMsg_ServerInfo{},
		&dota.CSVCMsg_SetPause{},
		&dota.CSVCMsg_SetView{},
		&dota.CSVCMsg_SoundsSounddataT{},
		&dota.CSVCMsg_Sounds{},
		&dota.CSVCMsg_SplitScreen{},
		&dota.CSVCMsg_StopSound{},
		&dota.CSVCMsg_TempEntities{},
		&dota.CSVCMsg_UpdateStringTable{},
		&dota.CSVCMsg_UserMessage{},
		&dota.CSVCMsg_VoiceData{},
		&dota.CSVCMsg_VoiceInit{},
		&dota.CSerializedCombatLog_Dictionary_DictString{},
		&dota.CSerializedCombatLog_Dictionary{},
		&dota.CSerializedCombatLog{},
		&dota.CSourceTVGame_Player{},
		&dota.CSourceTVGame{},
		&dota.CSteam_Voice_Encoding{},
		&dota.CUserMessageAchievementEvent{},
		&dota.CUserMessageAmmoDenied{},
		&dota.CUserMessageCameraTransition_Transition_DataDriven{},
		&dota.CUserMessageCameraTransition{},
		&dota.CUserMessageCloseCaptionDirect{},
		&dota.CUserMessageCloseCaptionPlaceholder{},
		&dota.CUserMessageCloseCaption{},
		&dota.CUserMessageColoredText{},
		&dota.CUserMessageCreditsMsg{},
		&dota.CUserMessageCrosshairAngle{},
		&dota.CUserMessageCurrentTimescale{},
		&dota.CUserMessageDesiredTimescale{},
		&dota.CUserMessageFade{},
		&dota.CUserMessageGameTitle{},
		&dota.CUserMessageHintText{},
		&dota.CUserMessageHudMsg{},
		&dota.CUserMessageHudText{},
		&dota.CUserMessageItemPickup{},
		&dota.CUserMessageKeyHintText{},
		&dota.CUserMessageRequestState{},
		&dota.CUserMessageResetHUD{},
		&dota.CUserMessageRumble{},
		&dota.CUserMessageSayText2{},
		&dota.CUserMessageSayTextChannel{},
		&dota.CUserMessageSayText{},
		&dota.CUserMessageScreenTilt{},
		&dota.CUserMessageSendAudio{},
		&dota.CUserMessageShakeDir{},
		&dota.CUserMessageShake{},
		&dota.CUserMessageShowMenu{},
		&dota.CUserMessageTextMsg{},
		&dota.CUserMessageTrain{},
		&dota.CUserMessageVGUIMenu_Keys{},
		&dota.CUserMessageVGUIMenu{},
		&dota.CUserMessageVoiceMask{},
		&dota.CUserMessageVoiceSubtitle{},
		&dota.CWorkshop_GetContributors_Request{},
		&dota.CWorkshop_GetContributors_Response{},
		&dota.CWorkshop_PopulateItemDescriptions_Request_ItemDescriptionsLanguageBlock{},
		&dota.CWorkshop_PopulateItemDescriptions_Request_SingleItemDescription{},
		&dota.CWorkshop_PopulateItemDescriptions_Request{},
		&dota.CWorkshop_SetItemPaymentRules_Request_PartnerItemPaymentRule{},
		&dota.CWorkshop_SetItemPaymentRules_Request_WorkshopItemPaymentRule{},
		&dota.CWorkshop_SetItemPaymentRules_Request{},
		&dota.CWorkshop_SetItemPaymentRules_Response{},
		&dota.ProtoFlattenedSerializerFieldT{},
		&dota.ProtoFlattenedSerializerT{},
	}
}

func dbgenum(n uint64) {
	dbgenumcheck(dota.Bidirectional_Messages(n))
	dbgenumcheck(dota.CLC_Messages(n))
	dbgenumcheck(dota.ConnectionLessMessageIds(n))
	dbgenumcheck(dota.DIALOG_TYPE(n))
	dbgenumcheck(dota.DOTABotDifficulty(n))
	dbgenumcheck(dota.DOTAGameVersion(n))
	dbgenumcheck(dota.DOTAJoinLobbyResult(n))
	dbgenumcheck(dota.DOTALobbyReadyState(n))
	dbgenumcheck(dota.DOTALowPriorityBanType(n))
	dbgenumcheck(dota.DOTAMatchVote(n))
	dbgenumcheck(dota.DOTA_2013PassportSelectionIndices(n))
	dbgenumcheck(dota.DOTA_ABILITY_PING_TYPE(n))
	dbgenumcheck(dota.DOTA_CHAT_INFORMATIONAL(n))
	dbgenumcheck(dota.DOTA_CHAT_MESSAGE(n))
	dbgenumcheck(dota.DOTA_CM_PICK(n))
	dbgenumcheck(dota.DOTA_COMBATLOG_TYPES(n))
	dbgenumcheck(dota.DOTA_GC_TEAM(n))
	dbgenumcheck(dota.DOTA_GameMode(n))
	dbgenumcheck(dota.DOTA_GameState(n))
	dbgenumcheck(dota.DOTA_LobbyMemberXPBonus(n))
	dbgenumcheck(dota.DOTA_MODIFIER_ENTRY_TYPE(n))
	dbgenumcheck(dota.DOTA_NO_BATTLE_POINTS_REASONS(n))
	dbgenumcheck(dota.DOTA_OVERHEAD_ALERT(n))
	dbgenumcheck(dota.DOTA_PARTICLE_MESSAGE(n))
	dbgenumcheck(dota.DOTA_TournamentEvents(n))
	dbgenumcheck(dota.DOTA_WatchReplayType(n))
	dbgenumcheck(dota.EAbilityAbuseType(n))
	dbgenumcheck(dota.EBaseEntityMessages(n))
	dbgenumcheck(dota.EBaseGameEvents(n))
	dbgenumcheck(dota.EBaseUserMessages(n))
	dbgenumcheck(dota.EDOTAChatWheelMessage(n))
	dbgenumcheck(dota.EDOTAGCMsg(n))
	dbgenumcheck(dota.EDOTAGCSessionNeed(n))
	dbgenumcheck(dota.EDOTAPlayerMMRType(n))
	dbgenumcheck(dota.EDOTAStatPopupTypes(n))
	dbgenumcheck(dota.EDemoCommands(n))
	dbgenumcheck(dota.EDotaBroadcastMessages(n))
	dbgenumcheck(dota.EDotaClientMessages(n))
	dbgenumcheck(dota.EDotaEntityMessages(n))
	dbgenumcheck(dota.EDotaUserMessages(n))
	dbgenumcheck(dota.EGCBaseClientMsg(n))
	dbgenumcheck(dota.EGCBaseMsg(n))
	dbgenumcheck(dota.EGCBaseProtoObjectTypes(n))
	dbgenumcheck(dota.EGCItemMsg(n))
	dbgenumcheck(dota.EGCMsgResponse(n))
	dbgenumcheck(dota.EGCPartnerRequestResponse(n))
	dbgenumcheck(dota.EGCSystemMsg(n))
	dbgenumcheck(dota.EGCToGCMsg(n))
	dbgenumcheck(dota.EIntentionalFeedingType(n))
	dbgenumcheck(dota.EItemEditorReservationResult(n))
	dbgenumcheck(dota.EItemPurgatoryResponse_Finalize(n))
	dbgenumcheck(dota.EItemPurgatoryResponse_Refund(n))
	dbgenumcheck(dota.ENetworkDisconnectionReason(n))
	dbgenumcheck(dota.EQueryCvarValueStatus(n))
	dbgenumcheck(dota.ESOMsg(n))
	dbgenumcheck(dota.ESourceEngine(n))
	dbgenumcheck(dota.ESplitScreenMessageType(n))
	dbgenumcheck(dota.ESuspiciousActivity(n))
	dbgenumcheck(dota.ESuspiciousBuildType(n))
	dbgenumcheck(dota.ETEProtobufIds(n))
	dbgenumcheck(dota.ETournamentGameState(n))
	dbgenumcheck(dota.ETournamentNodeState(n))
	dbgenumcheck(dota.ETournamentState(n))
	dbgenumcheck(dota.ETournamentTeamState(n))
	dbgenumcheck(dota.ETournamentTemplate(n))
	dbgenumcheck(dota.ETournamentType(n))
	dbgenumcheck(dota.EUnlockStyle(n))
	dbgenumcheck(dota.Fantasy_Roles(n))
	dbgenumcheck(dota.Fantasy_Selection_Mode(n))
	dbgenumcheck(dota.Fantasy_Team_Slots(n))
	dbgenumcheck(dota.GCConnectionStatus(n))
	dbgenumcheck(dota.GCProtoBufMsgSrc(n))
	dbgenumcheck(dota.GC_BannedWordType(n))
	dbgenumcheck(dota.LobbyDotaTVDelay(n))
	dbgenumcheck(dota.MatchLanguages(n))
	dbgenumcheck(dota.MatchType(n))
	dbgenumcheck(dota.NET_Messages(n))
	dbgenumcheck(dota.P2P_Messages(n))
	dbgenumcheck(dota.PartnerAccountType(n))
	dbgenumcheck(dota.PrefetchType(n))
	dbgenumcheck(dota.SVC_Messages(n))
	dbgenumcheck(dota.SVC_Messages_LowFrequency(n))
}

func dbgenumcheck(n interface{}) {
	f := spew.Sprintf("%s", n)
	if spew.Sprintf("%d", n) == f {
	} else {
		PP(f, n)
	}
}
