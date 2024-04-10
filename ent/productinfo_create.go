// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/productinfo"
	"github.com/thaiha1607/foursquare_server/ent/tag"
)

// ProductInfoCreate is the builder for creating a ProductInfo entity.
type ProductInfoCreate struct {
	config
	mutation *ProductInfoMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pic *ProductInfoCreate) SetCreatedAt(t time.Time) *ProductInfoCreate {
	pic.mutation.SetCreatedAt(t)
	return pic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pic *ProductInfoCreate) SetNillableCreatedAt(t *time.Time) *ProductInfoCreate {
	if t != nil {
		pic.SetCreatedAt(*t)
	}
	return pic
}

// SetUpdatedAt sets the "updated_at" field.
func (pic *ProductInfoCreate) SetUpdatedAt(t time.Time) *ProductInfoCreate {
	pic.mutation.SetUpdatedAt(t)
	return pic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pic *ProductInfoCreate) SetNillableUpdatedAt(t *time.Time) *ProductInfoCreate {
	if t != nil {
		pic.SetUpdatedAt(*t)
	}
	return pic
}

// SetName sets the "name" field.
func (pic *ProductInfoCreate) SetName(s string) *ProductInfoCreate {
	pic.mutation.SetName(s)
	return pic
}

// SetDescription sets the "description" field.
func (pic *ProductInfoCreate) SetDescription(s string) *ProductInfoCreate {
	pic.mutation.SetDescription(s)
	return pic
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pic *ProductInfoCreate) SetNillableDescription(s *string) *ProductInfoCreate {
	if s != nil {
		pic.SetDescription(*s)
	}
	return pic
}

// SetYear sets the "year" field.
func (pic *ProductInfoCreate) SetYear(i int) *ProductInfoCreate {
	pic.mutation.SetYear(i)
	return pic
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (pic *ProductInfoCreate) SetNillableYear(i *int) *ProductInfoCreate {
	if i != nil {
		pic.SetYear(*i)
	}
	return pic
}

// SetProvider sets the "provider" field.
func (pic *ProductInfoCreate) SetProvider(s string) *ProductInfoCreate {
	pic.mutation.SetProvider(s)
	return pic
}

// SetNillableProvider sets the "provider" field if the given value is not nil.
func (pic *ProductInfoCreate) SetNillableProvider(s *string) *ProductInfoCreate {
	if s != nil {
		pic.SetProvider(*s)
	}
	return pic
}

// SetID sets the "id" field.
func (pic *ProductInfoCreate) SetID(s string) *ProductInfoCreate {
	pic.mutation.SetID(s)
	return pic
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (pic *ProductInfoCreate) AddTagIDs(ids ...string) *ProductInfoCreate {
	pic.mutation.AddTagIDs(ids...)
	return pic
}

// AddTags adds the "tags" edges to the Tag entity.
func (pic *ProductInfoCreate) AddTags(t ...*Tag) *ProductInfoCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pic.AddTagIDs(ids...)
}

// Mutation returns the ProductInfoMutation object of the builder.
func (pic *ProductInfoCreate) Mutation() *ProductInfoMutation {
	return pic.mutation
}

// Save creates the ProductInfo in the database.
func (pic *ProductInfoCreate) Save(ctx context.Context) (*ProductInfo, error) {
	if err := pic.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pic.sqlSave, pic.mutation, pic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pic *ProductInfoCreate) SaveX(ctx context.Context) *ProductInfo {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *ProductInfoCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *ProductInfoCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *ProductInfoCreate) defaults() error {
	if _, ok := pic.mutation.CreatedAt(); !ok {
		if productinfo.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized productinfo.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := productinfo.DefaultCreatedAt()
		pic.mutation.SetCreatedAt(v)
	}
	if _, ok := pic.mutation.UpdatedAt(); !ok {
		if productinfo.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized productinfo.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := productinfo.DefaultUpdatedAt()
		pic.mutation.SetUpdatedAt(v)
	}
	if _, ok := pic.mutation.Description(); !ok {
		v := productinfo.DefaultDescription
		pic.mutation.SetDescription(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pic *ProductInfoCreate) check() error {
	if _, ok := pic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProductInfo.created_at"`)}
	}
	if _, ok := pic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProductInfo.updated_at"`)}
	}
	if _, ok := pic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ProductInfo.name"`)}
	}
	if v, ok := pic.mutation.Name(); ok {
		if err := productinfo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ProductInfo.name": %w`, err)}
		}
	}
	if v, ok := pic.mutation.Year(); ok {
		if err := productinfo.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "ProductInfo.year": %w`, err)}
		}
	}
	if v, ok := pic.mutation.ID(); ok {
		if err := productinfo.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "ProductInfo.id": %w`, err)}
		}
	}
	return nil
}

func (pic *ProductInfoCreate) sqlSave(ctx context.Context) (*ProductInfo, error) {
	if err := pic.check(); err != nil {
		return nil, err
	}
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ProductInfo.ID type: %T", _spec.ID.Value)
		}
	}
	pic.mutation.id = &_node.ID
	pic.mutation.done = true
	return _node, nil
}

func (pic *ProductInfoCreate) createSpec() (*ProductInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductInfo{config: pic.config}
		_spec = sqlgraph.NewCreateSpec(productinfo.Table, sqlgraph.NewFieldSpec(productinfo.FieldID, field.TypeString))
	)
	if id, ok := pic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pic.mutation.CreatedAt(); ok {
		_spec.SetField(productinfo.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := pic.mutation.UpdatedAt(); ok {
		_spec.SetField(productinfo.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := pic.mutation.Name(); ok {
		_spec.SetField(productinfo.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pic.mutation.Description(); ok {
		_spec.SetField(productinfo.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := pic.mutation.Year(); ok {
		_spec.SetField(productinfo.FieldYear, field.TypeInt, value)
		_node.Year = &value
	}
	if value, ok := pic.mutation.Provider(); ok {
		_spec.SetField(productinfo.FieldProvider, field.TypeString, value)
		_node.Provider = &value
	}
	if nodes := pic.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   productinfo.TagsTable,
			Columns: productinfo.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductInfoCreateBulk is the builder for creating many ProductInfo entities in bulk.
type ProductInfoCreateBulk struct {
	config
	err      error
	builders []*ProductInfoCreate
}

// Save creates the ProductInfo entities in the database.
func (picb *ProductInfoCreateBulk) Save(ctx context.Context) ([]*ProductInfo, error) {
	if picb.err != nil {
		return nil, picb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*ProductInfo, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *ProductInfoCreateBulk) SaveX(ctx context.Context) []*ProductInfo {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *ProductInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *ProductInfoCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}
