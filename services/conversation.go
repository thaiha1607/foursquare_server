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

func NewConversationService(r interfaces.ConversationRepository) interfaces.ConversationService {
	return &conversationService{
		conversationRepo: r,
	}
}

func (s *conversationService) Fetch(ctx context.Context) ([]*ent.Conversation, error) {
	return s.conversationRepo.Fetch(ctx)
}

func (s *conversationService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Conversation, error) {
	return s.conversationRepo.GetByID(ctx, id)
}

func (s *conversationService) Store(ctx context.Context, obj *ent.Conversation) (*ent.Conversation, error) {
	return s.conversationRepo.Store(ctx, obj)
}

func (s *conversationService) Update(ctx context.Context, id uuid.UUID, obj *ent.Conversation) (*ent.Conversation, error) {
	return s.conversationRepo.Update(ctx, id, obj)
}

func (s *conversationService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.conversationRepo.Delete(ctx, id)
}
