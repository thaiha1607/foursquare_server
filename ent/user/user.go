// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAvatarURL holds the string denoting the avatar_url field in the database.
	FieldAvatarURL = "avatar_url"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldLastReset holds the string denoting the last_reset field in the database.
	FieldLastReset = "last_reset"
	// FieldLastVerification holds the string denoting the last_verification field in the database.
	FieldLastVerification = "last_verification"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldVerified holds the string denoting the verified field in the database.
	FieldVerified = "verified"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// FieldOtherAddressInfo holds the string denoting the other_address_info field in the database.
	FieldOtherAddressInfo = "other_address_info"
	// EdgeUserRole holds the string denoting the user_role edge name in mutations.
	EdgeUserRole = "user_role"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UserRoleTable is the table that holds the user_role relation/edge.
	UserRoleTable = "users"
	// UserRoleInverseTable is the table name for the UserRole entity.
	// It exists in this package in order to avoid circular dependency with the "userrole" package.
	UserRoleInverseTable = "user_roles"
	// UserRoleColumn is the table column denoting the user_role relation/edge.
	UserRoleColumn = "role"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAvatarURL,
	FieldEmail,
	FieldLastReset,
	FieldLastVerification,
	FieldName,
	FieldPasswordHash,
	FieldUsername,
	FieldVerified,
	FieldPhone,
	FieldRole,
	FieldAddress,
	FieldPostalCode,
	FieldOtherAddressInfo,
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
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	PasswordHashValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DefaultVerified holds the default value on creation for the "verified" field.
	DefaultVerified bool
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
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

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByLastReset orders the results by the last_reset field.
func ByLastReset(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastReset, opts...).ToFunc()
}

// ByLastVerification orders the results by the last_verification field.
func ByLastVerification(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastVerification, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPasswordHash orders the results by the password_hash field.
func ByPasswordHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByVerified orders the results by the verified field.
func ByVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVerified, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByPostalCode orders the results by the postal_code field.
func ByPostalCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalCode, opts...).ToFunc()
}

// ByOtherAddressInfo orders the results by the other_address_info field.
func ByOtherAddressInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOtherAddressInfo, opts...).ToFunc()
}

// ByUserRoleField orders the results by user_role field.
func ByUserRoleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserRoleStep(), sql.OrderByField(field, opts...))
	}
}
func newUserRoleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserRoleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserRoleTable, UserRoleColumn),
	)
}
