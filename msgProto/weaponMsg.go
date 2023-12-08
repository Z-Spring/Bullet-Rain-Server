package msgProto

type MsgSwitchWeapon struct {
	ProtoName string `json:"protoName"`
	WeaponId  int    `json:"weaponId"`
	Id        string `json:"id"`
}

func (b *MsgSwitchWeapon) GetProtoName() string {
	return "MsgSwitchWeapon"
}

type MsgSyncWeaponPosition struct {
	ProtoName string `json:"protoName"`
	Id        string `json:"id"`
	IsScope   bool   `json:"isScope"`
}

func (b *MsgSyncWeaponPosition) GetProtoName() string {
	return "MsgSyncWeaponPosition"
}
