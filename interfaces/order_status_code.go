package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// OrderStatusCodeService describes the service.
type OrderStatusCodeService interface {
	Fetch(ctx context.Context) ([]*ent.OrderStatusCode, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderStatusCode, error)
	Update(ctx context.Context, obj *ent.OrderStatusCode) error
	Store(context.Context, *ent.OrderStatusCode) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// OrderStatusCodeRepository describes the repository.
type OrderStatusCodeRepository interface {
	Fetch(ctx context.Context) ([]*ent.OrderStatusCode, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderStatusCode, error)
	Update(ctx context.Context, obj *ent.OrderStatusCode) error
	Store(context.Context, *ent.OrderStatusCode) error
	Delete(ctx context.Context, id uuid.UUID) error
}
