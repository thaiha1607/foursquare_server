// Code generated by ent, DO NOT EDIT.

package conversation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the conversation type in the database.
	Label = "conversation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldUserOneID holds the string denoting the user_one_id field in the database.
	FieldUserOneID = "user_one_id"
	// FieldUserTwoID holds the string denoting the user_two_id field in the database.
	FieldUserTwoID = "user_two_id"
	// EdgeUserOne holds the string denoting the user_one edge name in mutations.
	EdgeUserOne = "user_one"
	// EdgeUserTwo holds the string denoting the user_two edge name in mutations.
	EdgeUserTwo = "user_two"
	// Table holds the table name of the conversation in the database.
	Table = "conversations"
	// UserOneTable is the table that holds the user_one relation/edge.
	UserOneTable = "conversations"
	// UserOneInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserOneInverseTable = "users"
	// UserOneColumn is the table column denoting the user_one relation/edge.
	UserOneColumn = "user_one_id"
	// UserTwoTable is the table that holds the user_two relation/edge.
	UserTwoTable = "conversations"
	// UserTwoInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserTwoInverseTable = "users"
	// UserTwoColumn is the table column denoting the user_two relation/edge.
	UserTwoColumn = "user_two_id"
)

// Columns holds all SQL columns for conversation fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTitle,
	FieldUserOneID,
	FieldUserTwoID,
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

// OrderOption defines the ordering options for the Conversation queries.
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

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByUserOneID orders the results by the user_one_id field.
func ByUserOneID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserOneID, opts...).ToFunc()
}

// ByUserTwoID orders the results by the user_two_id field.
func ByUserTwoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserTwoID, opts...).ToFunc()
}

// ByUserOneField orders the results by user_one field.
func ByUserOneField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserOneStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserTwoField orders the results by user_two field.
func ByUserTwoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserTwoStep(), sql.OrderByField(field, opts...))
	}
}
func newUserOneStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserOneInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserOneTable, UserOneColumn),
	)
}
func newUserTwoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserTwoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTwoTable, UserTwoColumn),
	)
}
