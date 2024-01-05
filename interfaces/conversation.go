package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
)

// ConversationService describes the service.
type ConversationService interface {
	Fetch(ctx context.Context) ([]*ent.Conversation, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Conversation, error)
	Update(ctx context.Context, obj *ent.Conversation) error
	Store(context.Context, *ent.Conversation) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// ConversationRepository describes the repository.
type ConversationRepository interface {
	Fetch(ctx context.Context) ([]*ent.Conversation, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Conversation, error)
	Update(ctx context.Context, obj *ent.Conversation) error
	Store(context.Context, *ent.Conversation) error
	Delete(ctx context.Context, id uuid.UUID) error
}
