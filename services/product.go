package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type productService struct {
	productRepo interfaces.ProductRepository
}

func NewProductService(p interfaces.ProductRepository) interfaces.ProductService {
	return &productService{
		productRepo: p,
	}
}

func (p *productService) Fetch(ctx context.Context) ([]*ent.Product, error) {
	return p.productRepo.Fetch(ctx)
}

func (p *productService) GetByID(ctx context.Context, id string) (*ent.Product, error) {
	return p.productRepo.GetByID(ctx, id)
}

func (p *productService) Store(ctx context.Context, obj *ent.Product) error {
	return p.productRepo.Store(ctx, obj)
}

func (p *productService) Update(ctx context.Context, id string, obj *ent.Product) error {
	return p.productRepo.Update(ctx, id, obj)
}

func (p *productService) Delete(ctx context.Context, id string) error {
	return p.productRepo.Delete(ctx, id)
}
