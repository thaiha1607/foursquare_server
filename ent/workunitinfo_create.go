// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/address"
	"github.com/thaiha1607/foursquare_server/ent/workunitinfo"
)

// WorkUnitInfoCreate is the builder for creating a WorkUnitInfo entity.
type WorkUnitInfoCreate struct {
	config
	mutation *WorkUnitInfoMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wuic *WorkUnitInfoCreate) SetName(s string) *WorkUnitInfoCreate {
	wuic.mutation.SetName(s)
	return wuic
}

// SetAddressID sets the "address_id" field.
func (wuic *WorkUnitInfoCreate) SetAddressID(s string) *WorkUnitInfoCreate {
	wuic.mutation.SetAddressID(s)
	return wuic
}

// SetNillableAddressID sets the "address_id" field if the given value is not nil.
func (wuic *WorkUnitInfoCreate) SetNillableAddressID(s *string) *WorkUnitInfoCreate {
	if s != nil {
		wuic.SetAddressID(*s)
	}
	return wuic
}

// SetType sets the "type" field.
func (wuic *WorkUnitInfoCreate) SetType(w workunitinfo.Type) *WorkUnitInfoCreate {
	wuic.mutation.SetType(w)
	return wuic
}

// SetNillableType sets the "type" field if the given value is not nil.
func (wuic *WorkUnitInfoCreate) SetNillableType(w *workunitinfo.Type) *WorkUnitInfoCreate {
	if w != nil {
		wuic.SetType(*w)
	}
	return wuic
}

// SetID sets the "id" field.
func (wuic *WorkUnitInfoCreate) SetID(u uuid.UUID) *WorkUnitInfoCreate {
	wuic.mutation.SetID(u)
	return wuic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wuic *WorkUnitInfoCreate) SetNillableID(u *uuid.UUID) *WorkUnitInfoCreate {
	if u != nil {
		wuic.SetID(*u)
	}
	return wuic
}

// SetAddress sets the "address" edge to the Address entity.
func (wuic *WorkUnitInfoCreate) SetAddress(a *Address) *WorkUnitInfoCreate {
	return wuic.SetAddressID(a.ID)
}

// Mutation returns the WorkUnitInfoMutation object of the builder.
func (wuic *WorkUnitInfoCreate) Mutation() *WorkUnitInfoMutation {
	return wuic.mutation
}

// Save creates the WorkUnitInfo in the database.
func (wuic *WorkUnitInfoCreate) Save(ctx context.Context) (*WorkUnitInfo, error) {
	wuic.defaults()
	return withHooks(ctx, wuic.sqlSave, wuic.mutation, wuic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wuic *WorkUnitInfoCreate) SaveX(ctx context.Context) *WorkUnitInfo {
	v, err := wuic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wuic *WorkUnitInfoCreate) Exec(ctx context.Context) error {
	_, err := wuic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuic *WorkUnitInfoCreate) ExecX(ctx context.Context) {
	if err := wuic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuic *WorkUnitInfoCreate) defaults() {
	if _, ok := wuic.mutation.GetType(); !ok {
		v := workunitinfo.DefaultType
		wuic.mutation.SetType(v)
	}
	if _, ok := wuic.mutation.ID(); !ok {
		v := workunitinfo.DefaultID()
		wuic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuic *WorkUnitInfoCreate) check() error {
	if _, ok := wuic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "WorkUnitInfo.name"`)}
	}
	if v, ok := wuic.mutation.Name(); ok {
		if err := workunitinfo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WorkUnitInfo.name": %w`, err)}
		}
	}
	if _, ok := wuic.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "WorkUnitInfo.type"`)}
	}
	if v, ok := wuic.mutation.GetType(); ok {
		if err := workunitinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "WorkUnitInfo.type": %w`, err)}
		}
	}
	return nil
}

func (wuic *WorkUnitInfoCreate) sqlSave(ctx context.Context) (*WorkUnitInfo, error) {
	if err := wuic.check(); err != nil {
		return nil, err
	}
	_node, _spec := wuic.createSpec()
	if err := sqlgraph.CreateNode(ctx, wuic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	wuic.mutation.id = &_node.ID
	wuic.mutation.done = true
	return _node, nil
}

func (wuic *WorkUnitInfoCreate) createSpec() (*WorkUnitInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkUnitInfo{config: wuic.config}
		_spec = sqlgraph.NewCreateSpec(workunitinfo.Table, sqlgraph.NewFieldSpec(workunitinfo.FieldID, field.TypeUUID))
	)
	if id, ok := wuic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wuic.mutation.Name(); ok {
		_spec.SetField(workunitinfo.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wuic.mutation.GetType(); ok {
		_spec.SetField(workunitinfo.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if nodes := wuic.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workunitinfo.AddressTable,
			Columns: []string{workunitinfo.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AddressID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkUnitInfoCreateBulk is the builder for creating many WorkUnitInfo entities in bulk.
type WorkUnitInfoCreateBulk struct {
	config
	err      error
	builders []*WorkUnitInfoCreate
}

// Save creates the WorkUnitInfo entities in the database.
func (wuicb *WorkUnitInfoCreateBulk) Save(ctx context.Context) ([]*WorkUnitInfo, error) {
	if wuicb.err != nil {
		return nil, wuicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wuicb.builders))
	nodes := make([]*WorkUnitInfo, len(wuicb.builders))
	mutators := make([]Mutator, len(wuicb.builders))
	for i := range wuicb.builders {
		func(i int, root context.Context) {
			builder := wuicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkUnitInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, wuicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wuicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wuicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wuicb *WorkUnitInfoCreateBulk) SaveX(ctx context.Context) []*WorkUnitInfo {
	v, err := wuicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wuicb *WorkUnitInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := wuicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuicb *WorkUnitInfoCreateBulk) ExecX(ctx context.Context) {
	if err := wuicb.Exec(ctx); err != nil {
		panic(err)
	}
}
