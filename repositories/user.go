package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entUserRepository struct {
	Client *ent.Client
}

func (e *entUserRepository) Fetch(ctx context.Context) ([]*ent.User, error) {
	users, err := e.Client.User.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (e *entUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	user, err := e.Client.User.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *entUserRepository) Store(ctx context.Context, obj *ent.User) (err error) {
	//lint:ignore SA4006 we want to return the result of creating operation
	obj, err = e.Client.User.
		Create().
		SetAvatarURL(obj.AvatarURL).
		SetNillableEmail(obj.Email).
		SetName(obj.Name).
		SetPhone(obj.Phone).
		SetNillableAddress(obj.Address).
		SetNillablePostalCode(obj.PostalCode).
		SetNillableOtherAddressInfo(obj.OtherAddressInfo).
		Save(ctx)
	return

}

func (e *entUserRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.User) (err error) {
	//lint:ignore SA4006 we want to return the result of updating operation
	obj, err = e.Client.User.
		UpdateOneID(id).
		SetAvatarURL(obj.AvatarURL).
		SetNillableEmail(obj.Email).
		SetNillableName(&obj.Name).
		SetNillablePhone(&obj.Phone).
		SetNillableAddress(obj.Address).
		SetNillablePostalCode(obj.PostalCode).
		SetNillableOtherAddressInfo(obj.OtherAddressInfo).
		Save(ctx)
	return
}

func (e *entUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.User.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntUserRepository(client *ent.Client) interfaces.UserRepository {
	return &entUserRepository{Client: client}
}
