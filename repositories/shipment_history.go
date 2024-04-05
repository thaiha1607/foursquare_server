package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entShipmentHistoryRepository struct {
	Client *ent.Client
}

func (e *entShipmentHistoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.ShipmentHistory, error) {
	shipment_history, err := e.Client.ShipmentHistory.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return shipment_history, nil
}

func (e *entShipmentHistoryRepository) Store(ctx context.Context, obj *ent.ShipmentHistory) (res *ent.ShipmentHistory, err error) {
	res, err = e.Client.ShipmentHistory.
		Create().
		SetShipmentID(obj.ShipmentID).
		SetPersonID(obj.PersonID).
		SetNillableOldStatusCode(obj.OldStatusCode).
		SetNillableNewStatusCode(obj.NewStatusCode).
		SetNillableDescription(obj.Description).
		Save(ctx)
	return
}

func (e *entShipmentHistoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.ShipmentHistory.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntShipmentHistoryRepository(client *ent.Client) interfaces.ShipmentHistoryRepository {
	return &entShipmentHistoryRepository{Client: client}
}
