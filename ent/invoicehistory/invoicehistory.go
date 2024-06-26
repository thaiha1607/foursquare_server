// Code generated by ent, DO NOT EDIT.

package invoicehistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the invoicehistory type in the database.
	Label = "invoice_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldInvoiceID holds the string denoting the invoice_id field in the database.
	FieldInvoiceID = "invoice_id"
	// FieldPersonID holds the string denoting the person_id field in the database.
	FieldPersonID = "person_id"
	// FieldOldStatusCode holds the string denoting the old_status_code field in the database.
	FieldOldStatusCode = "old_status_code"
	// FieldNewStatusCode holds the string denoting the new_status_code field in the database.
	FieldNewStatusCode = "new_status_code"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeInvoice holds the string denoting the invoice edge name in mutations.
	EdgeInvoice = "invoice"
	// EdgePerson holds the string denoting the person edge name in mutations.
	EdgePerson = "person"
	// EdgeOldStatus holds the string denoting the old_status edge name in mutations.
	EdgeOldStatus = "old_status"
	// EdgeNewStatus holds the string denoting the new_status edge name in mutations.
	EdgeNewStatus = "new_status"
	// Table holds the table name of the invoicehistory in the database.
	Table = "invoice_history"
	// InvoiceTable is the table that holds the invoice relation/edge.
	InvoiceTable = "invoice_history"
	// InvoiceInverseTable is the table name for the Invoice entity.
	// It exists in this package in order to avoid circular dependency with the "invoice" package.
	InvoiceInverseTable = "invoices"
	// InvoiceColumn is the table column denoting the invoice relation/edge.
	InvoiceColumn = "invoice_id"
	// PersonTable is the table that holds the person relation/edge.
	PersonTable = "invoice_history"
	// PersonInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	PersonInverseTable = "persons"
	// PersonColumn is the table column denoting the person relation/edge.
	PersonColumn = "person_id"
	// OldStatusTable is the table that holds the old_status relation/edge.
	OldStatusTable = "invoice_history"
	// OldStatusInverseTable is the table name for the InvoiceStatusCode entity.
	// It exists in this package in order to avoid circular dependency with the "invoicestatuscode" package.
	OldStatusInverseTable = "invoice_status_codes"
	// OldStatusColumn is the table column denoting the old_status relation/edge.
	OldStatusColumn = "old_status_code"
	// NewStatusTable is the table that holds the new_status relation/edge.
	NewStatusTable = "invoice_history"
	// NewStatusInverseTable is the table name for the InvoiceStatusCode entity.
	// It exists in this package in order to avoid circular dependency with the "invoicestatuscode" package.
	NewStatusInverseTable = "invoice_status_codes"
	// NewStatusColumn is the table column denoting the new_status relation/edge.
	NewStatusColumn = "new_status_code"
)

// Columns holds all SQL columns for invoicehistory fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldInvoiceID,
	FieldPersonID,
	FieldOldStatusCode,
	FieldNewStatusCode,
	FieldDescription,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the InvoiceHistory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByInvoiceID orders the results by the invoice_id field.
func ByInvoiceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceID, opts...).ToFunc()
}

// ByPersonID orders the results by the person_id field.
func ByPersonID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPersonID, opts...).ToFunc()
}

// ByOldStatusCode orders the results by the old_status_code field.
func ByOldStatusCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOldStatusCode, opts...).ToFunc()
}

// ByNewStatusCode orders the results by the new_status_code field.
func ByNewStatusCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNewStatusCode, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByInvoiceField orders the results by invoice field.
func ByInvoiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvoiceStep(), sql.OrderByField(field, opts...))
	}
}

// ByPersonField orders the results by person field.
func ByPersonField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPersonStep(), sql.OrderByField(field, opts...))
	}
}

// ByOldStatusField orders the results by old_status field.
func ByOldStatusField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOldStatusStep(), sql.OrderByField(field, opts...))
	}
}

// ByNewStatusField orders the results by new_status field.
func ByNewStatusField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNewStatusStep(), sql.OrderByField(field, opts...))
	}
}
func newInvoiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvoiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, InvoiceTable, InvoiceColumn),
	)
}
func newPersonStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PersonInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PersonTable, PersonColumn),
	)
}
func newOldStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OldStatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OldStatusTable, OldStatusColumn),
	)
}
func newNewStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NewStatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, NewStatusTable, NewStatusColumn),
	)
}
