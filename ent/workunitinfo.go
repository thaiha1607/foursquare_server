// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/address"
	"github.com/thaiha1607/foursquare_server/ent/workunitinfo"
)

// WorkUnitInfo is the model entity for the WorkUnitInfo schema.
type WorkUnitInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// AddressID holds the value of the "address_id" field.
	AddressID *uuid.UUID `json:"address_id,omitempty"`
	// Type holds the value of the "type" field.
	Type workunitinfo.Type `json:"type,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL *string `json:"image_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkUnitInfoQuery when eager-loading is set.
	Edges        WorkUnitInfoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WorkUnitInfoEdges holds the relations/edges for other nodes in the graph.
type WorkUnitInfoEdges struct {
	// Address holds the value of the address edge.
	Address *Address `json:"address,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AddressOrErr returns the Address value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkUnitInfoEdges) AddressOrErr() (*Address, error) {
	if e.Address != nil {
		return e.Address, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: address.Label}
	}
	return nil, &NotLoadedError{edge: "address"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkUnitInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workunitinfo.FieldAddressID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case workunitinfo.FieldName, workunitinfo.FieldType, workunitinfo.FieldImageURL:
			values[i] = new(sql.NullString)
		case workunitinfo.FieldCreatedAt, workunitinfo.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case workunitinfo.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkUnitInfo fields.
func (wui *WorkUnitInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workunitinfo.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				wui.ID = *value
			}
		case workunitinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wui.CreatedAt = value.Time
			}
		case workunitinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				wui.UpdatedAt = value.Time
			}
		case workunitinfo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				wui.Name = value.String
			}
		case workunitinfo.FieldAddressID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field address_id", values[i])
			} else if value.Valid {
				wui.AddressID = new(uuid.UUID)
				*wui.AddressID = *value.S.(*uuid.UUID)
			}
		case workunitinfo.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				wui.Type = workunitinfo.Type(value.String)
			}
		case workunitinfo.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				wui.ImageURL = new(string)
				*wui.ImageURL = value.String
			}
		default:
			wui.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WorkUnitInfo.
// This includes values selected through modifiers, order, etc.
func (wui *WorkUnitInfo) Value(name string) (ent.Value, error) {
	return wui.selectValues.Get(name)
}

// QueryAddress queries the "address" edge of the WorkUnitInfo entity.
func (wui *WorkUnitInfo) QueryAddress() *AddressQuery {
	return NewWorkUnitInfoClient(wui.config).QueryAddress(wui)
}

// Update returns a builder for updating this WorkUnitInfo.
// Note that you need to call WorkUnitInfo.Unwrap() before calling this method if this WorkUnitInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (wui *WorkUnitInfo) Update() *WorkUnitInfoUpdateOne {
	return NewWorkUnitInfoClient(wui.config).UpdateOne(wui)
}

// Unwrap unwraps the WorkUnitInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wui *WorkUnitInfo) Unwrap() *WorkUnitInfo {
	_tx, ok := wui.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkUnitInfo is not a transactional entity")
	}
	wui.config.driver = _tx.drv
	return wui
}

// String implements the fmt.Stringer.
func (wui *WorkUnitInfo) String() string {
	var builder strings.Builder
	builder.WriteString("WorkUnitInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wui.ID))
	builder.WriteString("created_at=")
	builder.WriteString(wui.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(wui.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(wui.Name)
	builder.WriteString(", ")
	if v := wui.AddressID; v != nil {
		builder.WriteString("address_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", wui.Type))
	builder.WriteString(", ")
	if v := wui.ImageURL; v != nil {
		builder.WriteString("image_url=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// WorkUnitInfos is a parsable slice of WorkUnitInfo.
type WorkUnitInfos []*WorkUnitInfo
