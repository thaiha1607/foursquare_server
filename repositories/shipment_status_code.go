package repositories

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entShipmentStatusCodeRepository struct {
	Client *ent.Client
}

func (e *entShipmentStatusCodeRepository) Fetch(ctx context.Context) ([]*ent.ShipmentStatusCode, error) {
	shipment_status_codes, err := e.Client.ShipmentStatusCode.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return shipment_status_codes, nil
}

func (e *entShipmentStatusCodeRepository) GetByID(ctx context.Context, id int) (*ent.ShipmentStatusCode, error) {
	shipment_status_code, err := e.Client.ShipmentStatusCode.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return shipment_status_code, nil
}

func (e *entShipmentStatusCodeRepository) Store(ctx context.Context, obj *ent.ShipmentStatusCode) (res *ent.ShipmentStatusCode, err error) {
	res, err = e.Client.ShipmentStatusCode.
		Create().
		SetShipmentStatus(obj.ShipmentStatus).
		Save(ctx)
	return
}

func (e *entShipmentStatusCodeRepository) Update(ctx context.Context, id int, obj *ent.ShipmentStatusCode) (res *ent.ShipmentStatusCode, err error) {
	res, err = e.Client.ShipmentStatusCode.
		UpdateOneID(id).
		SetNillableShipmentStatus(&obj.ShipmentStatus).
		Save(ctx)
	return
}

func (e *entShipmentStatusCodeRepository) Delete(ctx context.Context, id int) error {
	return e.Client.ShipmentStatusCode.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntShipmentStatusCodeRepository(client *ent.Client) interfaces.ShipmentStatusCodeRepository {
	return &entShipmentStatusCodeRepository{Client: client}
}
