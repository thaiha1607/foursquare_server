package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type orderService struct {
	orderRepo interfaces.OrderRepository
}

func NewOrderService(o interfaces.OrderRepository) interfaces.OrderService {
	return &orderService{
		orderRepo: o,
	}
}

func (o *orderService) Fetch(ctx context.Context) ([]*ent.Order, error) {
	return o.orderRepo.Fetch(ctx)
}

func (o *orderService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Order, error) {
	return o.orderRepo.GetByID(ctx, id)
}

func (o *orderService) Store(ctx context.Context, obj *ent.Order) error {
	return o.orderRepo.Store(ctx, obj)
}

func (o *orderService) Update(ctx context.Context, id uuid.UUID, obj *ent.Order) error {
	return o.orderRepo.Update(ctx, id, obj)
}

func (o *orderService) Delete(ctx context.Context, id uuid.UUID) error {
	return o.orderRepo.Delete(ctx, id)
}
