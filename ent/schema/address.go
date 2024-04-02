package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("line1").
			NotEmpty(),
		field.String("line2").
			Optional().
			Nillable(),
		field.String("city").
			NotEmpty(),
		field.String("state_or_province").
			NotEmpty(),
		field.String("zip_or_postcode").
			NotEmpty(),
		field.String("country").
			NotEmpty(),
		field.String("other_address_details").
			Optional().
			Nillable(),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("persons", Person.Type).
			Ref("addresses").
			Through("person_addresses", PersonAddress.Type),
	}
}

// Mixin of the Address.
func (Address) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
