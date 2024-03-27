package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// OrderLineItemService describes the service.
type OrderItemService interface {
	Fetch(ctx context.Context) ([]*ent.OrderItem, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderItem, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.OrderItem) (*ent.OrderItem, error)
	Store(ctx context.Context, obj *ent.OrderItem) (*ent.OrderItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// OrderLineItemRepository describes the repository.
type OrderItemRepository interface {
	Fetch(ctx context.Context) ([]*ent.OrderItem, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderItem, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.OrderItem) (*ent.OrderItem, error)
	Store(ctx context.Context, obj *ent.OrderItem) (*ent.OrderItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
