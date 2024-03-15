package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent"
	_ "github.com/thaiha1607/foursquare_server/ent/runtime"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type entPersonRepository struct {
	Client *ent.Client
}

func (e *entPersonRepository) Fetch(ctx context.Context) ([]*ent.Person, error) {
	persons, err := e.Client.Person.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func (e *entPersonRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Person, error) {
	person, err := e.Client.Person.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (e *entPersonRepository) Store(ctx context.Context, obj *ent.Person) (res *ent.Person, err error) {
	res, err = e.Client.Person.
		Create().
		SetAvatarURL(obj.AvatarURL).
		SetNillableEmail(obj.Email).
		SetName(obj.Name).
		SetPhone(obj.Phone).
		Save(ctx)
	return

}

func (e *entPersonRepository) Update(ctx context.Context, id uuid.UUID, obj *ent.Person) (res *ent.Person, err error) {
	res, err = e.Client.Person.
		UpdateOneID(id).
		SetAvatarURL(obj.AvatarURL).
		SetNillableEmail(obj.Email).
		SetNillableName(&obj.Name).
		SetNillablePhone(&obj.Phone).
		Save(ctx)
	return
}

func (e *entPersonRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.Client.Person.
		DeleteOneID(id).
		Exec(ctx)
}

func NewEntPersonRepository(client *ent.Client) interfaces.PersonRepository {
	return &entPersonRepository{Client: client}
}
