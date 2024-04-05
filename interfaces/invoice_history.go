package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// InvoiceHistoryService describes the service.
type InvoiceHistoryService interface {
	GetByInvoiceID(ctx context.Context, invoice_id uuid.UUID) ([]*ent.InvoiceHistory, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.InvoiceHistory, error)
	Store(ctx context.Context, obj *ent.InvoiceHistory) (*ent.InvoiceHistory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// InvoiceHistoryRepository describes the repository.
type InvoiceHistoryRepository interface {
	GetByInvoiceID(ctx context.Context, invoice_id uuid.UUID) ([]*ent.InvoiceHistory, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.InvoiceHistory, error)
	Store(ctx context.Context, obj *ent.InvoiceHistory) (*ent.InvoiceHistory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
