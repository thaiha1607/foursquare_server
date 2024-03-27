// Code generated by ent, DO NOT EDIT.

package productimage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEQ(FieldUpdatedAt, v))
}

// ProductID applies equality check predicate on the "product_id" field. It's identical to ProductIDEQ.
func ProductID(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEQ(FieldProductID, v))
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEQ(FieldImageURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldLTE(FieldUpdatedAt, v))
}

// ProductIDEQ applies the EQ predicate on the "product_id" field.
func ProductIDEQ(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEQ(FieldProductID, v))
}

// ProductIDNEQ applies the NEQ predicate on the "product_id" field.
func ProductIDNEQ(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldNEQ(FieldProductID, v))
}

// ProductIDIn applies the In predicate on the "product_id" field.
func ProductIDIn(vs ...string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldIn(FieldProductID, vs...))
}

// ProductIDNotIn applies the NotIn predicate on the "product_id" field.
func ProductIDNotIn(vs ...string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldNotIn(FieldProductID, vs...))
}

// ProductIDGT applies the GT predicate on the "product_id" field.
func ProductIDGT(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldGT(FieldProductID, v))
}

// ProductIDGTE applies the GTE predicate on the "product_id" field.
func ProductIDGTE(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldGTE(FieldProductID, v))
}

// ProductIDLT applies the LT predicate on the "product_id" field.
func ProductIDLT(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldLT(FieldProductID, v))
}

// ProductIDLTE applies the LTE predicate on the "product_id" field.
func ProductIDLTE(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldLTE(FieldProductID, v))
}

// ProductIDContains applies the Contains predicate on the "product_id" field.
func ProductIDContains(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldContains(FieldProductID, v))
}

// ProductIDHasPrefix applies the HasPrefix predicate on the "product_id" field.
func ProductIDHasPrefix(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldHasPrefix(FieldProductID, v))
}

// ProductIDHasSuffix applies the HasSuffix predicate on the "product_id" field.
func ProductIDHasSuffix(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldHasSuffix(FieldProductID, v))
}

// ProductIDEqualFold applies the EqualFold predicate on the "product_id" field.
func ProductIDEqualFold(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEqualFold(FieldProductID, v))
}

// ProductIDContainsFold applies the ContainsFold predicate on the "product_id" field.
func ProductIDContainsFold(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldContainsFold(FieldProductID, v))
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEQ(FieldImageURL, v))
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldNEQ(FieldImageURL, v))
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldIn(FieldImageURL, vs...))
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldNotIn(FieldImageURL, vs...))
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldGT(FieldImageURL, v))
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldGTE(FieldImageURL, v))
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldLT(FieldImageURL, v))
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldLTE(FieldImageURL, v))
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldContains(FieldImageURL, v))
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldHasPrefix(FieldImageURL, v))
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldHasSuffix(FieldImageURL, v))
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldEqualFold(FieldImageURL, v))
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.ProductImage {
	return predicate.ProductImage(sql.FieldContainsFold(FieldImageURL, v))
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.ProductImage {
	return predicate.ProductImage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.ProductInfo) predicate.ProductImage {
	return predicate.ProductImage(func(s *sql.Selector) {
		step := newProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProductImage) predicate.ProductImage {
	return predicate.ProductImage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProductImage) predicate.ProductImage {
	return predicate.ProductImage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProductImage) predicate.ProductImage {
	return predicate.ProductImage(sql.NotPredicates(p))
}
