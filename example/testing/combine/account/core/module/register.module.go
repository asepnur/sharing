package module

import (
	"errors"
	"log"

	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
)

var ErrAccountAlreadyExist = errors.New("account already exist")

func (a *accountModuleDependency) Register(account *model.Account) error {
	accountDB, err := a.repo.FindByUsername(account.Username)
	if err != nil && err != model.ErrAccountNotFound {
		log.Println("1, ", err)
		return err
	}
	if accountDB != nil {
		return ErrAccountAlreadyExist
	}
	err = a.repo.Insert(account)
	if err != nil {
		return err
	}
	return nil
}
