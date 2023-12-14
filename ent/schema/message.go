package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("conversation_id", uuid.UUID{}),
		field.UUID("sender_id", uuid.UUID{}),
		field.String("type").
			Optional(),
		field.String("content").Optional(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("conversation", Conversation.Type).
			Field("conversation_id").
			Unique().
			Required(),
		edge.To("sender", User.Type).
			Field("sender_id").
			Unique().
			Required(),
		edge.From("seen_recipients", User.Type).
			Ref("messages").
			Through("message_read_states", MessageReadState.Type),
	}
}

// Mixin of the Message.
func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
