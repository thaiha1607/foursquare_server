package interfaces

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
)

// ShipmentService describes the service.
type ShipmentService interface {
	Fetch(ctx context.Context) ([]*ent.Shipment, error)
	GetByID(ctx context.Context, id string) (*ent.Shipment, error)
	Update(ctx context.Context, id string, obj *ent.Shipment) (*ent.Shipment, error)
	Store(ctx context.Context, obj *ent.Shipment) (*ent.Shipment, error)
	Delete(ctx context.Context, id string) error
}

// ShipmentRepository describes the repository.
type ShipmentRepository interface {
	Fetch(ctx context.Context) ([]*ent.Shipment, error)
	GetByID(ctx context.Context, id string) (*ent.Shipment, error)
	Update(ctx context.Context, id string, obj *ent.Shipment) (*ent.Shipment, error)
	Store(ctx context.Context, obj *ent.Shipment) (*ent.Shipment, error)
	Delete(ctx context.Context, id string) error
}
