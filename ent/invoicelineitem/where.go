// Code generated by ent, DO NOT EDIT.

package invoicelineitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldLTE(FieldID, id))
}

// InvoiceID applies equality check predicate on the "invoice_id" field. It's identical to InvoiceIDEQ.
func InvoiceID(v uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldInvoiceID, v))
}

// OrderLineItemID applies equality check predicate on the "order_line_item_id" field. It's identical to OrderLineItemIDEQ.
func OrderLineItemID(v uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldOrderLineItemID, v))
}

// Qty applies equality check predicate on the "qty" field. It's identical to QtyEQ.
func Qty(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldQty, v))
}

// Total applies equality check predicate on the "total" field. It's identical to TotalEQ.
func Total(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldTotal, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldCreatedAt, v))
}

// InvoiceIDEQ applies the EQ predicate on the "invoice_id" field.
func InvoiceIDEQ(v uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldInvoiceID, v))
}

// InvoiceIDNEQ applies the NEQ predicate on the "invoice_id" field.
func InvoiceIDNEQ(v uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNEQ(FieldInvoiceID, v))
}

// InvoiceIDIn applies the In predicate on the "invoice_id" field.
func InvoiceIDIn(vs ...uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldIn(FieldInvoiceID, vs...))
}

// InvoiceIDNotIn applies the NotIn predicate on the "invoice_id" field.
func InvoiceIDNotIn(vs ...uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNotIn(FieldInvoiceID, vs...))
}

// OrderLineItemIDEQ applies the EQ predicate on the "order_line_item_id" field.
func OrderLineItemIDEQ(v uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldOrderLineItemID, v))
}

// OrderLineItemIDNEQ applies the NEQ predicate on the "order_line_item_id" field.
func OrderLineItemIDNEQ(v uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNEQ(FieldOrderLineItemID, v))
}

// OrderLineItemIDIn applies the In predicate on the "order_line_item_id" field.
func OrderLineItemIDIn(vs ...uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldIn(FieldOrderLineItemID, vs...))
}

// OrderLineItemIDNotIn applies the NotIn predicate on the "order_line_item_id" field.
func OrderLineItemIDNotIn(vs ...uuid.UUID) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNotIn(FieldOrderLineItemID, vs...))
}

// QtyEQ applies the EQ predicate on the "qty" field.
func QtyEQ(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldQty, v))
}

// QtyNEQ applies the NEQ predicate on the "qty" field.
func QtyNEQ(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNEQ(FieldQty, v))
}

// QtyIn applies the In predicate on the "qty" field.
func QtyIn(vs ...decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldIn(FieldQty, vs...))
}

// QtyNotIn applies the NotIn predicate on the "qty" field.
func QtyNotIn(vs ...decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNotIn(FieldQty, vs...))
}

// QtyGT applies the GT predicate on the "qty" field.
func QtyGT(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldGT(FieldQty, v))
}

// QtyGTE applies the GTE predicate on the "qty" field.
func QtyGTE(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldGTE(FieldQty, v))
}

// QtyLT applies the LT predicate on the "qty" field.
func QtyLT(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldLT(FieldQty, v))
}

// QtyLTE applies the LTE predicate on the "qty" field.
func QtyLTE(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldLTE(FieldQty, v))
}

// TotalEQ applies the EQ predicate on the "total" field.
func TotalEQ(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldTotal, v))
}

// TotalNEQ applies the NEQ predicate on the "total" field.
func TotalNEQ(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNEQ(FieldTotal, v))
}

// TotalIn applies the In predicate on the "total" field.
func TotalIn(vs ...decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldIn(FieldTotal, vs...))
}

// TotalNotIn applies the NotIn predicate on the "total" field.
func TotalNotIn(vs ...decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNotIn(FieldTotal, vs...))
}

// TotalGT applies the GT predicate on the "total" field.
func TotalGT(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldGT(FieldTotal, v))
}

// TotalGTE applies the GTE predicate on the "total" field.
func TotalGTE(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldGTE(FieldTotal, v))
}

// TotalLT applies the LT predicate on the "total" field.
func TotalLT(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldLT(FieldTotal, v))
}

// TotalLTE applies the LTE predicate on the "total" field.
func TotalLTE(v decimal.Decimal) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldLTE(FieldTotal, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.FieldLTE(FieldCreatedAt, v))
}

// HasInvoice applies the HasEdge predicate on the "invoice" edge.
func HasInvoice() predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, InvoiceTable, InvoiceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvoiceWith applies the HasEdge predicate on the "invoice" edge with a given conditions (other predicates).
func HasInvoiceWith(preds ...predicate.Invoice) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(func(s *sql.Selector) {
		step := newInvoiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderLineItem applies the HasEdge predicate on the "order_line_item" edge.
func HasOrderLineItem() predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrderLineItemTable, OrderLineItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderLineItemWith applies the HasEdge predicate on the "order_line_item" edge with a given conditions (other predicates).
func HasOrderLineItemWith(preds ...predicate.OrderLineItem) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(func(s *sql.Selector) {
		step := newOrderLineItemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InvoiceLineItem) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InvoiceLineItem) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.InvoiceLineItem) predicate.InvoiceLineItem {
	return predicate.InvoiceLineItem(sql.NotPredicates(p))
}
