package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProductTag holds the schema definition for the ProductTag entity.
type ProductTag struct {
	ent.Schema
}

// Annotations of the ProductTag.
func (ProductTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_tags"},
		field.ID("product_id", "tag_id"),
	}
}

// Fields of the ProductTag.
func (ProductTag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("product_id", uuid.UUID{}),
		field.UUID("tag_id", uuid.UUID{}),
	}
}

// Edges of the ProductTag.
func (ProductTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type).
			Field("product_id").
			Unique().
			Required(),
		edge.To("tags", Tag.Type).
			Field("tag_id").
			Unique().
			Required(),
	}
}

// Mixin of the ProductTag.
func (ProductTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
