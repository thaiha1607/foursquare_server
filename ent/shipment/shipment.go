// Code generated by ent, DO NOT EDIT.

package shipment

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the shipment type in the database.
	Label = "shipment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldInvoiceID holds the string denoting the invoice_id field in the database.
	FieldInvoiceID = "invoice_id"
	// FieldStaffID holds the string denoting the staff_id field in the database.
	FieldStaffID = "staff_id"
	// FieldShipmentDate holds the string denoting the shipment_date field in the database.
	FieldShipmentDate = "shipment_date"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldStatusCode holds the string denoting the status_code field in the database.
	FieldStatusCode = "status_code"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// EdgeInvoice holds the string denoting the invoice edge name in mutations.
	EdgeInvoice = "invoice"
	// EdgeStaff holds the string denoting the staff edge name in mutations.
	EdgeStaff = "staff"
	// EdgeShipmentStatus holds the string denoting the shipment_status edge name in mutations.
	EdgeShipmentStatus = "shipment_status"
	// Table holds the table name of the shipment in the database.
	Table = "shipments"
	// OrderTable is the table that holds the order relation/edge.
	OrderTable = "shipments"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
	// OrderColumn is the table column denoting the order relation/edge.
	OrderColumn = "order_id"
	// InvoiceTable is the table that holds the invoice relation/edge.
	InvoiceTable = "shipments"
	// InvoiceInverseTable is the table name for the Invoice entity.
	// It exists in this package in order to avoid circular dependency with the "invoice" package.
	InvoiceInverseTable = "invoices"
	// InvoiceColumn is the table column denoting the invoice relation/edge.
	InvoiceColumn = "invoice_id"
	// StaffTable is the table that holds the staff relation/edge.
	StaffTable = "shipments"
	// StaffInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	StaffInverseTable = "persons"
	// StaffColumn is the table column denoting the staff relation/edge.
	StaffColumn = "staff_id"
	// ShipmentStatusTable is the table that holds the shipment_status relation/edge.
	ShipmentStatusTable = "shipments"
	// ShipmentStatusInverseTable is the table name for the ShipmentStatusCode entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentstatuscode" package.
	ShipmentStatusInverseTable = "shipment_status_codes"
	// ShipmentStatusColumn is the table column denoting the shipment_status relation/edge.
	ShipmentStatusColumn = "status_code"
)

// Columns holds all SQL columns for shipment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldOrderID,
	FieldInvoiceID,
	FieldStaffID,
	FieldShipmentDate,
	FieldNote,
	FieldStatusCode,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/thaiha1607/foursquare_server/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatusCode holds the default value on creation for the "status_code" field.
	DefaultStatusCode int
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Shipment queries.
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

// ByInvoiceID orders the results by the invoice_id field.
func ByInvoiceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceID, opts...).ToFunc()
}

// ByStaffID orders the results by the staff_id field.
func ByStaffID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStaffID, opts...).ToFunc()
}

// ByShipmentDate orders the results by the shipment_date field.
func ByShipmentDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShipmentDate, opts...).ToFunc()
}

// ByNote orders the results by the note field.
func ByNote(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNote, opts...).ToFunc()
}

// ByStatusCode orders the results by the status_code field.
func ByStatusCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatusCode, opts...).ToFunc()
}

// ByOrderField orders the results by order field.
func ByOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStep(), sql.OrderByField(field, opts...))
	}
}

// ByInvoiceField orders the results by invoice field.
func ByInvoiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvoiceStep(), sql.OrderByField(field, opts...))
	}
}

// ByStaffField orders the results by staff field.
func ByStaffField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStaffStep(), sql.OrderByField(field, opts...))
	}
}

// ByShipmentStatusField orders the results by shipment_status field.
func ByShipmentStatusField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShipmentStatusStep(), sql.OrderByField(field, opts...))
	}
}
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OrderTable, OrderColumn),
	)
}
func newInvoiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvoiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, InvoiceTable, InvoiceColumn),
	)
}
func newStaffStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StaffInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, StaffTable, StaffColumn),
	)
}
func newShipmentStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShipmentStatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ShipmentStatusTable, ShipmentStatusColumn),
	)
}