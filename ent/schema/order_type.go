package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OrderType holds the schema definition for the OrderType entity.
type OrderType struct {
	ent.Schema
}

// Annotations of the OrderType.
func (OrderType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_types"},
	}
}

// Fields of the OrderType.
func (OrderType) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_type").
			NotEmpty().
			Unique(),
	}
}

// Edges of the OrderType.
func (OrderType) Edges() []ent.Edge {
	return nil
}

// Mixin of the OrderType.
func (OrderType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
