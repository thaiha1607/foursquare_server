package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

// Annotations of the UserRole.
func (UserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_roles"},
	}
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_role").
			NotEmpty().
			Unique(),
	}
}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return nil
}

// Mixin of the UserRole.
func (UserRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
