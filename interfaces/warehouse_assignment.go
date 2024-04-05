package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// WarehouseAssignmentService describes the service.
type WarehouseAssignmentService interface {
	Fetch(ctx context.Context) ([]*ent.WarehouseAssignment, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.WarehouseAssignment, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.WarehouseAssignment) (*ent.WarehouseAssignment, error)
	Store(ctx context.Context, obj *ent.WarehouseAssignment) (*ent.WarehouseAssignment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// WarehouseAssignmentRepository describes the repository.
type WarehouseAssignmentRepository interface {
	Fetch(ctx context.Context) ([]*ent.WarehouseAssignment, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.WarehouseAssignment, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.WarehouseAssignment) (*ent.WarehouseAssignment, error)
	Store(ctx context.Context, obj *ent.WarehouseAssignment) (*ent.WarehouseAssignment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
