package repositories

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entProductInfoRepository struct {
	Client *ent.Client
}

func (e *entProductInfoRepository) Fetch(ctx context.Context) ([]*ent.ProductInfo, error) {
	product_infos, err := e.Client.ProductInfo.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return product_infos, nil
}

func (e *entProductInfoRepository) GetByID(ctx context.Context, id string) (*ent.ProductInfo, error) {
	product_info, err := e.Client.ProductInfo.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return product_info, nil
}

func (e *entProductInfoRepository) Store(ctx context.Context, obj *ent.ProductInfo) (res *ent.ProductInfo, err error) {
	res, err = e.Client.ProductInfo.
		Create().
		SetName(obj.Name).
		SetNillableDescription(obj.Description).
		SetNillableYear(obj.Year).
		SetNillableProvider(obj.Provider).
		Save(ctx)
	return
}

func (e *entProductInfoRepository) Update(ctx context.Context, id string, obj *ent.ProductInfo) (res *ent.ProductInfo, err error) {
	res, err = e.Client.ProductInfo.
		UpdateOneID(id).
		SetNillableName(&obj.Name).
		SetNillableDescription(obj.Description).
		SetNillableYear(obj.Year).
		SetNillableProvider(obj.Provider).
		Save(ctx)
	return
}

func (e *entProductInfoRepository) Delete(ctx context.Context, id string) error {
	return e.Client.ProductInfo.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntProductInfoRepository(client *ent.Client) interfaces.ProductInfoRepository {
	return &entProductInfoRepository{Client: client}
}
