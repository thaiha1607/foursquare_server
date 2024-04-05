package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type productColorService struct {
	productColorRepo interfaces.ProductColorRepository
}

func NewProductColorService(r interfaces.ProductColorRepository) interfaces.ProductColorService {
	return &productColorService{
		productColorRepo: r,
	}
}

func (s *productColorService) Fetch(ctx context.Context) ([]*ent.ProductColor, error) {
	return s.productColorRepo.Fetch(ctx)
}

func (s *productColorService) GetByID(ctx context.Context, id string) (*ent.ProductColor, error) {
	return s.productColorRepo.GetByID(ctx, id)
}

func (s *productColorService) Store(ctx context.Context, obj *ent.ProductColor) (*ent.ProductColor, error) {
	return s.productColorRepo.Store(ctx, obj)
}

func (s *productColorService) Update(ctx context.Context, id string, obj *ent.ProductColor) (*ent.ProductColor, error) {
	return s.productColorRepo.Update(ctx, id, obj)
}

func (s *productColorService) Delete(ctx context.Context, id string) error {
	return s.productColorRepo.Delete(ctx, id)
}
