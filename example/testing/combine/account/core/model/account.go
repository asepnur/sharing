package model

type Account struct {
	Serial   string
	Username string
	Password string
	Type     string
}

func (Account) TableName() string {
	return "account"
}
