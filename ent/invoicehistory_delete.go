// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/invoicehistory"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// InvoiceHistoryDelete is the builder for deleting a InvoiceHistory entity.
type InvoiceHistoryDelete struct {
	config
	hooks    []Hook
	mutation *InvoiceHistoryMutation
}

// Where appends a list predicates to the InvoiceHistoryDelete builder.
func (ihd *InvoiceHistoryDelete) Where(ps ...predicate.InvoiceHistory) *InvoiceHistoryDelete {
	ihd.mutation.Where(ps...)
	return ihd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ihd *InvoiceHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ihd.sqlExec, ihd.mutation, ihd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ihd *InvoiceHistoryDelete) ExecX(ctx context.Context) int {
	n, err := ihd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ihd *InvoiceHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(invoicehistory.Table, sqlgraph.NewFieldSpec(invoicehistory.FieldID, field.TypeUUID))
	if ps := ihd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ihd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ihd.mutation.done = true
	return affected, err
}

// InvoiceHistoryDeleteOne is the builder for deleting a single InvoiceHistory entity.
type InvoiceHistoryDeleteOne struct {
	ihd *InvoiceHistoryDelete
}

// Where appends a list predicates to the InvoiceHistoryDelete builder.
func (ihdo *InvoiceHistoryDeleteOne) Where(ps ...predicate.InvoiceHistory) *InvoiceHistoryDeleteOne {
	ihdo.ihd.mutation.Where(ps...)
	return ihdo
}

// Exec executes the deletion query.
func (ihdo *InvoiceHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := ihdo.ihd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{invoicehistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ihdo *InvoiceHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := ihdo.Exec(ctx); err != nil {
		panic(err)
	}
}