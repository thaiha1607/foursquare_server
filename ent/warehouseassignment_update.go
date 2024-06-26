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
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/warehouseassignment"
)

// WarehouseAssignmentUpdate is the builder for updating WarehouseAssignment entities.
type WarehouseAssignmentUpdate struct {
	config
	hooks    []Hook
	mutation *WarehouseAssignmentMutation
}

// Where appends a list predicates to the WarehouseAssignmentUpdate builder.
func (wau *WarehouseAssignmentUpdate) Where(ps ...predicate.WarehouseAssignment) *WarehouseAssignmentUpdate {
	wau.mutation.Where(ps...)
	return wau
}

// SetUpdatedAt sets the "updated_at" field.
func (wau *WarehouseAssignmentUpdate) SetUpdatedAt(t time.Time) *WarehouseAssignmentUpdate {
	wau.mutation.SetUpdatedAt(t)
	return wau
}

// SetStaffID sets the "staff_id" field.
func (wau *WarehouseAssignmentUpdate) SetStaffID(u uuid.UUID) *WarehouseAssignmentUpdate {
	wau.mutation.SetStaffID(u)
	return wau
}

// SetNillableStaffID sets the "staff_id" field if the given value is not nil.
func (wau *WarehouseAssignmentUpdate) SetNillableStaffID(u *uuid.UUID) *WarehouseAssignmentUpdate {
	if u != nil {
		wau.SetStaffID(*u)
	}
	return wau
}

// ClearStaffID clears the value of the "staff_id" field.
func (wau *WarehouseAssignmentUpdate) ClearStaffID() *WarehouseAssignmentUpdate {
	wau.mutation.ClearStaffID()
	return wau
}

// SetStatus sets the "status" field.
func (wau *WarehouseAssignmentUpdate) SetStatus(w warehouseassignment.Status) *WarehouseAssignmentUpdate {
	wau.mutation.SetStatus(w)
	return wau
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wau *WarehouseAssignmentUpdate) SetNillableStatus(w *warehouseassignment.Status) *WarehouseAssignmentUpdate {
	if w != nil {
		wau.SetStatus(*w)
	}
	return wau
}

// SetNote sets the "note" field.
func (wau *WarehouseAssignmentUpdate) SetNote(s string) *WarehouseAssignmentUpdate {
	wau.mutation.SetNote(s)
	return wau
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (wau *WarehouseAssignmentUpdate) SetNillableNote(s *string) *WarehouseAssignmentUpdate {
	if s != nil {
		wau.SetNote(*s)
	}
	return wau
}

// ClearNote clears the value of the "note" field.
func (wau *WarehouseAssignmentUpdate) ClearNote() *WarehouseAssignmentUpdate {
	wau.mutation.ClearNote()
	return wau
}

// SetStaff sets the "staff" edge to the Person entity.
func (wau *WarehouseAssignmentUpdate) SetStaff(p *Person) *WarehouseAssignmentUpdate {
	return wau.SetStaffID(p.ID)
}

// Mutation returns the WarehouseAssignmentMutation object of the builder.
func (wau *WarehouseAssignmentUpdate) Mutation() *WarehouseAssignmentMutation {
	return wau.mutation
}

// ClearStaff clears the "staff" edge to the Person entity.
func (wau *WarehouseAssignmentUpdate) ClearStaff() *WarehouseAssignmentUpdate {
	wau.mutation.ClearStaff()
	return wau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wau *WarehouseAssignmentUpdate) Save(ctx context.Context) (int, error) {
	wau.defaults()
	return withHooks(ctx, wau.sqlSave, wau.mutation, wau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wau *WarehouseAssignmentUpdate) SaveX(ctx context.Context) int {
	affected, err := wau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wau *WarehouseAssignmentUpdate) Exec(ctx context.Context) error {
	_, err := wau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wau *WarehouseAssignmentUpdate) ExecX(ctx context.Context) {
	if err := wau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wau *WarehouseAssignmentUpdate) defaults() {
	if _, ok := wau.mutation.UpdatedAt(); !ok {
		v := warehouseassignment.UpdateDefaultUpdatedAt()
		wau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wau *WarehouseAssignmentUpdate) check() error {
	if v, ok := wau.mutation.Status(); ok {
		if err := warehouseassignment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "WarehouseAssignment.status": %w`, err)}
		}
	}
	if _, ok := wau.mutation.OrderID(); wau.mutation.OrderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WarehouseAssignment.order"`)
	}
	if _, ok := wau.mutation.WorkUnitID(); wau.mutation.WorkUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WarehouseAssignment.work_unit"`)
	}
	return nil
}

func (wau *WarehouseAssignmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(warehouseassignment.Table, warehouseassignment.Columns, sqlgraph.NewFieldSpec(warehouseassignment.FieldID, field.TypeUUID))
	if ps := wau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wau.mutation.UpdatedAt(); ok {
		_spec.SetField(warehouseassignment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wau.mutation.Status(); ok {
		_spec.SetField(warehouseassignment.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := wau.mutation.Note(); ok {
		_spec.SetField(warehouseassignment.FieldNote, field.TypeString, value)
	}
	if wau.mutation.NoteCleared() {
		_spec.ClearField(warehouseassignment.FieldNote, field.TypeString)
	}
	if wau.mutation.StaffCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   warehouseassignment.StaffTable,
			Columns: []string{warehouseassignment.StaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wau.mutation.StaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   warehouseassignment.StaffTable,
			Columns: []string{warehouseassignment.StaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{warehouseassignment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wau.mutation.done = true
	return n, nil
}

// WarehouseAssignmentUpdateOne is the builder for updating a single WarehouseAssignment entity.
type WarehouseAssignmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WarehouseAssignmentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (wauo *WarehouseAssignmentUpdateOne) SetUpdatedAt(t time.Time) *WarehouseAssignmentUpdateOne {
	wauo.mutation.SetUpdatedAt(t)
	return wauo
}

// SetStaffID sets the "staff_id" field.
func (wauo *WarehouseAssignmentUpdateOne) SetStaffID(u uuid.UUID) *WarehouseAssignmentUpdateOne {
	wauo.mutation.SetStaffID(u)
	return wauo
}

// SetNillableStaffID sets the "staff_id" field if the given value is not nil.
func (wauo *WarehouseAssignmentUpdateOne) SetNillableStaffID(u *uuid.UUID) *WarehouseAssignmentUpdateOne {
	if u != nil {
		wauo.SetStaffID(*u)
	}
	return wauo
}

// ClearStaffID clears the value of the "staff_id" field.
func (wauo *WarehouseAssignmentUpdateOne) ClearStaffID() *WarehouseAssignmentUpdateOne {
	wauo.mutation.ClearStaffID()
	return wauo
}

// SetStatus sets the "status" field.
func (wauo *WarehouseAssignmentUpdateOne) SetStatus(w warehouseassignment.Status) *WarehouseAssignmentUpdateOne {
	wauo.mutation.SetStatus(w)
	return wauo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wauo *WarehouseAssignmentUpdateOne) SetNillableStatus(w *warehouseassignment.Status) *WarehouseAssignmentUpdateOne {
	if w != nil {
		wauo.SetStatus(*w)
	}
	return wauo
}

// SetNote sets the "note" field.
func (wauo *WarehouseAssignmentUpdateOne) SetNote(s string) *WarehouseAssignmentUpdateOne {
	wauo.mutation.SetNote(s)
	return wauo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (wauo *WarehouseAssignmentUpdateOne) SetNillableNote(s *string) *WarehouseAssignmentUpdateOne {
	if s != nil {
		wauo.SetNote(*s)
	}
	return wauo
}

// ClearNote clears the value of the "note" field.
func (wauo *WarehouseAssignmentUpdateOne) ClearNote() *WarehouseAssignmentUpdateOne {
	wauo.mutation.ClearNote()
	return wauo
}

// SetStaff sets the "staff" edge to the Person entity.
func (wauo *WarehouseAssignmentUpdateOne) SetStaff(p *Person) *WarehouseAssignmentUpdateOne {
	return wauo.SetStaffID(p.ID)
}

// Mutation returns the WarehouseAssignmentMutation object of the builder.
func (wauo *WarehouseAssignmentUpdateOne) Mutation() *WarehouseAssignmentMutation {
	return wauo.mutation
}

// ClearStaff clears the "staff" edge to the Person entity.
func (wauo *WarehouseAssignmentUpdateOne) ClearStaff() *WarehouseAssignmentUpdateOne {
	wauo.mutation.ClearStaff()
	return wauo
}

// Where appends a list predicates to the WarehouseAssignmentUpdate builder.
func (wauo *WarehouseAssignmentUpdateOne) Where(ps ...predicate.WarehouseAssignment) *WarehouseAssignmentUpdateOne {
	wauo.mutation.Where(ps...)
	return wauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wauo *WarehouseAssignmentUpdateOne) Select(field string, fields ...string) *WarehouseAssignmentUpdateOne {
	wauo.fields = append([]string{field}, fields...)
	return wauo
}

// Save executes the query and returns the updated WarehouseAssignment entity.
func (wauo *WarehouseAssignmentUpdateOne) Save(ctx context.Context) (*WarehouseAssignment, error) {
	wauo.defaults()
	return withHooks(ctx, wauo.sqlSave, wauo.mutation, wauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wauo *WarehouseAssignmentUpdateOne) SaveX(ctx context.Context) *WarehouseAssignment {
	node, err := wauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wauo *WarehouseAssignmentUpdateOne) Exec(ctx context.Context) error {
	_, err := wauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wauo *WarehouseAssignmentUpdateOne) ExecX(ctx context.Context) {
	if err := wauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wauo *WarehouseAssignmentUpdateOne) defaults() {
	if _, ok := wauo.mutation.UpdatedAt(); !ok {
		v := warehouseassignment.UpdateDefaultUpdatedAt()
		wauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wauo *WarehouseAssignmentUpdateOne) check() error {
	if v, ok := wauo.mutation.Status(); ok {
		if err := warehouseassignment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "WarehouseAssignment.status": %w`, err)}
		}
	}
	if _, ok := wauo.mutation.OrderID(); wauo.mutation.OrderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WarehouseAssignment.order"`)
	}
	if _, ok := wauo.mutation.WorkUnitID(); wauo.mutation.WorkUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WarehouseAssignment.work_unit"`)
	}
	return nil
}

func (wauo *WarehouseAssignmentUpdateOne) sqlSave(ctx context.Context) (_node *WarehouseAssignment, err error) {
	if err := wauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(warehouseassignment.Table, warehouseassignment.Columns, sqlgraph.NewFieldSpec(warehouseassignment.FieldID, field.TypeUUID))
	id, ok := wauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WarehouseAssignment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, warehouseassignment.FieldID)
		for _, f := range fields {
			if !warehouseassignment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != warehouseassignment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wauo.mutation.UpdatedAt(); ok {
		_spec.SetField(warehouseassignment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wauo.mutation.Status(); ok {
		_spec.SetField(warehouseassignment.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := wauo.mutation.Note(); ok {
		_spec.SetField(warehouseassignment.FieldNote, field.TypeString, value)
	}
	if wauo.mutation.NoteCleared() {
		_spec.ClearField(warehouseassignment.FieldNote, field.TypeString)
	}
	if wauo.mutation.StaffCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   warehouseassignment.StaffTable,
			Columns: []string{warehouseassignment.StaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wauo.mutation.StaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   warehouseassignment.StaffTable,
			Columns: []string{warehouseassignment.StaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WarehouseAssignment{config: wauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{warehouseassignment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wauo.mutation.done = true
	return _node, nil
}
