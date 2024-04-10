package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entWorkUnitInfoRepository struct {
	Client *ent.Client
}

func (e *entWorkUnitInfoRepository) Fetch(ctx context.Context) ([]*ent.WorkUnitInfo, error) {
	work_unit_info_list, err := e.Client.WorkUnitInfo.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return work_unit_info_list, nil
}

func (e *entWorkUnitInfoRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.WorkUnitInfo, error) {
	work_unit_info, err := e.Client.WorkUnitInfo.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return work_unit_info, nil
}

func (e *entWorkUnitInfoRepository) Store(ctx context.Context, obj *ent.WorkUnitInfo) (res *ent.WorkUnitInfo, err error) {
	res, err = e.Client.WorkUnitInfo.
		Create().
		SetName(obj.Name).
		SetNillableAddressID(obj.AddressID).
		SetNillableType(obj.Type).
		SetNillableImageURL(obj.ImageURL).
		Save(ctx)
	return
}

func (e *entWorkUnitInfoRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.WorkUnitInfo) (res *ent.WorkUnitInfo, err error) {
	res, err = e.Client.WorkUnitInfo.
		UpdateOneID(id).
		SetNillableAddressID(obj.AddressID).
		SetNillableType(obj.Type).
		SetNillableImageURL(obj.ImageURL).
		Save(ctx)
	return
}

func (e *entWorkUnitInfoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.WorkUnitInfo.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntWorkUnitInfoRepository(client *ent.Client) interfaces.WorkUnitInfoRepository {
	return &entWorkUnitInfoRepository{Client: client}
}
