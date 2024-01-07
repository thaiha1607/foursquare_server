package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// ProductTagService describes the service.
type ProductTagService interface {
	Fetch(ctx context.Context) ([]*ent.ProductTag, error)
	Store(ctx context.Context, obj *ent.ProductTag) error
	Delete(ctx context.Context, product_id uuid.UUID, tag_id uuid.UUID) error
}

// ProductTagRepository describes the repository.
type ProductTagRepository interface {
	Fetch(ctx context.Context) ([]*ent.ProductTag, error)
	Store(ctx context.Context, obj *ent.ProductTag) error
	Delete(ctx context.Context, product_id uuid.UUID, tag_id uuid.UUID) error
}
