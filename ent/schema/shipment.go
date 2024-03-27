package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/thaiha1607/foursquare_server/ent/hook"
)

// Shipment holds the schema definition for the Shipment entity.
type Shipment struct {
	ent.Schema
}

// Fields of the Shipment.
func (Shipment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("order_id", uuid.UUID{}).
			Immutable(),
		field.UUID("invoice_id", uuid.UUID{}).
			Immutable(),
		field.String("shipment_tracking_number").
			MaxLen(26).
			NotEmpty(),
		field.Time("shipment_date"),
		field.String("note").
			Optional(),
	}
}

// Edges of the Shipment.
func (Shipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			Field("order_id").
			Unique().
			Required().
			Immutable(),
		edge.To("invoice", Invoice.Type).
			Field("invoice_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Hooks of the Shipment.
func (Shipment) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func() ent.Hook {
			type ShipmentTrackingNumberSetter interface {
				SetShipmentTrackingNumber(string)
			}
			return func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					is, ok := m.(ShipmentTrackingNumberSetter)
					if !ok {
						return nil, fmt.Errorf("unexpected mutation %T", m)
					}
					id := ulid.Make().String()
					is.SetShipmentTrackingNumber(id)
					return next.Mutate(ctx, m)
				})
			}
		}(), ent.OpCreate),
	}
}

// Mixin of the Shipment.
func (Shipment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
