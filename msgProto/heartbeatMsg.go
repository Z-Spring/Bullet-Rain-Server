package msgProto

type MsgPing struct {
	ProtoName string `json:"protoName"`
}

func (m *MsgPing) GetProtoName() string {
	return "MsgPing"
}

type MsgPong struct {
	ProtoName string `json:"protoName"`
}

func (m *MsgPong) GetProtoName() string {
	return "MsgPong"
}
