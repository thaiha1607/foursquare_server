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
			Default(uuid.New).
			Immutable(),
		field.UUID("invoice_id", uuid.UUID{}).
			Immutable(),
		field.Float("amount").GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}).
			Immutable(),
		field.String("comment").
			Optional().
			Nillable(),
		field.Bool("is_internal").
			Default(false).
			Immutable(),
		field.Enum("payment_method").
			NamedValues(
				"Cash", "CASH",
				"EFT", "ELECTRONIC_FUNDS_TRANSFER",
				"GiftCard", "GIFT_CARD",
				"CreditCard", "CREDIT_CARD",
				"DebitCard", "DEBIT_CARD",
				"PrepaidCard", "PREPAID_CARD",
				"Check", "CHECK",
				"Other", "OTHER",
			).Default("CASH").
			Immutable(),
	}
}

// Edges of the FinancialTransaction.
func (FinancialTransaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Mixin of the FinancialTransaction.
func (FinancialTransaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
