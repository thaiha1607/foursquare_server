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
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/personaddress"
)

// PersonAddress is the model entity for the PersonAddress schema.
type PersonAddress struct {
	config `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// PersonID holds the value of the "person_id" field.
	PersonID uuid.UUID `json:"person_id,omitempty"`
	// AddressID holds the value of the "address_id" field.
	AddressID uuid.UUID `json:"address_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonAddressQuery when eager-loading is set.
	Edges        PersonAddressEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PersonAddressEdges holds the relations/edges for other nodes in the graph.
type PersonAddressEdges struct {
	// Persons holds the value of the persons edge.
	Persons *Person `json:"persons,omitempty"`
	// Addresses holds the value of the addresses edge.
	Addresses *Address `json:"addresses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PersonsOrErr returns the Persons value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonAddressEdges) PersonsOrErr() (*Person, error) {
	if e.Persons != nil {
		return e.Persons, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: person.Label}
	}
	return nil, &NotLoadedError{edge: "persons"}
}

// AddressesOrErr returns the Addresses value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonAddressEdges) AddressesOrErr() (*Address, error) {
	if e.Addresses != nil {
		return e.Addresses, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: address.Label}
	}
	return nil, &NotLoadedError{edge: "addresses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PersonAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case personaddress.FieldCreatedAt, personaddress.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case personaddress.FieldPersonID, personaddress.FieldAddressID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PersonAddress fields.
func (pa *PersonAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case personaddress.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = new(time.Time)
				*pa.CreatedAt = value.Time
			}
		case personaddress.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pa.UpdatedAt = new(time.Time)
				*pa.UpdatedAt = value.Time
			}
		case personaddress.FieldPersonID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field person_id", values[i])
			} else if value != nil {
				pa.PersonID = *value
			}
		case personaddress.FieldAddressID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field address_id", values[i])
			} else if value != nil {
				pa.AddressID = *value
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PersonAddress.
// This includes values selected through modifiers, order, etc.
func (pa *PersonAddress) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// QueryPersons queries the "persons" edge of the PersonAddress entity.
func (pa *PersonAddress) QueryPersons() *PersonQuery {
	return NewPersonAddressClient(pa.config).QueryPersons(pa)
}

// QueryAddresses queries the "addresses" edge of the PersonAddress entity.
func (pa *PersonAddress) QueryAddresses() *AddressQuery {
	return NewPersonAddressClient(pa.config).QueryAddresses(pa)
}

// Update returns a builder for updating this PersonAddress.
// Note that you need to call PersonAddress.Unwrap() before calling this method if this PersonAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *PersonAddress) Update() *PersonAddressUpdateOne {
	return NewPersonAddressClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the PersonAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *PersonAddress) Unwrap() *PersonAddress {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: PersonAddress is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *PersonAddress) String() string {
	var builder strings.Builder
	builder.WriteString("PersonAddress(")
	if v := pa.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pa.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("person_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.PersonID))
	builder.WriteString(", ")
	builder.WriteString("address_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.AddressID))
	builder.WriteByte(')')
	return builder.String()
}

// PersonAddresses is a parsable slice of PersonAddress.
type PersonAddresses []*PersonAddress
