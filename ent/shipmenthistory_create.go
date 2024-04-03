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
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/shipment"
	"github.com/thaiha1607/foursquare_server/ent/shipmenthistory"
	"github.com/thaiha1607/foursquare_server/ent/shipmentstatuscode"
)

// ShipmentHistoryCreate is the builder for creating a ShipmentHistory entity.
type ShipmentHistoryCreate struct {
	config
	mutation *ShipmentHistoryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (shc *ShipmentHistoryCreate) SetCreatedAt(t time.Time) *ShipmentHistoryCreate {
	shc.mutation.SetCreatedAt(t)
	return shc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (shc *ShipmentHistoryCreate) SetNillableCreatedAt(t *time.Time) *ShipmentHistoryCreate {
	if t != nil {
		shc.SetCreatedAt(*t)
	}
	return shc
}

// SetUpdatedAt sets the "updated_at" field.
func (shc *ShipmentHistoryCreate) SetUpdatedAt(t time.Time) *ShipmentHistoryCreate {
	shc.mutation.SetUpdatedAt(t)
	return shc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (shc *ShipmentHistoryCreate) SetNillableUpdatedAt(t *time.Time) *ShipmentHistoryCreate {
	if t != nil {
		shc.SetUpdatedAt(*t)
	}
	return shc
}

// SetShipmentID sets the "shipment_id" field.
func (shc *ShipmentHistoryCreate) SetShipmentID(s string) *ShipmentHistoryCreate {
	shc.mutation.SetShipmentID(s)
	return shc
}

// SetPersonID sets the "person_id" field.
func (shc *ShipmentHistoryCreate) SetPersonID(u uuid.UUID) *ShipmentHistoryCreate {
	shc.mutation.SetPersonID(u)
	return shc
}

// SetOldStatusCode sets the "old_status_code" field.
func (shc *ShipmentHistoryCreate) SetOldStatusCode(i int) *ShipmentHistoryCreate {
	shc.mutation.SetOldStatusCode(i)
	return shc
}

// SetNillableOldStatusCode sets the "old_status_code" field if the given value is not nil.
func (shc *ShipmentHistoryCreate) SetNillableOldStatusCode(i *int) *ShipmentHistoryCreate {
	if i != nil {
		shc.SetOldStatusCode(*i)
	}
	return shc
}

// SetNewStatusCode sets the "new_status_code" field.
func (shc *ShipmentHistoryCreate) SetNewStatusCode(i int) *ShipmentHistoryCreate {
	shc.mutation.SetNewStatusCode(i)
	return shc
}

// SetNillableNewStatusCode sets the "new_status_code" field if the given value is not nil.
func (shc *ShipmentHistoryCreate) SetNillableNewStatusCode(i *int) *ShipmentHistoryCreate {
	if i != nil {
		shc.SetNewStatusCode(*i)
	}
	return shc
}

// SetDescription sets the "description" field.
func (shc *ShipmentHistoryCreate) SetDescription(s string) *ShipmentHistoryCreate {
	shc.mutation.SetDescription(s)
	return shc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (shc *ShipmentHistoryCreate) SetNillableDescription(s *string) *ShipmentHistoryCreate {
	if s != nil {
		shc.SetDescription(*s)
	}
	return shc
}

// SetID sets the "id" field.
func (shc *ShipmentHistoryCreate) SetID(u uuid.UUID) *ShipmentHistoryCreate {
	shc.mutation.SetID(u)
	return shc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (shc *ShipmentHistoryCreate) SetNillableID(u *uuid.UUID) *ShipmentHistoryCreate {
	if u != nil {
		shc.SetID(*u)
	}
	return shc
}

// SetShipment sets the "shipment" edge to the Shipment entity.
func (shc *ShipmentHistoryCreate) SetShipment(s *Shipment) *ShipmentHistoryCreate {
	return shc.SetShipmentID(s.ID)
}

// SetPerson sets the "person" edge to the Person entity.
func (shc *ShipmentHistoryCreate) SetPerson(p *Person) *ShipmentHistoryCreate {
	return shc.SetPersonID(p.ID)
}

// SetOldStatusID sets the "old_status" edge to the ShipmentStatusCode entity by ID.
func (shc *ShipmentHistoryCreate) SetOldStatusID(id int) *ShipmentHistoryCreate {
	shc.mutation.SetOldStatusID(id)
	return shc
}

// SetNillableOldStatusID sets the "old_status" edge to the ShipmentStatusCode entity by ID if the given value is not nil.
func (shc *ShipmentHistoryCreate) SetNillableOldStatusID(id *int) *ShipmentHistoryCreate {
	if id != nil {
		shc = shc.SetOldStatusID(*id)
	}
	return shc
}

// SetOldStatus sets the "old_status" edge to the ShipmentStatusCode entity.
func (shc *ShipmentHistoryCreate) SetOldStatus(s *ShipmentStatusCode) *ShipmentHistoryCreate {
	return shc.SetOldStatusID(s.ID)
}

// SetNewStatusID sets the "new_status" edge to the ShipmentStatusCode entity by ID.
func (shc *ShipmentHistoryCreate) SetNewStatusID(id int) *ShipmentHistoryCreate {
	shc.mutation.SetNewStatusID(id)
	return shc
}

// SetNillableNewStatusID sets the "new_status" edge to the ShipmentStatusCode entity by ID if the given value is not nil.
func (shc *ShipmentHistoryCreate) SetNillableNewStatusID(id *int) *ShipmentHistoryCreate {
	if id != nil {
		shc = shc.SetNewStatusID(*id)
	}
	return shc
}

// SetNewStatus sets the "new_status" edge to the ShipmentStatusCode entity.
func (shc *ShipmentHistoryCreate) SetNewStatus(s *ShipmentStatusCode) *ShipmentHistoryCreate {
	return shc.SetNewStatusID(s.ID)
}

// Mutation returns the ShipmentHistoryMutation object of the builder.
func (shc *ShipmentHistoryCreate) Mutation() *ShipmentHistoryMutation {
	return shc.mutation
}

// Save creates the ShipmentHistory in the database.
func (shc *ShipmentHistoryCreate) Save(ctx context.Context) (*ShipmentHistory, error) {
	shc.defaults()
	return withHooks(ctx, shc.sqlSave, shc.mutation, shc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (shc *ShipmentHistoryCreate) SaveX(ctx context.Context) *ShipmentHistory {
	v, err := shc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (shc *ShipmentHistoryCreate) Exec(ctx context.Context) error {
	_, err := shc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (shc *ShipmentHistoryCreate) ExecX(ctx context.Context) {
	if err := shc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (shc *ShipmentHistoryCreate) defaults() {
	if _, ok := shc.mutation.CreatedAt(); !ok {
		v := shipmenthistory.DefaultCreatedAt()
		shc.mutation.SetCreatedAt(v)
	}
	if _, ok := shc.mutation.UpdatedAt(); !ok {
		v := shipmenthistory.DefaultUpdatedAt()
		shc.mutation.SetUpdatedAt(v)
	}
	if _, ok := shc.mutation.ID(); !ok {
		v := shipmenthistory.DefaultID()
		shc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (shc *ShipmentHistoryCreate) check() error {
	if _, ok := shc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ShipmentHistory.created_at"`)}
	}
	if _, ok := shc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ShipmentHistory.updated_at"`)}
	}
	if _, ok := shc.mutation.ShipmentID(); !ok {
		return &ValidationError{Name: "shipment_id", err: errors.New(`ent: missing required field "ShipmentHistory.shipment_id"`)}
	}
	if v, ok := shc.mutation.ShipmentID(); ok {
		if err := shipmenthistory.ShipmentIDValidator(v); err != nil {
			return &ValidationError{Name: "shipment_id", err: fmt.Errorf(`ent: validator failed for field "ShipmentHistory.shipment_id": %w`, err)}
		}
	}
	if _, ok := shc.mutation.PersonID(); !ok {
		return &ValidationError{Name: "person_id", err: errors.New(`ent: missing required field "ShipmentHistory.person_id"`)}
	}
	if _, ok := shc.mutation.ShipmentID(); !ok {
		return &ValidationError{Name: "shipment", err: errors.New(`ent: missing required edge "ShipmentHistory.shipment"`)}
	}
	if _, ok := shc.mutation.PersonID(); !ok {
		return &ValidationError{Name: "person", err: errors.New(`ent: missing required edge "ShipmentHistory.person"`)}
	}
	return nil
}

func (shc *ShipmentHistoryCreate) sqlSave(ctx context.Context) (*ShipmentHistory, error) {
	if err := shc.check(); err != nil {
		return nil, err
	}
	_node, _spec := shc.createSpec()
	if err := sqlgraph.CreateNode(ctx, shc.driver, _spec); err != nil {
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
	shc.mutation.id = &_node.ID
	shc.mutation.done = true
	return _node, nil
}

func (shc *ShipmentHistoryCreate) createSpec() (*ShipmentHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &ShipmentHistory{config: shc.config}
		_spec = sqlgraph.NewCreateSpec(shipmenthistory.Table, sqlgraph.NewFieldSpec(shipmenthistory.FieldID, field.TypeUUID))
	)
	if id, ok := shc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := shc.mutation.CreatedAt(); ok {
		_spec.SetField(shipmenthistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := shc.mutation.UpdatedAt(); ok {
		_spec.SetField(shipmenthistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := shc.mutation.Description(); ok {
		_spec.SetField(shipmenthistory.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if nodes := shc.mutation.ShipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.ShipmentTable,
			Columns: []string{shipmenthistory.ShipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ShipmentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := shc.mutation.PersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.PersonTable,
			Columns: []string{shipmenthistory.PersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PersonID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := shc.mutation.OldStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.OldStatusTable,
			Columns: []string{shipmenthistory.OldStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipmentstatuscode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OldStatusCode = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := shc.mutation.NewStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.NewStatusTable,
			Columns: []string{shipmenthistory.NewStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipmentstatuscode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NewStatusCode = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ShipmentHistoryCreateBulk is the builder for creating many ShipmentHistory entities in bulk.
type ShipmentHistoryCreateBulk struct {
	config
	err      error
	builders []*ShipmentHistoryCreate
}

// Save creates the ShipmentHistory entities in the database.
func (shcb *ShipmentHistoryCreateBulk) Save(ctx context.Context) ([]*ShipmentHistory, error) {
	if shcb.err != nil {
		return nil, shcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(shcb.builders))
	nodes := make([]*ShipmentHistory, len(shcb.builders))
	mutators := make([]Mutator, len(shcb.builders))
	for i := range shcb.builders {
		func(i int, root context.Context) {
			builder := shcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShipmentHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, shcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, shcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, shcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (shcb *ShipmentHistoryCreateBulk) SaveX(ctx context.Context) []*ShipmentHistory {
	v, err := shcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (shcb *ShipmentHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := shcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (shcb *ShipmentHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := shcb.Exec(ctx); err != nil {
		panic(err)
	}
}
