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
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/order"
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/shipment"
	"github.com/thaiha1607/foursquare_server/ent/shipmentstatuscode"
)

// ShipmentCreate is the builder for creating a Shipment entity.
type ShipmentCreate struct {
	config
	mutation *ShipmentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *ShipmentCreate) SetCreatedAt(t time.Time) *ShipmentCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableCreatedAt(t *time.Time) *ShipmentCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *ShipmentCreate) SetUpdatedAt(t time.Time) *ShipmentCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableUpdatedAt(t *time.Time) *ShipmentCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetOrderID sets the "order_id" field.
func (sc *ShipmentCreate) SetOrderID(u uuid.UUID) *ShipmentCreate {
	sc.mutation.SetOrderID(u)
	return sc
}

// SetInvoiceID sets the "invoice_id" field.
func (sc *ShipmentCreate) SetInvoiceID(u uuid.UUID) *ShipmentCreate {
	sc.mutation.SetInvoiceID(u)
	return sc
}

// SetStaffID sets the "staff_id" field.
func (sc *ShipmentCreate) SetStaffID(u uuid.UUID) *ShipmentCreate {
	sc.mutation.SetStaffID(u)
	return sc
}

// SetShipmentDate sets the "shipment_date" field.
func (sc *ShipmentCreate) SetShipmentDate(t time.Time) *ShipmentCreate {
	sc.mutation.SetShipmentDate(t)
	return sc
}

// SetNote sets the "note" field.
func (sc *ShipmentCreate) SetNote(s string) *ShipmentCreate {
	sc.mutation.SetNote(s)
	return sc
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableNote(s *string) *ShipmentCreate {
	if s != nil {
		sc.SetNote(*s)
	}
	return sc
}

// SetStatusCode sets the "status_code" field.
func (sc *ShipmentCreate) SetStatusCode(i int) *ShipmentCreate {
	sc.mutation.SetStatusCode(i)
	return sc
}

// SetNillableStatusCode sets the "status_code" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableStatusCode(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetStatusCode(*i)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *ShipmentCreate) SetID(s string) *ShipmentCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetOrder sets the "order" edge to the Order entity.
func (sc *ShipmentCreate) SetOrder(o *Order) *ShipmentCreate {
	return sc.SetOrderID(o.ID)
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (sc *ShipmentCreate) SetInvoice(i *Invoice) *ShipmentCreate {
	return sc.SetInvoiceID(i.ID)
}

// SetStaff sets the "staff" edge to the Person entity.
func (sc *ShipmentCreate) SetStaff(p *Person) *ShipmentCreate {
	return sc.SetStaffID(p.ID)
}

// SetShipmentStatusID sets the "shipment_status" edge to the ShipmentStatusCode entity by ID.
func (sc *ShipmentCreate) SetShipmentStatusID(id int) *ShipmentCreate {
	sc.mutation.SetShipmentStatusID(id)
	return sc
}

// SetShipmentStatus sets the "shipment_status" edge to the ShipmentStatusCode entity.
func (sc *ShipmentCreate) SetShipmentStatus(s *ShipmentStatusCode) *ShipmentCreate {
	return sc.SetShipmentStatusID(s.ID)
}

// Mutation returns the ShipmentMutation object of the builder.
func (sc *ShipmentCreate) Mutation() *ShipmentMutation {
	return sc.mutation
}

// Save creates the Shipment in the database.
func (sc *ShipmentCreate) Save(ctx context.Context) (*Shipment, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShipmentCreate) SaveX(ctx context.Context) *Shipment {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ShipmentCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ShipmentCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ShipmentCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if shipment.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized shipment.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := shipment.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if shipment.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized shipment.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := shipment.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.StatusCode(); !ok {
		v := shipment.DefaultStatusCode
		sc.mutation.SetStatusCode(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *ShipmentCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Shipment.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Shipment.updated_at"`)}
	}
	if _, ok := sc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "Shipment.order_id"`)}
	}
	if _, ok := sc.mutation.InvoiceID(); !ok {
		return &ValidationError{Name: "invoice_id", err: errors.New(`ent: missing required field "Shipment.invoice_id"`)}
	}
	if _, ok := sc.mutation.StaffID(); !ok {
		return &ValidationError{Name: "staff_id", err: errors.New(`ent: missing required field "Shipment.staff_id"`)}
	}
	if _, ok := sc.mutation.ShipmentDate(); !ok {
		return &ValidationError{Name: "shipment_date", err: errors.New(`ent: missing required field "Shipment.shipment_date"`)}
	}
	if _, ok := sc.mutation.StatusCode(); !ok {
		return &ValidationError{Name: "status_code", err: errors.New(`ent: missing required field "Shipment.status_code"`)}
	}
	if v, ok := sc.mutation.ID(); ok {
		if err := shipment.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Shipment.id": %w`, err)}
		}
	}
	if _, ok := sc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required edge "Shipment.order"`)}
	}
	if _, ok := sc.mutation.InvoiceID(); !ok {
		return &ValidationError{Name: "invoice", err: errors.New(`ent: missing required edge "Shipment.invoice"`)}
	}
	if _, ok := sc.mutation.StaffID(); !ok {
		return &ValidationError{Name: "staff", err: errors.New(`ent: missing required edge "Shipment.staff"`)}
	}
	if _, ok := sc.mutation.ShipmentStatusID(); !ok {
		return &ValidationError{Name: "shipment_status", err: errors.New(`ent: missing required edge "Shipment.shipment_status"`)}
	}
	return nil
}

func (sc *ShipmentCreate) sqlSave(ctx context.Context) (*Shipment, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Shipment.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ShipmentCreate) createSpec() (*Shipment, *sqlgraph.CreateSpec) {
	var (
		_node = &Shipment{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(shipment.Table, sqlgraph.NewFieldSpec(shipment.FieldID, field.TypeString))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(shipment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(shipment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := sc.mutation.ShipmentDate(); ok {
		_spec.SetField(shipment.FieldShipmentDate, field.TypeTime, value)
		_node.ShipmentDate = &value
	}
	if value, ok := sc.mutation.Note(); ok {
		_spec.SetField(shipment.FieldNote, field.TypeString, value)
		_node.Note = &value
	}
	if nodes := sc.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipment.OrderTable,
			Columns: []string{shipment.OrderColumn},
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
	if nodes := sc.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipment.InvoiceTable,
			Columns: []string{shipment.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.InvoiceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipment.StaffTable,
			Columns: []string{shipment.StaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.StaffID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ShipmentStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipment.ShipmentStatusTable,
			Columns: []string{shipment.ShipmentStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipmentstatuscode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.StatusCode = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ShipmentCreateBulk is the builder for creating many Shipment entities in bulk.
type ShipmentCreateBulk struct {
	config
	err      error
	builders []*ShipmentCreate
}

// Save creates the Shipment entities in the database.
func (scb *ShipmentCreateBulk) Save(ctx context.Context) ([]*Shipment, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Shipment, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShipmentMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ShipmentCreateBulk) SaveX(ctx context.Context) []*Shipment {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ShipmentCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ShipmentCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}