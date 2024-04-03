package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OrderHistory holds the schema definition for the OrderHistory entity.
type OrderHistory struct {
	ent.Schema
}

// Annotations of the OrderHistory.
func (OrderHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_history"},
	}
}

// Fields of the OrderHistory.
func (OrderHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("order_id", uuid.UUID{}).
			Immutable(),
		field.UUID("person_id", uuid.UUID{}).
			Immutable(),
		field.Int("old_status_code").
			Optional().
			Nillable(),
		field.Int("new_status_code").
			Optional().
			Nillable(),
		field.String("description").
			Optional().
			Nillable(),
	}
}

// Edges of the OrderHistory.
func (OrderHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			Field("order_id").
			Unique().
			Required().
			Immutable(),
		edge.To("person", Person.Type).
			Field("person_id").
			Unique().
			Required().
			Immutable(),
		edge.To("old_status", OrderStatusCode.Type).
			Field("old_status_code").
			Unique(),
		edge.To("new_status", OrderStatusCode.Type).
			Field("new_status_code").
			Unique(),
	}
}

// Mixin of the OrderHistory.
func (OrderHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
