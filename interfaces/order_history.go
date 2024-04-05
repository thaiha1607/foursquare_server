package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// OrderHistoryService describes the service.
type OrderHistoryService interface {
	GetByOrderID(ctx context.Context, order_id uuid.UUID) ([]*ent.OrderHistory, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderHistory, error)
	Store(ctx context.Context, obj *ent.OrderHistory) (*ent.OrderHistory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// OrderHistoryRepository describes the repository.
type OrderHistoryRepository interface {
	GetByOrderID(ctx context.Context, order_id uuid.UUID) ([]*ent.OrderHistory, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.OrderHistory, error)
	Store(ctx context.Context, obj *ent.OrderHistory) (*ent.OrderHistory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
