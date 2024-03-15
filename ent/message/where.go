// Code generated by ent, DO NOT EDIT.

package message

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldUpdatedAt, v))
}

// ConversationID applies equality check predicate on the "conversation_id" field. It's identical to ConversationIDEQ.
func ConversationID(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldConversationID, v))
}

// SenderID applies equality check predicate on the "sender_id" field. It's identical to SenderIDEQ.
func SenderID(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSenderID, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldContent, v))
}

// IsRead applies equality check predicate on the "is_read" field. It's identical to IsReadEQ.
func IsRead(v bool) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldIsRead, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldUpdatedAt, v))
}

// ConversationIDEQ applies the EQ predicate on the "conversation_id" field.
func ConversationIDEQ(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldConversationID, v))
}

// ConversationIDNEQ applies the NEQ predicate on the "conversation_id" field.
func ConversationIDNEQ(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldConversationID, v))
}

// ConversationIDIn applies the In predicate on the "conversation_id" field.
func ConversationIDIn(vs ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldConversationID, vs...))
}

// ConversationIDNotIn applies the NotIn predicate on the "conversation_id" field.
func ConversationIDNotIn(vs ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldConversationID, vs...))
}

// SenderIDEQ applies the EQ predicate on the "sender_id" field.
func SenderIDEQ(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSenderID, v))
}

// SenderIDNEQ applies the NEQ predicate on the "sender_id" field.
func SenderIDNEQ(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldSenderID, v))
}

// SenderIDIn applies the In predicate on the "sender_id" field.
func SenderIDIn(vs ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldSenderID, vs...))
}

// SenderIDNotIn applies the NotIn predicate on the "sender_id" field.
func SenderIDNotIn(vs ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldSenderID, vs...))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldType, vs...))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldContent, v))
}

// IsReadEQ applies the EQ predicate on the "is_read" field.
func IsReadEQ(v bool) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldIsRead, v))
}

// IsReadNEQ applies the NEQ predicate on the "is_read" field.
func IsReadNEQ(v bool) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldIsRead, v))
}

// HasConversation applies the HasEdge predicate on the "conversation" edge.
func HasConversation() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ConversationTable, ConversationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConversationWith applies the HasEdge predicate on the "conversation" edge with a given conditions (other predicates).
func HasConversationWith(preds ...predicate.Conversation) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := newConversationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSender applies the HasEdge predicate on the "sender" edge.
func HasSender() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SenderTable, SenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSenderWith applies the HasEdge predicate on the "sender" edge with a given conditions (other predicates).
func HasSenderWith(preds ...predicate.Person) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := newSenderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(sql.NotPredicates(p))
}
