// Code generated by ent, DO NOT EDIT.

package transactiontype

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldEQ(FieldUpdatedAt, v))
}

// TransactionType applies equality check predicate on the "transaction_type" field. It's identical to TransactionTypeEQ.
func TransactionType(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldEQ(FieldTransactionType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldLTE(FieldUpdatedAt, v))
}

// TransactionTypeEQ applies the EQ predicate on the "transaction_type" field.
func TransactionTypeEQ(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldEQ(FieldTransactionType, v))
}

// TransactionTypeNEQ applies the NEQ predicate on the "transaction_type" field.
func TransactionTypeNEQ(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldNEQ(FieldTransactionType, v))
}

// TransactionTypeIn applies the In predicate on the "transaction_type" field.
func TransactionTypeIn(vs ...string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldIn(FieldTransactionType, vs...))
}

// TransactionTypeNotIn applies the NotIn predicate on the "transaction_type" field.
func TransactionTypeNotIn(vs ...string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldNotIn(FieldTransactionType, vs...))
}

// TransactionTypeGT applies the GT predicate on the "transaction_type" field.
func TransactionTypeGT(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldGT(FieldTransactionType, v))
}

// TransactionTypeGTE applies the GTE predicate on the "transaction_type" field.
func TransactionTypeGTE(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldGTE(FieldTransactionType, v))
}

// TransactionTypeLT applies the LT predicate on the "transaction_type" field.
func TransactionTypeLT(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldLT(FieldTransactionType, v))
}

// TransactionTypeLTE applies the LTE predicate on the "transaction_type" field.
func TransactionTypeLTE(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldLTE(FieldTransactionType, v))
}

// TransactionTypeContains applies the Contains predicate on the "transaction_type" field.
func TransactionTypeContains(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldContains(FieldTransactionType, v))
}

// TransactionTypeHasPrefix applies the HasPrefix predicate on the "transaction_type" field.
func TransactionTypeHasPrefix(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldHasPrefix(FieldTransactionType, v))
}

// TransactionTypeHasSuffix applies the HasSuffix predicate on the "transaction_type" field.
func TransactionTypeHasSuffix(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldHasSuffix(FieldTransactionType, v))
}

// TransactionTypeEqualFold applies the EqualFold predicate on the "transaction_type" field.
func TransactionTypeEqualFold(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldEqualFold(FieldTransactionType, v))
}

// TransactionTypeContainsFold applies the ContainsFold predicate on the "transaction_type" field.
func TransactionTypeContainsFold(v string) predicate.TransactionType {
	return predicate.TransactionType(sql.FieldContainsFold(FieldTransactionType, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TransactionType) predicate.TransactionType {
	return predicate.TransactionType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TransactionType) predicate.TransactionType {
	return predicate.TransactionType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TransactionType) predicate.TransactionType {
	return predicate.TransactionType(sql.NotPredicates(p))
}
