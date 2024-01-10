package repositories

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entAccountRepository struct {
	Client *ent.Client
}

func (e *entAccountRepository) Fetch(ctx context.Context) ([]*ent.Account, error) {
	conversations, err := e.Client.Account.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return conversations, nil
}

func (e *entAccountRepository) GetByID(ctx context.Context, id string) (*ent.Account, error) {
	conversation, err := e.Client.Account.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return conversation, nil
}

func (e *entAccountRepository) Store(ctx context.Context, obj *ent.Account) (err error) {
	//lint:ignore SA4006 we want to return the result of creating operation
	obj, err = e.Client.Account.
		Create().
		SetUserID(obj.UserID).
		SetNillablePasswordHash(obj.PasswordHash).
		Save(ctx)
	return
}

func (e *entAccountRepository) Update(ctx context.Context, id string, obj *ent.Account) (err error) {
	//lint:ignore SA4006 we want to return the result of updating operation
	obj, err = e.Client.Account.
		UpdateOneID(id).
		SetNillableLastReset(obj.LastReset).
		SetNillableLastEmailVerification(obj.LastEmailVerification).
		SetNillableLastPhoneVerification(obj.LastPhoneVerification).
		SetNillableIsEmailVerified(&obj.IsEmailVerified).
		SetNillableIsPhoneVerified(&obj.IsPhoneVerified).
		SetNillableRole(&obj.Role).
		SetNillablePasswordHash(obj.PasswordHash).
		Save(ctx)
	return
}

func (e *entAccountRepository) Delete(ctx context.Context, id string) error {
	return e.Client.Account.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntAccountRepository(client *ent.Client) interfaces.AccountRepository {
	return &entAccountRepository{Client: client}
}
