package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Conversation holds the schema definition for the Conversation entity.
type Conversation struct {
	ent.Schema
}

// Fields of the Conversation.
func (Conversation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("title").Optional(),
	}
}

// Edges of the Conversation.
func (Conversation) Edges() []ent.Edge {
	return nil
}

// Mixin of the Conversation.
func (Conversation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
