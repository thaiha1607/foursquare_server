// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/invoicetype"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// InvoiceTypeUpdate is the builder for updating InvoiceType entities.
type InvoiceTypeUpdate struct {
	config
	hooks    []Hook
	mutation *InvoiceTypeMutation
}

// Where appends a list predicates to the InvoiceTypeUpdate builder.
func (itu *InvoiceTypeUpdate) Where(ps ...predicate.InvoiceType) *InvoiceTypeUpdate {
	itu.mutation.Where(ps...)
	return itu
}

// SetUpdatedAt sets the "updated_at" field.
func (itu *InvoiceTypeUpdate) SetUpdatedAt(t time.Time) *InvoiceTypeUpdate {
	itu.mutation.SetUpdatedAt(t)
	return itu
}

// SetInvoiceType sets the "invoice_type" field.
func (itu *InvoiceTypeUpdate) SetInvoiceType(s string) *InvoiceTypeUpdate {
	itu.mutation.SetInvoiceType(s)
	return itu
}

// SetNillableInvoiceType sets the "invoice_type" field if the given value is not nil.
func (itu *InvoiceTypeUpdate) SetNillableInvoiceType(s *string) *InvoiceTypeUpdate {
	if s != nil {
		itu.SetInvoiceType(*s)
	}
	return itu
}

// Mutation returns the InvoiceTypeMutation object of the builder.
func (itu *InvoiceTypeUpdate) Mutation() *InvoiceTypeMutation {
	return itu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (itu *InvoiceTypeUpdate) Save(ctx context.Context) (int, error) {
	itu.defaults()
	return withHooks(ctx, itu.sqlSave, itu.mutation, itu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (itu *InvoiceTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := itu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (itu *InvoiceTypeUpdate) Exec(ctx context.Context) error {
	_, err := itu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itu *InvoiceTypeUpdate) ExecX(ctx context.Context) {
	if err := itu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (itu *InvoiceTypeUpdate) defaults() {
	if _, ok := itu.mutation.UpdatedAt(); !ok {
		v := invoicetype.UpdateDefaultUpdatedAt()
		itu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (itu *InvoiceTypeUpdate) check() error {
	if v, ok := itu.mutation.InvoiceType(); ok {
		if err := invoicetype.InvoiceTypeValidator(v); err != nil {
			return &ValidationError{Name: "invoice_type", err: fmt.Errorf(`ent: validator failed for field "InvoiceType.invoice_type": %w`, err)}
		}
	}
	return nil
}

func (itu *InvoiceTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := itu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoicetype.Table, invoicetype.Columns, sqlgraph.NewFieldSpec(invoicetype.FieldID, field.TypeInt))
	if ps := itu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := itu.mutation.UpdatedAt(); ok {
		_spec.SetField(invoicetype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := itu.mutation.InvoiceType(); ok {
		_spec.SetField(invoicetype.FieldInvoiceType, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, itu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoicetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	itu.mutation.done = true
	return n, nil
}

// InvoiceTypeUpdateOne is the builder for updating a single InvoiceType entity.
type InvoiceTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InvoiceTypeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ituo *InvoiceTypeUpdateOne) SetUpdatedAt(t time.Time) *InvoiceTypeUpdateOne {
	ituo.mutation.SetUpdatedAt(t)
	return ituo
}

// SetInvoiceType sets the "invoice_type" field.
func (ituo *InvoiceTypeUpdateOne) SetInvoiceType(s string) *InvoiceTypeUpdateOne {
	ituo.mutation.SetInvoiceType(s)
	return ituo
}

// SetNillableInvoiceType sets the "invoice_type" field if the given value is not nil.
func (ituo *InvoiceTypeUpdateOne) SetNillableInvoiceType(s *string) *InvoiceTypeUpdateOne {
	if s != nil {
		ituo.SetInvoiceType(*s)
	}
	return ituo
}

// Mutation returns the InvoiceTypeMutation object of the builder.
func (ituo *InvoiceTypeUpdateOne) Mutation() *InvoiceTypeMutation {
	return ituo.mutation
}

// Where appends a list predicates to the InvoiceTypeUpdate builder.
func (ituo *InvoiceTypeUpdateOne) Where(ps ...predicate.InvoiceType) *InvoiceTypeUpdateOne {
	ituo.mutation.Where(ps...)
	return ituo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ituo *InvoiceTypeUpdateOne) Select(field string, fields ...string) *InvoiceTypeUpdateOne {
	ituo.fields = append([]string{field}, fields...)
	return ituo
}

// Save executes the query and returns the updated InvoiceType entity.
func (ituo *InvoiceTypeUpdateOne) Save(ctx context.Context) (*InvoiceType, error) {
	ituo.defaults()
	return withHooks(ctx, ituo.sqlSave, ituo.mutation, ituo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ituo *InvoiceTypeUpdateOne) SaveX(ctx context.Context) *InvoiceType {
	node, err := ituo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ituo *InvoiceTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ituo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ituo *InvoiceTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ituo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ituo *InvoiceTypeUpdateOne) defaults() {
	if _, ok := ituo.mutation.UpdatedAt(); !ok {
		v := invoicetype.UpdateDefaultUpdatedAt()
		ituo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ituo *InvoiceTypeUpdateOne) check() error {
	if v, ok := ituo.mutation.InvoiceType(); ok {
		if err := invoicetype.InvoiceTypeValidator(v); err != nil {
			return &ValidationError{Name: "invoice_type", err: fmt.Errorf(`ent: validator failed for field "InvoiceType.invoice_type": %w`, err)}
		}
	}
	return nil
}

func (ituo *InvoiceTypeUpdateOne) sqlSave(ctx context.Context) (_node *InvoiceType, err error) {
	if err := ituo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoicetype.Table, invoicetype.Columns, sqlgraph.NewFieldSpec(invoicetype.FieldID, field.TypeInt))
	id, ok := ituo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "InvoiceType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ituo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoicetype.FieldID)
		for _, f := range fields {
			if !invoicetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != invoicetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ituo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ituo.mutation.UpdatedAt(); ok {
		_spec.SetField(invoicetype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ituo.mutation.InvoiceType(); ok {
		_spec.SetField(invoicetype.FieldInvoiceType, field.TypeString, value)
	}
	_node = &InvoiceType{config: ituo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ituo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoicetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ituo.mutation.done = true
	return _node, nil
}
