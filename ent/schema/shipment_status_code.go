package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ShipmentStatusCode holds the schema definition for the ShipmentStatusCode entity.
type ShipmentStatusCode struct {
	ent.Schema
}

// Annotations of the ShipmentStatusCode.
func (ShipmentStatusCode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "shipment_status_codes"},
	}
}

// Fields of the ShipmentStatusCode.
func (ShipmentStatusCode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable().
			Positive().
			Min(1),
		field.String("shipment_status").
			NotEmpty().
			Unique(),
	}
}

// Edges of the ShipmentStatusCode.
func (ShipmentStatusCode) Edges() []ent.Edge {
	return nil
}

// Mixin of the ShipmentStatusCode.
func (ShipmentStatusCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
