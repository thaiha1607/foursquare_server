package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// AddressService describes the service.
type AddressService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Address, error)
	Store(ctx context.Context, obj *ent.Address) (*ent.Address, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// AddressRepository describes the repository.
type AddressRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Address, error)
	Store(ctx context.Context, obj *ent.Address) (*ent.Address, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
