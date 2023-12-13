package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// InvoiceLineItem holds the schema definition for the InvoiceLineItem entity.
type InvoiceLineItem struct {
	ent.Schema
}

// Annotations of the InvoiceLineItem.
func (InvoiceLineItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "invoice_line_item"},
	}
}

// Fields of the InvoiceLineItem.
func (InvoiceLineItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("invoice_id", uuid.UUID{}),
		field.UUID("order_line_item_id", uuid.UUID{}),
		field.Float("qty").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
		field.Float("total_exclusive").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
		field.Float("total_inclusive").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
		field.Float("vat_amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
	}
}

// Edges of the InvoiceLineItem.
func (InvoiceLineItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required(),
		edge.To("order_line_item", OrderLineItem.Type).
			Field("order_line_item_id").
			Unique().
			Required(),
	}
}

// Mixin of the InvoiceLineItem.
func (InvoiceLineItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
