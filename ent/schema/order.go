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
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("customer_id", uuid.UUID{}).
			Immutable(),
		field.String("note").
			Optional().
			Nillable(),
		field.UUID("created_by", uuid.UUID{}).
			Immutable(),
		field.UUID("parent_order_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.Int("priority").
			Default(0).
			Range(0, 100),
		field.Enum("type").
			NamedValues(
				"Sale", "SALE",
				"Return", "RETURN",
				"Exchange", "EXCHANGE",
				"Transfer", "TRANSFER",
				"Other", "OTHER",
			).
			Default("SALE").
			Immutable(),
		field.Int("status_code").
			Default(1),
		field.UUID("staff_id", uuid.UUID{}),
		field.String("internal_note").
			Optional().
			Nillable(),
		field.Bool("is_internal").
			Default(false),
		field.String("address_id").
			NotEmpty(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer", Person.Type).
			Field("customer_id").
			Unique().
			Required().
			Immutable(),
		edge.To("creator", Person.Type).
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
		edge.To("staff", Person.Type).
			Field("staff_id").
			Unique().
			Required(),
		edge.To("order_address", Address.Type).
			Field("address_id").
			Unique().
			Required(),
	}
}

// Mixin of the Order.
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
