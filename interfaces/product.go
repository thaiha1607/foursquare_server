package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// ProductService describes the service.
type ProductService interface {
	Fetch(ctx context.Context) ([]*ent.Product, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Product, error)
	Update(ctx context.Context, obj *ent.Product) error
	Store(context.Context, *ent.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// ProductRepository describes the repository.
type ProductRepository interface {
	Fetch(ctx context.Context) ([]*ent.Product, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Product, error)
	Update(ctx context.Context, obj *ent.Product) error
	Store(context.Context, *ent.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}
