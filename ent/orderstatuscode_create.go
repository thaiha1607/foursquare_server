// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/orderstatuscode"
)

// OrderStatusCodeCreate is the builder for creating a OrderStatusCode entity.
type OrderStatusCodeCreate struct {
	config
	mutation *OrderStatusCodeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oscc *OrderStatusCodeCreate) SetCreatedAt(t time.Time) *OrderStatusCodeCreate {
	oscc.mutation.SetCreatedAt(t)
	return oscc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oscc *OrderStatusCodeCreate) SetNillableCreatedAt(t *time.Time) *OrderStatusCodeCreate {
	if t != nil {
		oscc.SetCreatedAt(*t)
	}
	return oscc
}

// SetUpdatedAt sets the "updated_at" field.
func (oscc *OrderStatusCodeCreate) SetUpdatedAt(t time.Time) *OrderStatusCodeCreate {
	oscc.mutation.SetUpdatedAt(t)
	return oscc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oscc *OrderStatusCodeCreate) SetNillableUpdatedAt(t *time.Time) *OrderStatusCodeCreate {
	if t != nil {
		oscc.SetUpdatedAt(*t)
	}
	return oscc
}

// SetOrderStatus sets the "order_status" field.
func (oscc *OrderStatusCodeCreate) SetOrderStatus(s string) *OrderStatusCodeCreate {
	oscc.mutation.SetOrderStatus(s)
	return oscc
}

// SetID sets the "id" field.
func (oscc *OrderStatusCodeCreate) SetID(i int) *OrderStatusCodeCreate {
	oscc.mutation.SetID(i)
	return oscc
}

// Mutation returns the OrderStatusCodeMutation object of the builder.
func (oscc *OrderStatusCodeCreate) Mutation() *OrderStatusCodeMutation {
	return oscc.mutation
}

// Save creates the OrderStatusCode in the database.
func (oscc *OrderStatusCodeCreate) Save(ctx context.Context) (*OrderStatusCode, error) {
	oscc.defaults()
	return withHooks(ctx, oscc.sqlSave, oscc.mutation, oscc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oscc *OrderStatusCodeCreate) SaveX(ctx context.Context) *OrderStatusCode {
	v, err := oscc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oscc *OrderStatusCodeCreate) Exec(ctx context.Context) error {
	_, err := oscc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oscc *OrderStatusCodeCreate) ExecX(ctx context.Context) {
	if err := oscc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oscc *OrderStatusCodeCreate) defaults() {
	if _, ok := oscc.mutation.CreatedAt(); !ok {
		v := orderstatuscode.DefaultCreatedAt()
		oscc.mutation.SetCreatedAt(v)
	}
	if _, ok := oscc.mutation.UpdatedAt(); !ok {
		v := orderstatuscode.DefaultUpdatedAt()
		oscc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oscc *OrderStatusCodeCreate) check() error {
	if _, ok := oscc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderStatusCode.created_at"`)}
	}
	if _, ok := oscc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderStatusCode.updated_at"`)}
	}
	if _, ok := oscc.mutation.OrderStatus(); !ok {
		return &ValidationError{Name: "order_status", err: errors.New(`ent: missing required field "OrderStatusCode.order_status"`)}
	}
	if v, ok := oscc.mutation.OrderStatus(); ok {
		if err := orderstatuscode.OrderStatusValidator(v); err != nil {
			return &ValidationError{Name: "order_status", err: fmt.Errorf(`ent: validator failed for field "OrderStatusCode.order_status": %w`, err)}
		}
	}
	if v, ok := oscc.mutation.ID(); ok {
		if err := orderstatuscode.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "OrderStatusCode.id": %w`, err)}
		}
	}
	return nil
}

func (oscc *OrderStatusCodeCreate) sqlSave(ctx context.Context) (*OrderStatusCode, error) {
	if err := oscc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oscc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oscc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	oscc.mutation.id = &_node.ID
	oscc.mutation.done = true
	return _node, nil
}

func (oscc *OrderStatusCodeCreate) createSpec() (*OrderStatusCode, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderStatusCode{config: oscc.config}
		_spec = sqlgraph.NewCreateSpec(orderstatuscode.Table, sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt))
	)
	if id, ok := oscc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oscc.mutation.CreatedAt(); ok {
		_spec.SetField(orderstatuscode.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := oscc.mutation.UpdatedAt(); ok {
		_spec.SetField(orderstatuscode.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := oscc.mutation.OrderStatus(); ok {
		_spec.SetField(orderstatuscode.FieldOrderStatus, field.TypeString, value)
		_node.OrderStatus = value
	}
	return _node, _spec
}

// OrderStatusCodeCreateBulk is the builder for creating many OrderStatusCode entities in bulk.
type OrderStatusCodeCreateBulk struct {
	config
	err      error
	builders []*OrderStatusCodeCreate
}

// Save creates the OrderStatusCode entities in the database.
func (osccb *OrderStatusCodeCreateBulk) Save(ctx context.Context) ([]*OrderStatusCode, error) {
	if osccb.err != nil {
		return nil, osccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(osccb.builders))
	nodes := make([]*OrderStatusCode, len(osccb.builders))
	mutators := make([]Mutator, len(osccb.builders))
	for i := range osccb.builders {
		func(i int, root context.Context) {
			builder := osccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderStatusCodeMutation)
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
					_, err = mutators[i+1].Mutate(root, osccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, osccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, osccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (osccb *OrderStatusCodeCreateBulk) SaveX(ctx context.Context) []*OrderStatusCode {
	v, err := osccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (osccb *OrderStatusCodeCreateBulk) Exec(ctx context.Context) error {
	_, err := osccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osccb *OrderStatusCodeCreateBulk) ExecX(ctx context.Context) {
	if err := osccb.Exec(ctx); err != nil {
		panic(err)
	}
}
