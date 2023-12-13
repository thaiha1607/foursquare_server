package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// DeliveryAssignmentStatus holds the schema definition for the DeliveryAssignmentStatus entity.
type DeliveryAssignmentStatus struct {
	ent.Schema
}

// Annotations of the DeliveryAssignmentStatus.
func (DeliveryAssignmentStatus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "delivery_assignment_status"},
	}
}

// Fields of the DeliveryAssignmentStatus.
func (DeliveryAssignmentStatus) Fields() []ent.Field {
	return nil
}

// Edges of the DeliveryAssignmentStatus.
func (DeliveryAssignmentStatus) Edges() []ent.Edge {
	return nil
}

// Mixin of the DeliveryAssignmentStatus.
func (DeliveryAssignmentStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TextCodeMixin{},
	}
}
