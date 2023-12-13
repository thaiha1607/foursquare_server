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
		field.UUID("provider_id", uuid.UUID{}).
			Optional(),
		field.String("unit").
			Optional(),
		field.String("type").
			Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provider", Provider.Type).
			Field("provider_id").
			Unique(),
		edge.To("product_type", ProductType.Type).
			Field("type").
			Unique(),
	}
}

// Mixin of the Product.
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
