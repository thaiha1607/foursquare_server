// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/messagetype"
)

// MessageTypeCreate is the builder for creating a MessageType entity.
type MessageTypeCreate struct {
	config
	mutation *MessageTypeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mtc *MessageTypeCreate) SetCreatedAt(t time.Time) *MessageTypeCreate {
	mtc.mutation.SetCreatedAt(t)
	return mtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mtc *MessageTypeCreate) SetNillableCreatedAt(t *time.Time) *MessageTypeCreate {
	if t != nil {
		mtc.SetCreatedAt(*t)
	}
	return mtc
}

// SetUpdatedAt sets the "updated_at" field.
func (mtc *MessageTypeCreate) SetUpdatedAt(t time.Time) *MessageTypeCreate {
	mtc.mutation.SetUpdatedAt(t)
	return mtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mtc *MessageTypeCreate) SetNillableUpdatedAt(t *time.Time) *MessageTypeCreate {
	if t != nil {
		mtc.SetUpdatedAt(*t)
	}
	return mtc
}

// SetMessageType sets the "message_type" field.
func (mtc *MessageTypeCreate) SetMessageType(s string) *MessageTypeCreate {
	mtc.mutation.SetMessageType(s)
	return mtc
}

// Mutation returns the MessageTypeMutation object of the builder.
func (mtc *MessageTypeCreate) Mutation() *MessageTypeMutation {
	return mtc.mutation
}

// Save creates the MessageType in the database.
func (mtc *MessageTypeCreate) Save(ctx context.Context) (*MessageType, error) {
	mtc.defaults()
	return withHooks(ctx, mtc.sqlSave, mtc.mutation, mtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mtc *MessageTypeCreate) SaveX(ctx context.Context) *MessageType {
	v, err := mtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtc *MessageTypeCreate) Exec(ctx context.Context) error {
	_, err := mtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtc *MessageTypeCreate) ExecX(ctx context.Context) {
	if err := mtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mtc *MessageTypeCreate) defaults() {
	if _, ok := mtc.mutation.CreatedAt(); !ok {
		v := messagetype.DefaultCreatedAt()
		mtc.mutation.SetCreatedAt(v)
	}
	if _, ok := mtc.mutation.UpdatedAt(); !ok {
		v := messagetype.DefaultUpdatedAt()
		mtc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mtc *MessageTypeCreate) check() error {
	if _, ok := mtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MessageType.created_at"`)}
	}
	if _, ok := mtc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MessageType.updated_at"`)}
	}
	if _, ok := mtc.mutation.MessageType(); !ok {
		return &ValidationError{Name: "message_type", err: errors.New(`ent: missing required field "MessageType.message_type"`)}
	}
	if v, ok := mtc.mutation.MessageType(); ok {
		if err := messagetype.MessageTypeValidator(v); err != nil {
			return &ValidationError{Name: "message_type", err: fmt.Errorf(`ent: validator failed for field "MessageType.message_type": %w`, err)}
		}
	}
	return nil
}

func (mtc *MessageTypeCreate) sqlSave(ctx context.Context) (*MessageType, error) {
	if err := mtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mtc.mutation.id = &_node.ID
	mtc.mutation.done = true
	return _node, nil
}

func (mtc *MessageTypeCreate) createSpec() (*MessageType, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageType{config: mtc.config}
		_spec = sqlgraph.NewCreateSpec(messagetype.Table, sqlgraph.NewFieldSpec(messagetype.FieldID, field.TypeInt))
	)
	if value, ok := mtc.mutation.CreatedAt(); ok {
		_spec.SetField(messagetype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mtc.mutation.UpdatedAt(); ok {
		_spec.SetField(messagetype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mtc.mutation.MessageType(); ok {
		_spec.SetField(messagetype.FieldMessageType, field.TypeString, value)
		_node.MessageType = value
	}
	return _node, _spec
}

// MessageTypeCreateBulk is the builder for creating many MessageType entities in bulk.
type MessageTypeCreateBulk struct {
	config
	err      error
	builders []*MessageTypeCreate
}

// Save creates the MessageType entities in the database.
func (mtcb *MessageTypeCreateBulk) Save(ctx context.Context) ([]*MessageType, error) {
	if mtcb.err != nil {
		return nil, mtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mtcb.builders))
	nodes := make([]*MessageType, len(mtcb.builders))
	mutators := make([]Mutator, len(mtcb.builders))
	for i := range mtcb.builders {
		func(i int, root context.Context) {
			builder := mtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, mtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mtcb *MessageTypeCreateBulk) SaveX(ctx context.Context) []*MessageType {
	v, err := mtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtcb *MessageTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := mtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtcb *MessageTypeCreateBulk) ExecX(ctx context.Context) {
	if err := mtcb.Exec(ctx); err != nil {
		panic(err)
	}
}