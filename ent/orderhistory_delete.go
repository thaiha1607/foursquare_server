// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/orderhistory"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// OrderHistoryDelete is the builder for deleting a OrderHistory entity.
type OrderHistoryDelete struct {
	config
	hooks    []Hook
	mutation *OrderHistoryMutation
}

// Where appends a list predicates to the OrderHistoryDelete builder.
func (ohd *OrderHistoryDelete) Where(ps ...predicate.OrderHistory) *OrderHistoryDelete {
	ohd.mutation.Where(ps...)
	return ohd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ohd *OrderHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ohd.sqlExec, ohd.mutation, ohd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ohd *OrderHistoryDelete) ExecX(ctx context.Context) int {
	n, err := ohd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ohd *OrderHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orderhistory.Table, sqlgraph.NewFieldSpec(orderhistory.FieldID, field.TypeUUID))
	if ps := ohd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ohd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ohd.mutation.done = true
	return affected, err
}

// OrderHistoryDeleteOne is the builder for deleting a single OrderHistory entity.
type OrderHistoryDeleteOne struct {
	ohd *OrderHistoryDelete
}

// Where appends a list predicates to the OrderHistoryDelete builder.
func (ohdo *OrderHistoryDeleteOne) Where(ps ...predicate.OrderHistory) *OrderHistoryDeleteOne {
	ohdo.ohd.mutation.Where(ps...)
	return ohdo
}

// Exec executes the deletion query.
func (ohdo *OrderHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := ohdo.ohd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ohdo *OrderHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := ohdo.Exec(ctx); err != nil {
		panic(err)
	}
}