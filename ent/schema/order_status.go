package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// OrderStatus holds the schema definition for the OrderStatus entity.
type OrderStatus struct {
	ent.Schema
}

// Annotations of the OrderStatus.
func (OrderStatus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_status"},
	}
}

// Fields of the OrderStatus.
func (OrderStatus) Fields() []ent.Field {
	return nil
}

// Edges of the OrderStatus.
func (OrderStatus) Edges() []ent.Edge {
	return nil
}

// Mixin of the OrderStatus.
func (OrderStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TextCodeMixin{},
	}
}
