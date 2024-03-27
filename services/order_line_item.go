package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type orderItemService struct {
	orderItemRepo interfaces.OrderItemRepository
}

func NewOrderItemService(o interfaces.OrderItemRepository) interfaces.OrderItemService {
	return &orderItemService{
		orderItemRepo: o,
	}
}

func (o *orderItemService) Fetch(ctx context.Context) ([]*ent.OrderItem, error) {
	return o.orderItemRepo.Fetch(ctx)
}

func (o *orderItemService) GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderItem, error) {
	return o.orderItemRepo.GetByID(ctx, id)
}

func (o *orderItemService) Store(ctx context.Context, obj *ent.OrderItem) (*ent.OrderItem, error) {
	return o.orderItemRepo.Store(ctx, obj)
}

func (o *orderItemService) Update(ctx context.Context, id uuid.UUID, obj *ent.OrderItem) (*ent.OrderItem, error) {
	return o.orderItemRepo.Update(ctx, id, obj)
}

func (o *orderItemService) Delete(ctx context.Context, id uuid.UUID) error {
	return o.orderItemRepo.Delete(ctx, id)
}
