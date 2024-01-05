package interfaces

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
)

// OrderTypeService describes the service.
type OrderTypeService interface {
	Fetch(ctx context.Context) ([]*ent.OrderType, error)
	GetByID(ctx context.Context, id int) (*ent.OrderType, error)
	Update(ctx context.Context, obj *ent.OrderType) error
	Store(context.Context, *ent.OrderType) error
	Delete(ctx context.Context, id int) error
}

// OrderTypeRepository describes the repository.
type OrderTypeRepository interface {
	Fetch(ctx context.Context) ([]*ent.OrderType, error)
	GetByID(ctx context.Context, id int) (*ent.OrderType, error)
	Update(ctx context.Context, obj *ent.OrderType) error
	Store(context.Context, *ent.OrderType) error
	Delete(ctx context.Context, id int) error
}
