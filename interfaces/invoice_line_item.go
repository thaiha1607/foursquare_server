package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// InvoiceLineItemService describes the service.
type InvoiceLineItemService interface {
	Fetch(ctx context.Context) ([]*ent.InvoiceLineItem, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.InvoiceLineItem, error)
	Store(ctx context.Context, obj *ent.InvoiceLineItem) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// InvoiceLineItemRepository describes the repository.
type InvoiceLineItemRepository interface {
	Fetch(ctx context.Context) ([]*ent.InvoiceLineItem, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.InvoiceLineItem, error)
	Store(ctx context.Context, obj *ent.InvoiceLineItem) error
	Delete(ctx context.Context, id uuid.UUID) error
}
