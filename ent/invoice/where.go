// Code generated by ent, DO NOT EDIT.

package invoice

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldOrderID, v))
}

// Total applies equality check predicate on the "total" field. It's identical to TotalEQ.
func Total(v decimal.Decimal) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldTotal, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldComment, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldNote, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Invoice {
	return predicate.Invoice(sql.FieldLTE(FieldUpdatedAt, v))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldOrderID, vs...))
}

// TotalEQ applies the EQ predicate on the "total" field.
func TotalEQ(v decimal.Decimal) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldTotal, v))
}

// TotalNEQ applies the NEQ predicate on the "total" field.
func TotalNEQ(v decimal.Decimal) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldTotal, v))
}

// TotalIn applies the In predicate on the "total" field.
func TotalIn(vs ...decimal.Decimal) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldTotal, vs...))
}

// TotalNotIn applies the NotIn predicate on the "total" field.
func TotalNotIn(vs ...decimal.Decimal) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldTotal, vs...))
}

// TotalGT applies the GT predicate on the "total" field.
func TotalGT(v decimal.Decimal) predicate.Invoice {
	return predicate.Invoice(sql.FieldGT(FieldTotal, v))
}

// TotalGTE applies the GTE predicate on the "total" field.
func TotalGTE(v decimal.Decimal) predicate.Invoice {
	return predicate.Invoice(sql.FieldGTE(FieldTotal, v))
}

// TotalLT applies the LT predicate on the "total" field.
func TotalLT(v decimal.Decimal) predicate.Invoice {
	return predicate.Invoice(sql.FieldLT(FieldTotal, v))
}

// TotalLTE applies the LTE predicate on the "total" field.
func TotalLTE(v decimal.Decimal) predicate.Invoice {
	return predicate.Invoice(sql.FieldLTE(FieldTotal, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldHasSuffix(FieldComment, v))
}

// CommentIsNil applies the IsNil predicate on the "comment" field.
func CommentIsNil() predicate.Invoice {
	return predicate.Invoice(sql.FieldIsNull(FieldComment))
}

// CommentNotNil applies the NotNil predicate on the "comment" field.
func CommentNotNil() predicate.Invoice {
	return predicate.Invoice(sql.FieldNotNull(FieldComment))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldContainsFold(FieldComment, v))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldHasSuffix(FieldNote, v))
}

// NoteIsNil applies the IsNil predicate on the "note" field.
func NoteIsNil() predicate.Invoice {
	return predicate.Invoice(sql.FieldIsNull(FieldNote))
}

// NoteNotNil applies the NotNil predicate on the "note" field.
func NoteNotNil() predicate.Invoice {
	return predicate.Invoice(sql.FieldNotNull(FieldNote))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldContainsFold(FieldNote, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldType, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldStatus, vs...))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.Invoice {
	return predicate.Invoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.Invoice {
	return predicate.Invoice(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInvoiceType applies the HasEdge predicate on the "invoice_type" edge.
func HasInvoiceType() predicate.Invoice {
	return predicate.Invoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, InvoiceTypeTable, InvoiceTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvoiceTypeWith applies the HasEdge predicate on the "invoice_type" edge with a given conditions (other predicates).
func HasInvoiceTypeWith(preds ...predicate.InvoiceType) predicate.Invoice {
	return predicate.Invoice(func(s *sql.Selector) {
		step := newInvoiceTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Invoice) predicate.Invoice {
	return predicate.Invoice(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Invoice) predicate.Invoice {
	return predicate.Invoice(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Invoice) predicate.Invoice {
	return predicate.Invoice(sql.NotPredicates(p))
}
