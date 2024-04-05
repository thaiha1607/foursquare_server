package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// ShipmentItemService describes the service.
type ShipmentItemService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*ent.ShipmentItem, error)
	Store(ctx context.Context, obj *ent.ShipmentItem) (*ent.ShipmentItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// ShipmentItemRepository describes the repository.
type ShipmentItemRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*ent.ShipmentItem, error)
	Store(ctx context.Context, obj *ent.ShipmentItem) (*ent.ShipmentItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
