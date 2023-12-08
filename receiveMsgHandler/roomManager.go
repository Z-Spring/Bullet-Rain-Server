package receiveMsgHandler

import (
	"BulletRain_server/msg"
	"BulletRain_server/msgProto"
)

type RoomManager struct {
	MaxId int
	Rooms map[int]*Room
}

func (m *RoomManager) GetRoomList() msg.MsgBase {
	msgGetRoomList := &msgProto.MsgGetRoomList{}
	count := len(m.Rooms)
	msgGetRoomList.Rooms = make([]msgProto.RoomInfo, count)

	i := 0
	for _, room := range m.Rooms {
		roomInfo := msgProto.RoomInfo{
			Id:     room.Id,
			Count:  len(room.Players),
			Status: room.Status,
		}
		msgGetRoomList.Rooms[i] = roomInfo
		i++
	}
	return msgGetRoomList
}

func (m *RoomManager) AddRoom() *Room {
	m.MaxId++
	room := &Room{
		Id:        m.MaxId,
		MaxPlayer: 6,
		Players:   make(map[string]bool, 6),
	}
	m.Rooms[m.MaxId] = room
	return room
}

func (m *RoomManager) RemoveRoom(room *Room) bool {
	delete(m.Rooms, room.Id)
	return true
}

func (m *RoomManager) GetRoom(id int) *Room {
	if room, ok := m.Rooms[id]; ok {
		return room
	}
	return nil
}

func (m *RoomManager) UpdateBattleResult() {
	for _, room := range m.Rooms {
		room.BroadcastBattleResult()
	}
}
