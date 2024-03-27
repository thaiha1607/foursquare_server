package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PersonAddress holds the schema definition for the PersonAddress entity.
type PersonAddress struct {
	ent.Schema
}

// Annotations of the PersonAddress.
func (PersonAddress) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "person_addresses"},
		field.ID("person_id", "address_id"),
	}
}

// Fields of the PersonAddress.
func (PersonAddress) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("person_id", uuid.UUID{}).
			Immutable(),
		field.String("address_id").
			Immutable().
			NotEmpty(),
	}
}

// Edges of the PersonAddress.
func (PersonAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("persons", Person.Type).
			Field("person_id").
			Unique().
			Required().
			Immutable(),
		edge.To("addresses", Address.Type).
			Field("address_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Mixin of the PersonAddress.
func (PersonAddress) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
