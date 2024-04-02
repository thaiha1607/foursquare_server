package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ShipmentHistory holds the schema definition for the ShipmentHistory entity.
type ShipmentHistory struct {
	ent.Schema
}

// Annotations of the ShipmentHistory.
func (ShipmentHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "shipment_histories"},
	}
}

// Fields of the ShipmentHistory.
func (ShipmentHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("shipment_id").
			NotEmpty().
			Immutable(),
		field.UUID("person_id", uuid.UUID{}).
			Immutable(),
		field.Int("prev_status_code").
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

// Edges of the ShipmentHistory.
func (ShipmentHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("shipment", Shipment.Type).
			Field("shipment_id").
			Unique().
			Required().
			Immutable(),
		edge.To("person", Person.Type).
			Field("person_id").
			Unique().
			Required().
			Immutable(),
		edge.To("prev_status", OrderStatusCode.Type).
			Field("prev_status_code").
			Unique(),
		edge.To("new_status", OrderStatusCode.Type).
			Field("new_status_code").
			Unique(),
	}
}
