package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entInvoiceLineItemRepository struct {
	Client *ent.Client
}

func (e *entInvoiceLineItemRepository) Fetch(ctx context.Context) ([]*ent.InvoiceLineItem, error) {
	invoice_line_items, err := e.Client.InvoiceLineItem.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return invoice_line_items, nil
}

func (e *entInvoiceLineItemRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.InvoiceLineItem, error) {
	invoice_line_item, err := e.Client.InvoiceLineItem.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return invoice_line_item, nil
}

func (e *entInvoiceLineItemRepository) Store(ctx context.Context, obj *ent.InvoiceLineItem) (err error) {
	//lint:ignore SA4006 we want to return the result of creating operation
	obj, err = e.Client.InvoiceLineItem.
		Create().
		SetInvoiceID(obj.InvoiceID).
		SetOrderLineItemID(obj.OrderLineItemID).
		SetQty(obj.Qty).
		SetTotal(obj.Total).
		Save(ctx)
	return
}

func (e *entInvoiceLineItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.InvoiceLineItem.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntInvoiceLineItemRepository(client *ent.Client) interfaces.InvoiceLineItemRepository {
	return &entInvoiceLineItemRepository{Client: client}
}
