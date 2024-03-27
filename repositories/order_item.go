package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entOrderItemRepository struct {
	Client *ent.Client
}

func (e *entOrderItemRepository) Fetch(ctx context.Context) ([]*ent.OrderItem, error) {
	order_line_items, err := e.Client.OrderItem.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return order_line_items, nil
}

func (e *entOrderItemRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderItem, error) {
	order_line_item, err := e.Client.OrderItem.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return order_line_item, nil
}

func (e *entOrderItemRepository) Store(ctx context.Context, obj *ent.OrderItem) (res *ent.OrderItem, err error) {
	res, err = e.Client.OrderItem.
		Create().
		SetOrderID(obj.OrderID).
		SetProductID(obj.ProductID).
		SetQty(obj.Qty).
		Save(ctx)
	return
}

func (e *entOrderItemRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.OrderItem) (res *ent.OrderItem, err error) {
	res, err = e.Client.OrderItem.
		UpdateOneID(id).
		SetNillableQty(&obj.Qty).
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
