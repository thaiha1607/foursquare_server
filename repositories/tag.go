package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entTagRepository struct {
	Client *ent.Client
}

func (e *entTagRepository) Fetch(ctx context.Context) ([]*ent.Tag, error) {
	tags, err := e.Client.Tag.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (e *entTagRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Tag, error) {
	tag, err := e.Client.Tag.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (e *entTagRepository) Store(ctx context.Context, obj *ent.Tag) (err error) {
	//lint:ignore SA4006 we want to return the result of creating operation
	obj, err = e.Client.Tag.
		Create().
		SetTitle(obj.Title).
		Save(ctx)
	return
}

func (e *entTagRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.Tag) (err error) {
	//lint:ignore SA4006 we want to return the result of updating operation
	obj, err = e.Client.Tag.
		UpdateOneID(id).
		SetTitle(obj.Title).
		Save(ctx)
	return
}

func (e *entTagRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.Tag.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntTagRepository(client *ent.Client) interfaces.TagRepository {
	return &entTagRepository{Client: client}
}