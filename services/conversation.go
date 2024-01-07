package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type conversationService struct {
	conversationRepo interfaces.ConversationRepository
}

func NewConversationService(c interfaces.ConversationRepository) interfaces.ConversationService {
	return &conversationService{
		conversationRepo: c,
	}
}

func (c *conversationService) Fetch(ctx context.Context) ([]*ent.Conversation, error) {
	return c.conversationRepo.Fetch(ctx)
}

func (c *conversationService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Conversation, error) {
	return c.conversationRepo.GetByID(ctx, id)
}

func (c *conversationService) Store(ctx context.Context, obj *ent.Conversation) error {
	return c.conversationRepo.Store(ctx, obj)
}

func (c *conversationService) Update(ctx context.Context, id uuid.UUID, obj *ent.Conversation) error {
	return c.conversationRepo.Update(ctx, id, obj)
}

func (c *conversationService) Delete(ctx context.Context, id uuid.UUID) error {
	return c.conversationRepo.Delete(ctx, id)
}
