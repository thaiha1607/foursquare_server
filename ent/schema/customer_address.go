package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CustomerAddress holds the schema definition for the CustomerAddress entity.
type CustomerAddress struct {
	ent.Schema
}

// Annotations of the CustomerAddress.
func (CustomerAddress) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "customer_addresses"},
	}
}

// Fields of the CustomerAddress.
func (CustomerAddress) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("customer_id", uuid.UUID{}),
		field.String("line_1").Optional(),
		field.String("line_2").Optional(),
		field.String("line_3").Optional(),
		field.String("line_4").Optional(),
		field.String("city").Optional(),
		field.String("state").Optional(),
		field.String("country").Optional(),
		field.String("postal_code").Optional(),
		field.String("other_details").Optional(),
	}
}

// Edges of the CustomerAddress.
func (CustomerAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer", User.Type).
			Field("customer_id").
			Unique().
			Required(),
	}
}

// Mixin of the CustomerAddress.
func (CustomerAddress) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
