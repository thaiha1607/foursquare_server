package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique(),
		field.String("warehouse_id").
			NotEmpty().
			MaxLen(10),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		// TODO: Link to user
		edge.To("warehouse", Warehouse.Type).
			Field("warehouse_id").
			Unique().
			Required(),
	}
}

// Mixin of the Employee.
func (Employee) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
