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

// FinancialTransaction holds the schema definition for the FinancialTransaction entity.
type FinancialTransaction struct {
	ent.Schema
}

// Annotations of the FinancialTransaction.
func (FinancialTransaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "financial_transactions"},
	}
}

// Fields of the FinancialTransaction.
func (FinancialTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("invoice_id", uuid.UUID{}),
		field.Float("amount").GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
		field.String("comment").
			Optional().
			Nillable(),
		field.Int("type"),
		field.Int("payment_method"),
	}
}

// Edges of the FinancialTransaction.
func (FinancialTransaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required(),
		edge.To("transaction_type", TransactionType.Type).
			Field("type").
			Unique().
			Required(),
		edge.To("payment", PaymentMethod.Type).
			Field("payment_method").
			Unique().
			Required(),
	}
}

// Mixin of the FinancialTransaction.
func (FinancialTransaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
