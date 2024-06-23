package repo

import (
	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
	"github.com/asepnur/sharing/example/testing/combine/account/core/repo"
)

type accountRepoImpl struct {
}

func NewAccountRepo() repo.AccountRepo {
	return &accountRepoImpl{}
}

func (a *accountRepoImpl) Insert(account *model.Account) error {
	panic("not implemented") // TODO: Implement
}

func (a *accountRepoImpl) FindByUsername(username string) (*model.Account, error) {
	panic("not implemented") // TODO: Implement
}
