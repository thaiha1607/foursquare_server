package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// PersonService describes the service.
type PersonService interface {
	Fetch(ctx context.Context) ([]*ent.Person, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Person, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.Person) (*ent.Person, error)
	Store(ctx context.Context, obj *ent.Person) (*ent.Person, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// PersonRepository describes the repository.
type PersonRepository interface {
	Fetch(ctx context.Context) ([]*ent.Person, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Person, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.Person) (*ent.Person, error)
	Store(ctx context.Context, obj *ent.Person) (*ent.Person, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
