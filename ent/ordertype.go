// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/thaiha1607/foursquare_server/ent/ordertype"
)

// OrderType is the model entity for the OrderType schema.
type OrderType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// OrderType holds the value of the "order_type" field.
	OrderType    string `json:"order_type,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ordertype.FieldID:
			values[i] = new(sql.NullInt64)
		case ordertype.FieldOrderType:
			values[i] = new(sql.NullString)
		case ordertype.FieldCreatedAt, ordertype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderType fields.
func (ot *OrderType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ordertype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ot.ID = int(value.Int64)
		case ordertype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ot.CreatedAt = value.Time
			}
		case ordertype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ot.UpdatedAt = value.Time
			}
		case ordertype.FieldOrderType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_type", values[i])
			} else if value.Valid {
				ot.OrderType = value.String
			}
		default:
			ot.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderType.
// This includes values selected through modifiers, order, etc.
func (ot *OrderType) Value(name string) (ent.Value, error) {
	return ot.selectValues.Get(name)
}

// Update returns a builder for updating this OrderType.
// Note that you need to call OrderType.Unwrap() before calling this method if this OrderType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ot *OrderType) Update() *OrderTypeUpdateOne {
	return NewOrderTypeClient(ot.config).UpdateOne(ot)
}

// Unwrap unwraps the OrderType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ot *OrderType) Unwrap() *OrderType {
	_tx, ok := ot.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderType is not a transactional entity")
	}
	ot.config.driver = _tx.drv
	return ot
}

// String implements the fmt.Stringer.
func (ot *OrderType) String() string {
	var builder strings.Builder
	builder.WriteString("OrderType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ot.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ot.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ot.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("order_type=")
	builder.WriteString(ot.OrderType)
	builder.WriteByte(')')
	return builder.String()
}

// OrderTypes is a parsable slice of OrderType.
type OrderTypes []*OrderType
