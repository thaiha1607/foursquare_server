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
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/productcolor"
	"github.com/thaiha1607/foursquare_server/ent/productinfo"
	"github.com/thaiha1607/foursquare_server/ent/productqty"
	"github.com/thaiha1607/foursquare_server/ent/workunitinfo"
)

// ProductQtyQuery is the builder for querying ProductQty entities.
type ProductQtyQuery struct {
	config
	ctx              *QueryContext
	order            []productqty.OrderOption
	inters           []Interceptor
	predicates       []predicate.ProductQty
	withWorkUnit     *WorkUnitInfoQuery
	withProduct      *ProductInfoQuery
	withProductColor *ProductColorQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductQtyQuery builder.
func (pqq *ProductQtyQuery) Where(ps ...predicate.ProductQty) *ProductQtyQuery {
	pqq.predicates = append(pqq.predicates, ps...)
	return pqq
}

// Limit the number of records to be returned by this query.
func (pqq *ProductQtyQuery) Limit(limit int) *ProductQtyQuery {
	pqq.ctx.Limit = &limit
	return pqq
}

// Offset to start from.
func (pqq *ProductQtyQuery) Offset(offset int) *ProductQtyQuery {
	pqq.ctx.Offset = &offset
	return pqq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pqq *ProductQtyQuery) Unique(unique bool) *ProductQtyQuery {
	pqq.ctx.Unique = &unique
	return pqq
}

// Order specifies how the records should be ordered.
func (pqq *ProductQtyQuery) Order(o ...productqty.OrderOption) *ProductQtyQuery {
	pqq.order = append(pqq.order, o...)
	return pqq
}

// QueryWorkUnit chains the current query on the "work_unit" edge.
func (pqq *ProductQtyQuery) QueryWorkUnit() *WorkUnitInfoQuery {
	query := (&WorkUnitInfoClient{config: pqq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pqq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productqty.Table, productqty.FieldID, selector),
			sqlgraph.To(workunitinfo.Table, workunitinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, productqty.WorkUnitTable, productqty.WorkUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(pqq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the "product" edge.
func (pqq *ProductQtyQuery) QueryProduct() *ProductInfoQuery {
	query := (&ProductInfoClient{config: pqq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pqq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productqty.Table, productqty.FieldID, selector),
			sqlgraph.To(productinfo.Table, productinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, productqty.ProductTable, productqty.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(pqq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProductColor chains the current query on the "product_color" edge.
func (pqq *ProductQtyQuery) QueryProductColor() *ProductColorQuery {
	query := (&ProductColorClient{config: pqq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pqq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productqty.Table, productqty.FieldID, selector),
			sqlgraph.To(productcolor.Table, productcolor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, productqty.ProductColorTable, productqty.ProductColorColumn),
		)
		fromU = sqlgraph.SetNeighbors(pqq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductQty entity from the query.
// Returns a *NotFoundError when no ProductQty was found.
func (pqq *ProductQtyQuery) First(ctx context.Context) (*ProductQty, error) {
	nodes, err := pqq.Limit(1).All(setContextOp(ctx, pqq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productqty.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pqq *ProductQtyQuery) FirstX(ctx context.Context) *ProductQty {
	node, err := pqq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductQty ID from the query.
// Returns a *NotFoundError when no ProductQty ID was found.
func (pqq *ProductQtyQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pqq.Limit(1).IDs(setContextOp(ctx, pqq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productqty.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pqq *ProductQtyQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pqq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductQty entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProductQty entity is found.
// Returns a *NotFoundError when no ProductQty entities are found.
func (pqq *ProductQtyQuery) Only(ctx context.Context) (*ProductQty, error) {
	nodes, err := pqq.Limit(2).All(setContextOp(ctx, pqq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productqty.Label}
	default:
		return nil, &NotSingularError{productqty.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pqq *ProductQtyQuery) OnlyX(ctx context.Context) *ProductQty {
	node, err := pqq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductQty ID in the query.
// Returns a *NotSingularError when more than one ProductQty ID is found.
// Returns a *NotFoundError when no entities are found.
func (pqq *ProductQtyQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pqq.Limit(2).IDs(setContextOp(ctx, pqq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productqty.Label}
	default:
		err = &NotSingularError{productqty.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pqq *ProductQtyQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pqq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductQties.
func (pqq *ProductQtyQuery) All(ctx context.Context) ([]*ProductQty, error) {
	ctx = setContextOp(ctx, pqq.ctx, "All")
	if err := pqq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProductQty, *ProductQtyQuery]()
	return withInterceptors[[]*ProductQty](ctx, pqq, qr, pqq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pqq *ProductQtyQuery) AllX(ctx context.Context) []*ProductQty {
	nodes, err := pqq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductQty IDs.
func (pqq *ProductQtyQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pqq.ctx.Unique == nil && pqq.path != nil {
		pqq.Unique(true)
	}
	ctx = setContextOp(ctx, pqq.ctx, "IDs")
	if err = pqq.Select(productqty.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pqq *ProductQtyQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pqq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pqq *ProductQtyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pqq.ctx, "Count")
	if err := pqq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pqq, querierCount[*ProductQtyQuery](), pqq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pqq *ProductQtyQuery) CountX(ctx context.Context) int {
	count, err := pqq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pqq *ProductQtyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pqq.ctx, "Exist")
	switch _, err := pqq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pqq *ProductQtyQuery) ExistX(ctx context.Context) bool {
	exist, err := pqq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductQtyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pqq *ProductQtyQuery) Clone() *ProductQtyQuery {
	if pqq == nil {
		return nil
	}
	return &ProductQtyQuery{
		config:           pqq.config,
		ctx:              pqq.ctx.Clone(),
		order:            append([]productqty.OrderOption{}, pqq.order...),
		inters:           append([]Interceptor{}, pqq.inters...),
		predicates:       append([]predicate.ProductQty{}, pqq.predicates...),
		withWorkUnit:     pqq.withWorkUnit.Clone(),
		withProduct:      pqq.withProduct.Clone(),
		withProductColor: pqq.withProductColor.Clone(),
		// clone intermediate query.
		sql:  pqq.sql.Clone(),
		path: pqq.path,
	}
}

// WithWorkUnit tells the query-builder to eager-load the nodes that are connected to
// the "work_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (pqq *ProductQtyQuery) WithWorkUnit(opts ...func(*WorkUnitInfoQuery)) *ProductQtyQuery {
	query := (&WorkUnitInfoClient{config: pqq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pqq.withWorkUnit = query
	return pqq
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (pqq *ProductQtyQuery) WithProduct(opts ...func(*ProductInfoQuery)) *ProductQtyQuery {
	query := (&ProductInfoClient{config: pqq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pqq.withProduct = query
	return pqq
}

// WithProductColor tells the query-builder to eager-load the nodes that are connected to
// the "product_color" edge. The optional arguments are used to configure the query builder of the edge.
func (pqq *ProductQtyQuery) WithProductColor(opts ...func(*ProductColorQuery)) *ProductQtyQuery {
	query := (&ProductColorClient{config: pqq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pqq.withProductColor = query
	return pqq
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
//	client.ProductQty.Query().
//		GroupBy(productqty.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pqq *ProductQtyQuery) GroupBy(field string, fields ...string) *ProductQtyGroupBy {
	pqq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProductQtyGroupBy{build: pqq}
	grbuild.flds = &pqq.ctx.Fields
	grbuild.label = productqty.Label
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
//	client.ProductQty.Query().
//		Select(productqty.FieldCreatedAt).
//		Scan(ctx, &v)
func (pqq *ProductQtyQuery) Select(fields ...string) *ProductQtySelect {
	pqq.ctx.Fields = append(pqq.ctx.Fields, fields...)
	sbuild := &ProductQtySelect{ProductQtyQuery: pqq}
	sbuild.label = productqty.Label
	sbuild.flds, sbuild.scan = &pqq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProductQtySelect configured with the given aggregations.
func (pqq *ProductQtyQuery) Aggregate(fns ...AggregateFunc) *ProductQtySelect {
	return pqq.Select().Aggregate(fns...)
}

func (pqq *ProductQtyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pqq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pqq); err != nil {
				return err
			}
		}
	}
	for _, f := range pqq.ctx.Fields {
		if !productqty.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pqq.path != nil {
		prev, err := pqq.path(ctx)
		if err != nil {
			return err
		}
		pqq.sql = prev
	}
	return nil
}

func (pqq *ProductQtyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProductQty, error) {
	var (
		nodes       = []*ProductQty{}
		_spec       = pqq.querySpec()
		loadedTypes = [3]bool{
			pqq.withWorkUnit != nil,
			pqq.withProduct != nil,
			pqq.withProductColor != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProductQty).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProductQty{config: pqq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pqq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pqq.withWorkUnit; query != nil {
		if err := pqq.loadWorkUnit(ctx, query, nodes, nil,
			func(n *ProductQty, e *WorkUnitInfo) { n.Edges.WorkUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := pqq.withProduct; query != nil {
		if err := pqq.loadProduct(ctx, query, nodes, nil,
			func(n *ProductQty, e *ProductInfo) { n.Edges.Product = e }); err != nil {
			return nil, err
		}
	}
	if query := pqq.withProductColor; query != nil {
		if err := pqq.loadProductColor(ctx, query, nodes, nil,
			func(n *ProductQty, e *ProductColor) { n.Edges.ProductColor = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pqq *ProductQtyQuery) loadWorkUnit(ctx context.Context, query *WorkUnitInfoQuery, nodes []*ProductQty, init func(*ProductQty), assign func(*ProductQty, *WorkUnitInfo)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ProductQty)
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
func (pqq *ProductQtyQuery) loadProduct(ctx context.Context, query *ProductInfoQuery, nodes []*ProductQty, init func(*ProductQty), assign func(*ProductQty, *ProductInfo)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ProductQty)
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
	query.Where(productinfo.IDIn(ids...))
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
func (pqq *ProductQtyQuery) loadProductColor(ctx context.Context, query *ProductColorQuery, nodes []*ProductQty, init func(*ProductQty), assign func(*ProductQty, *ProductColor)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ProductQty)
	for i := range nodes {
		fk := nodes[i].ProductColorID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(productcolor.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "product_color_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pqq *ProductQtyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pqq.querySpec()
	_spec.Node.Columns = pqq.ctx.Fields
	if len(pqq.ctx.Fields) > 0 {
		_spec.Unique = pqq.ctx.Unique != nil && *pqq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pqq.driver, _spec)
}

func (pqq *ProductQtyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(productqty.Table, productqty.Columns, sqlgraph.NewFieldSpec(productqty.FieldID, field.TypeUUID))
	_spec.From = pqq.sql
	if unique := pqq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pqq.path != nil {
		_spec.Unique = true
	}
	if fields := pqq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productqty.FieldID)
		for i := range fields {
			if fields[i] != productqty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pqq.withWorkUnit != nil {
			_spec.Node.AddColumnOnce(productqty.FieldWorkUnitID)
		}
		if pqq.withProduct != nil {
			_spec.Node.AddColumnOnce(productqty.FieldProductID)
		}
		if pqq.withProductColor != nil {
			_spec.Node.AddColumnOnce(productqty.FieldProductColorID)
		}
	}
	if ps := pqq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pqq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pqq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pqq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pqq *ProductQtyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pqq.driver.Dialect())
	t1 := builder.Table(productqty.Table)
	columns := pqq.ctx.Fields
	if len(columns) == 0 {
		columns = productqty.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pqq.sql != nil {
		selector = pqq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pqq.ctx.Unique != nil && *pqq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pqq.predicates {
		p(selector)
	}
	for _, p := range pqq.order {
		p(selector)
	}
	if offset := pqq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pqq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductQtyGroupBy is the group-by builder for ProductQty entities.
type ProductQtyGroupBy struct {
	selector
	build *ProductQtyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pqgb *ProductQtyGroupBy) Aggregate(fns ...AggregateFunc) *ProductQtyGroupBy {
	pqgb.fns = append(pqgb.fns, fns...)
	return pqgb
}

// Scan applies the selector query and scans the result into the given value.
func (pqgb *ProductQtyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pqgb.build.ctx, "GroupBy")
	if err := pqgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductQtyQuery, *ProductQtyGroupBy](ctx, pqgb.build, pqgb, pqgb.build.inters, v)
}

func (pqgb *ProductQtyGroupBy) sqlScan(ctx context.Context, root *ProductQtyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pqgb.fns))
	for _, fn := range pqgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pqgb.flds)+len(pqgb.fns))
		for _, f := range *pqgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pqgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pqgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProductQtySelect is the builder for selecting fields of ProductQty entities.
type ProductQtySelect struct {
	*ProductQtyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pqs *ProductQtySelect) Aggregate(fns ...AggregateFunc) *ProductQtySelect {
	pqs.fns = append(pqs.fns, fns...)
	return pqs
}

// Scan applies the selector query and scans the result into the given value.
func (pqs *ProductQtySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pqs.ctx, "Select")
	if err := pqs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductQtyQuery, *ProductQtySelect](ctx, pqs.ProductQtyQuery, pqs, pqs.inters, v)
}

func (pqs *ProductQtySelect) sqlScan(ctx context.Context, root *ProductQtyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pqs.fns))
	for _, fn := range pqs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pqs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pqs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
