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
	"github.com/thaiha1607/foursquare_server/ent/financialtransaction"
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/transactiontype"
)

// FinancialTransactionUpdate is the builder for updating FinancialTransaction entities.
type FinancialTransactionUpdate struct {
	config
	hooks    []Hook
	mutation *FinancialTransactionMutation
}

// Where appends a list predicates to the FinancialTransactionUpdate builder.
func (ftu *FinancialTransactionUpdate) Where(ps ...predicate.FinancialTransaction) *FinancialTransactionUpdate {
	ftu.mutation.Where(ps...)
	return ftu
}

// SetUpdatedAt sets the "updated_at" field.
func (ftu *FinancialTransactionUpdate) SetUpdatedAt(t time.Time) *FinancialTransactionUpdate {
	ftu.mutation.SetUpdatedAt(t)
	return ftu
}

// SetInvoiceID sets the "invoice_id" field.
func (ftu *FinancialTransactionUpdate) SetInvoiceID(u uuid.UUID) *FinancialTransactionUpdate {
	ftu.mutation.SetInvoiceID(u)
	return ftu
}

// SetNillableInvoiceID sets the "invoice_id" field if the given value is not nil.
func (ftu *FinancialTransactionUpdate) SetNillableInvoiceID(u *uuid.UUID) *FinancialTransactionUpdate {
	if u != nil {
		ftu.SetInvoiceID(*u)
	}
	return ftu
}

// SetAmount sets the "amount" field.
func (ftu *FinancialTransactionUpdate) SetAmount(d decimal.Decimal) *FinancialTransactionUpdate {
	ftu.mutation.ResetAmount()
	ftu.mutation.SetAmount(d)
	return ftu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (ftu *FinancialTransactionUpdate) SetNillableAmount(d *decimal.Decimal) *FinancialTransactionUpdate {
	if d != nil {
		ftu.SetAmount(*d)
	}
	return ftu
}

// AddAmount adds d to the "amount" field.
func (ftu *FinancialTransactionUpdate) AddAmount(d decimal.Decimal) *FinancialTransactionUpdate {
	ftu.mutation.AddAmount(d)
	return ftu
}

// SetComment sets the "comment" field.
func (ftu *FinancialTransactionUpdate) SetComment(s string) *FinancialTransactionUpdate {
	ftu.mutation.SetComment(s)
	return ftu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (ftu *FinancialTransactionUpdate) SetNillableComment(s *string) *FinancialTransactionUpdate {
	if s != nil {
		ftu.SetComment(*s)
	}
	return ftu
}

// ClearComment clears the value of the "comment" field.
func (ftu *FinancialTransactionUpdate) ClearComment() *FinancialTransactionUpdate {
	ftu.mutation.ClearComment()
	return ftu
}

// SetType sets the "type" field.
func (ftu *FinancialTransactionUpdate) SetType(i int) *FinancialTransactionUpdate {
	ftu.mutation.SetType(i)
	return ftu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftu *FinancialTransactionUpdate) SetNillableType(i *int) *FinancialTransactionUpdate {
	if i != nil {
		ftu.SetType(*i)
	}
	return ftu
}

// SetPaymentMethod sets the "payment_method" field.
func (ftu *FinancialTransactionUpdate) SetPaymentMethod(fm financialtransaction.PaymentMethod) *FinancialTransactionUpdate {
	ftu.mutation.SetPaymentMethod(fm)
	return ftu
}

// SetNillablePaymentMethod sets the "payment_method" field if the given value is not nil.
func (ftu *FinancialTransactionUpdate) SetNillablePaymentMethod(fm *financialtransaction.PaymentMethod) *FinancialTransactionUpdate {
	if fm != nil {
		ftu.SetPaymentMethod(*fm)
	}
	return ftu
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (ftu *FinancialTransactionUpdate) SetInvoice(i *Invoice) *FinancialTransactionUpdate {
	return ftu.SetInvoiceID(i.ID)
}

// SetTransactionTypeID sets the "transaction_type" edge to the TransactionType entity by ID.
func (ftu *FinancialTransactionUpdate) SetTransactionTypeID(id int) *FinancialTransactionUpdate {
	ftu.mutation.SetTransactionTypeID(id)
	return ftu
}

// SetTransactionType sets the "transaction_type" edge to the TransactionType entity.
func (ftu *FinancialTransactionUpdate) SetTransactionType(t *TransactionType) *FinancialTransactionUpdate {
	return ftu.SetTransactionTypeID(t.ID)
}

// Mutation returns the FinancialTransactionMutation object of the builder.
func (ftu *FinancialTransactionUpdate) Mutation() *FinancialTransactionMutation {
	return ftu.mutation
}

// ClearInvoice clears the "invoice" edge to the Invoice entity.
func (ftu *FinancialTransactionUpdate) ClearInvoice() *FinancialTransactionUpdate {
	ftu.mutation.ClearInvoice()
	return ftu
}

// ClearTransactionType clears the "transaction_type" edge to the TransactionType entity.
func (ftu *FinancialTransactionUpdate) ClearTransactionType() *FinancialTransactionUpdate {
	ftu.mutation.ClearTransactionType()
	return ftu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ftu *FinancialTransactionUpdate) Save(ctx context.Context) (int, error) {
	ftu.defaults()
	return withHooks(ctx, ftu.sqlSave, ftu.mutation, ftu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FinancialTransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FinancialTransactionUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FinancialTransactionUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftu *FinancialTransactionUpdate) defaults() {
	if _, ok := ftu.mutation.UpdatedAt(); !ok {
		v := financialtransaction.UpdateDefaultUpdatedAt()
		ftu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftu *FinancialTransactionUpdate) check() error {
	if v, ok := ftu.mutation.PaymentMethod(); ok {
		if err := financialtransaction.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "FinancialTransaction.payment_method": %w`, err)}
		}
	}
	if _, ok := ftu.mutation.InvoiceID(); ftu.mutation.InvoiceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FinancialTransaction.invoice"`)
	}
	if _, ok := ftu.mutation.TransactionTypeID(); ftu.mutation.TransactionTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FinancialTransaction.transaction_type"`)
	}
	return nil
}

func (ftu *FinancialTransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ftu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(financialtransaction.Table, financialtransaction.Columns, sqlgraph.NewFieldSpec(financialtransaction.FieldID, field.TypeUUID))
	if ps := ftu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftu.mutation.UpdatedAt(); ok {
		_spec.SetField(financialtransaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ftu.mutation.Amount(); ok {
		_spec.SetField(financialtransaction.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := ftu.mutation.AddedAmount(); ok {
		_spec.AddField(financialtransaction.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := ftu.mutation.Comment(); ok {
		_spec.SetField(financialtransaction.FieldComment, field.TypeString, value)
	}
	if ftu.mutation.CommentCleared() {
		_spec.ClearField(financialtransaction.FieldComment, field.TypeString)
	}
	if value, ok := ftu.mutation.PaymentMethod(); ok {
		_spec.SetField(financialtransaction.FieldPaymentMethod, field.TypeEnum, value)
	}
	if ftu.mutation.InvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   financialtransaction.InvoiceTable,
			Columns: []string{financialtransaction.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   financialtransaction.InvoiceTable,
			Columns: []string{financialtransaction.InvoiceColumn},
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
	if ftu.mutation.TransactionTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   financialtransaction.TransactionTypeTable,
			Columns: []string{financialtransaction.TransactionTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transactiontype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.TransactionTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   financialtransaction.TransactionTypeTable,
			Columns: []string{financialtransaction.TransactionTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transactiontype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{financialtransaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ftu.mutation.done = true
	return n, nil
}

// FinancialTransactionUpdateOne is the builder for updating a single FinancialTransaction entity.
type FinancialTransactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FinancialTransactionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ftuo *FinancialTransactionUpdateOne) SetUpdatedAt(t time.Time) *FinancialTransactionUpdateOne {
	ftuo.mutation.SetUpdatedAt(t)
	return ftuo
}

// SetInvoiceID sets the "invoice_id" field.
func (ftuo *FinancialTransactionUpdateOne) SetInvoiceID(u uuid.UUID) *FinancialTransactionUpdateOne {
	ftuo.mutation.SetInvoiceID(u)
	return ftuo
}

// SetNillableInvoiceID sets the "invoice_id" field if the given value is not nil.
func (ftuo *FinancialTransactionUpdateOne) SetNillableInvoiceID(u *uuid.UUID) *FinancialTransactionUpdateOne {
	if u != nil {
		ftuo.SetInvoiceID(*u)
	}
	return ftuo
}

// SetAmount sets the "amount" field.
func (ftuo *FinancialTransactionUpdateOne) SetAmount(d decimal.Decimal) *FinancialTransactionUpdateOne {
	ftuo.mutation.ResetAmount()
	ftuo.mutation.SetAmount(d)
	return ftuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (ftuo *FinancialTransactionUpdateOne) SetNillableAmount(d *decimal.Decimal) *FinancialTransactionUpdateOne {
	if d != nil {
		ftuo.SetAmount(*d)
	}
	return ftuo
}

// AddAmount adds d to the "amount" field.
func (ftuo *FinancialTransactionUpdateOne) AddAmount(d decimal.Decimal) *FinancialTransactionUpdateOne {
	ftuo.mutation.AddAmount(d)
	return ftuo
}

// SetComment sets the "comment" field.
func (ftuo *FinancialTransactionUpdateOne) SetComment(s string) *FinancialTransactionUpdateOne {
	ftuo.mutation.SetComment(s)
	return ftuo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (ftuo *FinancialTransactionUpdateOne) SetNillableComment(s *string) *FinancialTransactionUpdateOne {
	if s != nil {
		ftuo.SetComment(*s)
	}
	return ftuo
}

// ClearComment clears the value of the "comment" field.
func (ftuo *FinancialTransactionUpdateOne) ClearComment() *FinancialTransactionUpdateOne {
	ftuo.mutation.ClearComment()
	return ftuo
}

// SetType sets the "type" field.
func (ftuo *FinancialTransactionUpdateOne) SetType(i int) *FinancialTransactionUpdateOne {
	ftuo.mutation.SetType(i)
	return ftuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftuo *FinancialTransactionUpdateOne) SetNillableType(i *int) *FinancialTransactionUpdateOne {
	if i != nil {
		ftuo.SetType(*i)
	}
	return ftuo
}

// SetPaymentMethod sets the "payment_method" field.
func (ftuo *FinancialTransactionUpdateOne) SetPaymentMethod(fm financialtransaction.PaymentMethod) *FinancialTransactionUpdateOne {
	ftuo.mutation.SetPaymentMethod(fm)
	return ftuo
}

// SetNillablePaymentMethod sets the "payment_method" field if the given value is not nil.
func (ftuo *FinancialTransactionUpdateOne) SetNillablePaymentMethod(fm *financialtransaction.PaymentMethod) *FinancialTransactionUpdateOne {
	if fm != nil {
		ftuo.SetPaymentMethod(*fm)
	}
	return ftuo
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (ftuo *FinancialTransactionUpdateOne) SetInvoice(i *Invoice) *FinancialTransactionUpdateOne {
	return ftuo.SetInvoiceID(i.ID)
}

// SetTransactionTypeID sets the "transaction_type" edge to the TransactionType entity by ID.
func (ftuo *FinancialTransactionUpdateOne) SetTransactionTypeID(id int) *FinancialTransactionUpdateOne {
	ftuo.mutation.SetTransactionTypeID(id)
	return ftuo
}

// SetTransactionType sets the "transaction_type" edge to the TransactionType entity.
func (ftuo *FinancialTransactionUpdateOne) SetTransactionType(t *TransactionType) *FinancialTransactionUpdateOne {
	return ftuo.SetTransactionTypeID(t.ID)
}

// Mutation returns the FinancialTransactionMutation object of the builder.
func (ftuo *FinancialTransactionUpdateOne) Mutation() *FinancialTransactionMutation {
	return ftuo.mutation
}

// ClearInvoice clears the "invoice" edge to the Invoice entity.
func (ftuo *FinancialTransactionUpdateOne) ClearInvoice() *FinancialTransactionUpdateOne {
	ftuo.mutation.ClearInvoice()
	return ftuo
}

// ClearTransactionType clears the "transaction_type" edge to the TransactionType entity.
func (ftuo *FinancialTransactionUpdateOne) ClearTransactionType() *FinancialTransactionUpdateOne {
	ftuo.mutation.ClearTransactionType()
	return ftuo
}

// Where appends a list predicates to the FinancialTransactionUpdate builder.
func (ftuo *FinancialTransactionUpdateOne) Where(ps ...predicate.FinancialTransaction) *FinancialTransactionUpdateOne {
	ftuo.mutation.Where(ps...)
	return ftuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ftuo *FinancialTransactionUpdateOne) Select(field string, fields ...string) *FinancialTransactionUpdateOne {
	ftuo.fields = append([]string{field}, fields...)
	return ftuo
}

// Save executes the query and returns the updated FinancialTransaction entity.
func (ftuo *FinancialTransactionUpdateOne) Save(ctx context.Context) (*FinancialTransaction, error) {
	ftuo.defaults()
	return withHooks(ctx, ftuo.sqlSave, ftuo.mutation, ftuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FinancialTransactionUpdateOne) SaveX(ctx context.Context) *FinancialTransaction {
	node, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ftuo *FinancialTransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FinancialTransactionUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftuo *FinancialTransactionUpdateOne) defaults() {
	if _, ok := ftuo.mutation.UpdatedAt(); !ok {
		v := financialtransaction.UpdateDefaultUpdatedAt()
		ftuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftuo *FinancialTransactionUpdateOne) check() error {
	if v, ok := ftuo.mutation.PaymentMethod(); ok {
		if err := financialtransaction.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "FinancialTransaction.payment_method": %w`, err)}
		}
	}
	if _, ok := ftuo.mutation.InvoiceID(); ftuo.mutation.InvoiceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FinancialTransaction.invoice"`)
	}
	if _, ok := ftuo.mutation.TransactionTypeID(); ftuo.mutation.TransactionTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FinancialTransaction.transaction_type"`)
	}
	return nil
}

func (ftuo *FinancialTransactionUpdateOne) sqlSave(ctx context.Context) (_node *FinancialTransaction, err error) {
	if err := ftuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(financialtransaction.Table, financialtransaction.Columns, sqlgraph.NewFieldSpec(financialtransaction.FieldID, field.TypeUUID))
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FinancialTransaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ftuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, financialtransaction.FieldID)
		for _, f := range fields {
			if !financialtransaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != financialtransaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ftuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftuo.mutation.UpdatedAt(); ok {
		_spec.SetField(financialtransaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ftuo.mutation.Amount(); ok {
		_spec.SetField(financialtransaction.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := ftuo.mutation.AddedAmount(); ok {
		_spec.AddField(financialtransaction.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := ftuo.mutation.Comment(); ok {
		_spec.SetField(financialtransaction.FieldComment, field.TypeString, value)
	}
	if ftuo.mutation.CommentCleared() {
		_spec.ClearField(financialtransaction.FieldComment, field.TypeString)
	}
	if value, ok := ftuo.mutation.PaymentMethod(); ok {
		_spec.SetField(financialtransaction.FieldPaymentMethod, field.TypeEnum, value)
	}
	if ftuo.mutation.InvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   financialtransaction.InvoiceTable,
			Columns: []string{financialtransaction.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   financialtransaction.InvoiceTable,
			Columns: []string{financialtransaction.InvoiceColumn},
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
	if ftuo.mutation.TransactionTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   financialtransaction.TransactionTypeTable,
			Columns: []string{financialtransaction.TransactionTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transactiontype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.TransactionTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   financialtransaction.TransactionTypeTable,
			Columns: []string{financialtransaction.TransactionTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transactiontype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FinancialTransaction{config: ftuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{financialtransaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ftuo.mutation.done = true
	return _node, nil
}
