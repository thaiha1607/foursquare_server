package schema

import (
	"net/mail"
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("avatar_url", &url.URL{}).
			Optional(),
		field.String("email").
			NotEmpty().
			Validate(
				func(s string) error {
					_, err := mail.ParseAddress(s)
					return err
				},
			),
		field.Bool("email_visible").
			Default(true),
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("last_reset").
			Optional(),
		field.Time("last_verification").
			Optional(),
		field.String("name").
			NotEmpty(),
		field.Text("password_hash").
			NotEmpty().
			Sensitive(),
		field.String("username").
			NotEmpty(),
		field.Bool("verified").
			Default(false),
		field.String("phone").
			NotEmpty().
			Unique(),
		field.String("role").
			NotEmpty().
			Default("customer"),
		field.String("workplace_id").
			MaxLen(10).
			MinLen(10).
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("workplace", Workplace.Type).
			Unique().
			Field("workplace_id"),
		edge.To("delivery_stage_invoices", Invoice.Type).
			Through("delivery_assignments", DeliveryAssignment.Type),
		edge.To("conversations", Conversation.Type).
			Through("participants", Participant.Type),
		edge.To("messages", Message.Type).
			Through("message_read_states", MessageReadState.Type),
		edge.To("warehouse_stage_invoices", Invoice.Type).
			Through("warehouse_assignments", WarehouseAssignment.Type),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
