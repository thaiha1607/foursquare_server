package repositories

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/ent/producttag"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entProductTagRepository struct {
	Client *ent.Client
}

func (e *entProductTagRepository) Fetch(ctx context.Context) ([]*ent.ProductTag, error) {
	product_tags, err := e.Client.ProductTag.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return product_tags, nil
}

func (e *entProductTagRepository) Store(ctx context.Context, obj *ent.ProductTag) (res *ent.ProductTag, err error) {
	res, err = e.Client.ProductTag.
		Create().
		SetProductID(obj.ProductID).
		SetTagID(obj.TagID).
		Save(ctx)
	return
}

func (e *entProductTagRepository) Delete(ctx context.Context, product_id string, tag_id string) (err error) {
	_, err = e.Client.ProductTag.
		Delete().
		Where(
			producttag.And(
				producttag.ProductID(product_id),
				producttag.TagID(tag_id),
			),
		).
		Exec(ctx)
	return
}

func NewEntProductTagRepository(client *ent.Client) interfaces.ProductTagRepository {
	return &entProductTagRepository{Client: client}
}
