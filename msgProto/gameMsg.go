package msgProto

type MsgGetGameTime struct {
	ProtoName string  `json:"protoName"`
	GameTime  float32 `json:"gameTime"`
}

func (b *MsgGetGameTime) GetProtoName() string {
	return "MsgGetGameTime"
}
