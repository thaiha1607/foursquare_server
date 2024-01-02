// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/ordertype"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// OrderTypeDelete is the builder for deleting a OrderType entity.
type OrderTypeDelete struct {
	config
	hooks    []Hook
	mutation *OrderTypeMutation
}

// Where appends a list predicates to the OrderTypeDelete builder.
func (otd *OrderTypeDelete) Where(ps ...predicate.OrderType) *OrderTypeDelete {
	otd.mutation.Where(ps...)
	return otd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (otd *OrderTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, otd.sqlExec, otd.mutation, otd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (otd *OrderTypeDelete) ExecX(ctx context.Context) int {
	n, err := otd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (otd *OrderTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(ordertype.Table, sqlgraph.NewFieldSpec(ordertype.FieldID, field.TypeInt))
	if ps := otd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, otd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	otd.mutation.done = true
	return affected, err
}

// OrderTypeDeleteOne is the builder for deleting a single OrderType entity.
type OrderTypeDeleteOne struct {
	otd *OrderTypeDelete
}

// Where appends a list predicates to the OrderTypeDelete builder.
func (otdo *OrderTypeDeleteOne) Where(ps ...predicate.OrderType) *OrderTypeDeleteOne {
	otdo.otd.mutation.Where(ps...)
	return otdo
}

// Exec executes the deletion query.
func (otdo *OrderTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := otdo.otd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ordertype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (otdo *OrderTypeDeleteOne) ExecX(ctx context.Context) {
	if err := otdo.Exec(ctx); err != nil {
		panic(err)
	}
}