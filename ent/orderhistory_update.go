// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/orderhistory"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// OrderHistoryUpdate is the builder for updating OrderHistory entities.
type OrderHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *OrderHistoryMutation
}

// Where appends a list predicates to the OrderHistoryUpdate builder.
func (ohu *OrderHistoryUpdate) Where(ps ...predicate.OrderHistory) *OrderHistoryUpdate {
	ohu.mutation.Where(ps...)
	return ohu
}

// Mutation returns the OrderHistoryMutation object of the builder.
func (ohu *OrderHistoryUpdate) Mutation() *OrderHistoryMutation {
	return ohu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ohu *OrderHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ohu.sqlSave, ohu.mutation, ohu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ohu *OrderHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ohu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ohu *OrderHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ohu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohu *OrderHistoryUpdate) ExecX(ctx context.Context) {
	if err := ohu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ohu *OrderHistoryUpdate) check() error {
	if _, ok := ohu.mutation.OrderID(); ohu.mutation.OrderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderHistory.order"`)
	}
	if _, ok := ohu.mutation.PersonID(); ohu.mutation.PersonCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderHistory.person"`)
	}
	return nil
}

func (ohu *OrderHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ohu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderhistory.Table, orderhistory.Columns, sqlgraph.NewFieldSpec(orderhistory.FieldID, field.TypeUUID))
	if ps := ohu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ohu.mutation.DescriptionCleared() {
		_spec.ClearField(orderhistory.FieldDescription, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ohu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ohu.mutation.done = true
	return n, nil
}

// OrderHistoryUpdateOne is the builder for updating a single OrderHistory entity.
type OrderHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderHistoryMutation
}

// Mutation returns the OrderHistoryMutation object of the builder.
func (ohuo *OrderHistoryUpdateOne) Mutation() *OrderHistoryMutation {
	return ohuo.mutation
}

// Where appends a list predicates to the OrderHistoryUpdate builder.
func (ohuo *OrderHistoryUpdateOne) Where(ps ...predicate.OrderHistory) *OrderHistoryUpdateOne {
	ohuo.mutation.Where(ps...)
	return ohuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ohuo *OrderHistoryUpdateOne) Select(field string, fields ...string) *OrderHistoryUpdateOne {
	ohuo.fields = append([]string{field}, fields...)
	return ohuo
}

// Save executes the query and returns the updated OrderHistory entity.
func (ohuo *OrderHistoryUpdateOne) Save(ctx context.Context) (*OrderHistory, error) {
	return withHooks(ctx, ohuo.sqlSave, ohuo.mutation, ohuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ohuo *OrderHistoryUpdateOne) SaveX(ctx context.Context) *OrderHistory {
	node, err := ohuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ohuo *OrderHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ohuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohuo *OrderHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ohuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ohuo *OrderHistoryUpdateOne) check() error {
	if _, ok := ohuo.mutation.OrderID(); ohuo.mutation.OrderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderHistory.order"`)
	}
	if _, ok := ohuo.mutation.PersonID(); ohuo.mutation.PersonCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrderHistory.person"`)
	}
	return nil
}

func (ohuo *OrderHistoryUpdateOne) sqlSave(ctx context.Context) (_node *OrderHistory, err error) {
	if err := ohuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderhistory.Table, orderhistory.Columns, sqlgraph.NewFieldSpec(orderhistory.FieldID, field.TypeUUID))
	id, ok := ohuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ohuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderhistory.FieldID)
		for _, f := range fields {
			if !orderhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ohuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ohuo.mutation.DescriptionCleared() {
		_spec.ClearField(orderhistory.FieldDescription, field.TypeString)
	}
	_node = &OrderHistory{config: ohuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ohuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ohuo.mutation.done = true
	return _node, nil
}
