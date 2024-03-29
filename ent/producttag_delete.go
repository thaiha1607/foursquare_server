// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/producttag"
)

// ProductTagDelete is the builder for deleting a ProductTag entity.
type ProductTagDelete struct {
	config
	hooks    []Hook
	mutation *ProductTagMutation
}

// Where appends a list predicates to the ProductTagDelete builder.
func (ptd *ProductTagDelete) Where(ps ...predicate.ProductTag) *ProductTagDelete {
	ptd.mutation.Where(ps...)
	return ptd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ptd *ProductTagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ptd.sqlExec, ptd.mutation, ptd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ptd *ProductTagDelete) ExecX(ctx context.Context) int {
	n, err := ptd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ptd *ProductTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(producttag.Table, nil)
	if ps := ptd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ptd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ptd.mutation.done = true
	return affected, err
}

// ProductTagDeleteOne is the builder for deleting a single ProductTag entity.
type ProductTagDeleteOne struct {
	ptd *ProductTagDelete
}

// Where appends a list predicates to the ProductTagDelete builder.
func (ptdo *ProductTagDeleteOne) Where(ps ...predicate.ProductTag) *ProductTagDeleteOne {
	ptdo.ptd.mutation.Where(ps...)
	return ptdo
}

// Exec executes the deletion query.
func (ptdo *ProductTagDeleteOne) Exec(ctx context.Context) error {
	n, err := ptdo.ptd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{producttag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ptdo *ProductTagDeleteOne) ExecX(ctx context.Context) {
	if err := ptdo.Exec(ctx); err != nil {
		panic(err)
	}
}
