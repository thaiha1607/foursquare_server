package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type accountService struct {
	accountRepo interfaces.AccountRepository
}

func NewAccountService(a interfaces.AccountRepository) interfaces.AccountService {
	return &accountService{
		accountRepo: a,
	}
}

func (a *accountService) Fetch(ctx context.Context) ([]*ent.Account, error) {
	return a.accountRepo.Fetch(ctx)
}

func (a *accountService) GetByID(ctx context.Context, id string) (*ent.Account, error) {
	return a.accountRepo.GetByID(ctx, id)
}

func (a *accountService) Store(ctx context.Context, obj *ent.Account) error {
	return a.accountRepo.Store(ctx, obj)
}

func (a *accountService) Update(ctx context.Context, id string, obj *ent.Account) error {
	return a.accountRepo.Update(ctx, id, obj)
}

func (a *accountService) Delete(ctx context.Context, id string) error {
	return a.accountRepo.Delete(ctx, id)
}
