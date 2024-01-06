// Code generated by ent, DO NOT EDIT.

package order

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldParentOrderID holds the string denoting the parent_order_id field in the database.
	FieldParentOrderID = "parent_order_id"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStatusCode holds the string denoting the status_code field in the database.
	FieldStatusCode = "status_code"
	// FieldManagementStaffID holds the string denoting the management_staff_id field in the database.
	FieldManagementStaffID = "management_staff_id"
	// FieldWarehouseStaffID holds the string denoting the warehouse_staff_id field in the database.
	FieldWarehouseStaffID = "warehouse_staff_id"
	// FieldDeliveryStaffID holds the string denoting the delivery_staff_id field in the database.
	FieldDeliveryStaffID = "delivery_staff_id"
	// FieldInternalNote holds the string denoting the internal_note field in the database.
	FieldInternalNote = "internal_note"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeParentOrder holds the string denoting the parent_order edge name in mutations.
	EdgeParentOrder = "parent_order"
	// EdgeOrderStatus holds the string denoting the order_status edge name in mutations.
	EdgeOrderStatus = "order_status"
	// EdgeManagementStaff holds the string denoting the management_staff edge name in mutations.
	EdgeManagementStaff = "management_staff"
	// EdgeWarehouseStaff holds the string denoting the warehouse_staff edge name in mutations.
	EdgeWarehouseStaff = "warehouse_staff"
	// EdgeDeliveryStaff holds the string denoting the delivery_staff edge name in mutations.
	EdgeDeliveryStaff = "delivery_staff"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "orders"
	// CustomerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CustomerInverseTable = "users"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_id"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "orders"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "created_by"
	// ParentOrderTable is the table that holds the parent_order relation/edge.
	ParentOrderTable = "orders"
	// ParentOrderColumn is the table column denoting the parent_order relation/edge.
	ParentOrderColumn = "parent_order_id"
	// OrderStatusTable is the table that holds the order_status relation/edge.
	OrderStatusTable = "orders"
	// OrderStatusInverseTable is the table name for the OrderStatusCode entity.
	// It exists in this package in order to avoid circular dependency with the "orderstatuscode" package.
	OrderStatusInverseTable = "order_status_codes"
	// OrderStatusColumn is the table column denoting the order_status relation/edge.
	OrderStatusColumn = "status_code"
	// ManagementStaffTable is the table that holds the management_staff relation/edge.
	ManagementStaffTable = "orders"
	// ManagementStaffInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ManagementStaffInverseTable = "users"
	// ManagementStaffColumn is the table column denoting the management_staff relation/edge.
	ManagementStaffColumn = "management_staff_id"
	// WarehouseStaffTable is the table that holds the warehouse_staff relation/edge.
	WarehouseStaffTable = "orders"
	// WarehouseStaffInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	WarehouseStaffInverseTable = "users"
	// WarehouseStaffColumn is the table column denoting the warehouse_staff relation/edge.
	WarehouseStaffColumn = "warehouse_staff_id"
	// DeliveryStaffTable is the table that holds the delivery_staff relation/edge.
	DeliveryStaffTable = "orders"
	// DeliveryStaffInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	DeliveryStaffInverseTable = "users"
	// DeliveryStaffColumn is the table column denoting the delivery_staff relation/edge.
	DeliveryStaffColumn = "delivery_staff_id"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCustomerID,
	FieldNote,
	FieldCreatedBy,
	FieldParentOrderID,
	FieldPriority,
	FieldType,
	FieldStatusCode,
	FieldManagementStaffID,
	FieldWarehouseStaffID,
	FieldDeliveryStaffID,
	FieldInternalNote,
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
	// DefaultPriority holds the default value on creation for the "priority" field.
	DefaultPriority int
	// PriorityValidator is a validator for the "priority" field. It is called by the builders before save.
	PriorityValidator func(int) error
	// DefaultStatusCode holds the default value on creation for the "status_code" field.
	DefaultStatusCode int
)

// Type defines the type for the "type" enum field.
type Type string

// TypeSale is the default value of the Type enum.
const DefaultType = TypeSale

// Type values.
const (
	TypeSale     Type = "SALE"
	TypeReturn   Type = "RETURN"
	TypeExchange Type = "EXCHANGE"
	TypeTransfer Type = "TRANSFER"
	TypeInternal Type = "INTERNAL"
	TypeOther    Type = "OTHER"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeSale, TypeReturn, TypeExchange, TypeTransfer, TypeInternal, TypeOther:
		return nil
	default:
		return fmt.Errorf("order: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Order queries.
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

// ByCustomerID orders the results by the customer_id field.
func ByCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerID, opts...).ToFunc()
}

// ByNote orders the results by the note field.
func ByNote(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNote, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByParentOrderID orders the results by the parent_order_id field.
func ByParentOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentOrderID, opts...).ToFunc()
}

// ByPriority orders the results by the priority field.
func ByPriority(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPriority, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByStatusCode orders the results by the status_code field.
func ByStatusCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatusCode, opts...).ToFunc()
}

// ByManagementStaffID orders the results by the management_staff_id field.
func ByManagementStaffID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldManagementStaffID, opts...).ToFunc()
}

// ByWarehouseStaffID orders the results by the warehouse_staff_id field.
func ByWarehouseStaffID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWarehouseStaffID, opts...).ToFunc()
}

// ByDeliveryStaffID orders the results by the delivery_staff_id field.
func ByDeliveryStaffID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeliveryStaffID, opts...).ToFunc()
}

// ByInternalNote orders the results by the internal_note field.
func ByInternalNote(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInternalNote, opts...).ToFunc()
}

// ByCustomerField orders the results by customer field.
func ByCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), sql.OrderByField(field, opts...))
	}
}

// ByCreatorField orders the results by creator field.
func ByCreatorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatorStep(), sql.OrderByField(field, opts...))
	}
}

// ByParentOrderField orders the results by parent_order field.
func ByParentOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentOrderStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrderStatusField orders the results by order_status field.
func ByOrderStatusField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStatusStep(), sql.OrderByField(field, opts...))
	}
}

// ByManagementStaffField orders the results by management_staff field.
func ByManagementStaffField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newManagementStaffStep(), sql.OrderByField(field, opts...))
	}
}

// ByWarehouseStaffField orders the results by warehouse_staff field.
func ByWarehouseStaffField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWarehouseStaffStep(), sql.OrderByField(field, opts...))
	}
}

// ByDeliveryStaffField orders the results by delivery_staff field.
func ByDeliveryStaffField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeliveryStaffStep(), sql.OrderByField(field, opts...))
	}
}
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CustomerTable, CustomerColumn),
	)
}
func newCreatorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CreatorTable, CreatorColumn),
	)
}
func newParentOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ParentOrderTable, ParentOrderColumn),
	)
}
func newOrderStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderStatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OrderStatusTable, OrderStatusColumn),
	)
}
func newManagementStaffStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ManagementStaffInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ManagementStaffTable, ManagementStaffColumn),
	)
}
func newWarehouseStaffStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WarehouseStaffInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, WarehouseStaffTable, WarehouseStaffColumn),
	)
}
func newDeliveryStaffStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeliveryStaffInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DeliveryStaffTable, DeliveryStaffColumn),
	)
}
