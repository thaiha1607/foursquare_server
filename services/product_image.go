package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type productImageService struct {
	productImageRepo interfaces.ProductImageRepository
}

func NewProductImageService(r interfaces.ProductImageRepository) interfaces.ProductImageService {
	return &productImageService{
		productImageRepo: r,
	}
}

func (s *productImageService) Fetch(ctx context.Context) ([]*ent.ProductImage, error) {
	return s.productImageRepo.Fetch(ctx)
}

func (s *productImageService) GetByID(ctx context.Context, id uuid.UUID) (*ent.ProductImage, error) {
	return s.productImageRepo.GetByID(ctx, id)
}

func (s *productImageService) Store(ctx context.Context, obj *ent.ProductImage) (*ent.ProductImage, error) {
	return s.productImageRepo.Store(ctx, obj)
}

func (s *productImageService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.productImageRepo.Delete(ctx, id)
}
