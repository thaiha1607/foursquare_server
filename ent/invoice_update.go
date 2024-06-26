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
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/invoicestatuscode"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// InvoiceUpdate is the builder for updating Invoice entities.
type InvoiceUpdate struct {
	config
	hooks    []Hook
	mutation *InvoiceMutation
}

// Where appends a list predicates to the InvoiceUpdate builder.
func (iu *InvoiceUpdate) Where(ps ...predicate.Invoice) *InvoiceUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *InvoiceUpdate) SetUpdatedAt(t time.Time) *InvoiceUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetNote sets the "note" field.
func (iu *InvoiceUpdate) SetNote(s string) *InvoiceUpdate {
	iu.mutation.SetNote(s)
	return iu
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableNote(s *string) *InvoiceUpdate {
	if s != nil {
		iu.SetNote(*s)
	}
	return iu
}

// ClearNote clears the value of the "note" field.
func (iu *InvoiceUpdate) ClearNote() *InvoiceUpdate {
	iu.mutation.ClearNote()
	return iu
}

// SetStatusCode sets the "status_code" field.
func (iu *InvoiceUpdate) SetStatusCode(i int) *InvoiceUpdate {
	iu.mutation.SetStatusCode(i)
	return iu
}

// SetNillableStatusCode sets the "status_code" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillableStatusCode(i *int) *InvoiceUpdate {
	if i != nil {
		iu.SetStatusCode(*i)
	}
	return iu
}

// SetPaymentMethod sets the "payment_method" field.
func (iu *InvoiceUpdate) SetPaymentMethod(im invoice.PaymentMethod) *InvoiceUpdate {
	iu.mutation.SetPaymentMethod(im)
	return iu
}

// SetNillablePaymentMethod sets the "payment_method" field if the given value is not nil.
func (iu *InvoiceUpdate) SetNillablePaymentMethod(im *invoice.PaymentMethod) *InvoiceUpdate {
	if im != nil {
		iu.SetPaymentMethod(*im)
	}
	return iu
}

// SetInvoiceStatusID sets the "invoice_status" edge to the InvoiceStatusCode entity by ID.
func (iu *InvoiceUpdate) SetInvoiceStatusID(id int) *InvoiceUpdate {
	iu.mutation.SetInvoiceStatusID(id)
	return iu
}

// SetInvoiceStatus sets the "invoice_status" edge to the InvoiceStatusCode entity.
func (iu *InvoiceUpdate) SetInvoiceStatus(i *InvoiceStatusCode) *InvoiceUpdate {
	return iu.SetInvoiceStatusID(i.ID)
}

// Mutation returns the InvoiceMutation object of the builder.
func (iu *InvoiceUpdate) Mutation() *InvoiceMutation {
	return iu.mutation
}

// ClearInvoiceStatus clears the "invoice_status" edge to the InvoiceStatusCode entity.
func (iu *InvoiceUpdate) ClearInvoiceStatus() *InvoiceUpdate {
	iu.mutation.ClearInvoiceStatus()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InvoiceUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InvoiceUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InvoiceUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InvoiceUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *InvoiceUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := invoice.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *InvoiceUpdate) check() error {
	if v, ok := iu.mutation.PaymentMethod(); ok {
		if err := invoice.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "Invoice.payment_method": %w`, err)}
		}
	}
	if _, ok := iu.mutation.OrderID(); iu.mutation.OrderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Invoice.order"`)
	}
	if _, ok := iu.mutation.InvoiceStatusID(); iu.mutation.InvoiceStatusCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Invoice.invoice_status"`)
	}
	return nil
}

func (iu *InvoiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoice.Table, invoice.Columns, sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(invoice.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iu.mutation.Note(); ok {
		_spec.SetField(invoice.FieldNote, field.TypeString, value)
	}
	if iu.mutation.NoteCleared() {
		_spec.ClearField(invoice.FieldNote, field.TypeString)
	}
	if value, ok := iu.mutation.PaymentMethod(); ok {
		_spec.SetField(invoice.FieldPaymentMethod, field.TypeEnum, value)
	}
	if iu.mutation.InvoiceStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoice.InvoiceStatusTable,
			Columns: []string{invoice.InvoiceStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoicestatuscode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.InvoiceStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoice.InvoiceStatusTable,
			Columns: []string{invoice.InvoiceStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoicestatuscode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// InvoiceUpdateOne is the builder for updating a single Invoice entity.
type InvoiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InvoiceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *InvoiceUpdateOne) SetUpdatedAt(t time.Time) *InvoiceUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetNote sets the "note" field.
func (iuo *InvoiceUpdateOne) SetNote(s string) *InvoiceUpdateOne {
	iuo.mutation.SetNote(s)
	return iuo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableNote(s *string) *InvoiceUpdateOne {
	if s != nil {
		iuo.SetNote(*s)
	}
	return iuo
}

// ClearNote clears the value of the "note" field.
func (iuo *InvoiceUpdateOne) ClearNote() *InvoiceUpdateOne {
	iuo.mutation.ClearNote()
	return iuo
}

// SetStatusCode sets the "status_code" field.
func (iuo *InvoiceUpdateOne) SetStatusCode(i int) *InvoiceUpdateOne {
	iuo.mutation.SetStatusCode(i)
	return iuo
}

// SetNillableStatusCode sets the "status_code" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillableStatusCode(i *int) *InvoiceUpdateOne {
	if i != nil {
		iuo.SetStatusCode(*i)
	}
	return iuo
}

// SetPaymentMethod sets the "payment_method" field.
func (iuo *InvoiceUpdateOne) SetPaymentMethod(im invoice.PaymentMethod) *InvoiceUpdateOne {
	iuo.mutation.SetPaymentMethod(im)
	return iuo
}

// SetNillablePaymentMethod sets the "payment_method" field if the given value is not nil.
func (iuo *InvoiceUpdateOne) SetNillablePaymentMethod(im *invoice.PaymentMethod) *InvoiceUpdateOne {
	if im != nil {
		iuo.SetPaymentMethod(*im)
	}
	return iuo
}

// SetInvoiceStatusID sets the "invoice_status" edge to the InvoiceStatusCode entity by ID.
func (iuo *InvoiceUpdateOne) SetInvoiceStatusID(id int) *InvoiceUpdateOne {
	iuo.mutation.SetInvoiceStatusID(id)
	return iuo
}

// SetInvoiceStatus sets the "invoice_status" edge to the InvoiceStatusCode entity.
func (iuo *InvoiceUpdateOne) SetInvoiceStatus(i *InvoiceStatusCode) *InvoiceUpdateOne {
	return iuo.SetInvoiceStatusID(i.ID)
}

// Mutation returns the InvoiceMutation object of the builder.
func (iuo *InvoiceUpdateOne) Mutation() *InvoiceMutation {
	return iuo.mutation
}

// ClearInvoiceStatus clears the "invoice_status" edge to the InvoiceStatusCode entity.
func (iuo *InvoiceUpdateOne) ClearInvoiceStatus() *InvoiceUpdateOne {
	iuo.mutation.ClearInvoiceStatus()
	return iuo
}

// Where appends a list predicates to the InvoiceUpdate builder.
func (iuo *InvoiceUpdateOne) Where(ps ...predicate.Invoice) *InvoiceUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InvoiceUpdateOne) Select(field string, fields ...string) *InvoiceUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Invoice entity.
func (iuo *InvoiceUpdateOne) Save(ctx context.Context) (*Invoice, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InvoiceUpdateOne) SaveX(ctx context.Context) *Invoice {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InvoiceUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InvoiceUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *InvoiceUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := invoice.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *InvoiceUpdateOne) check() error {
	if v, ok := iuo.mutation.PaymentMethod(); ok {
		if err := invoice.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "Invoice.payment_method": %w`, err)}
		}
	}
	if _, ok := iuo.mutation.OrderID(); iuo.mutation.OrderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Invoice.order"`)
	}
	if _, ok := iuo.mutation.InvoiceStatusID(); iuo.mutation.InvoiceStatusCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Invoice.invoice_status"`)
	}
	return nil
}

func (iuo *InvoiceUpdateOne) sqlSave(ctx context.Context) (_node *Invoice, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoice.Table, invoice.Columns, sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Invoice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoice.FieldID)
		for _, f := range fields {
			if !invoice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != invoice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(invoice.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.Note(); ok {
		_spec.SetField(invoice.FieldNote, field.TypeString, value)
	}
	if iuo.mutation.NoteCleared() {
		_spec.ClearField(invoice.FieldNote, field.TypeString)
	}
	if value, ok := iuo.mutation.PaymentMethod(); ok {
		_spec.SetField(invoice.FieldPaymentMethod, field.TypeEnum, value)
	}
	if iuo.mutation.InvoiceStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoice.InvoiceStatusTable,
			Columns: []string{invoice.InvoiceStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoicestatuscode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.InvoiceStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoice.InvoiceStatusTable,
			Columns: []string{invoice.InvoiceStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoicestatuscode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Invoice{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
