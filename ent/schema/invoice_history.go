package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// InvoiceHistory holds the schema definition for the InvoiceHistory entity.
type InvoiceHistory struct {
	ent.Schema
}

// Annotations of the InvoiceHistory.
func (InvoiceHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "invoice_histories"},
	}
}

// Fields of the InvoiceHistory.
func (InvoiceHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("invoice_id", uuid.UUID{}).
			Immutable(),
		field.UUID("person_id", uuid.UUID{}).
			Immutable(),
		field.Int("old_status_code").
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

// Edges of the InvoiceHistory.
func (InvoiceHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required().
			Immutable(),
		edge.To("person", Person.Type).
			Field("person_id").
			Unique().
			Required().
			Immutable(),
		edge.To("old_status", OrderStatusCode.Type).
			Field("old_status_code").
			Unique(),
		edge.To("new_status", OrderStatusCode.Type).
			Field("new_status_code").
			Unique(),
	}
}

// Mixin of the InvoiceHistory.
func (InvoiceHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
