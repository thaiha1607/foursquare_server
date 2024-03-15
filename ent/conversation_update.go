// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/conversation"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ConversationUpdate is the builder for updating Conversation entities.
type ConversationUpdate struct {
	config
	hooks    []Hook
	mutation *ConversationMutation
}

// Where appends a list predicates to the ConversationUpdate builder.
func (cu *ConversationUpdate) Where(ps ...predicate.Conversation) *ConversationUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ConversationUpdate) SetUpdatedAt(t time.Time) *ConversationUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetTitle sets the "title" field.
func (cu *ConversationUpdate) SetTitle(s string) *ConversationUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cu *ConversationUpdate) SetNillableTitle(s *string) *ConversationUpdate {
	if s != nil {
		cu.SetTitle(*s)
	}
	return cu
}

// ClearTitle clears the value of the "title" field.
func (cu *ConversationUpdate) ClearTitle() *ConversationUpdate {
	cu.mutation.ClearTitle()
	return cu
}

// Mutation returns the ConversationMutation object of the builder.
func (cu *ConversationUpdate) Mutation() *ConversationMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConversationUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConversationUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConversationUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConversationUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ConversationUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := conversation.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ConversationUpdate) check() error {
	if _, ok := cu.mutation.PersonOneID(); cu.mutation.PersonOneCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Conversation.person_one"`)
	}
	if _, ok := cu.mutation.PersonTwoID(); cu.mutation.PersonTwoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Conversation.person_two"`)
	}
	return nil
}

func (cu *ConversationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(conversation.Table, conversation.Columns, sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(conversation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.SetField(conversation.FieldTitle, field.TypeString, value)
	}
	if cu.mutation.TitleCleared() {
		_spec.ClearField(conversation.FieldTitle, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conversation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConversationUpdateOne is the builder for updating a single Conversation entity.
type ConversationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConversationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ConversationUpdateOne) SetUpdatedAt(t time.Time) *ConversationUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetTitle sets the "title" field.
func (cuo *ConversationUpdateOne) SetTitle(s string) *ConversationUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cuo *ConversationUpdateOne) SetNillableTitle(s *string) *ConversationUpdateOne {
	if s != nil {
		cuo.SetTitle(*s)
	}
	return cuo
}

// ClearTitle clears the value of the "title" field.
func (cuo *ConversationUpdateOne) ClearTitle() *ConversationUpdateOne {
	cuo.mutation.ClearTitle()
	return cuo
}

// Mutation returns the ConversationMutation object of the builder.
func (cuo *ConversationUpdateOne) Mutation() *ConversationMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ConversationUpdate builder.
func (cuo *ConversationUpdateOne) Where(ps ...predicate.Conversation) *ConversationUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConversationUpdateOne) Select(field string, fields ...string) *ConversationUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Conversation entity.
func (cuo *ConversationUpdateOne) Save(ctx context.Context) (*Conversation, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConversationUpdateOne) SaveX(ctx context.Context) *Conversation {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConversationUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConversationUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ConversationUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := conversation.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ConversationUpdateOne) check() error {
	if _, ok := cuo.mutation.PersonOneID(); cuo.mutation.PersonOneCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Conversation.person_one"`)
	}
	if _, ok := cuo.mutation.PersonTwoID(); cuo.mutation.PersonTwoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Conversation.person_two"`)
	}
	return nil
}

func (cuo *ConversationUpdateOne) sqlSave(ctx context.Context) (_node *Conversation, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(conversation.Table, conversation.Columns, sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Conversation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, conversation.FieldID)
		for _, f := range fields {
			if !conversation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != conversation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(conversation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.SetField(conversation.FieldTitle, field.TypeString, value)
	}
	if cuo.mutation.TitleCleared() {
		_spec.ClearField(conversation.FieldTitle, field.TypeString)
	}
	_node = &Conversation{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conversation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
