package interfaces

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
)

// ShipmentStatusCodeService describes the service.
type ShipmentStatusCodeService interface {
	Fetch(ctx context.Context) ([]*ent.ShipmentStatusCode, error)
	GetByID(ctx context.Context, id int) (*ent.ShipmentStatusCode, error)
	Update(ctx context.Context, id int, obj *ent.ShipmentStatusCode) (*ent.ShipmentStatusCode, error)
	Store(ctx context.Context, obj *ent.ShipmentStatusCode) (*ent.ShipmentStatusCode, error)
	Delete(ctx context.Context, id int) error
}

// ShipmentStatusCodeRepository describes the repository.
type ShipmentStatusCodeRepository interface {
	Fetch(ctx context.Context) ([]*ent.ShipmentStatusCode, error)
	GetByID(ctx context.Context, id int) (*ent.ShipmentStatusCode, error)
	Update(ctx context.Context, id int, obj *ent.ShipmentStatusCode) (*ent.ShipmentStatusCode, error)
	Store(ctx context.Context, obj *ent.ShipmentStatusCode) (*ent.ShipmentStatusCode, error)
	Delete(ctx context.Context, id int) error
}
