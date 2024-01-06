package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.UUID("customer_id", uuid.UUID{}).
			Immutable(),
		field.String("note").
			Optional().
			Nillable(),
		field.UUID("created_by", uuid.UUID{}).
			Immutable(),
		field.UUID("parent_order_id", uuid.UUID{}).
			Optional(),
		field.Int("priority").
			Default(0).
			Range(0, 10),
		field.Enum("type").
			NamedValues(
				"Sale", "SALE",
				"Return", "RETURN",
				"Exchange", "EXCHANGE",
				"Transfer", "TRANSFER",
				"Internal", "INTERNAL",
				"Other", "OTHER",
			).
			Default("SALE").
			Immutable(),
		field.Int("status_code").
			Default(1),
		field.UUID("management_staff_id", uuid.UUID{}),
		field.UUID("warehouse_staff_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.UUID("delivery_staff_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.String("internal_note").
			Optional().
			Nillable(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer", User.Type).
			Field("customer_id").
			Unique().
			Required().
			Immutable(),
		edge.To("creator", User.Type).
			Field("created_by").
			Unique().
			Required().
			Immutable(),
		edge.To("parent_order", Order.Type).
			Field("parent_order_id").
			Unique(),
		edge.To("order_status", OrderStatusCode.Type).
			Field("status_code").
			Unique().
			Required(),
		edge.To("management_staff", User.Type).
			Field("management_staff_id").
			Unique().
			Required(),
		edge.To("warehouse_staff", User.Type).
			Field("warehouse_staff_id").
			Unique(),
		edge.To("delivery_staff", User.Type).
			Field("delivery_staff_id").
			Unique(),
	}
}

// Mixin of the Order.
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
