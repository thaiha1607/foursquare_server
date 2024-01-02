// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/thaiha1607/foursquare_server/ent/producttype"
)

// ProductType is the model entity for the ProductType schema.
type ProductType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ProductType holds the value of the "product_type" field.
	ProductType  string `json:"product_type,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case producttype.FieldID:
			values[i] = new(sql.NullInt64)
		case producttype.FieldProductType:
			values[i] = new(sql.NullString)
		case producttype.FieldCreatedAt, producttype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductType fields.
func (pt *ProductType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case producttype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pt.ID = int(value.Int64)
		case producttype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pt.CreatedAt = value.Time
			}
		case producttype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pt.UpdatedAt = value.Time
			}
		case producttype.FieldProductType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_type", values[i])
			} else if value.Valid {
				pt.ProductType = value.String
			}
		default:
			pt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductType.
// This includes values selected through modifiers, order, etc.
func (pt *ProductType) Value(name string) (ent.Value, error) {
	return pt.selectValues.Get(name)
}

// Update returns a builder for updating this ProductType.
// Note that you need to call ProductType.Unwrap() before calling this method if this ProductType
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *ProductType) Update() *ProductTypeUpdateOne {
	return NewProductTypeClient(pt.config).UpdateOne(pt)
}

// Unwrap unwraps the ProductType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *ProductType) Unwrap() *ProductType {
	_tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductType is not a transactional entity")
	}
	pt.config.driver = _tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *ProductType) String() string {
	var builder strings.Builder
	builder.WriteString("ProductType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pt.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("product_type=")
	builder.WriteString(pt.ProductType)
	builder.WriteByte(')')
	return builder.String()
}

// ProductTypes is a parsable slice of ProductType.
type ProductTypes []*ProductType
