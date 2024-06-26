// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/productinfo"
	"github.com/thaiha1607/foursquare_server/ent/producttag"
	"github.com/thaiha1607/foursquare_server/ent/tag"
)

// ProductInfoQuery is the builder for querying ProductInfo entities.
type ProductInfoQuery struct {
	config
	ctx             *QueryContext
	order           []productinfo.OrderOption
	inters          []Interceptor
	predicates      []predicate.ProductInfo
	withTags        *TagQuery
	withProductTags *ProductTagQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductInfoQuery builder.
func (piq *ProductInfoQuery) Where(ps ...predicate.ProductInfo) *ProductInfoQuery {
	piq.predicates = append(piq.predicates, ps...)
	return piq
}

// Limit the number of records to be returned by this query.
func (piq *ProductInfoQuery) Limit(limit int) *ProductInfoQuery {
	piq.ctx.Limit = &limit
	return piq
}

// Offset to start from.
func (piq *ProductInfoQuery) Offset(offset int) *ProductInfoQuery {
	piq.ctx.Offset = &offset
	return piq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piq *ProductInfoQuery) Unique(unique bool) *ProductInfoQuery {
	piq.ctx.Unique = &unique
	return piq
}

// Order specifies how the records should be ordered.
func (piq *ProductInfoQuery) Order(o ...productinfo.OrderOption) *ProductInfoQuery {
	piq.order = append(piq.order, o...)
	return piq
}

// QueryTags chains the current query on the "tags" edge.
func (piq *ProductInfoQuery) QueryTags() *TagQuery {
	query := (&TagClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productinfo.Table, productinfo.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, productinfo.TagsTable, productinfo.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProductTags chains the current query on the "product_tags" edge.
func (piq *ProductInfoQuery) QueryProductTags() *ProductTagQuery {
	query := (&ProductTagClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productinfo.Table, productinfo.FieldID, selector),
			sqlgraph.To(producttag.Table, producttag.ProductsColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, productinfo.ProductTagsTable, productinfo.ProductTagsColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductInfo entity from the query.
// Returns a *NotFoundError when no ProductInfo was found.
func (piq *ProductInfoQuery) First(ctx context.Context) (*ProductInfo, error) {
	nodes, err := piq.Limit(1).All(setContextOp(ctx, piq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piq *ProductInfoQuery) FirstX(ctx context.Context) *ProductInfo {
	node, err := piq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductInfo ID from the query.
// Returns a *NotFoundError when no ProductInfo ID was found.
func (piq *ProductInfoQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = piq.Limit(1).IDs(setContextOp(ctx, piq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piq *ProductInfoQuery) FirstIDX(ctx context.Context) string {
	id, err := piq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProductInfo entity is found.
// Returns a *NotFoundError when no ProductInfo entities are found.
func (piq *ProductInfoQuery) Only(ctx context.Context) (*ProductInfo, error) {
	nodes, err := piq.Limit(2).All(setContextOp(ctx, piq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productinfo.Label}
	default:
		return nil, &NotSingularError{productinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piq *ProductInfoQuery) OnlyX(ctx context.Context) *ProductInfo {
	node, err := piq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductInfo ID in the query.
// Returns a *NotSingularError when more than one ProductInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (piq *ProductInfoQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = piq.Limit(2).IDs(setContextOp(ctx, piq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productinfo.Label}
	default:
		err = &NotSingularError{productinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piq *ProductInfoQuery) OnlyIDX(ctx context.Context) string {
	id, err := piq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductInfos.
func (piq *ProductInfoQuery) All(ctx context.Context) ([]*ProductInfo, error) {
	ctx = setContextOp(ctx, piq.ctx, "All")
	if err := piq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProductInfo, *ProductInfoQuery]()
	return withInterceptors[[]*ProductInfo](ctx, piq, qr, piq.inters)
}

// AllX is like All, but panics if an error occurs.
func (piq *ProductInfoQuery) AllX(ctx context.Context) []*ProductInfo {
	nodes, err := piq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductInfo IDs.
func (piq *ProductInfoQuery) IDs(ctx context.Context) (ids []string, err error) {
	if piq.ctx.Unique == nil && piq.path != nil {
		piq.Unique(true)
	}
	ctx = setContextOp(ctx, piq.ctx, "IDs")
	if err = piq.Select(productinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piq *ProductInfoQuery) IDsX(ctx context.Context) []string {
	ids, err := piq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piq *ProductInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, piq.ctx, "Count")
	if err := piq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, piq, querierCount[*ProductInfoQuery](), piq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (piq *ProductInfoQuery) CountX(ctx context.Context) int {
	count, err := piq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piq *ProductInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, piq.ctx, "Exist")
	switch _, err := piq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (piq *ProductInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := piq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piq *ProductInfoQuery) Clone() *ProductInfoQuery {
	if piq == nil {
		return nil
	}
	return &ProductInfoQuery{
		config:          piq.config,
		ctx:             piq.ctx.Clone(),
		order:           append([]productinfo.OrderOption{}, piq.order...),
		inters:          append([]Interceptor{}, piq.inters...),
		predicates:      append([]predicate.ProductInfo{}, piq.predicates...),
		withTags:        piq.withTags.Clone(),
		withProductTags: piq.withProductTags.Clone(),
		// clone intermediate query.
		sql:  piq.sql.Clone(),
		path: piq.path,
	}
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *ProductInfoQuery) WithTags(opts ...func(*TagQuery)) *ProductInfoQuery {
	query := (&TagClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withTags = query
	return piq
}

// WithProductTags tells the query-builder to eager-load the nodes that are connected to
// the "product_tags" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *ProductInfoQuery) WithProductTags(opts ...func(*ProductTagQuery)) *ProductInfoQuery {
	query := (&ProductTagClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withProductTags = query
	return piq
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
//	client.ProductInfo.Query().
//		GroupBy(productinfo.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (piq *ProductInfoQuery) GroupBy(field string, fields ...string) *ProductInfoGroupBy {
	piq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProductInfoGroupBy{build: piq}
	grbuild.flds = &piq.ctx.Fields
	grbuild.label = productinfo.Label
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
//	client.ProductInfo.Query().
//		Select(productinfo.FieldCreatedAt).
//		Scan(ctx, &v)
func (piq *ProductInfoQuery) Select(fields ...string) *ProductInfoSelect {
	piq.ctx.Fields = append(piq.ctx.Fields, fields...)
	sbuild := &ProductInfoSelect{ProductInfoQuery: piq}
	sbuild.label = productinfo.Label
	sbuild.flds, sbuild.scan = &piq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProductInfoSelect configured with the given aggregations.
func (piq *ProductInfoQuery) Aggregate(fns ...AggregateFunc) *ProductInfoSelect {
	return piq.Select().Aggregate(fns...)
}

func (piq *ProductInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range piq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, piq); err != nil {
				return err
			}
		}
	}
	for _, f := range piq.ctx.Fields {
		if !productinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if piq.path != nil {
		prev, err := piq.path(ctx)
		if err != nil {
			return err
		}
		piq.sql = prev
	}
	return nil
}

func (piq *ProductInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProductInfo, error) {
	var (
		nodes       = []*ProductInfo{}
		_spec       = piq.querySpec()
		loadedTypes = [2]bool{
			piq.withTags != nil,
			piq.withProductTags != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProductInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProductInfo{config: piq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, piq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := piq.withTags; query != nil {
		if err := piq.loadTags(ctx, query, nodes,
			func(n *ProductInfo) { n.Edges.Tags = []*Tag{} },
			func(n *ProductInfo, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	if query := piq.withProductTags; query != nil {
		if err := piq.loadProductTags(ctx, query, nodes,
			func(n *ProductInfo) { n.Edges.ProductTags = []*ProductTag{} },
			func(n *ProductInfo, e *ProductTag) { n.Edges.ProductTags = append(n.Edges.ProductTags, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (piq *ProductInfoQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*ProductInfo, init func(*ProductInfo), assign func(*ProductInfo, *Tag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*ProductInfo)
	nids := make(map[string]map[*ProductInfo]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(productinfo.TagsTable)
		s.Join(joinT).On(s.C(tag.FieldID), joinT.C(productinfo.TagsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(productinfo.TagsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(productinfo.TagsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*ProductInfo]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tag](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (piq *ProductInfoQuery) loadProductTags(ctx context.Context, query *ProductTagQuery, nodes []*ProductInfo, init func(*ProductInfo), assign func(*ProductInfo, *ProductTag)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*ProductInfo)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(producttag.FieldProductID)
	}
	query.Where(predicate.ProductTag(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(productinfo.ProductTagsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProductID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "product_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (piq *ProductInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piq.querySpec()
	_spec.Node.Columns = piq.ctx.Fields
	if len(piq.ctx.Fields) > 0 {
		_spec.Unique = piq.ctx.Unique != nil && *piq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, piq.driver, _spec)
}

func (piq *ProductInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(productinfo.Table, productinfo.Columns, sqlgraph.NewFieldSpec(productinfo.FieldID, field.TypeString))
	_spec.From = piq.sql
	if unique := piq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if piq.path != nil {
		_spec.Unique = true
	}
	if fields := piq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productinfo.FieldID)
		for i := range fields {
			if fields[i] != productinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := piq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piq *ProductInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piq.driver.Dialect())
	t1 := builder.Table(productinfo.Table)
	columns := piq.ctx.Fields
	if len(columns) == 0 {
		columns = productinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if piq.sql != nil {
		selector = piq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if piq.ctx.Unique != nil && *piq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range piq.predicates {
		p(selector)
	}
	for _, p := range piq.order {
		p(selector)
	}
	if offset := piq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductInfoGroupBy is the group-by builder for ProductInfo entities.
type ProductInfoGroupBy struct {
	selector
	build *ProductInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pigb *ProductInfoGroupBy) Aggregate(fns ...AggregateFunc) *ProductInfoGroupBy {
	pigb.fns = append(pigb.fns, fns...)
	return pigb
}

// Scan applies the selector query and scans the result into the given value.
func (pigb *ProductInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pigb.build.ctx, "GroupBy")
	if err := pigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductInfoQuery, *ProductInfoGroupBy](ctx, pigb.build, pigb, pigb.build.inters, v)
}

func (pigb *ProductInfoGroupBy) sqlScan(ctx context.Context, root *ProductInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pigb.fns))
	for _, fn := range pigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pigb.flds)+len(pigb.fns))
		for _, f := range *pigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProductInfoSelect is the builder for selecting fields of ProductInfo entities.
type ProductInfoSelect struct {
	*ProductInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pis *ProductInfoSelect) Aggregate(fns ...AggregateFunc) *ProductInfoSelect {
	pis.fns = append(pis.fns, fns...)
	return pis
}

// Scan applies the selector query and scans the result into the given value.
func (pis *ProductInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pis.ctx, "Select")
	if err := pis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductInfoQuery, *ProductInfoSelect](ctx, pis.ProductInfoQuery, pis, pis.inters, v)
}

func (pis *ProductInfoSelect) sqlScan(ctx context.Context, root *ProductInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pis.fns))
	for _, fn := range pis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
