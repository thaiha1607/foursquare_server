package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type addressService struct {
	addressRepo interfaces.AddressRepository
}

func NewAddressService(r interfaces.AddressRepository) interfaces.AddressService {
	return &addressService{
		addressRepo: r,
	}
}

func (s *addressService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Address, error) {
	return s.addressRepo.GetByID(ctx, id)
}

func (s *addressService) Store(ctx context.Context, obj *ent.Address) (*ent.Address, error) {
	return s.addressRepo.Store(ctx, obj)
}

func (s *addressService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.addressRepo.Delete(ctx, id)
}
