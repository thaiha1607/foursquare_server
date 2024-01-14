package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// ProductImageService describes the service.
type ProductImageService interface {
	Fetch(ctx context.Context) ([]*ent.ProductImage, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.ProductImage, error)
	Store(ctx context.Context, obj *ent.ProductImage) (*ent.ProductImage, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// ProductImageRepository describes the repository.
type ProductImageRepository interface {
	Fetch(ctx context.Context) ([]*ent.ProductImage, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.ProductImage, error)
	Store(ctx context.Context, obj *ent.ProductImage) (*ent.ProductImage, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
