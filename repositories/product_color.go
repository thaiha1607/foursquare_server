package repositories

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entProductColorRepository struct {
	Client *ent.Client
}

func (e *entProductColorRepository) Fetch(ctx context.Context) ([]*ent.ProductColor, error) {
	product_colors, err := e.Client.ProductColor.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return product_colors, nil
}

func (e *entProductColorRepository) GetByID(ctx context.Context, id string) (*ent.ProductColor, error) {
	product_color, err := e.Client.ProductColor.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return product_color, nil
}

func (e *entProductColorRepository) Store(ctx context.Context, obj *ent.ProductColor) (res *ent.ProductColor, err error) {
	res, err = e.Client.ProductColor.
		Create().
		SetName(obj.Name).
		SetColorCode(obj.ColorCode).
		Save(ctx)
	return
}

func (e *entProductColorRepository) Update(ctx context.Context, id string, obj *ent.ProductColor) (res *ent.ProductColor, err error) {
	res, err = e.Client.ProductColor.
		UpdateOneID(id).
		SetNillableName(&obj.Name).
		SetNillableColorCode(&obj.ColorCode).
		Save(ctx)
	return
}

func (e *entProductColorRepository) Delete(ctx context.Context, id string) error {
	return e.Client.ProductColor.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntProductColorRepository(client *ent.Client) interfaces.ProductColorRepository {
	return &entProductColorRepository{Client: client}
}
