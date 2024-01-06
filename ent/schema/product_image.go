package schema

import (
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProductImage holds the schema definition for the ProductImage entity.
type ProductImage struct {
	ent.Schema
}

// Annotations of the ProductImage.
func (ProductImage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_images"},
	}
}

// Fields of the ProductImage.
func (ProductImage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("product_id", uuid.UUID{}).
			Immutable(),
		field.JSON("image_url", &url.URL{}).
			Immutable(),
	}
}

// Edges of the ProductImage.
func (ProductImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).
			Field("product_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Mixin of the ProductImage
func (ProductImage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
