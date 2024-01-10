package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entFinancialTransactionRepository struct {
	Client *ent.Client
}

func (e *entFinancialTransactionRepository) Fetch(ctx context.Context) ([]*ent.FinancialTransaction, error) {
	conversations, err := e.Client.FinancialTransaction.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return conversations, nil
}

func (e *entFinancialTransactionRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.FinancialTransaction, error) {
	conversation, err := e.Client.FinancialTransaction.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return conversation, nil
}

func (e *entFinancialTransactionRepository) Store(ctx context.Context, obj *ent.FinancialTransaction) (err error) {
	//lint:ignore SA4006 we want to return the result of creating operation
	obj, err = e.Client.FinancialTransaction.
		Create().
		SetInvoiceID(obj.InvoiceID).
		SetAmount(obj.Amount).
		SetNillableComment(obj.Comment).
		SetNillableIsInternal(&obj.IsInternal).
		SetNillablePaymentMethod(&obj.PaymentMethod).
		Save(ctx)
	return
}

func (e *entFinancialTransactionRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.FinancialTransaction) (err error) {
	//lint:ignore SA4006 we want to return the result of updating operation
	obj, err = e.Client.FinancialTransaction.
		UpdateOneID(id).
		SetNillableComment(obj.Comment).
		Save(ctx)
	return
}

func (e *entFinancialTransactionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.FinancialTransaction.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntFinancialTransactionRepository(client *ent.Client) interfaces.FinancialTransactionRepository {
	return &entFinancialTransactionRepository{Client: client}
}
