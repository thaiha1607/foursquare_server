package interfaces

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
)

// AccountService describes the service.
type AccountService interface {
	Fetch(ctx context.Context) ([]*ent.Account, error)
	GetByID(ctx context.Context, id string) (*ent.Account, error)
	Update(ctx context.Context, id string, obj *ent.Account) (*ent.Account, error)
	Store(ctx context.Context, obj *ent.Account) (*ent.Account, error)
	Delete(ctx context.Context, id string) error
}

// AccountRepository describes the repository.
type AccountRepository interface {
	Fetch(ctx context.Context) ([]*ent.Account, error)
	GetByID(ctx context.Context, id string) (*ent.Account, error)
	Update(ctx context.Context, id string, obj *ent.Account) (*ent.Account, error)
	Store(ctx context.Context, obj *ent.Account) (*ent.Account, error)
	Delete(ctx context.Context, id string) error
}
