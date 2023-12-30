package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// InvoiceStatusCode holds the schema definition for the InvoiceStatusCode entity.
type InvoiceStatusCode struct {
	ent.Schema
}

// Annotations of the InvoiceStatusCode.
func (InvoiceStatusCode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "invoice_status_codes"},
	}
}

// Fields of the InvoiceStatusCode.
func (InvoiceStatusCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("invoice_status").
			NotEmpty().
			Unique(),
	}
}

// Edges of the InvoiceStatusCode.
func (InvoiceStatusCode) Edges() []ent.Edge {
	return nil
}

// Mixin of the InvoiceStatusCode.
func (InvoiceStatusCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
