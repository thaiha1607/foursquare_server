// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/productinfo"
	"github.com/thaiha1607/foursquare_server/ent/producttag"
	"github.com/thaiha1607/foursquare_server/ent/tag"
)

// ProductTagCreate is the builder for creating a ProductTag entity.
type ProductTagCreate struct {
	config
	mutation *ProductTagMutation
	hooks    []Hook
}

// SetProductID sets the "product_id" field.
func (ptc *ProductTagCreate) SetProductID(s string) *ProductTagCreate {
	ptc.mutation.SetProductID(s)
	return ptc
}

// SetTagID sets the "tag_id" field.
func (ptc *ProductTagCreate) SetTagID(s string) *ProductTagCreate {
	ptc.mutation.SetTagID(s)
	return ptc
}

// SetProductsID sets the "products" edge to the ProductInfo entity by ID.
func (ptc *ProductTagCreate) SetProductsID(id string) *ProductTagCreate {
	ptc.mutation.SetProductsID(id)
	return ptc
}

// SetProducts sets the "products" edge to the ProductInfo entity.
func (ptc *ProductTagCreate) SetProducts(p *ProductInfo) *ProductTagCreate {
	return ptc.SetProductsID(p.ID)
}

// SetTagsID sets the "tags" edge to the Tag entity by ID.
func (ptc *ProductTagCreate) SetTagsID(id string) *ProductTagCreate {
	ptc.mutation.SetTagsID(id)
	return ptc
}

// SetTags sets the "tags" edge to the Tag entity.
func (ptc *ProductTagCreate) SetTags(t *Tag) *ProductTagCreate {
	return ptc.SetTagsID(t.ID)
}

// Mutation returns the ProductTagMutation object of the builder.
func (ptc *ProductTagCreate) Mutation() *ProductTagMutation {
	return ptc.mutation
}

// Save creates the ProductTag in the database.
func (ptc *ProductTagCreate) Save(ctx context.Context) (*ProductTag, error) {
	return withHooks(ctx, ptc.sqlSave, ptc.mutation, ptc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *ProductTagCreate) SaveX(ctx context.Context) *ProductTag {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptc *ProductTagCreate) Exec(ctx context.Context) error {
	_, err := ptc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptc *ProductTagCreate) ExecX(ctx context.Context) {
	if err := ptc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptc *ProductTagCreate) check() error {
	if _, ok := ptc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`ent: missing required field "ProductTag.product_id"`)}
	}
	if _, ok := ptc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag_id", err: errors.New(`ent: missing required field "ProductTag.tag_id"`)}
	}
	if _, ok := ptc.mutation.ProductsID(); !ok {
		return &ValidationError{Name: "products", err: errors.New(`ent: missing required edge "ProductTag.products"`)}
	}
	if _, ok := ptc.mutation.TagsID(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`ent: missing required edge "ProductTag.tags"`)}
	}
	return nil
}

func (ptc *ProductTagCreate) sqlSave(ctx context.Context) (*ProductTag, error) {
	if err := ptc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (ptc *ProductTagCreate) createSpec() (*ProductTag, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductTag{config: ptc.config}
		_spec = sqlgraph.NewCreateSpec(producttag.Table, nil)
	)
	if nodes := ptc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   producttag.ProductsTable,
			Columns: []string{producttag.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productinfo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProductID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ptc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   producttag.TagsTable,
			Columns: []string{producttag.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TagID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductTagCreateBulk is the builder for creating many ProductTag entities in bulk.
type ProductTagCreateBulk struct {
	config
	err      error
	builders []*ProductTagCreate
}

// Save creates the ProductTag entities in the database.
func (ptcb *ProductTagCreateBulk) Save(ctx context.Context) ([]*ProductTag, error) {
	if ptcb.err != nil {
		return nil, ptcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ptcb.builders))
	nodes := make([]*ProductTag, len(ptcb.builders))
	mutators := make([]Mutator, len(ptcb.builders))
	for i := range ptcb.builders {
		func(i int, root context.Context) {
			builder := ptcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductTagMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ptcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptcb *ProductTagCreateBulk) SaveX(ctx context.Context) []*ProductTag {
	v, err := ptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptcb *ProductTagCreateBulk) Exec(ctx context.Context) error {
	_, err := ptcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcb *ProductTagCreateBulk) ExecX(ctx context.Context) {
	if err := ptcb.Exec(ctx); err != nil {
		panic(err)
	}
}
