package schema

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
			NotEmpty().
			Unique(),
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			NotEmpty(),
		field.String("phone").
			Optional().
			Nillable().
			Validate(
				func(s string) error {
					e164Regex := `^\+[1-9]\d{1,14}$`
					re := regexp.MustCompile(e164Regex)
					s = strings.ReplaceAll(s, " ", "")
					if re.Find([]byte(s)) == nil {
						return errors.New("invalid phone number (phone number must be in E.164 format)")
					}
					return nil
				},
			).
			Unique(),
		field.Enum("role").
			NamedValues(
				"Salesperson", "SALESPERSON",
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
		field.UUID("work_unit_id", uuid.UUID{}).
			Optional().
			Nillable(),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("work_unit", WorkUnitInfo.Type).
			Field("work_unit_id").
			Unique(),
		edge.To("addresses", Address.Type).
			Through("person_addresses", PersonAddress.Type),
	}
}

// Mixin of the Person.
func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
