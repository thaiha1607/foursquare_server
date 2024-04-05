package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type invoiceService struct {
	invoiceRepo interfaces.InvoiceRepository
}

func NewInvoiceService(r interfaces.InvoiceRepository) interfaces.InvoiceService {
	return &invoiceService{
		invoiceRepo: r,
	}
}

func (s *invoiceService) Fetch(ctx context.Context) ([]*ent.Invoice, error) {
	return s.invoiceRepo.Fetch(ctx)
}

func (s *invoiceService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Invoice, error) {
	return s.invoiceRepo.GetByID(ctx, id)
}

func (s *invoiceService) Store(ctx context.Context, obj *ent.Invoice) (*ent.Invoice, error) {
	return s.invoiceRepo.Store(ctx, obj)
}

func (s *invoiceService) Update(ctx context.Context, id uuid.UUID, obj *ent.Invoice) (*ent.Invoice, error) {
	return s.invoiceRepo.Update(ctx, id, obj)
}

func (s *invoiceService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.invoiceRepo.Delete(ctx, id)
}
