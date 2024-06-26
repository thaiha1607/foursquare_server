// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/invoicehistory"
	"github.com/thaiha1607/foursquare_server/ent/invoicestatuscode"
	"github.com/thaiha1607/foursquare_server/ent/person"
)

// InvoiceHistory is the model entity for the InvoiceHistory schema.
type InvoiceHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// InvoiceID holds the value of the "invoice_id" field.
	InvoiceID uuid.UUID `json:"invoice_id,omitempty"`
	// PersonID holds the value of the "person_id" field.
	PersonID uuid.UUID `json:"person_id,omitempty"`
	// OldStatusCode holds the value of the "old_status_code" field.
	OldStatusCode *int `json:"old_status_code,omitempty"`
	// NewStatusCode holds the value of the "new_status_code" field.
	NewStatusCode *int `json:"new_status_code,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InvoiceHistoryQuery when eager-loading is set.
	Edges        InvoiceHistoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// InvoiceHistoryEdges holds the relations/edges for other nodes in the graph.
type InvoiceHistoryEdges struct {
	// Invoice holds the value of the invoice edge.
	Invoice *Invoice `json:"invoice,omitempty"`
	// Person holds the value of the person edge.
	Person *Person `json:"person,omitempty"`
	// OldStatus holds the value of the old_status edge.
	OldStatus *InvoiceStatusCode `json:"old_status,omitempty"`
	// NewStatus holds the value of the new_status edge.
	NewStatus *InvoiceStatusCode `json:"new_status,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// InvoiceOrErr returns the Invoice value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvoiceHistoryEdges) InvoiceOrErr() (*Invoice, error) {
	if e.Invoice != nil {
		return e.Invoice, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: invoice.Label}
	}
	return nil, &NotLoadedError{edge: "invoice"}
}

// PersonOrErr returns the Person value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvoiceHistoryEdges) PersonOrErr() (*Person, error) {
	if e.Person != nil {
		return e.Person, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: person.Label}
	}
	return nil, &NotLoadedError{edge: "person"}
}

// OldStatusOrErr returns the OldStatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvoiceHistoryEdges) OldStatusOrErr() (*InvoiceStatusCode, error) {
	if e.OldStatus != nil {
		return e.OldStatus, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: invoicestatuscode.Label}
	}
	return nil, &NotLoadedError{edge: "old_status"}
}

// NewStatusOrErr returns the NewStatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvoiceHistoryEdges) NewStatusOrErr() (*InvoiceStatusCode, error) {
	if e.NewStatus != nil {
		return e.NewStatus, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: invoicestatuscode.Label}
	}
	return nil, &NotLoadedError{edge: "new_status"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InvoiceHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case invoicehistory.FieldOldStatusCode, invoicehistory.FieldNewStatusCode:
			values[i] = new(sql.NullInt64)
		case invoicehistory.FieldDescription:
			values[i] = new(sql.NullString)
		case invoicehistory.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case invoicehistory.FieldID, invoicehistory.FieldInvoiceID, invoicehistory.FieldPersonID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InvoiceHistory fields.
func (ih *InvoiceHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case invoicehistory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ih.ID = *value
			}
		case invoicehistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ih.CreatedAt = new(time.Time)
				*ih.CreatedAt = value.Time
			}
		case invoicehistory.FieldInvoiceID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_id", values[i])
			} else if value != nil {
				ih.InvoiceID = *value
			}
		case invoicehistory.FieldPersonID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field person_id", values[i])
			} else if value != nil {
				ih.PersonID = *value
			}
		case invoicehistory.FieldOldStatusCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field old_status_code", values[i])
			} else if value.Valid {
				ih.OldStatusCode = new(int)
				*ih.OldStatusCode = int(value.Int64)
			}
		case invoicehistory.FieldNewStatusCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field new_status_code", values[i])
			} else if value.Valid {
				ih.NewStatusCode = new(int)
				*ih.NewStatusCode = int(value.Int64)
			}
		case invoicehistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ih.Description = new(string)
				*ih.Description = value.String
			}
		default:
			ih.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the InvoiceHistory.
// This includes values selected through modifiers, order, etc.
func (ih *InvoiceHistory) Value(name string) (ent.Value, error) {
	return ih.selectValues.Get(name)
}

// QueryInvoice queries the "invoice" edge of the InvoiceHistory entity.
func (ih *InvoiceHistory) QueryInvoice() *InvoiceQuery {
	return NewInvoiceHistoryClient(ih.config).QueryInvoice(ih)
}

// QueryPerson queries the "person" edge of the InvoiceHistory entity.
func (ih *InvoiceHistory) QueryPerson() *PersonQuery {
	return NewInvoiceHistoryClient(ih.config).QueryPerson(ih)
}

// QueryOldStatus queries the "old_status" edge of the InvoiceHistory entity.
func (ih *InvoiceHistory) QueryOldStatus() *InvoiceStatusCodeQuery {
	return NewInvoiceHistoryClient(ih.config).QueryOldStatus(ih)
}

// QueryNewStatus queries the "new_status" edge of the InvoiceHistory entity.
func (ih *InvoiceHistory) QueryNewStatus() *InvoiceStatusCodeQuery {
	return NewInvoiceHistoryClient(ih.config).QueryNewStatus(ih)
}

// Update returns a builder for updating this InvoiceHistory.
// Note that you need to call InvoiceHistory.Unwrap() before calling this method if this InvoiceHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ih *InvoiceHistory) Update() *InvoiceHistoryUpdateOne {
	return NewInvoiceHistoryClient(ih.config).UpdateOne(ih)
}

// Unwrap unwraps the InvoiceHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ih *InvoiceHistory) Unwrap() *InvoiceHistory {
	_tx, ok := ih.config.driver.(*txDriver)
	if !ok {
		panic("ent: InvoiceHistory is not a transactional entity")
	}
	ih.config.driver = _tx.drv
	return ih
}

// String implements the fmt.Stringer.
func (ih *InvoiceHistory) String() string {
	var builder strings.Builder
	builder.WriteString("InvoiceHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ih.ID))
	if v := ih.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("invoice_id=")
	builder.WriteString(fmt.Sprintf("%v", ih.InvoiceID))
	builder.WriteString(", ")
	builder.WriteString("person_id=")
	builder.WriteString(fmt.Sprintf("%v", ih.PersonID))
	builder.WriteString(", ")
	if v := ih.OldStatusCode; v != nil {
		builder.WriteString("old_status_code=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ih.NewStatusCode; v != nil {
		builder.WriteString("new_status_code=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ih.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// InvoiceHistories is a parsable slice of InvoiceHistory.
type InvoiceHistories []*InvoiceHistory
