package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/hook"
)

// ProductInfo holds the schema definition for the ProductInfo entity.
type ProductInfo struct {
	ent.Schema
}

// Annotations of the ProductInfo.
func (ProductInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_info"},
	}
}

// Fields of the ProductInfo.
func (ProductInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			MinLen(6),
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional().
			Nillable().
			Default(""),
		field.Int("year").
			Optional().
			Nillable().
			Positive().
			Min(1900),
		field.String("provider").
			Optional().
			Nillable(),
	}
}

// Edges of the ProductInfo.
func (ProductInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type).
			Through("product_tags", ProductTag.Type),
	}
}

// Hooks of the ProductInfo.
func (ProductInfo) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(NanoIDHook(6), ent.OpCreate),
	}
}

// Mixin of the ProductInfo.
func (ProductInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
