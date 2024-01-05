package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		entsql.Annotation{Table: "invoice_line_items"},
	}
}

// Fields of the InvoiceLineItem.
func (InvoiceLineItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("invoice_id", uuid.UUID{}).
			Immutable(),
		field.UUID("order_line_item_id", uuid.UUID{}).
			Immutable(),
		field.Float("qty").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}).
			Immutable(),
		field.Float("total").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}).
			Immutable(),
	}
}

// Edges of the InvoiceLineItem.
func (InvoiceLineItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required().
			Immutable(),
		edge.To("order_line_item", OrderLineItem.Type).
			Field("order_line_item_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the InvoiceLineItem.
func (InvoiceLineItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("invoice_id", "order_line_item_id").
			Unique(),
	}
}

// Mixin of the InvoiceLineItem.
func (InvoiceLineItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
