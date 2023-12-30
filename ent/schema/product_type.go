package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ProductType holds the schema definition for the ProductType entity.
type ProductType struct {
	ent.Schema
}

// Annotations of the ProductType.
func (ProductType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_types"},
	}
}

// Fields of the ProductType.
func (ProductType) Fields() []ent.Field {
	return []ent.Field{
		field.String("product_type").
			NotEmpty().
			Unique(),
	}
}

// Edges of the ProductType.
func (ProductType) Edges() []ent.Edge {
	return nil
}

// Mixin of the ProductType.
func (ProductType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
