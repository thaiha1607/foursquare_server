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
			Default(uuid.New).
			Immutable(),
		field.UUID("order_id", uuid.UUID{}).
			Immutable(),
		field.Float("total").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}).
			Immutable(),
		field.String("note").
			Optional().
			Nillable(),
		field.Enum("type").
			NamedValues(
				"ProForma", "PRO_FORMA",
				"Regular", "REGULAR",
				"PastDue", "PAST_DUE",
				"Interim", "INTERIM",
				"Timesheet", "TIMESHEET",
				"Final", "FINAL",
				"Credit", "CREDIT",
				"Debit", "DEBIT",
				"Mixed", "MIXED",
				"Commercial", "COMMERCIAL",
				"Recurring", "RECURRING",
				"Other", "OTHER",
			).Default("PRO_FORMA").
			Immutable(),
		/*
			Status we might want to consider:
			- Draft
			- Active
			- Sent
			- Disputed
			- Overdue
			- Partial
			- Paid
			- Void
			- Debt
			- Other
		*/
		field.Int("status_code").
			Default(1),
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
			Optional().
			Nillable(),
	}
}

// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			Field("order_id").
			Unique().
			Required().
			Immutable(),
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
