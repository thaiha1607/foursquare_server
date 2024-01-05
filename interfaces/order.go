package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// OrderService describes the service.
type OrderService interface {
	Fetch(ctx context.Context) ([]*ent.Order, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Order, error)
	Update(ctx context.Context, obj *ent.Order) error
	Store(context.Context, *ent.Order) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// OrderRepository describes the repository.
type OrderRepository interface {
	Fetch(ctx context.Context) ([]*ent.Order, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Order, error)
	Update(ctx context.Context, obj *ent.Order) error
	Store(context.Context, *ent.Order) error
	Delete(ctx context.Context, id uuid.UUID) error
}
