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
		entsql.Annotation{Table: "financial_transaction"},
	}
}

// Fields of the FinancialTransaction.
func (FinancialTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("customer_id", uuid.UUID{}),
		field.UUID("invoice_id", uuid.UUID{}),
		field.String("transaction_type"),
		field.Float("amount").GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
		field.String("comment"),
		field.String("payment_method"),
	}
}

// Edges of the FinancialTransaction.
func (FinancialTransaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer", User.Type).
			Field("customer_id").
			Unique().
			Required(),
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required(),
		edge.To("type", TransactionType.Type).
			Field("transaction_type").
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
