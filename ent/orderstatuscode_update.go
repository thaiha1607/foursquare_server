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
	"github.com/thaiha1607/foursquare_server/ent/orderstatuscode"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// OrderStatusCodeUpdate is the builder for updating OrderStatusCode entities.
type OrderStatusCodeUpdate struct {
	config
	hooks    []Hook
	mutation *OrderStatusCodeMutation
}

// Where appends a list predicates to the OrderStatusCodeUpdate builder.
func (oscu *OrderStatusCodeUpdate) Where(ps ...predicate.OrderStatusCode) *OrderStatusCodeUpdate {
	oscu.mutation.Where(ps...)
	return oscu
}

// SetUpdatedAt sets the "updated_at" field.
func (oscu *OrderStatusCodeUpdate) SetUpdatedAt(t time.Time) *OrderStatusCodeUpdate {
	oscu.mutation.SetUpdatedAt(t)
	return oscu
}

// SetOrderStatus sets the "order_status" field.
func (oscu *OrderStatusCodeUpdate) SetOrderStatus(s string) *OrderStatusCodeUpdate {
	oscu.mutation.SetOrderStatus(s)
	return oscu
}

// SetNillableOrderStatus sets the "order_status" field if the given value is not nil.
func (oscu *OrderStatusCodeUpdate) SetNillableOrderStatus(s *string) *OrderStatusCodeUpdate {
	if s != nil {
		oscu.SetOrderStatus(*s)
	}
	return oscu
}

// Mutation returns the OrderStatusCodeMutation object of the builder.
func (oscu *OrderStatusCodeUpdate) Mutation() *OrderStatusCodeMutation {
	return oscu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oscu *OrderStatusCodeUpdate) Save(ctx context.Context) (int, error) {
	oscu.defaults()
	return withHooks(ctx, oscu.sqlSave, oscu.mutation, oscu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oscu *OrderStatusCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := oscu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oscu *OrderStatusCodeUpdate) Exec(ctx context.Context) error {
	_, err := oscu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oscu *OrderStatusCodeUpdate) ExecX(ctx context.Context) {
	if err := oscu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oscu *OrderStatusCodeUpdate) defaults() {
	if _, ok := oscu.mutation.UpdatedAt(); !ok {
		v := orderstatuscode.UpdateDefaultUpdatedAt()
		oscu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oscu *OrderStatusCodeUpdate) check() error {
	if v, ok := oscu.mutation.OrderStatus(); ok {
		if err := orderstatuscode.OrderStatusValidator(v); err != nil {
			return &ValidationError{Name: "order_status", err: fmt.Errorf(`ent: validator failed for field "OrderStatusCode.order_status": %w`, err)}
		}
	}
	return nil
}

func (oscu *OrderStatusCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oscu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderstatuscode.Table, orderstatuscode.Columns, sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt))
	if ps := oscu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oscu.mutation.UpdatedAt(); ok {
		_spec.SetField(orderstatuscode.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := oscu.mutation.OrderStatus(); ok {
		_spec.SetField(orderstatuscode.FieldOrderStatus, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oscu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderstatuscode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oscu.mutation.done = true
	return n, nil
}

// OrderStatusCodeUpdateOne is the builder for updating a single OrderStatusCode entity.
type OrderStatusCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderStatusCodeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (oscuo *OrderStatusCodeUpdateOne) SetUpdatedAt(t time.Time) *OrderStatusCodeUpdateOne {
	oscuo.mutation.SetUpdatedAt(t)
	return oscuo
}

// SetOrderStatus sets the "order_status" field.
func (oscuo *OrderStatusCodeUpdateOne) SetOrderStatus(s string) *OrderStatusCodeUpdateOne {
	oscuo.mutation.SetOrderStatus(s)
	return oscuo
}

// SetNillableOrderStatus sets the "order_status" field if the given value is not nil.
func (oscuo *OrderStatusCodeUpdateOne) SetNillableOrderStatus(s *string) *OrderStatusCodeUpdateOne {
	if s != nil {
		oscuo.SetOrderStatus(*s)
	}
	return oscuo
}

// Mutation returns the OrderStatusCodeMutation object of the builder.
func (oscuo *OrderStatusCodeUpdateOne) Mutation() *OrderStatusCodeMutation {
	return oscuo.mutation
}

// Where appends a list predicates to the OrderStatusCodeUpdate builder.
func (oscuo *OrderStatusCodeUpdateOne) Where(ps ...predicate.OrderStatusCode) *OrderStatusCodeUpdateOne {
	oscuo.mutation.Where(ps...)
	return oscuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oscuo *OrderStatusCodeUpdateOne) Select(field string, fields ...string) *OrderStatusCodeUpdateOne {
	oscuo.fields = append([]string{field}, fields...)
	return oscuo
}

// Save executes the query and returns the updated OrderStatusCode entity.
func (oscuo *OrderStatusCodeUpdateOne) Save(ctx context.Context) (*OrderStatusCode, error) {
	oscuo.defaults()
	return withHooks(ctx, oscuo.sqlSave, oscuo.mutation, oscuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oscuo *OrderStatusCodeUpdateOne) SaveX(ctx context.Context) *OrderStatusCode {
	node, err := oscuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oscuo *OrderStatusCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := oscuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oscuo *OrderStatusCodeUpdateOne) ExecX(ctx context.Context) {
	if err := oscuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oscuo *OrderStatusCodeUpdateOne) defaults() {
	if _, ok := oscuo.mutation.UpdatedAt(); !ok {
		v := orderstatuscode.UpdateDefaultUpdatedAt()
		oscuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oscuo *OrderStatusCodeUpdateOne) check() error {
	if v, ok := oscuo.mutation.OrderStatus(); ok {
		if err := orderstatuscode.OrderStatusValidator(v); err != nil {
			return &ValidationError{Name: "order_status", err: fmt.Errorf(`ent: validator failed for field "OrderStatusCode.order_status": %w`, err)}
		}
	}
	return nil
}

func (oscuo *OrderStatusCodeUpdateOne) sqlSave(ctx context.Context) (_node *OrderStatusCode, err error) {
	if err := oscuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderstatuscode.Table, orderstatuscode.Columns, sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt))
	id, ok := oscuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderStatusCode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oscuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderstatuscode.FieldID)
		for _, f := range fields {
			if !orderstatuscode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderstatuscode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oscuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oscuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orderstatuscode.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := oscuo.mutation.OrderStatus(); ok {
		_spec.SetField(orderstatuscode.FieldOrderStatus, field.TypeString, value)
	}
	_node = &OrderStatusCode{config: oscuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oscuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderstatuscode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oscuo.mutation.done = true
	return _node, nil
}
