package receiveMsgHandler

import (
	"log"
	"time"

	"BulletRain_server/database"
	"BulletRain_server/databaseModel"
	"BulletRain_server/msg"
	"BulletRain_server/msgProto"
)

type Room struct {
	Id        int
	MaxPlayer int
	Players   map[string]bool
	OwnerId   string
	Status    int
}

const (
	Prepare = iota
	Fighting
)

var lastJudgeTime time.Time

// 阵营、每阵营人数、出生点
var playerBirthConfig = [2][3][6]float32{
	{
		{-10, 1, -25, 0, 90, 0},      //出生点1
		{-10, 1, -28, 0, 90, 0},      //出生点2
		{-7.533, 1, -27.61, 0, 0, 0}, //出生点3
	},
	{
		{16, 1, 33, 0, -90, 0},  //出生点1
		{14, 1, 33, 0, -180, 0}, //出生点2
		{14, 1, 29, 0, 0, 0},    //出生点3
	},
}

func setPlayerPos(player *Player, index int) {
	camp := player.Camp
	player.X = playerBirthConfig[camp-1][index][0]
	player.Y = playerBirthConfig[camp-1][index][1]
	player.Z = playerBirthConfig[camp-1][index][2]
	player.EX = playerBirthConfig[camp-1][index][3]
	player.EY = playerBirthConfig[camp-1][index][4]
	player.EZ = playerBirthConfig[camp-1][index][5]
}

func convertPlayerInfoToTankInfo(player *Player) msgProto.TankInfo {
	tankInfo := msgProto.TankInfo{
		Id:   player.Id,
		Camp: player.Camp,
		Hp:   player.HP,
		X:    player.X,
		Y:    player.Y,
		Z:    player.Z,
		EX:   player.EX,
		EY:   player.EY,
		EZ:   player.EZ,
	}
	return tankInfo
}

func (r *Room) ResetPlayer() {
	count1, count2 := 0, 0
	for k := range r.Players {
		player := Pm.GetPlayer(k)
		if player == nil {
			continue
		}

		player.HP = 100

		if player.Camp == 1 {
			setPlayerPos(player, count1)
			count1++
		} else {
			setPlayerPos(player, count2)
			count2++
		}
	}
}

func (r *Room) AddPlayer(id string) bool {
	player := Pm.GetPlayer(id)
	if player == nil {
		return false
	}

	if len(r.Players) >= r.MaxPlayer {
		log.Println("room is full")
		return false
	}

	if r.Status == Fighting {
		log.Println("room is fighting")
		return false
	}

	if r.Players[id] {
		log.Printf("player %s already in room", id)
		return false
	}

	r.Players[id] = true

	player.Camp = r.SwitchCamp()
	player.RoomId = r.Id
	if r.OwnerId == "" {
		r.OwnerId = id
	}
	r.broadcastMsgToAllPlayers(r.GetRoomInfo())
	return true
}

func (r *Room) RemovePlayer(id string) bool {
	player := Pm.GetPlayer(id)
	if player == nil {
		return false
	}

	if !r.Players[id] {
		log.Println("Remove player fail, player is not in this room.")
		return false
	}

	delete(r.Players, id)

	player.Camp = 0
	player.RoomId = -1

	if r.OwnerId == id {
		r.OwnerId = r.switchOwner()
	}

	if r.Status == Fighting {
		playerData := &databaseModel.PlayerData{}
		playerData.Loss++
		database.UpdatePlayerData(id, playerData)
		msgLeaveBattle := &msgProto.MsgLeaveBattle{
			Id: player.Id,
		}
		r.broadcastMsgToAllPlayers(msgLeaveBattle)
	}

	if len(r.Players) == 0 {
		RM.RemoveRoom(r)
	}
	r.broadcastMsgToAllPlayers(r.GetRoomInfo())
	return true
}

func (r *Room) IsOwner(player *Player) bool {
	return r.OwnerId == player.Id
}

func (r *Room) GetRoomInfo() msg.MsgBase {
	msgGetRoomInfo := &msgProto.MsgGetRoomInfo{}
	count := len(r.Players)
	msgGetRoomInfo.PlayerInfos = make([]*msgProto.PlayerInfo, count)

	i := 0

	for k := range r.Players {
		player := Pm.GetPlayer(k)
		if player == nil {
			continue
		}

		playerInfos := &msgProto.PlayerInfo{
			Id:      player.Id,
			Camp:    player.Camp,
			Win:     player.PlayerData.Win,
			Lost:    player.PlayerData.Loss,
			IsOwner: 0,
		}
		if r.OwnerId == player.Id {
			playerInfos.IsOwner = 1
		}
		msgGetRoomInfo.PlayerInfos[i] = playerInfos
		i++
	}

	return msgGetRoomInfo
}

func (r *Room) SwitchCamp() int {
	var (
		count1 int
		count2 int
	)

	for k := range r.Players {
		player := Pm.GetPlayer(k)
		if player.Camp == 1 {
			count1++
		} else if player.Camp == 2 {
			count2++
		}
	}
	if count1 >= count2 {
		return 2
	} else {
		return 1
	}
}

func (r *Room) broadcastMsgToAllPlayers(toMsg msg.MsgBase) {
	for k := range r.Players {
		player := Pm.GetPlayer(k)
		if player == nil {
			continue
		}
		player.Send(toMsg)

	}
}

func (r *Room) switchOwner() string {
	for k := range r.Players {
		return k
	}
	return ""
}
