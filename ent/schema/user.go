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
			),
		field.Bool("email_visible").
			Default(true),
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("last_reset").
			Optional(),
		field.Time("last_verification").
			Optional(),
		field.String("name").
			NotEmpty(),
		field.Text("password_hash").
			NotEmpty().
			Sensitive(),
		field.String("username").
			NotEmpty(),
		field.Bool("verified").
			Default(false),
		field.String("phone").
			NotEmpty().
			Unique(),
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
