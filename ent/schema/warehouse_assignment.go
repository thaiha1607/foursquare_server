package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WarehouseAssignment holds the schema definition for the WarehouseAssignment entity.
type WarehouseAssignment struct {
	ent.Schema
}

// Annotations of the WarehouseAssignment.
func (WarehouseAssignment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "warehouse_assignments"},
	}
}

// Fields of the WarehouseAssignment.
func (WarehouseAssignment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("order_id", uuid.UUID{}).
			Immutable(),
		field.UUID("work_unit_id", uuid.UUID{}).
			Immutable(),
		field.UUID("staff_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.Enum("status").
			NamedValues(
				"Pending", "PENDING",
				"Accepted", "ACCEPTED",
				"Rejected", "REJECTED",
			).
			Nillable().
			Default("PENDING"),
		field.String("note").
			Optional().
			Nillable(),
	}
}

// Edges of the WarehouseAssignment.
func (WarehouseAssignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			Field("order_id").
			Unique().
			Required().
			Immutable(),
		edge.To("work_unit", WorkUnitInfo.Type).
			Field("work_unit_id").
			Unique().
			Required().
			Immutable(),
		edge.To("staff", Person.Type).
			Field("staff_id").
			Unique(),
	}
}

// Mixin of the WarehouseAssignment
func (WarehouseAssignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
