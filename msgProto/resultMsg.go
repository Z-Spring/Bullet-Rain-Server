package msgProto

import "BulletRain_server/databaseModel"

type MsgResult struct {
	ProtoName  string                     `json:"protoName,omitempty"`
	PlayerData []databaseModel.PlayerData `json:"playerData"`
}

func (b *MsgResult) GetProtoName() string {
	return "MsgResult"
}
