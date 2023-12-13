package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// InvoiceType holds the schema definition for the InvoiceType entity.
type InvoiceType struct {
	ent.Schema
}

// Annotations of the InvoiceType.
func (InvoiceType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "invoice_type"},
	}
}

// Fields of the InvoiceType.
func (InvoiceType) Fields() []ent.Field {
	return nil
}

// Edges of the InvoiceType.
func (InvoiceType) Edges() []ent.Edge {
	return nil
}

// Mixin of the InvoiceType.
func (InvoiceType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TextCodeMixin{},
	}
}
