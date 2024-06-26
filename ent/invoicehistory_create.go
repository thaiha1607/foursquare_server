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
	"github.com/thaiha1607/foursquare_server/ent/invoicehistory"
	"github.com/thaiha1607/foursquare_server/ent/invoicestatuscode"
	"github.com/thaiha1607/foursquare_server/ent/person"
)

// InvoiceHistoryCreate is the builder for creating a InvoiceHistory entity.
type InvoiceHistoryCreate struct {
	config
	mutation *InvoiceHistoryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ihc *InvoiceHistoryCreate) SetCreatedAt(t time.Time) *InvoiceHistoryCreate {
	ihc.mutation.SetCreatedAt(t)
	return ihc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ihc *InvoiceHistoryCreate) SetNillableCreatedAt(t *time.Time) *InvoiceHistoryCreate {
	if t != nil {
		ihc.SetCreatedAt(*t)
	}
	return ihc
}

// SetInvoiceID sets the "invoice_id" field.
func (ihc *InvoiceHistoryCreate) SetInvoiceID(u uuid.UUID) *InvoiceHistoryCreate {
	ihc.mutation.SetInvoiceID(u)
	return ihc
}

// SetPersonID sets the "person_id" field.
func (ihc *InvoiceHistoryCreate) SetPersonID(u uuid.UUID) *InvoiceHistoryCreate {
	ihc.mutation.SetPersonID(u)
	return ihc
}

// SetOldStatusCode sets the "old_status_code" field.
func (ihc *InvoiceHistoryCreate) SetOldStatusCode(i int) *InvoiceHistoryCreate {
	ihc.mutation.SetOldStatusCode(i)
	return ihc
}

// SetNillableOldStatusCode sets the "old_status_code" field if the given value is not nil.
func (ihc *InvoiceHistoryCreate) SetNillableOldStatusCode(i *int) *InvoiceHistoryCreate {
	if i != nil {
		ihc.SetOldStatusCode(*i)
	}
	return ihc
}

// SetNewStatusCode sets the "new_status_code" field.
func (ihc *InvoiceHistoryCreate) SetNewStatusCode(i int) *InvoiceHistoryCreate {
	ihc.mutation.SetNewStatusCode(i)
	return ihc
}

// SetNillableNewStatusCode sets the "new_status_code" field if the given value is not nil.
func (ihc *InvoiceHistoryCreate) SetNillableNewStatusCode(i *int) *InvoiceHistoryCreate {
	if i != nil {
		ihc.SetNewStatusCode(*i)
	}
	return ihc
}

// SetDescription sets the "description" field.
func (ihc *InvoiceHistoryCreate) SetDescription(s string) *InvoiceHistoryCreate {
	ihc.mutation.SetDescription(s)
	return ihc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ihc *InvoiceHistoryCreate) SetNillableDescription(s *string) *InvoiceHistoryCreate {
	if s != nil {
		ihc.SetDescription(*s)
	}
	return ihc
}

// SetID sets the "id" field.
func (ihc *InvoiceHistoryCreate) SetID(u uuid.UUID) *InvoiceHistoryCreate {
	ihc.mutation.SetID(u)
	return ihc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ihc *InvoiceHistoryCreate) SetNillableID(u *uuid.UUID) *InvoiceHistoryCreate {
	if u != nil {
		ihc.SetID(*u)
	}
	return ihc
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (ihc *InvoiceHistoryCreate) SetInvoice(i *Invoice) *InvoiceHistoryCreate {
	return ihc.SetInvoiceID(i.ID)
}

// SetPerson sets the "person" edge to the Person entity.
func (ihc *InvoiceHistoryCreate) SetPerson(p *Person) *InvoiceHistoryCreate {
	return ihc.SetPersonID(p.ID)
}

// SetOldStatusID sets the "old_status" edge to the InvoiceStatusCode entity by ID.
func (ihc *InvoiceHistoryCreate) SetOldStatusID(id int) *InvoiceHistoryCreate {
	ihc.mutation.SetOldStatusID(id)
	return ihc
}

// SetNillableOldStatusID sets the "old_status" edge to the InvoiceStatusCode entity by ID if the given value is not nil.
func (ihc *InvoiceHistoryCreate) SetNillableOldStatusID(id *int) *InvoiceHistoryCreate {
	if id != nil {
		ihc = ihc.SetOldStatusID(*id)
	}
	return ihc
}

// SetOldStatus sets the "old_status" edge to the InvoiceStatusCode entity.
func (ihc *InvoiceHistoryCreate) SetOldStatus(i *InvoiceStatusCode) *InvoiceHistoryCreate {
	return ihc.SetOldStatusID(i.ID)
}

// SetNewStatusID sets the "new_status" edge to the InvoiceStatusCode entity by ID.
func (ihc *InvoiceHistoryCreate) SetNewStatusID(id int) *InvoiceHistoryCreate {
	ihc.mutation.SetNewStatusID(id)
	return ihc
}

// SetNillableNewStatusID sets the "new_status" edge to the InvoiceStatusCode entity by ID if the given value is not nil.
func (ihc *InvoiceHistoryCreate) SetNillableNewStatusID(id *int) *InvoiceHistoryCreate {
	if id != nil {
		ihc = ihc.SetNewStatusID(*id)
	}
	return ihc
}

// SetNewStatus sets the "new_status" edge to the InvoiceStatusCode entity.
func (ihc *InvoiceHistoryCreate) SetNewStatus(i *InvoiceStatusCode) *InvoiceHistoryCreate {
	return ihc.SetNewStatusID(i.ID)
}

// Mutation returns the InvoiceHistoryMutation object of the builder.
func (ihc *InvoiceHistoryCreate) Mutation() *InvoiceHistoryMutation {
	return ihc.mutation
}

// Save creates the InvoiceHistory in the database.
func (ihc *InvoiceHistoryCreate) Save(ctx context.Context) (*InvoiceHistory, error) {
	ihc.defaults()
	return withHooks(ctx, ihc.sqlSave, ihc.mutation, ihc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ihc *InvoiceHistoryCreate) SaveX(ctx context.Context) *InvoiceHistory {
	v, err := ihc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ihc *InvoiceHistoryCreate) Exec(ctx context.Context) error {
	_, err := ihc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ihc *InvoiceHistoryCreate) ExecX(ctx context.Context) {
	if err := ihc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ihc *InvoiceHistoryCreate) defaults() {
	if _, ok := ihc.mutation.CreatedAt(); !ok {
		v := invoicehistory.DefaultCreatedAt()
		ihc.mutation.SetCreatedAt(v)
	}
	if _, ok := ihc.mutation.ID(); !ok {
		v := invoicehistory.DefaultID()
		ihc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ihc *InvoiceHistoryCreate) check() error {
	if _, ok := ihc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "InvoiceHistory.created_at"`)}
	}
	if _, ok := ihc.mutation.InvoiceID(); !ok {
		return &ValidationError{Name: "invoice_id", err: errors.New(`ent: missing required field "InvoiceHistory.invoice_id"`)}
	}
	if _, ok := ihc.mutation.PersonID(); !ok {
		return &ValidationError{Name: "person_id", err: errors.New(`ent: missing required field "InvoiceHistory.person_id"`)}
	}
	if _, ok := ihc.mutation.InvoiceID(); !ok {
		return &ValidationError{Name: "invoice", err: errors.New(`ent: missing required edge "InvoiceHistory.invoice"`)}
	}
	if _, ok := ihc.mutation.PersonID(); !ok {
		return &ValidationError{Name: "person", err: errors.New(`ent: missing required edge "InvoiceHistory.person"`)}
	}
	return nil
}

func (ihc *InvoiceHistoryCreate) sqlSave(ctx context.Context) (*InvoiceHistory, error) {
	if err := ihc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ihc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ihc.driver, _spec); err != nil {
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
	ihc.mutation.id = &_node.ID
	ihc.mutation.done = true
	return _node, nil
}

func (ihc *InvoiceHistoryCreate) createSpec() (*InvoiceHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &InvoiceHistory{config: ihc.config}
		_spec = sqlgraph.NewCreateSpec(invoicehistory.Table, sqlgraph.NewFieldSpec(invoicehistory.FieldID, field.TypeUUID))
	)
	if id, ok := ihc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ihc.mutation.CreatedAt(); ok {
		_spec.SetField(invoicehistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := ihc.mutation.Description(); ok {
		_spec.SetField(invoicehistory.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if nodes := ihc.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicehistory.InvoiceTable,
			Columns: []string{invoicehistory.InvoiceColumn},
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
	if nodes := ihc.mutation.PersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicehistory.PersonTable,
			Columns: []string{invoicehistory.PersonColumn},
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
	if nodes := ihc.mutation.OldStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicehistory.OldStatusTable,
			Columns: []string{invoicehistory.OldStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoicestatuscode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OldStatusCode = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ihc.mutation.NewStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicehistory.NewStatusTable,
			Columns: []string{invoicehistory.NewStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoicestatuscode.FieldID, field.TypeInt),
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

// InvoiceHistoryCreateBulk is the builder for creating many InvoiceHistory entities in bulk.
type InvoiceHistoryCreateBulk struct {
	config
	err      error
	builders []*InvoiceHistoryCreate
}

// Save creates the InvoiceHistory entities in the database.
func (ihcb *InvoiceHistoryCreateBulk) Save(ctx context.Context) ([]*InvoiceHistory, error) {
	if ihcb.err != nil {
		return nil, ihcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ihcb.builders))
	nodes := make([]*InvoiceHistory, len(ihcb.builders))
	mutators := make([]Mutator, len(ihcb.builders))
	for i := range ihcb.builders {
		func(i int, root context.Context) {
			builder := ihcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvoiceHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ihcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ihcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ihcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ihcb *InvoiceHistoryCreateBulk) SaveX(ctx context.Context) []*InvoiceHistory {
	v, err := ihcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ihcb *InvoiceHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ihcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ihcb *InvoiceHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ihcb.Exec(ctx); err != nil {
		panic(err)
	}
}
