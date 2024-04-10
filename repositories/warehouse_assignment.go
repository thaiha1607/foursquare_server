package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entWarehouseAssignmentRepository struct {
	Client *ent.Client
}

func (e *entWarehouseAssignmentRepository) Fetch(ctx context.Context) ([]*ent.WarehouseAssignment, error) {
	warehouse_assignments, err := e.Client.WarehouseAssignment.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return warehouse_assignments, nil
}

func (e *entWarehouseAssignmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.WarehouseAssignment, error) {
	warehouse_assignment, err := e.Client.WarehouseAssignment.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return warehouse_assignment, nil
}

func (e *entWarehouseAssignmentRepository) Store(ctx context.Context, obj *ent.WarehouseAssignment) (res *ent.WarehouseAssignment, err error) {
	res, err = e.Client.WarehouseAssignment.
		Create().
		SetOrderID(obj.OrderID).
		SetWorkUnitID(obj.WorkUnitID).
		SetNillableStaffID(obj.StaffID).
		SetNillableStatus(obj.Status).
		SetNillableNote(obj.Note).
		Save(ctx)
	return
}

func (e *entWarehouseAssignmentRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.WarehouseAssignment) (res *ent.WarehouseAssignment, err error) {
	res, err = e.Client.WarehouseAssignment.
		UpdateOneID(id).
		SetNillableStaffID(obj.StaffID).
		SetNillableStatus(obj.Status).
		SetNillableNote(obj.Note).
		Save(ctx)
	return
}

func (e *entWarehouseAssignmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.WarehouseAssignment.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntWarehouseAssignmentRepository(client *ent.Client) interfaces.WarehouseAssignmentRepository {
	return &entWarehouseAssignmentRepository{Client: client}
}
