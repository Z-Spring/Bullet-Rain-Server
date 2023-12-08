package receiveMsgHandler

import (
	"fmt"
	"log"

	"BulletRain_server/database"
	"BulletRain_server/msgProto"
)

var RM = &RoomManager{
	MaxId: 1,
	Rooms: make(map[int]*Room, 5),
}

func (h *MsgHandler) MsgGetAchieve(info *ClientInfo, msgGetAchieve *msgProto.MsgGetAchieve) {
	player := info.Player
	if player == nil {
		return
	}

	playerInfo := database.GetPlayerData(player.Id)
	msgGetAchieve.Win = playerInfo.Win
	msgGetAchieve.Lost = playerInfo.Loss
	player.Send(msgGetAchieve)
}

func (h *MsgHandler) MsgGetRoomList(info *ClientInfo, msgGetRoomList *msgProto.MsgGetRoomList) {
	//log.Println("MsgGetRoomList")
	player := info.Player
	if player == nil {
		return
	}
	player.Send(RM.GetRoomList())
}

func (h *MsgHandler) MsgCreateRoom(info *ClientInfo, msgCreateRoom *msgProto.MsgCreateRoom) {
	player := info.Player
	if player == nil {
		return
	}
	// todo: roomId > 0 ???
	if IsInRoom(player) {
		msgCreateRoom.Result = 1
		msgCreateRoom.Reason = fmt.Sprintf("player %s is already in room", player.Id)
		log.Printf("player %s is already in room\n", player.Id)
		player.Send(msgCreateRoom)
		return
	}
	room := RM.AddRoom()
	room.AddPlayer(player.Id)
	msgCreateRoom.Result = 0
	player.Send(msgCreateRoom)
}

func (h *MsgHandler) MsgEnterRoom(info *ClientInfo, msgEnterRoom *msgProto.MsgEnterRoom) {
	player := info.Player
	if player == nil {
		return
	}
	if IsInRoom(player) {
		msgEnterRoom.Result = 1
		player.Send(msgEnterRoom)
		return
	}
	room := RM.GetRoom(msgEnterRoom.Id)
	if room == nil {
		msgEnterRoom.Result = 1
		player.Send(msgEnterRoom)
		return
	}
	if !room.AddPlayer(player.Id) {
		msgEnterRoom.Result = 1
		player.Send(msgEnterRoom)
		return
	}
	msgEnterRoom.Result = 0
	player.Send(msgEnterRoom)

}

func (h *MsgHandler) MsgLeaveRoom(info *ClientInfo, msgLeaveRoom *msgProto.MsgLeaveRoom) {
	player := info.Player
	if player == nil {
		return
	}
	if !IsInRoom(player) {
		msgLeaveRoom.Result = 1
		player.Send(msgLeaveRoom)
		return
	}
	room := RM.GetRoom(player.RoomId)
	if room == nil {
		msgLeaveRoom.Result = 1
		player.Send(msgLeaveRoom)
		return
	}

	room.RemovePlayer(player.Id)
	msgLeaveRoom.Result = 0
	player.Send(msgLeaveRoom)
}

func (h *MsgHandler) MsgGetRoomInfo(info *ClientInfo, roomInfo *msgProto.MsgGetRoomInfo) {
	player := info.Player
	if player == nil {
		return
	}
	room := RM.GetRoom(player.RoomId)
	if room == nil {
		player.Send(roomInfo)
		return
	}

	player.Send(room.GetRoomInfo())
}

func IsInRoom(player *Player) bool {
	return player.RoomId >= 0
}
