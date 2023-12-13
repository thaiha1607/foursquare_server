package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ProductQuantity holds the schema definition for the ProductQuantity entity.
type ProductQuantity struct {
	ent.Schema
}

// Annotations of the ProductQuantity.
func (ProductQuantity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_quantity"},
	}
}

// Fields of the ProductQuantity.
func (ProductQuantity) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.UUID("category_id", uuid.UUID{}),
		field.String("warehouse_id").
			NotEmpty().
			MaxLen(10),
		field.Float("qty").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
	}
}

// Edges of the ProductQuantity.
func (ProductQuantity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("category", Category.Type).
			Field("category_id").
			Unique().
			Required(),
		edge.To("warehouse", Warehouse.Type).
			Field("warehouse_id").
			Unique().
			Required(),
	}
}

// Mixin of the ProductQuantity.
func (ProductQuantity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
