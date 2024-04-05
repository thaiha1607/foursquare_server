package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type shipmentStatusCodeService struct {
	shipmentStatusCodeRepo interfaces.ShipmentStatusCodeRepository
}

func NewShipmentStatusCodeService(r interfaces.ShipmentStatusCodeRepository) interfaces.ShipmentStatusCodeService {
	return &shipmentStatusCodeService{
		shipmentStatusCodeRepo: r,
	}
}

func (s *shipmentStatusCodeService) Fetch(ctx context.Context) ([]*ent.ShipmentStatusCode, error) {
	return s.shipmentStatusCodeRepo.Fetch(ctx)
}

func (s *shipmentStatusCodeService) GetByID(ctx context.Context, id int) (*ent.ShipmentStatusCode, error) {
	return s.shipmentStatusCodeRepo.GetByID(ctx, id)
}

func (s *shipmentStatusCodeService) Store(ctx context.Context, obj *ent.ShipmentStatusCode) (*ent.ShipmentStatusCode, error) {
	return s.shipmentStatusCodeRepo.Store(ctx, obj)
}

func (s *shipmentStatusCodeService) Update(ctx context.Context, id int, obj *ent.ShipmentStatusCode) (*ent.ShipmentStatusCode, error) {
	return s.shipmentStatusCodeRepo.Update(ctx, id, obj)
}

func (s *shipmentStatusCodeService) Delete(ctx context.Context, id int) error {
	return s.shipmentStatusCodeRepo.Delete(ctx, id)
}
