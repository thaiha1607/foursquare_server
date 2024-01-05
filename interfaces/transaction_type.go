package interfaces

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
)

// TransactionTypeService describes the service.
type TransactionTypeService interface {
	Fetch(ctx context.Context) ([]*ent.TransactionType, error)
	GetByID(ctx context.Context, id int) (*ent.TransactionType, error)
	Update(ctx context.Context, obj *ent.TransactionType) error
	Store(context.Context, *ent.TransactionType) error
	Delete(ctx context.Context, id int) error
}

// TransactionTypeRepository describes the repository.
type TransactionTypeRepository interface {
	Fetch(ctx context.Context) ([]*ent.TransactionType, error)
	GetByID(ctx context.Context, id int) (*ent.TransactionType, error)
	Update(ctx context.Context, obj *ent.TransactionType) error
	Store(context.Context, *ent.TransactionType) error
	Delete(ctx context.Context, id int) error
}
