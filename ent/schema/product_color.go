package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/hook"
)

// ProductColor holds the schema definition for the ProductColor entity.
type ProductColor struct {
	ent.Schema
}

// Annotations of the ProductColor.
func (ProductColor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_color"},
	}
}

// Fields of the ProductColor.
func (ProductColor) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			MinLen(4),
		field.String("name").
			NotEmpty(),
		field.String("color_code").
			NotEmpty(),
	}
}

// Edges of the ProductColor.
func (ProductColor) Edges() []ent.Edge {
	return nil
}

// Hooks of the ProductColor.
func (ProductColor) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(NanoIDHook(4), ent.OpCreate),
	}
}

// Mixin of the ProductColor.
func (ProductColor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
