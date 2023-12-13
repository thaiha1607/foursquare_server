package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// WarehouseAssignmentStatus holds the schema definition for the WarehouseAssignmentStatus entity.
type WarehouseAssignmentStatus struct {
	ent.Schema
}

// Annotations of the WarehouseAssignmentStatus.
func (WarehouseAssignmentStatus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "warehouse_assignment_status"},
	}
}

// Fields of the WarehouseAssignmentStatus.
func (WarehouseAssignmentStatus) Fields() []ent.Field {
	return nil
}

// Edges of the WarehouseAssignmentStatus.
func (WarehouseAssignmentStatus) Edges() []ent.Edge {
	return nil
}

// Mixin of the WarehouseAssignmentStatus.
func (WarehouseAssignmentStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TextCodeMixin{},
	}
}
