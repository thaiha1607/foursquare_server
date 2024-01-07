package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entMessageRepository struct {
	Client *ent.Client
}

func (e *entMessageRepository) Fetch(ctx context.Context) ([]*ent.Message, error) {
	messages, err := e.Client.Message.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (e *entMessageRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Message, error) {
	message, err := e.Client.Message.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (e *entMessageRepository) Store(ctx context.Context, obj *ent.Message) (err error) {
	//lint:ignore SA4006 we want to return the result of creating operation
	obj, err = e.Client.Message.
		Create().
		SetConversationID(obj.ConversationID).
		SetSenderID(obj.SenderID).
		SetNillableType(&obj.Type).
		SetContent(obj.Content).
		SetNillableIsRead(&obj.IsRead).
		Save(ctx)
	return
}

func (e *entMessageRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.Message) (err error) {
	//lint:ignore SA4006 we want to return the result of updating operation
	obj, err = e.Client.Message.
		UpdateOneID(id).
		SetNillableContent(&obj.Content).
		SetNillableIsRead(&obj.IsRead).
		Save(ctx)
	return
}

func (e *entMessageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.Message.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntMessageRepository(client *ent.Client) interfaces.MessageRepository {
	return &entMessageRepository{Client: client}
}
