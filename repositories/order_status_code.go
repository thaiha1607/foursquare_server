package repositories

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entOrderStatusCodeRepository struct {
	Client *ent.Client
}

func (e *entOrderStatusCodeRepository) Fetch(ctx context.Context) ([]*ent.OrderStatusCode, error) {
	order_status_codes, err := e.Client.OrderStatusCode.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return order_status_codes, nil
}

func (e *entOrderStatusCodeRepository) GetByID(ctx context.Context, id int) (*ent.OrderStatusCode, error) {
	order_status_code, err := e.Client.OrderStatusCode.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return order_status_code, nil
}

func (e *entOrderStatusCodeRepository) Store(ctx context.Context, obj *ent.OrderStatusCode) (res *ent.OrderStatusCode, err error) {
	res, err = e.Client.OrderStatusCode.
		Create().
		SetOrderStatus(obj.OrderStatus).
		Save(ctx)
	return
}

func (e *entOrderStatusCodeRepository) Update(ctx context.Context, id int, obj *ent.OrderStatusCode) (res *ent.OrderStatusCode, err error) {
	res, err = e.Client.OrderStatusCode.
		UpdateOneID(id).
		SetNillableOrderStatus(&obj.OrderStatus).
		Save(ctx)
	return
}

func (e *entOrderStatusCodeRepository) Delete(ctx context.Context, id int) error {
	return e.Client.OrderStatusCode.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntOrderStatusCodeRepository(client *ent.Client) interfaces.OrderStatusCodeRepository {
	return &entOrderStatusCodeRepository{Client: client}
}
