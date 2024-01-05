package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// TagService describes the service.
type TagService interface {
	Fetch(ctx context.Context) ([]*ent.Tag, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Tag, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.Tag) error
	Store(ctx context.Context, obj *ent.Tag) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TagRepository describes the repository.
type TagRepository interface {
	Fetch(ctx context.Context) ([]*ent.Tag, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Tag, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.Tag) error
	Store(ctx context.Context, obj *ent.Tag) error
	Delete(ctx context.Context, id uuid.UUID) error
}
