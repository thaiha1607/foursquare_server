package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// OrderLineItemService describes the service.
type OrderLineItemService interface {
	Fetch(ctx context.Context) ([]*ent.OrderLineItem, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderLineItem, error)
	Update(ctx context.Context, obj *ent.OrderLineItem) error
	Store(context.Context, *ent.OrderLineItem) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// OrderLineItemRepository describes the repository.
type OrderLineItemRepository interface {
	Fetch(ctx context.Context) ([]*ent.OrderLineItem, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderLineItem, error)
	Update(ctx context.Context, obj *ent.OrderLineItem) error
	Store(context.Context, *ent.OrderLineItem) error
	Delete(ctx context.Context, id uuid.UUID) error
}
