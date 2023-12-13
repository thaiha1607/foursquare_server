package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// ProductType holds the schema definition for the ProductType entity.
type ProductType struct {
	ent.Schema
}

// Annotations of the ProductType.
func (ProductType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_type"},
	}
}

// Fields of the ProductType.
func (ProductType) Fields() []ent.Field {
	return nil
}

// Edges of the ProductType.
func (ProductType) Edges() []ent.Edge {
	return nil
}

// Mixin of the ProductType.
func (ProductType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TextCodeMixin{},
	}
}
