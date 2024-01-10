package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entProductImageRepository struct {
	Client *ent.Client
}

func (e *entProductImageRepository) Fetch(ctx context.Context) ([]*ent.ProductImage, error) {
	orders, err := e.Client.ProductImage.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (e *entProductImageRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.ProductImage, error) {
	order, err := e.Client.ProductImage.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (e *entProductImageRepository) Store(ctx context.Context, obj *ent.ProductImage) (err error) {
	//lint:ignore SA4006 we want to return the result of creating operation
	obj, err = e.Client.ProductImage.
		Create().
		SetProductID(obj.ProductID).
		SetImageURL(obj.ImageURL).
		Save(ctx)
	return
}

func (e *entProductImageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.ProductImage.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntProductImageRepository(client *ent.Client) interfaces.ProductImageRepository {
	return &entProductImageRepository{Client: client}
}
