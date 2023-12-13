package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// PaymentMethod holds the schema definition for the PaymentMethod entity.
type PaymentMethod struct {
	ent.Schema
}

// Annotations of the PaymentMethod.
func (PaymentMethod) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "payment_method"},
	}
}

// Fields of the PaymentMethod.
func (PaymentMethod) Fields() []ent.Field {
	return nil
}

// Edges of the PaymentMethod.
func (PaymentMethod) Edges() []ent.Edge {
	return nil
}

// Mixin of the PaymentMethod.
func (PaymentMethod) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TextCodeMixin{},
	}
}
