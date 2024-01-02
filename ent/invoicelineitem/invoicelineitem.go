// Code generated by ent, DO NOT EDIT.

package invoicelineitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the invoicelineitem type in the database.
	Label = "invoice_line_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldInvoiceID holds the string denoting the invoice_id field in the database.
	FieldInvoiceID = "invoice_id"
	// FieldOrderLineItemID holds the string denoting the order_line_item_id field in the database.
	FieldOrderLineItemID = "order_line_item_id"
	// FieldQty holds the string denoting the qty field in the database.
	FieldQty = "qty"
	// FieldTotal holds the string denoting the total field in the database.
	FieldTotal = "total"
	// EdgeInvoice holds the string denoting the invoice edge name in mutations.
	EdgeInvoice = "invoice"
	// EdgeOrderLineItem holds the string denoting the order_line_item edge name in mutations.
	EdgeOrderLineItem = "order_line_item"
	// Table holds the table name of the invoicelineitem in the database.
	Table = "invoice_line_items"
	// InvoiceTable is the table that holds the invoice relation/edge.
	InvoiceTable = "invoice_line_items"
	// InvoiceInverseTable is the table name for the Invoice entity.
	// It exists in this package in order to avoid circular dependency with the "invoice" package.
	InvoiceInverseTable = "invoices"
	// InvoiceColumn is the table column denoting the invoice relation/edge.
	InvoiceColumn = "invoice_id"
	// OrderLineItemTable is the table that holds the order_line_item relation/edge.
	OrderLineItemTable = "invoice_line_items"
	// OrderLineItemInverseTable is the table name for the OrderLineItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderlineitem" package.
	OrderLineItemInverseTable = "order_line_items"
	// OrderLineItemColumn is the table column denoting the order_line_item relation/edge.
	OrderLineItemColumn = "order_line_item_id"
)

// Columns holds all SQL columns for invoicelineitem fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldInvoiceID,
	FieldOrderLineItemID,
	FieldQty,
	FieldTotal,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the InvoiceLineItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByInvoiceID orders the results by the invoice_id field.
func ByInvoiceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceID, opts...).ToFunc()
}

// ByOrderLineItemID orders the results by the order_line_item_id field.
func ByOrderLineItemID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderLineItemID, opts...).ToFunc()
}

// ByQty orders the results by the qty field.
func ByQty(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQty, opts...).ToFunc()
}

// ByTotal orders the results by the total field.
func ByTotal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotal, opts...).ToFunc()
}

// ByInvoiceField orders the results by invoice field.
func ByInvoiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvoiceStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrderLineItemField orders the results by order_line_item field.
func ByOrderLineItemField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderLineItemStep(), sql.OrderByField(field, opts...))
	}
}
func newInvoiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvoiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, InvoiceTable, InvoiceColumn),
	)
}
func newOrderLineItemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderLineItemInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OrderLineItemTable, OrderLineItemColumn),
	)
}
