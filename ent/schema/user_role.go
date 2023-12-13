package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

// Annotations of the UserRole.
func (UserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_role"},
	}
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique(),
		field.String("role").NotEmpty(),
	}
}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		// TODO: Link to user
	}
}

// Mixin of the UserRole.
func (UserRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
