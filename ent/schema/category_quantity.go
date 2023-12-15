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

// CategoryQuantity holds the schema definition for the CategoryQuantity entity.
type CategoryQuantity struct {
	ent.Schema
}

// Annotations of the CategoryQuantity.
func (CategoryQuantity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "category_quantities"},
		field.ID("workplace_id", "category_id"),
	}
}

// Fields of the CategoryQuantity.
func (CategoryQuantity) Fields() []ent.Field {
	return []ent.Field{
		field.String("workplace_id").
			NotEmpty(),
		field.UUID("category_id", uuid.UUID{}),
		field.Float("qty").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
	}
}

// Edges of the CategoryQuantity.
func (CategoryQuantity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("warehouses", Workplace.Type).
			Field("workplace_id").
			Unique().
			Required(),
		edge.To("categories", Category.Type).
			Field("category_id").
			Unique().
			Required(),
	}
}

// Mixin of the CategoryQuantity.
func (CategoryQuantity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
