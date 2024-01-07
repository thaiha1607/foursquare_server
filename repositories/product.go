package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
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

func (e *entProductRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Product, error) {
	product, err := e.Client.Product.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (e *entProductRepository) Store(ctx context.Context, obj *ent.Product) (err error) {
	//lint:ignore SA4006 we want to return the result of creating operation
	obj, err = e.Client.Product.
		Create().
		SetName(obj.Name).
		SetNillableDescription(obj.Description).
		SetNillableYear(obj.Year).
		SetPrice(obj.Price).
		SetQty(obj.Qty).
		SetNillableUnitOfMeasurement(obj.UnitOfMeasurement).
		SetNillableType(obj.Type).
		SetNillableProvider(obj.Provider).
		Save(ctx)
	return err
}

func (e *entProductRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.Product) error {
	//lint:ignore SA4006 we want to return the result of updating operation
	obj, err := e.Client.Product.
		UpdateOneID(id).
		SetNillableName(&obj.Name).
		SetNillableDescription(obj.Description).
		SetNillableYear(obj.Year).
		SetNillablePrice(&obj.Price).
		SetNillableQty(&obj.Qty).
		SetNillableUnitOfMeasurement(obj.UnitOfMeasurement).
		SetNillableType(obj.Type).
		SetNillableProvider(obj.Provider).
		Save(ctx)
	return err
}

func (e *entProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.Product.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntProductRepository(client *ent.Client) interfaces.ProductRepository {
	return &entProductRepository{Client: client}
}
