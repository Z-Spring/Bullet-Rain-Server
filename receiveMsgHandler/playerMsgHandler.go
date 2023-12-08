package receiveMsgHandler

import (
	"errors"
	"log"

	"BulletRain_server/msgProto"
)

const DAMAGE = 10

func (h *MsgHandler) MsgSyncPlayer(info *ClientInfo, msgSyncPlayer *msgProto.MsgSyncPlayer) {
	player := info.Player
	if player == nil {
		return
	}
	room, err := checkRoomInfos(player)
	if err != nil {
		//log.Println(err)
		return
	}

	player.X = msgSyncPlayer.X
	player.Y = msgSyncPlayer.Y
	player.Z = msgSyncPlayer.Z
	player.EX = msgSyncPlayer.EX
	player.EY = msgSyncPlayer.EY
	player.EZ = msgSyncPlayer.EZ

	msgSyncPlayer.Id = player.Id
	room.broadcastMsgToAllPlayers(msgSyncPlayer)
}

func (h *MsgHandler) MsgFire(info *ClientInfo, msgFire *msgProto.MsgFire) {
	player := info.Player
	if player == nil {
		return
	}
	room, err := checkRoomInfos(player)
	if err != nil {
		log.Println(err)
		return
	}

	msgFire.Id = player.Id
	room.broadcastMsgToAllPlayers(msgFire)
}

// MsgHit todo: 发送一条回复两条？ 还是有其他的也发送过来了消息？
func (h *MsgHandler) MsgHit(info *ClientInfo, msgHit *msgProto.MsgHit) {
	player := info.Player
	if player == nil {
		return
	}

	targetPlayer := Pm.GetPlayer(msgHit.TargetId)
	if targetPlayer == nil {
		return
	}

	room, err := checkRoomInfos(player)
	if err != nil {
		log.Println(err)
		return
	}

	if player.Id != msgHit.Id {
		return
	}

	damage := DAMAGE

	targetPlayer.HP -= damage

	msgHit.Hp = player.HP
	msgHit.Id = player.Id
	msgHit.Damage = damage

	room.broadcastMsgToAllPlayers(msgHit)
}

func checkRoomInfos(player *Player) (*Room, error) {
	room := RM.GetRoom(player.RoomId)
	if room == nil {
		return nil, errors.New("room is nil")
	}

	if room.Status != Fighting {
		return nil, errors.New("room status is not fighting")
	}

	return room, nil
}
