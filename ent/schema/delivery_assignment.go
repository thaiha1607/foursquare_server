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
		field.ID("user_id", "invoice_id"),
	}
}

// Fields of the DeliveryAssignment.
func (DeliveryAssignment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("invoice_id", uuid.UUID{}),
		field.UUID("assigned_by", uuid.UUID{}),
		field.String("status").
			NotEmpty(),
	}
}

// Edges of the DeliveryAssignment.
func (DeliveryAssignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("staff", User.Type).
			Field("user_id").
			Unique().
			Required(),
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required(),
		edge.To("manager", User.Type).
			Field("assigned_by").
			Unique().
			Required(),
	}
}

// Mixin of the DeliveryAssignment.
func (DeliveryAssignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
