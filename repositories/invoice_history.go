package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/ent/invoicehistory"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entInvoiceHistoryRepository struct {
	Client *ent.Client
}

func (e *entInvoiceHistoryRepository) GetByInvoiceID(ctx context.Context, invoice_id uuid.UUID) ([]*ent.InvoiceHistory, error) {
	invoice_history, err := e.Client.InvoiceHistory.
		Query().
		Where(
			invoicehistory.InvoiceID(invoice_id),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return invoice_history, nil
}

func (e *entInvoiceHistoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.InvoiceHistory, error) {
	invoice_history_record, err := e.Client.InvoiceHistory.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return invoice_history_record, nil
}

func (e *entInvoiceHistoryRepository) Store(ctx context.Context, obj *ent.InvoiceHistory) (res *ent.InvoiceHistory, err error) {
	res, err = e.Client.InvoiceHistory.
		Create().
		SetInvoiceID(obj.InvoiceID).
		SetPersonID(obj.PersonID).
		SetNillableOldStatusCode(obj.OldStatusCode).
		SetNillableNewStatusCode(obj.NewStatusCode).
		Save(ctx)
	return
}

func (e *entInvoiceHistoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.InvoiceHistory.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntInvoiceHistoryRepository(client *ent.Client) interfaces.InvoiceHistoryRepository {
	return &entInvoiceHistoryRepository{Client: client}
}
