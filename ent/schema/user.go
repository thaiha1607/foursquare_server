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
