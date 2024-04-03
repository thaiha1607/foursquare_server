// Code generated by ent, DO NOT EDIT.

package shipmenthistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the shipmenthistory type in the database.
	Label = "shipment_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldShipmentID holds the string denoting the shipment_id field in the database.
	FieldShipmentID = "shipment_id"
	// FieldPersonID holds the string denoting the person_id field in the database.
	FieldPersonID = "person_id"
	// FieldOldStatusCode holds the string denoting the old_status_code field in the database.
	FieldOldStatusCode = "old_status_code"
	// FieldNewStatusCode holds the string denoting the new_status_code field in the database.
	FieldNewStatusCode = "new_status_code"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeShipment holds the string denoting the shipment edge name in mutations.
	EdgeShipment = "shipment"
	// EdgePerson holds the string denoting the person edge name in mutations.
	EdgePerson = "person"
	// EdgeOldStatus holds the string denoting the old_status edge name in mutations.
	EdgeOldStatus = "old_status"
	// EdgeNewStatus holds the string denoting the new_status edge name in mutations.
	EdgeNewStatus = "new_status"
	// Table holds the table name of the shipmenthistory in the database.
	Table = "shipment_history"
	// ShipmentTable is the table that holds the shipment relation/edge.
	ShipmentTable = "shipment_history"
	// ShipmentInverseTable is the table name for the Shipment entity.
	// It exists in this package in order to avoid circular dependency with the "shipment" package.
	ShipmentInverseTable = "shipments"
	// ShipmentColumn is the table column denoting the shipment relation/edge.
	ShipmentColumn = "shipment_id"
	// PersonTable is the table that holds the person relation/edge.
	PersonTable = "shipment_history"
	// PersonInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	PersonInverseTable = "persons"
	// PersonColumn is the table column denoting the person relation/edge.
	PersonColumn = "person_id"
	// OldStatusTable is the table that holds the old_status relation/edge.
	OldStatusTable = "shipment_history"
	// OldStatusInverseTable is the table name for the ShipmentStatusCode entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentstatuscode" package.
	OldStatusInverseTable = "shipment_status_codes"
	// OldStatusColumn is the table column denoting the old_status relation/edge.
	OldStatusColumn = "old_status_code"
	// NewStatusTable is the table that holds the new_status relation/edge.
	NewStatusTable = "shipment_history"
	// NewStatusInverseTable is the table name for the ShipmentStatusCode entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentstatuscode" package.
	NewStatusInverseTable = "shipment_status_codes"
	// NewStatusColumn is the table column denoting the new_status relation/edge.
	NewStatusColumn = "new_status_code"
)

// Columns holds all SQL columns for shipmenthistory fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldShipmentID,
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
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// ShipmentIDValidator is a validator for the "shipment_id" field. It is called by the builders before save.
	ShipmentIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the ShipmentHistory queries.
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

// ByShipmentID orders the results by the shipment_id field.
func ByShipmentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShipmentID, opts...).ToFunc()
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

// ByShipmentField orders the results by shipment field.
func ByShipmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShipmentStep(), sql.OrderByField(field, opts...))
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
func newShipmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShipmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ShipmentTable, ShipmentColumn),
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
