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
	"github.com/shopspring/decimal"
	"github.com/thaiha1607/foursquare_server/ent/order"
	"github.com/thaiha1607/foursquare_server/ent/orderlineitem"
	"github.com/thaiha1607/foursquare_server/ent/product"
)

// OrderLineItemCreate is the builder for creating a OrderLineItem entity.
type OrderLineItemCreate struct {
	config
	mutation *OrderLineItemMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (olic *OrderLineItemCreate) SetCreatedAt(t time.Time) *OrderLineItemCreate {
	olic.mutation.SetCreatedAt(t)
	return olic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (olic *OrderLineItemCreate) SetNillableCreatedAt(t *time.Time) *OrderLineItemCreate {
	if t != nil {
		olic.SetCreatedAt(*t)
	}
	return olic
}

// SetUpdatedAt sets the "updated_at" field.
func (olic *OrderLineItemCreate) SetUpdatedAt(t time.Time) *OrderLineItemCreate {
	olic.mutation.SetUpdatedAt(t)
	return olic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (olic *OrderLineItemCreate) SetNillableUpdatedAt(t *time.Time) *OrderLineItemCreate {
	if t != nil {
		olic.SetUpdatedAt(*t)
	}
	return olic
}

// SetOrderID sets the "order_id" field.
func (olic *OrderLineItemCreate) SetOrderID(u uuid.UUID) *OrderLineItemCreate {
	olic.mutation.SetOrderID(u)
	return olic
}

// SetProductID sets the "product_id" field.
func (olic *OrderLineItemCreate) SetProductID(u uuid.UUID) *OrderLineItemCreate {
	olic.mutation.SetProductID(u)
	return olic
}

// SetQty sets the "qty" field.
func (olic *OrderLineItemCreate) SetQty(d decimal.Decimal) *OrderLineItemCreate {
	olic.mutation.SetQty(d)
	return olic
}

// SetID sets the "id" field.
func (olic *OrderLineItemCreate) SetID(u uuid.UUID) *OrderLineItemCreate {
	olic.mutation.SetID(u)
	return olic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (olic *OrderLineItemCreate) SetNillableID(u *uuid.UUID) *OrderLineItemCreate {
	if u != nil {
		olic.SetID(*u)
	}
	return olic
}

// SetOrder sets the "order" edge to the Order entity.
func (olic *OrderLineItemCreate) SetOrder(o *Order) *OrderLineItemCreate {
	return olic.SetOrderID(o.ID)
}

// SetProduct sets the "product" edge to the Product entity.
func (olic *OrderLineItemCreate) SetProduct(p *Product) *OrderLineItemCreate {
	return olic.SetProductID(p.ID)
}

// Mutation returns the OrderLineItemMutation object of the builder.
func (olic *OrderLineItemCreate) Mutation() *OrderLineItemMutation {
	return olic.mutation
}

// Save creates the OrderLineItem in the database.
func (olic *OrderLineItemCreate) Save(ctx context.Context) (*OrderLineItem, error) {
	olic.defaults()
	return withHooks(ctx, olic.sqlSave, olic.mutation, olic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (olic *OrderLineItemCreate) SaveX(ctx context.Context) *OrderLineItem {
	v, err := olic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (olic *OrderLineItemCreate) Exec(ctx context.Context) error {
	_, err := olic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olic *OrderLineItemCreate) ExecX(ctx context.Context) {
	if err := olic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (olic *OrderLineItemCreate) defaults() {
	if _, ok := olic.mutation.CreatedAt(); !ok {
		v := orderlineitem.DefaultCreatedAt()
		olic.mutation.SetCreatedAt(v)
	}
	if _, ok := olic.mutation.UpdatedAt(); !ok {
		v := orderlineitem.DefaultUpdatedAt()
		olic.mutation.SetUpdatedAt(v)
	}
	if _, ok := olic.mutation.ID(); !ok {
		v := orderlineitem.DefaultID()
		olic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (olic *OrderLineItemCreate) check() error {
	if _, ok := olic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderLineItem.created_at"`)}
	}
	if _, ok := olic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderLineItem.updated_at"`)}
	}
	if _, ok := olic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "OrderLineItem.order_id"`)}
	}
	if _, ok := olic.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`ent: missing required field "OrderLineItem.product_id"`)}
	}
	if _, ok := olic.mutation.Qty(); !ok {
		return &ValidationError{Name: "qty", err: errors.New(`ent: missing required field "OrderLineItem.qty"`)}
	}
	if _, ok := olic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required edge "OrderLineItem.order"`)}
	}
	if _, ok := olic.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product", err: errors.New(`ent: missing required edge "OrderLineItem.product"`)}
	}
	return nil
}

func (olic *OrderLineItemCreate) sqlSave(ctx context.Context) (*OrderLineItem, error) {
	if err := olic.check(); err != nil {
		return nil, err
	}
	_node, _spec := olic.createSpec()
	if err := sqlgraph.CreateNode(ctx, olic.driver, _spec); err != nil {
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
	olic.mutation.id = &_node.ID
	olic.mutation.done = true
	return _node, nil
}

func (olic *OrderLineItemCreate) createSpec() (*OrderLineItem, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderLineItem{config: olic.config}
		_spec = sqlgraph.NewCreateSpec(orderlineitem.Table, sqlgraph.NewFieldSpec(orderlineitem.FieldID, field.TypeUUID))
	)
	if id, ok := olic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := olic.mutation.CreatedAt(); ok {
		_spec.SetField(orderlineitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := olic.mutation.UpdatedAt(); ok {
		_spec.SetField(orderlineitem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := olic.mutation.Qty(); ok {
		_spec.SetField(orderlineitem.FieldQty, field.TypeFloat64, value)
		_node.Qty = value
	}
	if nodes := olic.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderlineitem.OrderTable,
			Columns: []string{orderlineitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := olic.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderlineitem.ProductTable,
			Columns: []string{orderlineitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProductID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderLineItemCreateBulk is the builder for creating many OrderLineItem entities in bulk.
type OrderLineItemCreateBulk struct {
	config
	err      error
	builders []*OrderLineItemCreate
}

// Save creates the OrderLineItem entities in the database.
func (olicb *OrderLineItemCreateBulk) Save(ctx context.Context) ([]*OrderLineItem, error) {
	if olicb.err != nil {
		return nil, olicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(olicb.builders))
	nodes := make([]*OrderLineItem, len(olicb.builders))
	mutators := make([]Mutator, len(olicb.builders))
	for i := range olicb.builders {
		func(i int, root context.Context) {
			builder := olicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderLineItemMutation)
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
					_, err = mutators[i+1].Mutate(root, olicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, olicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, olicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (olicb *OrderLineItemCreateBulk) SaveX(ctx context.Context) []*OrderLineItem {
	v, err := olicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (olicb *OrderLineItemCreateBulk) Exec(ctx context.Context) error {
	_, err := olicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olicb *OrderLineItemCreateBulk) ExecX(ctx context.Context) {
	if err := olicb.Exec(ctx); err != nil {
		panic(err)
	}
}
