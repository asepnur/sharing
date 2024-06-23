package module

import (
	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
	"github.com/asepnur/sharing/example/testing/combine/account/core/repo"
)

type AccountModule interface {
	Register(account *model.Account) error
	Authenticate(account *model.Account) error
}

type accountModuleDependency struct {
	repo repo.AccountRepo
}

func NewAccountModule(repo repo.AccountRepo) AccountModule {
	return &accountModuleDependency{
		repo: repo,
	}
}
