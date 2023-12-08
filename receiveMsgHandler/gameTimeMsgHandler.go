package receiveMsgHandler

func (m *RoomManager) UpdateGameTime() {
	for _, room := range m.Rooms {
		if room.Status != Fighting {
			continue
		}
		MsgGetGameTime.GameTime -= 1
		if MsgGetGameTime.GameTime <= 0 {
			MsgGetGameTime.GameTime = 0
		}
		room.broadcastMsgToAllPlayers(MsgGetGameTime)
	}

}
