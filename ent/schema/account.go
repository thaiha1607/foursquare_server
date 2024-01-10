package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/hook"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			MinLen(10),
		field.UUID("user_id", uuid.UUID{}).
			Immutable(),
		field.Time("last_reset").
			Optional().
			Nillable(),
		field.Time("last_email_verification").
			Optional().
			Nillable(),
		field.Time("last_phone_verification").
			Optional().
			Nillable(),
		field.Bool("is_email_verified").
			Default(false),
		field.Bool("is_phone_verified").
			Default(false),
		field.Enum("role").
			NamedValues(
				"Admin", "ADMIN",
				"Customer", "CUSTOMER",
				"Warehouse", "WAREHOUSE",
				"Delivery", "DELIVERY",
				"Management", "MANAGEMENT",
			).Default("CUSTOMER"),
		field.Text("password_hash").
			NotEmpty().
			Optional().
			Nillable().
			Sensitive(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Mixin of the Account.
func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Hooks of the Account.
func (Account) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(10), ent.OpCreate),
	}
}
