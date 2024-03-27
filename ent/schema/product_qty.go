package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ProductQty holds the schema definition for the ProductQty entity.
type ProductQty struct {
	ent.Schema
}

// Annotations of the ProductQty.
func (ProductQty) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_qty"},
	}
}

// Fields of the ProductQty.
func (ProductQty) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("work_unit_id", uuid.UUID{}).
			Immutable(),
		field.String("product_id").
			Immutable().
			NotEmpty(),
		field.String("product_color_id").
			Immutable().
			NotEmpty(),
		field.Float("price_per_unit").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
		field.Float("qty").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
	}
}

// Edges of the ProductQty.
func (ProductQty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("work_unit", WorkUnitInfo.Type).
			Field("work_unit_id").
			Unique().
			Required().
			Immutable(),
		edge.To("product", ProductInfo.Type).
			Field("product_id").
			Unique().
			Required().
			Immutable(),
		edge.To("product_color", ProductColor.Type).
			Field("product_color_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the ProductQty.
func (ProductQty) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("work_unit_id", "product_id", "product_color_id").
			Unique(),
	}
}

// Mixin of the ProductQty
func (ProductQty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
