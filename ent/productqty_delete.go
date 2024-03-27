// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/productqty"
)

// ProductQtyDelete is the builder for deleting a ProductQty entity.
type ProductQtyDelete struct {
	config
	hooks    []Hook
	mutation *ProductQtyMutation
}

// Where appends a list predicates to the ProductQtyDelete builder.
func (pqd *ProductQtyDelete) Where(ps ...predicate.ProductQty) *ProductQtyDelete {
	pqd.mutation.Where(ps...)
	return pqd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pqd *ProductQtyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pqd.sqlExec, pqd.mutation, pqd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pqd *ProductQtyDelete) ExecX(ctx context.Context) int {
	n, err := pqd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pqd *ProductQtyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(productqty.Table, sqlgraph.NewFieldSpec(productqty.FieldID, field.TypeUUID))
	if ps := pqd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pqd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pqd.mutation.done = true
	return affected, err
}

// ProductQtyDeleteOne is the builder for deleting a single ProductQty entity.
type ProductQtyDeleteOne struct {
	pqd *ProductQtyDelete
}

// Where appends a list predicates to the ProductQtyDelete builder.
func (pqdo *ProductQtyDeleteOne) Where(ps ...predicate.ProductQty) *ProductQtyDeleteOne {
	pqdo.pqd.mutation.Where(ps...)
	return pqdo
}

// Exec executes the deletion query.
func (pqdo *ProductQtyDeleteOne) Exec(ctx context.Context) error {
	n, err := pqdo.pqd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productqty.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pqdo *ProductQtyDeleteOne) ExecX(ctx context.Context) {
	if err := pqdo.Exec(ctx); err != nil {
		panic(err)
	}
}
