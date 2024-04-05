package interfaces

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
)

// ProductInfoService describes the service.
type ProductInfoService interface {
	Fetch(ctx context.Context) ([]*ent.ProductInfo, error)
	GetByID(ctx context.Context, id string) (*ent.ProductInfo, error)
	Update(ctx context.Context, id string, obj *ent.ProductInfo) (*ent.ProductInfo, error)
	Store(ctx context.Context, obj *ent.ProductInfo) (*ent.ProductInfo, error)
	Delete(ctx context.Context, id string) error
}

// ProductInfoRepository describes the repository.
type ProductInfoRepository interface {
	Fetch(ctx context.Context) ([]*ent.ProductInfo, error)
	GetByID(ctx context.Context, id string) (*ent.ProductInfo, error)
	Update(ctx context.Context, id string, obj *ent.ProductInfo) (*ent.ProductInfo, error)
	Store(ctx context.Context, obj *ent.ProductInfo) (*ent.ProductInfo, error)
	Delete(ctx context.Context, id string) error
}
