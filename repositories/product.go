package repositories

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entProductRepository struct {
	Client *ent.Client
}

func (e *entProductRepository) Fetch(ctx context.Context) ([]*ent.Product, error) {
	products, err := e.Client.Product.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (e *entProductRepository) GetByID(ctx context.Context, id string) (*ent.Product, error) {
	product, err := e.Client.Product.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (e *entProductRepository) Store(ctx context.Context, obj *ent.Product) (res *ent.Product, err error) {
	res, err = e.Client.Product.
		Create().
		SetName(obj.Name).
		SetNillableDescription(obj.Description).
		SetNillableYear(obj.Year).
		SetPrice(obj.Price).
		SetQty(obj.Qty).
		SetNillableProvider(obj.Provider).
		Save(ctx)
	return
}

func (e *entProductRepository) Update(ctx context.Context, id string, obj *ent.Product) (res *ent.Product, err error) {
	res, err = e.Client.Product.
		UpdateOneID(id).
		SetNillableName(&obj.Name).
		SetNillableDescription(obj.Description).
		SetNillableYear(obj.Year).
		SetNillablePrice(&obj.Price).
		SetNillableQty(&obj.Qty).
		SetNillableProvider(obj.Provider).
		Save(ctx)
	return
}

func (e *entProductRepository) Delete(ctx context.Context, id string) error {
	return e.Client.Product.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntProductRepository(client *ent.Client) interfaces.ProductRepository {
	return &entProductRepository{Client: client}
}
