package repo

import (
	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
)

type AccountRepo interface {
	Insert(account *model.Account) error
	FindByUsername(username string) (*model.Account, error)
}

type stubRepo struct {
}

func NewStubRepo() AccountRepo {
	return &stubRepo{}
}

func (s *stubRepo) Insert(account *model.Account) error {
	return nil
}
func (s *stubRepo) FindByUsername(username string) (*model.Account, error) {
	if username == "new_user" {
		return &model.Account{Username: "new_user", Password: "password"}, nil
	}
	return nil, nil
}
