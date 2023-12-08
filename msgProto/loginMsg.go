package msgProto

type MsgLogin struct {
	ProtoName string `json:"protoName"`
	Id        string `json:"id"`
	Pw        string `json:"pw"`
	Result    int    `json:"result"`
}

func (login *MsgLogin) GetProtoName() string {
	return "MsgLogin"
}

type MsgRegister struct {
	ProtoName string `json:"protoName"`
	Id        string `json:"id"`
	Pw        string `json:"pw"`
	Result    int    `json:"result"`
}

func (register *MsgRegister) GetProtoName() string {
	return "MsgRegister"
}

type MsgKick struct {
	ProtoName string `json:"protoName"`
	Reason    string `json:"reason"`
}

func (kick *MsgKick) GetProtoName() string {
	return "MsgKick"
}
