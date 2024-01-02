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
	"github.com/shopspring/decimal"
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/invoicelineitem"
	"github.com/thaiha1607/foursquare_server/ent/orderlineitem"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// InvoiceLineItemUpdate is the builder for updating InvoiceLineItem entities.
type InvoiceLineItemUpdate struct {
	config
	hooks    []Hook
	mutation *InvoiceLineItemMutation
}

// Where appends a list predicates to the InvoiceLineItemUpdate builder.
func (iliu *InvoiceLineItemUpdate) Where(ps ...predicate.InvoiceLineItem) *InvoiceLineItemUpdate {
	iliu.mutation.Where(ps...)
	return iliu
}

// SetUpdatedAt sets the "updated_at" field.
func (iliu *InvoiceLineItemUpdate) SetUpdatedAt(t time.Time) *InvoiceLineItemUpdate {
	iliu.mutation.SetUpdatedAt(t)
	return iliu
}

// SetInvoiceID sets the "invoice_id" field.
func (iliu *InvoiceLineItemUpdate) SetInvoiceID(u uuid.UUID) *InvoiceLineItemUpdate {
	iliu.mutation.SetInvoiceID(u)
	return iliu
}

// SetNillableInvoiceID sets the "invoice_id" field if the given value is not nil.
func (iliu *InvoiceLineItemUpdate) SetNillableInvoiceID(u *uuid.UUID) *InvoiceLineItemUpdate {
	if u != nil {
		iliu.SetInvoiceID(*u)
	}
	return iliu
}

// SetOrderLineItemID sets the "order_line_item_id" field.
func (iliu *InvoiceLineItemUpdate) SetOrderLineItemID(u uuid.UUID) *InvoiceLineItemUpdate {
	iliu.mutation.SetOrderLineItemID(u)
	return iliu
}

// SetNillableOrderLineItemID sets the "order_line_item_id" field if the given value is not nil.
func (iliu *InvoiceLineItemUpdate) SetNillableOrderLineItemID(u *uuid.UUID) *InvoiceLineItemUpdate {
	if u != nil {
		iliu.SetOrderLineItemID(*u)
	}
	return iliu
}

// SetQty sets the "qty" field.
func (iliu *InvoiceLineItemUpdate) SetQty(d decimal.Decimal) *InvoiceLineItemUpdate {
	iliu.mutation.ResetQty()
	iliu.mutation.SetQty(d)
	return iliu
}

// SetNillableQty sets the "qty" field if the given value is not nil.
func (iliu *InvoiceLineItemUpdate) SetNillableQty(d *decimal.Decimal) *InvoiceLineItemUpdate {
	if d != nil {
		iliu.SetQty(*d)
	}
	return iliu
}

// AddQty adds d to the "qty" field.
func (iliu *InvoiceLineItemUpdate) AddQty(d decimal.Decimal) *InvoiceLineItemUpdate {
	iliu.mutation.AddQty(d)
	return iliu
}

// SetTotal sets the "total" field.
func (iliu *InvoiceLineItemUpdate) SetTotal(d decimal.Decimal) *InvoiceLineItemUpdate {
	iliu.mutation.ResetTotal()
	iliu.mutation.SetTotal(d)
	return iliu
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (iliu *InvoiceLineItemUpdate) SetNillableTotal(d *decimal.Decimal) *InvoiceLineItemUpdate {
	if d != nil {
		iliu.SetTotal(*d)
	}
	return iliu
}

// AddTotal adds d to the "total" field.
func (iliu *InvoiceLineItemUpdate) AddTotal(d decimal.Decimal) *InvoiceLineItemUpdate {
	iliu.mutation.AddTotal(d)
	return iliu
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (iliu *InvoiceLineItemUpdate) SetInvoice(i *Invoice) *InvoiceLineItemUpdate {
	return iliu.SetInvoiceID(i.ID)
}

// SetOrderLineItem sets the "order_line_item" edge to the OrderLineItem entity.
func (iliu *InvoiceLineItemUpdate) SetOrderLineItem(o *OrderLineItem) *InvoiceLineItemUpdate {
	return iliu.SetOrderLineItemID(o.ID)
}

// Mutation returns the InvoiceLineItemMutation object of the builder.
func (iliu *InvoiceLineItemUpdate) Mutation() *InvoiceLineItemMutation {
	return iliu.mutation
}

// ClearInvoice clears the "invoice" edge to the Invoice entity.
func (iliu *InvoiceLineItemUpdate) ClearInvoice() *InvoiceLineItemUpdate {
	iliu.mutation.ClearInvoice()
	return iliu
}

// ClearOrderLineItem clears the "order_line_item" edge to the OrderLineItem entity.
func (iliu *InvoiceLineItemUpdate) ClearOrderLineItem() *InvoiceLineItemUpdate {
	iliu.mutation.ClearOrderLineItem()
	return iliu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iliu *InvoiceLineItemUpdate) Save(ctx context.Context) (int, error) {
	iliu.defaults()
	return withHooks(ctx, iliu.sqlSave, iliu.mutation, iliu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iliu *InvoiceLineItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iliu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iliu *InvoiceLineItemUpdate) Exec(ctx context.Context) error {
	_, err := iliu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iliu *InvoiceLineItemUpdate) ExecX(ctx context.Context) {
	if err := iliu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iliu *InvoiceLineItemUpdate) defaults() {
	if _, ok := iliu.mutation.UpdatedAt(); !ok {
		v := invoicelineitem.UpdateDefaultUpdatedAt()
		iliu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iliu *InvoiceLineItemUpdate) check() error {
	if _, ok := iliu.mutation.InvoiceID(); iliu.mutation.InvoiceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceLineItem.invoice"`)
	}
	if _, ok := iliu.mutation.OrderLineItemID(); iliu.mutation.OrderLineItemCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceLineItem.order_line_item"`)
	}
	return nil
}

func (iliu *InvoiceLineItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iliu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoicelineitem.Table, invoicelineitem.Columns, sqlgraph.NewFieldSpec(invoicelineitem.FieldID, field.TypeUUID))
	if ps := iliu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iliu.mutation.UpdatedAt(); ok {
		_spec.SetField(invoicelineitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iliu.mutation.Qty(); ok {
		_spec.SetField(invoicelineitem.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := iliu.mutation.AddedQty(); ok {
		_spec.AddField(invoicelineitem.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := iliu.mutation.Total(); ok {
		_spec.SetField(invoicelineitem.FieldTotal, field.TypeFloat64, value)
	}
	if value, ok := iliu.mutation.AddedTotal(); ok {
		_spec.AddField(invoicelineitem.FieldTotal, field.TypeFloat64, value)
	}
	if iliu.mutation.InvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicelineitem.InvoiceTable,
			Columns: []string{invoicelineitem.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iliu.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicelineitem.InvoiceTable,
			Columns: []string{invoicelineitem.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iliu.mutation.OrderLineItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicelineitem.OrderLineItemTable,
			Columns: []string{invoicelineitem.OrderLineItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderlineitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iliu.mutation.OrderLineItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicelineitem.OrderLineItemTable,
			Columns: []string{invoicelineitem.OrderLineItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderlineitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iliu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoicelineitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iliu.mutation.done = true
	return n, nil
}

// InvoiceLineItemUpdateOne is the builder for updating a single InvoiceLineItem entity.
type InvoiceLineItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InvoiceLineItemMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (iliuo *InvoiceLineItemUpdateOne) SetUpdatedAt(t time.Time) *InvoiceLineItemUpdateOne {
	iliuo.mutation.SetUpdatedAt(t)
	return iliuo
}

// SetInvoiceID sets the "invoice_id" field.
func (iliuo *InvoiceLineItemUpdateOne) SetInvoiceID(u uuid.UUID) *InvoiceLineItemUpdateOne {
	iliuo.mutation.SetInvoiceID(u)
	return iliuo
}

// SetNillableInvoiceID sets the "invoice_id" field if the given value is not nil.
func (iliuo *InvoiceLineItemUpdateOne) SetNillableInvoiceID(u *uuid.UUID) *InvoiceLineItemUpdateOne {
	if u != nil {
		iliuo.SetInvoiceID(*u)
	}
	return iliuo
}

// SetOrderLineItemID sets the "order_line_item_id" field.
func (iliuo *InvoiceLineItemUpdateOne) SetOrderLineItemID(u uuid.UUID) *InvoiceLineItemUpdateOne {
	iliuo.mutation.SetOrderLineItemID(u)
	return iliuo
}

// SetNillableOrderLineItemID sets the "order_line_item_id" field if the given value is not nil.
func (iliuo *InvoiceLineItemUpdateOne) SetNillableOrderLineItemID(u *uuid.UUID) *InvoiceLineItemUpdateOne {
	if u != nil {
		iliuo.SetOrderLineItemID(*u)
	}
	return iliuo
}

// SetQty sets the "qty" field.
func (iliuo *InvoiceLineItemUpdateOne) SetQty(d decimal.Decimal) *InvoiceLineItemUpdateOne {
	iliuo.mutation.ResetQty()
	iliuo.mutation.SetQty(d)
	return iliuo
}

// SetNillableQty sets the "qty" field if the given value is not nil.
func (iliuo *InvoiceLineItemUpdateOne) SetNillableQty(d *decimal.Decimal) *InvoiceLineItemUpdateOne {
	if d != nil {
		iliuo.SetQty(*d)
	}
	return iliuo
}

// AddQty adds d to the "qty" field.
func (iliuo *InvoiceLineItemUpdateOne) AddQty(d decimal.Decimal) *InvoiceLineItemUpdateOne {
	iliuo.mutation.AddQty(d)
	return iliuo
}

// SetTotal sets the "total" field.
func (iliuo *InvoiceLineItemUpdateOne) SetTotal(d decimal.Decimal) *InvoiceLineItemUpdateOne {
	iliuo.mutation.ResetTotal()
	iliuo.mutation.SetTotal(d)
	return iliuo
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (iliuo *InvoiceLineItemUpdateOne) SetNillableTotal(d *decimal.Decimal) *InvoiceLineItemUpdateOne {
	if d != nil {
		iliuo.SetTotal(*d)
	}
	return iliuo
}

// AddTotal adds d to the "total" field.
func (iliuo *InvoiceLineItemUpdateOne) AddTotal(d decimal.Decimal) *InvoiceLineItemUpdateOne {
	iliuo.mutation.AddTotal(d)
	return iliuo
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (iliuo *InvoiceLineItemUpdateOne) SetInvoice(i *Invoice) *InvoiceLineItemUpdateOne {
	return iliuo.SetInvoiceID(i.ID)
}

// SetOrderLineItem sets the "order_line_item" edge to the OrderLineItem entity.
func (iliuo *InvoiceLineItemUpdateOne) SetOrderLineItem(o *OrderLineItem) *InvoiceLineItemUpdateOne {
	return iliuo.SetOrderLineItemID(o.ID)
}

// Mutation returns the InvoiceLineItemMutation object of the builder.
func (iliuo *InvoiceLineItemUpdateOne) Mutation() *InvoiceLineItemMutation {
	return iliuo.mutation
}

// ClearInvoice clears the "invoice" edge to the Invoice entity.
func (iliuo *InvoiceLineItemUpdateOne) ClearInvoice() *InvoiceLineItemUpdateOne {
	iliuo.mutation.ClearInvoice()
	return iliuo
}

// ClearOrderLineItem clears the "order_line_item" edge to the OrderLineItem entity.
func (iliuo *InvoiceLineItemUpdateOne) ClearOrderLineItem() *InvoiceLineItemUpdateOne {
	iliuo.mutation.ClearOrderLineItem()
	return iliuo
}

// Where appends a list predicates to the InvoiceLineItemUpdate builder.
func (iliuo *InvoiceLineItemUpdateOne) Where(ps ...predicate.InvoiceLineItem) *InvoiceLineItemUpdateOne {
	iliuo.mutation.Where(ps...)
	return iliuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iliuo *InvoiceLineItemUpdateOne) Select(field string, fields ...string) *InvoiceLineItemUpdateOne {
	iliuo.fields = append([]string{field}, fields...)
	return iliuo
}

// Save executes the query and returns the updated InvoiceLineItem entity.
func (iliuo *InvoiceLineItemUpdateOne) Save(ctx context.Context) (*InvoiceLineItem, error) {
	iliuo.defaults()
	return withHooks(ctx, iliuo.sqlSave, iliuo.mutation, iliuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iliuo *InvoiceLineItemUpdateOne) SaveX(ctx context.Context) *InvoiceLineItem {
	node, err := iliuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iliuo *InvoiceLineItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iliuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iliuo *InvoiceLineItemUpdateOne) ExecX(ctx context.Context) {
	if err := iliuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iliuo *InvoiceLineItemUpdateOne) defaults() {
	if _, ok := iliuo.mutation.UpdatedAt(); !ok {
		v := invoicelineitem.UpdateDefaultUpdatedAt()
		iliuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iliuo *InvoiceLineItemUpdateOne) check() error {
	if _, ok := iliuo.mutation.InvoiceID(); iliuo.mutation.InvoiceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceLineItem.invoice"`)
	}
	if _, ok := iliuo.mutation.OrderLineItemID(); iliuo.mutation.OrderLineItemCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceLineItem.order_line_item"`)
	}
	return nil
}

func (iliuo *InvoiceLineItemUpdateOne) sqlSave(ctx context.Context) (_node *InvoiceLineItem, err error) {
	if err := iliuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoicelineitem.Table, invoicelineitem.Columns, sqlgraph.NewFieldSpec(invoicelineitem.FieldID, field.TypeUUID))
	id, ok := iliuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "InvoiceLineItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iliuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoicelineitem.FieldID)
		for _, f := range fields {
			if !invoicelineitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != invoicelineitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iliuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iliuo.mutation.UpdatedAt(); ok {
		_spec.SetField(invoicelineitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iliuo.mutation.Qty(); ok {
		_spec.SetField(invoicelineitem.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := iliuo.mutation.AddedQty(); ok {
		_spec.AddField(invoicelineitem.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := iliuo.mutation.Total(); ok {
		_spec.SetField(invoicelineitem.FieldTotal, field.TypeFloat64, value)
	}
	if value, ok := iliuo.mutation.AddedTotal(); ok {
		_spec.AddField(invoicelineitem.FieldTotal, field.TypeFloat64, value)
	}
	if iliuo.mutation.InvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicelineitem.InvoiceTable,
			Columns: []string{invoicelineitem.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iliuo.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicelineitem.InvoiceTable,
			Columns: []string{invoicelineitem.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iliuo.mutation.OrderLineItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicelineitem.OrderLineItemTable,
			Columns: []string{invoicelineitem.OrderLineItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderlineitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iliuo.mutation.OrderLineItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicelineitem.OrderLineItemTable,
			Columns: []string{invoicelineitem.OrderLineItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderlineitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &InvoiceLineItem{config: iliuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iliuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoicelineitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iliuo.mutation.done = true
	return _node, nil
}
