// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/transactiontype"
)

// TransactionTypeCreate is the builder for creating a TransactionType entity.
type TransactionTypeCreate struct {
	config
	mutation *TransactionTypeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ttc *TransactionTypeCreate) SetCreatedAt(t time.Time) *TransactionTypeCreate {
	ttc.mutation.SetCreatedAt(t)
	return ttc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ttc *TransactionTypeCreate) SetNillableCreatedAt(t *time.Time) *TransactionTypeCreate {
	if t != nil {
		ttc.SetCreatedAt(*t)
	}
	return ttc
}

// SetUpdatedAt sets the "updated_at" field.
func (ttc *TransactionTypeCreate) SetUpdatedAt(t time.Time) *TransactionTypeCreate {
	ttc.mutation.SetUpdatedAt(t)
	return ttc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ttc *TransactionTypeCreate) SetNillableUpdatedAt(t *time.Time) *TransactionTypeCreate {
	if t != nil {
		ttc.SetUpdatedAt(*t)
	}
	return ttc
}

// SetTransactionType sets the "transaction_type" field.
func (ttc *TransactionTypeCreate) SetTransactionType(s string) *TransactionTypeCreate {
	ttc.mutation.SetTransactionType(s)
	return ttc
}

// Mutation returns the TransactionTypeMutation object of the builder.
func (ttc *TransactionTypeCreate) Mutation() *TransactionTypeMutation {
	return ttc.mutation
}

// Save creates the TransactionType in the database.
func (ttc *TransactionTypeCreate) Save(ctx context.Context) (*TransactionType, error) {
	ttc.defaults()
	return withHooks(ctx, ttc.sqlSave, ttc.mutation, ttc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ttc *TransactionTypeCreate) SaveX(ctx context.Context) *TransactionType {
	v, err := ttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttc *TransactionTypeCreate) Exec(ctx context.Context) error {
	_, err := ttc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttc *TransactionTypeCreate) ExecX(ctx context.Context) {
	if err := ttc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttc *TransactionTypeCreate) defaults() {
	if _, ok := ttc.mutation.CreatedAt(); !ok {
		v := transactiontype.DefaultCreatedAt()
		ttc.mutation.SetCreatedAt(v)
	}
	if _, ok := ttc.mutation.UpdatedAt(); !ok {
		v := transactiontype.DefaultUpdatedAt()
		ttc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttc *TransactionTypeCreate) check() error {
	if _, ok := ttc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TransactionType.created_at"`)}
	}
	if _, ok := ttc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TransactionType.updated_at"`)}
	}
	if _, ok := ttc.mutation.TransactionType(); !ok {
		return &ValidationError{Name: "transaction_type", err: errors.New(`ent: missing required field "TransactionType.transaction_type"`)}
	}
	if v, ok := ttc.mutation.TransactionType(); ok {
		if err := transactiontype.TransactionTypeValidator(v); err != nil {
			return &ValidationError{Name: "transaction_type", err: fmt.Errorf(`ent: validator failed for field "TransactionType.transaction_type": %w`, err)}
		}
	}
	return nil
}

func (ttc *TransactionTypeCreate) sqlSave(ctx context.Context) (*TransactionType, error) {
	if err := ttc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ttc.mutation.id = &_node.ID
	ttc.mutation.done = true
	return _node, nil
}

func (ttc *TransactionTypeCreate) createSpec() (*TransactionType, *sqlgraph.CreateSpec) {
	var (
		_node = &TransactionType{config: ttc.config}
		_spec = sqlgraph.NewCreateSpec(transactiontype.Table, sqlgraph.NewFieldSpec(transactiontype.FieldID, field.TypeInt))
	)
	if value, ok := ttc.mutation.CreatedAt(); ok {
		_spec.SetField(transactiontype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ttc.mutation.UpdatedAt(); ok {
		_spec.SetField(transactiontype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ttc.mutation.TransactionType(); ok {
		_spec.SetField(transactiontype.FieldTransactionType, field.TypeString, value)
		_node.TransactionType = value
	}
	return _node, _spec
}

// TransactionTypeCreateBulk is the builder for creating many TransactionType entities in bulk.
type TransactionTypeCreateBulk struct {
	config
	err      error
	builders []*TransactionTypeCreate
}

// Save creates the TransactionType entities in the database.
func (ttcb *TransactionTypeCreateBulk) Save(ctx context.Context) ([]*TransactionType, error) {
	if ttcb.err != nil {
		return nil, ttcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ttcb.builders))
	nodes := make([]*TransactionType, len(ttcb.builders))
	mutators := make([]Mutator, len(ttcb.builders))
	for i := range ttcb.builders {
		func(i int, root context.Context) {
			builder := ttcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, ttcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ttcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, ttcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ttcb *TransactionTypeCreateBulk) SaveX(ctx context.Context) []*TransactionType {
	v, err := ttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttcb *TransactionTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := ttcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttcb *TransactionTypeCreateBulk) ExecX(ctx context.Context) {
	if err := ttcb.Exec(ctx); err != nil {
		panic(err)
	}
}
