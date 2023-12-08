package msg

import (
	"BulletRain_server/msgProto"
)

var nameToType2 = map[string]MsgBase{
	"MsgPing": &msgProto.MsgPing{},
	"MsgPong": &msgProto.MsgPong{},

	"MsgLogin":    &msgProto.MsgLogin{},
	"MsgRegister": &msgProto.MsgRegister{},
	"MsgKick":     &msgProto.MsgKick{},

	"MsgEnterBattle":  &msgProto.MsgEnterBattle{},
	"MsgStartBattle":  &msgProto.MsgStartBattle{},
	"MsgLeaveBattle":  &msgProto.MsgLeaveBattle{},
	"MsgBattleResult": &msgProto.MsgBattleResult{},

	"MsgEnterRoom":   &msgProto.MsgEnterRoom{},
	"MsgLeaveRoom":   &msgProto.MsgLeaveRoom{},
	"MsgGetRoomList": &msgProto.MsgGetRoomList{},
	"MsgCreateRoom":  &msgProto.MsgCreateRoom{},
	"MsgGetAchieve":  &msgProto.MsgGetAchieve{},
	"MsgGetRoomInfo": &msgProto.MsgGetRoomInfo{},

	"MsgSyncPlayer": &msgProto.MsgSyncPlayer{},
	"MsgFire":       &msgProto.MsgFire{},
	"MsgHit":        &msgProto.MsgHit{},

	"MsgResult":             &msgProto.MsgResult{},
	"MsgGetGameTime":        &msgProto.MsgGetGameTime{},
	"MsgSwitchWeapon":       &msgProto.MsgSwitchWeapon{},
	"MsgSyncWeaponPosition": &msgProto.MsgSyncWeaponPosition{},
}

/*func ConvertMsgBase(msgBase MsgBase) MsgBase {
	base := reflect.TypeOf(msgBase).Elem()
	base = msgBase.GetProtoName()
	return base

}*/
