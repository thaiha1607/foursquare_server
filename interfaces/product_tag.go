package interfaces

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
)

// ProductTagService describes the service.
type ProductTagService interface {
	Fetch(ctx context.Context) ([]*ent.ProductTag, error)
	Store(ctx context.Context, obj *ent.ProductTag) (*ent.ProductTag, error)
	Delete(ctx context.Context, product_id string, tag_id string) error
}

// ProductTagRepository describes the repository.
type ProductTagRepository interface {
	Fetch(ctx context.Context) ([]*ent.ProductTag, error)
	Store(ctx context.Context, obj *ent.ProductTag) (*ent.ProductTag, error)
	Delete(ctx context.Context, product_id string, tag_id string) error
}
