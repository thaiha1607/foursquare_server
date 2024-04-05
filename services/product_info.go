package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type productInfoService struct {
	productInfoRepo interfaces.ProductInfoRepository
}

func NewProductInfoService(r interfaces.ProductInfoRepository) interfaces.ProductInfoService {
	return &productInfoService{
		productInfoRepo: r,
	}
}

func (s *productInfoService) Fetch(ctx context.Context) ([]*ent.ProductInfo, error) {
	return s.productInfoRepo.Fetch(ctx)
}

func (s *productInfoService) GetByID(ctx context.Context, id string) (*ent.ProductInfo, error) {
	return s.productInfoRepo.GetByID(ctx, id)
}

func (s *productInfoService) Store(ctx context.Context, obj *ent.ProductInfo) (*ent.ProductInfo, error) {
	return s.productInfoRepo.Store(ctx, obj)
}

func (s *productInfoService) Update(ctx context.Context, id string, obj *ent.ProductInfo) (*ent.ProductInfo, error) {
	return s.productInfoRepo.Update(ctx, id, obj)
}

func (s *productInfoService) Delete(ctx context.Context, id string) error {
	return s.productInfoRepo.Delete(ctx, id)
}
