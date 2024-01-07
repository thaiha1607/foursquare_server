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

func NewProductImageService(p interfaces.ProductImageRepository) interfaces.ProductImageService {
	return &productImageService{
		productImageRepo: p,
	}
}

func (p *productImageService) Fetch(ctx context.Context) ([]*ent.ProductImage, error) {
	return p.productImageRepo.Fetch(ctx)
}

func (p *productImageService) GetByID(ctx context.Context, id uuid.UUID) (*ent.ProductImage, error) {
	return p.productImageRepo.GetByID(ctx, id)
}

func (p *productImageService) Store(ctx context.Context, obj *ent.ProductImage) error {
	return p.productImageRepo.Store(ctx, obj)
}

func (p *productImageService) Delete(ctx context.Context, id uuid.UUID) error {
	return p.productImageRepo.Delete(ctx, id)
}
