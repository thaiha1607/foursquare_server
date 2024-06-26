// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/productcolor"
)

// ProductColorDelete is the builder for deleting a ProductColor entity.
type ProductColorDelete struct {
	config
	hooks    []Hook
	mutation *ProductColorMutation
}

// Where appends a list predicates to the ProductColorDelete builder.
func (pcd *ProductColorDelete) Where(ps ...predicate.ProductColor) *ProductColorDelete {
	pcd.mutation.Where(ps...)
	return pcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pcd *ProductColorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pcd.sqlExec, pcd.mutation, pcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pcd *ProductColorDelete) ExecX(ctx context.Context) int {
	n, err := pcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pcd *ProductColorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(productcolor.Table, sqlgraph.NewFieldSpec(productcolor.FieldID, field.TypeString))
	if ps := pcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pcd.mutation.done = true
	return affected, err
}

// ProductColorDeleteOne is the builder for deleting a single ProductColor entity.
type ProductColorDeleteOne struct {
	pcd *ProductColorDelete
}

// Where appends a list predicates to the ProductColorDelete builder.
func (pcdo *ProductColorDeleteOne) Where(ps ...predicate.ProductColor) *ProductColorDeleteOne {
	pcdo.pcd.mutation.Where(ps...)
	return pcdo
}

// Exec executes the deletion query.
func (pcdo *ProductColorDeleteOne) Exec(ctx context.Context) error {
	n, err := pcdo.pcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productcolor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pcdo *ProductColorDeleteOne) ExecX(ctx context.Context) {
	if err := pcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
