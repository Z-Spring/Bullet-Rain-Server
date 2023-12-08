package receiveMsgHandler

import (
	"fmt"
	"log"

	"BulletRain_server/database"
	"BulletRain_server/msgProto"
)

var Pm = PlayerManager{}

func (h *MsgHandler) MsgLogin(info *ClientInfo, msgLogin *msgProto.MsgLogin) {
	if !database.CheckPassword(msgLogin.Id, msgLogin.Pw) {
		msgLogin.Result = 1
		Send(info, msgLogin)
		return
	}
	if info.Player != nil {
		log.Printf("player %s already login\n", msgLogin.Id)
		return
	}
	if Pm.IsOnline(msgLogin.Id) {
		log.Printf("player %s already login\n", msgLogin.Id)
		player := Pm.GetPlayer(msgLogin.Id)
		msgKick := &msgProto.MsgKick{
			Reason: fmt.Sprintf("您的账号在其他地方登录，请退出后重新登录或修改密码"),
		}
		player.Send(msgKick)
		//netHandle.CloseConnection(player.ClientInfo)
	}

	playerData := database.GetPlayerData(msgLogin.Id)
	player := &Player{
		Id:         msgLogin.Id,
		ClientInfo: info,
		PlayerData: playerData,
		RoomId:     -1,
	}
	Pm.AddPlayer(msgLogin.Id, player)

	info.Player = player
	msgLogin.Result = 0
	log.Printf("Player %s login\n", msgLogin.Id)
	player.Send(msgLogin)
}

func (h *MsgHandler) MsgRegister(info *ClientInfo, msgRegister *msgProto.MsgRegister) {
	if database.Register(msgRegister.Id, msgRegister.Pw) {
		database.CreatePlayerData(msgRegister.Id)
		msgRegister.Result = 0
	} else {
		msgRegister.Result = 1
	}
	Send(info, msgRegister)

}
