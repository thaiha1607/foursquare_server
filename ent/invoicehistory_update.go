// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/invoicehistory"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// InvoiceHistoryUpdate is the builder for updating InvoiceHistory entities.
type InvoiceHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *InvoiceHistoryMutation
}

// Where appends a list predicates to the InvoiceHistoryUpdate builder.
func (ihu *InvoiceHistoryUpdate) Where(ps ...predicate.InvoiceHistory) *InvoiceHistoryUpdate {
	ihu.mutation.Where(ps...)
	return ihu
}

// Mutation returns the InvoiceHistoryMutation object of the builder.
func (ihu *InvoiceHistoryUpdate) Mutation() *InvoiceHistoryMutation {
	return ihu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ihu *InvoiceHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ihu.sqlSave, ihu.mutation, ihu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ihu *InvoiceHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ihu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ihu *InvoiceHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ihu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ihu *InvoiceHistoryUpdate) ExecX(ctx context.Context) {
	if err := ihu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ihu *InvoiceHistoryUpdate) check() error {
	if _, ok := ihu.mutation.InvoiceID(); ihu.mutation.InvoiceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceHistory.invoice"`)
	}
	if _, ok := ihu.mutation.PersonID(); ihu.mutation.PersonCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceHistory.person"`)
	}
	return nil
}

func (ihu *InvoiceHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ihu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoicehistory.Table, invoicehistory.Columns, sqlgraph.NewFieldSpec(invoicehistory.FieldID, field.TypeUUID))
	if ps := ihu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ihu.mutation.DescriptionCleared() {
		_spec.ClearField(invoicehistory.FieldDescription, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ihu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoicehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ihu.mutation.done = true
	return n, nil
}

// InvoiceHistoryUpdateOne is the builder for updating a single InvoiceHistory entity.
type InvoiceHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InvoiceHistoryMutation
}

// Mutation returns the InvoiceHistoryMutation object of the builder.
func (ihuo *InvoiceHistoryUpdateOne) Mutation() *InvoiceHistoryMutation {
	return ihuo.mutation
}

// Where appends a list predicates to the InvoiceHistoryUpdate builder.
func (ihuo *InvoiceHistoryUpdateOne) Where(ps ...predicate.InvoiceHistory) *InvoiceHistoryUpdateOne {
	ihuo.mutation.Where(ps...)
	return ihuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ihuo *InvoiceHistoryUpdateOne) Select(field string, fields ...string) *InvoiceHistoryUpdateOne {
	ihuo.fields = append([]string{field}, fields...)
	return ihuo
}

// Save executes the query and returns the updated InvoiceHistory entity.
func (ihuo *InvoiceHistoryUpdateOne) Save(ctx context.Context) (*InvoiceHistory, error) {
	return withHooks(ctx, ihuo.sqlSave, ihuo.mutation, ihuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ihuo *InvoiceHistoryUpdateOne) SaveX(ctx context.Context) *InvoiceHistory {
	node, err := ihuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ihuo *InvoiceHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ihuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ihuo *InvoiceHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ihuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ihuo *InvoiceHistoryUpdateOne) check() error {
	if _, ok := ihuo.mutation.InvoiceID(); ihuo.mutation.InvoiceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceHistory.invoice"`)
	}
	if _, ok := ihuo.mutation.PersonID(); ihuo.mutation.PersonCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceHistory.person"`)
	}
	return nil
}

func (ihuo *InvoiceHistoryUpdateOne) sqlSave(ctx context.Context) (_node *InvoiceHistory, err error) {
	if err := ihuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoicehistory.Table, invoicehistory.Columns, sqlgraph.NewFieldSpec(invoicehistory.FieldID, field.TypeUUID))
	id, ok := ihuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "InvoiceHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ihuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoicehistory.FieldID)
		for _, f := range fields {
			if !invoicehistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != invoicehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ihuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ihuo.mutation.DescriptionCleared() {
		_spec.ClearField(invoicehistory.FieldDescription, field.TypeString)
	}
	_node = &InvoiceHistory{config: ihuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ihuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoicehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ihuo.mutation.done = true
	return _node, nil
}
