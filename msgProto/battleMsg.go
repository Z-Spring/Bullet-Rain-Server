package msgProto

type TankInfo struct {
	Id   string  `json:"id"`
	Camp int     `json:"camp"`
	Hp   int     `json:"hp"`
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	Z    float32 `json:"z"`
	EX   float32 `json:"ex"`
	EY   float32 `json:"ey"`
	EZ   float32 `json:"ez"`
}

type MsgEnterBattle struct {
	ProtoName string     `json:"protoName"`
	TankInfos []TankInfo `json:"tanks"`
}

func (b *MsgEnterBattle) GetProtoName() string {
	return "MsgEnterBattle"
}

type MsgLeaveBattle struct {
	ProtoName string `json:"protoName"`
	Id        string `json:"id"`
}

func (b *MsgLeaveBattle) GetProtoName() string {
	return "MsgLeaveBattle"
}

type MsgBattleResult struct {
	ProtoName string `json:"protoName"`
	WinCamp   int    `json:"winCamp"`
}

func (b *MsgBattleResult) GetProtoName() string {
	return "MsgBattleResult"
}
