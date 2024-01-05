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
	"github.com/thaiha1607/foursquare_server/ent/financialtransaction"
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// FinancialTransactionQuery is the builder for querying FinancialTransaction entities.
type FinancialTransactionQuery struct {
	config
	ctx         *QueryContext
	order       []financialtransaction.OrderOption
	inters      []Interceptor
	predicates  []predicate.FinancialTransaction
	withInvoice *InvoiceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FinancialTransactionQuery builder.
func (ftq *FinancialTransactionQuery) Where(ps ...predicate.FinancialTransaction) *FinancialTransactionQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit the number of records to be returned by this query.
func (ftq *FinancialTransactionQuery) Limit(limit int) *FinancialTransactionQuery {
	ftq.ctx.Limit = &limit
	return ftq
}

// Offset to start from.
func (ftq *FinancialTransactionQuery) Offset(offset int) *FinancialTransactionQuery {
	ftq.ctx.Offset = &offset
	return ftq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ftq *FinancialTransactionQuery) Unique(unique bool) *FinancialTransactionQuery {
	ftq.ctx.Unique = &unique
	return ftq
}

// Order specifies how the records should be ordered.
func (ftq *FinancialTransactionQuery) Order(o ...financialtransaction.OrderOption) *FinancialTransactionQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// QueryInvoice chains the current query on the "invoice" edge.
func (ftq *FinancialTransactionQuery) QueryInvoice() *InvoiceQuery {
	query := (&InvoiceClient{config: ftq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ftq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(financialtransaction.Table, financialtransaction.FieldID, selector),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, financialtransaction.InvoiceTable, financialtransaction.InvoiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(ftq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FinancialTransaction entity from the query.
// Returns a *NotFoundError when no FinancialTransaction was found.
func (ftq *FinancialTransactionQuery) First(ctx context.Context) (*FinancialTransaction, error) {
	nodes, err := ftq.Limit(1).All(setContextOp(ctx, ftq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{financialtransaction.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FinancialTransactionQuery) FirstX(ctx context.Context) *FinancialTransaction {
	node, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FinancialTransaction ID from the query.
// Returns a *NotFoundError when no FinancialTransaction ID was found.
func (ftq *FinancialTransactionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ftq.Limit(1).IDs(setContextOp(ctx, ftq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{financialtransaction.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ftq *FinancialTransactionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FinancialTransaction entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FinancialTransaction entity is found.
// Returns a *NotFoundError when no FinancialTransaction entities are found.
func (ftq *FinancialTransactionQuery) Only(ctx context.Context) (*FinancialTransaction, error) {
	nodes, err := ftq.Limit(2).All(setContextOp(ctx, ftq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{financialtransaction.Label}
	default:
		return nil, &NotSingularError{financialtransaction.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FinancialTransactionQuery) OnlyX(ctx context.Context) *FinancialTransaction {
	node, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FinancialTransaction ID in the query.
// Returns a *NotSingularError when more than one FinancialTransaction ID is found.
// Returns a *NotFoundError when no entities are found.
func (ftq *FinancialTransactionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ftq.Limit(2).IDs(setContextOp(ctx, ftq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{financialtransaction.Label}
	default:
		err = &NotSingularError{financialtransaction.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ftq *FinancialTransactionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FinancialTransactions.
func (ftq *FinancialTransactionQuery) All(ctx context.Context) ([]*FinancialTransaction, error) {
	ctx = setContextOp(ctx, ftq.ctx, "All")
	if err := ftq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FinancialTransaction, *FinancialTransactionQuery]()
	return withInterceptors[[]*FinancialTransaction](ctx, ftq, qr, ftq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ftq *FinancialTransactionQuery) AllX(ctx context.Context) []*FinancialTransaction {
	nodes, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FinancialTransaction IDs.
func (ftq *FinancialTransactionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ftq.ctx.Unique == nil && ftq.path != nil {
		ftq.Unique(true)
	}
	ctx = setContextOp(ctx, ftq.ctx, "IDs")
	if err = ftq.Select(financialtransaction.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FinancialTransactionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FinancialTransactionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ftq.ctx, "Count")
	if err := ftq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ftq, querierCount[*FinancialTransactionQuery](), ftq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FinancialTransactionQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FinancialTransactionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ftq.ctx, "Exist")
	switch _, err := ftq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FinancialTransactionQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FinancialTransactionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftq *FinancialTransactionQuery) Clone() *FinancialTransactionQuery {
	if ftq == nil {
		return nil
	}
	return &FinancialTransactionQuery{
		config:      ftq.config,
		ctx:         ftq.ctx.Clone(),
		order:       append([]financialtransaction.OrderOption{}, ftq.order...),
		inters:      append([]Interceptor{}, ftq.inters...),
		predicates:  append([]predicate.FinancialTransaction{}, ftq.predicates...),
		withInvoice: ftq.withInvoice.Clone(),
		// clone intermediate query.
		sql:  ftq.sql.Clone(),
		path: ftq.path,
	}
}

// WithInvoice tells the query-builder to eager-load the nodes that are connected to
// the "invoice" edge. The optional arguments are used to configure the query builder of the edge.
func (ftq *FinancialTransactionQuery) WithInvoice(opts ...func(*InvoiceQuery)) *FinancialTransactionQuery {
	query := (&InvoiceClient{config: ftq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ftq.withInvoice = query
	return ftq
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
//	client.FinancialTransaction.Query().
//		GroupBy(financialtransaction.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ftq *FinancialTransactionQuery) GroupBy(field string, fields ...string) *FinancialTransactionGroupBy {
	ftq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FinancialTransactionGroupBy{build: ftq}
	grbuild.flds = &ftq.ctx.Fields
	grbuild.label = financialtransaction.Label
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
//	client.FinancialTransaction.Query().
//		Select(financialtransaction.FieldCreatedAt).
//		Scan(ctx, &v)
func (ftq *FinancialTransactionQuery) Select(fields ...string) *FinancialTransactionSelect {
	ftq.ctx.Fields = append(ftq.ctx.Fields, fields...)
	sbuild := &FinancialTransactionSelect{FinancialTransactionQuery: ftq}
	sbuild.label = financialtransaction.Label
	sbuild.flds, sbuild.scan = &ftq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FinancialTransactionSelect configured with the given aggregations.
func (ftq *FinancialTransactionQuery) Aggregate(fns ...AggregateFunc) *FinancialTransactionSelect {
	return ftq.Select().Aggregate(fns...)
}

func (ftq *FinancialTransactionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ftq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ftq); err != nil {
				return err
			}
		}
	}
	for _, f := range ftq.ctx.Fields {
		if !financialtransaction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ftq.path != nil {
		prev, err := ftq.path(ctx)
		if err != nil {
			return err
		}
		ftq.sql = prev
	}
	return nil
}

func (ftq *FinancialTransactionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FinancialTransaction, error) {
	var (
		nodes       = []*FinancialTransaction{}
		_spec       = ftq.querySpec()
		loadedTypes = [1]bool{
			ftq.withInvoice != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FinancialTransaction).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FinancialTransaction{config: ftq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ftq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ftq.withInvoice; query != nil {
		if err := ftq.loadInvoice(ctx, query, nodes, nil,
			func(n *FinancialTransaction, e *Invoice) { n.Edges.Invoice = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ftq *FinancialTransactionQuery) loadInvoice(ctx context.Context, query *InvoiceQuery, nodes []*FinancialTransaction, init func(*FinancialTransaction), assign func(*FinancialTransaction, *Invoice)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*FinancialTransaction)
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

func (ftq *FinancialTransactionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ftq.querySpec()
	_spec.Node.Columns = ftq.ctx.Fields
	if len(ftq.ctx.Fields) > 0 {
		_spec.Unique = ftq.ctx.Unique != nil && *ftq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ftq.driver, _spec)
}

func (ftq *FinancialTransactionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(financialtransaction.Table, financialtransaction.Columns, sqlgraph.NewFieldSpec(financialtransaction.FieldID, field.TypeUUID))
	_spec.From = ftq.sql
	if unique := ftq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ftq.path != nil {
		_spec.Unique = true
	}
	if fields := ftq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, financialtransaction.FieldID)
		for i := range fields {
			if fields[i] != financialtransaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ftq.withInvoice != nil {
			_spec.Node.AddColumnOnce(financialtransaction.FieldInvoiceID)
		}
	}
	if ps := ftq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ftq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ftq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ftq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ftq *FinancialTransactionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ftq.driver.Dialect())
	t1 := builder.Table(financialtransaction.Table)
	columns := ftq.ctx.Fields
	if len(columns) == 0 {
		columns = financialtransaction.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ftq.sql != nil {
		selector = ftq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ftq.ctx.Unique != nil && *ftq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ftq.predicates {
		p(selector)
	}
	for _, p := range ftq.order {
		p(selector)
	}
	if offset := ftq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ftq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FinancialTransactionGroupBy is the group-by builder for FinancialTransaction entities.
type FinancialTransactionGroupBy struct {
	selector
	build *FinancialTransactionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FinancialTransactionGroupBy) Aggregate(fns ...AggregateFunc) *FinancialTransactionGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the selector query and scans the result into the given value.
func (ftgb *FinancialTransactionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ftgb.build.ctx, "GroupBy")
	if err := ftgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FinancialTransactionQuery, *FinancialTransactionGroupBy](ctx, ftgb.build, ftgb, ftgb.build.inters, v)
}

func (ftgb *FinancialTransactionGroupBy) sqlScan(ctx context.Context, root *FinancialTransactionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ftgb.fns))
	for _, fn := range ftgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ftgb.flds)+len(ftgb.fns))
		for _, f := range *ftgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ftgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ftgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FinancialTransactionSelect is the builder for selecting fields of FinancialTransaction entities.
type FinancialTransactionSelect struct {
	*FinancialTransactionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fts *FinancialTransactionSelect) Aggregate(fns ...AggregateFunc) *FinancialTransactionSelect {
	fts.fns = append(fts.fns, fns...)
	return fts
}

// Scan applies the selector query and scans the result into the given value.
func (fts *FinancialTransactionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fts.ctx, "Select")
	if err := fts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FinancialTransactionQuery, *FinancialTransactionSelect](ctx, fts.FinancialTransactionQuery, fts, fts.inters, v)
}

func (fts *FinancialTransactionSelect) sqlScan(ctx context.Context, root *FinancialTransactionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fts.fns))
	for _, fn := range fts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
