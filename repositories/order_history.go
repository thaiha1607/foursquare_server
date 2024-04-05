package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/ent/orderhistory"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entOrderHistoryRepository struct {
	Client *ent.Client
}

func (e *entOrderHistoryRepository) GetByOrderID(ctx context.Context, order_id uuid.UUID) ([]*ent.OrderHistory, error) {
	order_history, err := e.Client.OrderHistory.
		Query().
		Where(
			orderhistory.OrderID(order_id),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return order_history, nil
}

func (e *entOrderHistoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderHistory, error) {
	order_history_record, err := e.Client.OrderHistory.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return order_history_record, nil
}

func (e *entOrderHistoryRepository) Store(ctx context.Context, obj *ent.OrderHistory) (res *ent.OrderHistory, err error) {
	res, err = e.Client.OrderHistory.
		Create().
		SetOrderID(obj.OrderID).
		SetPersonID(obj.PersonID).
		SetNillableOldStatusCode(obj.OldStatusCode).
		SetNillableNewStatusCode(obj.NewStatusCode).
		SetNillableDescription(obj.Description).
		Save(ctx)
	return
}

func (e *entOrderHistoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.OrderHistory.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntOrderHistoryRepository(client *ent.Client) interfaces.OrderHistoryRepository {
	return &entOrderHistoryRepository{Client: client}
}
