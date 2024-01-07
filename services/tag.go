package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type tagService struct {
	tagRepo interfaces.TagRepository
}

func NewTagService(t interfaces.TagRepository) interfaces.TagService {
	return &tagService{
		tagRepo: t,
	}
}

func (t *tagService) Fetch(ctx context.Context) ([]*ent.Tag, error) {
	return t.tagRepo.Fetch(ctx)
}

func (t *tagService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Tag, error) {
	return t.tagRepo.GetByID(ctx, id)
}

func (t *tagService) Store(ctx context.Context, obj *ent.Tag) error {
	return t.tagRepo.Store(ctx, obj)
}

func (t *tagService) Update(ctx context.Context, id uuid.UUID, obj *ent.Tag) error {
	return t.tagRepo.Update(ctx, id, obj)
}

func (t *tagService) Delete(ctx context.Context, id uuid.UUID) error {
	return t.tagRepo.Delete(ctx, id)
}
