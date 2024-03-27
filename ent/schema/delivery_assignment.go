package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DeliveryAssignment holds the schema definition for the DeliveryAssignment entity.
type DeliveryAssignment struct {
	ent.Schema
}

// Annotations of the DeliveryAssignment.
func (DeliveryAssignment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "delivery_assignments"},
	}
}

// Fields of the DeliveryAssignment.
func (DeliveryAssignment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("shipment_id", uuid.UUID{}).
			Immutable(),
		field.UUID("staff_id", uuid.UUID{}).
			Immutable(),
		field.Enum("status").
			NamedValues(
				"Pending", "PENDING",
				"Accepted", "ACCEPTED",
				"Rejected", "REJECTED",
			).
			Default("PENDING"),
		field.String("note").
			Optional().
			Nillable(),
	}
}

// Edges of the DeliveryAssignment.
func (DeliveryAssignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("shipment", Shipment.Type).
			Field("shipment_id").
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

// Mixin of the DeliveryAssignment
func (DeliveryAssignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
