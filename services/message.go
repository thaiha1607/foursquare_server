package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type messageService struct {
	messageRepo interfaces.MessageRepository
}

func NewMessageService(r interfaces.MessageRepository) interfaces.MessageService {
	return &messageService{
		messageRepo: r,
	}
}

func (s *messageService) Fetch(ctx context.Context) ([]*ent.Message, error) {
	return s.messageRepo.Fetch(ctx)
}

func (s *messageService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Message, error) {
	return s.messageRepo.GetByID(ctx, id)
}

func (s *messageService) Store(ctx context.Context, obj *ent.Message) (*ent.Message, error) {
	return s.messageRepo.Store(ctx, obj)
}

func (s *messageService) Update(ctx context.Context, id uuid.UUID, obj *ent.Message) (*ent.Message, error) {
	return s.messageRepo.Update(ctx, id, obj)
}

func (s *messageService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.messageRepo.Delete(ctx, id)
}
