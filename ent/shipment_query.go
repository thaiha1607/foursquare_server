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
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/order"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/shipment"
)

// ShipmentQuery is the builder for querying Shipment entities.
type ShipmentQuery struct {
	config
	ctx         *QueryContext
	order       []shipment.OrderOption
	inters      []Interceptor
	predicates  []predicate.Shipment
	withOrder   *OrderQuery
	withInvoice *InvoiceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShipmentQuery builder.
func (sq *ShipmentQuery) Where(ps ...predicate.Shipment) *ShipmentQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *ShipmentQuery) Limit(limit int) *ShipmentQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *ShipmentQuery) Offset(offset int) *ShipmentQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *ShipmentQuery) Unique(unique bool) *ShipmentQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *ShipmentQuery) Order(o ...shipment.OrderOption) *ShipmentQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryOrder chains the current query on the "order" edge.
func (sq *ShipmentQuery) QueryOrder() *OrderQuery {
	query := (&OrderClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipment.Table, shipment.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, shipment.OrderTable, shipment.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInvoice chains the current query on the "invoice" edge.
func (sq *ShipmentQuery) QueryInvoice() *InvoiceQuery {
	query := (&InvoiceClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipment.Table, shipment.FieldID, selector),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, shipment.InvoiceTable, shipment.InvoiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Shipment entity from the query.
// Returns a *NotFoundError when no Shipment was found.
func (sq *ShipmentQuery) First(ctx context.Context) (*Shipment, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shipment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ShipmentQuery) FirstX(ctx context.Context) *Shipment {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Shipment ID from the query.
// Returns a *NotFoundError when no Shipment ID was found.
func (sq *ShipmentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shipment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *ShipmentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Shipment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Shipment entity is found.
// Returns a *NotFoundError when no Shipment entities are found.
func (sq *ShipmentQuery) Only(ctx context.Context) (*Shipment, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shipment.Label}
	default:
		return nil, &NotSingularError{shipment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ShipmentQuery) OnlyX(ctx context.Context) *Shipment {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Shipment ID in the query.
// Returns a *NotSingularError when more than one Shipment ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *ShipmentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shipment.Label}
	default:
		err = &NotSingularError{shipment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *ShipmentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Shipments.
func (sq *ShipmentQuery) All(ctx context.Context) ([]*Shipment, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Shipment, *ShipmentQuery]()
	return withInterceptors[[]*Shipment](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *ShipmentQuery) AllX(ctx context.Context) []*Shipment {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Shipment IDs.
func (sq *ShipmentQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(shipment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ShipmentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ShipmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*ShipmentQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ShipmentQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ShipmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ShipmentQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShipmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ShipmentQuery) Clone() *ShipmentQuery {
	if sq == nil {
		return nil
	}
	return &ShipmentQuery{
		config:      sq.config,
		ctx:         sq.ctx.Clone(),
		order:       append([]shipment.OrderOption{}, sq.order...),
		inters:      append([]Interceptor{}, sq.inters...),
		predicates:  append([]predicate.Shipment{}, sq.predicates...),
		withOrder:   sq.withOrder.Clone(),
		withInvoice: sq.withInvoice.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithOrder tells the query-builder to eager-load the nodes that are connected to
// the "order" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ShipmentQuery) WithOrder(opts ...func(*OrderQuery)) *ShipmentQuery {
	query := (&OrderClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withOrder = query
	return sq
}

// WithInvoice tells the query-builder to eager-load the nodes that are connected to
// the "invoice" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ShipmentQuery) WithInvoice(opts ...func(*InvoiceQuery)) *ShipmentQuery {
	query := (&InvoiceClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withInvoice = query
	return sq
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
//	client.Shipment.Query().
//		GroupBy(shipment.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *ShipmentQuery) GroupBy(field string, fields ...string) *ShipmentGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ShipmentGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = shipment.Label
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
//	client.Shipment.Query().
//		Select(shipment.FieldCreatedAt).
//		Scan(ctx, &v)
func (sq *ShipmentQuery) Select(fields ...string) *ShipmentSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &ShipmentSelect{ShipmentQuery: sq}
	sbuild.label = shipment.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ShipmentSelect configured with the given aggregations.
func (sq *ShipmentQuery) Aggregate(fns ...AggregateFunc) *ShipmentSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *ShipmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !shipment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *ShipmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Shipment, error) {
	var (
		nodes       = []*Shipment{}
		_spec       = sq.querySpec()
		loadedTypes = [2]bool{
			sq.withOrder != nil,
			sq.withInvoice != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Shipment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Shipment{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withOrder; query != nil {
		if err := sq.loadOrder(ctx, query, nodes, nil,
			func(n *Shipment, e *Order) { n.Edges.Order = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withInvoice; query != nil {
		if err := sq.loadInvoice(ctx, query, nodes, nil,
			func(n *Shipment, e *Invoice) { n.Edges.Invoice = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *ShipmentQuery) loadOrder(ctx context.Context, query *OrderQuery, nodes []*Shipment, init func(*Shipment), assign func(*Shipment, *Order)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Shipment)
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
func (sq *ShipmentQuery) loadInvoice(ctx context.Context, query *InvoiceQuery, nodes []*Shipment, init func(*Shipment), assign func(*Shipment, *Invoice)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Shipment)
	for i := range nodes {
		fk := nodes[i].InvoiceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(invoice.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "invoice_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *ShipmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ShipmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(shipment.Table, shipment.Columns, sqlgraph.NewFieldSpec(shipment.FieldID, field.TypeUUID))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shipment.FieldID)
		for i := range fields {
			if fields[i] != shipment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if sq.withOrder != nil {
			_spec.Node.AddColumnOnce(shipment.FieldOrderID)
		}
		if sq.withInvoice != nil {
			_spec.Node.AddColumnOnce(shipment.FieldInvoiceID)
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ShipmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(shipment.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = shipment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShipmentGroupBy is the group-by builder for Shipment entities.
type ShipmentGroupBy struct {
	selector
	build *ShipmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ShipmentGroupBy) Aggregate(fns ...AggregateFunc) *ShipmentGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *ShipmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShipmentQuery, *ShipmentGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *ShipmentGroupBy) sqlScan(ctx context.Context, root *ShipmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ShipmentSelect is the builder for selecting fields of Shipment entities.
type ShipmentSelect struct {
	*ShipmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *ShipmentSelect) Aggregate(fns ...AggregateFunc) *ShipmentSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *ShipmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShipmentQuery, *ShipmentSelect](ctx, ss.ShipmentQuery, ss, ss.inters, v)
}

func (ss *ShipmentSelect) sqlScan(ctx context.Context, root *ShipmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
