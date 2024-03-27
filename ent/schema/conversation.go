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

// Fields of the Conversation.
func (Conversation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("title").
			Optional().
			Nillable(),
		field.UUID("person_one_id", uuid.UUID{}).
			Immutable(),
		field.UUID("person_two_id", uuid.UUID{}).
			Immutable(),
	}
}

// Edges of the Conversation.
func (Conversation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("person_one", Person.Type).
			Field("person_one_id").
			Unique().
			Required().
			Immutable(),
		edge.To("person_two", Person.Type).
			Field("person_two_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the Conversation.
func (Conversation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("person_one_id", "person_two_id").
			Unique(),
	}
}

// Mixin of the Conversation.
func (Conversation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
