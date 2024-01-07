package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type invoiceLineItemService struct {
	invoiceLineItemRepo interfaces.InvoiceLineItemRepository
}

func NewInvoiceLineItemService(i interfaces.InvoiceLineItemRepository) interfaces.InvoiceLineItemService {
	return &invoiceLineItemService{
		invoiceLineItemRepo: i,
	}
}

func (i *invoiceLineItemService) Fetch(ctx context.Context) ([]*ent.InvoiceLineItem, error) {
	return i.invoiceLineItemRepo.Fetch(ctx)
}

func (i *invoiceLineItemService) GetByID(ctx context.Context, id uuid.UUID) (*ent.InvoiceLineItem, error) {
	return i.invoiceLineItemRepo.GetByID(ctx, id)
}

func (i *invoiceLineItemService) Store(ctx context.Context, obj *ent.InvoiceLineItem) error {
	return i.invoiceLineItemRepo.Store(ctx, obj)
}

func (i *invoiceLineItemService) Delete(ctx context.Context, id uuid.UUID) error {
	return i.invoiceLineItemRepo.Delete(ctx, id)
}
