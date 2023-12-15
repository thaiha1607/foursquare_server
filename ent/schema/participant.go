package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Participant holds the schema definition for the Participant entity.
type Participant struct {
	ent.Schema
}

// Annotations of the Participant.
func (Participant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "conversation_id"),
	}
}

// Fields of the Participant.
func (Participant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("conversation_id", uuid.UUID{}),
	}
}

// Edges of the Participant.
func (Participant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Field("user_id").
			Unique().
			Required(),
		edge.To("conversations", Conversation.Type).
			Field("conversation_id").
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
