// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUpdatedAt, v))
}

// CustomerID applies equality check predicate on the "customer_id" field. It's identical to CustomerIDEQ.
func CustomerID(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCustomerID, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldNote, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedBy, v))
}

// ParentOrderID applies equality check predicate on the "parent_order_id" field. It's identical to ParentOrderIDEQ.
func ParentOrderID(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldParentOrderID, v))
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldPriority, v))
}

// StatusCode applies equality check predicate on the "status_code" field. It's identical to StatusCodeEQ.
func StatusCode(v int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStatusCode, v))
}

// ManagementStaffID applies equality check predicate on the "management_staff_id" field. It's identical to ManagementStaffIDEQ.
func ManagementStaffID(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldManagementStaffID, v))
}

// WarehouseStaffID applies equality check predicate on the "warehouse_staff_id" field. It's identical to WarehouseStaffIDEQ.
func WarehouseStaffID(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldWarehouseStaffID, v))
}

// DeliveryStaffID applies equality check predicate on the "delivery_staff_id" field. It's identical to DeliveryStaffIDEQ.
func DeliveryStaffID(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDeliveryStaffID, v))
}

// InternalNote applies equality check predicate on the "internal_note" field. It's identical to InternalNoteEQ.
func InternalNote(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldInternalNote, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldUpdatedAt, v))
}

// CustomerIDEQ applies the EQ predicate on the "customer_id" field.
func CustomerIDEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCustomerID, v))
}

// CustomerIDNEQ applies the NEQ predicate on the "customer_id" field.
func CustomerIDNEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldCustomerID, v))
}

// CustomerIDIn applies the In predicate on the "customer_id" field.
func CustomerIDIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldCustomerID, vs...))
}

// CustomerIDNotIn applies the NotIn predicate on the "customer_id" field.
func CustomerIDNotIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldCustomerID, vs...))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldNote, v))
}

// NoteIsNil applies the IsNil predicate on the "note" field.
func NoteIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldNote))
}

// NoteNotNil applies the NotNil predicate on the "note" field.
func NoteNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldNote))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldNote, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// ParentOrderIDEQ applies the EQ predicate on the "parent_order_id" field.
func ParentOrderIDEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldParentOrderID, v))
}

// ParentOrderIDNEQ applies the NEQ predicate on the "parent_order_id" field.
func ParentOrderIDNEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldParentOrderID, v))
}

// ParentOrderIDIn applies the In predicate on the "parent_order_id" field.
func ParentOrderIDIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldParentOrderID, vs...))
}

// ParentOrderIDNotIn applies the NotIn predicate on the "parent_order_id" field.
func ParentOrderIDNotIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldParentOrderID, vs...))
}

// ParentOrderIDIsNil applies the IsNil predicate on the "parent_order_id" field.
func ParentOrderIDIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldParentOrderID))
}

// ParentOrderIDNotNil applies the NotNil predicate on the "parent_order_id" field.
func ParentOrderIDNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldParentOrderID))
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldPriority, v))
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v int) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldPriority, v))
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...int) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldPriority, vs...))
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...int) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldPriority, vs...))
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v int) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldPriority, v))
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v int) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldPriority, v))
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v int) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldPriority, v))
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v int) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldPriority, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldType, vs...))
}

// StatusCodeEQ applies the EQ predicate on the "status_code" field.
func StatusCodeEQ(v int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStatusCode, v))
}

// StatusCodeNEQ applies the NEQ predicate on the "status_code" field.
func StatusCodeNEQ(v int) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldStatusCode, v))
}

// StatusCodeIn applies the In predicate on the "status_code" field.
func StatusCodeIn(vs ...int) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldStatusCode, vs...))
}

// StatusCodeNotIn applies the NotIn predicate on the "status_code" field.
func StatusCodeNotIn(vs ...int) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldStatusCode, vs...))
}

// ManagementStaffIDEQ applies the EQ predicate on the "management_staff_id" field.
func ManagementStaffIDEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldManagementStaffID, v))
}

// ManagementStaffIDNEQ applies the NEQ predicate on the "management_staff_id" field.
func ManagementStaffIDNEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldManagementStaffID, v))
}

// ManagementStaffIDIn applies the In predicate on the "management_staff_id" field.
func ManagementStaffIDIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldManagementStaffID, vs...))
}

// ManagementStaffIDNotIn applies the NotIn predicate on the "management_staff_id" field.
func ManagementStaffIDNotIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldManagementStaffID, vs...))
}

// WarehouseStaffIDEQ applies the EQ predicate on the "warehouse_staff_id" field.
func WarehouseStaffIDEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldWarehouseStaffID, v))
}

// WarehouseStaffIDNEQ applies the NEQ predicate on the "warehouse_staff_id" field.
func WarehouseStaffIDNEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldWarehouseStaffID, v))
}

// WarehouseStaffIDIn applies the In predicate on the "warehouse_staff_id" field.
func WarehouseStaffIDIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldWarehouseStaffID, vs...))
}

// WarehouseStaffIDNotIn applies the NotIn predicate on the "warehouse_staff_id" field.
func WarehouseStaffIDNotIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldWarehouseStaffID, vs...))
}

// WarehouseStaffIDIsNil applies the IsNil predicate on the "warehouse_staff_id" field.
func WarehouseStaffIDIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldWarehouseStaffID))
}

// WarehouseStaffIDNotNil applies the NotNil predicate on the "warehouse_staff_id" field.
func WarehouseStaffIDNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldWarehouseStaffID))
}

// DeliveryStaffIDEQ applies the EQ predicate on the "delivery_staff_id" field.
func DeliveryStaffIDEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDeliveryStaffID, v))
}

// DeliveryStaffIDNEQ applies the NEQ predicate on the "delivery_staff_id" field.
func DeliveryStaffIDNEQ(v uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldDeliveryStaffID, v))
}

// DeliveryStaffIDIn applies the In predicate on the "delivery_staff_id" field.
func DeliveryStaffIDIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldDeliveryStaffID, vs...))
}

// DeliveryStaffIDNotIn applies the NotIn predicate on the "delivery_staff_id" field.
func DeliveryStaffIDNotIn(vs ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldDeliveryStaffID, vs...))
}

// DeliveryStaffIDIsNil applies the IsNil predicate on the "delivery_staff_id" field.
func DeliveryStaffIDIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldDeliveryStaffID))
}

// DeliveryStaffIDNotNil applies the NotNil predicate on the "delivery_staff_id" field.
func DeliveryStaffIDNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldDeliveryStaffID))
}

// InternalNoteEQ applies the EQ predicate on the "internal_note" field.
func InternalNoteEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldInternalNote, v))
}

// InternalNoteNEQ applies the NEQ predicate on the "internal_note" field.
func InternalNoteNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldInternalNote, v))
}

// InternalNoteIn applies the In predicate on the "internal_note" field.
func InternalNoteIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldInternalNote, vs...))
}

// InternalNoteNotIn applies the NotIn predicate on the "internal_note" field.
func InternalNoteNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldInternalNote, vs...))
}

// InternalNoteGT applies the GT predicate on the "internal_note" field.
func InternalNoteGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldInternalNote, v))
}

// InternalNoteGTE applies the GTE predicate on the "internal_note" field.
func InternalNoteGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldInternalNote, v))
}

// InternalNoteLT applies the LT predicate on the "internal_note" field.
func InternalNoteLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldInternalNote, v))
}

// InternalNoteLTE applies the LTE predicate on the "internal_note" field.
func InternalNoteLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldInternalNote, v))
}

// InternalNoteContains applies the Contains predicate on the "internal_note" field.
func InternalNoteContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldInternalNote, v))
}

// InternalNoteHasPrefix applies the HasPrefix predicate on the "internal_note" field.
func InternalNoteHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldInternalNote, v))
}

// InternalNoteHasSuffix applies the HasSuffix predicate on the "internal_note" field.
func InternalNoteHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldInternalNote, v))
}

// InternalNoteIsNil applies the IsNil predicate on the "internal_note" field.
func InternalNoteIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldInternalNote))
}

// InternalNoteNotNil applies the NotNil predicate on the "internal_note" field.
func InternalNoteNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldInternalNote))
}

// InternalNoteEqualFold applies the EqualFold predicate on the "internal_note" field.
func InternalNoteEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldInternalNote, v))
}

// InternalNoteContainsFold applies the ContainsFold predicate on the "internal_note" field.
func InternalNoteContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldInternalNote, v))
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.User) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.User) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParentOrder applies the HasEdge predicate on the "parent_order" edge.
func HasParentOrder() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ParentOrderTable, ParentOrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentOrderWith applies the HasEdge predicate on the "parent_order" edge with a given conditions (other predicates).
func HasParentOrderWith(preds ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newParentOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderStatus applies the HasEdge predicate on the "order_status" edge.
func HasOrderStatus() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrderStatusTable, OrderStatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderStatusWith applies the HasEdge predicate on the "order_status" edge with a given conditions (other predicates).
func HasOrderStatusWith(preds ...predicate.OrderStatusCode) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newOrderStatusStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasManagementStaff applies the HasEdge predicate on the "management_staff" edge.
func HasManagementStaff() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ManagementStaffTable, ManagementStaffColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasManagementStaffWith applies the HasEdge predicate on the "management_staff" edge with a given conditions (other predicates).
func HasManagementStaffWith(preds ...predicate.User) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newManagementStaffStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWarehouseStaff applies the HasEdge predicate on the "warehouse_staff" edge.
func HasWarehouseStaff() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, WarehouseStaffTable, WarehouseStaffColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWarehouseStaffWith applies the HasEdge predicate on the "warehouse_staff" edge with a given conditions (other predicates).
func HasWarehouseStaffWith(preds ...predicate.User) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newWarehouseStaffStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeliveryStaff applies the HasEdge predicate on the "delivery_staff" edge.
func HasDeliveryStaff() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DeliveryStaffTable, DeliveryStaffColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeliveryStaffWith applies the HasEdge predicate on the "delivery_staff" edge with a given conditions (other predicates).
func HasDeliveryStaffWith(preds ...predicate.User) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newDeliveryStaffStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(sql.NotPredicates(p))
}
