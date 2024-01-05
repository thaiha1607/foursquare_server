package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// MessageService describes the service.
type MessageService interface {
	Fetch(ctx context.Context) ([]*ent.Message, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Message, error)
	Update(ctx context.Context, obj *ent.Message) error
	Store(context.Context, *ent.Message) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// MessageRepository describes the repository.
type MessageRepository interface {
	Fetch(ctx context.Context) ([]*ent.Message, string, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Message, error)
	Update(ctx context.Context, obj *ent.Message) error
	Store(context.Context, *ent.Message) error
	Delete(ctx context.Context, id uuid.UUID) error
}
