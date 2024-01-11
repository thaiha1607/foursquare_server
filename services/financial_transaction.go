package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type financialTransactionService struct {
	financialTransactionRepo interfaces.FinancialTransactionRepository
}

func NewFinancialTransactionService(f interfaces.FinancialTransactionRepository) interfaces.FinancialTransactionService {
	return &financialTransactionService{
		financialTransactionRepo: f,
	}
}

func (f *financialTransactionService) Fetch(ctx context.Context) ([]*ent.FinancialTransaction, error) {
	return f.financialTransactionRepo.Fetch(ctx)
}

func (f *financialTransactionService) GetByID(ctx context.Context, id uuid.UUID) (*ent.FinancialTransaction, error) {
	return f.financialTransactionRepo.GetByID(ctx, id)
}

func (f *financialTransactionService) Store(ctx context.Context, obj *ent.FinancialTransaction) (*ent.FinancialTransaction, error) {
	return f.financialTransactionRepo.Store(ctx, obj)
}

func (f *financialTransactionService) Update(ctx context.Context, id uuid.UUID, obj *ent.FinancialTransaction) (*ent.FinancialTransaction, error) {
	return f.financialTransactionRepo.Update(ctx, id, obj)
}

func (f *financialTransactionService) Delete(ctx context.Context, id uuid.UUID) error {
	return f.financialTransactionRepo.Delete(ctx, id)
}
