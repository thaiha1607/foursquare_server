// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/order"
)

// Invoice is the model entity for the Invoice schema.
type Invoice struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// Total holds the value of the "total" field.
	Total decimal.Decimal `json:"total,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment *string `json:"comment,omitempty"`
	// Note holds the value of the "note" field.
	Note *string `json:"note,omitempty"`
	// Type holds the value of the "type" field.
	Type invoice.Type `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status invoice.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InvoiceQuery when eager-loading is set.
	Edges        InvoiceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// InvoiceEdges holds the relations/edges for other nodes in the graph.
type InvoiceEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvoiceEdges) OrderOrErr() (*Order, error) {
	if e.Order != nil {
		return e.Order, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: order.Label}
	}
	return nil, &NotLoadedError{edge: "order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Invoice) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case invoice.FieldTotal:
			values[i] = new(decimal.Decimal)
		case invoice.FieldComment, invoice.FieldNote, invoice.FieldType, invoice.FieldStatus:
			values[i] = new(sql.NullString)
		case invoice.FieldCreatedAt, invoice.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case invoice.FieldID, invoice.FieldOrderID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Invoice fields.
func (i *Invoice) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case invoice.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case invoice.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case invoice.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case invoice.FieldOrderID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[j])
			} else if value != nil {
				i.OrderID = *value
			}
		case invoice.FieldTotal:
			if value, ok := values[j].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total", values[j])
			} else if value != nil {
				i.Total = *value
			}
		case invoice.FieldComment:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[j])
			} else if value.Valid {
				i.Comment = new(string)
				*i.Comment = value.String
			}
		case invoice.FieldNote:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[j])
			} else if value.Valid {
				i.Note = new(string)
				*i.Note = value.String
			}
		case invoice.FieldType:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[j])
			} else if value.Valid {
				i.Type = invoice.Type(value.String)
			}
		case invoice.FieldStatus:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[j])
			} else if value.Valid {
				i.Status = invoice.Status(value.String)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Invoice.
// This includes values selected through modifiers, order, etc.
func (i *Invoice) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryOrder queries the "order" edge of the Invoice entity.
func (i *Invoice) QueryOrder() *OrderQuery {
	return NewInvoiceClient(i.config).QueryOrder(i)
}

// Update returns a builder for updating this Invoice.
// Note that you need to call Invoice.Unwrap() before calling this method if this Invoice
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Invoice) Update() *InvoiceUpdateOne {
	return NewInvoiceClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Invoice entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Invoice) Unwrap() *Invoice {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Invoice is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Invoice) String() string {
	var builder strings.Builder
	builder.WriteString("Invoice(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", i.OrderID))
	builder.WriteString(", ")
	builder.WriteString("total=")
	builder.WriteString(fmt.Sprintf("%v", i.Total))
	builder.WriteString(", ")
	if v := i.Comment; v != nil {
		builder.WriteString("comment=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := i.Note; v != nil {
		builder.WriteString("note=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", i.Type))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", i.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Invoices is a parsable slice of Invoice.
type Invoices []*Invoice
