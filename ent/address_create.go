// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/address"
	"github.com/thaiha1607/foursquare_server/ent/person"
)

// AddressCreate is the builder for creating a Address entity.
type AddressCreate struct {
	config
	mutation *AddressMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AddressCreate) SetCreatedAt(t time.Time) *AddressCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AddressCreate) SetNillableCreatedAt(t *time.Time) *AddressCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetLine1 sets the "line1" field.
func (ac *AddressCreate) SetLine1(s string) *AddressCreate {
	ac.mutation.SetLine1(s)
	return ac
}

// SetLine2 sets the "line2" field.
func (ac *AddressCreate) SetLine2(s string) *AddressCreate {
	ac.mutation.SetLine2(s)
	return ac
}

// SetNillableLine2 sets the "line2" field if the given value is not nil.
func (ac *AddressCreate) SetNillableLine2(s *string) *AddressCreate {
	if s != nil {
		ac.SetLine2(*s)
	}
	return ac
}

// SetCity sets the "city" field.
func (ac *AddressCreate) SetCity(s string) *AddressCreate {
	ac.mutation.SetCity(s)
	return ac
}

// SetStateOrProvince sets the "state_or_province" field.
func (ac *AddressCreate) SetStateOrProvince(s string) *AddressCreate {
	ac.mutation.SetStateOrProvince(s)
	return ac
}

// SetZipOrPostcode sets the "zip_or_postcode" field.
func (ac *AddressCreate) SetZipOrPostcode(s string) *AddressCreate {
	ac.mutation.SetZipOrPostcode(s)
	return ac
}

// SetCountry sets the "country" field.
func (ac *AddressCreate) SetCountry(s string) *AddressCreate {
	ac.mutation.SetCountry(s)
	return ac
}

// SetOtherAddressDetails sets the "other_address_details" field.
func (ac *AddressCreate) SetOtherAddressDetails(s string) *AddressCreate {
	ac.mutation.SetOtherAddressDetails(s)
	return ac
}

// SetNillableOtherAddressDetails sets the "other_address_details" field if the given value is not nil.
func (ac *AddressCreate) SetNillableOtherAddressDetails(s *string) *AddressCreate {
	if s != nil {
		ac.SetOtherAddressDetails(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AddressCreate) SetID(u uuid.UUID) *AddressCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AddressCreate) SetNillableID(u *uuid.UUID) *AddressCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// AddPersonIDs adds the "persons" edge to the Person entity by IDs.
func (ac *AddressCreate) AddPersonIDs(ids ...uuid.UUID) *AddressCreate {
	ac.mutation.AddPersonIDs(ids...)
	return ac
}

// AddPersons adds the "persons" edges to the Person entity.
func (ac *AddressCreate) AddPersons(p ...*Person) *AddressCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPersonIDs(ids...)
}

// Mutation returns the AddressMutation object of the builder.
func (ac *AddressCreate) Mutation() *AddressMutation {
	return ac.mutation
}

// Save creates the Address in the database.
func (ac *AddressCreate) Save(ctx context.Context) (*Address, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AddressCreate) SaveX(ctx context.Context) *Address {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AddressCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AddressCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AddressCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := address.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := address.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AddressCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Address.created_at"`)}
	}
	if _, ok := ac.mutation.Line1(); !ok {
		return &ValidationError{Name: "line1", err: errors.New(`ent: missing required field "Address.line1"`)}
	}
	if v, ok := ac.mutation.Line1(); ok {
		if err := address.Line1Validator(v); err != nil {
			return &ValidationError{Name: "line1", err: fmt.Errorf(`ent: validator failed for field "Address.line1": %w`, err)}
		}
	}
	if _, ok := ac.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "Address.city"`)}
	}
	if v, ok := ac.mutation.City(); ok {
		if err := address.CityValidator(v); err != nil {
			return &ValidationError{Name: "city", err: fmt.Errorf(`ent: validator failed for field "Address.city": %w`, err)}
		}
	}
	if _, ok := ac.mutation.StateOrProvince(); !ok {
		return &ValidationError{Name: "state_or_province", err: errors.New(`ent: missing required field "Address.state_or_province"`)}
	}
	if v, ok := ac.mutation.StateOrProvince(); ok {
		if err := address.StateOrProvinceValidator(v); err != nil {
			return &ValidationError{Name: "state_or_province", err: fmt.Errorf(`ent: validator failed for field "Address.state_or_province": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ZipOrPostcode(); !ok {
		return &ValidationError{Name: "zip_or_postcode", err: errors.New(`ent: missing required field "Address.zip_or_postcode"`)}
	}
	if v, ok := ac.mutation.ZipOrPostcode(); ok {
		if err := address.ZipOrPostcodeValidator(v); err != nil {
			return &ValidationError{Name: "zip_or_postcode", err: fmt.Errorf(`ent: validator failed for field "Address.zip_or_postcode": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`ent: missing required field "Address.country"`)}
	}
	if v, ok := ac.mutation.Country(); ok {
		if err := address.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`ent: validator failed for field "Address.country": %w`, err)}
		}
	}
	return nil
}

func (ac *AddressCreate) sqlSave(ctx context.Context) (*Address, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AddressCreate) createSpec() (*Address, *sqlgraph.CreateSpec) {
	var (
		_node = &Address{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(address.Table, sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(address.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := ac.mutation.Line1(); ok {
		_spec.SetField(address.FieldLine1, field.TypeString, value)
		_node.Line1 = value
	}
	if value, ok := ac.mutation.Line2(); ok {
		_spec.SetField(address.FieldLine2, field.TypeString, value)
		_node.Line2 = &value
	}
	if value, ok := ac.mutation.City(); ok {
		_spec.SetField(address.FieldCity, field.TypeString, value)
		_node.City = value
	}
	if value, ok := ac.mutation.StateOrProvince(); ok {
		_spec.SetField(address.FieldStateOrProvince, field.TypeString, value)
		_node.StateOrProvince = value
	}
	if value, ok := ac.mutation.ZipOrPostcode(); ok {
		_spec.SetField(address.FieldZipOrPostcode, field.TypeString, value)
		_node.ZipOrPostcode = value
	}
	if value, ok := ac.mutation.Country(); ok {
		_spec.SetField(address.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if value, ok := ac.mutation.OtherAddressDetails(); ok {
		_spec.SetField(address.FieldOtherAddressDetails, field.TypeString, value)
		_node.OtherAddressDetails = &value
	}
	if nodes := ac.mutation.PersonsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   address.PersonsTable,
			Columns: address.PersonsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &PersonAddressCreate{config: ac.config, mutation: newPersonAddressMutation(ac.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AddressCreateBulk is the builder for creating many Address entities in bulk.
type AddressCreateBulk struct {
	config
	err      error
	builders []*AddressCreate
}

// Save creates the Address entities in the database.
func (acb *AddressCreateBulk) Save(ctx context.Context) ([]*Address, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Address, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AddressMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AddressCreateBulk) SaveX(ctx context.Context) []*Address {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AddressCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AddressCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}