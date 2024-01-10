package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OrderStatusCode holds the schema definition for the OrderStatusCode entity.
type OrderStatusCode struct {
	ent.Schema
}

// Annotations of the OrderStatusCode.
func (OrderStatusCode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_status_codes"},
	}
}

// Fields of the OrderStatusCode.
func (OrderStatusCode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable().
			Positive().
			Min(1),
		field.String("order_status").
			NotEmpty().
			Unique(),
	}
}

// Edges of the OrderStatusCode.
func (OrderStatusCode) Edges() []ent.Edge {
	return nil
}

// Mixin of the OrderStatusCode.
func (OrderStatusCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
