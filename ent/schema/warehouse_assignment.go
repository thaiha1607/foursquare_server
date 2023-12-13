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
		entsql.Annotation{Table: "warehouse_assignment"},
	}
}

// Fields of the WarehouseAssignment.
func (WarehouseAssignment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("invoice_id", uuid.UUID{}),
		field.UUID("staff_id", uuid.UUID{}),
		field.UUID("assigned_by", uuid.UUID{}),
		field.String("status"),
	}
}

// Edges of the WarehouseAssignment.
func (WarehouseAssignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required(),
		edge.To("staff", Employee.Type).
			Field("staff_id").
			Unique().
			Required(),
		edge.To("manager", Employee.Type).
			Field("assigned_by").
			Unique().
			Required(),
		edge.To("warehouse_assignment_status", WarehouseAssignmentStatus.Type).
			Field("status").
			Unique().
			Required(),
	}
}

// Mixin of the WarehouseAssignment.
func (WarehouseAssignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
