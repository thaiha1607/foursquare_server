// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/orderstatuscode"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/shipmenthistory"
)

// ShipmentHistoryUpdate is the builder for updating ShipmentHistory entities.
type ShipmentHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *ShipmentHistoryMutation
}

// Where appends a list predicates to the ShipmentHistoryUpdate builder.
func (shu *ShipmentHistoryUpdate) Where(ps ...predicate.ShipmentHistory) *ShipmentHistoryUpdate {
	shu.mutation.Where(ps...)
	return shu
}

// SetPrevStatusCode sets the "prev_status_code" field.
func (shu *ShipmentHistoryUpdate) SetPrevStatusCode(i int) *ShipmentHistoryUpdate {
	shu.mutation.SetPrevStatusCode(i)
	return shu
}

// SetNillablePrevStatusCode sets the "prev_status_code" field if the given value is not nil.
func (shu *ShipmentHistoryUpdate) SetNillablePrevStatusCode(i *int) *ShipmentHistoryUpdate {
	if i != nil {
		shu.SetPrevStatusCode(*i)
	}
	return shu
}

// ClearPrevStatusCode clears the value of the "prev_status_code" field.
func (shu *ShipmentHistoryUpdate) ClearPrevStatusCode() *ShipmentHistoryUpdate {
	shu.mutation.ClearPrevStatusCode()
	return shu
}

// SetNewStatusCode sets the "new_status_code" field.
func (shu *ShipmentHistoryUpdate) SetNewStatusCode(i int) *ShipmentHistoryUpdate {
	shu.mutation.SetNewStatusCode(i)
	return shu
}

// SetNillableNewStatusCode sets the "new_status_code" field if the given value is not nil.
func (shu *ShipmentHistoryUpdate) SetNillableNewStatusCode(i *int) *ShipmentHistoryUpdate {
	if i != nil {
		shu.SetNewStatusCode(*i)
	}
	return shu
}

// ClearNewStatusCode clears the value of the "new_status_code" field.
func (shu *ShipmentHistoryUpdate) ClearNewStatusCode() *ShipmentHistoryUpdate {
	shu.mutation.ClearNewStatusCode()
	return shu
}

// SetDescription sets the "description" field.
func (shu *ShipmentHistoryUpdate) SetDescription(s string) *ShipmentHistoryUpdate {
	shu.mutation.SetDescription(s)
	return shu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (shu *ShipmentHistoryUpdate) SetNillableDescription(s *string) *ShipmentHistoryUpdate {
	if s != nil {
		shu.SetDescription(*s)
	}
	return shu
}

// ClearDescription clears the value of the "description" field.
func (shu *ShipmentHistoryUpdate) ClearDescription() *ShipmentHistoryUpdate {
	shu.mutation.ClearDescription()
	return shu
}

// SetPrevStatusID sets the "prev_status" edge to the OrderStatusCode entity by ID.
func (shu *ShipmentHistoryUpdate) SetPrevStatusID(id int) *ShipmentHistoryUpdate {
	shu.mutation.SetPrevStatusID(id)
	return shu
}

// SetNillablePrevStatusID sets the "prev_status" edge to the OrderStatusCode entity by ID if the given value is not nil.
func (shu *ShipmentHistoryUpdate) SetNillablePrevStatusID(id *int) *ShipmentHistoryUpdate {
	if id != nil {
		shu = shu.SetPrevStatusID(*id)
	}
	return shu
}

// SetPrevStatus sets the "prev_status" edge to the OrderStatusCode entity.
func (shu *ShipmentHistoryUpdate) SetPrevStatus(o *OrderStatusCode) *ShipmentHistoryUpdate {
	return shu.SetPrevStatusID(o.ID)
}

// SetNewStatusID sets the "new_status" edge to the OrderStatusCode entity by ID.
func (shu *ShipmentHistoryUpdate) SetNewStatusID(id int) *ShipmentHistoryUpdate {
	shu.mutation.SetNewStatusID(id)
	return shu
}

// SetNillableNewStatusID sets the "new_status" edge to the OrderStatusCode entity by ID if the given value is not nil.
func (shu *ShipmentHistoryUpdate) SetNillableNewStatusID(id *int) *ShipmentHistoryUpdate {
	if id != nil {
		shu = shu.SetNewStatusID(*id)
	}
	return shu
}

// SetNewStatus sets the "new_status" edge to the OrderStatusCode entity.
func (shu *ShipmentHistoryUpdate) SetNewStatus(o *OrderStatusCode) *ShipmentHistoryUpdate {
	return shu.SetNewStatusID(o.ID)
}

// Mutation returns the ShipmentHistoryMutation object of the builder.
func (shu *ShipmentHistoryUpdate) Mutation() *ShipmentHistoryMutation {
	return shu.mutation
}

// ClearPrevStatus clears the "prev_status" edge to the OrderStatusCode entity.
func (shu *ShipmentHistoryUpdate) ClearPrevStatus() *ShipmentHistoryUpdate {
	shu.mutation.ClearPrevStatus()
	return shu
}

// ClearNewStatus clears the "new_status" edge to the OrderStatusCode entity.
func (shu *ShipmentHistoryUpdate) ClearNewStatus() *ShipmentHistoryUpdate {
	shu.mutation.ClearNewStatus()
	return shu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (shu *ShipmentHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, shu.sqlSave, shu.mutation, shu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (shu *ShipmentHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := shu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (shu *ShipmentHistoryUpdate) Exec(ctx context.Context) error {
	_, err := shu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (shu *ShipmentHistoryUpdate) ExecX(ctx context.Context) {
	if err := shu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (shu *ShipmentHistoryUpdate) check() error {
	if _, ok := shu.mutation.ShipmentID(); shu.mutation.ShipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ShipmentHistory.shipment"`)
	}
	if _, ok := shu.mutation.PersonID(); shu.mutation.PersonCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ShipmentHistory.person"`)
	}
	return nil
}

func (shu *ShipmentHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := shu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(shipmenthistory.Table, shipmenthistory.Columns, sqlgraph.NewFieldSpec(shipmenthistory.FieldID, field.TypeUUID))
	if ps := shu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := shu.mutation.Description(); ok {
		_spec.SetField(shipmenthistory.FieldDescription, field.TypeString, value)
	}
	if shu.mutation.DescriptionCleared() {
		_spec.ClearField(shipmenthistory.FieldDescription, field.TypeString)
	}
	if shu.mutation.PrevStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.PrevStatusTable,
			Columns: []string{shipmenthistory.PrevStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := shu.mutation.PrevStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.PrevStatusTable,
			Columns: []string{shipmenthistory.PrevStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if shu.mutation.NewStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.NewStatusTable,
			Columns: []string{shipmenthistory.NewStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := shu.mutation.NewStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.NewStatusTable,
			Columns: []string{shipmenthistory.NewStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, shu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shipmenthistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	shu.mutation.done = true
	return n, nil
}

// ShipmentHistoryUpdateOne is the builder for updating a single ShipmentHistory entity.
type ShipmentHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShipmentHistoryMutation
}

// SetPrevStatusCode sets the "prev_status_code" field.
func (shuo *ShipmentHistoryUpdateOne) SetPrevStatusCode(i int) *ShipmentHistoryUpdateOne {
	shuo.mutation.SetPrevStatusCode(i)
	return shuo
}

// SetNillablePrevStatusCode sets the "prev_status_code" field if the given value is not nil.
func (shuo *ShipmentHistoryUpdateOne) SetNillablePrevStatusCode(i *int) *ShipmentHistoryUpdateOne {
	if i != nil {
		shuo.SetPrevStatusCode(*i)
	}
	return shuo
}

// ClearPrevStatusCode clears the value of the "prev_status_code" field.
func (shuo *ShipmentHistoryUpdateOne) ClearPrevStatusCode() *ShipmentHistoryUpdateOne {
	shuo.mutation.ClearPrevStatusCode()
	return shuo
}

// SetNewStatusCode sets the "new_status_code" field.
func (shuo *ShipmentHistoryUpdateOne) SetNewStatusCode(i int) *ShipmentHistoryUpdateOne {
	shuo.mutation.SetNewStatusCode(i)
	return shuo
}

// SetNillableNewStatusCode sets the "new_status_code" field if the given value is not nil.
func (shuo *ShipmentHistoryUpdateOne) SetNillableNewStatusCode(i *int) *ShipmentHistoryUpdateOne {
	if i != nil {
		shuo.SetNewStatusCode(*i)
	}
	return shuo
}

// ClearNewStatusCode clears the value of the "new_status_code" field.
func (shuo *ShipmentHistoryUpdateOne) ClearNewStatusCode() *ShipmentHistoryUpdateOne {
	shuo.mutation.ClearNewStatusCode()
	return shuo
}

// SetDescription sets the "description" field.
func (shuo *ShipmentHistoryUpdateOne) SetDescription(s string) *ShipmentHistoryUpdateOne {
	shuo.mutation.SetDescription(s)
	return shuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (shuo *ShipmentHistoryUpdateOne) SetNillableDescription(s *string) *ShipmentHistoryUpdateOne {
	if s != nil {
		shuo.SetDescription(*s)
	}
	return shuo
}

// ClearDescription clears the value of the "description" field.
func (shuo *ShipmentHistoryUpdateOne) ClearDescription() *ShipmentHistoryUpdateOne {
	shuo.mutation.ClearDescription()
	return shuo
}

// SetPrevStatusID sets the "prev_status" edge to the OrderStatusCode entity by ID.
func (shuo *ShipmentHistoryUpdateOne) SetPrevStatusID(id int) *ShipmentHistoryUpdateOne {
	shuo.mutation.SetPrevStatusID(id)
	return shuo
}

// SetNillablePrevStatusID sets the "prev_status" edge to the OrderStatusCode entity by ID if the given value is not nil.
func (shuo *ShipmentHistoryUpdateOne) SetNillablePrevStatusID(id *int) *ShipmentHistoryUpdateOne {
	if id != nil {
		shuo = shuo.SetPrevStatusID(*id)
	}
	return shuo
}

// SetPrevStatus sets the "prev_status" edge to the OrderStatusCode entity.
func (shuo *ShipmentHistoryUpdateOne) SetPrevStatus(o *OrderStatusCode) *ShipmentHistoryUpdateOne {
	return shuo.SetPrevStatusID(o.ID)
}

// SetNewStatusID sets the "new_status" edge to the OrderStatusCode entity by ID.
func (shuo *ShipmentHistoryUpdateOne) SetNewStatusID(id int) *ShipmentHistoryUpdateOne {
	shuo.mutation.SetNewStatusID(id)
	return shuo
}

// SetNillableNewStatusID sets the "new_status" edge to the OrderStatusCode entity by ID if the given value is not nil.
func (shuo *ShipmentHistoryUpdateOne) SetNillableNewStatusID(id *int) *ShipmentHistoryUpdateOne {
	if id != nil {
		shuo = shuo.SetNewStatusID(*id)
	}
	return shuo
}

// SetNewStatus sets the "new_status" edge to the OrderStatusCode entity.
func (shuo *ShipmentHistoryUpdateOne) SetNewStatus(o *OrderStatusCode) *ShipmentHistoryUpdateOne {
	return shuo.SetNewStatusID(o.ID)
}

// Mutation returns the ShipmentHistoryMutation object of the builder.
func (shuo *ShipmentHistoryUpdateOne) Mutation() *ShipmentHistoryMutation {
	return shuo.mutation
}

// ClearPrevStatus clears the "prev_status" edge to the OrderStatusCode entity.
func (shuo *ShipmentHistoryUpdateOne) ClearPrevStatus() *ShipmentHistoryUpdateOne {
	shuo.mutation.ClearPrevStatus()
	return shuo
}

// ClearNewStatus clears the "new_status" edge to the OrderStatusCode entity.
func (shuo *ShipmentHistoryUpdateOne) ClearNewStatus() *ShipmentHistoryUpdateOne {
	shuo.mutation.ClearNewStatus()
	return shuo
}

// Where appends a list predicates to the ShipmentHistoryUpdate builder.
func (shuo *ShipmentHistoryUpdateOne) Where(ps ...predicate.ShipmentHistory) *ShipmentHistoryUpdateOne {
	shuo.mutation.Where(ps...)
	return shuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (shuo *ShipmentHistoryUpdateOne) Select(field string, fields ...string) *ShipmentHistoryUpdateOne {
	shuo.fields = append([]string{field}, fields...)
	return shuo
}

// Save executes the query and returns the updated ShipmentHistory entity.
func (shuo *ShipmentHistoryUpdateOne) Save(ctx context.Context) (*ShipmentHistory, error) {
	return withHooks(ctx, shuo.sqlSave, shuo.mutation, shuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (shuo *ShipmentHistoryUpdateOne) SaveX(ctx context.Context) *ShipmentHistory {
	node, err := shuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (shuo *ShipmentHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := shuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (shuo *ShipmentHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := shuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (shuo *ShipmentHistoryUpdateOne) check() error {
	if _, ok := shuo.mutation.ShipmentID(); shuo.mutation.ShipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ShipmentHistory.shipment"`)
	}
	if _, ok := shuo.mutation.PersonID(); shuo.mutation.PersonCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ShipmentHistory.person"`)
	}
	return nil
}

func (shuo *ShipmentHistoryUpdateOne) sqlSave(ctx context.Context) (_node *ShipmentHistory, err error) {
	if err := shuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(shipmenthistory.Table, shipmenthistory.Columns, sqlgraph.NewFieldSpec(shipmenthistory.FieldID, field.TypeUUID))
	id, ok := shuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ShipmentHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := shuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shipmenthistory.FieldID)
		for _, f := range fields {
			if !shipmenthistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != shipmenthistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := shuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := shuo.mutation.Description(); ok {
		_spec.SetField(shipmenthistory.FieldDescription, field.TypeString, value)
	}
	if shuo.mutation.DescriptionCleared() {
		_spec.ClearField(shipmenthistory.FieldDescription, field.TypeString)
	}
	if shuo.mutation.PrevStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.PrevStatusTable,
			Columns: []string{shipmenthistory.PrevStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := shuo.mutation.PrevStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.PrevStatusTable,
			Columns: []string{shipmenthistory.PrevStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if shuo.mutation.NewStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.NewStatusTable,
			Columns: []string{shipmenthistory.NewStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := shuo.mutation.NewStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenthistory.NewStatusTable,
			Columns: []string{shipmenthistory.NewStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ShipmentHistory{config: shuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, shuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shipmenthistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	shuo.mutation.done = true
	return _node, nil
}
