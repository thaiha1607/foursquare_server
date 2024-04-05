package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type orderHistoryService struct {
	orderHistoryRepo interfaces.OrderHistoryRepository
}

func NewOrderHistoryService(r interfaces.OrderHistoryRepository) interfaces.OrderHistoryService {
	return &orderHistoryService{
		orderHistoryRepo: r,
	}
}

func (s *orderHistoryService) GetByOrderID(ctx context.Context, order_id uuid.UUID) ([]*ent.OrderHistory, error) {
	return s.orderHistoryRepo.GetByOrderID(ctx, order_id)
}

func (s *orderHistoryService) GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderHistory, error) {
	return s.orderHistoryRepo.GetByID(ctx, id)
}

func (s *orderHistoryService) Store(ctx context.Context, obj *ent.OrderHistory) (*ent.OrderHistory, error) {
	return s.orderHistoryRepo.Store(ctx, obj)
}

func (s *orderHistoryService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.orderHistoryRepo.Delete(ctx, id)
}
