// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/conversation"
	"github.com/thaiha1607/foursquare_server/ent/user"
)

// ConversationCreate is the builder for creating a Conversation entity.
type ConversationCreate struct {
	config
	mutation *ConversationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *ConversationCreate) SetCreatedAt(t time.Time) *ConversationCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ConversationCreate) SetNillableCreatedAt(t *time.Time) *ConversationCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ConversationCreate) SetUpdatedAt(t time.Time) *ConversationCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ConversationCreate) SetNillableUpdatedAt(t *time.Time) *ConversationCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetTitle sets the "title" field.
func (cc *ConversationCreate) SetTitle(s string) *ConversationCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cc *ConversationCreate) SetNillableTitle(s *string) *ConversationCreate {
	if s != nil {
		cc.SetTitle(*s)
	}
	return cc
}

// SetUserOneID sets the "user_one_id" field.
func (cc *ConversationCreate) SetUserOneID(u uuid.UUID) *ConversationCreate {
	cc.mutation.SetUserOneID(u)
	return cc
}

// SetUserTwoID sets the "user_two_id" field.
func (cc *ConversationCreate) SetUserTwoID(u uuid.UUID) *ConversationCreate {
	cc.mutation.SetUserTwoID(u)
	return cc
}

// SetID sets the "id" field.
func (cc *ConversationCreate) SetID(u uuid.UUID) *ConversationCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ConversationCreate) SetNillableID(u *uuid.UUID) *ConversationCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// SetUserOne sets the "user_one" edge to the User entity.
func (cc *ConversationCreate) SetUserOne(u *User) *ConversationCreate {
	return cc.SetUserOneID(u.ID)
}

// SetUserTwo sets the "user_two" edge to the User entity.
func (cc *ConversationCreate) SetUserTwo(u *User) *ConversationCreate {
	return cc.SetUserTwoID(u.ID)
}

// Mutation returns the ConversationMutation object of the builder.
func (cc *ConversationCreate) Mutation() *ConversationMutation {
	return cc.mutation
}

// Save creates the Conversation in the database.
func (cc *ConversationCreate) Save(ctx context.Context) (*Conversation, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ConversationCreate) SaveX(ctx context.Context) *Conversation {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ConversationCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ConversationCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ConversationCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := conversation.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := conversation.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := conversation.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConversationCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Conversation.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Conversation.updated_at"`)}
	}
	if _, ok := cc.mutation.UserOneID(); !ok {
		return &ValidationError{Name: "user_one_id", err: errors.New(`ent: missing required field "Conversation.user_one_id"`)}
	}
	if _, ok := cc.mutation.UserTwoID(); !ok {
		return &ValidationError{Name: "user_two_id", err: errors.New(`ent: missing required field "Conversation.user_two_id"`)}
	}
	if _, ok := cc.mutation.UserOneID(); !ok {
		return &ValidationError{Name: "user_one", err: errors.New(`ent: missing required edge "Conversation.user_one"`)}
	}
	if _, ok := cc.mutation.UserTwoID(); !ok {
		return &ValidationError{Name: "user_two", err: errors.New(`ent: missing required edge "Conversation.user_two"`)}
	}
	return nil
}

func (cc *ConversationCreate) sqlSave(ctx context.Context) (*Conversation, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ConversationCreate) createSpec() (*Conversation, *sqlgraph.CreateSpec) {
	var (
		_node = &Conversation{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(conversation.Table, sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(conversation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(conversation.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Title(); ok {
		_spec.SetField(conversation.FieldTitle, field.TypeString, value)
		_node.Title = &value
	}
	if nodes := cc.mutation.UserOneIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   conversation.UserOneTable,
			Columns: []string{conversation.UserOneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserOneID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.UserTwoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   conversation.UserTwoTable,
			Columns: []string{conversation.UserTwoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserTwoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ConversationCreateBulk is the builder for creating many Conversation entities in bulk.
type ConversationCreateBulk struct {
	config
	err      error
	builders []*ConversationCreate
}

// Save creates the Conversation entities in the database.
func (ccb *ConversationCreateBulk) Save(ctx context.Context) ([]*Conversation, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Conversation, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConversationMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ConversationCreateBulk) SaveX(ctx context.Context) []*Conversation {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ConversationCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ConversationCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
