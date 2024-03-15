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
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// PersonUpdate is the builder for updating Person entities.
type PersonUpdate struct {
	config
	hooks    []Hook
	mutation *PersonMutation
}

// Where appends a list predicates to the PersonUpdate builder.
func (pu *PersonUpdate) Where(ps ...predicate.Person) *PersonUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PersonUpdate) SetUpdatedAt(t time.Time) *PersonUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetAvatarURL sets the "avatar_url" field.
func (pu *PersonUpdate) SetAvatarURL(s string) *PersonUpdate {
	pu.mutation.SetAvatarURL(s)
	return pu
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableAvatarURL(s *string) *PersonUpdate {
	if s != nil {
		pu.SetAvatarURL(*s)
	}
	return pu
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (pu *PersonUpdate) ClearAvatarURL() *PersonUpdate {
	pu.mutation.ClearAvatarURL()
	return pu
}

// SetEmail sets the "email" field.
func (pu *PersonUpdate) SetEmail(s string) *PersonUpdate {
	pu.mutation.SetEmail(s)
	return pu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableEmail(s *string) *PersonUpdate {
	if s != nil {
		pu.SetEmail(*s)
	}
	return pu
}

// ClearEmail clears the value of the "email" field.
func (pu *PersonUpdate) ClearEmail() *PersonUpdate {
	pu.mutation.ClearEmail()
	return pu
}

// SetName sets the "name" field.
func (pu *PersonUpdate) SetName(s string) *PersonUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableName(s *string) *PersonUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetPhone sets the "phone" field.
func (pu *PersonUpdate) SetPhone(s string) *PersonUpdate {
	pu.mutation.SetPhone(s)
	return pu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (pu *PersonUpdate) SetNillablePhone(s *string) *PersonUpdate {
	if s != nil {
		pu.SetPhone(*s)
	}
	return pu
}

// SetRole sets the "role" field.
func (pu *PersonUpdate) SetRole(pe person.Role) *PersonUpdate {
	pu.mutation.SetRole(pe)
	return pu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableRole(pe *person.Role) *PersonUpdate {
	if pe != nil {
		pu.SetRole(*pe)
	}
	return pu
}

// SetPasswordHash sets the "password_hash" field.
func (pu *PersonUpdate) SetPasswordHash(b []byte) *PersonUpdate {
	pu.mutation.SetPasswordHash(b)
	return pu
}

// ClearPasswordHash clears the value of the "password_hash" field.
func (pu *PersonUpdate) ClearPasswordHash() *PersonUpdate {
	pu.mutation.ClearPasswordHash()
	return pu
}

// SetIsEmailVerified sets the "is_email_verified" field.
func (pu *PersonUpdate) SetIsEmailVerified(b bool) *PersonUpdate {
	pu.mutation.SetIsEmailVerified(b)
	return pu
}

// SetNillableIsEmailVerified sets the "is_email_verified" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableIsEmailVerified(b *bool) *PersonUpdate {
	if b != nil {
		pu.SetIsEmailVerified(*b)
	}
	return pu
}

// SetIsPhoneVerified sets the "is_phone_verified" field.
func (pu *PersonUpdate) SetIsPhoneVerified(b bool) *PersonUpdate {
	pu.mutation.SetIsPhoneVerified(b)
	return pu
}

// SetNillableIsPhoneVerified sets the "is_phone_verified" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableIsPhoneVerified(b *bool) *PersonUpdate {
	if b != nil {
		pu.SetIsPhoneVerified(*b)
	}
	return pu
}

// Mutation returns the PersonMutation object of the builder.
func (pu *PersonUpdate) Mutation() *PersonMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PersonUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PersonUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PersonUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PersonUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PersonUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := person.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PersonUpdate) check() error {
	if v, ok := pu.mutation.AvatarURL(); ok {
		if err := person.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf(`ent: validator failed for field "Person.avatar_url": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Email(); ok {
		if err := person.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Person.email": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Name(); ok {
		if err := person.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Person.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Phone(); ok {
		if err := person.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Person.phone": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Role(); ok {
		if err := person.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Person.role": %w`, err)}
		}
	}
	if v, ok := pu.mutation.PasswordHash(); ok {
		if err := person.PasswordHashValidator(v); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf(`ent: validator failed for field "Person.password_hash": %w`, err)}
		}
	}
	return nil
}

func (pu *PersonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(person.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.AvatarURL(); ok {
		_spec.SetField(person.FieldAvatarURL, field.TypeString, value)
	}
	if pu.mutation.AvatarURLCleared() {
		_spec.ClearField(person.FieldAvatarURL, field.TypeString)
	}
	if value, ok := pu.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
	}
	if pu.mutation.EmailCleared() {
		_spec.ClearField(person.FieldEmail, field.TypeString)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(person.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Phone(); ok {
		_spec.SetField(person.FieldPhone, field.TypeString, value)
	}
	if value, ok := pu.mutation.Role(); ok {
		_spec.SetField(person.FieldRole, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.PasswordHash(); ok {
		_spec.SetField(person.FieldPasswordHash, field.TypeBytes, value)
	}
	if pu.mutation.PasswordHashCleared() {
		_spec.ClearField(person.FieldPasswordHash, field.TypeBytes)
	}
	if value, ok := pu.mutation.IsEmailVerified(); ok {
		_spec.SetField(person.FieldIsEmailVerified, field.TypeBool, value)
	}
	if value, ok := pu.mutation.IsPhoneVerified(); ok {
		_spec.SetField(person.FieldIsPhoneVerified, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{person.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PersonUpdateOne is the builder for updating a single Person entity.
type PersonUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PersonMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PersonUpdateOne) SetUpdatedAt(t time.Time) *PersonUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetAvatarURL sets the "avatar_url" field.
func (puo *PersonUpdateOne) SetAvatarURL(s string) *PersonUpdateOne {
	puo.mutation.SetAvatarURL(s)
	return puo
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableAvatarURL(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetAvatarURL(*s)
	}
	return puo
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (puo *PersonUpdateOne) ClearAvatarURL() *PersonUpdateOne {
	puo.mutation.ClearAvatarURL()
	return puo
}

// SetEmail sets the "email" field.
func (puo *PersonUpdateOne) SetEmail(s string) *PersonUpdateOne {
	puo.mutation.SetEmail(s)
	return puo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableEmail(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetEmail(*s)
	}
	return puo
}

// ClearEmail clears the value of the "email" field.
func (puo *PersonUpdateOne) ClearEmail() *PersonUpdateOne {
	puo.mutation.ClearEmail()
	return puo
}

// SetName sets the "name" field.
func (puo *PersonUpdateOne) SetName(s string) *PersonUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableName(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetPhone sets the "phone" field.
func (puo *PersonUpdateOne) SetPhone(s string) *PersonUpdateOne {
	puo.mutation.SetPhone(s)
	return puo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillablePhone(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetPhone(*s)
	}
	return puo
}

// SetRole sets the "role" field.
func (puo *PersonUpdateOne) SetRole(pe person.Role) *PersonUpdateOne {
	puo.mutation.SetRole(pe)
	return puo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableRole(pe *person.Role) *PersonUpdateOne {
	if pe != nil {
		puo.SetRole(*pe)
	}
	return puo
}

// SetPasswordHash sets the "password_hash" field.
func (puo *PersonUpdateOne) SetPasswordHash(b []byte) *PersonUpdateOne {
	puo.mutation.SetPasswordHash(b)
	return puo
}

// ClearPasswordHash clears the value of the "password_hash" field.
func (puo *PersonUpdateOne) ClearPasswordHash() *PersonUpdateOne {
	puo.mutation.ClearPasswordHash()
	return puo
}

// SetIsEmailVerified sets the "is_email_verified" field.
func (puo *PersonUpdateOne) SetIsEmailVerified(b bool) *PersonUpdateOne {
	puo.mutation.SetIsEmailVerified(b)
	return puo
}

// SetNillableIsEmailVerified sets the "is_email_verified" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableIsEmailVerified(b *bool) *PersonUpdateOne {
	if b != nil {
		puo.SetIsEmailVerified(*b)
	}
	return puo
}

// SetIsPhoneVerified sets the "is_phone_verified" field.
func (puo *PersonUpdateOne) SetIsPhoneVerified(b bool) *PersonUpdateOne {
	puo.mutation.SetIsPhoneVerified(b)
	return puo
}

// SetNillableIsPhoneVerified sets the "is_phone_verified" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableIsPhoneVerified(b *bool) *PersonUpdateOne {
	if b != nil {
		puo.SetIsPhoneVerified(*b)
	}
	return puo
}

// Mutation returns the PersonMutation object of the builder.
func (puo *PersonUpdateOne) Mutation() *PersonMutation {
	return puo.mutation
}

// Where appends a list predicates to the PersonUpdate builder.
func (puo *PersonUpdateOne) Where(ps ...predicate.Person) *PersonUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PersonUpdateOne) Select(field string, fields ...string) *PersonUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Person entity.
func (puo *PersonUpdateOne) Save(ctx context.Context) (*Person, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PersonUpdateOne) SaveX(ctx context.Context) *Person {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PersonUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PersonUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PersonUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := person.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PersonUpdateOne) check() error {
	if v, ok := puo.mutation.AvatarURL(); ok {
		if err := person.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf(`ent: validator failed for field "Person.avatar_url": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Email(); ok {
		if err := person.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Person.email": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Name(); ok {
		if err := person.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Person.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Phone(); ok {
		if err := person.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Person.phone": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Role(); ok {
		if err := person.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Person.role": %w`, err)}
		}
	}
	if v, ok := puo.mutation.PasswordHash(); ok {
		if err := person.PasswordHashValidator(v); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf(`ent: validator failed for field "Person.password_hash": %w`, err)}
		}
	}
	return nil
}

func (puo *PersonUpdateOne) sqlSave(ctx context.Context) (_node *Person, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Person.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, person.FieldID)
		for _, f := range fields {
			if !person.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != person.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(person.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.AvatarURL(); ok {
		_spec.SetField(person.FieldAvatarURL, field.TypeString, value)
	}
	if puo.mutation.AvatarURLCleared() {
		_spec.ClearField(person.FieldAvatarURL, field.TypeString)
	}
	if value, ok := puo.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
	}
	if puo.mutation.EmailCleared() {
		_spec.ClearField(person.FieldEmail, field.TypeString)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(person.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Phone(); ok {
		_spec.SetField(person.FieldPhone, field.TypeString, value)
	}
	if value, ok := puo.mutation.Role(); ok {
		_spec.SetField(person.FieldRole, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.PasswordHash(); ok {
		_spec.SetField(person.FieldPasswordHash, field.TypeBytes, value)
	}
	if puo.mutation.PasswordHashCleared() {
		_spec.ClearField(person.FieldPasswordHash, field.TypeBytes)
	}
	if value, ok := puo.mutation.IsEmailVerified(); ok {
		_spec.SetField(person.FieldIsEmailVerified, field.TypeBool, value)
	}
	if value, ok := puo.mutation.IsPhoneVerified(); ok {
		_spec.SetField(person.FieldIsPhoneVerified, field.TypeBool, value)
	}
	_node = &Person{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{person.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
