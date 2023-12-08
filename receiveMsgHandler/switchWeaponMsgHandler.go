package receiveMsgHandler

import "BulletRain_server/msgProto"

func (h *MsgHandler) MsgSwitchWeapon(info *ClientInfo, msgSwitchWeapon *msgProto.MsgSwitchWeapon) {
	player := info.Player
	if player == nil {
		return
	}
	if !IsInRoom(player) {
		return
	}
	room := RM.GetRoom(player.RoomId)
	if room == nil {
		return
	}
	if room.Status != Fighting {
		return
	}
	msgSwitchWeapon.Id = player.Id
	room.broadcastMsgToAllPlayers(msgSwitchWeapon)

}

func (h *MsgHandler) MsgSyncWeaponPosition(info *ClientInfo, msgSyncWeaponPosition *msgProto.MsgSyncWeaponPosition) {
	player := info.Player
	if player == nil {
		return
	}
	if !IsInRoom(player) {
		return
	}
	room := RM.GetRoom(player.RoomId)
	if room == nil {
		return
	}
	if room.Status != Fighting {
		return
	}
	msgSyncWeaponPosition.Id = player.Id
	room.broadcastMsgToAllPlayers(msgSyncWeaponPosition)
}
