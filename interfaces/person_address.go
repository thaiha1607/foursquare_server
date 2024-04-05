package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// PersonAddressService describes the service.
type PersonAddressService interface {
	Fetch(ctx context.Context) ([]*ent.PersonAddress, error)
	GetByPersonID(ctx context.Context, person_id uuid.UUID) ([]*ent.PersonAddress, error)
	Store(ctx context.Context, obj *ent.PersonAddress) (*ent.PersonAddress, error)
	Delete(ctx context.Context, person_id uuid.UUID, address_id uuid.UUID) error
}

// PersonAddressRepository describes the repository.
type PersonAddressRepository interface {
	Fetch(ctx context.Context) ([]*ent.PersonAddress, error)
	GetByPersonID(ctx context.Context, person_id uuid.UUID) ([]*ent.PersonAddress, error)
	Store(ctx context.Context, obj *ent.PersonAddress) (*ent.PersonAddress, error)
	Delete(ctx context.Context, person_id uuid.UUID, address_id uuid.UUID) error
}
