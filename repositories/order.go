package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entOrderRepository struct {
	Client *ent.Client
}

func (e *entOrderRepository) Fetch(ctx context.Context) ([]*ent.Order, error) {
	orders, err := e.Client.Order.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (e *entOrderRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Order, error) {
	order, err := e.Client.Order.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (e *entOrderRepository) Store(ctx context.Context, obj *ent.Order) (res *ent.Order, err error) {
	res, err = e.Client.Order.
		Create().
		SetCustomerID(obj.CustomerID).
		SetNillableNote(obj.Note).
		SetCreatedBy(obj.CreatedBy).
		SetNillableParentOrderID(obj.ParentOrderID).
		SetNillablePriority(obj.Priority).
		SetNillableType(obj.Type).
		SetNillableStatusCode(&obj.StatusCode).
		SetNillableStaffID(obj.StaffID).
		SetNillableInternalNote(obj.InternalNote).
		SetNillableIsInternal(obj.IsInternal).
		SetAddressID(obj.AddressID).
		Save(ctx)
	return
}

func (e *entOrderRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.Order) (res *ent.Order, err error) {
	res, err = e.Client.Order.
		UpdateOneID(id).
		SetNillableNote(obj.Note).
		SetNillablePriority(obj.Priority).
		SetNillableStatusCode(&obj.StatusCode).
		SetNillableStaffID(obj.StaffID).
		SetNillableInternalNote(obj.InternalNote).
		SetNillableAddressID(&obj.AddressID).
		Save(ctx)
	return
}

func (e *entOrderRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.Order.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntOrderRepository(client *ent.Client) interfaces.OrderRepository {
	return &entOrderRepository{Client: client}
}
