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
	"github.com/thaiha1607/foursquare_server/ent/orderitem"
	"github.com/thaiha1607/foursquare_server/ent/productcolor"
	"github.com/thaiha1607/foursquare_server/ent/productinfo"
	"github.com/thaiha1607/foursquare_server/ent/workunitinfo"
)

// OrderItemCreate is the builder for creating a OrderItem entity.
type OrderItemCreate struct {
	config
	mutation *OrderItemMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oic *OrderItemCreate) SetCreatedAt(t time.Time) *OrderItemCreate {
	oic.mutation.SetCreatedAt(t)
	return oic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableCreatedAt(t *time.Time) *OrderItemCreate {
	if t != nil {
		oic.SetCreatedAt(*t)
	}
	return oic
}

// SetUpdatedAt sets the "updated_at" field.
func (oic *OrderItemCreate) SetUpdatedAt(t time.Time) *OrderItemCreate {
	oic.mutation.SetUpdatedAt(t)
	return oic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableUpdatedAt(t *time.Time) *OrderItemCreate {
	if t != nil {
		oic.SetUpdatedAt(*t)
	}
	return oic
}

// SetOrderID sets the "order_id" field.
func (oic *OrderItemCreate) SetOrderID(u uuid.UUID) *OrderItemCreate {
	oic.mutation.SetOrderID(u)
	return oic
}

// SetProductID sets the "product_id" field.
func (oic *OrderItemCreate) SetProductID(s string) *OrderItemCreate {
	oic.mutation.SetProductID(s)
	return oic
}

// SetProductColorID sets the "product_color_id" field.
func (oic *OrderItemCreate) SetProductColorID(s string) *OrderItemCreate {
	oic.mutation.SetProductColorID(s)
	return oic
}

// SetSrcUnitID sets the "src_unit_id" field.
func (oic *OrderItemCreate) SetSrcUnitID(u uuid.UUID) *OrderItemCreate {
	oic.mutation.SetSrcUnitID(u)
	return oic
}

// SetNillableSrcUnitID sets the "src_unit_id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableSrcUnitID(u *uuid.UUID) *OrderItemCreate {
	if u != nil {
		oic.SetSrcUnitID(*u)
	}
	return oic
}

// SetDstUnitID sets the "dst_unit_id" field.
func (oic *OrderItemCreate) SetDstUnitID(u uuid.UUID) *OrderItemCreate {
	oic.mutation.SetDstUnitID(u)
	return oic
}

// SetNillableDstUnitID sets the "dst_unit_id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableDstUnitID(u *uuid.UUID) *OrderItemCreate {
	if u != nil {
		oic.SetDstUnitID(*u)
	}
	return oic
}

// SetQty sets the "qty" field.
func (oic *OrderItemCreate) SetQty(d decimal.Decimal) *OrderItemCreate {
	oic.mutation.SetQty(d)
	return oic
}

// SetPricePerUnit sets the "price_per_unit" field.
func (oic *OrderItemCreate) SetPricePerUnit(d decimal.Decimal) *OrderItemCreate {
	oic.mutation.SetPricePerUnit(d)
	return oic
}

// SetStatus sets the "status" field.
func (oic *OrderItemCreate) SetStatus(o orderitem.Status) *OrderItemCreate {
	oic.mutation.SetStatus(o)
	return oic
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableStatus(o *orderitem.Status) *OrderItemCreate {
	if o != nil {
		oic.SetStatus(*o)
	}
	return oic
}

// SetID sets the "id" field.
func (oic *OrderItemCreate) SetID(u uuid.UUID) *OrderItemCreate {
	oic.mutation.SetID(u)
	return oic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableID(u *uuid.UUID) *OrderItemCreate {
	if u != nil {
		oic.SetID(*u)
	}
	return oic
}

// SetOrder sets the "order" edge to the Order entity.
func (oic *OrderItemCreate) SetOrder(o *Order) *OrderItemCreate {
	return oic.SetOrderID(o.ID)
}

// SetProduct sets the "product" edge to the ProductInfo entity.
func (oic *OrderItemCreate) SetProduct(p *ProductInfo) *OrderItemCreate {
	return oic.SetProductID(p.ID)
}

// SetProductColor sets the "product_color" edge to the ProductColor entity.
func (oic *OrderItemCreate) SetProductColor(p *ProductColor) *OrderItemCreate {
	return oic.SetProductColorID(p.ID)
}

// SetSourceWorkUnitID sets the "source_work_unit" edge to the WorkUnitInfo entity by ID.
func (oic *OrderItemCreate) SetSourceWorkUnitID(id uuid.UUID) *OrderItemCreate {
	oic.mutation.SetSourceWorkUnitID(id)
	return oic
}

// SetNillableSourceWorkUnitID sets the "source_work_unit" edge to the WorkUnitInfo entity by ID if the given value is not nil.
func (oic *OrderItemCreate) SetNillableSourceWorkUnitID(id *uuid.UUID) *OrderItemCreate {
	if id != nil {
		oic = oic.SetSourceWorkUnitID(*id)
	}
	return oic
}

// SetSourceWorkUnit sets the "source_work_unit" edge to the WorkUnitInfo entity.
func (oic *OrderItemCreate) SetSourceWorkUnit(w *WorkUnitInfo) *OrderItemCreate {
	return oic.SetSourceWorkUnitID(w.ID)
}

// SetDestinationWorkUnitID sets the "destination_work_unit" edge to the WorkUnitInfo entity by ID.
func (oic *OrderItemCreate) SetDestinationWorkUnitID(id uuid.UUID) *OrderItemCreate {
	oic.mutation.SetDestinationWorkUnitID(id)
	return oic
}

// SetNillableDestinationWorkUnitID sets the "destination_work_unit" edge to the WorkUnitInfo entity by ID if the given value is not nil.
func (oic *OrderItemCreate) SetNillableDestinationWorkUnitID(id *uuid.UUID) *OrderItemCreate {
	if id != nil {
		oic = oic.SetDestinationWorkUnitID(*id)
	}
	return oic
}

// SetDestinationWorkUnit sets the "destination_work_unit" edge to the WorkUnitInfo entity.
func (oic *OrderItemCreate) SetDestinationWorkUnit(w *WorkUnitInfo) *OrderItemCreate {
	return oic.SetDestinationWorkUnitID(w.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oic *OrderItemCreate) Mutation() *OrderItemMutation {
	return oic.mutation
}

// Save creates the OrderItem in the database.
func (oic *OrderItemCreate) Save(ctx context.Context) (*OrderItem, error) {
	oic.defaults()
	return withHooks(ctx, oic.sqlSave, oic.mutation, oic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oic *OrderItemCreate) SaveX(ctx context.Context) *OrderItem {
	v, err := oic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oic *OrderItemCreate) Exec(ctx context.Context) error {
	_, err := oic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oic *OrderItemCreate) ExecX(ctx context.Context) {
	if err := oic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oic *OrderItemCreate) defaults() {
	if _, ok := oic.mutation.CreatedAt(); !ok {
		v := orderitem.DefaultCreatedAt()
		oic.mutation.SetCreatedAt(v)
	}
	if _, ok := oic.mutation.UpdatedAt(); !ok {
		v := orderitem.DefaultUpdatedAt()
		oic.mutation.SetUpdatedAt(v)
	}
	if _, ok := oic.mutation.Status(); !ok {
		v := orderitem.DefaultStatus
		oic.mutation.SetStatus(v)
	}
	if _, ok := oic.mutation.ID(); !ok {
		v := orderitem.DefaultID()
		oic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oic *OrderItemCreate) check() error {
	if _, ok := oic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderItem.created_at"`)}
	}
	if _, ok := oic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderItem.updated_at"`)}
	}
	if _, ok := oic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "OrderItem.order_id"`)}
	}
	if _, ok := oic.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`ent: missing required field "OrderItem.product_id"`)}
	}
	if _, ok := oic.mutation.ProductColorID(); !ok {
		return &ValidationError{Name: "product_color_id", err: errors.New(`ent: missing required field "OrderItem.product_color_id"`)}
	}
	if _, ok := oic.mutation.Qty(); !ok {
		return &ValidationError{Name: "qty", err: errors.New(`ent: missing required field "OrderItem.qty"`)}
	}
	if _, ok := oic.mutation.PricePerUnit(); !ok {
		return &ValidationError{Name: "price_per_unit", err: errors.New(`ent: missing required field "OrderItem.price_per_unit"`)}
	}
	if _, ok := oic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "OrderItem.status"`)}
	}
	if v, ok := oic.mutation.Status(); ok {
		if err := orderitem.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "OrderItem.status": %w`, err)}
		}
	}
	if _, ok := oic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required edge "OrderItem.order"`)}
	}
	if _, ok := oic.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product", err: errors.New(`ent: missing required edge "OrderItem.product"`)}
	}
	if _, ok := oic.mutation.ProductColorID(); !ok {
		return &ValidationError{Name: "product_color", err: errors.New(`ent: missing required edge "OrderItem.product_color"`)}
	}
	return nil
}

func (oic *OrderItemCreate) sqlSave(ctx context.Context) (*OrderItem, error) {
	if err := oic.check(); err != nil {
		return nil, err
	}
	_node, _spec := oic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oic.driver, _spec); err != nil {
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
	oic.mutation.id = &_node.ID
	oic.mutation.done = true
	return _node, nil
}

func (oic *OrderItemCreate) createSpec() (*OrderItem, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderItem{config: oic.config}
		_spec = sqlgraph.NewCreateSpec(orderitem.Table, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeUUID))
	)
	if id, ok := oic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oic.mutation.CreatedAt(); ok {
		_spec.SetField(orderitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := oic.mutation.UpdatedAt(); ok {
		_spec.SetField(orderitem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := oic.mutation.Qty(); ok {
		_spec.SetField(orderitem.FieldQty, field.TypeFloat64, value)
		_node.Qty = &value
	}
	if value, ok := oic.mutation.PricePerUnit(); ok {
		_spec.SetField(orderitem.FieldPricePerUnit, field.TypeFloat64, value)
		_node.PricePerUnit = &value
	}
	if value, ok := oic.mutation.Status(); ok {
		_spec.SetField(orderitem.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if nodes := oic.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
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
	if nodes := oic.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderitem.ProductTable,
			Columns: []string{orderitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productinfo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProductID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.ProductColorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderitem.ProductColorTable,
			Columns: []string{orderitem.ProductColorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productcolor.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProductColorID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.SourceWorkUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderitem.SourceWorkUnitTable,
			Columns: []string{orderitem.SourceWorkUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workunitinfo.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SrcUnitID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.DestinationWorkUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderitem.DestinationWorkUnitTable,
			Columns: []string{orderitem.DestinationWorkUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workunitinfo.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DstUnitID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderItemCreateBulk is the builder for creating many OrderItem entities in bulk.
type OrderItemCreateBulk struct {
	config
	err      error
	builders []*OrderItemCreate
}

// Save creates the OrderItem entities in the database.
func (oicb *OrderItemCreateBulk) Save(ctx context.Context) ([]*OrderItem, error) {
	if oicb.err != nil {
		return nil, oicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oicb.builders))
	nodes := make([]*OrderItem, len(oicb.builders))
	mutators := make([]Mutator, len(oicb.builders))
	for i := range oicb.builders {
		func(i int, root context.Context) {
			builder := oicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderItemMutation)
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
					_, err = mutators[i+1].Mutate(root, oicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) SaveX(ctx context.Context) []*OrderItem {
	v, err := oicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oicb *OrderItemCreateBulk) Exec(ctx context.Context) error {
	_, err := oicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) ExecX(ctx context.Context) {
	if err := oicb.Exec(ctx); err != nil {
		panic(err)
	}
}
