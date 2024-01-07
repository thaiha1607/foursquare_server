// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/account"
	"github.com/thaiha1607/foursquare_server/ent/user"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// LastReset holds the value of the "last_reset" field.
	LastReset *time.Time `json:"last_reset,omitempty"`
	// LastEmailVerification holds the value of the "last_email_verification" field.
	LastEmailVerification *time.Time `json:"last_email_verification,omitempty"`
	// LastPhoneVerification holds the value of the "last_phone_verification" field.
	LastPhoneVerification *time.Time `json:"last_phone_verification,omitempty"`
	// IsEmailVerified holds the value of the "is_email_verified" field.
	IsEmailVerified bool `json:"is_email_verified,omitempty"`
	// IsPhoneVerified holds the value of the "is_phone_verified" field.
	IsPhoneVerified bool `json:"is_phone_verified,omitempty"`
	// Role holds the value of the "role" field.
	Role account.Role `json:"role,omitempty"`
	// PasswordHash holds the value of the "password_hash" field.
	PasswordHash *string `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges        AccountEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldIsEmailVerified, account.FieldIsPhoneVerified:
			values[i] = new(sql.NullBool)
		case account.FieldID, account.FieldRole, account.FieldPasswordHash:
			values[i] = new(sql.NullString)
		case account.FieldCreatedAt, account.FieldUpdatedAt, account.FieldLastReset, account.FieldLastEmailVerification, account.FieldLastPhoneVerification:
			values[i] = new(sql.NullTime)
		case account.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case account.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				a.UserID = *value
			}
		case account.FieldLastReset:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_reset", values[i])
			} else if value.Valid {
				a.LastReset = new(time.Time)
				*a.LastReset = value.Time
			}
		case account.FieldLastEmailVerification:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_email_verification", values[i])
			} else if value.Valid {
				a.LastEmailVerification = new(time.Time)
				*a.LastEmailVerification = value.Time
			}
		case account.FieldLastPhoneVerification:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_phone_verification", values[i])
			} else if value.Valid {
				a.LastPhoneVerification = new(time.Time)
				*a.LastPhoneVerification = value.Time
			}
		case account.FieldIsEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_email_verified", values[i])
			} else if value.Valid {
				a.IsEmailVerified = value.Bool
			}
		case account.FieldIsPhoneVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_phone_verified", values[i])
			} else if value.Valid {
				a.IsPhoneVerified = value.Bool
			}
		case account.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				a.Role = account.Role(value.String)
			}
		case account.FieldPasswordHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_hash", values[i])
			} else if value.Valid {
				a.PasswordHash = new(string)
				*a.PasswordHash = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Account.
// This includes values selected through modifiers, order, etc.
func (a *Account) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Account entity.
func (a *Account) QueryUser() *UserQuery {
	return NewAccountClient(a.config).QueryUser(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return NewAccountClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteString(", ")
	if v := a.LastReset; v != nil {
		builder.WriteString("last_reset=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := a.LastEmailVerification; v != nil {
		builder.WriteString("last_email_verification=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := a.LastPhoneVerification; v != nil {
		builder.WriteString("last_phone_verification=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_email_verified=")
	builder.WriteString(fmt.Sprintf("%v", a.IsEmailVerified))
	builder.WriteString(", ")
	builder.WriteString("is_phone_verified=")
	builder.WriteString(fmt.Sprintf("%v", a.IsPhoneVerified))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", a.Role))
	builder.WriteString(", ")
	builder.WriteString("password_hash=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account
