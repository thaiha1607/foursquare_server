package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type shipmentItemService struct {
	shipmentItemRepo interfaces.ShipmentItemRepository
}

func NewShipmentItemService(r interfaces.ShipmentItemRepository) interfaces.ShipmentItemService {
	return &shipmentItemService{
		shipmentItemRepo: r,
	}
}

func (s *shipmentItemService) GetByID(ctx context.Context, id uuid.UUID) (*ent.ShipmentItem, error) {
	return s.shipmentItemRepo.GetByID(ctx, id)
}

func (s *shipmentItemService) Store(ctx context.Context, obj *ent.ShipmentItem) (*ent.ShipmentItem, error) {
	return s.shipmentItemRepo.Store(ctx, obj)
}

func (s *shipmentItemService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.shipmentItemRepo.Delete(ctx, id)
}
