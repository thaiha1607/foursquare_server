package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// ProductQtyService describes the service.
type ProductQtyService interface {
	Fetch(ctx context.Context) ([]*ent.ProductQty, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.ProductQty, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.ProductQty) (*ent.ProductQty, error)
	Store(ctx context.Context, obj *ent.ProductQty) (*ent.ProductQty, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// ProductQtyRepository describes the repository.
type ProductQtyRepository interface {
	Fetch(ctx context.Context) ([]*ent.ProductQty, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.ProductQty, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.ProductQty) (*ent.ProductQty, error)
	Store(ctx context.Context, obj *ent.ProductQty) (*ent.ProductQty, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
