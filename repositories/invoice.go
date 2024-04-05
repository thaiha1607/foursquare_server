package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entInvoiceRepository struct {
	Client *ent.Client
}

func (e *entInvoiceRepository) Fetch(ctx context.Context) ([]*ent.Invoice, error) {
	invoices, err := e.Client.Invoice.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (e *entInvoiceRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Invoice, error) {
	invoice, err := e.Client.Invoice.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

func (e *entInvoiceRepository) Store(ctx context.Context, obj *ent.Invoice) (res *ent.Invoice, err error) {
	res, err = e.Client.Invoice.
		Create().
		SetOrderID(obj.OrderID).
		SetTotal(obj.Total).
		SetNillableNote(obj.Note).
		SetNillableType(&obj.Type).
		SetNillableStatusCode(&obj.StatusCode).
		SetNillablePaymentMethod(obj.PaymentMethod).
		Save(ctx)
	return
}

func (e *entInvoiceRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.Invoice) (res *ent.Invoice, err error) {
	res, err = e.Client.Invoice.
		UpdateOneID(id).
		SetNillableNote(obj.Note).
		SetNillableStatusCode(&obj.StatusCode).
		SetNillablePaymentMethod(obj.PaymentMethod).
		Save(ctx)
	return
}

func (e *entInvoiceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.Invoice.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntInvoiceRepository(client *ent.Client) interfaces.InvoiceRepository {
	return &entInvoiceRepository{Client: client}
}
