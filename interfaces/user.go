package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// UserService describes the service.
type UserService interface {
	Fetch(ctx context.Context) ([]*ent.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.User) (*ent.User, error)
	Store(ctx context.Context, obj *ent.User) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// UserRepository describes the repository.
type UserRepository interface {
	Fetch(ctx context.Context) ([]*ent.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.User) (*ent.User, error)
	Store(ctx context.Context, obj *ent.User) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
