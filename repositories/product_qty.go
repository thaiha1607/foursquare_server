package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entProductQtyRepository struct {
	Client *ent.Client
}

func (e *entProductQtyRepository) Fetch(ctx context.Context) ([]*ent.ProductQty, error) {
	product_qtys, err := e.Client.ProductQty.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return product_qtys, nil
}

func (e *entProductQtyRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.ProductQty, error) {
	product_qty, err := e.Client.ProductQty.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return product_qty, nil
}

func (e *entProductQtyRepository) Store(ctx context.Context, obj *ent.ProductQty) (res *ent.ProductQty, err error) {
	res, err = e.Client.ProductQty.
		Create().
		SetWorkUnitID(obj.WorkUnitID).
		SetProductID(obj.ProductID).
		SetProductColorID(obj.ProductColorID).
		SetPricePerUnit(*obj.PricePerUnit).
		SetQty(*obj.Qty).
		Save(ctx)
	return
}

func (e *entProductQtyRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.ProductQty) (res *ent.ProductQty, err error) {
	res, err = e.Client.ProductQty.
		UpdateOneID(id).
		SetNillablePricePerUnit(obj.PricePerUnit).
		SetNillableQty(obj.Qty).
		Save(ctx)
	return
}

func (e *entProductQtyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.ProductQty.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntProductQtyRepository(client *ent.Client) interfaces.ProductQtyRepository {
	return &entProductQtyRepository{Client: client}
}
