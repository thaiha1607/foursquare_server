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
	"github.com/shopspring/decimal"
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/invoicestatuscode"
	"github.com/thaiha1607/foursquare_server/ent/order"
)

// InvoiceCreate is the builder for creating a Invoice entity.
type InvoiceCreate struct {
	config
	mutation *InvoiceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *InvoiceCreate) SetCreatedAt(t time.Time) *InvoiceCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCreatedAt(t *time.Time) *InvoiceCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *InvoiceCreate) SetUpdatedAt(t time.Time) *InvoiceCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableUpdatedAt(t *time.Time) *InvoiceCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetOrderID sets the "order_id" field.
func (ic *InvoiceCreate) SetOrderID(u uuid.UUID) *InvoiceCreate {
	ic.mutation.SetOrderID(u)
	return ic
}

// SetTotal sets the "total" field.
func (ic *InvoiceCreate) SetTotal(d decimal.Decimal) *InvoiceCreate {
	ic.mutation.SetTotal(d)
	return ic
}

// SetNote sets the "note" field.
func (ic *InvoiceCreate) SetNote(s string) *InvoiceCreate {
	ic.mutation.SetNote(s)
	return ic
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableNote(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetNote(*s)
	}
	return ic
}

// SetType sets the "type" field.
func (ic *InvoiceCreate) SetType(i invoice.Type) *InvoiceCreate {
	ic.mutation.SetType(i)
	return ic
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableType(i *invoice.Type) *InvoiceCreate {
	if i != nil {
		ic.SetType(*i)
	}
	return ic
}

// SetStatusCode sets the "status_code" field.
func (ic *InvoiceCreate) SetStatusCode(i int) *InvoiceCreate {
	ic.mutation.SetStatusCode(i)
	return ic
}

// SetNillableStatusCode sets the "status_code" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableStatusCode(i *int) *InvoiceCreate {
	if i != nil {
		ic.SetStatusCode(*i)
	}
	return ic
}

// SetPaymentMethod sets the "payment_method" field.
func (ic *InvoiceCreate) SetPaymentMethod(im invoice.PaymentMethod) *InvoiceCreate {
	ic.mutation.SetPaymentMethod(im)
	return ic
}

// SetNillablePaymentMethod sets the "payment_method" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillablePaymentMethod(im *invoice.PaymentMethod) *InvoiceCreate {
	if im != nil {
		ic.SetPaymentMethod(*im)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *InvoiceCreate) SetID(u uuid.UUID) *InvoiceCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableID(u *uuid.UUID) *InvoiceCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// SetOrder sets the "order" edge to the Order entity.
func (ic *InvoiceCreate) SetOrder(o *Order) *InvoiceCreate {
	return ic.SetOrderID(o.ID)
}

// SetInvoiceStatusID sets the "invoice_status" edge to the InvoiceStatusCode entity by ID.
func (ic *InvoiceCreate) SetInvoiceStatusID(id int) *InvoiceCreate {
	ic.mutation.SetInvoiceStatusID(id)
	return ic
}

// SetInvoiceStatus sets the "invoice_status" edge to the InvoiceStatusCode entity.
func (ic *InvoiceCreate) SetInvoiceStatus(i *InvoiceStatusCode) *InvoiceCreate {
	return ic.SetInvoiceStatusID(i.ID)
}

// Mutation returns the InvoiceMutation object of the builder.
func (ic *InvoiceCreate) Mutation() *InvoiceMutation {
	return ic.mutation
}

// Save creates the Invoice in the database.
func (ic *InvoiceCreate) Save(ctx context.Context) (*Invoice, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InvoiceCreate) SaveX(ctx context.Context) *Invoice {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InvoiceCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InvoiceCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InvoiceCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := invoice.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := invoice.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.GetType(); !ok {
		v := invoice.DefaultType
		ic.mutation.SetType(v)
	}
	if _, ok := ic.mutation.StatusCode(); !ok {
		v := invoice.DefaultStatusCode
		ic.mutation.SetStatusCode(v)
	}
	if _, ok := ic.mutation.PaymentMethod(); !ok {
		v := invoice.DefaultPaymentMethod
		ic.mutation.SetPaymentMethod(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := invoice.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InvoiceCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Invoice.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Invoice.updated_at"`)}
	}
	if _, ok := ic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "Invoice.order_id"`)}
	}
	if _, ok := ic.mutation.Total(); !ok {
		return &ValidationError{Name: "total", err: errors.New(`ent: missing required field "Invoice.total"`)}
	}
	if _, ok := ic.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Invoice.type"`)}
	}
	if v, ok := ic.mutation.GetType(); ok {
		if err := invoice.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Invoice.type": %w`, err)}
		}
	}
	if _, ok := ic.mutation.StatusCode(); !ok {
		return &ValidationError{Name: "status_code", err: errors.New(`ent: missing required field "Invoice.status_code"`)}
	}
	if v, ok := ic.mutation.PaymentMethod(); ok {
		if err := invoice.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "Invoice.payment_method": %w`, err)}
		}
	}
	if _, ok := ic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required edge "Invoice.order"`)}
	}
	if _, ok := ic.mutation.InvoiceStatusID(); !ok {
		return &ValidationError{Name: "invoice_status", err: errors.New(`ent: missing required edge "Invoice.invoice_status"`)}
	}
	return nil
}

func (ic *InvoiceCreate) sqlSave(ctx context.Context) (*Invoice, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
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
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InvoiceCreate) createSpec() (*Invoice, *sqlgraph.CreateSpec) {
	var (
		_node = &Invoice{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(invoice.Table, sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(invoice.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(invoice.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.Total(); ok {
		_spec.SetField(invoice.FieldTotal, field.TypeFloat64, value)
		_node.Total = value
	}
	if value, ok := ic.mutation.Note(); ok {
		_spec.SetField(invoice.FieldNote, field.TypeString, value)
		_node.Note = &value
	}
	if value, ok := ic.mutation.GetType(); ok {
		_spec.SetField(invoice.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ic.mutation.PaymentMethod(); ok {
		_spec.SetField(invoice.FieldPaymentMethod, field.TypeEnum, value)
		_node.PaymentMethod = &value
	}
	if nodes := ic.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoice.OrderTable,
			Columns: []string{invoice.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.InvoiceStatusIDs(); len(nodes) > 0 {
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
		_node.StatusCode = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InvoiceCreateBulk is the builder for creating many Invoice entities in bulk.
type InvoiceCreateBulk struct {
	config
	err      error
	builders []*InvoiceCreate
}

// Save creates the Invoice entities in the database.
func (icb *InvoiceCreateBulk) Save(ctx context.Context) ([]*Invoice, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Invoice, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvoiceMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InvoiceCreateBulk) SaveX(ctx context.Context) []*Invoice {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InvoiceCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InvoiceCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
