package repositories

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entShipmentRepository struct {
	Client *ent.Client
}

func (e *entShipmentRepository) Fetch(ctx context.Context) ([]*ent.Shipment, error) {
	shipments, err := e.Client.Shipment.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return shipments, nil
}

func (e *entShipmentRepository) GetByID(ctx context.Context, id string) (*ent.Shipment, error) {
	shipment, err := e.Client.Shipment.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return shipment, nil
}

func (e *entShipmentRepository) Store(ctx context.Context, obj *ent.Shipment) (res *ent.Shipment, err error) {
	res, err = e.Client.Shipment.
		Create().
		SetOrderID(obj.OrderID).
		SetInvoiceID(obj.InvoiceID).
		SetStaffID(obj.StaffID).
		SetShipmentDate(*obj.ShipmentDate).
		SetNillableNote(obj.Note).
		SetNillableStatusCode(obj.StatusCode).
		Save(ctx)
	return
}

func (e *entShipmentRepository) Update(ctx context.Context, id string, obj *ent.Shipment) (res *ent.Shipment, err error) {
	res, err = e.Client.Shipment.
		UpdateOneID(id).
		SetNillableShipmentDate(obj.ShipmentDate).
		SetNillableNote(obj.Note).
		SetNillableStatusCode(obj.StatusCode).
		Save(ctx)
	return
}

func (e *entShipmentRepository) Delete(ctx context.Context, id string) error {
	return e.Client.Shipment.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntShipmentRepository(client *ent.Client) interfaces.ShipmentRepository {
	return &entShipmentRepository{Client: client}
}
