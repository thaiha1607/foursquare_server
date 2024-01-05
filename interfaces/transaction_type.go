package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// TransactionTypeService describes the service.
type TransactionTypeService interface {
	Fetch(ctx context.Context) ([]*ent.TransactionType, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.TransactionType, error)
	Update(ctx context.Context, obj *ent.TransactionType) error
	Store(context.Context, *ent.TransactionType) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TransactionTypeRepository describes the repository.
type TransactionTypeRepository interface {
	Fetch(ctx context.Context) ([]*ent.TransactionType, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.TransactionType, error)
	Update(ctx context.Context, obj *ent.TransactionType) error
	Store(context.Context, *ent.TransactionType) error
	Delete(ctx context.Context, id uuid.UUID) error
}
