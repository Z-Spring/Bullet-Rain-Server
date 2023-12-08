package msgProto

type MsgSyncPlayer struct {
	ProtoName string `json:"protoName"`
	Id        string `json:"id"`

	X  float32 `json:"x"`
	Y  float32 `json:"y"`
	Z  float32 `json:"z"`
	EX float32 `json:"ex"`
	EY float32 `json:"ey"`
	EZ float32 `json:"ez"`
}

func (b *MsgSyncPlayer) GetProtoName() string {
	return "MsgSyncPlayer"
}

type MsgFire struct {
	ProtoName string  `json:"protoName"`
	Id        string  `json:"id"`
	X         float32 `json:"x"`
	Y         float32 `json:"y"`
	Z         float32 `json:"z"`
	EX        float32 `json:"ex"`
	EY        float32 `json:"ey"`
	EZ        float32 `json:"ez"`
}

func (b *MsgFire) GetProtoName() string {
	return "MsgFire"
}

type MsgHit struct {
	ProtoName string `json:"protoName"`
	Id        string `json:"id"`
	Hp        int    `json:"hp"`
	TargetId  string `json:"targetId"`

	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`

	Damage int `json:"damage"`
}

func (b *MsgHit) GetProtoName() string {
	return "MsgHit"
}
