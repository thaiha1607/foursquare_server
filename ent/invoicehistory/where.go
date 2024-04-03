// Code generated by ent, DO NOT EDIT.

package invoicehistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// InvoiceID applies equality check predicate on the "invoice_id" field. It's identical to InvoiceIDEQ.
func InvoiceID(v uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldInvoiceID, v))
}

// PersonID applies equality check predicate on the "person_id" field. It's identical to PersonIDEQ.
func PersonID(v uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldPersonID, v))
}

// OldStatusCode applies equality check predicate on the "old_status_code" field. It's identical to OldStatusCodeEQ.
func OldStatusCode(v int) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldOldStatusCode, v))
}

// NewStatusCode applies equality check predicate on the "new_status_code" field. It's identical to NewStatusCodeEQ.
func NewStatusCode(v int) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldNewStatusCode, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldLTE(FieldUpdatedAt, v))
}

// InvoiceIDEQ applies the EQ predicate on the "invoice_id" field.
func InvoiceIDEQ(v uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldInvoiceID, v))
}

// InvoiceIDNEQ applies the NEQ predicate on the "invoice_id" field.
func InvoiceIDNEQ(v uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNEQ(FieldInvoiceID, v))
}

// InvoiceIDIn applies the In predicate on the "invoice_id" field.
func InvoiceIDIn(vs ...uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIn(FieldInvoiceID, vs...))
}

// InvoiceIDNotIn applies the NotIn predicate on the "invoice_id" field.
func InvoiceIDNotIn(vs ...uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotIn(FieldInvoiceID, vs...))
}

// PersonIDEQ applies the EQ predicate on the "person_id" field.
func PersonIDEQ(v uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldPersonID, v))
}

// PersonIDNEQ applies the NEQ predicate on the "person_id" field.
func PersonIDNEQ(v uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNEQ(FieldPersonID, v))
}

// PersonIDIn applies the In predicate on the "person_id" field.
func PersonIDIn(vs ...uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIn(FieldPersonID, vs...))
}

// PersonIDNotIn applies the NotIn predicate on the "person_id" field.
func PersonIDNotIn(vs ...uuid.UUID) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotIn(FieldPersonID, vs...))
}

// OldStatusCodeEQ applies the EQ predicate on the "old_status_code" field.
func OldStatusCodeEQ(v int) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldOldStatusCode, v))
}

// OldStatusCodeNEQ applies the NEQ predicate on the "old_status_code" field.
func OldStatusCodeNEQ(v int) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNEQ(FieldOldStatusCode, v))
}

// OldStatusCodeIn applies the In predicate on the "old_status_code" field.
func OldStatusCodeIn(vs ...int) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIn(FieldOldStatusCode, vs...))
}

// OldStatusCodeNotIn applies the NotIn predicate on the "old_status_code" field.
func OldStatusCodeNotIn(vs ...int) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotIn(FieldOldStatusCode, vs...))
}

// OldStatusCodeIsNil applies the IsNil predicate on the "old_status_code" field.
func OldStatusCodeIsNil() predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIsNull(FieldOldStatusCode))
}

// OldStatusCodeNotNil applies the NotNil predicate on the "old_status_code" field.
func OldStatusCodeNotNil() predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotNull(FieldOldStatusCode))
}

// NewStatusCodeEQ applies the EQ predicate on the "new_status_code" field.
func NewStatusCodeEQ(v int) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldNewStatusCode, v))
}

// NewStatusCodeNEQ applies the NEQ predicate on the "new_status_code" field.
func NewStatusCodeNEQ(v int) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNEQ(FieldNewStatusCode, v))
}

// NewStatusCodeIn applies the In predicate on the "new_status_code" field.
func NewStatusCodeIn(vs ...int) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIn(FieldNewStatusCode, vs...))
}

// NewStatusCodeNotIn applies the NotIn predicate on the "new_status_code" field.
func NewStatusCodeNotIn(vs ...int) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotIn(FieldNewStatusCode, vs...))
}

// NewStatusCodeIsNil applies the IsNil predicate on the "new_status_code" field.
func NewStatusCodeIsNil() predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIsNull(FieldNewStatusCode))
}

// NewStatusCodeNotNil applies the NotNil predicate on the "new_status_code" field.
func NewStatusCodeNotNil() predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotNull(FieldNewStatusCode))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.FieldContainsFold(FieldDescription, v))
}

// HasInvoice applies the HasEdge predicate on the "invoice" edge.
func HasInvoice() predicate.InvoiceHistory {
	return predicate.InvoiceHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, InvoiceTable, InvoiceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvoiceWith applies the HasEdge predicate on the "invoice" edge with a given conditions (other predicates).
func HasInvoiceWith(preds ...predicate.Invoice) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(func(s *sql.Selector) {
		step := newInvoiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPerson applies the HasEdge predicate on the "person" edge.
func HasPerson() predicate.InvoiceHistory {
	return predicate.InvoiceHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PersonTable, PersonColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonWith applies the HasEdge predicate on the "person" edge with a given conditions (other predicates).
func HasPersonWith(preds ...predicate.Person) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(func(s *sql.Selector) {
		step := newPersonStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOldStatus applies the HasEdge predicate on the "old_status" edge.
func HasOldStatus() predicate.InvoiceHistory {
	return predicate.InvoiceHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OldStatusTable, OldStatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOldStatusWith applies the HasEdge predicate on the "old_status" edge with a given conditions (other predicates).
func HasOldStatusWith(preds ...predicate.InvoiceStatusCode) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(func(s *sql.Selector) {
		step := newOldStatusStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNewStatus applies the HasEdge predicate on the "new_status" edge.
func HasNewStatus() predicate.InvoiceHistory {
	return predicate.InvoiceHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NewStatusTable, NewStatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNewStatusWith applies the HasEdge predicate on the "new_status" edge with a given conditions (other predicates).
func HasNewStatusWith(preds ...predicate.InvoiceStatusCode) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(func(s *sql.Selector) {
		step := newNewStatusStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InvoiceHistory) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InvoiceHistory) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.InvoiceHistory) predicate.InvoiceHistory {
	return predicate.InvoiceHistory(sql.NotPredicates(p))
}
