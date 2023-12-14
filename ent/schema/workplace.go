package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Workplace holds the schema definition for the Workplace entity.
type Workplace struct {
	ent.Schema
}

// Fields of the Workplace.
func (Workplace) Fields() []ent.Field {
	return []ent.Field{
		// Length of ID is 10 characters.
		// ID has the following format:
		// <ISO 3166-1 alpha-2 country code>-<3 characters of city name>-<5 random characters>
		// Example: VN-HCM-12345
		field.String("id").
			Match(regexp.MustCompile(`^[A-Z]{2}-[A-Z]{3}-[0-9]{5}$`)).
			NotEmpty().
			Unique(),
		field.String("name").
			NotEmpty(),
		field.String("address").
			NotEmpty(),
		field.Bool("is_warehouse").
			Default(false),
		field.UUID("manager_id", uuid.UUID{}).
			Optional(),
	}
}

// Edges of the Workplace.
func (Workplace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("manager", User.Type).
			Field("manager_id").
			Unique(),
		edge.To("categories", Category.Type).
			Through("category_quantities", CategoryQuantity.Type),
	}
}

// Mixin of the Workplace.
func (Workplace) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
