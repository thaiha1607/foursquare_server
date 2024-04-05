package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// WorkUnitInfoService describes the service.
type WorkUnitInfoService interface {
	Fetch(ctx context.Context) ([]*ent.WorkUnitInfo, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.WorkUnitInfo, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.WorkUnitInfo) (*ent.WorkUnitInfo, error)
	Store(ctx context.Context, obj *ent.WorkUnitInfo) (*ent.WorkUnitInfo, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// WorkUnitInfoRepository describes the repository.
type WorkUnitInfoRepository interface {
	Fetch(ctx context.Context) ([]*ent.WorkUnitInfo, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.WorkUnitInfo, error)
	Update(ctx context.Context, id uuid.UUID, obj *ent.WorkUnitInfo) (*ent.WorkUnitInfo, error)
	Store(ctx context.Context, obj *ent.WorkUnitInfo) (*ent.WorkUnitInfo, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
