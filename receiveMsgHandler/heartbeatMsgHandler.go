package receiveMsgHandler

import (
	"time"

	"BulletRain_server/msgProto"
)

func (h *MsgHandler) MsgPing(info *ClientInfo, msgPing *msgProto.MsgPing) {
	player := info.Player
	if player == nil {
		return
	}
	msgPong := &msgProto.MsgPong{
		ProtoName: "MsgPong",
	}
	info.LastPingTime = time.Now()
	Send(info, msgPong)
}
