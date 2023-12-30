package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PaymentMethod holds the schema definition for the PaymentMethod entity.
type PaymentMethod struct {
	ent.Schema
}

// Annotations of the PaymentMethod.
func (PaymentMethod) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "payment_methods"},
	}
}

// Fields of the PaymentMethod.
func (PaymentMethod) Fields() []ent.Field {
	return []ent.Field{
		field.String("payment_method").
			NotEmpty().
			Unique(),
	}
}

// Edges of the PaymentMethod.
func (PaymentMethod) Edges() []ent.Edge {
	return nil
}

// Mixin of the PaymentMethod.
func (PaymentMethod) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
