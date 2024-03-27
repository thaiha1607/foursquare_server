// Code generated by ent, DO NOT EDIT.

package productqty

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the productqty type in the database.
	Label = "product_qty"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldWorkUnitID holds the string denoting the work_unit_id field in the database.
	FieldWorkUnitID = "work_unit_id"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldProductColorID holds the string denoting the product_color_id field in the database.
	FieldProductColorID = "product_color_id"
	// FieldPricePerUnit holds the string denoting the price_per_unit field in the database.
	FieldPricePerUnit = "price_per_unit"
	// FieldQty holds the string denoting the qty field in the database.
	FieldQty = "qty"
	// EdgeWorkUnit holds the string denoting the work_unit edge name in mutations.
	EdgeWorkUnit = "work_unit"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeProductColor holds the string denoting the product_color edge name in mutations.
	EdgeProductColor = "product_color"
	// Table holds the table name of the productqty in the database.
	Table = "product_qty"
	// WorkUnitTable is the table that holds the work_unit relation/edge.
	WorkUnitTable = "product_qty"
	// WorkUnitInverseTable is the table name for the WorkUnitInfo entity.
	// It exists in this package in order to avoid circular dependency with the "workunitinfo" package.
	WorkUnitInverseTable = "work_unit_info"
	// WorkUnitColumn is the table column denoting the work_unit relation/edge.
	WorkUnitColumn = "work_unit_id"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "product_qty"
	// ProductInverseTable is the table name for the ProductInfo entity.
	// It exists in this package in order to avoid circular dependency with the "productinfo" package.
	ProductInverseTable = "product_info"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "product_id"
	// ProductColorTable is the table that holds the product_color relation/edge.
	ProductColorTable = "product_qty"
	// ProductColorInverseTable is the table name for the ProductColor entity.
	// It exists in this package in order to avoid circular dependency with the "productcolor" package.
	ProductColorInverseTable = "product_color"
	// ProductColorColumn is the table column denoting the product_color relation/edge.
	ProductColorColumn = "product_color_id"
)

// Columns holds all SQL columns for productqty fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldWorkUnitID,
	FieldProductID,
	FieldProductColorID,
	FieldPricePerUnit,
	FieldQty,
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
	// ProductIDValidator is a validator for the "product_id" field. It is called by the builders before save.
	ProductIDValidator func(string) error
	// ProductColorIDValidator is a validator for the "product_color_id" field. It is called by the builders before save.
	ProductColorIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the ProductQty queries.
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

// ByWorkUnitID orders the results by the work_unit_id field.
func ByWorkUnitID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkUnitID, opts...).ToFunc()
}

// ByProductID orders the results by the product_id field.
func ByProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductID, opts...).ToFunc()
}

// ByProductColorID orders the results by the product_color_id field.
func ByProductColorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductColorID, opts...).ToFunc()
}

// ByPricePerUnit orders the results by the price_per_unit field.
func ByPricePerUnit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPricePerUnit, opts...).ToFunc()
}

// ByQty orders the results by the qty field.
func ByQty(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQty, opts...).ToFunc()
}

// ByWorkUnitField orders the results by work_unit field.
func ByWorkUnitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkUnitStep(), sql.OrderByField(field, opts...))
	}
}

// ByProductField orders the results by product field.
func ByProductField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductStep(), sql.OrderByField(field, opts...))
	}
}

// ByProductColorField orders the results by product_color field.
func ByProductColorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductColorStep(), sql.OrderByField(field, opts...))
	}
}
func newWorkUnitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkUnitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, WorkUnitTable, WorkUnitColumn),
	)
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProductTable, ProductColumn),
	)
}
func newProductColorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductColorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProductColorTable, ProductColorColumn),
	)
}
