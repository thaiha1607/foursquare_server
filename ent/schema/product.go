package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"github.com/thaiha1607/foursquare_server/ent/hook"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional().
			Nillable().
			Default(""),
		field.String("id").
			NotEmpty().
			Unique().
			MinLen(6),
		field.Int("year").
			Optional().
			Nillable().
			Positive().
			Min(1900),
		field.Float("price").
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
		field.Strings("colors").
			Optional(),
		field.String("provider").
			Optional().
			Nillable(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type).
			Through("product_tags", ProductTag.Type),
	}
}

// Hooks of the Product.
func (Product) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(NanoIDHook(6), ent.OpCreate),
	}
}

// Mixin of the Product.
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
