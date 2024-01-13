package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type orderStatusCodeService struct {
	orderStatusCodeRepo interfaces.OrderStatusCodeRepository
}

func NewOrderStatusCodeService(o interfaces.OrderStatusCodeRepository) interfaces.OrderStatusCodeService {
	return &orderStatusCodeService{
		orderStatusCodeRepo: o,
	}
}

func (o *orderStatusCodeService) Fetch(ctx context.Context) ([]*ent.OrderStatusCode, error) {
	return o.orderStatusCodeRepo.Fetch(ctx)
}

func (o *orderStatusCodeService) GetByID(ctx context.Context, id int) (*ent.OrderStatusCode, error) {
	return o.orderStatusCodeRepo.GetByID(ctx, id)
}

func (o *orderStatusCodeService) Store(ctx context.Context, obj *ent.OrderStatusCode) (*ent.OrderStatusCode, error) {
	return o.orderStatusCodeRepo.Store(ctx, obj)
}

func (o *orderStatusCodeService) Update(ctx context.Context, id int, obj *ent.OrderStatusCode) (*ent.OrderStatusCode, error) {
	return o.orderStatusCodeRepo.Update(ctx, id, obj)
}

func (o *orderStatusCodeService) Delete(ctx context.Context, id int) error {
	return o.orderStatusCodeRepo.Delete(ctx, id)
}
