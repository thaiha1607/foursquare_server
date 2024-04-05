package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entAddressRepository struct {
	Client *ent.Client
}

func (e *entAddressRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Address, error) {
	address, err := e.Client.Address.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (e *entAddressRepository) Store(ctx context.Context, obj *ent.Address) (res *ent.Address, err error) {
	res, err = e.Client.Address.
		Create().
		SetLine1(obj.Line1).
		SetNillableLine2(obj.Line2).
		SetCity(obj.City).
		SetStateOrProvince(obj.StateOrProvince).
		SetZipOrPostcode(obj.ZipOrPostcode).
		SetCountry(obj.Country).
		SetNillableOtherAddressDetails(obj.OtherAddressDetails).
		Save(ctx)
	return
}

func (e *entAddressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.Address.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntAddressRepository(client *ent.Client) interfaces.AddressRepository {
	return &entAddressRepository{Client: client}
}
