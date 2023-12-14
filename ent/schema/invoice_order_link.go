package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// InvoiceOrderLink holds the schema definition for the InvoiceOrderLink entity.
type InvoiceOrderLink struct {
	ent.Schema
}

// Annotations of the InvoiceOrderLink.
func (InvoiceOrderLink) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "invoice_order_links"},
		field.ID("order_id", "invoice_id"),
	}
}

// Fields of the InvoiceOrderLink.
func (InvoiceOrderLink) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("order_id", uuid.UUID{}),
		field.UUID("invoice_id", uuid.UUID{}),
	}
}

// Edges of the InvoiceOrderLink.
func (InvoiceOrderLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			Field("order_id").
			Unique().
			Required(),
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required(),
	}
}

// Mixin of the InvoiceOrderLink.
func (InvoiceOrderLink) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
