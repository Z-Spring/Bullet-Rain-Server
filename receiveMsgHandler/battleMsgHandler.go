package receiveMsgHandler

import (
	"errors"
	"fmt"
	"log"
	"time"

	"BulletRain_server/databaseModel"
	"BulletRain_server/msgProto"
)

var MsgGetGameTime = &msgProto.MsgGetGameTime{
	GameTime: 300.00,
}

func (h *MsgHandler) MsgStartBattle(info *ClientInfo, msgStartBattle *msgProto.MsgStartBattle) {
	player := info.Player
	if player == nil {
		return
	}
	if !IsInRoom(player) {
		log.Println("player is not in room")
		msgStartBattle.Reason = fmt.Sprintf("您不在房间中！请先加入房间！")
		msgStartBattle.Result = 1
		player.Send(msgStartBattle)
		return
	}
	room := RM.GetRoom(player.RoomId)
	if room == nil {
		log.Println("room is nil")
		msgStartBattle.Reason = fmt.Sprintf("房间不存在！请创建房间后开始游戏！")
		msgStartBattle.Result = 1
		player.Send(msgStartBattle)
		return
	}
	if !room.IsOwner(player) {
		log.Println("player is not owner")
		msgStartBattle.Reason = fmt.Sprintf("您不是房主，无法开始游戏！")
		msgStartBattle.Result = 1
		player.Send(msgStartBattle)
		return
	}
	if ok, err := room.StartBattle(); !ok {
		log.Println("start battle fail")
		msgStartBattle.Reason = fmt.Sprintf("开始游戏失败：%s", err)
		msgStartBattle.Result = 1
		player.Send(msgStartBattle)
		return
	}

	msgStartBattle.Result = 0
	log.Println("Can Start battle!")
	MsgGetGameTime.GameTime = 300.00
	player.Send(msgStartBattle)
}

func (h *MsgHandler) MsgBattleResult(info *ClientInfo, msgBattleResult *msgProto.MsgBattleResult) {
	player := info.Player
	if player == nil {
		return
	}
	if !IsInRoom(player) {
		log.Println("player is not in room")
		return
	}
	room := RM.GetRoom(player.RoomId)
	if room == nil {
		log.Println("room is nil")
		return
	}
	winCamp := room.JudgeBattleResult()
	if winCamp == 0 {
		return
	}
	msgBattleResult.WinCamp = winCamp

	room.broadcastMsgToAllPlayers(msgBattleResult)
}

/*func MsgLeaveBattle(info *ClientInfo, msgLeaveBattle *msgProto.MsgLeaveBattle) {
	player := info.Player
	if player == nil {
		return
	}

	room := RM.GetRoom(player.RoomId)
	if !room.RemovePlayer(player.Id) {
		log.Println("remove player fail")
		player.Send(msgLeaveBattle)
		return
	}

	msgLeaveBattle.Id = player.Id
	player.Send(msgLeaveBattle)
}*/

func (r *Room) canStartBattle() (bool, error) {
	if r.Status != Prepare {
		return false, errors.New("房间状态不是在准备中")
	}
	if len(r.Players) < 2 {
		return false, errors.New("房间人数少于2人")
	}
	var (
		count1 int
		count2 int
	)
	for k := range r.Players {
		player := Pm.GetPlayer(k)
		if player == nil {
			continue
		}

		if player.Camp == 1 {
			count1++
		} else {
			count2++
		}

		if count1 > 0 && count2 > 0 {
			return true, nil
		}

	}
	return false, errors.New("房间人数不足")

}

func (r *Room) StartBattle() (bool, error) {
	if ok, err := r.canStartBattle(); !ok {
		return false, err
	}
	r.Status = Fighting

	r.ResetPlayer()
	msgEnterBattle := &msgProto.MsgEnterBattle{
		TankInfos: make([]msgProto.TankInfo, len(r.Players)),
	}
	i := 0
	for k := range r.Players {
		player := Pm.GetPlayer(k)
		if player == nil {
			continue
		}
		msgEnterBattle.TankInfos[i] = convertPlayerInfoToTankInfo(player)
		i++
	}
	r.broadcastMsgToAllPlayers(msgEnterBattle)
	return true, nil
}

func (r *Room) BroadcastBattleResult() {
	if r.Status != Fighting {
		return
	}

	if time.Since(lastJudgeTime) < 10*time.Second {
		return
	}
	lastJudgeTime = time.Now()

	winCamp := r.JudgeBattleResult()

	if winCamp == 0 {
		return
	}

	r.Status = Prepare

	i := 0
	resultMsg := &msgProto.MsgResult{
		PlayerData: make([]databaseModel.PlayerData, len(r.Players)),
	}

	for k := range r.Players {
		player := Pm.GetPlayer(k)
		if player == nil {
			continue
		}

		playerData := &player.PlayerData

		if player.Camp == winCamp {
			playerData.Win++
		} else {
			playerData.Loss++
		}
		playerData.Id = player.Id
		log.Println("ID:", playerData.Id, "Win:", playerData.Win, "Loss:", playerData.Loss)

		// todo: 编码解码很复杂，需要优化
		//jsonInfo, err := json.Marshal(playerData)
		//if err != nil {
		//	log.Println("json marshal fail")
		//}
		resultMsg.PlayerData[i] = *playerData
		i++
	}
	r.broadcastMsgToAllPlayers(resultMsg)
	msgBattleResult := &msgProto.MsgBattleResult{
		WinCamp: winCamp,
	}
	log.Println("winCamp:", winCamp)

	r.broadcastMsgToAllPlayers(msgBattleResult)
}

func (r *Room) JudgeBattleResult() int {
	var (
		count1 int
		count2 int
	)

	for k := range r.Players {
		player := Pm.GetPlayer(k)
		if player == nil {
			continue
		}
		if !IsDie(player) {
			if player.Camp == 1 {
				count1++
			} else {
				count2++
			}
		}

	}
	if count1 == 0 {
		return 2
	} else if count2 == 0 {
		return 1
	}
	return 0
}

func IsDie(player *Player) bool {
	return player.HP <= 0
}
