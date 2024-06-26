// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/thaiha1607/foursquare_server/ent/productinfo"
	"github.com/thaiha1607/foursquare_server/ent/producttag"
	"github.com/thaiha1607/foursquare_server/ent/tag"
)

// ProductTag is the model entity for the ProductTag schema.
type ProductTag struct {
	config `json:"-"`
	// ProductID holds the value of the "product_id" field.
	ProductID string `json:"product_id,omitempty"`
	// TagID holds the value of the "tag_id" field.
	TagID string `json:"tag_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductTagQuery when eager-loading is set.
	Edges        ProductTagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductTagEdges holds the relations/edges for other nodes in the graph.
type ProductTagEdges struct {
	// Products holds the value of the products edge.
	Products *ProductInfo `json:"products,omitempty"`
	// Tags holds the value of the tags edge.
	Tags *Tag `json:"tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductTagEdges) ProductsOrErr() (*ProductInfo, error) {
	if e.Products != nil {
		return e.Products, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: productinfo.Label}
	}
	return nil, &NotLoadedError{edge: "products"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductTagEdges) TagsOrErr() (*Tag, error) {
	if e.Tags != nil {
		return e.Tags, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: tag.Label}
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case producttag.FieldProductID, producttag.FieldTagID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductTag fields.
func (pt *ProductTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case producttag.FieldProductID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				pt.ProductID = value.String
			}
		case producttag.FieldTagID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag_id", values[i])
			} else if value.Valid {
				pt.TagID = value.String
			}
		default:
			pt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductTag.
// This includes values selected through modifiers, order, etc.
func (pt *ProductTag) Value(name string) (ent.Value, error) {
	return pt.selectValues.Get(name)
}

// QueryProducts queries the "products" edge of the ProductTag entity.
func (pt *ProductTag) QueryProducts() *ProductInfoQuery {
	return NewProductTagClient(pt.config).QueryProducts(pt)
}

// QueryTags queries the "tags" edge of the ProductTag entity.
func (pt *ProductTag) QueryTags() *TagQuery {
	return NewProductTagClient(pt.config).QueryTags(pt)
}

// Update returns a builder for updating this ProductTag.
// Note that you need to call ProductTag.Unwrap() before calling this method if this ProductTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *ProductTag) Update() *ProductTagUpdateOne {
	return NewProductTagClient(pt.config).UpdateOne(pt)
}

// Unwrap unwraps the ProductTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *ProductTag) Unwrap() *ProductTag {
	_tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductTag is not a transactional entity")
	}
	pt.config.driver = _tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *ProductTag) String() string {
	var builder strings.Builder
	builder.WriteString("ProductTag(")
	builder.WriteString("product_id=")
	builder.WriteString(pt.ProductID)
	builder.WriteString(", ")
	builder.WriteString("tag_id=")
	builder.WriteString(pt.TagID)
	builder.WriteByte(')')
	return builder.String()
}

// ProductTags is a parsable slice of ProductTag.
type ProductTags []*ProductTag
