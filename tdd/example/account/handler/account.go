package handler

import (
	"context"

	"github.com/asepnur/sharing/tdd/example/account/model"
)

type AccountHandler interface {
	Register(ctx context.Context, account *model.Account) error
	Authenticate(ctx context.Context, account *model.Account) error
}

type accountHandlerDependency struct {
}

func NewAccountHandler() AccountHandler {
	return &accountHandlerDependency{}
}

func (a *accountHandlerDependency) Register(ctx context.Context, account *model.Account) error {
	return nil
}

func (a *accountHandlerDependency) Authenticate(ctx context.Context, account *model.Account) error {
	return nil
}
