package receiveMsgHandler

import (
	"log"

	"BulletRain_server/databaseModel"
	"BulletRain_server/msg"
)

type Player struct {
	Id         string
	X, Y, Z    float32
	EX, EY, EZ float32
	HP         int
	RoomId     int // -1 means not in room
	Camp       int
	ClientInfo *ClientInfo
	databaseModel.PlayerData
}

func (p *Player) Send(base msg.MsgBase) {
	Send(p.ClientInfo, base)
}

func Send(clientInfo *ClientInfo, base msg.MsgBase) {
	if clientInfo == nil {
		return
	}
	// todo: modify
	//if clientInfo.Conn != nil {
	//	return
	//}
	base = msg.ConvertMsgBase(base)

	name := msg.EncodeName(base)
	if string(name) == "MsgResult" {
		log.Println("MsgResult")
	}

	body := msg.Encode(base)
	msgLength := len(name) + len(body)
	sendMsg := make([]byte, msgLength+2)

	sendMsg[0] = byte(msgLength % 256)
	sendMsg[1] = byte(msgLength / 256)
	copy(sendMsg[2:], name)
	copy(sendMsg[2+len(name):], body)

	_, err := clientInfo.Conn.Write(sendMsg)
	if err != nil {
		log.Println(err)
		clientInfo.Conn.Close()
		return
	}
}
