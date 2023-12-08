package databaseModel

type Account struct {
	Id string `json:"id,omitempty"`
	Pw string `json:"pw,omitempty"`
}

func (Account) TableName() string {
	return "account"
}
