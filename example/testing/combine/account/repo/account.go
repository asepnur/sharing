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
	for rows.Next() {
		var account model.Account
		err = a.conn.ScanRows(rows, &account)
		if err != nil {
			return nil, err
		}
		return &account, nil
	}
	return nil, nil
}
