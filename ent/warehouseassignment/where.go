// Code generated by ent, DO NOT EDIT.

package warehouseassignment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldOrderID, v))
}

// WorkUnitID applies equality check predicate on the "work_unit_id" field. It's identical to WorkUnitIDEQ.
func WorkUnitID(v uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldWorkUnitID, v))
}

// StaffID applies equality check predicate on the "staff_id" field. It's identical to StaffIDEQ.
func StaffID(v uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldStaffID, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldNote, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldLTE(FieldUpdatedAt, v))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNotIn(FieldOrderID, vs...))
}

// WorkUnitIDEQ applies the EQ predicate on the "work_unit_id" field.
func WorkUnitIDEQ(v uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldWorkUnitID, v))
}

// WorkUnitIDNEQ applies the NEQ predicate on the "work_unit_id" field.
func WorkUnitIDNEQ(v uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNEQ(FieldWorkUnitID, v))
}

// WorkUnitIDIn applies the In predicate on the "work_unit_id" field.
func WorkUnitIDIn(vs ...uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldIn(FieldWorkUnitID, vs...))
}

// WorkUnitIDNotIn applies the NotIn predicate on the "work_unit_id" field.
func WorkUnitIDNotIn(vs ...uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNotIn(FieldWorkUnitID, vs...))
}

// StaffIDEQ applies the EQ predicate on the "staff_id" field.
func StaffIDEQ(v uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldStaffID, v))
}

// StaffIDNEQ applies the NEQ predicate on the "staff_id" field.
func StaffIDNEQ(v uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNEQ(FieldStaffID, v))
}

// StaffIDIn applies the In predicate on the "staff_id" field.
func StaffIDIn(vs ...uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldIn(FieldStaffID, vs...))
}

// StaffIDNotIn applies the NotIn predicate on the "staff_id" field.
func StaffIDNotIn(vs ...uuid.UUID) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNotIn(FieldStaffID, vs...))
}

// StaffIDIsNil applies the IsNil predicate on the "staff_id" field.
func StaffIDIsNil() predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldIsNull(FieldStaffID))
}

// StaffIDNotNil applies the NotNil predicate on the "staff_id" field.
func StaffIDNotNil() predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNotNull(FieldStaffID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNotIn(FieldStatus, vs...))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldHasSuffix(FieldNote, v))
}

// NoteIsNil applies the IsNil predicate on the "note" field.
func NoteIsNil() predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldIsNull(FieldNote))
}

// NoteNotNil applies the NotNil predicate on the "note" field.
func NoteNotNil() predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldNotNull(FieldNote))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.FieldContainsFold(FieldNote, v))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkUnit applies the HasEdge predicate on the "work_unit" edge.
func HasWorkUnit() predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, WorkUnitTable, WorkUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkUnitWith applies the HasEdge predicate on the "work_unit" edge with a given conditions (other predicates).
func HasWorkUnitWith(preds ...predicate.WorkUnitInfo) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(func(s *sql.Selector) {
		step := newWorkUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStaff applies the HasEdge predicate on the "staff" edge.
func HasStaff() predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, StaffTable, StaffColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStaffWith applies the HasEdge predicate on the "staff" edge with a given conditions (other predicates).
func HasStaffWith(preds ...predicate.Person) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(func(s *sql.Selector) {
		step := newStaffStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WarehouseAssignment) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WarehouseAssignment) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WarehouseAssignment) predicate.WarehouseAssignment {
	return predicate.WarehouseAssignment(sql.NotPredicates(p))
}
