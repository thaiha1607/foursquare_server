package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type productTagService struct {
	productTagRepo interfaces.ProductTagRepository
}

func NewProductTagService(p interfaces.ProductTagRepository) interfaces.ProductTagService {
	return &productTagService{
		productTagRepo: p,
	}
}

func (p *productTagService) Fetch(ctx context.Context) ([]*ent.ProductTag, error) {
	return p.productTagRepo.Fetch(ctx)
}

func (p *productTagService) Store(ctx context.Context, obj *ent.ProductTag) (*ent.ProductTag, error) {
	return p.productTagRepo.Store(ctx, obj)
}

func (p *productTagService) Delete(ctx context.Context, product_id string, tag_id string) error {
	return p.productTagRepo.Delete(ctx, product_id, tag_id)
}
