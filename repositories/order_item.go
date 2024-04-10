package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/ent/orderitem"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entOrderItemRepository struct {
	Client *ent.Client
}

func (e *entOrderItemRepository) GetByOrderID(ctx context.Context, order_id uuid.UUID) ([]*ent.OrderItem, error) {
	order_items, err := e.Client.OrderItem.
		Query().
		Where(
			orderitem.OrderID(order_id),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return order_items, nil
}

func (e *entOrderItemRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderItem, error) {
	order_item, err := e.Client.OrderItem.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return order_item, nil
}

func (e *entOrderItemRepository) Store(ctx context.Context, obj *ent.OrderItem) (res *ent.OrderItem, err error) {
	res, err = e.Client.OrderItem.
		Create().
		SetOrderID(obj.OrderID).
		SetProductID(obj.ProductID).
		SetProductColorID(obj.ProductColorID).
		SetNillableSrcUnitID(obj.SrcUnitID).
		SetNillableDstUnitID(obj.DstUnitID).
		SetQty(*obj.Qty).
		SetPricePerUnit(*obj.PricePerUnit).
		SetNillableStatus(obj.Status).
		Save(ctx)
	return
}

func (e *entOrderItemRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.OrderItem) (res *ent.OrderItem, err error) {
	res, err = e.Client.OrderItem.
		UpdateOneID(id).
		SetNillableQty(obj.Qty).
		SetNillablePricePerUnit(obj.PricePerUnit).
		SetNillableStatus(obj.Status).
		Save(ctx)
	return
}

func (e *entOrderItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.OrderItem.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntOrderItemRepository(client *ent.Client) interfaces.OrderItemRepository {
	return &entOrderItemRepository{Client: client}
}
