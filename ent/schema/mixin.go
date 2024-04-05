package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid/v2"
	"github.com/thaiha1607/foursquare_server/ent/hook"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

type ULIDMixin struct {
	mixin.Schema
}

func (ULIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			NotEmpty().
			Validate(func(s string) (err error) {
				_, err = ulid.ParseStrict(s)
				return
			}),
	}
}

func (ULIDMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(ULIDHook(), ent.OpCreate),
	}
}
