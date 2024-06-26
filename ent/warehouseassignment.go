// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/order"
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/warehouseassignment"
	"github.com/thaiha1607/foursquare_server/ent/workunitinfo"
)

// WarehouseAssignment is the model entity for the WarehouseAssignment schema.
type WarehouseAssignment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// WorkUnitID holds the value of the "work_unit_id" field.
	WorkUnitID uuid.UUID `json:"work_unit_id,omitempty"`
	// StaffID holds the value of the "staff_id" field.
	StaffID *uuid.UUID `json:"staff_id,omitempty"`
	// Status holds the value of the "status" field.
	Status *warehouseassignment.Status `json:"status,omitempty"`
	// Note holds the value of the "note" field.
	Note *string `json:"note,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WarehouseAssignmentQuery when eager-loading is set.
	Edges        WarehouseAssignmentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WarehouseAssignmentEdges holds the relations/edges for other nodes in the graph.
type WarehouseAssignmentEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// WorkUnit holds the value of the work_unit edge.
	WorkUnit *WorkUnitInfo `json:"work_unit,omitempty"`
	// Staff holds the value of the staff edge.
	Staff *Person `json:"staff,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WarehouseAssignmentEdges) OrderOrErr() (*Order, error) {
	if e.Order != nil {
		return e.Order, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: order.Label}
	}
	return nil, &NotLoadedError{edge: "order"}
}

// WorkUnitOrErr returns the WorkUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WarehouseAssignmentEdges) WorkUnitOrErr() (*WorkUnitInfo, error) {
	if e.WorkUnit != nil {
		return e.WorkUnit, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: workunitinfo.Label}
	}
	return nil, &NotLoadedError{edge: "work_unit"}
}

// StaffOrErr returns the Staff value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WarehouseAssignmentEdges) StaffOrErr() (*Person, error) {
	if e.Staff != nil {
		return e.Staff, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: person.Label}
	}
	return nil, &NotLoadedError{edge: "staff"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WarehouseAssignment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case warehouseassignment.FieldStaffID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case warehouseassignment.FieldStatus, warehouseassignment.FieldNote:
			values[i] = new(sql.NullString)
		case warehouseassignment.FieldCreatedAt, warehouseassignment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case warehouseassignment.FieldID, warehouseassignment.FieldOrderID, warehouseassignment.FieldWorkUnitID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WarehouseAssignment fields.
func (wa *WarehouseAssignment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case warehouseassignment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				wa.ID = *value
			}
		case warehouseassignment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wa.CreatedAt = new(time.Time)
				*wa.CreatedAt = value.Time
			}
		case warehouseassignment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				wa.UpdatedAt = new(time.Time)
				*wa.UpdatedAt = value.Time
			}
		case warehouseassignment.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				wa.OrderID = *value
			}
		case warehouseassignment.FieldWorkUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field work_unit_id", values[i])
			} else if value != nil {
				wa.WorkUnitID = *value
			}
		case warehouseassignment.FieldStaffID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field staff_id", values[i])
			} else if value.Valid {
				wa.StaffID = new(uuid.UUID)
				*wa.StaffID = *value.S.(*uuid.UUID)
			}
		case warehouseassignment.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				wa.Status = new(warehouseassignment.Status)
				*wa.Status = warehouseassignment.Status(value.String)
			}
		case warehouseassignment.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				wa.Note = new(string)
				*wa.Note = value.String
			}
		default:
			wa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WarehouseAssignment.
// This includes values selected through modifiers, order, etc.
func (wa *WarehouseAssignment) Value(name string) (ent.Value, error) {
	return wa.selectValues.Get(name)
}

// QueryOrder queries the "order" edge of the WarehouseAssignment entity.
func (wa *WarehouseAssignment) QueryOrder() *OrderQuery {
	return NewWarehouseAssignmentClient(wa.config).QueryOrder(wa)
}

// QueryWorkUnit queries the "work_unit" edge of the WarehouseAssignment entity.
func (wa *WarehouseAssignment) QueryWorkUnit() *WorkUnitInfoQuery {
	return NewWarehouseAssignmentClient(wa.config).QueryWorkUnit(wa)
}

// QueryStaff queries the "staff" edge of the WarehouseAssignment entity.
func (wa *WarehouseAssignment) QueryStaff() *PersonQuery {
	return NewWarehouseAssignmentClient(wa.config).QueryStaff(wa)
}

// Update returns a builder for updating this WarehouseAssignment.
// Note that you need to call WarehouseAssignment.Unwrap() before calling this method if this WarehouseAssignment
// was returned from a transaction, and the transaction was committed or rolled back.
func (wa *WarehouseAssignment) Update() *WarehouseAssignmentUpdateOne {
	return NewWarehouseAssignmentClient(wa.config).UpdateOne(wa)
}

// Unwrap unwraps the WarehouseAssignment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wa *WarehouseAssignment) Unwrap() *WarehouseAssignment {
	_tx, ok := wa.config.driver.(*txDriver)
	if !ok {
		panic("ent: WarehouseAssignment is not a transactional entity")
	}
	wa.config.driver = _tx.drv
	return wa
}

// String implements the fmt.Stringer.
func (wa *WarehouseAssignment) String() string {
	var builder strings.Builder
	builder.WriteString("WarehouseAssignment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wa.ID))
	if v := wa.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := wa.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", wa.OrderID))
	builder.WriteString(", ")
	builder.WriteString("work_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", wa.WorkUnitID))
	builder.WriteString(", ")
	if v := wa.StaffID; v != nil {
		builder.WriteString("staff_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := wa.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := wa.Note; v != nil {
		builder.WriteString("note=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// WarehouseAssignments is a parsable slice of WarehouseAssignment.
type WarehouseAssignments []*WarehouseAssignment
