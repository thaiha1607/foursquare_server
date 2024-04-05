package interfaces

import (
	"context"

	"github.com/thaiha1607/foursquare_server/ent"
)

// InvoiceStatusCodeService describes the service.
type InvoiceStatusCodeService interface {
	Fetch(ctx context.Context) ([]*ent.InvoiceStatusCode, error)
	GetByID(ctx context.Context, id int) (*ent.InvoiceStatusCode, error)
	Update(ctx context.Context, id int, obj *ent.InvoiceStatusCode) (*ent.InvoiceStatusCode, error)
	Store(ctx context.Context, obj *ent.InvoiceStatusCode) (*ent.InvoiceStatusCode, error)
	Delete(ctx context.Context, id int) error
}

// InvoiceStatusCodeRepository describes the repository.
type InvoiceStatusCodeRepository interface {
	Fetch(ctx context.Context) ([]*ent.InvoiceStatusCode, error)
	GetByID(ctx context.Context, id int) (*ent.InvoiceStatusCode, error)
	Update(ctx context.Context, id int, obj *ent.InvoiceStatusCode) (*ent.InvoiceStatusCode, error)
	Store(ctx context.Context, obj *ent.InvoiceStatusCode) (*ent.InvoiceStatusCode, error)
	Delete(ctx context.Context, id int) error
}
