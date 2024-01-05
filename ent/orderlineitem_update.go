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
	"github.com/thaiha1607/foursquare_server/ent/orderlineitem"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// OrderLineItemUpdate is the builder for updating OrderLineItem entities.
type OrderLineItemUpdate struct {
	config
	hooks    []Hook
	mutation *OrderLineItemMutation
}

// Where appends a list predicates to the OrderLineItemUpdate builder.
func (oliu *OrderLineItemUpdate) Where(ps ...predicate.OrderLineItem) *OrderLineItemUpdate {
	oliu.mutation.Where(ps...)
	return oliu
}

// SetUpdatedAt sets the "updated_at" field.
func (oliu *OrderLineItemUpdate) SetUpdatedAt(t time.Time) *OrderLineItemUpdate {
	oliu.mutation.SetUpdatedAt(t)
	return oliu
}

// SetQty sets the "qty" field.
func (oliu *OrderLineItemUpdate) SetQty(d decimal.Decimal) *OrderLineItemUpdate {
	oliu.mutation.ResetQty()
	oliu.mutation.SetQty(d)
	return oliu
}

// SetNillableQty sets the "qty" field if the given value is not nil.
func (oliu *OrderLineItemUpdate) SetNillableQty(d *decimal.Decimal) *OrderLineItemUpdate {
	if d != nil {
		oliu.SetQty(*d)
	}
	return oliu
}

// AddQty adds d to the "qty" field.
func (oliu *OrderLineItemUpdate) AddQty(d decimal.Decimal) *OrderLineItemUpdate {
	oliu.mutation.AddQty(d)
	return oliu
}

// Mutation returns the OrderLineItemMutation object of the builder.
func (oliu *OrderLineItemUpdate) Mutation() *OrderLineItemMutation {
	return oliu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oliu *OrderLineItemUpdate) Save(ctx context.Context) (int, error) {
	oliu.defaults()
	return withHooks(ctx, oliu.sqlSave, oliu.mutation, oliu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oliu *OrderLineItemUpdate) SaveX(ctx context.Context) int {
	affected, err := oliu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oliu *OrderLineItemUpdate) Exec(ctx context.Context) error {
	_, err := oliu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oliu *OrderLineItemUpdate) ExecX(ctx context.Context) {
	if err := oliu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oliu *OrderLineItemUpdate) defaults() {
	if _, ok := oliu.mutation.UpdatedAt(); !ok {
		v := orderlineitem.UpdateDefaultUpdatedAt()
		oliu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oliu *OrderLineItemUpdate) check() error {
	if _, ok := oliu.mutation.OrderID(); oliu.mutation.OrderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderLineItem.order"`)
	}
	if _, ok := oliu.mutation.ProductID(); oliu.mutation.ProductCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderLineItem.product"`)
	}
	return nil
}

func (oliu *OrderLineItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oliu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderlineitem.Table, orderlineitem.Columns, sqlgraph.NewFieldSpec(orderlineitem.FieldID, field.TypeUUID))
	if ps := oliu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oliu.mutation.UpdatedAt(); ok {
		_spec.SetField(orderlineitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := oliu.mutation.Qty(); ok {
		_spec.SetField(orderlineitem.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := oliu.mutation.AddedQty(); ok {
		_spec.AddField(orderlineitem.FieldQty, field.TypeFloat64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oliu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderlineitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oliu.mutation.done = true
	return n, nil
}

// OrderLineItemUpdateOne is the builder for updating a single OrderLineItem entity.
type OrderLineItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderLineItemMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (oliuo *OrderLineItemUpdateOne) SetUpdatedAt(t time.Time) *OrderLineItemUpdateOne {
	oliuo.mutation.SetUpdatedAt(t)
	return oliuo
}

// SetQty sets the "qty" field.
func (oliuo *OrderLineItemUpdateOne) SetQty(d decimal.Decimal) *OrderLineItemUpdateOne {
	oliuo.mutation.ResetQty()
	oliuo.mutation.SetQty(d)
	return oliuo
}

// SetNillableQty sets the "qty" field if the given value is not nil.
func (oliuo *OrderLineItemUpdateOne) SetNillableQty(d *decimal.Decimal) *OrderLineItemUpdateOne {
	if d != nil {
		oliuo.SetQty(*d)
	}
	return oliuo
}

// AddQty adds d to the "qty" field.
func (oliuo *OrderLineItemUpdateOne) AddQty(d decimal.Decimal) *OrderLineItemUpdateOne {
	oliuo.mutation.AddQty(d)
	return oliuo
}

// Mutation returns the OrderLineItemMutation object of the builder.
func (oliuo *OrderLineItemUpdateOne) Mutation() *OrderLineItemMutation {
	return oliuo.mutation
}

// Where appends a list predicates to the OrderLineItemUpdate builder.
func (oliuo *OrderLineItemUpdateOne) Where(ps ...predicate.OrderLineItem) *OrderLineItemUpdateOne {
	oliuo.mutation.Where(ps...)
	return oliuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oliuo *OrderLineItemUpdateOne) Select(field string, fields ...string) *OrderLineItemUpdateOne {
	oliuo.fields = append([]string{field}, fields...)
	return oliuo
}

// Save executes the query and returns the updated OrderLineItem entity.
func (oliuo *OrderLineItemUpdateOne) Save(ctx context.Context) (*OrderLineItem, error) {
	oliuo.defaults()
	return withHooks(ctx, oliuo.sqlSave, oliuo.mutation, oliuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oliuo *OrderLineItemUpdateOne) SaveX(ctx context.Context) *OrderLineItem {
	node, err := oliuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oliuo *OrderLineItemUpdateOne) Exec(ctx context.Context) error {
	_, err := oliuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oliuo *OrderLineItemUpdateOne) ExecX(ctx context.Context) {
	if err := oliuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oliuo *OrderLineItemUpdateOne) defaults() {
	if _, ok := oliuo.mutation.UpdatedAt(); !ok {
		v := orderlineitem.UpdateDefaultUpdatedAt()
		oliuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oliuo *OrderLineItemUpdateOne) check() error {
	if _, ok := oliuo.mutation.OrderID(); oliuo.mutation.OrderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderLineItem.order"`)
	}
	if _, ok := oliuo.mutation.ProductID(); oliuo.mutation.ProductCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderLineItem.product"`)
	}
	return nil
}

func (oliuo *OrderLineItemUpdateOne) sqlSave(ctx context.Context) (_node *OrderLineItem, err error) {
	if err := oliuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderlineitem.Table, orderlineitem.Columns, sqlgraph.NewFieldSpec(orderlineitem.FieldID, field.TypeUUID))
	id, ok := oliuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderLineItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oliuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderlineitem.FieldID)
		for _, f := range fields {
			if !orderlineitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderlineitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oliuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oliuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orderlineitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := oliuo.mutation.Qty(); ok {
		_spec.SetField(orderlineitem.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := oliuo.mutation.AddedQty(); ok {
		_spec.AddField(orderlineitem.FieldQty, field.TypeFloat64, value)
	}
	_node = &OrderLineItem{config: oliuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oliuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderlineitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oliuo.mutation.done = true
	return _node, nil
}
