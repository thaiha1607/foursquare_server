package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type orderLineItemService struct {
	orderLineItemRepo interfaces.OrderLineItemRepository
}

func NewOrderLineItemService(o interfaces.OrderLineItemRepository) interfaces.OrderLineItemService {
	return &orderLineItemService{
		orderLineItemRepo: o,
	}
}

func (o *orderLineItemService) Fetch(ctx context.Context) ([]*ent.OrderLineItem, error) {
	return o.orderLineItemRepo.Fetch(ctx)
}

func (o *orderLineItemService) GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderLineItem, error) {
	return o.orderLineItemRepo.GetByID(ctx, id)
}

func (o *orderLineItemService) Store(ctx context.Context, obj *ent.OrderLineItem) (*ent.OrderLineItem, error) {
	return o.orderLineItemRepo.Store(ctx, obj)
}

func (o *orderLineItemService) Update(ctx context.Context, id uuid.UUID, obj *ent.OrderLineItem) (*ent.OrderLineItem, error) {
	return o.orderLineItemRepo.Update(ctx, id, obj)
}

func (o *orderLineItemService) Delete(ctx context.Context, id uuid.UUID) error {
	return o.orderLineItemRepo.Delete(ctx, id)
}
