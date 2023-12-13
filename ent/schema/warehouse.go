package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Warehouse holds the schema definition for the Warehouse entity.
type Warehouse struct {
	ent.Schema
}

// Fields of the Warehouse.
func (Warehouse) Fields() []ent.Field {
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
		field.UUID("manager_id", uuid.UUID{}).
			Optional(),
	}
}

// Edges of the Warehouse.
func (Warehouse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("manager", Employee.Type).
			Field("manager_id").
			Unique(),
	}
}

// Mixin of the Warehouse.
func (Warehouse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
