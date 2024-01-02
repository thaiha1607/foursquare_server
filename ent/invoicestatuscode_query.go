// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thaiha1607/foursquare_server/ent/invoicestatuscode"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// InvoiceStatusCodeQuery is the builder for querying InvoiceStatusCode entities.
type InvoiceStatusCodeQuery struct {
	config
	ctx        *QueryContext
	order      []invoicestatuscode.OrderOption
	inters     []Interceptor
	predicates []predicate.InvoiceStatusCode
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InvoiceStatusCodeQuery builder.
func (iscq *InvoiceStatusCodeQuery) Where(ps ...predicate.InvoiceStatusCode) *InvoiceStatusCodeQuery {
	iscq.predicates = append(iscq.predicates, ps...)
	return iscq
}

// Limit the number of records to be returned by this query.
func (iscq *InvoiceStatusCodeQuery) Limit(limit int) *InvoiceStatusCodeQuery {
	iscq.ctx.Limit = &limit
	return iscq
}

// Offset to start from.
func (iscq *InvoiceStatusCodeQuery) Offset(offset int) *InvoiceStatusCodeQuery {
	iscq.ctx.Offset = &offset
	return iscq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iscq *InvoiceStatusCodeQuery) Unique(unique bool) *InvoiceStatusCodeQuery {
	iscq.ctx.Unique = &unique
	return iscq
}

// Order specifies how the records should be ordered.
func (iscq *InvoiceStatusCodeQuery) Order(o ...invoicestatuscode.OrderOption) *InvoiceStatusCodeQuery {
	iscq.order = append(iscq.order, o...)
	return iscq
}

// First returns the first InvoiceStatusCode entity from the query.
// Returns a *NotFoundError when no InvoiceStatusCode was found.
func (iscq *InvoiceStatusCodeQuery) First(ctx context.Context) (*InvoiceStatusCode, error) {
	nodes, err := iscq.Limit(1).All(setContextOp(ctx, iscq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invoicestatuscode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iscq *InvoiceStatusCodeQuery) FirstX(ctx context.Context) *InvoiceStatusCode {
	node, err := iscq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InvoiceStatusCode ID from the query.
// Returns a *NotFoundError when no InvoiceStatusCode ID was found.
func (iscq *InvoiceStatusCodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iscq.Limit(1).IDs(setContextOp(ctx, iscq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invoicestatuscode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iscq *InvoiceStatusCodeQuery) FirstIDX(ctx context.Context) int {
	id, err := iscq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InvoiceStatusCode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InvoiceStatusCode entity is found.
// Returns a *NotFoundError when no InvoiceStatusCode entities are found.
func (iscq *InvoiceStatusCodeQuery) Only(ctx context.Context) (*InvoiceStatusCode, error) {
	nodes, err := iscq.Limit(2).All(setContextOp(ctx, iscq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invoicestatuscode.Label}
	default:
		return nil, &NotSingularError{invoicestatuscode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iscq *InvoiceStatusCodeQuery) OnlyX(ctx context.Context) *InvoiceStatusCode {
	node, err := iscq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InvoiceStatusCode ID in the query.
// Returns a *NotSingularError when more than one InvoiceStatusCode ID is found.
// Returns a *NotFoundError when no entities are found.
func (iscq *InvoiceStatusCodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iscq.Limit(2).IDs(setContextOp(ctx, iscq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invoicestatuscode.Label}
	default:
		err = &NotSingularError{invoicestatuscode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iscq *InvoiceStatusCodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := iscq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InvoiceStatusCodes.
func (iscq *InvoiceStatusCodeQuery) All(ctx context.Context) ([]*InvoiceStatusCode, error) {
	ctx = setContextOp(ctx, iscq.ctx, "All")
	if err := iscq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*InvoiceStatusCode, *InvoiceStatusCodeQuery]()
	return withInterceptors[[]*InvoiceStatusCode](ctx, iscq, qr, iscq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iscq *InvoiceStatusCodeQuery) AllX(ctx context.Context) []*InvoiceStatusCode {
	nodes, err := iscq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InvoiceStatusCode IDs.
func (iscq *InvoiceStatusCodeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if iscq.ctx.Unique == nil && iscq.path != nil {
		iscq.Unique(true)
	}
	ctx = setContextOp(ctx, iscq.ctx, "IDs")
	if err = iscq.Select(invoicestatuscode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iscq *InvoiceStatusCodeQuery) IDsX(ctx context.Context) []int {
	ids, err := iscq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iscq *InvoiceStatusCodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iscq.ctx, "Count")
	if err := iscq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iscq, querierCount[*InvoiceStatusCodeQuery](), iscq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iscq *InvoiceStatusCodeQuery) CountX(ctx context.Context) int {
	count, err := iscq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iscq *InvoiceStatusCodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iscq.ctx, "Exist")
	switch _, err := iscq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iscq *InvoiceStatusCodeQuery) ExistX(ctx context.Context) bool {
	exist, err := iscq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InvoiceStatusCodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iscq *InvoiceStatusCodeQuery) Clone() *InvoiceStatusCodeQuery {
	if iscq == nil {
		return nil
	}
	return &InvoiceStatusCodeQuery{
		config:     iscq.config,
		ctx:        iscq.ctx.Clone(),
		order:      append([]invoicestatuscode.OrderOption{}, iscq.order...),
		inters:     append([]Interceptor{}, iscq.inters...),
		predicates: append([]predicate.InvoiceStatusCode{}, iscq.predicates...),
		// clone intermediate query.
		sql:  iscq.sql.Clone(),
		path: iscq.path,
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
//	client.InvoiceStatusCode.Query().
//		GroupBy(invoicestatuscode.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iscq *InvoiceStatusCodeQuery) GroupBy(field string, fields ...string) *InvoiceStatusCodeGroupBy {
	iscq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InvoiceStatusCodeGroupBy{build: iscq}
	grbuild.flds = &iscq.ctx.Fields
	grbuild.label = invoicestatuscode.Label
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
//	client.InvoiceStatusCode.Query().
//		Select(invoicestatuscode.FieldCreatedAt).
//		Scan(ctx, &v)
func (iscq *InvoiceStatusCodeQuery) Select(fields ...string) *InvoiceStatusCodeSelect {
	iscq.ctx.Fields = append(iscq.ctx.Fields, fields...)
	sbuild := &InvoiceStatusCodeSelect{InvoiceStatusCodeQuery: iscq}
	sbuild.label = invoicestatuscode.Label
	sbuild.flds, sbuild.scan = &iscq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InvoiceStatusCodeSelect configured with the given aggregations.
func (iscq *InvoiceStatusCodeQuery) Aggregate(fns ...AggregateFunc) *InvoiceStatusCodeSelect {
	return iscq.Select().Aggregate(fns...)
}

func (iscq *InvoiceStatusCodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iscq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iscq); err != nil {
				return err
			}
		}
	}
	for _, f := range iscq.ctx.Fields {
		if !invoicestatuscode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iscq.path != nil {
		prev, err := iscq.path(ctx)
		if err != nil {
			return err
		}
		iscq.sql = prev
	}
	return nil
}

func (iscq *InvoiceStatusCodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*InvoiceStatusCode, error) {
	var (
		nodes = []*InvoiceStatusCode{}
		_spec = iscq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*InvoiceStatusCode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &InvoiceStatusCode{config: iscq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iscq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (iscq *InvoiceStatusCodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iscq.querySpec()
	_spec.Node.Columns = iscq.ctx.Fields
	if len(iscq.ctx.Fields) > 0 {
		_spec.Unique = iscq.ctx.Unique != nil && *iscq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iscq.driver, _spec)
}

func (iscq *InvoiceStatusCodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(invoicestatuscode.Table, invoicestatuscode.Columns, sqlgraph.NewFieldSpec(invoicestatuscode.FieldID, field.TypeInt))
	_spec.From = iscq.sql
	if unique := iscq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iscq.path != nil {
		_spec.Unique = true
	}
	if fields := iscq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoicestatuscode.FieldID)
		for i := range fields {
			if fields[i] != invoicestatuscode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iscq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iscq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iscq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iscq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iscq *InvoiceStatusCodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iscq.driver.Dialect())
	t1 := builder.Table(invoicestatuscode.Table)
	columns := iscq.ctx.Fields
	if len(columns) == 0 {
		columns = invoicestatuscode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iscq.sql != nil {
		selector = iscq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iscq.ctx.Unique != nil && *iscq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iscq.predicates {
		p(selector)
	}
	for _, p := range iscq.order {
		p(selector)
	}
	if offset := iscq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iscq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InvoiceStatusCodeGroupBy is the group-by builder for InvoiceStatusCode entities.
type InvoiceStatusCodeGroupBy struct {
	selector
	build *InvoiceStatusCodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iscgb *InvoiceStatusCodeGroupBy) Aggregate(fns ...AggregateFunc) *InvoiceStatusCodeGroupBy {
	iscgb.fns = append(iscgb.fns, fns...)
	return iscgb
}

// Scan applies the selector query and scans the result into the given value.
func (iscgb *InvoiceStatusCodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iscgb.build.ctx, "GroupBy")
	if err := iscgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceStatusCodeQuery, *InvoiceStatusCodeGroupBy](ctx, iscgb.build, iscgb, iscgb.build.inters, v)
}

func (iscgb *InvoiceStatusCodeGroupBy) sqlScan(ctx context.Context, root *InvoiceStatusCodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(iscgb.fns))
	for _, fn := range iscgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*iscgb.flds)+len(iscgb.fns))
		for _, f := range *iscgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*iscgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iscgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InvoiceStatusCodeSelect is the builder for selecting fields of InvoiceStatusCode entities.
type InvoiceStatusCodeSelect struct {
	*InvoiceStatusCodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (iscs *InvoiceStatusCodeSelect) Aggregate(fns ...AggregateFunc) *InvoiceStatusCodeSelect {
	iscs.fns = append(iscs.fns, fns...)
	return iscs
}

// Scan applies the selector query and scans the result into the given value.
func (iscs *InvoiceStatusCodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iscs.ctx, "Select")
	if err := iscs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceStatusCodeQuery, *InvoiceStatusCodeSelect](ctx, iscs.InvoiceStatusCodeQuery, iscs, iscs.inters, v)
}

func (iscs *InvoiceStatusCodeSelect) sqlScan(ctx context.Context, root *InvoiceStatusCodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(iscs.fns))
	for _, fn := range iscs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*iscs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iscs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}