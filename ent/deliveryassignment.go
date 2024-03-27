// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/deliveryassignment"
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/shipment"
)

// DeliveryAssignment is the model entity for the DeliveryAssignment schema.
type DeliveryAssignment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ShipmentID holds the value of the "shipment_id" field.
	ShipmentID uuid.UUID `json:"shipment_id,omitempty"`
	// StaffID holds the value of the "staff_id" field.
	StaffID uuid.UUID `json:"staff_id,omitempty"`
	// Status holds the value of the "status" field.
	Status deliveryassignment.Status `json:"status,omitempty"`
	// Note holds the value of the "note" field.
	Note *string `json:"note,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeliveryAssignmentQuery when eager-loading is set.
	Edges        DeliveryAssignmentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DeliveryAssignmentEdges holds the relations/edges for other nodes in the graph.
type DeliveryAssignmentEdges struct {
	// Shipment holds the value of the shipment edge.
	Shipment *Shipment `json:"shipment,omitempty"`
	// Staff holds the value of the staff edge.
	Staff *Person `json:"staff,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ShipmentOrErr returns the Shipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeliveryAssignmentEdges) ShipmentOrErr() (*Shipment, error) {
	if e.Shipment != nil {
		return e.Shipment, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: shipment.Label}
	}
	return nil, &NotLoadedError{edge: "shipment"}
}

// StaffOrErr returns the Staff value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeliveryAssignmentEdges) StaffOrErr() (*Person, error) {
	if e.Staff != nil {
		return e.Staff, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: person.Label}
	}
	return nil, &NotLoadedError{edge: "staff"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeliveryAssignment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case deliveryassignment.FieldStatus, deliveryassignment.FieldNote:
			values[i] = new(sql.NullString)
		case deliveryassignment.FieldCreatedAt, deliveryassignment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case deliveryassignment.FieldID, deliveryassignment.FieldShipmentID, deliveryassignment.FieldStaffID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeliveryAssignment fields.
func (da *DeliveryAssignment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deliveryassignment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				da.ID = *value
			}
		case deliveryassignment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				da.CreatedAt = value.Time
			}
		case deliveryassignment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				da.UpdatedAt = value.Time
			}
		case deliveryassignment.FieldShipmentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field shipment_id", values[i])
			} else if value != nil {
				da.ShipmentID = *value
			}
		case deliveryassignment.FieldStaffID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field staff_id", values[i])
			} else if value != nil {
				da.StaffID = *value
			}
		case deliveryassignment.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				da.Status = deliveryassignment.Status(value.String)
			}
		case deliveryassignment.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				da.Note = new(string)
				*da.Note = value.String
			}
		default:
			da.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DeliveryAssignment.
// This includes values selected through modifiers, order, etc.
func (da *DeliveryAssignment) Value(name string) (ent.Value, error) {
	return da.selectValues.Get(name)
}

// QueryShipment queries the "shipment" edge of the DeliveryAssignment entity.
func (da *DeliveryAssignment) QueryShipment() *ShipmentQuery {
	return NewDeliveryAssignmentClient(da.config).QueryShipment(da)
}

// QueryStaff queries the "staff" edge of the DeliveryAssignment entity.
func (da *DeliveryAssignment) QueryStaff() *PersonQuery {
	return NewDeliveryAssignmentClient(da.config).QueryStaff(da)
}

// Update returns a builder for updating this DeliveryAssignment.
// Note that you need to call DeliveryAssignment.Unwrap() before calling this method if this DeliveryAssignment
// was returned from a transaction, and the transaction was committed or rolled back.
func (da *DeliveryAssignment) Update() *DeliveryAssignmentUpdateOne {
	return NewDeliveryAssignmentClient(da.config).UpdateOne(da)
}

// Unwrap unwraps the DeliveryAssignment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (da *DeliveryAssignment) Unwrap() *DeliveryAssignment {
	_tx, ok := da.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeliveryAssignment is not a transactional entity")
	}
	da.config.driver = _tx.drv
	return da
}

// String implements the fmt.Stringer.
func (da *DeliveryAssignment) String() string {
	var builder strings.Builder
	builder.WriteString("DeliveryAssignment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", da.ID))
	builder.WriteString("created_at=")
	builder.WriteString(da.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(da.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("shipment_id=")
	builder.WriteString(fmt.Sprintf("%v", da.ShipmentID))
	builder.WriteString(", ")
	builder.WriteString("staff_id=")
	builder.WriteString(fmt.Sprintf("%v", da.StaffID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", da.Status))
	builder.WriteString(", ")
	if v := da.Note; v != nil {
		builder.WriteString("note=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// DeliveryAssignments is a parsable slice of DeliveryAssignment.
type DeliveryAssignments []*DeliveryAssignment
