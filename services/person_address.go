package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type personAddressService struct {
	personAddressRepo interfaces.PersonAddressRepository
}

func NewPersonAddressService(r interfaces.PersonAddressRepository) interfaces.PersonAddressService {
	return &personAddressService{
		personAddressRepo: r,
	}
}

func (s *personAddressService) Fetch(ctx context.Context) ([]*ent.PersonAddress, error) {
	return s.personAddressRepo.Fetch(ctx)
}

func (s *personAddressService) GetByPersonID(ctx context.Context, person_id uuid.UUID) ([]*ent.PersonAddress, error) {
	return s.personAddressRepo.GetByPersonID(ctx, person_id)
}

func (s *personAddressService) Store(ctx context.Context, obj *ent.PersonAddress) (*ent.PersonAddress, error) {
	return s.personAddressRepo.Store(ctx, obj)
}

func (s *personAddressService) Delete(ctx context.Context, person_id uuid.UUID, address_id uuid.UUID) error {
	return s.personAddressRepo.Delete(ctx, person_id, address_id)
}
