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
	"github.com/thaiha1607/foursquare_server/ent/orderitem"
	"github.com/thaiha1607/foursquare_server/ent/shipment"
	"github.com/thaiha1607/foursquare_server/ent/shipmentitem"
)

// ShipmentItem is the model entity for the ShipmentItem schema.
type ShipmentItem struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ShipmentID holds the value of the "shipment_id" field.
	ShipmentID string `json:"shipment_id,omitempty"`
	// OrderItemID holds the value of the "order_item_id" field.
	OrderItemID uuid.UUID `json:"order_item_id,omitempty"`
	// Qty holds the value of the "qty" field.
	Qty decimal.Decimal `json:"qty,omitempty"`
	// Total holds the value of the "total" field.
	Total decimal.Decimal `json:"total,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShipmentItemQuery when eager-loading is set.
	Edges        ShipmentItemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ShipmentItemEdges holds the relations/edges for other nodes in the graph.
type ShipmentItemEdges struct {
	// Shipment holds the value of the shipment edge.
	Shipment *Shipment `json:"shipment,omitempty"`
	// OrderItem holds the value of the order_item edge.
	OrderItem *OrderItem `json:"order_item,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ShipmentOrErr returns the Shipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShipmentItemEdges) ShipmentOrErr() (*Shipment, error) {
	if e.Shipment != nil {
		return e.Shipment, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: shipment.Label}
	}
	return nil, &NotLoadedError{edge: "shipment"}
}

// OrderItemOrErr returns the OrderItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShipmentItemEdges) OrderItemOrErr() (*OrderItem, error) {
	if e.OrderItem != nil {
		return e.OrderItem, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: orderitem.Label}
	}
	return nil, &NotLoadedError{edge: "order_item"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShipmentItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case shipmentitem.FieldQty, shipmentitem.FieldTotal:
			values[i] = new(decimal.Decimal)
		case shipmentitem.FieldShipmentID:
			values[i] = new(sql.NullString)
		case shipmentitem.FieldCreatedAt, shipmentitem.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case shipmentitem.FieldID, shipmentitem.FieldOrderItemID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShipmentItem fields.
func (si *ShipmentItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shipmentitem.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				si.ID = *value
			}
		case shipmentitem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				si.CreatedAt = value.Time
			}
		case shipmentitem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				si.UpdatedAt = value.Time
			}
		case shipmentitem.FieldShipmentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shipment_id", values[i])
			} else if value.Valid {
				si.ShipmentID = value.String
			}
		case shipmentitem.FieldOrderItemID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_item_id", values[i])
			} else if value != nil {
				si.OrderItemID = *value
			}
		case shipmentitem.FieldQty:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field qty", values[i])
			} else if value != nil {
				si.Qty = *value
			}
		case shipmentitem.FieldTotal:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total", values[i])
			} else if value != nil {
				si.Total = *value
			}
		default:
			si.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ShipmentItem.
// This includes values selected through modifiers, order, etc.
func (si *ShipmentItem) Value(name string) (ent.Value, error) {
	return si.selectValues.Get(name)
}

// QueryShipment queries the "shipment" edge of the ShipmentItem entity.
func (si *ShipmentItem) QueryShipment() *ShipmentQuery {
	return NewShipmentItemClient(si.config).QueryShipment(si)
}

// QueryOrderItem queries the "order_item" edge of the ShipmentItem entity.
func (si *ShipmentItem) QueryOrderItem() *OrderItemQuery {
	return NewShipmentItemClient(si.config).QueryOrderItem(si)
}

// Update returns a builder for updating this ShipmentItem.
// Note that you need to call ShipmentItem.Unwrap() before calling this method if this ShipmentItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (si *ShipmentItem) Update() *ShipmentItemUpdateOne {
	return NewShipmentItemClient(si.config).UpdateOne(si)
}

// Unwrap unwraps the ShipmentItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (si *ShipmentItem) Unwrap() *ShipmentItem {
	_tx, ok := si.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShipmentItem is not a transactional entity")
	}
	si.config.driver = _tx.drv
	return si
}

// String implements the fmt.Stringer.
func (si *ShipmentItem) String() string {
	var builder strings.Builder
	builder.WriteString("ShipmentItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", si.ID))
	builder.WriteString("created_at=")
	builder.WriteString(si.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(si.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("shipment_id=")
	builder.WriteString(si.ShipmentID)
	builder.WriteString(", ")
	builder.WriteString("order_item_id=")
	builder.WriteString(fmt.Sprintf("%v", si.OrderItemID))
	builder.WriteString(", ")
	builder.WriteString("qty=")
	builder.WriteString(fmt.Sprintf("%v", si.Qty))
	builder.WriteString(", ")
	builder.WriteString("total=")
	builder.WriteString(fmt.Sprintf("%v", si.Total))
	builder.WriteByte(')')
	return builder.String()
}

// ShipmentItems is a parsable slice of ShipmentItem.
type ShipmentItems []*ShipmentItem
