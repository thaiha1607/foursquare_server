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

func NewOrderItemService(r interfaces.OrderItemRepository) interfaces.OrderItemService {
	return &orderItemService{
		orderItemRepo: r,
	}
}

func (s *orderItemService) GetByOrderID(ctx context.Context, order_id uuid.UUID) ([]*ent.OrderItem, error) {
	return s.orderItemRepo.GetByOrderID(ctx, order_id)
}

func (s *orderItemService) GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderItem, error) {
	return s.orderItemRepo.GetByID(ctx, id)
}

func (s *orderItemService) Store(ctx context.Context, obj *ent.OrderItem) (*ent.OrderItem, error) {
	return s.orderItemRepo.Store(ctx, obj)
}

func (s *orderItemService) Update(ctx context.Context, id uuid.UUID, obj *ent.OrderItem) (*ent.OrderItem, error) {
	return s.orderItemRepo.Update(ctx, id, obj)
}

func (s *orderItemService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.orderItemRepo.Delete(ctx, id)
}
