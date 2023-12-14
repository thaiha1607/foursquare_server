package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MessageReadState holds the schema definition for the MessageReadState entity.
type MessageReadState struct {
	ent.Schema
}

// Annotations of the MessageReadState.
func (MessageReadState) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "message_read_states"},
		field.ID("user_id", "message_id"),
	}
}

// Fields of the MessageReadState.
func (MessageReadState) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("message_id", uuid.UUID{}),
		field.Time("read_at").
			Immutable().
			Default(time.Now),
	}
}

// Edges of the MessageReadState.
func (MessageReadState) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required(),
		edge.To("message", Message.Type).
			Field("message_id").
			Unique().
			Required(),
	}
}
