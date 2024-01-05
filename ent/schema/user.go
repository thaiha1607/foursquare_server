package schema

import (
	"net/mail"
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("avatar_url", &url.URL{}).
			Optional(),
		field.String("email").
			NotEmpty().
			Validate(
				func(s string) error {
					_, err := mail.ParseAddress(s)
					return err
				},
			).
			Unique(),
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("last_reset").
			Optional().
			Nillable(),
		field.Time("last_verification").
			Optional().
			Nillable(),
		field.String("name").
			NotEmpty(),
		field.Text("password_hash").
			NotEmpty().
			Sensitive(),
		field.Bool("verified").
			Default(false),
		field.String("phone").
			NotEmpty().
			Unique(),
		field.Enum("role").
			NamedValues(
				"Admin", "ADMIN",
				"Customer", "CUSTOMER",
				"Warehouse", "WAREHOUSE",
				"Delivery", "DELIVERY",
				"Management", "MANAGEMENT",
			).Default("CUSTOMER"),
		field.String("address").
			Optional().
			Nillable(),
		field.String("postal_code").
			Optional().
			Nillable(),
		field.String("other_address_info").
			Optional().
			Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
