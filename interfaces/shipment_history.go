package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// ShipmentHistoryService describes the service.
type ShipmentHistoryService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*ent.ShipmentHistory, error)
	Store(ctx context.Context, obj *ent.ShipmentHistory) (*ent.ShipmentHistory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// ShipmentHistoryRepository describes the repository.
type ShipmentHistoryRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*ent.ShipmentHistory, error)
	Store(ctx context.Context, obj *ent.ShipmentHistory) (*ent.ShipmentHistory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
