package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type orderStatusCodeService struct {
	orderStatusCodeRepo interfaces.OrderStatusCodeRepository
}

func NewOrderStatusCodeService(r interfaces.OrderStatusCodeRepository) interfaces.OrderStatusCodeService {
	return &orderStatusCodeService{
		orderStatusCodeRepo: r,
	}
}

func (s *orderStatusCodeService) Fetch(ctx context.Context) ([]*ent.OrderStatusCode, error) {
	return s.orderStatusCodeRepo.Fetch(ctx)
}

func (s *orderStatusCodeService) GetByID(ctx context.Context, id int) (*ent.OrderStatusCode, error) {
	return s.orderStatusCodeRepo.GetByID(ctx, id)
}

func (s *orderStatusCodeService) Store(ctx context.Context, obj *ent.OrderStatusCode) (*ent.OrderStatusCode, error) {
	return s.orderStatusCodeRepo.Store(ctx, obj)
}

func (s *orderStatusCodeService) Update(ctx context.Context, id int, obj *ent.OrderStatusCode) (*ent.OrderStatusCode, error) {
	return s.orderStatusCodeRepo.Update(ctx, id, obj)
}

func (s *orderStatusCodeService) Delete(ctx context.Context, id int) error {
	return s.orderStatusCodeRepo.Delete(ctx, id)
}
