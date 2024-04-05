package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/ent/personaddress"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entPersonAddressRepository struct {
	Client *ent.Client
}

func (e *entPersonAddressRepository) Fetch(ctx context.Context) ([]*ent.PersonAddress, error) {
	person_addresses, err := e.Client.PersonAddress.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return person_addresses, nil
}

func (e *entPersonAddressRepository) GetByPersonID(ctx context.Context, person_id uuid.UUID) ([]*ent.PersonAddress, error) {
	person_address, err := e.Client.PersonAddress.
		Query().
		Where(
			personaddress.PersonID(person_id),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return person_address, nil
}

func (e *entPersonAddressRepository) Store(ctx context.Context, obj *ent.PersonAddress) (res *ent.PersonAddress, err error) {
	res, err = e.Client.PersonAddress.
		Create().
		SetPersonID(obj.PersonID).
		SetAddressID(obj.AddressID).
		Save(ctx)
	return
}

func (e *entPersonAddressRepository) Delete(ctx context.Context, person_id uuid.UUID, address_id uuid.UUID) error {
	_, err := e.Client.PersonAddress.
		Delete().
		Where(
			personaddress.And(
				personaddress.PersonID(person_id),
				personaddress.AddressID(address_id),
			),
		).
		Exec(ctx)
	return err
}

func NewEntPersonAddressRepository(client *ent.Client) interfaces.PersonAddressRepository {
	return &entPersonAddressRepository{Client: client}
}
