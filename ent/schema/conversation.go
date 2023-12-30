package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Conversation holds the schema definition for the Conversation entity.
type Conversation struct {
	ent.Schema
}

// Indexes of the Conversation.
func (Conversation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_one_id", "user_two_id").
			Unique(),
	}
}

// Fields of the Conversation.
func (Conversation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("title").
			Optional().
			Nillable(),
		field.UUID("user_one_id", uuid.UUID{}),
		field.UUID("user_two_id", uuid.UUID{}),
	}
}

// Edges of the Conversation.
func (Conversation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_one", User.Type).
			Field("user_one_id").
			Unique().
			Required(),
		edge.To("user_two", User.Type).
			Field("user_two_id").
			Unique().
			Required(),
	}
}

// Mixin of the Conversation.
func (Conversation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
