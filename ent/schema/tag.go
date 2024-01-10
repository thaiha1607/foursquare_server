package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/hook"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			MinLen(6),
		field.String("title").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).
			Ref("tags").
			Through("product_tags", ProductTag.Type),
	}
}

// Hooks of the Tag.
func (Tag) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(6), ent.OpCreate),
	}
}

// Mixin of the Tag.
func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
