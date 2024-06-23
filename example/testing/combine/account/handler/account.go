package handler

import (
	"context"

	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
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
	panic("not implemented") // TODO: Implement
}

func (a *accountHandlerDependency) Authenticate(ctx context.Context, account *model.Account) error {
	panic("not implemented")
}
