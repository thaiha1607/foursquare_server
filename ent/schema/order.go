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
		field.String("status").
			NotEmpty(),
		field.UUID("customer_id", uuid.UUID{}),
		field.UUID("address_id", uuid.UUID{}),
		field.String("phone").
			NotEmpty(),
		field.UUID("parent_order_id", uuid.UUID{}).
			Optional(),
		field.Int("priority").
			Default(0),
		field.String("note").
			Optional(),
		field.String("internal_note").
			Optional(),
		field.String("type").
			NotEmpty(),
		field.UUID("created_by", uuid.UUID{}).
			Optional(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer", User.Type).
			Field("customer_id").
			Unique().
			Required(),
		edge.To("address", CustomerAddress.Type).
			Field("address_id").
			Unique().
			Required(),
		edge.To("parent_order", Order.Type).
			Field("parent_order_id").
			Unique(),
		edge.To("creator", User.Type).
			Field("created_by").
			Unique(),
		edge.To("invoices", Invoice.Type).
			Through("invoice_order_links", InvoiceOrderLink.Type),
	}
}

// Mixin of the Order.
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
