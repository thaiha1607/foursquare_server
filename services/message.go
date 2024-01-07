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

func NewMessageService(m interfaces.MessageRepository) interfaces.MessageService {
	return &messageService{
		messageRepo: m,
	}
}

func (m *messageService) Fetch(ctx context.Context) ([]*ent.Message, error) {
	return m.messageRepo.Fetch(ctx)
}

func (m *messageService) GetByID(ctx context.Context, id uuid.UUID) (*ent.Message, error) {
	return m.messageRepo.GetByID(ctx, id)
}

func (m *messageService) Store(ctx context.Context, obj *ent.Message) error {
	return m.messageRepo.Store(ctx, obj)
}

func (m *messageService) Update(ctx context.Context, id uuid.UUID, obj *ent.Message) error {
	return m.messageRepo.Update(ctx, id, obj)
}

func (m *messageService) Delete(ctx context.Context, id uuid.UUID) error {
	return m.messageRepo.Delete(ctx, id)
}
