package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// FinancialTransactionService describes the service.
type FinancialTransactionService interface {
	Fetch(ctx context.Context) ([]*ent.FinancialTransaction, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.FinancialTransaction, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.FinancialTransaction) (*ent.FinancialTransaction, error)
	Store(ctx context.Context, obj *ent.FinancialTransaction) (*ent.FinancialTransaction, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// FinancialTransactionRepository describes the repository.
type FinancialTransactionRepository interface {
	Fetch(ctx context.Context) ([]*ent.FinancialTransaction, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.FinancialTransaction, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.FinancialTransaction) (*ent.FinancialTransaction, error)
	Store(ctx context.Context, obj *ent.FinancialTransaction) (*ent.FinancialTransaction, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
