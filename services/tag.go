package services

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type tagService struct {
	tagRepo interfaces.TagRepository
}

func NewTagService(r interfaces.TagRepository) interfaces.TagService {
	return &tagService{
		tagRepo: r,
	}
}

func (s *tagService) Fetch(ctx context.Context) ([]*ent.Tag, error) {
	return s.tagRepo.Fetch(ctx)
}

func (s *tagService) GetByID(ctx context.Context, id string) (*ent.Tag, error) {
	return s.tagRepo.GetByID(ctx, id)
}

func (s *tagService) Store(ctx context.Context, obj *ent.Tag) (*ent.Tag, error) {
	return s.tagRepo.Store(ctx, obj)
}

func (s *tagService) Update(ctx context.Context, id string, obj *ent.Tag) (*ent.Tag, error) {
	return s.tagRepo.Update(ctx, id, obj)
}

func (s *tagService) Delete(ctx context.Context, id string) error {
	return s.tagRepo.Delete(ctx, id)
}
