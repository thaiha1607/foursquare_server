// Code generated by ent, DO NOT EDIT.

package invoicetype

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldEQ(FieldUpdatedAt, v))
}

// InvoiceType applies equality check predicate on the "invoice_type" field. It's identical to InvoiceTypeEQ.
func InvoiceType(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldEQ(FieldInvoiceType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldLTE(FieldUpdatedAt, v))
}

// InvoiceTypeEQ applies the EQ predicate on the "invoice_type" field.
func InvoiceTypeEQ(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldEQ(FieldInvoiceType, v))
}

// InvoiceTypeNEQ applies the NEQ predicate on the "invoice_type" field.
func InvoiceTypeNEQ(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldNEQ(FieldInvoiceType, v))
}

// InvoiceTypeIn applies the In predicate on the "invoice_type" field.
func InvoiceTypeIn(vs ...string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldIn(FieldInvoiceType, vs...))
}

// InvoiceTypeNotIn applies the NotIn predicate on the "invoice_type" field.
func InvoiceTypeNotIn(vs ...string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldNotIn(FieldInvoiceType, vs...))
}

// InvoiceTypeGT applies the GT predicate on the "invoice_type" field.
func InvoiceTypeGT(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldGT(FieldInvoiceType, v))
}

// InvoiceTypeGTE applies the GTE predicate on the "invoice_type" field.
func InvoiceTypeGTE(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldGTE(FieldInvoiceType, v))
}

// InvoiceTypeLT applies the LT predicate on the "invoice_type" field.
func InvoiceTypeLT(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldLT(FieldInvoiceType, v))
}

// InvoiceTypeLTE applies the LTE predicate on the "invoice_type" field.
func InvoiceTypeLTE(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldLTE(FieldInvoiceType, v))
}

// InvoiceTypeContains applies the Contains predicate on the "invoice_type" field.
func InvoiceTypeContains(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldContains(FieldInvoiceType, v))
}

// InvoiceTypeHasPrefix applies the HasPrefix predicate on the "invoice_type" field.
func InvoiceTypeHasPrefix(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldHasPrefix(FieldInvoiceType, v))
}

// InvoiceTypeHasSuffix applies the HasSuffix predicate on the "invoice_type" field.
func InvoiceTypeHasSuffix(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldHasSuffix(FieldInvoiceType, v))
}

// InvoiceTypeEqualFold applies the EqualFold predicate on the "invoice_type" field.
func InvoiceTypeEqualFold(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldEqualFold(FieldInvoiceType, v))
}

// InvoiceTypeContainsFold applies the ContainsFold predicate on the "invoice_type" field.
func InvoiceTypeContainsFold(v string) predicate.InvoiceType {
	return predicate.InvoiceType(sql.FieldContainsFold(FieldInvoiceType, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InvoiceType) predicate.InvoiceType {
	return predicate.InvoiceType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InvoiceType) predicate.InvoiceType {
	return predicate.InvoiceType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.InvoiceType) predicate.InvoiceType {
	return predicate.InvoiceType(sql.NotPredicates(p))
}
