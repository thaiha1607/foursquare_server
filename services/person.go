package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type personService struct {
	personRepo interfaces.PersonRepository
}

func NewPersonService(r interfaces.PersonRepository) interfaces.PersonService {
	return &personService{
		personRepo: r,
	}
}

func (s *personService) Fetch(ctx context.Context) ([]*ent.Person, error) {
	return s.personRepo.Fetch(ctx)
}

func (s *personService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Person, error) {
	return s.personRepo.GetByID(ctx, id)
}

func (s *personService) Store(ctx context.Context, obj *ent.Person) (*ent.Person, error) {
	return s.personRepo.Store(ctx, obj)
}

func (s *personService) Update(ctx context.Context, id uuid.UUID, obj *ent.Person) (*ent.Person, error) {
	return s.personRepo.Update(ctx, id, obj)
}

func (s *personService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.personRepo.Delete(ctx, id)
}
