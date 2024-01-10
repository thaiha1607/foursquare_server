package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entConversationRepository struct {
	Client *ent.Client
}

func (e *entConversationRepository) Fetch(ctx context.Context) ([]*ent.Conversation, error) {
	conversations, err := e.Client.Conversation.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return conversations, nil
}

func (e *entConversationRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Conversation, error) {
	conversation, err := e.Client.Conversation.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return conversation, nil
}

func (e *entConversationRepository) Store(ctx context.Context, obj *ent.Conversation) (err error) {
	//lint:ignore SA4006 we want to return the result of creating operation
	obj, err = e.Client.Conversation.
		Create().
		SetNillableTitle(obj.Title).
		SetUserOneID(obj.UserOneID).
		SetUserTwoID(obj.UserTwoID).
		Save(ctx)
	return
}

func (e *entConversationRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.Conversation) (err error) {
	//lint:ignore SA4006 we want to return the result of updating operation
	obj, err = e.Client.Conversation.
		UpdateOneID(id).
		SetNillableTitle(obj.Title).
		Save(ctx)
	return
}

func (e *entConversationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.Conversation.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntConversationRepository(client *ent.Client) interfaces.ConversationRepository {
	return &entConversationRepository{Client: client}
}
