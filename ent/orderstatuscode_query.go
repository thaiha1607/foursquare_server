// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/orderstatuscode"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// OrderStatusCodeQuery is the builder for querying OrderStatusCode entities.
type OrderStatusCodeQuery struct {
	config
	ctx        *QueryContext
	order      []orderstatuscode.OrderOption
	inters     []Interceptor
	predicates []predicate.OrderStatusCode
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderStatusCodeQuery builder.
func (oscq *OrderStatusCodeQuery) Where(ps ...predicate.OrderStatusCode) *OrderStatusCodeQuery {
	oscq.predicates = append(oscq.predicates, ps...)
	return oscq
}

// Limit the number of records to be returned by this query.
func (oscq *OrderStatusCodeQuery) Limit(limit int) *OrderStatusCodeQuery {
	oscq.ctx.Limit = &limit
	return oscq
}

// Offset to start from.
func (oscq *OrderStatusCodeQuery) Offset(offset int) *OrderStatusCodeQuery {
	oscq.ctx.Offset = &offset
	return oscq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oscq *OrderStatusCodeQuery) Unique(unique bool) *OrderStatusCodeQuery {
	oscq.ctx.Unique = &unique
	return oscq
}

// Order specifies how the records should be ordered.
func (oscq *OrderStatusCodeQuery) Order(o ...orderstatuscode.OrderOption) *OrderStatusCodeQuery {
	oscq.order = append(oscq.order, o...)
	return oscq
}

// First returns the first OrderStatusCode entity from the query.
// Returns a *NotFoundError when no OrderStatusCode was found.
func (oscq *OrderStatusCodeQuery) First(ctx context.Context) (*OrderStatusCode, error) {
	nodes, err := oscq.Limit(1).All(setContextOp(ctx, oscq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderstatuscode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oscq *OrderStatusCodeQuery) FirstX(ctx context.Context) *OrderStatusCode {
	node, err := oscq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderStatusCode ID from the query.
// Returns a *NotFoundError when no OrderStatusCode ID was found.
func (oscq *OrderStatusCodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oscq.Limit(1).IDs(setContextOp(ctx, oscq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderstatuscode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oscq *OrderStatusCodeQuery) FirstIDX(ctx context.Context) int {
	id, err := oscq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderStatusCode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderStatusCode entity is found.
// Returns a *NotFoundError when no OrderStatusCode entities are found.
func (oscq *OrderStatusCodeQuery) Only(ctx context.Context) (*OrderStatusCode, error) {
	nodes, err := oscq.Limit(2).All(setContextOp(ctx, oscq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderstatuscode.Label}
	default:
		return nil, &NotSingularError{orderstatuscode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oscq *OrderStatusCodeQuery) OnlyX(ctx context.Context) *OrderStatusCode {
	node, err := oscq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderStatusCode ID in the query.
// Returns a *NotSingularError when more than one OrderStatusCode ID is found.
// Returns a *NotFoundError when no entities are found.
func (oscq *OrderStatusCodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oscq.Limit(2).IDs(setContextOp(ctx, oscq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderstatuscode.Label}
	default:
		err = &NotSingularError{orderstatuscode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oscq *OrderStatusCodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := oscq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderStatusCodes.
func (oscq *OrderStatusCodeQuery) All(ctx context.Context) ([]*OrderStatusCode, error) {
	ctx = setContextOp(ctx, oscq.ctx, "All")
	if err := oscq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderStatusCode, *OrderStatusCodeQuery]()
	return withInterceptors[[]*OrderStatusCode](ctx, oscq, qr, oscq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oscq *OrderStatusCodeQuery) AllX(ctx context.Context) []*OrderStatusCode {
	nodes, err := oscq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderStatusCode IDs.
func (oscq *OrderStatusCodeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if oscq.ctx.Unique == nil && oscq.path != nil {
		oscq.Unique(true)
	}
	ctx = setContextOp(ctx, oscq.ctx, "IDs")
	if err = oscq.Select(orderstatuscode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oscq *OrderStatusCodeQuery) IDsX(ctx context.Context) []int {
	ids, err := oscq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oscq *OrderStatusCodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oscq.ctx, "Count")
	if err := oscq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oscq, querierCount[*OrderStatusCodeQuery](), oscq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oscq *OrderStatusCodeQuery) CountX(ctx context.Context) int {
	count, err := oscq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oscq *OrderStatusCodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oscq.ctx, "Exist")
	switch _, err := oscq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oscq *OrderStatusCodeQuery) ExistX(ctx context.Context) bool {
	exist, err := oscq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderStatusCodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oscq *OrderStatusCodeQuery) Clone() *OrderStatusCodeQuery {
	if oscq == nil {
		return nil
	}
	return &OrderStatusCodeQuery{
		config:     oscq.config,
		ctx:        oscq.ctx.Clone(),
		order:      append([]orderstatuscode.OrderOption{}, oscq.order...),
		inters:     append([]Interceptor{}, oscq.inters...),
		predicates: append([]predicate.OrderStatusCode{}, oscq.predicates...),
		// clone intermediate query.
		sql:  oscq.sql.Clone(),
		path: oscq.path,
	}
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
//	client.OrderStatusCode.Query().
//		GroupBy(orderstatuscode.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oscq *OrderStatusCodeQuery) GroupBy(field string, fields ...string) *OrderStatusCodeGroupBy {
	oscq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderStatusCodeGroupBy{build: oscq}
	grbuild.flds = &oscq.ctx.Fields
	grbuild.label = orderstatuscode.Label
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
//	client.OrderStatusCode.Query().
//		Select(orderstatuscode.FieldCreatedAt).
//		Scan(ctx, &v)
func (oscq *OrderStatusCodeQuery) Select(fields ...string) *OrderStatusCodeSelect {
	oscq.ctx.Fields = append(oscq.ctx.Fields, fields...)
	sbuild := &OrderStatusCodeSelect{OrderStatusCodeQuery: oscq}
	sbuild.label = orderstatuscode.Label
	sbuild.flds, sbuild.scan = &oscq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderStatusCodeSelect configured with the given aggregations.
func (oscq *OrderStatusCodeQuery) Aggregate(fns ...AggregateFunc) *OrderStatusCodeSelect {
	return oscq.Select().Aggregate(fns...)
}

func (oscq *OrderStatusCodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oscq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oscq); err != nil {
				return err
			}
		}
	}
	for _, f := range oscq.ctx.Fields {
		if !orderstatuscode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oscq.path != nil {
		prev, err := oscq.path(ctx)
		if err != nil {
			return err
		}
		oscq.sql = prev
	}
	return nil
}

func (oscq *OrderStatusCodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderStatusCode, error) {
	var (
		nodes = []*OrderStatusCode{}
		_spec = oscq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderStatusCode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderStatusCode{config: oscq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oscq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oscq *OrderStatusCodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oscq.querySpec()
	_spec.Node.Columns = oscq.ctx.Fields
	if len(oscq.ctx.Fields) > 0 {
		_spec.Unique = oscq.ctx.Unique != nil && *oscq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oscq.driver, _spec)
}

func (oscq *OrderStatusCodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orderstatuscode.Table, orderstatuscode.Columns, sqlgraph.NewFieldSpec(orderstatuscode.FieldID, field.TypeInt))
	_spec.From = oscq.sql
	if unique := oscq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oscq.path != nil {
		_spec.Unique = true
	}
	if fields := oscq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderstatuscode.FieldID)
		for i := range fields {
			if fields[i] != orderstatuscode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oscq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oscq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oscq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oscq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oscq *OrderStatusCodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oscq.driver.Dialect())
	t1 := builder.Table(orderstatuscode.Table)
	columns := oscq.ctx.Fields
	if len(columns) == 0 {
		columns = orderstatuscode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oscq.sql != nil {
		selector = oscq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oscq.ctx.Unique != nil && *oscq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oscq.predicates {
		p(selector)
	}
	for _, p := range oscq.order {
		p(selector)
	}
	if offset := oscq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oscq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderStatusCodeGroupBy is the group-by builder for OrderStatusCode entities.
type OrderStatusCodeGroupBy struct {
	selector
	build *OrderStatusCodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oscgb *OrderStatusCodeGroupBy) Aggregate(fns ...AggregateFunc) *OrderStatusCodeGroupBy {
	oscgb.fns = append(oscgb.fns, fns...)
	return oscgb
}

// Scan applies the selector query and scans the result into the given value.
func (oscgb *OrderStatusCodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oscgb.build.ctx, "GroupBy")
	if err := oscgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderStatusCodeQuery, *OrderStatusCodeGroupBy](ctx, oscgb.build, oscgb, oscgb.build.inters, v)
}

func (oscgb *OrderStatusCodeGroupBy) sqlScan(ctx context.Context, root *OrderStatusCodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(oscgb.fns))
	for _, fn := range oscgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*oscgb.flds)+len(oscgb.fns))
		for _, f := range *oscgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*oscgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oscgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderStatusCodeSelect is the builder for selecting fields of OrderStatusCode entities.
type OrderStatusCodeSelect struct {
	*OrderStatusCodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oscs *OrderStatusCodeSelect) Aggregate(fns ...AggregateFunc) *OrderStatusCodeSelect {
	oscs.fns = append(oscs.fns, fns...)
	return oscs
}

// Scan applies the selector query and scans the result into the given value.
func (oscs *OrderStatusCodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oscs.ctx, "Select")
	if err := oscs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderStatusCodeQuery, *OrderStatusCodeSelect](ctx, oscs.OrderStatusCodeQuery, oscs, oscs.inters, v)
}

func (oscs *OrderStatusCodeSelect) sqlScan(ctx context.Context, root *OrderStatusCodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oscs.fns))
	for _, fn := range oscs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oscs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oscs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
