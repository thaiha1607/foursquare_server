// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/orderlineitem"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// OrderLineItemDelete is the builder for deleting a OrderLineItem entity.
type OrderLineItemDelete struct {
	config
	hooks    []Hook
	mutation *OrderLineItemMutation
}

// Where appends a list predicates to the OrderLineItemDelete builder.
func (olid *OrderLineItemDelete) Where(ps ...predicate.OrderLineItem) *OrderLineItemDelete {
	olid.mutation.Where(ps...)
	return olid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (olid *OrderLineItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, olid.sqlExec, olid.mutation, olid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (olid *OrderLineItemDelete) ExecX(ctx context.Context) int {
	n, err := olid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (olid *OrderLineItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orderlineitem.Table, sqlgraph.NewFieldSpec(orderlineitem.FieldID, field.TypeUUID))
	if ps := olid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, olid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	olid.mutation.done = true
	return affected, err
}

// OrderLineItemDeleteOne is the builder for deleting a single OrderLineItem entity.
type OrderLineItemDeleteOne struct {
	olid *OrderLineItemDelete
}

// Where appends a list predicates to the OrderLineItemDelete builder.
func (olido *OrderLineItemDeleteOne) Where(ps ...predicate.OrderLineItem) *OrderLineItemDeleteOne {
	olido.olid.mutation.Where(ps...)
	return olido
}

// Exec executes the deletion query.
func (olido *OrderLineItemDeleteOne) Exec(ctx context.Context) error {
	n, err := olido.olid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderlineitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (olido *OrderLineItemDeleteOne) ExecX(ctx context.Context) {
	if err := olido.Exec(ctx); err != nil {
		panic(err)
	}
}
