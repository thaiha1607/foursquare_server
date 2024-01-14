package interfaces

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
)

// TagService describes the service.
type TagService interface {
	Fetch(ctx context.Context) ([]*ent.Tag, error)
	GetByID(ctx context.Context, id string) (*ent.Tag, error)
	Update(ctx context.Context, id string, obj *ent.Tag) (*ent.Tag, error)
	Store(ctx context.Context, obj *ent.Tag) (*ent.Tag, error)
	Delete(ctx context.Context, id string) error
}

// TagRepository describes the repository.
type TagRepository interface {
	Fetch(ctx context.Context) ([]*ent.Tag, error)
	GetByID(ctx context.Context, id string) (*ent.Tag, error)
	Update(ctx context.Context, id string, obj *ent.Tag) (*ent.Tag, error)
	Store(ctx context.Context, obj *ent.Tag) (*ent.Tag, error)
	Delete(ctx context.Context, id string) error
}
