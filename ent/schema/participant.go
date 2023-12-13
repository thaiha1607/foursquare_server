package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Participant holds the schema definition for the Participant entity.
type Participant struct {
	ent.Schema
}

// Fields of the Participant.
func (Participant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("conversation_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
	}
}

// Edges of the Participant.
func (Participant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("conversation", Conversation.Type).
			Field("conversation_id").
			Unique().
			Required(),
		edge.To("user", Employee.Type).
			Field("user_id").
			Unique().
			Required(),
	}
}

// Mixin of the Participant.
func (Participant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
