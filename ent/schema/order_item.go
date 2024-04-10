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

// OrderItem holds the schema definition for the OrderItem entity.
type OrderItem struct {
	ent.Schema
}

// Annotations of the OrderItem.
func (OrderItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_items"},
	}
}

// Fields of the OrderItem.
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("order_id", uuid.UUID{}).
			Immutable(),
		field.String("product_id").
			Immutable(),
		field.String("product_color_id").
			Immutable(),
		field.UUID("src_unit_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.UUID("dst_unit_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.Float("qty").
			GoType(decimal.Decimal{}).
			Nillable().
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
		field.Float("price_per_unit").
			GoType(decimal.Decimal{}).
			Nillable().
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}),
		field.Enum("status").
			NamedValues(
				"Delivered", "DELIVERED",
				"OutOfStock", "OUT_OF_STOCK",
				"InTransit", "IN_TRANSIT",
				"InStock", "IN_STOCK",
				"PartiallyDelivered", "PARTIALLY_DELIVERED",
			).
			Nillable().
			Default("IN_STOCK"),
	}
}

// Edges of the OrderItem.
func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			Field("order_id").
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
		edge.To("source_work_unit", WorkUnitInfo.Type).
			Field("src_unit_id").
			Unique(),
		edge.To("destination_work_unit", WorkUnitInfo.Type).
			Field("dst_unit_id").
			Unique(),
	}
}

// Indexes of the OrderItem.
func (OrderItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id", "product_id", "product_color_id").
			Unique(),
	}
}

// Mixin of the OrderItem.
func (OrderItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
