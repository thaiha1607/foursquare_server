package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entOrderLineItemRepository struct {
	Client *ent.Client
}

func (e *entOrderLineItemRepository) Fetch(ctx context.Context) ([]*ent.OrderLineItem, error) {
	order_line_items, err := e.Client.OrderLineItem.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return order_line_items, nil
}

func (e *entOrderLineItemRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderLineItem, error) {
	order_line_item, err := e.Client.OrderLineItem.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return order_line_item, nil
}

func (e *entOrderLineItemRepository) Store(ctx context.Context, obj *ent.OrderLineItem) (res *ent.OrderLineItem, err error) {
	res, err = e.Client.OrderLineItem.
		Create().
		SetOrderID(obj.OrderID).
		SetProductID(obj.ProductID).
		SetQty(obj.Qty).
		Save(ctx)
	return
}

func (e *entOrderLineItemRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.OrderLineItem) (res *ent.OrderLineItem, err error) {
	res, err = e.Client.OrderLineItem.
		UpdateOneID(id).
		SetNillableQty(&obj.Qty).
		Save(ctx)
	return
}

func (e *entOrderLineItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.OrderLineItem.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntOrderLineItemRepository(client *ent.Client) interfaces.OrderLineItemRepository {
	return &entOrderLineItemRepository{Client: client}
}
