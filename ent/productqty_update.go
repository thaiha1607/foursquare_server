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
	"github.com/shopspring/decimal"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/productqty"
)

// ProductQtyUpdate is the builder for updating ProductQty entities.
type ProductQtyUpdate struct {
	config
	hooks    []Hook
	mutation *ProductQtyMutation
}

// Where appends a list predicates to the ProductQtyUpdate builder.
func (pqu *ProductQtyUpdate) Where(ps ...predicate.ProductQty) *ProductQtyUpdate {
	pqu.mutation.Where(ps...)
	return pqu
}

// SetUpdatedAt sets the "updated_at" field.
func (pqu *ProductQtyUpdate) SetUpdatedAt(t time.Time) *ProductQtyUpdate {
	pqu.mutation.SetUpdatedAt(t)
	return pqu
}

// SetPricePerUnit sets the "price_per_unit" field.
func (pqu *ProductQtyUpdate) SetPricePerUnit(d decimal.Decimal) *ProductQtyUpdate {
	pqu.mutation.ResetPricePerUnit()
	pqu.mutation.SetPricePerUnit(d)
	return pqu
}

// SetNillablePricePerUnit sets the "price_per_unit" field if the given value is not nil.
func (pqu *ProductQtyUpdate) SetNillablePricePerUnit(d *decimal.Decimal) *ProductQtyUpdate {
	if d != nil {
		pqu.SetPricePerUnit(*d)
	}
	return pqu
}

// AddPricePerUnit adds d to the "price_per_unit" field.
func (pqu *ProductQtyUpdate) AddPricePerUnit(d decimal.Decimal) *ProductQtyUpdate {
	pqu.mutation.AddPricePerUnit(d)
	return pqu
}

// SetQty sets the "qty" field.
func (pqu *ProductQtyUpdate) SetQty(d decimal.Decimal) *ProductQtyUpdate {
	pqu.mutation.ResetQty()
	pqu.mutation.SetQty(d)
	return pqu
}

// SetNillableQty sets the "qty" field if the given value is not nil.
func (pqu *ProductQtyUpdate) SetNillableQty(d *decimal.Decimal) *ProductQtyUpdate {
	if d != nil {
		pqu.SetQty(*d)
	}
	return pqu
}

// AddQty adds d to the "qty" field.
func (pqu *ProductQtyUpdate) AddQty(d decimal.Decimal) *ProductQtyUpdate {
	pqu.mutation.AddQty(d)
	return pqu
}

// Mutation returns the ProductQtyMutation object of the builder.
func (pqu *ProductQtyUpdate) Mutation() *ProductQtyMutation {
	return pqu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pqu *ProductQtyUpdate) Save(ctx context.Context) (int, error) {
	pqu.defaults()
	return withHooks(ctx, pqu.sqlSave, pqu.mutation, pqu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pqu *ProductQtyUpdate) SaveX(ctx context.Context) int {
	affected, err := pqu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pqu *ProductQtyUpdate) Exec(ctx context.Context) error {
	_, err := pqu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pqu *ProductQtyUpdate) ExecX(ctx context.Context) {
	if err := pqu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pqu *ProductQtyUpdate) defaults() {
	if _, ok := pqu.mutation.UpdatedAt(); !ok {
		v := productqty.UpdateDefaultUpdatedAt()
		pqu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pqu *ProductQtyUpdate) check() error {
	if _, ok := pqu.mutation.WorkUnitID(); pqu.mutation.WorkUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProductQty.work_unit"`)
	}
	if _, ok := pqu.mutation.ProductID(); pqu.mutation.ProductCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProductQty.product"`)
	}
	if _, ok := pqu.mutation.ProductColorID(); pqu.mutation.ProductColorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProductQty.product_color"`)
	}
	return nil
}

func (pqu *ProductQtyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pqu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(productqty.Table, productqty.Columns, sqlgraph.NewFieldSpec(productqty.FieldID, field.TypeUUID))
	if ps := pqu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pqu.mutation.UpdatedAt(); ok {
		_spec.SetField(productqty.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pqu.mutation.PricePerUnit(); ok {
		_spec.SetField(productqty.FieldPricePerUnit, field.TypeFloat64, value)
	}
	if value, ok := pqu.mutation.AddedPricePerUnit(); ok {
		_spec.AddField(productqty.FieldPricePerUnit, field.TypeFloat64, value)
	}
	if value, ok := pqu.mutation.Qty(); ok {
		_spec.SetField(productqty.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := pqu.mutation.AddedQty(); ok {
		_spec.AddField(productqty.FieldQty, field.TypeFloat64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pqu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productqty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pqu.mutation.done = true
	return n, nil
}

// ProductQtyUpdateOne is the builder for updating a single ProductQty entity.
type ProductQtyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductQtyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pquo *ProductQtyUpdateOne) SetUpdatedAt(t time.Time) *ProductQtyUpdateOne {
	pquo.mutation.SetUpdatedAt(t)
	return pquo
}

// SetPricePerUnit sets the "price_per_unit" field.
func (pquo *ProductQtyUpdateOne) SetPricePerUnit(d decimal.Decimal) *ProductQtyUpdateOne {
	pquo.mutation.ResetPricePerUnit()
	pquo.mutation.SetPricePerUnit(d)
	return pquo
}

// SetNillablePricePerUnit sets the "price_per_unit" field if the given value is not nil.
func (pquo *ProductQtyUpdateOne) SetNillablePricePerUnit(d *decimal.Decimal) *ProductQtyUpdateOne {
	if d != nil {
		pquo.SetPricePerUnit(*d)
	}
	return pquo
}

// AddPricePerUnit adds d to the "price_per_unit" field.
func (pquo *ProductQtyUpdateOne) AddPricePerUnit(d decimal.Decimal) *ProductQtyUpdateOne {
	pquo.mutation.AddPricePerUnit(d)
	return pquo
}

// SetQty sets the "qty" field.
func (pquo *ProductQtyUpdateOne) SetQty(d decimal.Decimal) *ProductQtyUpdateOne {
	pquo.mutation.ResetQty()
	pquo.mutation.SetQty(d)
	return pquo
}

// SetNillableQty sets the "qty" field if the given value is not nil.
func (pquo *ProductQtyUpdateOne) SetNillableQty(d *decimal.Decimal) *ProductQtyUpdateOne {
	if d != nil {
		pquo.SetQty(*d)
	}
	return pquo
}

// AddQty adds d to the "qty" field.
func (pquo *ProductQtyUpdateOne) AddQty(d decimal.Decimal) *ProductQtyUpdateOne {
	pquo.mutation.AddQty(d)
	return pquo
}

// Mutation returns the ProductQtyMutation object of the builder.
func (pquo *ProductQtyUpdateOne) Mutation() *ProductQtyMutation {
	return pquo.mutation
}

// Where appends a list predicates to the ProductQtyUpdate builder.
func (pquo *ProductQtyUpdateOne) Where(ps ...predicate.ProductQty) *ProductQtyUpdateOne {
	pquo.mutation.Where(ps...)
	return pquo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pquo *ProductQtyUpdateOne) Select(field string, fields ...string) *ProductQtyUpdateOne {
	pquo.fields = append([]string{field}, fields...)
	return pquo
}

// Save executes the query and returns the updated ProductQty entity.
func (pquo *ProductQtyUpdateOne) Save(ctx context.Context) (*ProductQty, error) {
	pquo.defaults()
	return withHooks(ctx, pquo.sqlSave, pquo.mutation, pquo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pquo *ProductQtyUpdateOne) SaveX(ctx context.Context) *ProductQty {
	node, err := pquo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pquo *ProductQtyUpdateOne) Exec(ctx context.Context) error {
	_, err := pquo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pquo *ProductQtyUpdateOne) ExecX(ctx context.Context) {
	if err := pquo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pquo *ProductQtyUpdateOne) defaults() {
	if _, ok := pquo.mutation.UpdatedAt(); !ok {
		v := productqty.UpdateDefaultUpdatedAt()
		pquo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pquo *ProductQtyUpdateOne) check() error {
	if _, ok := pquo.mutation.WorkUnitID(); pquo.mutation.WorkUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProductQty.work_unit"`)
	}
	if _, ok := pquo.mutation.ProductID(); pquo.mutation.ProductCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProductQty.product"`)
	}
	if _, ok := pquo.mutation.ProductColorID(); pquo.mutation.ProductColorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProductQty.product_color"`)
	}
	return nil
}

func (pquo *ProductQtyUpdateOne) sqlSave(ctx context.Context) (_node *ProductQty, err error) {
	if err := pquo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(productqty.Table, productqty.Columns, sqlgraph.NewFieldSpec(productqty.FieldID, field.TypeUUID))
	id, ok := pquo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProductQty.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pquo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productqty.FieldID)
		for _, f := range fields {
			if !productqty.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productqty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pquo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pquo.mutation.UpdatedAt(); ok {
		_spec.SetField(productqty.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pquo.mutation.PricePerUnit(); ok {
		_spec.SetField(productqty.FieldPricePerUnit, field.TypeFloat64, value)
	}
	if value, ok := pquo.mutation.AddedPricePerUnit(); ok {
		_spec.AddField(productqty.FieldPricePerUnit, field.TypeFloat64, value)
	}
	if value, ok := pquo.mutation.Qty(); ok {
		_spec.SetField(productqty.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := pquo.mutation.AddedQty(); ok {
		_spec.AddField(productqty.FieldQty, field.TypeFloat64, value)
	}
	_node = &ProductQty{config: pquo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pquo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productqty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pquo.mutation.done = true
	return _node, nil
}