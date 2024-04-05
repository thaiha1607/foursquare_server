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

func NewOrderService(r interfaces.OrderRepository) interfaces.OrderService {
	return &orderService{
		orderRepo: r,
	}
}

func (s *orderService) Fetch(ctx context.Context) ([]*ent.Order, error) {
	return s.orderRepo.Fetch(ctx)
}

func (s *orderService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Order, error) {
	return s.orderRepo.GetByID(ctx, id)
}

func (s *orderService) Store(ctx context.Context, obj *ent.Order) (*ent.Order, error) {
	return s.orderRepo.Store(ctx, obj)
}

func (s *orderService) Update(ctx context.Context, id uuid.UUID, obj *ent.Order) (*ent.Order, error) {
	return s.orderRepo.Update(ctx, id, obj)
}

func (s *orderService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.orderRepo.Delete(ctx, id)
}
