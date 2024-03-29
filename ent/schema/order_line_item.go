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

// OrderLineItem holds the schema definition for the OrderLineItem entity.
type OrderLineItem struct {
	ent.Schema
}

// Annotations of the OrderLineItem.
func (OrderLineItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_line_items"},
	}
}

// Fields of the OrderLineItem.
func (OrderLineItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("order_id", uuid.UUID{}).
			Immutable(),
		field.String("product_id").
			Immutable(),
		field.Float("qty").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
	}
}

// Edges of the OrderLineItem.
func (OrderLineItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			Field("order_id").
			Unique().
			Required().
			Immutable(),
		edge.To("product", Product.Type).
			Field("product_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the OrderLineItem.
func (OrderLineItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id", "product_id").
			Unique(),
	}
}

// Mixin of the OrderLineItem.
func (OrderLineItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
