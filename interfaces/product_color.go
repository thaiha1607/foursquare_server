package interfaces

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
)

// ProductColorService describes the service.
type ProductColorService interface {
	Fetch(ctx context.Context) ([]*ent.ProductColor, error)
	GetByID(ctx context.Context, id string) (*ent.ProductColor, error)
	Update(ctx context.Context, id string, obj *ent.ProductColor) (*ent.ProductColor, error)
	Store(ctx context.Context, obj *ent.ProductColor) (*ent.ProductColor, error)
	Delete(ctx context.Context, id string) error
}

// ProductColorRepository describes the repository.
type ProductColorRepository interface {
	Fetch(ctx context.Context) ([]*ent.ProductColor, error)
	GetByID(ctx context.Context, id string) (*ent.ProductColor, error)
	Update(ctx context.Context, id string, obj *ent.ProductColor) (*ent.ProductColor, error)
	Store(ctx context.Context, obj *ent.ProductColor) (*ent.ProductColor, error)
	Delete(ctx context.Context, id string) error
}
