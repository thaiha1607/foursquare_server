package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// InvoiceService describes the service.
type InvoiceService interface {
	Fetch(ctx context.Context) ([]*ent.Invoice, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Invoice, error)
	Update(ctx context.Context, obj *ent.Invoice) error
	Store(context.Context, *ent.Invoice) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// InvoiceRepository describes the repository.
type InvoiceRepository interface {
	Fetch(ctx context.Context) ([]*ent.Invoice, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Invoice, error)
	Update(ctx context.Context, obj *ent.Invoice) error
	Store(context.Context, *ent.Invoice) error
	Delete(ctx context.Context, id uuid.UUID) error
}
