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
		entsql.Annotation{Table: "invoice_history"},
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
			Nillable().
			Immutable(),
		field.Int("new_status_code").
			Optional().
			Nillable().
			Immutable(),
		field.String("description").
			Optional().
			Nillable().
			Immutable(),
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
		edge.To("old_status", InvoiceStatusCode.Type).
			Field("old_status_code").
			Unique().
			Immutable(),
		edge.To("new_status", InvoiceStatusCode.Type).
			Field("new_status_code").
			Unique().
			Immutable(),
	}
}

// Mixin of the InvoiceHistory.
func (InvoiceHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
