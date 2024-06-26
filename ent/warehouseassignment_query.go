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
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/warehouseassignment"
	"github.com/thaiha1607/foursquare_server/ent/workunitinfo"
)

// WarehouseAssignmentQuery is the builder for querying WarehouseAssignment entities.
type WarehouseAssignmentQuery struct {
	config
	ctx          *QueryContext
	order        []warehouseassignment.OrderOption
	inters       []Interceptor
	predicates   []predicate.WarehouseAssignment
	withOrder    *OrderQuery
	withWorkUnit *WorkUnitInfoQuery
	withStaff    *PersonQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WarehouseAssignmentQuery builder.
func (waq *WarehouseAssignmentQuery) Where(ps ...predicate.WarehouseAssignment) *WarehouseAssignmentQuery {
	waq.predicates = append(waq.predicates, ps...)
	return waq
}

// Limit the number of records to be returned by this query.
func (waq *WarehouseAssignmentQuery) Limit(limit int) *WarehouseAssignmentQuery {
	waq.ctx.Limit = &limit
	return waq
}

// Offset to start from.
func (waq *WarehouseAssignmentQuery) Offset(offset int) *WarehouseAssignmentQuery {
	waq.ctx.Offset = &offset
	return waq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (waq *WarehouseAssignmentQuery) Unique(unique bool) *WarehouseAssignmentQuery {
	waq.ctx.Unique = &unique
	return waq
}

// Order specifies how the records should be ordered.
func (waq *WarehouseAssignmentQuery) Order(o ...warehouseassignment.OrderOption) *WarehouseAssignmentQuery {
	waq.order = append(waq.order, o...)
	return waq
}

// QueryOrder chains the current query on the "order" edge.
func (waq *WarehouseAssignmentQuery) QueryOrder() *OrderQuery {
	query := (&OrderClient{config: waq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := waq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := waq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(warehouseassignment.Table, warehouseassignment.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, warehouseassignment.OrderTable, warehouseassignment.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(waq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkUnit chains the current query on the "work_unit" edge.
func (waq *WarehouseAssignmentQuery) QueryWorkUnit() *WorkUnitInfoQuery {
	query := (&WorkUnitInfoClient{config: waq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := waq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := waq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(warehouseassignment.Table, warehouseassignment.FieldID, selector),
			sqlgraph.To(workunitinfo.Table, workunitinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, warehouseassignment.WorkUnitTable, warehouseassignment.WorkUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(waq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStaff chains the current query on the "staff" edge.
func (waq *WarehouseAssignmentQuery) QueryStaff() *PersonQuery {
	query := (&PersonClient{config: waq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := waq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := waq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(warehouseassignment.Table, warehouseassignment.FieldID, selector),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, warehouseassignment.StaffTable, warehouseassignment.StaffColumn),
		)
		fromU = sqlgraph.SetNeighbors(waq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WarehouseAssignment entity from the query.
// Returns a *NotFoundError when no WarehouseAssignment was found.
func (waq *WarehouseAssignmentQuery) First(ctx context.Context) (*WarehouseAssignment, error) {
	nodes, err := waq.Limit(1).All(setContextOp(ctx, waq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{warehouseassignment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (waq *WarehouseAssignmentQuery) FirstX(ctx context.Context) *WarehouseAssignment {
	node, err := waq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WarehouseAssignment ID from the query.
// Returns a *NotFoundError when no WarehouseAssignment ID was found.
func (waq *WarehouseAssignmentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = waq.Limit(1).IDs(setContextOp(ctx, waq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{warehouseassignment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (waq *WarehouseAssignmentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := waq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WarehouseAssignment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WarehouseAssignment entity is found.
// Returns a *NotFoundError when no WarehouseAssignment entities are found.
func (waq *WarehouseAssignmentQuery) Only(ctx context.Context) (*WarehouseAssignment, error) {
	nodes, err := waq.Limit(2).All(setContextOp(ctx, waq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{warehouseassignment.Label}
	default:
		return nil, &NotSingularError{warehouseassignment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (waq *WarehouseAssignmentQuery) OnlyX(ctx context.Context) *WarehouseAssignment {
	node, err := waq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WarehouseAssignment ID in the query.
// Returns a *NotSingularError when more than one WarehouseAssignment ID is found.
// Returns a *NotFoundError when no entities are found.
func (waq *WarehouseAssignmentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = waq.Limit(2).IDs(setContextOp(ctx, waq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{warehouseassignment.Label}
	default:
		err = &NotSingularError{warehouseassignment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (waq *WarehouseAssignmentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := waq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WarehouseAssignments.
func (waq *WarehouseAssignmentQuery) All(ctx context.Context) ([]*WarehouseAssignment, error) {
	ctx = setContextOp(ctx, waq.ctx, "All")
	if err := waq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WarehouseAssignment, *WarehouseAssignmentQuery]()
	return withInterceptors[[]*WarehouseAssignment](ctx, waq, qr, waq.inters)
}

// AllX is like All, but panics if an error occurs.
func (waq *WarehouseAssignmentQuery) AllX(ctx context.Context) []*WarehouseAssignment {
	nodes, err := waq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WarehouseAssignment IDs.
func (waq *WarehouseAssignmentQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if waq.ctx.Unique == nil && waq.path != nil {
		waq.Unique(true)
	}
	ctx = setContextOp(ctx, waq.ctx, "IDs")
	if err = waq.Select(warehouseassignment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (waq *WarehouseAssignmentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := waq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (waq *WarehouseAssignmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, waq.ctx, "Count")
	if err := waq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, waq, querierCount[*WarehouseAssignmentQuery](), waq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (waq *WarehouseAssignmentQuery) CountX(ctx context.Context) int {
	count, err := waq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (waq *WarehouseAssignmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, waq.ctx, "Exist")
	switch _, err := waq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (waq *WarehouseAssignmentQuery) ExistX(ctx context.Context) bool {
	exist, err := waq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WarehouseAssignmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (waq *WarehouseAssignmentQuery) Clone() *WarehouseAssignmentQuery {
	if waq == nil {
		return nil
	}
	return &WarehouseAssignmentQuery{
		config:       waq.config,
		ctx:          waq.ctx.Clone(),
		order:        append([]warehouseassignment.OrderOption{}, waq.order...),
		inters:       append([]Interceptor{}, waq.inters...),
		predicates:   append([]predicate.WarehouseAssignment{}, waq.predicates...),
		withOrder:    waq.withOrder.Clone(),
		withWorkUnit: waq.withWorkUnit.Clone(),
		withStaff:    waq.withStaff.Clone(),
		// clone intermediate query.
		sql:  waq.sql.Clone(),
		path: waq.path,
	}
}

// WithOrder tells the query-builder to eager-load the nodes that are connected to
// the "order" edge. The optional arguments are used to configure the query builder of the edge.
func (waq *WarehouseAssignmentQuery) WithOrder(opts ...func(*OrderQuery)) *WarehouseAssignmentQuery {
	query := (&OrderClient{config: waq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	waq.withOrder = query
	return waq
}

// WithWorkUnit tells the query-builder to eager-load the nodes that are connected to
// the "work_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (waq *WarehouseAssignmentQuery) WithWorkUnit(opts ...func(*WorkUnitInfoQuery)) *WarehouseAssignmentQuery {
	query := (&WorkUnitInfoClient{config: waq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	waq.withWorkUnit = query
	return waq
}

// WithStaff tells the query-builder to eager-load the nodes that are connected to
// the "staff" edge. The optional arguments are used to configure the query builder of the edge.
func (waq *WarehouseAssignmentQuery) WithStaff(opts ...func(*PersonQuery)) *WarehouseAssignmentQuery {
	query := (&PersonClient{config: waq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	waq.withStaff = query
	return waq
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
//	client.WarehouseAssignment.Query().
//		GroupBy(warehouseassignment.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (waq *WarehouseAssignmentQuery) GroupBy(field string, fields ...string) *WarehouseAssignmentGroupBy {
	waq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WarehouseAssignmentGroupBy{build: waq}
	grbuild.flds = &waq.ctx.Fields
	grbuild.label = warehouseassignment.Label
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
//	client.WarehouseAssignment.Query().
//		Select(warehouseassignment.FieldCreatedAt).
//		Scan(ctx, &v)
func (waq *WarehouseAssignmentQuery) Select(fields ...string) *WarehouseAssignmentSelect {
	waq.ctx.Fields = append(waq.ctx.Fields, fields...)
	sbuild := &WarehouseAssignmentSelect{WarehouseAssignmentQuery: waq}
	sbuild.label = warehouseassignment.Label
	sbuild.flds, sbuild.scan = &waq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WarehouseAssignmentSelect configured with the given aggregations.
func (waq *WarehouseAssignmentQuery) Aggregate(fns ...AggregateFunc) *WarehouseAssignmentSelect {
	return waq.Select().Aggregate(fns...)
}

func (waq *WarehouseAssignmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range waq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, waq); err != nil {
				return err
			}
		}
	}
	for _, f := range waq.ctx.Fields {
		if !warehouseassignment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if waq.path != nil {
		prev, err := waq.path(ctx)
		if err != nil {
			return err
		}
		waq.sql = prev
	}
	return nil
}

func (waq *WarehouseAssignmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WarehouseAssignment, error) {
	var (
		nodes       = []*WarehouseAssignment{}
		_spec       = waq.querySpec()
		loadedTypes = [3]bool{
			waq.withOrder != nil,
			waq.withWorkUnit != nil,
			waq.withStaff != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WarehouseAssignment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WarehouseAssignment{config: waq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, waq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := waq.withOrder; query != nil {
		if err := waq.loadOrder(ctx, query, nodes, nil,
			func(n *WarehouseAssignment, e *Order) { n.Edges.Order = e }); err != nil {
			return nil, err
		}
	}
	if query := waq.withWorkUnit; query != nil {
		if err := waq.loadWorkUnit(ctx, query, nodes, nil,
			func(n *WarehouseAssignment, e *WorkUnitInfo) { n.Edges.WorkUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := waq.withStaff; query != nil {
		if err := waq.loadStaff(ctx, query, nodes, nil,
			func(n *WarehouseAssignment, e *Person) { n.Edges.Staff = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (waq *WarehouseAssignmentQuery) loadOrder(ctx context.Context, query *OrderQuery, nodes []*WarehouseAssignment, init func(*WarehouseAssignment), assign func(*WarehouseAssignment, *Order)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WarehouseAssignment)
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
func (waq *WarehouseAssignmentQuery) loadWorkUnit(ctx context.Context, query *WorkUnitInfoQuery, nodes []*WarehouseAssignment, init func(*WarehouseAssignment), assign func(*WarehouseAssignment, *WorkUnitInfo)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WarehouseAssignment)
	for i := range nodes {
		fk := nodes[i].WorkUnitID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workunitinfo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "work_unit_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (waq *WarehouseAssignmentQuery) loadStaff(ctx context.Context, query *PersonQuery, nodes []*WarehouseAssignment, init func(*WarehouseAssignment), assign func(*WarehouseAssignment, *Person)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WarehouseAssignment)
	for i := range nodes {
		if nodes[i].StaffID == nil {
			continue
		}
		fk := *nodes[i].StaffID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(person.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "staff_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (waq *WarehouseAssignmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := waq.querySpec()
	_spec.Node.Columns = waq.ctx.Fields
	if len(waq.ctx.Fields) > 0 {
		_spec.Unique = waq.ctx.Unique != nil && *waq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, waq.driver, _spec)
}

func (waq *WarehouseAssignmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(warehouseassignment.Table, warehouseassignment.Columns, sqlgraph.NewFieldSpec(warehouseassignment.FieldID, field.TypeUUID))
	_spec.From = waq.sql
	if unique := waq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if waq.path != nil {
		_spec.Unique = true
	}
	if fields := waq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, warehouseassignment.FieldID)
		for i := range fields {
			if fields[i] != warehouseassignment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if waq.withOrder != nil {
			_spec.Node.AddColumnOnce(warehouseassignment.FieldOrderID)
		}
		if waq.withWorkUnit != nil {
			_spec.Node.AddColumnOnce(warehouseassignment.FieldWorkUnitID)
		}
		if waq.withStaff != nil {
			_spec.Node.AddColumnOnce(warehouseassignment.FieldStaffID)
		}
	}
	if ps := waq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := waq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := waq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := waq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (waq *WarehouseAssignmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(waq.driver.Dialect())
	t1 := builder.Table(warehouseassignment.Table)
	columns := waq.ctx.Fields
	if len(columns) == 0 {
		columns = warehouseassignment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if waq.sql != nil {
		selector = waq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if waq.ctx.Unique != nil && *waq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range waq.predicates {
		p(selector)
	}
	for _, p := range waq.order {
		p(selector)
	}
	if offset := waq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := waq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WarehouseAssignmentGroupBy is the group-by builder for WarehouseAssignment entities.
type WarehouseAssignmentGroupBy struct {
	selector
	build *WarehouseAssignmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wagb *WarehouseAssignmentGroupBy) Aggregate(fns ...AggregateFunc) *WarehouseAssignmentGroupBy {
	wagb.fns = append(wagb.fns, fns...)
	return wagb
}

// Scan applies the selector query and scans the result into the given value.
func (wagb *WarehouseAssignmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wagb.build.ctx, "GroupBy")
	if err := wagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WarehouseAssignmentQuery, *WarehouseAssignmentGroupBy](ctx, wagb.build, wagb, wagb.build.inters, v)
}

func (wagb *WarehouseAssignmentGroupBy) sqlScan(ctx context.Context, root *WarehouseAssignmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wagb.fns))
	for _, fn := range wagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wagb.flds)+len(wagb.fns))
		for _, f := range *wagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WarehouseAssignmentSelect is the builder for selecting fields of WarehouseAssignment entities.
type WarehouseAssignmentSelect struct {
	*WarehouseAssignmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (was *WarehouseAssignmentSelect) Aggregate(fns ...AggregateFunc) *WarehouseAssignmentSelect {
	was.fns = append(was.fns, fns...)
	return was
}

// Scan applies the selector query and scans the result into the given value.
func (was *WarehouseAssignmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, was.ctx, "Select")
	if err := was.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WarehouseAssignmentQuery, *WarehouseAssignmentSelect](ctx, was.WarehouseAssignmentQuery, was, was.inters, v)
}

func (was *WarehouseAssignmentSelect) sqlScan(ctx context.Context, root *WarehouseAssignmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(was.fns))
	for _, fn := range was.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*was.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := was.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
