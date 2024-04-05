package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type warehouseAssignmentService struct {
	warehouseAssignmentRepo interfaces.WarehouseAssignmentRepository
}

func NewWarehouseAssignmentService(r interfaces.WarehouseAssignmentRepository) interfaces.WarehouseAssignmentService {
	return &warehouseAssignmentService{
		warehouseAssignmentRepo: r,
	}
}

func (s *warehouseAssignmentService) Fetch(ctx context.Context) ([]*ent.WarehouseAssignment, error) {
	return s.warehouseAssignmentRepo.Fetch(ctx)
}

func (s *warehouseAssignmentService) GetByID(ctx context.Context, id uuid.UUID) (*ent.WarehouseAssignment, error) {
	return s.warehouseAssignmentRepo.GetByID(ctx, id)
}

func (s *warehouseAssignmentService) Store(ctx context.Context, obj *ent.WarehouseAssignment) (*ent.WarehouseAssignment, error) {
	return s.warehouseAssignmentRepo.Store(ctx, obj)
}

func (s *warehouseAssignmentService) Update(ctx context.Context, id uuid.UUID, obj *ent.WarehouseAssignment) (*ent.WarehouseAssignment, error) {
	return s.warehouseAssignmentRepo.Update(ctx, id, obj)
}

func (s *warehouseAssignmentService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.warehouseAssignmentRepo.Delete(ctx, id)
}
