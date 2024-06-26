// Code generated by ent, DO NOT EDIT.

package invoice

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the invoice type in the database.
	Label = "invoice"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldTotal holds the string denoting the total field in the database.
	FieldTotal = "total"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStatusCode holds the string denoting the status_code field in the database.
	FieldStatusCode = "status_code"
	// FieldPaymentMethod holds the string denoting the payment_method field in the database.
	FieldPaymentMethod = "payment_method"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// EdgeInvoiceStatus holds the string denoting the invoice_status edge name in mutations.
	EdgeInvoiceStatus = "invoice_status"
	// Table holds the table name of the invoice in the database.
	Table = "invoices"
	// OrderTable is the table that holds the order relation/edge.
	OrderTable = "invoices"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
	// OrderColumn is the table column denoting the order relation/edge.
	OrderColumn = "order_id"
	// InvoiceStatusTable is the table that holds the invoice_status relation/edge.
	InvoiceStatusTable = "invoices"
	// InvoiceStatusInverseTable is the table name for the InvoiceStatusCode entity.
	// It exists in this package in order to avoid circular dependency with the "invoicestatuscode" package.
	InvoiceStatusInverseTable = "invoice_status_codes"
	// InvoiceStatusColumn is the table column denoting the invoice_status relation/edge.
	InvoiceStatusColumn = "status_code"
)

// Columns holds all SQL columns for invoice fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldOrderID,
	FieldTotal,
	FieldNote,
	FieldType,
	FieldStatusCode,
	FieldPaymentMethod,
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
	// DefaultStatusCode holds the default value on creation for the "status_code" field.
	DefaultStatusCode int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// TypeProForma is the default value of the Type enum.
const DefaultType = TypeProForma

// Type values.
const (
	TypeProForma   Type = "PRO_FORMA"
	TypeRegular    Type = "REGULAR"
	TypePastDue    Type = "PAST_DUE"
	TypeInterim    Type = "INTERIM"
	TypeTimesheet  Type = "TIMESHEET"
	TypeFinal      Type = "FINAL"
	TypeCredit     Type = "CREDIT"
	TypeDebit      Type = "DEBIT"
	TypeMixed      Type = "MIXED"
	TypeCommercial Type = "COMMERCIAL"
	TypeRecurring  Type = "RECURRING"
	TypeOther      Type = "OTHER"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeProForma, TypeRegular, TypePastDue, TypeInterim, TypeTimesheet, TypeFinal, TypeCredit, TypeDebit, TypeMixed, TypeCommercial, TypeRecurring, TypeOther:
		return nil
	default:
		return fmt.Errorf("invoice: invalid enum value for type field: %q", _type)
	}
}

// PaymentMethod defines the type for the "payment_method" enum field.
type PaymentMethod string

// PaymentMethodCash is the default value of the PaymentMethod enum.
const DefaultPaymentMethod = PaymentMethodCash

// PaymentMethod values.
const (
	PaymentMethodCash        PaymentMethod = "CASH"
	PaymentMethodEFT         PaymentMethod = "ELECTRONIC_FUNDS_TRANSFER"
	PaymentMethodGiftCard    PaymentMethod = "GIFT_CARD"
	PaymentMethodCreditCard  PaymentMethod = "CREDIT_CARD"
	PaymentMethodDebitCard   PaymentMethod = "DEBIT_CARD"
	PaymentMethodPrepaidCard PaymentMethod = "PREPAID_CARD"
	PaymentMethodCheck       PaymentMethod = "CHECK"
	PaymentMethodOther       PaymentMethod = "OTHER"
)

func (pm PaymentMethod) String() string {
	return string(pm)
}

// PaymentMethodValidator is a validator for the "payment_method" field enum values. It is called by the builders before save.
func PaymentMethodValidator(pm PaymentMethod) error {
	switch pm {
	case PaymentMethodCash, PaymentMethodEFT, PaymentMethodGiftCard, PaymentMethodCreditCard, PaymentMethodDebitCard, PaymentMethodPrepaidCard, PaymentMethodCheck, PaymentMethodOther:
		return nil
	default:
		return fmt.Errorf("invoice: invalid enum value for payment_method field: %q", pm)
	}
}

// OrderOption defines the ordering options for the Invoice queries.
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

// ByOrderID orders the results by the order_id field.
func ByOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderID, opts...).ToFunc()
}

// ByTotal orders the results by the total field.
func ByTotal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotal, opts...).ToFunc()
}

// ByNote orders the results by the note field.
func ByNote(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNote, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByStatusCode orders the results by the status_code field.
func ByStatusCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatusCode, opts...).ToFunc()
}

// ByPaymentMethod orders the results by the payment_method field.
func ByPaymentMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentMethod, opts...).ToFunc()
}

// ByOrderField orders the results by order field.
func ByOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStep(), sql.OrderByField(field, opts...))
	}
}

// ByInvoiceStatusField orders the results by invoice_status field.
func ByInvoiceStatusField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvoiceStatusStep(), sql.OrderByField(field, opts...))
	}
}
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OrderTable, OrderColumn),
	)
}
func newInvoiceStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvoiceStatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, InvoiceStatusTable, InvoiceStatusColumn),
	)
}
