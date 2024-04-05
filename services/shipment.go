package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type shipmentService struct {
	shipmentRepo interfaces.ShipmentRepository
}

func NewShipmentService(r interfaces.ShipmentRepository) interfaces.ShipmentService {
	return &shipmentService{
		shipmentRepo: r,
	}
}

func (s *shipmentService) Fetch(ctx context.Context) ([]*ent.Shipment, error) {
	return s.shipmentRepo.Fetch(ctx)
}

func (s *shipmentService) GetByID(ctx context.Context, id string) (*ent.Shipment, error) {
	return s.shipmentRepo.GetByID(ctx, id)
}

func (s *shipmentService) Store(ctx context.Context, obj *ent.Shipment) (*ent.Shipment, error) {
	return s.shipmentRepo.Store(ctx, obj)
}

func (s *shipmentService) Update(ctx context.Context, id string, obj *ent.Shipment) (*ent.Shipment, error) {
	return s.shipmentRepo.Update(ctx, id, obj)
}

func (s *shipmentService) Delete(ctx context.Context, id string) error {
	return s.shipmentRepo.Delete(ctx, id)
}
