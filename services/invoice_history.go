package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type invoiceHistoryService struct {
	invoiceHistoryRepo interfaces.InvoiceHistoryRepository
}

func NewInvoiceHistoryService(r interfaces.InvoiceHistoryRepository) interfaces.InvoiceHistoryService {
	return &invoiceHistoryService{
		invoiceHistoryRepo: r,
	}
}

func (s *invoiceHistoryService) GetByInvoiceID(ctx context.Context, invoice_id uuid.UUID) ([]*ent.InvoiceHistory, error) {
	return s.invoiceHistoryRepo.GetByInvoiceID(ctx, invoice_id)
}

func (s *invoiceHistoryService) GetByID(ctx context.Context, id uuid.UUID) (*ent.InvoiceHistory, error) {
	return s.invoiceHistoryRepo.GetByID(ctx, id)
}

func (s *invoiceHistoryService) Store(ctx context.Context, obj *ent.InvoiceHistory) (*ent.InvoiceHistory, error) {
	return s.invoiceHistoryRepo.Store(ctx, obj)
}

func (s *invoiceHistoryService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.invoiceHistoryRepo.Delete(ctx, id)
}
