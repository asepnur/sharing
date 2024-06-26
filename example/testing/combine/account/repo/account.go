package repo

import (
	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
	"github.com/asepnur/sharing/example/testing/combine/account/core/repo"
	"github.com/jinzhu/gorm"
)

type accountRepoImpl struct {
	conn *gorm.DB
}

func NewAccountRepo(conn *gorm.DB) repo.AccountRepo {
	return &accountRepoImpl{
		conn: conn,
	}
}

func (a *accountRepoImpl) Insert(account *model.Account) error {
	panic("not implemented") // TODO: Implement
}

func (a *accountRepoImpl) FindByUsername(username string) (*model.Account, error) {
	query := "SELECT * FROM account WHERE username = ?"
	rows, err := a.conn.Raw(query, username).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var account model.Account
	for rows.Next() {
		err = rows.Scan(&account.Serial, &account.Username, &account.Password, &account.Type)
		if err != nil {
			return nil, err
		}
	}
	return &account, nil
}
