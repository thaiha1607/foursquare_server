package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/thaiha1607/foursquare_server/ent/hook"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Nillable().
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Nillable().
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

type CreationTimeMixin struct {
	mixin.Schema
}

func (CreationTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Nillable().
			Immutable().
			Default(time.Now),
	}
}

type ULIDMixin struct {
	mixin.Schema
}

func (ULIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			NotEmpty(),
	}
}

func (ULIDMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(ULIDHook(), ent.OpCreate),
	}
}
