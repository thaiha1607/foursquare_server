package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entShipmentItemRepository struct {
	Client *ent.Client
}

func (e *entShipmentItemRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.ShipmentItem, error) {
	shipment_item, err := e.Client.ShipmentItem.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return shipment_item, nil
}

func (e *entShipmentItemRepository) Store(ctx context.Context, obj *ent.ShipmentItem) (res *ent.ShipmentItem, err error) {
	res, err = e.Client.ShipmentItem.
		Create().
		SetShipmentID(obj.ShipmentID).
		SetOrderItemID(obj.OrderItemID).
		SetQty(*obj.Qty).
		SetTotal(*obj.Total).
		Save(ctx)
	return
}

func (e *entShipmentItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.ShipmentItem.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntShipmentItemRepository(client *ent.Client) interfaces.ShipmentItemRepository {
	return &entShipmentItemRepository{Client: client}
}
