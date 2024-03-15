package schema

import (
	"net/mail"
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("avatar_url").
			Validate(
				func(s string) error {
					_, err := url.Parse(s)
					return err
				},
			).
			Optional(),
		field.String("email").
			Validate(
				func(s string) error {
					_, err := mail.ParseAddress(s)
					return err
				},
			).
			Optional().
			Nillable().
			Unique(),
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			NotEmpty(),
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
		field.Bytes("password_hash").
			NotEmpty().
			Optional().
			Nillable().
			Sensitive(),
		field.Bool("is_email_verified").
			Default(false),
		field.Bool("is_phone_verified").
			Default(false),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return nil
}

// Mixin of the Person.
func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
