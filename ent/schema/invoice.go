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
		field.String("type").
			NotEmpty(),
	}
}

// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orders", Order.Type).
			Ref("invoices").
			Through("invoice_order_links", InvoiceOrderLink.Type),
		edge.From("delivery_employees", User.Type).
			Ref("delivery_stage_invoices").
			Through("delivery_assignments", DeliveryAssignment.Type),
		edge.From("warehouse_employees", User.Type).
			Ref("warehouse_stage_invoices").
			Through("warehouse_assignments", WarehouseAssignment.Type),
	}
}

// Mixin of the Invoice.
func (Invoice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}