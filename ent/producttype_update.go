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
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/producttype"
)

// ProductTypeUpdate is the builder for updating ProductType entities.
type ProductTypeUpdate struct {
	config
	hooks    []Hook
	mutation *ProductTypeMutation
}

// Where appends a list predicates to the ProductTypeUpdate builder.
func (ptu *ProductTypeUpdate) Where(ps ...predicate.ProductType) *ProductTypeUpdate {
	ptu.mutation.Where(ps...)
	return ptu
}

// SetUpdatedAt sets the "updated_at" field.
func (ptu *ProductTypeUpdate) SetUpdatedAt(t time.Time) *ProductTypeUpdate {
	ptu.mutation.SetUpdatedAt(t)
	return ptu
}

// SetProductType sets the "product_type" field.
func (ptu *ProductTypeUpdate) SetProductType(s string) *ProductTypeUpdate {
	ptu.mutation.SetProductType(s)
	return ptu
}

// SetNillableProductType sets the "product_type" field if the given value is not nil.
func (ptu *ProductTypeUpdate) SetNillableProductType(s *string) *ProductTypeUpdate {
	if s != nil {
		ptu.SetProductType(*s)
	}
	return ptu
}

// Mutation returns the ProductTypeMutation object of the builder.
func (ptu *ProductTypeUpdate) Mutation() *ProductTypeMutation {
	return ptu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptu *ProductTypeUpdate) Save(ctx context.Context) (int, error) {
	ptu.defaults()
	return withHooks(ctx, ptu.sqlSave, ptu.mutation, ptu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptu *ProductTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptu *ProductTypeUpdate) Exec(ctx context.Context) error {
	_, err := ptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptu *ProductTypeUpdate) ExecX(ctx context.Context) {
	if err := ptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptu *ProductTypeUpdate) defaults() {
	if _, ok := ptu.mutation.UpdatedAt(); !ok {
		v := producttype.UpdateDefaultUpdatedAt()
		ptu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptu *ProductTypeUpdate) check() error {
	if v, ok := ptu.mutation.ProductType(); ok {
		if err := producttype.ProductTypeValidator(v); err != nil {
			return &ValidationError{Name: "product_type", err: fmt.Errorf(`ent: validator failed for field "ProductType.product_type": %w`, err)}
		}
	}
	return nil
}

func (ptu *ProductTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ptu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(producttype.Table, producttype.Columns, sqlgraph.NewFieldSpec(producttype.FieldID, field.TypeInt))
	if ps := ptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptu.mutation.UpdatedAt(); ok {
		_spec.SetField(producttype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ptu.mutation.ProductType(); ok {
		_spec.SetField(producttype.FieldProductType, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{producttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ptu.mutation.done = true
	return n, nil
}

// ProductTypeUpdateOne is the builder for updating a single ProductType entity.
type ProductTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductTypeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ptuo *ProductTypeUpdateOne) SetUpdatedAt(t time.Time) *ProductTypeUpdateOne {
	ptuo.mutation.SetUpdatedAt(t)
	return ptuo
}

// SetProductType sets the "product_type" field.
func (ptuo *ProductTypeUpdateOne) SetProductType(s string) *ProductTypeUpdateOne {
	ptuo.mutation.SetProductType(s)
	return ptuo
}

// SetNillableProductType sets the "product_type" field if the given value is not nil.
func (ptuo *ProductTypeUpdateOne) SetNillableProductType(s *string) *ProductTypeUpdateOne {
	if s != nil {
		ptuo.SetProductType(*s)
	}
	return ptuo
}

// Mutation returns the ProductTypeMutation object of the builder.
func (ptuo *ProductTypeUpdateOne) Mutation() *ProductTypeMutation {
	return ptuo.mutation
}

// Where appends a list predicates to the ProductTypeUpdate builder.
func (ptuo *ProductTypeUpdateOne) Where(ps ...predicate.ProductType) *ProductTypeUpdateOne {
	ptuo.mutation.Where(ps...)
	return ptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptuo *ProductTypeUpdateOne) Select(field string, fields ...string) *ProductTypeUpdateOne {
	ptuo.fields = append([]string{field}, fields...)
	return ptuo
}

// Save executes the query and returns the updated ProductType entity.
func (ptuo *ProductTypeUpdateOne) Save(ctx context.Context) (*ProductType, error) {
	ptuo.defaults()
	return withHooks(ctx, ptuo.sqlSave, ptuo.mutation, ptuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptuo *ProductTypeUpdateOne) SaveX(ctx context.Context) *ProductType {
	node, err := ptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptuo *ProductTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptuo *ProductTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptuo *ProductTypeUpdateOne) defaults() {
	if _, ok := ptuo.mutation.UpdatedAt(); !ok {
		v := producttype.UpdateDefaultUpdatedAt()
		ptuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptuo *ProductTypeUpdateOne) check() error {
	if v, ok := ptuo.mutation.ProductType(); ok {
		if err := producttype.ProductTypeValidator(v); err != nil {
			return &ValidationError{Name: "product_type", err: fmt.Errorf(`ent: validator failed for field "ProductType.product_type": %w`, err)}
		}
	}
	return nil
}

func (ptuo *ProductTypeUpdateOne) sqlSave(ctx context.Context) (_node *ProductType, err error) {
	if err := ptuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(producttype.Table, producttype.Columns, sqlgraph.NewFieldSpec(producttype.FieldID, field.TypeInt))
	id, ok := ptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProductType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, producttype.FieldID)
		for _, f := range fields {
			if !producttype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != producttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptuo.mutation.UpdatedAt(); ok {
		_spec.SetField(producttype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ptuo.mutation.ProductType(); ok {
		_spec.SetField(producttype.FieldProductType, field.TypeString, value)
	}
	_node = &ProductType{config: ptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{producttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ptuo.mutation.done = true
	return _node, nil
}
