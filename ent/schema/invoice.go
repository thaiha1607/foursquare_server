package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Invoice holds the schema definition for the Invoice entity.
type Invoice struct {
	ent.Schema
}

// Fields of the Invoice.
func (Invoice) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("order_id", uuid.UUID{}),
		field.Float("total").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
		field.Int("type"),
		field.Int("status_code"),
	}
}

// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			Field("order_id").
			Unique().
			Required(),
		edge.To("invoice_type", InvoiceType.Type).
			Field("type").
			Unique().
			Required(),
		edge.To("invoice_status", InvoiceStatusCode.Type).
			Field("status_code").
			Unique().
			Required(),
	}
}

// Mixin of the Invoice.
func (Invoice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
