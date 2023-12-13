package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Provider holds the schema definition for the Provider entity.
type Provider struct {
	ent.Schema
}

// Fields of the Provider.
func (Provider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique(),
		field.String("description").
			Optional(),
		field.String("address").
			NotEmpty(),
		field.String("phone").
			NotEmpty(),
		field.String("email").
			Optional(),
		field.String("website").
			Optional(),
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
	}
}

// Edges of the Provider.
func (Provider) Edges() []ent.Edge {
	return nil
}

// Mixin of the Provider.
func (Provider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
