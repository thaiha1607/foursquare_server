package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type shipmentHistoryService struct {
	shipmentHistoryRepo interfaces.ShipmentHistoryRepository
}

func NewShipmentHistoryService(r interfaces.ShipmentHistoryRepository) interfaces.ShipmentHistoryService {
	return &shipmentHistoryService{
		shipmentHistoryRepo: r,
	}
}

func (s *shipmentHistoryService) GetByID(ctx context.Context, id uuid.UUID) (*ent.ShipmentHistory, error) {
	return s.shipmentHistoryRepo.GetByID(ctx, id)
}

func (s *shipmentHistoryService) Store(ctx context.Context, obj *ent.ShipmentHistory) (*ent.ShipmentHistory, error) {
	return s.shipmentHistoryRepo.Store(ctx, obj)
}

func (s *shipmentHistoryService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.shipmentHistoryRepo.Delete(ctx, id)
}
