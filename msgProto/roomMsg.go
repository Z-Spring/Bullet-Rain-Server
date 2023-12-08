package msgProto

type MsgGetAchieve struct {
	ProtoName string `json:"protoName"`
	Win       int    `json:"win"`
	Lost      int    `json:"lost"`
}

func (b *MsgGetAchieve) GetProtoName() string {
	return "MsgGetAchieve"
}

type RoomInfo struct {
	ProtoName string `json:"protoName"`
	Id        int    `json:"id"`
	Count     int    `json:"count"`
	Status    int    `json:"status"`
}

type MsgGetRoomList struct {
	ProtoName string     `json:"protoName"`
	Rooms     []RoomInfo `json:"rooms"`
}

func (b *MsgGetRoomList) GetProtoName() string {
	return "MsgGetRoomList"
}

type MsgCreateRoom struct {
	ProtoName string `json:"protoName"`
	Reason    string `json:"reason"`
	Result    int    `json:"result"`
}

func (b *MsgCreateRoom) GetProtoName() string {
	return "MsgCreateRoom"
}

type MsgEnterRoom struct {
	ProtoName string `json:"protoName"`
	Id        int    `json:"id"`
	Result    int    `json:"result"`
}

func (b *MsgEnterRoom) GetProtoName() string {
	return "MsgEnterRoom"
}

type PlayerInfo struct {
	Id      string `json:"id"`
	Camp    int    `json:"camp"`
	Win     int    `json:"win"`
	Lost    int    `json:"lost"`
	IsOwner int    `json:"isOwner"`
}

type MsgGetRoomInfo struct {
	ProtoName   string        `json:"protoName"`
	PlayerInfos []*PlayerInfo `json:"players"`
}

func (b *MsgGetRoomInfo) GetProtoName() string {
	return "MsgGetRoomInfo"
}

type MsgLeaveRoom struct {
	ProtoName string `json:"protoName"`
	Result    int    `json:"result"`
}

func (b *MsgLeaveRoom) GetProtoName() string {
	return "MsgLeaveRoom"
}

type MsgStartBattle struct {
	ProtoName string `json:"protoName"`
	Result    int    `json:"result"`
	Reason    string `json:"reason"`
}

func (b *MsgStartBattle) GetProtoName() string {
	return "MsgStartBattle"
}
