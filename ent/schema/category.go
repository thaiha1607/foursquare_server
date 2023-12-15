package schema

import (
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.UUID("product_id", uuid.UUID{}),
		field.String("name").
			NotEmpty(),
		field.Int("year").
			Positive(),
		field.Float("price").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
		field.JSON("image_url", &url.URL{}).
			Optional(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).
			Field("product_id").
			Unique().
			Required(),
		edge.From("warehouses", Workplace.Type).
			Ref("categories").
			Through("category_quantities", CategoryQuantity.Type),
	}
}

// Mixin of the Category.
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
