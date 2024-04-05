package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type workUnitInfoService struct {
	workUnitInfoRepo interfaces.WorkUnitInfoRepository
}

func NewWorkUnitInfoService(r interfaces.WorkUnitInfoRepository) interfaces.WorkUnitInfoService {
	return &workUnitInfoService{
		workUnitInfoRepo: r,
	}
}

func (s *workUnitInfoService) Fetch(ctx context.Context) ([]*ent.WorkUnitInfo, error) {
	return s.workUnitInfoRepo.Fetch(ctx)
}

func (s *workUnitInfoService) GetByID(ctx context.Context, id uuid.UUID) (*ent.WorkUnitInfo, error) {
	return s.workUnitInfoRepo.GetByID(ctx, id)
}

func (s *workUnitInfoService) Store(ctx context.Context, obj *ent.WorkUnitInfo) (*ent.WorkUnitInfo, error) {
	return s.workUnitInfoRepo.Store(ctx, obj)
}

func (s *workUnitInfoService) Update(ctx context.Context, id uuid.UUID, obj *ent.WorkUnitInfo) (*ent.WorkUnitInfo, error) {
	return s.workUnitInfoRepo.Update(ctx, id, obj)
}

func (s *workUnitInfoService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.workUnitInfoRepo.Delete(ctx, id)
}
