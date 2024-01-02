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
	"github.com/thaiha1607/foursquare_server/ent/order"
	"github.com/thaiha1607/foursquare_server/ent/orderlineitem"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/product"
)

// OrderLineItemQuery is the builder for querying OrderLineItem entities.
type OrderLineItemQuery struct {
	config
	ctx         *QueryContext
	order       []orderlineitem.OrderOption
	inters      []Interceptor
	predicates  []predicate.OrderLineItem
	withOrder   *OrderQuery
	withProduct *ProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderLineItemQuery builder.
func (oliq *OrderLineItemQuery) Where(ps ...predicate.OrderLineItem) *OrderLineItemQuery {
	oliq.predicates = append(oliq.predicates, ps...)
	return oliq
}

// Limit the number of records to be returned by this query.
func (oliq *OrderLineItemQuery) Limit(limit int) *OrderLineItemQuery {
	oliq.ctx.Limit = &limit
	return oliq
}

// Offset to start from.
func (oliq *OrderLineItemQuery) Offset(offset int) *OrderLineItemQuery {
	oliq.ctx.Offset = &offset
	return oliq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oliq *OrderLineItemQuery) Unique(unique bool) *OrderLineItemQuery {
	oliq.ctx.Unique = &unique
	return oliq
}

// Order specifies how the records should be ordered.
func (oliq *OrderLineItemQuery) Order(o ...orderlineitem.OrderOption) *OrderLineItemQuery {
	oliq.order = append(oliq.order, o...)
	return oliq
}

// QueryOrder chains the current query on the "order" edge.
func (oliq *OrderLineItemQuery) QueryOrder() *OrderQuery {
	query := (&OrderClient{config: oliq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oliq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oliq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderlineitem.Table, orderlineitem.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orderlineitem.OrderTable, orderlineitem.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(oliq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the "product" edge.
func (oliq *OrderLineItemQuery) QueryProduct() *ProductQuery {
	query := (&ProductClient{config: oliq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oliq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oliq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderlineitem.Table, orderlineitem.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orderlineitem.ProductTable, orderlineitem.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(oliq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderLineItem entity from the query.
// Returns a *NotFoundError when no OrderLineItem was found.
func (oliq *OrderLineItemQuery) First(ctx context.Context) (*OrderLineItem, error) {
	nodes, err := oliq.Limit(1).All(setContextOp(ctx, oliq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderlineitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oliq *OrderLineItemQuery) FirstX(ctx context.Context) *OrderLineItem {
	node, err := oliq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderLineItem ID from the query.
// Returns a *NotFoundError when no OrderLineItem ID was found.
func (oliq *OrderLineItemQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oliq.Limit(1).IDs(setContextOp(ctx, oliq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderlineitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oliq *OrderLineItemQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := oliq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderLineItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderLineItem entity is found.
// Returns a *NotFoundError when no OrderLineItem entities are found.
func (oliq *OrderLineItemQuery) Only(ctx context.Context) (*OrderLineItem, error) {
	nodes, err := oliq.Limit(2).All(setContextOp(ctx, oliq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderlineitem.Label}
	default:
		return nil, &NotSingularError{orderlineitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oliq *OrderLineItemQuery) OnlyX(ctx context.Context) *OrderLineItem {
	node, err := oliq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderLineItem ID in the query.
// Returns a *NotSingularError when more than one OrderLineItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (oliq *OrderLineItemQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oliq.Limit(2).IDs(setContextOp(ctx, oliq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderlineitem.Label}
	default:
		err = &NotSingularError{orderlineitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oliq *OrderLineItemQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := oliq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderLineItems.
func (oliq *OrderLineItemQuery) All(ctx context.Context) ([]*OrderLineItem, error) {
	ctx = setContextOp(ctx, oliq.ctx, "All")
	if err := oliq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderLineItem, *OrderLineItemQuery]()
	return withInterceptors[[]*OrderLineItem](ctx, oliq, qr, oliq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oliq *OrderLineItemQuery) AllX(ctx context.Context) []*OrderLineItem {
	nodes, err := oliq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderLineItem IDs.
func (oliq *OrderLineItemQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if oliq.ctx.Unique == nil && oliq.path != nil {
		oliq.Unique(true)
	}
	ctx = setContextOp(ctx, oliq.ctx, "IDs")
	if err = oliq.Select(orderlineitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oliq *OrderLineItemQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := oliq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oliq *OrderLineItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oliq.ctx, "Count")
	if err := oliq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oliq, querierCount[*OrderLineItemQuery](), oliq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oliq *OrderLineItemQuery) CountX(ctx context.Context) int {
	count, err := oliq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oliq *OrderLineItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oliq.ctx, "Exist")
	switch _, err := oliq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oliq *OrderLineItemQuery) ExistX(ctx context.Context) bool {
	exist, err := oliq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderLineItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oliq *OrderLineItemQuery) Clone() *OrderLineItemQuery {
	if oliq == nil {
		return nil
	}
	return &OrderLineItemQuery{
		config:      oliq.config,
		ctx:         oliq.ctx.Clone(),
		order:       append([]orderlineitem.OrderOption{}, oliq.order...),
		inters:      append([]Interceptor{}, oliq.inters...),
		predicates:  append([]predicate.OrderLineItem{}, oliq.predicates...),
		withOrder:   oliq.withOrder.Clone(),
		withProduct: oliq.withProduct.Clone(),
		// clone intermediate query.
		sql:  oliq.sql.Clone(),
		path: oliq.path,
	}
}

// WithOrder tells the query-builder to eager-load the nodes that are connected to
// the "order" edge. The optional arguments are used to configure the query builder of the edge.
func (oliq *OrderLineItemQuery) WithOrder(opts ...func(*OrderQuery)) *OrderLineItemQuery {
	query := (&OrderClient{config: oliq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oliq.withOrder = query
	return oliq
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (oliq *OrderLineItemQuery) WithProduct(opts ...func(*ProductQuery)) *OrderLineItemQuery {
	query := (&ProductClient{config: oliq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oliq.withProduct = query
	return oliq
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
//	client.OrderLineItem.Query().
//		GroupBy(orderlineitem.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oliq *OrderLineItemQuery) GroupBy(field string, fields ...string) *OrderLineItemGroupBy {
	oliq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderLineItemGroupBy{build: oliq}
	grbuild.flds = &oliq.ctx.Fields
	grbuild.label = orderlineitem.Label
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
//	client.OrderLineItem.Query().
//		Select(orderlineitem.FieldCreatedAt).
//		Scan(ctx, &v)
func (oliq *OrderLineItemQuery) Select(fields ...string) *OrderLineItemSelect {
	oliq.ctx.Fields = append(oliq.ctx.Fields, fields...)
	sbuild := &OrderLineItemSelect{OrderLineItemQuery: oliq}
	sbuild.label = orderlineitem.Label
	sbuild.flds, sbuild.scan = &oliq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderLineItemSelect configured with the given aggregations.
func (oliq *OrderLineItemQuery) Aggregate(fns ...AggregateFunc) *OrderLineItemSelect {
	return oliq.Select().Aggregate(fns...)
}

func (oliq *OrderLineItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oliq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oliq); err != nil {
				return err
			}
		}
	}
	for _, f := range oliq.ctx.Fields {
		if !orderlineitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oliq.path != nil {
		prev, err := oliq.path(ctx)
		if err != nil {
			return err
		}
		oliq.sql = prev
	}
	return nil
}

func (oliq *OrderLineItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderLineItem, error) {
	var (
		nodes       = []*OrderLineItem{}
		_spec       = oliq.querySpec()
		loadedTypes = [2]bool{
			oliq.withOrder != nil,
			oliq.withProduct != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderLineItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderLineItem{config: oliq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oliq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oliq.withOrder; query != nil {
		if err := oliq.loadOrder(ctx, query, nodes, nil,
			func(n *OrderLineItem, e *Order) { n.Edges.Order = e }); err != nil {
			return nil, err
		}
	}
	if query := oliq.withProduct; query != nil {
		if err := oliq.loadProduct(ctx, query, nodes, nil,
			func(n *OrderLineItem, e *Product) { n.Edges.Product = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oliq *OrderLineItemQuery) loadOrder(ctx context.Context, query *OrderQuery, nodes []*OrderLineItem, init func(*OrderLineItem), assign func(*OrderLineItem, *Order)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*OrderLineItem)
	for i := range nodes {
		fk := nodes[i].OrderID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(order.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (oliq *OrderLineItemQuery) loadProduct(ctx context.Context, query *ProductQuery, nodes []*OrderLineItem, init func(*OrderLineItem), assign func(*OrderLineItem, *Product)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*OrderLineItem)
	for i := range nodes {
		fk := nodes[i].ProductID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(product.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "product_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (oliq *OrderLineItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oliq.querySpec()
	_spec.Node.Columns = oliq.ctx.Fields
	if len(oliq.ctx.Fields) > 0 {
		_spec.Unique = oliq.ctx.Unique != nil && *oliq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oliq.driver, _spec)
}

func (oliq *OrderLineItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orderlineitem.Table, orderlineitem.Columns, sqlgraph.NewFieldSpec(orderlineitem.FieldID, field.TypeUUID))
	_spec.From = oliq.sql
	if unique := oliq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oliq.path != nil {
		_spec.Unique = true
	}
	if fields := oliq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderlineitem.FieldID)
		for i := range fields {
			if fields[i] != orderlineitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if oliq.withOrder != nil {
			_spec.Node.AddColumnOnce(orderlineitem.FieldOrderID)
		}
		if oliq.withProduct != nil {
			_spec.Node.AddColumnOnce(orderlineitem.FieldProductID)
		}
	}
	if ps := oliq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oliq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oliq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oliq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oliq *OrderLineItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oliq.driver.Dialect())
	t1 := builder.Table(orderlineitem.Table)
	columns := oliq.ctx.Fields
	if len(columns) == 0 {
		columns = orderlineitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oliq.sql != nil {
		selector = oliq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oliq.ctx.Unique != nil && *oliq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oliq.predicates {
		p(selector)
	}
	for _, p := range oliq.order {
		p(selector)
	}
	if offset := oliq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oliq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderLineItemGroupBy is the group-by builder for OrderLineItem entities.
type OrderLineItemGroupBy struct {
	selector
	build *OrderLineItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oligb *OrderLineItemGroupBy) Aggregate(fns ...AggregateFunc) *OrderLineItemGroupBy {
	oligb.fns = append(oligb.fns, fns...)
	return oligb
}

// Scan applies the selector query and scans the result into the given value.
func (oligb *OrderLineItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oligb.build.ctx, "GroupBy")
	if err := oligb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderLineItemQuery, *OrderLineItemGroupBy](ctx, oligb.build, oligb, oligb.build.inters, v)
}

func (oligb *OrderLineItemGroupBy) sqlScan(ctx context.Context, root *OrderLineItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(oligb.fns))
	for _, fn := range oligb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*oligb.flds)+len(oligb.fns))
		for _, f := range *oligb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*oligb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oligb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderLineItemSelect is the builder for selecting fields of OrderLineItem entities.
type OrderLineItemSelect struct {
	*OrderLineItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (olis *OrderLineItemSelect) Aggregate(fns ...AggregateFunc) *OrderLineItemSelect {
	olis.fns = append(olis.fns, fns...)
	return olis
}

// Scan applies the selector query and scans the result into the given value.
func (olis *OrderLineItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, olis.ctx, "Select")
	if err := olis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderLineItemQuery, *OrderLineItemSelect](ctx, olis.OrderLineItemQuery, olis, olis.inters, v)
}

func (olis *OrderLineItemSelect) sqlScan(ctx context.Context, root *OrderLineItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(olis.fns))
	for _, fn := range olis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*olis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := olis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}