package receiveMsgHandler

import (
	"net"
	"time"

	"BulletRain_server/msg"
)

type ClientInfo struct {
	Conn net.Conn
	msg.ByteArray
	LastPingTime time.Time
	Player       *Player
}

type MsgHandler struct{}
