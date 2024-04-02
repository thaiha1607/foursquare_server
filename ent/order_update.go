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
	"github.com/thaiha1607/foursquare_server/ent/order"
	"github.com/thaiha1607/foursquare_server/ent/orderstatuscode"
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrderUpdate) SetUpdatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetNote sets the "note" field.
func (ou *OrderUpdate) SetNote(s string) *OrderUpdate {
	ou.mutation.SetNote(s)
	return ou
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableNote(s *string) *OrderUpdate {
	if s != nil {
		ou.SetNote(*s)
	}
	return ou
}

// ClearNote clears the value of the "note" field.
func (ou *OrderUpdate) ClearNote() *OrderUpdate {
	ou.mutation.ClearNote()
	return ou
}

// SetParentOrderID sets the "parent_order_id" field.
func (ou *OrderUpdate) SetParentOrderID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetParentOrderID(u)
	return ou
}

// SetNillableParentOrderID sets the "parent_order_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableParentOrderID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetParentOrderID(*u)
	}
	return ou
}

// ClearParentOrderID clears the value of the "parent_order_id" field.
func (ou *OrderUpdate) ClearParentOrderID() *OrderUpdate {
	ou.mutation.ClearParentOrderID()
	return ou
}

// SetPriority sets the "priority" field.
func (ou *OrderUpdate) SetPriority(i int) *OrderUpdate {
	ou.mutation.ResetPriority()
	ou.mutation.SetPriority(i)
	return ou
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (ou *OrderUpdate) SetNillablePriority(i *int) *OrderUpdate {
	if i != nil {
		ou.SetPriority(*i)
	}
	return ou
}

// AddPriority adds i to the "priority" field.
func (ou *OrderUpdate) AddPriority(i int) *OrderUpdate {
	ou.mutation.AddPriority(i)
	return ou
}

// SetStatusCode sets the "status_code" field.
func (ou *OrderUpdate) SetStatusCode(i int) *OrderUpdate {
	ou.mutation.SetStatusCode(i)
	return ou
}

// SetNillableStatusCode sets the "status_code" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableStatusCode(i *int) *OrderUpdate {
	if i != nil {
		ou.SetStatusCode(*i)
	}
	return ou
}

// SetStaffID sets the "staff_id" field.
func (ou *OrderUpdate) SetStaffID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetStaffID(u)
	return ou
}

// SetNillableStaffID sets the "staff_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableStaffID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetStaffID(*u)
	}
	return ou
}

// SetInternalNote sets the "internal_note" field.
func (ou *OrderUpdate) SetInternalNote(s string) *OrderUpdate {
	ou.mutation.SetInternalNote(s)
	return ou
}

// SetNillableInternalNote sets the "internal_note" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableInternalNote(s *string) *OrderUpdate {
	if s != nil {
		ou.SetInternalNote(*s)
	}
	return ou
}

// ClearInternalNote clears the value of the "internal_note" field.
func (ou *OrderUpdate) ClearInternalNote() *OrderUpdate {
	ou.mutation.ClearInternalNote()
	return ou
}

// SetIsInternal sets the "is_internal" field.
func (ou *OrderUpdate) SetIsInternal(b bool) *OrderUpdate {
	ou.mutation.SetIsInternal(b)
	return ou
}

// SetNillableIsInternal sets the "is_internal" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableIsInternal(b *bool) *OrderUpdate {
	if b != nil {
		ou.SetIsInternal(*b)
	}
	return ou
}

// SetAddressID sets the "address_id" field.
func (ou *OrderUpdate) SetAddressID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetAddressID(u)
	return ou
}

// SetNillableAddressID sets the "address_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableAddressID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetAddressID(*u)
	}
	return ou
}

// SetParentOrder sets the "parent_order" edge to the Order entity.
func (ou *OrderUpdate) SetParentOrder(o *Order) *OrderUpdate {
	return ou.SetParentOrderID(o.ID)
}

// SetOrderStatusID sets the "order_status" edge to the OrderStatusCode entity by ID.
func (ou *OrderUpdate) SetOrderStatusID(id int) *OrderUpdate {
	ou.mutation.SetOrderStatusID(id)
	return ou
}

// SetOrderStatus sets the "order_status" edge to the OrderStatusCode entity.
func (ou *OrderUpdate) SetOrderStatus(o *OrderStatusCode) *OrderUpdate {
	return ou.SetOrderStatusID(o.ID)
}

// SetStaff sets the "staff" edge to the Person entity.
func (ou *OrderUpdate) SetStaff(p *Person) *OrderUpdate {
	return ou.SetStaffID(p.ID)
}

// SetOrderAddressID sets the "order_address" edge to the Address entity by ID.
func (ou *OrderUpdate) SetOrderAddressID(id uuid.UUID) *OrderUpdate {
	ou.mutation.SetOrderAddressID(id)
	return ou
}

// SetOrderAddress sets the "order_address" edge to the Address entity.
func (ou *OrderUpdate) SetOrderAddress(a *Address) *OrderUpdate {
	return ou.SetOrderAddressID(a.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// ClearParentOrder clears the "parent_order" edge to the Order entity.
func (ou *OrderUpdate) ClearParentOrder() *OrderUpdate {
	ou.mutation.ClearParentOrder()
	return ou
}

// ClearOrderStatus clears the "order_status" edge to the OrderStatusCode entity.
func (ou *OrderUpdate) ClearOrderStatus() *OrderUpdate {
	ou.mutation.ClearOrderStatus()
	return ou
}

// ClearStaff clears the "staff" edge to the Person entity.
func (ou *OrderUpdate) ClearStaff() *OrderUpdate {
	ou.mutation.ClearStaff()
	return ou
}

// ClearOrderAddress clears the "order_address" edge to the Address entity.
func (ou *OrderUpdate) ClearOrderAddress() *OrderUpdate {
	ou.mutation.ClearOrderAddress()
	return ou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	ou.defaults()
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.Priority(); ok {
		if err := order.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Order.priority": %w`, err)}
		}
	}
	if _, ok := ou.mutation.CustomerID(); ou.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.customer"`)
	}
	if _, ok := ou.mutation.CreatorID(); ou.mutation.CreatorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.creator"`)
	}
	if _, ok := ou.mutation.OrderStatusID(); ou.mutation.OrderStatusCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.order_status"`)
	}
	if _, ok := ou.mutation.StaffID(); ou.mutation.StaffCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.staff"`)
	}
	if _, ok := ou.mutation.OrderAddressID(); ou.mutation.OrderAddressCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.order_address"`)
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.Note(); ok {
		_spec.SetField(order.FieldNote, field.TypeString, value)
	}
	if ou.mutation.NoteCleared() {
		_spec.ClearField(order.FieldNote, field.TypeString)
	}
	if value, ok := ou.mutation.Priority(); ok {
		_spec.SetField(order.FieldPriority, field.TypeInt, value)
	}
	if value, ok := ou.mutation.AddedPriority(); ok {
		_spec.AddField(order.FieldPriority, field.TypeInt, value)
	}
	if value, ok := ou.mutation.InternalNote(); ok {
		_spec.SetField(order.FieldInternalNote, field.TypeString, value)
	}
	if ou.mutation.InternalNoteCleared() {
		_spec.ClearField(order.FieldInternalNote, field.TypeString)
	}
	if value, ok := ou.mutation.IsInternal(); ok {
		_spec.SetField(order.FieldIsInternal, field.TypeBool, value)
	}
	if ou.mutation.ParentOrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   order.ParentOrderTable,
			Columns: []string{order.ParentOrderColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ParentOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   order.ParentOrderTable,
			Columns: []string{order.ParentOrderColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.OrderStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.OrderStatusTable,
			Columns: []string{order.OrderStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.OrderStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.OrderStatusTable,
			Columns: []string{order.OrderStatusColumn},
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
	if ou.mutation.StaffCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.StaffTable,
			Columns: []string{order.StaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.StaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.StaffTable,
			Columns: []string{order.StaffColumn},
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
	if ou.mutation.OrderAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.OrderAddressTable,
			Columns: []string{order.OrderAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.OrderAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.OrderAddressTable,
			Columns: []string{order.OrderAddressColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrderUpdateOne) SetUpdatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetNote sets the "note" field.
func (ouo *OrderUpdateOne) SetNote(s string) *OrderUpdateOne {
	ouo.mutation.SetNote(s)
	return ouo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableNote(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetNote(*s)
	}
	return ouo
}

// ClearNote clears the value of the "note" field.
func (ouo *OrderUpdateOne) ClearNote() *OrderUpdateOne {
	ouo.mutation.ClearNote()
	return ouo
}

// SetParentOrderID sets the "parent_order_id" field.
func (ouo *OrderUpdateOne) SetParentOrderID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetParentOrderID(u)
	return ouo
}

// SetNillableParentOrderID sets the "parent_order_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableParentOrderID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetParentOrderID(*u)
	}
	return ouo
}

// ClearParentOrderID clears the value of the "parent_order_id" field.
func (ouo *OrderUpdateOne) ClearParentOrderID() *OrderUpdateOne {
	ouo.mutation.ClearParentOrderID()
	return ouo
}

// SetPriority sets the "priority" field.
func (ouo *OrderUpdateOne) SetPriority(i int) *OrderUpdateOne {
	ouo.mutation.ResetPriority()
	ouo.mutation.SetPriority(i)
	return ouo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillablePriority(i *int) *OrderUpdateOne {
	if i != nil {
		ouo.SetPriority(*i)
	}
	return ouo
}

// AddPriority adds i to the "priority" field.
func (ouo *OrderUpdateOne) AddPriority(i int) *OrderUpdateOne {
	ouo.mutation.AddPriority(i)
	return ouo
}

// SetStatusCode sets the "status_code" field.
func (ouo *OrderUpdateOne) SetStatusCode(i int) *OrderUpdateOne {
	ouo.mutation.SetStatusCode(i)
	return ouo
}

// SetNillableStatusCode sets the "status_code" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableStatusCode(i *int) *OrderUpdateOne {
	if i != nil {
		ouo.SetStatusCode(*i)
	}
	return ouo
}

// SetStaffID sets the "staff_id" field.
func (ouo *OrderUpdateOne) SetStaffID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetStaffID(u)
	return ouo
}

// SetNillableStaffID sets the "staff_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableStaffID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetStaffID(*u)
	}
	return ouo
}

// SetInternalNote sets the "internal_note" field.
func (ouo *OrderUpdateOne) SetInternalNote(s string) *OrderUpdateOne {
	ouo.mutation.SetInternalNote(s)
	return ouo
}

// SetNillableInternalNote sets the "internal_note" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableInternalNote(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetInternalNote(*s)
	}
	return ouo
}

// ClearInternalNote clears the value of the "internal_note" field.
func (ouo *OrderUpdateOne) ClearInternalNote() *OrderUpdateOne {
	ouo.mutation.ClearInternalNote()
	return ouo
}

// SetIsInternal sets the "is_internal" field.
func (ouo *OrderUpdateOne) SetIsInternal(b bool) *OrderUpdateOne {
	ouo.mutation.SetIsInternal(b)
	return ouo
}

// SetNillableIsInternal sets the "is_internal" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableIsInternal(b *bool) *OrderUpdateOne {
	if b != nil {
		ouo.SetIsInternal(*b)
	}
	return ouo
}

// SetAddressID sets the "address_id" field.
func (ouo *OrderUpdateOne) SetAddressID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetAddressID(u)
	return ouo
}

// SetNillableAddressID sets the "address_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableAddressID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetAddressID(*u)
	}
	return ouo
}

// SetParentOrder sets the "parent_order" edge to the Order entity.
func (ouo *OrderUpdateOne) SetParentOrder(o *Order) *OrderUpdateOne {
	return ouo.SetParentOrderID(o.ID)
}

// SetOrderStatusID sets the "order_status" edge to the OrderStatusCode entity by ID.
func (ouo *OrderUpdateOne) SetOrderStatusID(id int) *OrderUpdateOne {
	ouo.mutation.SetOrderStatusID(id)
	return ouo
}

// SetOrderStatus sets the "order_status" edge to the OrderStatusCode entity.
func (ouo *OrderUpdateOne) SetOrderStatus(o *OrderStatusCode) *OrderUpdateOne {
	return ouo.SetOrderStatusID(o.ID)
}

// SetStaff sets the "staff" edge to the Person entity.
func (ouo *OrderUpdateOne) SetStaff(p *Person) *OrderUpdateOne {
	return ouo.SetStaffID(p.ID)
}

// SetOrderAddressID sets the "order_address" edge to the Address entity by ID.
func (ouo *OrderUpdateOne) SetOrderAddressID(id uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetOrderAddressID(id)
	return ouo
}

// SetOrderAddress sets the "order_address" edge to the Address entity.
func (ouo *OrderUpdateOne) SetOrderAddress(a *Address) *OrderUpdateOne {
	return ouo.SetOrderAddressID(a.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// ClearParentOrder clears the "parent_order" edge to the Order entity.
func (ouo *OrderUpdateOne) ClearParentOrder() *OrderUpdateOne {
	ouo.mutation.ClearParentOrder()
	return ouo
}

// ClearOrderStatus clears the "order_status" edge to the OrderStatusCode entity.
func (ouo *OrderUpdateOne) ClearOrderStatus() *OrderUpdateOne {
	ouo.mutation.ClearOrderStatus()
	return ouo
}

// ClearStaff clears the "staff" edge to the Person entity.
func (ouo *OrderUpdateOne) ClearStaff() *OrderUpdateOne {
	ouo.mutation.ClearStaff()
	return ouo
}

// ClearOrderAddress clears the "order_address" edge to the Address entity.
func (ouo *OrderUpdateOne) ClearOrderAddress() *OrderUpdateOne {
	ouo.mutation.ClearOrderAddress()
	return ouo
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	ouo.defaults()
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.Priority(); ok {
		if err := order.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Order.priority": %w`, err)}
		}
	}
	if _, ok := ouo.mutation.CustomerID(); ouo.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.customer"`)
	}
	if _, ok := ouo.mutation.CreatorID(); ouo.mutation.CreatorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.creator"`)
	}
	if _, ok := ouo.mutation.OrderStatusID(); ouo.mutation.OrderStatusCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.order_status"`)
	}
	if _, ok := ouo.mutation.StaffID(); ouo.mutation.StaffCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.staff"`)
	}
	if _, ok := ouo.mutation.OrderAddressID(); ouo.mutation.OrderAddressCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.order_address"`)
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.Note(); ok {
		_spec.SetField(order.FieldNote, field.TypeString, value)
	}
	if ouo.mutation.NoteCleared() {
		_spec.ClearField(order.FieldNote, field.TypeString)
	}
	if value, ok := ouo.mutation.Priority(); ok {
		_spec.SetField(order.FieldPriority, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.AddedPriority(); ok {
		_spec.AddField(order.FieldPriority, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.InternalNote(); ok {
		_spec.SetField(order.FieldInternalNote, field.TypeString, value)
	}
	if ouo.mutation.InternalNoteCleared() {
		_spec.ClearField(order.FieldInternalNote, field.TypeString)
	}
	if value, ok := ouo.mutation.IsInternal(); ok {
		_spec.SetField(order.FieldIsInternal, field.TypeBool, value)
	}
	if ouo.mutation.ParentOrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   order.ParentOrderTable,
			Columns: []string{order.ParentOrderColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ParentOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   order.ParentOrderTable,
			Columns: []string{order.ParentOrderColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.OrderStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.OrderStatusTable,
			Columns: []string{order.OrderStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.OrderStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.OrderStatusTable,
			Columns: []string{order.OrderStatusColumn},
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
	if ouo.mutation.StaffCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.StaffTable,
			Columns: []string{order.StaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.StaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.StaffTable,
			Columns: []string{order.StaffColumn},
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
	if ouo.mutation.OrderAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.OrderAddressTable,
			Columns: []string{order.OrderAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.OrderAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.OrderAddressTable,
			Columns: []string{order.OrderAddressColumn},
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
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
