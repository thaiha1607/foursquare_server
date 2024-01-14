package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type userService struct {
	userRepo interfaces.UserRepository
}

func NewUserService(u interfaces.UserRepository) interfaces.UserService {
	return &userService{
		userRepo: u,
	}
}

func (u *userService) Fetch(ctx context.Context) ([]*ent.User, error) {
	return u.userRepo.Fetch(ctx)
}

func (u *userService) GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userService) Store(ctx context.Context, obj *ent.User) (*ent.User, error) {
	return u.userRepo.Store(ctx, obj)
}

func (u *userService) Update(ctx context.Context, id uuid.UUID, obj *ent.User) (*ent.User, error) {
	return u.userRepo.Update(ctx, id, obj)
}

func (u *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return u.userRepo.Delete(ctx, id)
}
