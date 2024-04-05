package repositories

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entInvoiceStatusCodeRepository struct {
	Client *ent.Client
}

func (e *entInvoiceStatusCodeRepository) Fetch(ctx context.Context) ([]*ent.InvoiceStatusCode, error) {
	invoice_status_codes, err := e.Client.InvoiceStatusCode.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return invoice_status_codes, nil
}

func (e *entInvoiceStatusCodeRepository) GetByID(ctx context.Context, id int) (*ent.InvoiceStatusCode, error) {
	invoice_status_code, err := e.Client.InvoiceStatusCode.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return invoice_status_code, nil
}

func (e *entInvoiceStatusCodeRepository) Store(ctx context.Context, obj *ent.InvoiceStatusCode) (res *ent.InvoiceStatusCode, err error) {
	res, err = e.Client.InvoiceStatusCode.
		Create().
		SetInvoiceStatus(obj.InvoiceStatus).
		Save(ctx)
	return
}

func (e *entInvoiceStatusCodeRepository) Update(ctx context.Context, id int, obj *ent.InvoiceStatusCode) (res *ent.InvoiceStatusCode, err error) {
	res, err = e.Client.InvoiceStatusCode.
		UpdateOneID(id).
		SetNillableInvoiceStatus(&obj.InvoiceStatus).
		Save(ctx)
	return
}

func (e *entInvoiceStatusCodeRepository) Delete(ctx context.Context, id int) error {
	return e.Client.InvoiceStatusCode.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntInvoiceStatusCodeRepository(client *ent.Client) interfaces.InvoiceStatusCodeRepository {
	return &entInvoiceStatusCodeRepository{Client: client}
}
