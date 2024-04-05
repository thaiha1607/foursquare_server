package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type productQtyService struct {
	productQtyRepo interfaces.ProductQtyRepository
}

func NewProductQtyService(r interfaces.ProductQtyRepository) interfaces.ProductQtyService {
	return &productQtyService{
		productQtyRepo: r,
	}
}

func (s *productQtyService) Fetch(ctx context.Context) ([]*ent.ProductQty, error) {
	return s.productQtyRepo.Fetch(ctx)
}

func (s *productQtyService) GetByID(ctx context.Context, id uuid.UUID) (*ent.ProductQty, error) {
	return s.productQtyRepo.GetByID(ctx, id)
}

func (s *productQtyService) Store(ctx context.Context, obj *ent.ProductQty) (*ent.ProductQty, error) {
	return s.productQtyRepo.Store(ctx, obj)
}

func (s *productQtyService) Update(ctx context.Context, id uuid.UUID, obj *ent.ProductQty) (*ent.ProductQty, error) {
	return s.productQtyRepo.Update(ctx, id, obj)
}

func (s *productQtyService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.productQtyRepo.Delete(ctx, id)
}
