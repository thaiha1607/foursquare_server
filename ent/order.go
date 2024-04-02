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
	"github.com/thaiha1607/foursquare_server/ent/order"
	"github.com/thaiha1607/foursquare_server/ent/orderstatuscode"
	"github.com/thaiha1607/foursquare_server/ent/person"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID uuid.UUID `json:"customer_id,omitempty"`
	// Note holds the value of the "note" field.
	Note *string `json:"note,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	// ParentOrderID holds the value of the "parent_order_id" field.
	ParentOrderID *uuid.UUID `json:"parent_order_id,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty"`
	// Type holds the value of the "type" field.
	Type order.Type `json:"type,omitempty"`
	// StatusCode holds the value of the "status_code" field.
	StatusCode int `json:"status_code,omitempty"`
	// StaffID holds the value of the "staff_id" field.
	StaffID uuid.UUID `json:"staff_id,omitempty"`
	// InternalNote holds the value of the "internal_note" field.
	InternalNote *string `json:"internal_note,omitempty"`
	// IsInternal holds the value of the "is_internal" field.
	IsInternal bool `json:"is_internal,omitempty"`
	// AddressID holds the value of the "address_id" field.
	AddressID uuid.UUID `json:"address_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges        OrderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// Customer holds the value of the customer edge.
	Customer *Person `json:"customer,omitempty"`
	// Creator holds the value of the creator edge.
	Creator *Person `json:"creator,omitempty"`
	// ParentOrder holds the value of the parent_order edge.
	ParentOrder *Order `json:"parent_order,omitempty"`
	// OrderStatus holds the value of the order_status edge.
	OrderStatus *OrderStatusCode `json:"order_status,omitempty"`
	// Staff holds the value of the staff edge.
	Staff *Person `json:"staff,omitempty"`
	// OrderAddress holds the value of the order_address edge.
	OrderAddress *Address `json:"order_address,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) CustomerOrErr() (*Person, error) {
	if e.Customer != nil {
		return e.Customer, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: person.Label}
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) CreatorOrErr() (*Person, error) {
	if e.Creator != nil {
		return e.Creator, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: person.Label}
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// ParentOrderOrErr returns the ParentOrder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) ParentOrderOrErr() (*Order, error) {
	if e.ParentOrder != nil {
		return e.ParentOrder, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: order.Label}
	}
	return nil, &NotLoadedError{edge: "parent_order"}
}

// OrderStatusOrErr returns the OrderStatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) OrderStatusOrErr() (*OrderStatusCode, error) {
	if e.OrderStatus != nil {
		return e.OrderStatus, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: orderstatuscode.Label}
	}
	return nil, &NotLoadedError{edge: "order_status"}
}

// StaffOrErr returns the Staff value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) StaffOrErr() (*Person, error) {
	if e.Staff != nil {
		return e.Staff, nil
	} else if e.loadedTypes[4] {
		return nil, &NotFoundError{label: person.Label}
	}
	return nil, &NotLoadedError{edge: "staff"}
}

// OrderAddressOrErr returns the OrderAddress value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) OrderAddressOrErr() (*Address, error) {
	if e.OrderAddress != nil {
		return e.OrderAddress, nil
	} else if e.loadedTypes[5] {
		return nil, &NotFoundError{label: address.Label}
	}
	return nil, &NotLoadedError{edge: "order_address"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldParentOrderID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case order.FieldIsInternal:
			values[i] = new(sql.NullBool)
		case order.FieldPriority, order.FieldStatusCode:
			values[i] = new(sql.NullInt64)
		case order.FieldNote, order.FieldType, order.FieldInternalNote:
			values[i] = new(sql.NullString)
		case order.FieldCreatedAt, order.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case order.FieldID, order.FieldCustomerID, order.FieldCreatedBy, order.FieldStaffID, order.FieldAddressID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				o.ID = *value
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case order.FieldCustomerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value != nil {
				o.CustomerID = *value
			}
		case order.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				o.Note = new(string)
				*o.Note = value.String
			}
		case order.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				o.CreatedBy = *value
			}
		case order.FieldParentOrderID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field parent_order_id", values[i])
			} else if value.Valid {
				o.ParentOrderID = new(uuid.UUID)
				*o.ParentOrderID = *value.S.(*uuid.UUID)
			}
		case order.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				o.Priority = int(value.Int64)
			}
		case order.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				o.Type = order.Type(value.String)
			}
		case order.FieldStatusCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status_code", values[i])
			} else if value.Valid {
				o.StatusCode = int(value.Int64)
			}
		case order.FieldStaffID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field staff_id", values[i])
			} else if value != nil {
				o.StaffID = *value
			}
		case order.FieldInternalNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field internal_note", values[i])
			} else if value.Valid {
				o.InternalNote = new(string)
				*o.InternalNote = value.String
			}
		case order.FieldIsInternal:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_internal", values[i])
			} else if value.Valid {
				o.IsInternal = value.Bool
			}
		case order.FieldAddressID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field address_id", values[i])
			} else if value != nil {
				o.AddressID = *value
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryCustomer queries the "customer" edge of the Order entity.
func (o *Order) QueryCustomer() *PersonQuery {
	return NewOrderClient(o.config).QueryCustomer(o)
}

// QueryCreator queries the "creator" edge of the Order entity.
func (o *Order) QueryCreator() *PersonQuery {
	return NewOrderClient(o.config).QueryCreator(o)
}

// QueryParentOrder queries the "parent_order" edge of the Order entity.
func (o *Order) QueryParentOrder() *OrderQuery {
	return NewOrderClient(o.config).QueryParentOrder(o)
}

// QueryOrderStatus queries the "order_status" edge of the Order entity.
func (o *Order) QueryOrderStatus() *OrderStatusCodeQuery {
	return NewOrderClient(o.config).QueryOrderStatus(o)
}

// QueryStaff queries the "staff" edge of the Order entity.
func (o *Order) QueryStaff() *PersonQuery {
	return NewOrderClient(o.config).QueryStaff(o)
}

// QueryOrderAddress queries the "order_address" edge of the Order entity.
func (o *Order) QueryOrderAddress() *AddressQuery {
	return NewOrderClient(o.config).QueryOrderAddress(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("customer_id=")
	builder.WriteString(fmt.Sprintf("%v", o.CustomerID))
	builder.WriteString(", ")
	if v := o.Note; v != nil {
		builder.WriteString("note=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", o.CreatedBy))
	builder.WriteString(", ")
	if v := o.ParentOrderID; v != nil {
		builder.WriteString("parent_order_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", o.Priority))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", o.Type))
	builder.WriteString(", ")
	builder.WriteString("status_code=")
	builder.WriteString(fmt.Sprintf("%v", o.StatusCode))
	builder.WriteString(", ")
	builder.WriteString("staff_id=")
	builder.WriteString(fmt.Sprintf("%v", o.StaffID))
	builder.WriteString(", ")
	if v := o.InternalNote; v != nil {
		builder.WriteString("internal_note=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("is_internal=")
	builder.WriteString(fmt.Sprintf("%v", o.IsInternal))
	builder.WriteString(", ")
	builder.WriteString("address_id=")
	builder.WriteString(fmt.Sprintf("%v", o.AddressID))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order
