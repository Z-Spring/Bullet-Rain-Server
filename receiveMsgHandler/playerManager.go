package receiveMsgHandler

type PlayerManager struct {
}

var players = make(map[string]*Player)

func (pm PlayerManager) IsOnline(id string) bool {
	_, ok := players[id]
	return ok
}

func (pm PlayerManager) GetPlayer(id string) *Player {
	if player, ok := players[id]; ok {
		return player
	}
	return &Player{}
}

func (pm PlayerManager) AddPlayer(id string, player *Player) {
	players[id] = player
}

func (pm PlayerManager) RemovePlayer(id string) {
	delete(players, id)
}
