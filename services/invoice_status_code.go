package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type invoiceStatusCodeService struct {
	invoiceStatusCodeRepo interfaces.InvoiceStatusCodeRepository
}

func NewInvoiceStatusCodeService(r interfaces.InvoiceStatusCodeRepository) interfaces.InvoiceStatusCodeService {
	return &invoiceStatusCodeService{
		invoiceStatusCodeRepo: r,
	}
}

func (s *invoiceStatusCodeService) Fetch(ctx context.Context) ([]*ent.InvoiceStatusCode, error) {
	return s.invoiceStatusCodeRepo.Fetch(ctx)
}

func (s *invoiceStatusCodeService) GetByID(ctx context.Context, id int) (*ent.InvoiceStatusCode, error) {
	return s.invoiceStatusCodeRepo.GetByID(ctx, id)
}

func (s *invoiceStatusCodeService) Store(ctx context.Context, obj *ent.InvoiceStatusCode) (*ent.InvoiceStatusCode, error) {
	return s.invoiceStatusCodeRepo.Store(ctx, obj)
}

func (s *invoiceStatusCodeService) Update(ctx context.Context, id int, obj *ent.InvoiceStatusCode) (*ent.InvoiceStatusCode, error) {
	return s.invoiceStatusCodeRepo.Update(ctx, id, obj)
}

func (s *invoiceStatusCodeService) Delete(ctx context.Context, id int) error {
	return s.invoiceStatusCodeRepo.Delete(ctx, id)
}
