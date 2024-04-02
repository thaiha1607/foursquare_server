package schema

import (
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WorkUnitInfo holds the schema definition for the WorkUnitInfo entity.
type WorkUnitInfo struct {
	ent.Schema
}

// Annotations of the WorkUnitInfo.
func (WorkUnitInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "work_unit_info"},
	}
}

// Fields of the WorkUnitInfo.
func (WorkUnitInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			NotEmpty(),
		field.UUID("address_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.Enum("type").
			NamedValues(
				"Warehouse", "WAREHOUSE",
				"Office", "OFFICE",
				"Delivery", "DELIVERY",
			).
			Default("WAREHOUSE"),
		field.String("image_url").
			Validate(
				func(s string) error {
					_, err := url.Parse(s)
					return err
				},
			).
			Optional().
			Nillable(),
	}
}

// Edges of the WorkUnitInfo.
func (WorkUnitInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("address", Address.Type).
			Field("address_id").
			Unique(),
	}
}
