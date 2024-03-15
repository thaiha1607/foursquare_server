package schema

import (
	"context"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/thaiha1607/foursquare_server/ent"
)

// NanoIDHook is a hook that generates a random ID for the entity.
// Only applies to ID fields that are strings.
// l is the length of the generated ID (optional, default is 21).
func NanoIDHook(l ...int) ent.Hook {
	type IDSetter interface {
		SetID(string)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			is, ok := m.(IDSetter)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation %T", m)
			}
			id, err := gonanoid.New(l...)
			if err != nil {
				return nil, err
			}
			is.SetID(id)
			return next.Mutate(ctx, m)
		})
	}
}
