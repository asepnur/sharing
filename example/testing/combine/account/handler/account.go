package handler

import (
	"context"

	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
	"github.com/asepnur/sharing/example/testing/combine/account/core/module"
)

type AccountHandler interface {
	Register(ctx context.Context, account *model.Account) error
	Authenticate(ctx context.Context, account *model.Account) error
}

type accountHandlerDependency struct {
	account module.AccountModule
}

func NewAccountHandler(account module.AccountModule) AccountHandler {
	return &accountHandlerDependency{
		account: account,
	}
}

func (a *accountHandlerDependency) Register(ctx context.Context, account *model.Account) error {
	if err := a.account.Register(account); err != nil {
		return err
	}
	return nil
}

func (a *accountHandlerDependency) Authenticate(ctx context.Context, account *model.Account) error {
	if err := a.account.Authenticate(account); err != nil {
		return err
	}
	return nil
}
