// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/orderitem"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/shipment"
	"github.com/thaiha1607/foursquare_server/ent/shipmentitem"
)

// ShipmentItemQuery is the builder for querying ShipmentItem entities.
type ShipmentItemQuery struct {
	config
	ctx           *QueryContext
	order         []shipmentitem.OrderOption
	inters        []Interceptor
	predicates    []predicate.ShipmentItem
	withShipment  *ShipmentQuery
	withOrderItem *OrderItemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShipmentItemQuery builder.
func (siq *ShipmentItemQuery) Where(ps ...predicate.ShipmentItem) *ShipmentItemQuery {
	siq.predicates = append(siq.predicates, ps...)
	return siq
}

// Limit the number of records to be returned by this query.
func (siq *ShipmentItemQuery) Limit(limit int) *ShipmentItemQuery {
	siq.ctx.Limit = &limit
	return siq
}

// Offset to start from.
func (siq *ShipmentItemQuery) Offset(offset int) *ShipmentItemQuery {
	siq.ctx.Offset = &offset
	return siq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (siq *ShipmentItemQuery) Unique(unique bool) *ShipmentItemQuery {
	siq.ctx.Unique = &unique
	return siq
}

// Order specifies how the records should be ordered.
func (siq *ShipmentItemQuery) Order(o ...shipmentitem.OrderOption) *ShipmentItemQuery {
	siq.order = append(siq.order, o...)
	return siq
}

// QueryShipment chains the current query on the "shipment" edge.
func (siq *ShipmentItemQuery) QueryShipment() *ShipmentQuery {
	query := (&ShipmentClient{config: siq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := siq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := siq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentitem.Table, shipmentitem.FieldID, selector),
			sqlgraph.To(shipment.Table, shipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, shipmentitem.ShipmentTable, shipmentitem.ShipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(siq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderItem chains the current query on the "order_item" edge.
func (siq *ShipmentItemQuery) QueryOrderItem() *OrderItemQuery {
	query := (&OrderItemClient{config: siq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := siq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := siq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentitem.Table, shipmentitem.FieldID, selector),
			sqlgraph.To(orderitem.Table, orderitem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, shipmentitem.OrderItemTable, shipmentitem.OrderItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(siq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShipmentItem entity from the query.
// Returns a *NotFoundError when no ShipmentItem was found.
func (siq *ShipmentItemQuery) First(ctx context.Context) (*ShipmentItem, error) {
	nodes, err := siq.Limit(1).All(setContextOp(ctx, siq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shipmentitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (siq *ShipmentItemQuery) FirstX(ctx context.Context) *ShipmentItem {
	node, err := siq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShipmentItem ID from the query.
// Returns a *NotFoundError when no ShipmentItem ID was found.
func (siq *ShipmentItemQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = siq.Limit(1).IDs(setContextOp(ctx, siq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shipmentitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (siq *ShipmentItemQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := siq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShipmentItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ShipmentItem entity is found.
// Returns a *NotFoundError when no ShipmentItem entities are found.
func (siq *ShipmentItemQuery) Only(ctx context.Context) (*ShipmentItem, error) {
	nodes, err := siq.Limit(2).All(setContextOp(ctx, siq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shipmentitem.Label}
	default:
		return nil, &NotSingularError{shipmentitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (siq *ShipmentItemQuery) OnlyX(ctx context.Context) *ShipmentItem {
	node, err := siq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShipmentItem ID in the query.
// Returns a *NotSingularError when more than one ShipmentItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (siq *ShipmentItemQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = siq.Limit(2).IDs(setContextOp(ctx, siq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shipmentitem.Label}
	default:
		err = &NotSingularError{shipmentitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (siq *ShipmentItemQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := siq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShipmentItems.
func (siq *ShipmentItemQuery) All(ctx context.Context) ([]*ShipmentItem, error) {
	ctx = setContextOp(ctx, siq.ctx, "All")
	if err := siq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ShipmentItem, *ShipmentItemQuery]()
	return withInterceptors[[]*ShipmentItem](ctx, siq, qr, siq.inters)
}

// AllX is like All, but panics if an error occurs.
func (siq *ShipmentItemQuery) AllX(ctx context.Context) []*ShipmentItem {
	nodes, err := siq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShipmentItem IDs.
func (siq *ShipmentItemQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if siq.ctx.Unique == nil && siq.path != nil {
		siq.Unique(true)
	}
	ctx = setContextOp(ctx, siq.ctx, "IDs")
	if err = siq.Select(shipmentitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (siq *ShipmentItemQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := siq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (siq *ShipmentItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, siq.ctx, "Count")
	if err := siq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, siq, querierCount[*ShipmentItemQuery](), siq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (siq *ShipmentItemQuery) CountX(ctx context.Context) int {
	count, err := siq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (siq *ShipmentItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, siq.ctx, "Exist")
	switch _, err := siq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (siq *ShipmentItemQuery) ExistX(ctx context.Context) bool {
	exist, err := siq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShipmentItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (siq *ShipmentItemQuery) Clone() *ShipmentItemQuery {
	if siq == nil {
		return nil
	}
	return &ShipmentItemQuery{
		config:        siq.config,
		ctx:           siq.ctx.Clone(),
		order:         append([]shipmentitem.OrderOption{}, siq.order...),
		inters:        append([]Interceptor{}, siq.inters...),
		predicates:    append([]predicate.ShipmentItem{}, siq.predicates...),
		withShipment:  siq.withShipment.Clone(),
		withOrderItem: siq.withOrderItem.Clone(),
		// clone intermediate query.
		sql:  siq.sql.Clone(),
		path: siq.path,
	}
}

// WithShipment tells the query-builder to eager-load the nodes that are connected to
// the "shipment" edge. The optional arguments are used to configure the query builder of the edge.
func (siq *ShipmentItemQuery) WithShipment(opts ...func(*ShipmentQuery)) *ShipmentItemQuery {
	query := (&ShipmentClient{config: siq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	siq.withShipment = query
	return siq
}

// WithOrderItem tells the query-builder to eager-load the nodes that are connected to
// the "order_item" edge. The optional arguments are used to configure the query builder of the edge.
func (siq *ShipmentItemQuery) WithOrderItem(opts ...func(*OrderItemQuery)) *ShipmentItemQuery {
	query := (&OrderItemClient{config: siq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	siq.withOrderItem = query
	return siq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ShipmentItem.Query().
//		GroupBy(shipmentitem.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (siq *ShipmentItemQuery) GroupBy(field string, fields ...string) *ShipmentItemGroupBy {
	siq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ShipmentItemGroupBy{build: siq}
	grbuild.flds = &siq.ctx.Fields
	grbuild.label = shipmentitem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.ShipmentItem.Query().
//		Select(shipmentitem.FieldCreatedAt).
//		Scan(ctx, &v)
func (siq *ShipmentItemQuery) Select(fields ...string) *ShipmentItemSelect {
	siq.ctx.Fields = append(siq.ctx.Fields, fields...)
	sbuild := &ShipmentItemSelect{ShipmentItemQuery: siq}
	sbuild.label = shipmentitem.Label
	sbuild.flds, sbuild.scan = &siq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ShipmentItemSelect configured with the given aggregations.
func (siq *ShipmentItemQuery) Aggregate(fns ...AggregateFunc) *ShipmentItemSelect {
	return siq.Select().Aggregate(fns...)
}

func (siq *ShipmentItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range siq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, siq); err != nil {
				return err
			}
		}
	}
	for _, f := range siq.ctx.Fields {
		if !shipmentitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if siq.path != nil {
		prev, err := siq.path(ctx)
		if err != nil {
			return err
		}
		siq.sql = prev
	}
	return nil
}

func (siq *ShipmentItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ShipmentItem, error) {
	var (
		nodes       = []*ShipmentItem{}
		_spec       = siq.querySpec()
		loadedTypes = [2]bool{
			siq.withShipment != nil,
			siq.withOrderItem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ShipmentItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ShipmentItem{config: siq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, siq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := siq.withShipment; query != nil {
		if err := siq.loadShipment(ctx, query, nodes, nil,
			func(n *ShipmentItem, e *Shipment) { n.Edges.Shipment = e }); err != nil {
			return nil, err
		}
	}
	if query := siq.withOrderItem; query != nil {
		if err := siq.loadOrderItem(ctx, query, nodes, nil,
			func(n *ShipmentItem, e *OrderItem) { n.Edges.OrderItem = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (siq *ShipmentItemQuery) loadShipment(ctx context.Context, query *ShipmentQuery, nodes []*ShipmentItem, init func(*ShipmentItem), assign func(*ShipmentItem, *Shipment)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ShipmentItem)
	for i := range nodes {
		fk := nodes[i].ShipmentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(shipment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "shipment_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (siq *ShipmentItemQuery) loadOrderItem(ctx context.Context, query *OrderItemQuery, nodes []*ShipmentItem, init func(*ShipmentItem), assign func(*ShipmentItem, *OrderItem)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentItem)
	for i := range nodes {
		fk := nodes[i].OrderItemID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(orderitem.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_item_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (siq *ShipmentItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := siq.querySpec()
	_spec.Node.Columns = siq.ctx.Fields
	if len(siq.ctx.Fields) > 0 {
		_spec.Unique = siq.ctx.Unique != nil && *siq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, siq.driver, _spec)
}

func (siq *ShipmentItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(shipmentitem.Table, shipmentitem.Columns, sqlgraph.NewFieldSpec(shipmentitem.FieldID, field.TypeUUID))
	_spec.From = siq.sql
	if unique := siq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if siq.path != nil {
		_spec.Unique = true
	}
	if fields := siq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentitem.FieldID)
		for i := range fields {
			if fields[i] != shipmentitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if siq.withShipment != nil {
			_spec.Node.AddColumnOnce(shipmentitem.FieldShipmentID)
		}
		if siq.withOrderItem != nil {
			_spec.Node.AddColumnOnce(shipmentitem.FieldOrderItemID)
		}
	}
	if ps := siq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := siq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := siq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := siq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (siq *ShipmentItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(siq.driver.Dialect())
	t1 := builder.Table(shipmentitem.Table)
	columns := siq.ctx.Fields
	if len(columns) == 0 {
		columns = shipmentitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if siq.sql != nil {
		selector = siq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if siq.ctx.Unique != nil && *siq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range siq.predicates {
		p(selector)
	}
	for _, p := range siq.order {
		p(selector)
	}
	if offset := siq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := siq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShipmentItemGroupBy is the group-by builder for ShipmentItem entities.
type ShipmentItemGroupBy struct {
	selector
	build *ShipmentItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sigb *ShipmentItemGroupBy) Aggregate(fns ...AggregateFunc) *ShipmentItemGroupBy {
	sigb.fns = append(sigb.fns, fns...)
	return sigb
}

// Scan applies the selector query and scans the result into the given value.
func (sigb *ShipmentItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sigb.build.ctx, "GroupBy")
	if err := sigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShipmentItemQuery, *ShipmentItemGroupBy](ctx, sigb.build, sigb, sigb.build.inters, v)
}

func (sigb *ShipmentItemGroupBy) sqlScan(ctx context.Context, root *ShipmentItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sigb.fns))
	for _, fn := range sigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sigb.flds)+len(sigb.fns))
		for _, f := range *sigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ShipmentItemSelect is the builder for selecting fields of ShipmentItem entities.
type ShipmentItemSelect struct {
	*ShipmentItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sis *ShipmentItemSelect) Aggregate(fns ...AggregateFunc) *ShipmentItemSelect {
	sis.fns = append(sis.fns, fns...)
	return sis
}

// Scan applies the selector query and scans the result into the given value.
func (sis *ShipmentItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sis.ctx, "Select")
	if err := sis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShipmentItemQuery, *ShipmentItemSelect](ctx, sis.ShipmentItemQuery, sis, sis.inters, v)
}

func (sis *ShipmentItemSelect) sqlScan(ctx context.Context, root *ShipmentItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sis.fns))
	for _, fn := range sis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
