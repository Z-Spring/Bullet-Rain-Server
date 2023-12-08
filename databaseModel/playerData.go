package databaseModel

type PlayerData struct {
	Id   string `json:"id"`
	Coin int    `json:"coin"`
	Text string `json:"text"`
	Win  int    `json:"win"`
	Loss int    `json:"loss"`
}

func (PlayerData) TableName() string {
	return "player_data"
}
