package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Shipment holds the schema definition for the Shipment entity.
type Shipment struct {
	ent.Schema
}

// Fields of the Shipment.
func (Shipment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("order_id", uuid.UUID{}).
			Immutable(),
		field.UUID("invoice_id", uuid.UUID{}).
			Immutable(),
		field.UUID("staff_id", uuid.UUID{}).
			Immutable(),
		field.Time("shipment_date"),
		field.String("note").
			Optional(),
		field.Enum("status").
			NamedValues(
				"Pending", "PENDING",
				"Accepted", "ACCEPTED",
				"Rejected", "REJECTED",
			).
			Default("PENDING"),
	}
}

// Edges of the Shipment.
func (Shipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			Field("order_id").
			Unique().
			Required().
			Immutable(),
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required().
			Immutable(),
		edge.To("staff", Person.Type).
			Field("staff_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Mixin of the Shipment.
func (Shipment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		ULIDMixin{},
	}
}
