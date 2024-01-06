package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/thaiha1607/foursquare_server/ent/hook"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("user_id", uuid.UUID{}).
			Immutable(),
		field.String("username").
			NotEmpty().
			Unique().
			MaxLen(10),
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
		hook.On(CreateUsernameHook(), ent.OpCreate),
	}
}

func CreateUsernameHook() ent.Hook {
	type UsernameSetter interface {
		SetUsername(string)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			is, ok := m.(UsernameSetter)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation %T", m)
			}
			id, err := gonanoid.New(10)
			if err != nil {
				return nil, err
			}
			is.SetUsername(id)
			return next.Mutate(ctx, m)
		})
	}
}
