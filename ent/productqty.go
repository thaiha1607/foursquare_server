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
	"github.com/thaiha1607/foursquare_server/ent/productcolor"
	"github.com/thaiha1607/foursquare_server/ent/productinfo"
	"github.com/thaiha1607/foursquare_server/ent/productqty"
	"github.com/thaiha1607/foursquare_server/ent/workunitinfo"
)

// ProductQty is the model entity for the ProductQty schema.
type ProductQty struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// WorkUnitID holds the value of the "work_unit_id" field.
	WorkUnitID uuid.UUID `json:"work_unit_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID string `json:"product_id,omitempty"`
	// ProductColorID holds the value of the "product_color_id" field.
	ProductColorID string `json:"product_color_id,omitempty"`
	// PricePerUnit holds the value of the "price_per_unit" field.
	PricePerUnit *decimal.Decimal `json:"price_per_unit,omitempty"`
	// Qty holds the value of the "qty" field.
	Qty *decimal.Decimal `json:"qty,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQtyQuery when eager-loading is set.
	Edges        ProductQtyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductQtyEdges holds the relations/edges for other nodes in the graph.
type ProductQtyEdges struct {
	// WorkUnit holds the value of the work_unit edge.
	WorkUnit *WorkUnitInfo `json:"work_unit,omitempty"`
	// Product holds the value of the product edge.
	Product *ProductInfo `json:"product,omitempty"`
	// ProductColor holds the value of the product_color edge.
	ProductColor *ProductColor `json:"product_color,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// WorkUnitOrErr returns the WorkUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductQtyEdges) WorkUnitOrErr() (*WorkUnitInfo, error) {
	if e.WorkUnit != nil {
		return e.WorkUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: workunitinfo.Label}
	}
	return nil, &NotLoadedError{edge: "work_unit"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductQtyEdges) ProductOrErr() (*ProductInfo, error) {
	if e.Product != nil {
		return e.Product, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: productinfo.Label}
	}
	return nil, &NotLoadedError{edge: "product"}
}

// ProductColorOrErr returns the ProductColor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductQtyEdges) ProductColorOrErr() (*ProductColor, error) {
	if e.ProductColor != nil {
		return e.ProductColor, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: productcolor.Label}
	}
	return nil, &NotLoadedError{edge: "product_color"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductQty) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case productqty.FieldPricePerUnit, productqty.FieldQty:
			values[i] = &sql.NullScanner{S: new(decimal.Decimal)}
		case productqty.FieldProductID, productqty.FieldProductColorID:
			values[i] = new(sql.NullString)
		case productqty.FieldCreatedAt, productqty.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case productqty.FieldID, productqty.FieldWorkUnitID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductQty fields.
func (pq *ProductQty) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productqty.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pq.ID = *value
			}
		case productqty.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pq.CreatedAt = new(time.Time)
				*pq.CreatedAt = value.Time
			}
		case productqty.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pq.UpdatedAt = new(time.Time)
				*pq.UpdatedAt = value.Time
			}
		case productqty.FieldWorkUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field work_unit_id", values[i])
			} else if value != nil {
				pq.WorkUnitID = *value
			}
		case productqty.FieldProductID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				pq.ProductID = value.String
			}
		case productqty.FieldProductColorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_color_id", values[i])
			} else if value.Valid {
				pq.ProductColorID = value.String
			}
		case productqty.FieldPricePerUnit:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field price_per_unit", values[i])
			} else if value.Valid {
				pq.PricePerUnit = new(decimal.Decimal)
				*pq.PricePerUnit = *value.S.(*decimal.Decimal)
			}
		case productqty.FieldQty:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field qty", values[i])
			} else if value.Valid {
				pq.Qty = new(decimal.Decimal)
				*pq.Qty = *value.S.(*decimal.Decimal)
			}
		default:
			pq.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductQty.
// This includes values selected through modifiers, order, etc.
func (pq *ProductQty) Value(name string) (ent.Value, error) {
	return pq.selectValues.Get(name)
}

// QueryWorkUnit queries the "work_unit" edge of the ProductQty entity.
func (pq *ProductQty) QueryWorkUnit() *WorkUnitInfoQuery {
	return NewProductQtyClient(pq.config).QueryWorkUnit(pq)
}

// QueryProduct queries the "product" edge of the ProductQty entity.
func (pq *ProductQty) QueryProduct() *ProductInfoQuery {
	return NewProductQtyClient(pq.config).QueryProduct(pq)
}

// QueryProductColor queries the "product_color" edge of the ProductQty entity.
func (pq *ProductQty) QueryProductColor() *ProductColorQuery {
	return NewProductQtyClient(pq.config).QueryProductColor(pq)
}

// Update returns a builder for updating this ProductQty.
// Note that you need to call ProductQty.Unwrap() before calling this method if this ProductQty
// was returned from a transaction, and the transaction was committed or rolled back.
func (pq *ProductQty) Update() *ProductQtyUpdateOne {
	return NewProductQtyClient(pq.config).UpdateOne(pq)
}

// Unwrap unwraps the ProductQty entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pq *ProductQty) Unwrap() *ProductQty {
	_tx, ok := pq.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductQty is not a transactional entity")
	}
	pq.config.driver = _tx.drv
	return pq
}

// String implements the fmt.Stringer.
func (pq *ProductQty) String() string {
	var builder strings.Builder
	builder.WriteString("ProductQty(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pq.ID))
	if v := pq.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pq.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("work_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", pq.WorkUnitID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(pq.ProductID)
	builder.WriteString(", ")
	builder.WriteString("product_color_id=")
	builder.WriteString(pq.ProductColorID)
	builder.WriteString(", ")
	if v := pq.PricePerUnit; v != nil {
		builder.WriteString("price_per_unit=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := pq.Qty; v != nil {
		builder.WriteString("qty=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// ProductQties is a parsable slice of ProductQty.
type ProductQties []*ProductQty
