package repo

import "github.com/asepnur/sharing/example/testing/combine/account/core/model"

type AccountRepo interface {
	Insert(account *model.Account) error
	FindByUsername(username string) (*model.Account, error)
}
