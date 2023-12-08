package msg

import (
	"encoding/json"
	"log"
	"reflect"

	"BulletRain_server/msgProto"
)

type MsgBase interface {
	GetProtoName() string
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
func Encode(msgBase MsgBase) []byte {
	s, err := json.Marshal(msgBase)
	if err != nil {
		log.Println("json marshal fail")
	}
	return s
}

var nameToType = map[string]reflect.Type{
	"MsgPing": reflect.TypeOf((*msgProto.MsgPing)(nil)).Elem(),
	"MsgPong": reflect.TypeOf((*msgProto.MsgPong)(nil)).Elem(),

	"MsgLogin":    reflect.TypeOf((*msgProto.MsgLogin)(nil)).Elem(),
	"MsgRegister": reflect.TypeOf((*msgProto.MsgRegister)(nil)).Elem(),
	"MsgKick":     reflect.TypeOf((*msgProto.MsgKick)(nil)).Elem(),

	"MsgEnterBattle":  reflect.TypeOf((*msgProto.MsgEnterBattle)(nil)).Elem(),
	"MsgStartBattle":  reflect.TypeOf((*msgProto.MsgStartBattle)(nil)).Elem(),
	"MsgLeaveBattle":  reflect.TypeOf((*msgProto.MsgLeaveBattle)(nil)).Elem(),
	"MsgBattleResult": reflect.TypeOf((*msgProto.MsgBattleResult)(nil)).Elem(),

	"MsgEnterRoom":   reflect.TypeOf((*msgProto.MsgEnterRoom)(nil)).Elem(),
	"MsgLeaveRoom":   reflect.TypeOf((*msgProto.MsgLeaveRoom)(nil)).Elem(),
	"MsgGetRoomList": reflect.TypeOf((*msgProto.MsgGetRoomList)(nil)).Elem(),
	"MsgCreateRoom":  reflect.TypeOf((*msgProto.MsgCreateRoom)(nil)).Elem(),
	"MsgGetAchieve":  reflect.TypeOf((*msgProto.MsgGetAchieve)(nil)).Elem(),
	"MsgGetRoomInfo": reflect.TypeOf((*msgProto.MsgGetRoomInfo)(nil)).Elem(),

	"MsgSyncPlayer": reflect.TypeOf((*msgProto.MsgSyncPlayer)(nil)).Elem(),
	"MsgFire":       reflect.TypeOf((*msgProto.MsgFire)(nil)).Elem(),
	"MsgHit":        reflect.TypeOf((*msgProto.MsgHit)(nil)).Elem(),

	"MsgResult":             reflect.TypeOf((*msgProto.MsgResult)(nil)).Elem(),
	"MsgGetGameTime":        reflect.TypeOf((*msgProto.MsgGetGameTime)(nil)).Elem(),
	"MsgSwitchWeapon":       reflect.TypeOf((*msgProto.MsgSwitchWeapon)(nil)).Elem(),
	"MsgSyncWeaponPosition": reflect.TypeOf((*msgProto.MsgSyncWeaponPosition)(nil)).Elem(),
}

func Decode(protoName string, bytes []byte, beginIndex, count int) MsgBase {
	newBytes := bytes[beginIndex : count+beginIndex]
	typeObj, ok := nameToType[protoName]
	if !ok {
		log.Printf("protoName [%s] is not exist\n", protoName)
		return nil
	}
	msgBase := reflect.New(typeObj).Interface().(MsgBase)
	if err := json.Unmarshal(newBytes, &msgBase); err != nil {
		log.Println("json unmarshal fail")
	}

	return msgBase

}

func EncodeName(msgBase MsgBase) []byte {

	s := []byte(msgBase.GetProtoName())
	byteLength := int16(len(s))

	bytes := make([]byte, 0, byteLength+2)
	bytes = append(bytes, byte(byteLength%256))
	bytes = append(bytes, byte(byteLength/256))
	bytes = append(bytes, s...)

	return bytes
}

func DecodeName(bytes []byte, beginIndex int, count *int) string {
	*count = 0
	if beginIndex+2 > len(bytes) {
		return ""
	}

	byteLength := int(uint(bytes[beginIndex+1])<<8 | uint(bytes[beginIndex]))
	if beginIndex+2+byteLength > len(bytes) {
		return ""
	}
	*count = 2 + byteLength
	name := string(bytes[beginIndex+2 : beginIndex+2+byteLength])
	return name
}
