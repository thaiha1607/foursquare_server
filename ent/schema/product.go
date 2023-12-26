package schema

import (
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.JSON("image_urls", []url.URL{}).
			Optional(),
		field.String("unit").
			Optional(),
		field.String("type").
			NotEmpty(),
		field.String("provider").
			Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type).
			Through("product_tags", ProductTag.Type),
	}
}

// Mixin of the Product.
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
