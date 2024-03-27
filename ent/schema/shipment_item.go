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

// ShipmentItem holds the schema definition for the ShipmentItem entity.
type ShipmentItem struct {
	ent.Schema
}

// Annotations of the ShipmentItem.
func (ShipmentItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "invoice_line_items"},
	}
}

// Fields of the ShipmentItem.
func (ShipmentItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("shipment_id", uuid.UUID{}).
			Immutable(),
		field.UUID("order_item_id", uuid.UUID{}).
			Immutable(),
		field.Float("qty").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}).
			Immutable(),
		field.Float("total").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
				dialect.MySQL:    "decimal(12,2)",
			}).
			Immutable(),
	}
}

// Edges of the ShipmentItem.
func (ShipmentItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("shipment", Shipment.Type).
			Field("shipment_id").
			Unique().
			Required().
			Immutable(),
		edge.To("order_item", OrderItem.Type).
			Field("order_item_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the ShipmentItem.
func (ShipmentItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("shipment_id", "order_item_id").
			Unique(),
	}
}

// Mixin of the ShipmentItem.
func (ShipmentItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
