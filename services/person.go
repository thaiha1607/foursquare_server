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

func NewPersonService(u interfaces.PersonRepository) interfaces.PersonService {
	return &personService{
		personRepo: u,
	}
}

func (u *personService) Fetch(ctx context.Context) ([]*ent.Person, error) {
	return u.personRepo.Fetch(ctx)
}

func (u *personService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Person, error) {
	return u.personRepo.GetByID(ctx, id)
}

func (u *personService) Store(ctx context.Context, obj *ent.Person) (*ent.Person, error) {
	return u.personRepo.Store(ctx, obj)
}

func (u *personService) Update(ctx context.Context, id uuid.UUID, obj *ent.Person) (*ent.Person, error) {
	return u.personRepo.Update(ctx, id, obj)
}

func (u *personService) Delete(ctx context.Context, id uuid.UUID) error {
	return u.personRepo.Delete(ctx, id)
}
