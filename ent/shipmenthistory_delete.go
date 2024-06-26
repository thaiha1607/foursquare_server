// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/shipmenthistory"
)

// ShipmentHistoryDelete is the builder for deleting a ShipmentHistory entity.
type ShipmentHistoryDelete struct {
	config
	hooks    []Hook
	mutation *ShipmentHistoryMutation
}

// Where appends a list predicates to the ShipmentHistoryDelete builder.
func (shd *ShipmentHistoryDelete) Where(ps ...predicate.ShipmentHistory) *ShipmentHistoryDelete {
	shd.mutation.Where(ps...)
	return shd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (shd *ShipmentHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, shd.sqlExec, shd.mutation, shd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (shd *ShipmentHistoryDelete) ExecX(ctx context.Context) int {
	n, err := shd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (shd *ShipmentHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(shipmenthistory.Table, sqlgraph.NewFieldSpec(shipmenthistory.FieldID, field.TypeUUID))
	if ps := shd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, shd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	shd.mutation.done = true
	return affected, err
}

// ShipmentHistoryDeleteOne is the builder for deleting a single ShipmentHistory entity.
type ShipmentHistoryDeleteOne struct {
	shd *ShipmentHistoryDelete
}

// Where appends a list predicates to the ShipmentHistoryDelete builder.
func (shdo *ShipmentHistoryDeleteOne) Where(ps ...predicate.ShipmentHistory) *ShipmentHistoryDeleteOne {
	shdo.shd.mutation.Where(ps...)
	return shdo
}

// Exec executes the deletion query.
func (shdo *ShipmentHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := shdo.shd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shipmenthistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (shdo *ShipmentHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := shdo.Exec(ctx); err != nil {
		panic(err)
	}
}
