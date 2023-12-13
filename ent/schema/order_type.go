package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// OrderType holds the schema definition for the OrderType entity.
type OrderType struct {
	ent.Schema
}

// Annotations of the OrderType.
func (OrderType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_type"},
	}
}

// Fields of the OrderType.
func (OrderType) Fields() []ent.Field {
	return nil
}

// Edges of the OrderType.
func (OrderType) Edges() []ent.Edge {
	return nil
}

// Mixin of the OrderType.
func (OrderType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TextCodeMixin{},
	}
}
