package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type productTagService struct {
	productTagRepo interfaces.ProductTagRepository
}

func NewProductTagService(r interfaces.ProductTagRepository) interfaces.ProductTagService {
	return &productTagService{
		productTagRepo: r,
	}
}

func (s *productTagService) Fetch(ctx context.Context) ([]*ent.ProductTag, error) {
	return s.productTagRepo.Fetch(ctx)
}

func (s *productTagService) Store(ctx context.Context, obj *ent.ProductTag) (*ent.ProductTag, error) {
	return s.productTagRepo.Store(ctx, obj)
}

func (s *productTagService) Delete(ctx context.Context, product_id string, tag_id string) error {
	return s.productTagRepo.Delete(ctx, product_id, tag_id)
}
