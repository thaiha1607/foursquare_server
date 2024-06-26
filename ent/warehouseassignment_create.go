// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/order"
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/warehouseassignment"
	"github.com/thaiha1607/foursquare_server/ent/workunitinfo"
)

// WarehouseAssignmentCreate is the builder for creating a WarehouseAssignment entity.
type WarehouseAssignmentCreate struct {
	config
	mutation *WarehouseAssignmentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (wac *WarehouseAssignmentCreate) SetCreatedAt(t time.Time) *WarehouseAssignmentCreate {
	wac.mutation.SetCreatedAt(t)
	return wac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wac *WarehouseAssignmentCreate) SetNillableCreatedAt(t *time.Time) *WarehouseAssignmentCreate {
	if t != nil {
		wac.SetCreatedAt(*t)
	}
	return wac
}

// SetUpdatedAt sets the "updated_at" field.
func (wac *WarehouseAssignmentCreate) SetUpdatedAt(t time.Time) *WarehouseAssignmentCreate {
	wac.mutation.SetUpdatedAt(t)
	return wac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wac *WarehouseAssignmentCreate) SetNillableUpdatedAt(t *time.Time) *WarehouseAssignmentCreate {
	if t != nil {
		wac.SetUpdatedAt(*t)
	}
	return wac
}

// SetOrderID sets the "order_id" field.
func (wac *WarehouseAssignmentCreate) SetOrderID(u uuid.UUID) *WarehouseAssignmentCreate {
	wac.mutation.SetOrderID(u)
	return wac
}

// SetWorkUnitID sets the "work_unit_id" field.
func (wac *WarehouseAssignmentCreate) SetWorkUnitID(u uuid.UUID) *WarehouseAssignmentCreate {
	wac.mutation.SetWorkUnitID(u)
	return wac
}

// SetStaffID sets the "staff_id" field.
func (wac *WarehouseAssignmentCreate) SetStaffID(u uuid.UUID) *WarehouseAssignmentCreate {
	wac.mutation.SetStaffID(u)
	return wac
}

// SetNillableStaffID sets the "staff_id" field if the given value is not nil.
func (wac *WarehouseAssignmentCreate) SetNillableStaffID(u *uuid.UUID) *WarehouseAssignmentCreate {
	if u != nil {
		wac.SetStaffID(*u)
	}
	return wac
}

// SetStatus sets the "status" field.
func (wac *WarehouseAssignmentCreate) SetStatus(w warehouseassignment.Status) *WarehouseAssignmentCreate {
	wac.mutation.SetStatus(w)
	return wac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wac *WarehouseAssignmentCreate) SetNillableStatus(w *warehouseassignment.Status) *WarehouseAssignmentCreate {
	if w != nil {
		wac.SetStatus(*w)
	}
	return wac
}

// SetNote sets the "note" field.
func (wac *WarehouseAssignmentCreate) SetNote(s string) *WarehouseAssignmentCreate {
	wac.mutation.SetNote(s)
	return wac
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (wac *WarehouseAssignmentCreate) SetNillableNote(s *string) *WarehouseAssignmentCreate {
	if s != nil {
		wac.SetNote(*s)
	}
	return wac
}

// SetID sets the "id" field.
func (wac *WarehouseAssignmentCreate) SetID(u uuid.UUID) *WarehouseAssignmentCreate {
	wac.mutation.SetID(u)
	return wac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wac *WarehouseAssignmentCreate) SetNillableID(u *uuid.UUID) *WarehouseAssignmentCreate {
	if u != nil {
		wac.SetID(*u)
	}
	return wac
}

// SetOrder sets the "order" edge to the Order entity.
func (wac *WarehouseAssignmentCreate) SetOrder(o *Order) *WarehouseAssignmentCreate {
	return wac.SetOrderID(o.ID)
}

// SetWorkUnit sets the "work_unit" edge to the WorkUnitInfo entity.
func (wac *WarehouseAssignmentCreate) SetWorkUnit(w *WorkUnitInfo) *WarehouseAssignmentCreate {
	return wac.SetWorkUnitID(w.ID)
}

// SetStaff sets the "staff" edge to the Person entity.
func (wac *WarehouseAssignmentCreate) SetStaff(p *Person) *WarehouseAssignmentCreate {
	return wac.SetStaffID(p.ID)
}

// Mutation returns the WarehouseAssignmentMutation object of the builder.
func (wac *WarehouseAssignmentCreate) Mutation() *WarehouseAssignmentMutation {
	return wac.mutation
}

// Save creates the WarehouseAssignment in the database.
func (wac *WarehouseAssignmentCreate) Save(ctx context.Context) (*WarehouseAssignment, error) {
	wac.defaults()
	return withHooks(ctx, wac.sqlSave, wac.mutation, wac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wac *WarehouseAssignmentCreate) SaveX(ctx context.Context) *WarehouseAssignment {
	v, err := wac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wac *WarehouseAssignmentCreate) Exec(ctx context.Context) error {
	_, err := wac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wac *WarehouseAssignmentCreate) ExecX(ctx context.Context) {
	if err := wac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wac *WarehouseAssignmentCreate) defaults() {
	if _, ok := wac.mutation.CreatedAt(); !ok {
		v := warehouseassignment.DefaultCreatedAt()
		wac.mutation.SetCreatedAt(v)
	}
	if _, ok := wac.mutation.UpdatedAt(); !ok {
		v := warehouseassignment.DefaultUpdatedAt()
		wac.mutation.SetUpdatedAt(v)
	}
	if _, ok := wac.mutation.Status(); !ok {
		v := warehouseassignment.DefaultStatus
		wac.mutation.SetStatus(v)
	}
	if _, ok := wac.mutation.ID(); !ok {
		v := warehouseassignment.DefaultID()
		wac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wac *WarehouseAssignmentCreate) check() error {
	if _, ok := wac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WarehouseAssignment.created_at"`)}
	}
	if _, ok := wac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WarehouseAssignment.updated_at"`)}
	}
	if _, ok := wac.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "WarehouseAssignment.order_id"`)}
	}
	if _, ok := wac.mutation.WorkUnitID(); !ok {
		return &ValidationError{Name: "work_unit_id", err: errors.New(`ent: missing required field "WarehouseAssignment.work_unit_id"`)}
	}
	if _, ok := wac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "WarehouseAssignment.status"`)}
	}
	if v, ok := wac.mutation.Status(); ok {
		if err := warehouseassignment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "WarehouseAssignment.status": %w`, err)}
		}
	}
	if _, ok := wac.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required edge "WarehouseAssignment.order"`)}
	}
	if _, ok := wac.mutation.WorkUnitID(); !ok {
		return &ValidationError{Name: "work_unit", err: errors.New(`ent: missing required edge "WarehouseAssignment.work_unit"`)}
	}
	return nil
}

func (wac *WarehouseAssignmentCreate) sqlSave(ctx context.Context) (*WarehouseAssignment, error) {
	if err := wac.check(); err != nil {
		return nil, err
	}
	_node, _spec := wac.createSpec()
	if err := sqlgraph.CreateNode(ctx, wac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	wac.mutation.id = &_node.ID
	wac.mutation.done = true
	return _node, nil
}

func (wac *WarehouseAssignmentCreate) createSpec() (*WarehouseAssignment, *sqlgraph.CreateSpec) {
	var (
		_node = &WarehouseAssignment{config: wac.config}
		_spec = sqlgraph.NewCreateSpec(warehouseassignment.Table, sqlgraph.NewFieldSpec(warehouseassignment.FieldID, field.TypeUUID))
	)
	if id, ok := wac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wac.mutation.CreatedAt(); ok {
		_spec.SetField(warehouseassignment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := wac.mutation.UpdatedAt(); ok {
		_spec.SetField(warehouseassignment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := wac.mutation.Status(); ok {
		_spec.SetField(warehouseassignment.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if value, ok := wac.mutation.Note(); ok {
		_spec.SetField(warehouseassignment.FieldNote, field.TypeString, value)
		_node.Note = &value
	}
	if nodes := wac.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   warehouseassignment.OrderTable,
			Columns: []string{warehouseassignment.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wac.mutation.WorkUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   warehouseassignment.WorkUnitTable,
			Columns: []string{warehouseassignment.WorkUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workunitinfo.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.WorkUnitID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wac.mutation.StaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   warehouseassignment.StaffTable,
			Columns: []string{warehouseassignment.StaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.StaffID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WarehouseAssignmentCreateBulk is the builder for creating many WarehouseAssignment entities in bulk.
type WarehouseAssignmentCreateBulk struct {
	config
	err      error
	builders []*WarehouseAssignmentCreate
}

// Save creates the WarehouseAssignment entities in the database.
func (wacb *WarehouseAssignmentCreateBulk) Save(ctx context.Context) ([]*WarehouseAssignment, error) {
	if wacb.err != nil {
		return nil, wacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wacb.builders))
	nodes := make([]*WarehouseAssignment, len(wacb.builders))
	mutators := make([]Mutator, len(wacb.builders))
	for i := range wacb.builders {
		func(i int, root context.Context) {
			builder := wacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WarehouseAssignmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wacb *WarehouseAssignmentCreateBulk) SaveX(ctx context.Context) []*WarehouseAssignment {
	v, err := wacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wacb *WarehouseAssignmentCreateBulk) Exec(ctx context.Context) error {
	_, err := wacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wacb *WarehouseAssignmentCreateBulk) ExecX(ctx context.Context) {
	if err := wacb.Exec(ctx); err != nil {
		panic(err)
	}
}
