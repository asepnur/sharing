package module

import "github.com/asepnur/sharing/example/testing/combine/account/core/model"

func (a *accountModuleDependency) Authenticate(account *model.Account) error {
	accountDB, err := a.repo.FindByUsername(account.Username)
	if err != nil {
		return err
	}
	if accountDB == nil {
		return model.ErrAccountNotFound
	}
	if accountDB.Password != account.Password {
		return model.ErrInvalidPassword
	}
	return nil
}
