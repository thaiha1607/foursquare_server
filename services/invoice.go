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

func NewInvoiceService(i interfaces.InvoiceRepository) interfaces.InvoiceService {
	return &invoiceService{
		invoiceRepo: i,
	}
}

func (i *invoiceService) Fetch(ctx context.Context) ([]*ent.Invoice, error) {
	return i.invoiceRepo.Fetch(ctx)
}

func (i *invoiceService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Invoice, error) {
	return i.invoiceRepo.GetByID(ctx, id)
}

func (i *invoiceService) Store(ctx context.Context, obj *ent.Invoice) (*ent.Invoice, error) {
	return i.invoiceRepo.Store(ctx, obj)
}

func (i *invoiceService) Update(ctx context.Context, id uuid.UUID, obj *ent.Invoice) (*ent.Invoice, error) {
	return i.invoiceRepo.Update(ctx, id, obj)
}

func (i *invoiceService) Delete(ctx context.Context, id uuid.UUID) error {
	return i.invoiceRepo.Delete(ctx, id)
}
