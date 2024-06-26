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
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/address"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/workunitinfo"
)

// WorkUnitInfoUpdate is the builder for updating WorkUnitInfo entities.
type WorkUnitInfoUpdate struct {
	config
	hooks    []Hook
	mutation *WorkUnitInfoMutation
}

// Where appends a list predicates to the WorkUnitInfoUpdate builder.
func (wuiu *WorkUnitInfoUpdate) Where(ps ...predicate.WorkUnitInfo) *WorkUnitInfoUpdate {
	wuiu.mutation.Where(ps...)
	return wuiu
}

// SetUpdatedAt sets the "updated_at" field.
func (wuiu *WorkUnitInfoUpdate) SetUpdatedAt(t time.Time) *WorkUnitInfoUpdate {
	wuiu.mutation.SetUpdatedAt(t)
	return wuiu
}

// SetName sets the "name" field.
func (wuiu *WorkUnitInfoUpdate) SetName(s string) *WorkUnitInfoUpdate {
	wuiu.mutation.SetName(s)
	return wuiu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wuiu *WorkUnitInfoUpdate) SetNillableName(s *string) *WorkUnitInfoUpdate {
	if s != nil {
		wuiu.SetName(*s)
	}
	return wuiu
}

// SetAddressID sets the "address_id" field.
func (wuiu *WorkUnitInfoUpdate) SetAddressID(u uuid.UUID) *WorkUnitInfoUpdate {
	wuiu.mutation.SetAddressID(u)
	return wuiu
}

// SetNillableAddressID sets the "address_id" field if the given value is not nil.
func (wuiu *WorkUnitInfoUpdate) SetNillableAddressID(u *uuid.UUID) *WorkUnitInfoUpdate {
	if u != nil {
		wuiu.SetAddressID(*u)
	}
	return wuiu
}

// ClearAddressID clears the value of the "address_id" field.
func (wuiu *WorkUnitInfoUpdate) ClearAddressID() *WorkUnitInfoUpdate {
	wuiu.mutation.ClearAddressID()
	return wuiu
}

// SetType sets the "type" field.
func (wuiu *WorkUnitInfoUpdate) SetType(w workunitinfo.Type) *WorkUnitInfoUpdate {
	wuiu.mutation.SetType(w)
	return wuiu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (wuiu *WorkUnitInfoUpdate) SetNillableType(w *workunitinfo.Type) *WorkUnitInfoUpdate {
	if w != nil {
		wuiu.SetType(*w)
	}
	return wuiu
}

// SetImageURL sets the "image_url" field.
func (wuiu *WorkUnitInfoUpdate) SetImageURL(s string) *WorkUnitInfoUpdate {
	wuiu.mutation.SetImageURL(s)
	return wuiu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (wuiu *WorkUnitInfoUpdate) SetNillableImageURL(s *string) *WorkUnitInfoUpdate {
	if s != nil {
		wuiu.SetImageURL(*s)
	}
	return wuiu
}

// ClearImageURL clears the value of the "image_url" field.
func (wuiu *WorkUnitInfoUpdate) ClearImageURL() *WorkUnitInfoUpdate {
	wuiu.mutation.ClearImageURL()
	return wuiu
}

// SetAddress sets the "address" edge to the Address entity.
func (wuiu *WorkUnitInfoUpdate) SetAddress(a *Address) *WorkUnitInfoUpdate {
	return wuiu.SetAddressID(a.ID)
}

// Mutation returns the WorkUnitInfoMutation object of the builder.
func (wuiu *WorkUnitInfoUpdate) Mutation() *WorkUnitInfoMutation {
	return wuiu.mutation
}

// ClearAddress clears the "address" edge to the Address entity.
func (wuiu *WorkUnitInfoUpdate) ClearAddress() *WorkUnitInfoUpdate {
	wuiu.mutation.ClearAddress()
	return wuiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wuiu *WorkUnitInfoUpdate) Save(ctx context.Context) (int, error) {
	wuiu.defaults()
	return withHooks(ctx, wuiu.sqlSave, wuiu.mutation, wuiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuiu *WorkUnitInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := wuiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wuiu *WorkUnitInfoUpdate) Exec(ctx context.Context) error {
	_, err := wuiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuiu *WorkUnitInfoUpdate) ExecX(ctx context.Context) {
	if err := wuiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuiu *WorkUnitInfoUpdate) defaults() {
	if _, ok := wuiu.mutation.UpdatedAt(); !ok {
		v := workunitinfo.UpdateDefaultUpdatedAt()
		wuiu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuiu *WorkUnitInfoUpdate) check() error {
	if v, ok := wuiu.mutation.Name(); ok {
		if err := workunitinfo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WorkUnitInfo.name": %w`, err)}
		}
	}
	if v, ok := wuiu.mutation.GetType(); ok {
		if err := workunitinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "WorkUnitInfo.type": %w`, err)}
		}
	}
	if v, ok := wuiu.mutation.ImageURL(); ok {
		if err := workunitinfo.ImageURLValidator(v); err != nil {
			return &ValidationError{Name: "image_url", err: fmt.Errorf(`ent: validator failed for field "WorkUnitInfo.image_url": %w`, err)}
		}
	}
	return nil
}

func (wuiu *WorkUnitInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wuiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workunitinfo.Table, workunitinfo.Columns, sqlgraph.NewFieldSpec(workunitinfo.FieldID, field.TypeUUID))
	if ps := wuiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuiu.mutation.UpdatedAt(); ok {
		_spec.SetField(workunitinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wuiu.mutation.Name(); ok {
		_spec.SetField(workunitinfo.FieldName, field.TypeString, value)
	}
	if value, ok := wuiu.mutation.GetType(); ok {
		_spec.SetField(workunitinfo.FieldType, field.TypeEnum, value)
	}
	if value, ok := wuiu.mutation.ImageURL(); ok {
		_spec.SetField(workunitinfo.FieldImageURL, field.TypeString, value)
	}
	if wuiu.mutation.ImageURLCleared() {
		_spec.ClearField(workunitinfo.FieldImageURL, field.TypeString)
	}
	if wuiu.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workunitinfo.AddressTable,
			Columns: []string{workunitinfo.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuiu.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workunitinfo.AddressTable,
			Columns: []string{workunitinfo.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wuiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workunitinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wuiu.mutation.done = true
	return n, nil
}

// WorkUnitInfoUpdateOne is the builder for updating a single WorkUnitInfo entity.
type WorkUnitInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkUnitInfoMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (wuiuo *WorkUnitInfoUpdateOne) SetUpdatedAt(t time.Time) *WorkUnitInfoUpdateOne {
	wuiuo.mutation.SetUpdatedAt(t)
	return wuiuo
}

// SetName sets the "name" field.
func (wuiuo *WorkUnitInfoUpdateOne) SetName(s string) *WorkUnitInfoUpdateOne {
	wuiuo.mutation.SetName(s)
	return wuiuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wuiuo *WorkUnitInfoUpdateOne) SetNillableName(s *string) *WorkUnitInfoUpdateOne {
	if s != nil {
		wuiuo.SetName(*s)
	}
	return wuiuo
}

// SetAddressID sets the "address_id" field.
func (wuiuo *WorkUnitInfoUpdateOne) SetAddressID(u uuid.UUID) *WorkUnitInfoUpdateOne {
	wuiuo.mutation.SetAddressID(u)
	return wuiuo
}

// SetNillableAddressID sets the "address_id" field if the given value is not nil.
func (wuiuo *WorkUnitInfoUpdateOne) SetNillableAddressID(u *uuid.UUID) *WorkUnitInfoUpdateOne {
	if u != nil {
		wuiuo.SetAddressID(*u)
	}
	return wuiuo
}

// ClearAddressID clears the value of the "address_id" field.
func (wuiuo *WorkUnitInfoUpdateOne) ClearAddressID() *WorkUnitInfoUpdateOne {
	wuiuo.mutation.ClearAddressID()
	return wuiuo
}

// SetType sets the "type" field.
func (wuiuo *WorkUnitInfoUpdateOne) SetType(w workunitinfo.Type) *WorkUnitInfoUpdateOne {
	wuiuo.mutation.SetType(w)
	return wuiuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (wuiuo *WorkUnitInfoUpdateOne) SetNillableType(w *workunitinfo.Type) *WorkUnitInfoUpdateOne {
	if w != nil {
		wuiuo.SetType(*w)
	}
	return wuiuo
}

// SetImageURL sets the "image_url" field.
func (wuiuo *WorkUnitInfoUpdateOne) SetImageURL(s string) *WorkUnitInfoUpdateOne {
	wuiuo.mutation.SetImageURL(s)
	return wuiuo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (wuiuo *WorkUnitInfoUpdateOne) SetNillableImageURL(s *string) *WorkUnitInfoUpdateOne {
	if s != nil {
		wuiuo.SetImageURL(*s)
	}
	return wuiuo
}

// ClearImageURL clears the value of the "image_url" field.
func (wuiuo *WorkUnitInfoUpdateOne) ClearImageURL() *WorkUnitInfoUpdateOne {
	wuiuo.mutation.ClearImageURL()
	return wuiuo
}

// SetAddress sets the "address" edge to the Address entity.
func (wuiuo *WorkUnitInfoUpdateOne) SetAddress(a *Address) *WorkUnitInfoUpdateOne {
	return wuiuo.SetAddressID(a.ID)
}

// Mutation returns the WorkUnitInfoMutation object of the builder.
func (wuiuo *WorkUnitInfoUpdateOne) Mutation() *WorkUnitInfoMutation {
	return wuiuo.mutation
}

// ClearAddress clears the "address" edge to the Address entity.
func (wuiuo *WorkUnitInfoUpdateOne) ClearAddress() *WorkUnitInfoUpdateOne {
	wuiuo.mutation.ClearAddress()
	return wuiuo
}

// Where appends a list predicates to the WorkUnitInfoUpdate builder.
func (wuiuo *WorkUnitInfoUpdateOne) Where(ps ...predicate.WorkUnitInfo) *WorkUnitInfoUpdateOne {
	wuiuo.mutation.Where(ps...)
	return wuiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuiuo *WorkUnitInfoUpdateOne) Select(field string, fields ...string) *WorkUnitInfoUpdateOne {
	wuiuo.fields = append([]string{field}, fields...)
	return wuiuo
}

// Save executes the query and returns the updated WorkUnitInfo entity.
func (wuiuo *WorkUnitInfoUpdateOne) Save(ctx context.Context) (*WorkUnitInfo, error) {
	wuiuo.defaults()
	return withHooks(ctx, wuiuo.sqlSave, wuiuo.mutation, wuiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuiuo *WorkUnitInfoUpdateOne) SaveX(ctx context.Context) *WorkUnitInfo {
	node, err := wuiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuiuo *WorkUnitInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := wuiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuiuo *WorkUnitInfoUpdateOne) ExecX(ctx context.Context) {
	if err := wuiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuiuo *WorkUnitInfoUpdateOne) defaults() {
	if _, ok := wuiuo.mutation.UpdatedAt(); !ok {
		v := workunitinfo.UpdateDefaultUpdatedAt()
		wuiuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuiuo *WorkUnitInfoUpdateOne) check() error {
	if v, ok := wuiuo.mutation.Name(); ok {
		if err := workunitinfo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WorkUnitInfo.name": %w`, err)}
		}
	}
	if v, ok := wuiuo.mutation.GetType(); ok {
		if err := workunitinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "WorkUnitInfo.type": %w`, err)}
		}
	}
	if v, ok := wuiuo.mutation.ImageURL(); ok {
		if err := workunitinfo.ImageURLValidator(v); err != nil {
			return &ValidationError{Name: "image_url", err: fmt.Errorf(`ent: validator failed for field "WorkUnitInfo.image_url": %w`, err)}
		}
	}
	return nil
}

func (wuiuo *WorkUnitInfoUpdateOne) sqlSave(ctx context.Context) (_node *WorkUnitInfo, err error) {
	if err := wuiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workunitinfo.Table, workunitinfo.Columns, sqlgraph.NewFieldSpec(workunitinfo.FieldID, field.TypeUUID))
	id, ok := wuiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WorkUnitInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workunitinfo.FieldID)
		for _, f := range fields {
			if !workunitinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workunitinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(workunitinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wuiuo.mutation.Name(); ok {
		_spec.SetField(workunitinfo.FieldName, field.TypeString, value)
	}
	if value, ok := wuiuo.mutation.GetType(); ok {
		_spec.SetField(workunitinfo.FieldType, field.TypeEnum, value)
	}
	if value, ok := wuiuo.mutation.ImageURL(); ok {
		_spec.SetField(workunitinfo.FieldImageURL, field.TypeString, value)
	}
	if wuiuo.mutation.ImageURLCleared() {
		_spec.ClearField(workunitinfo.FieldImageURL, field.TypeString)
	}
	if wuiuo.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workunitinfo.AddressTable,
			Columns: []string{workunitinfo.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuiuo.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workunitinfo.AddressTable,
			Columns: []string{workunitinfo.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkUnitInfo{config: wuiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workunitinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuiuo.mutation.done = true
	return _node, nil
}
