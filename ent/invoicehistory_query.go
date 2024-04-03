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
	"github.com/thaiha1607/foursquare_server/ent/invoicehistory"
	"github.com/thaiha1607/foursquare_server/ent/invoicestatuscode"
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// InvoiceHistoryQuery is the builder for querying InvoiceHistory entities.
type InvoiceHistoryQuery struct {
	config
	ctx           *QueryContext
	order         []invoicehistory.OrderOption
	inters        []Interceptor
	predicates    []predicate.InvoiceHistory
	withInvoice   *InvoiceQuery
	withPerson    *PersonQuery
	withOldStatus *InvoiceStatusCodeQuery
	withNewStatus *InvoiceStatusCodeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InvoiceHistoryQuery builder.
func (ihq *InvoiceHistoryQuery) Where(ps ...predicate.InvoiceHistory) *InvoiceHistoryQuery {
	ihq.predicates = append(ihq.predicates, ps...)
	return ihq
}

// Limit the number of records to be returned by this query.
func (ihq *InvoiceHistoryQuery) Limit(limit int) *InvoiceHistoryQuery {
	ihq.ctx.Limit = &limit
	return ihq
}

// Offset to start from.
func (ihq *InvoiceHistoryQuery) Offset(offset int) *InvoiceHistoryQuery {
	ihq.ctx.Offset = &offset
	return ihq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ihq *InvoiceHistoryQuery) Unique(unique bool) *InvoiceHistoryQuery {
	ihq.ctx.Unique = &unique
	return ihq
}

// Order specifies how the records should be ordered.
func (ihq *InvoiceHistoryQuery) Order(o ...invoicehistory.OrderOption) *InvoiceHistoryQuery {
	ihq.order = append(ihq.order, o...)
	return ihq
}

// QueryInvoice chains the current query on the "invoice" edge.
func (ihq *InvoiceHistoryQuery) QueryInvoice() *InvoiceQuery {
	query := (&InvoiceClient{config: ihq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ihq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ihq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoicehistory.Table, invoicehistory.FieldID, selector),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, invoicehistory.InvoiceTable, invoicehistory.InvoiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(ihq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPerson chains the current query on the "person" edge.
func (ihq *InvoiceHistoryQuery) QueryPerson() *PersonQuery {
	query := (&PersonClient{config: ihq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ihq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ihq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoicehistory.Table, invoicehistory.FieldID, selector),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, invoicehistory.PersonTable, invoicehistory.PersonColumn),
		)
		fromU = sqlgraph.SetNeighbors(ihq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOldStatus chains the current query on the "old_status" edge.
func (ihq *InvoiceHistoryQuery) QueryOldStatus() *InvoiceStatusCodeQuery {
	query := (&InvoiceStatusCodeClient{config: ihq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ihq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ihq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoicehistory.Table, invoicehistory.FieldID, selector),
			sqlgraph.To(invoicestatuscode.Table, invoicestatuscode.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, invoicehistory.OldStatusTable, invoicehistory.OldStatusColumn),
		)
		fromU = sqlgraph.SetNeighbors(ihq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNewStatus chains the current query on the "new_status" edge.
func (ihq *InvoiceHistoryQuery) QueryNewStatus() *InvoiceStatusCodeQuery {
	query := (&InvoiceStatusCodeClient{config: ihq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ihq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ihq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoicehistory.Table, invoicehistory.FieldID, selector),
			sqlgraph.To(invoicestatuscode.Table, invoicestatuscode.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, invoicehistory.NewStatusTable, invoicehistory.NewStatusColumn),
		)
		fromU = sqlgraph.SetNeighbors(ihq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InvoiceHistory entity from the query.
// Returns a *NotFoundError when no InvoiceHistory was found.
func (ihq *InvoiceHistoryQuery) First(ctx context.Context) (*InvoiceHistory, error) {
	nodes, err := ihq.Limit(1).All(setContextOp(ctx, ihq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invoicehistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ihq *InvoiceHistoryQuery) FirstX(ctx context.Context) *InvoiceHistory {
	node, err := ihq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InvoiceHistory ID from the query.
// Returns a *NotFoundError when no InvoiceHistory ID was found.
func (ihq *InvoiceHistoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ihq.Limit(1).IDs(setContextOp(ctx, ihq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invoicehistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ihq *InvoiceHistoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ihq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InvoiceHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InvoiceHistory entity is found.
// Returns a *NotFoundError when no InvoiceHistory entities are found.
func (ihq *InvoiceHistoryQuery) Only(ctx context.Context) (*InvoiceHistory, error) {
	nodes, err := ihq.Limit(2).All(setContextOp(ctx, ihq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invoicehistory.Label}
	default:
		return nil, &NotSingularError{invoicehistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ihq *InvoiceHistoryQuery) OnlyX(ctx context.Context) *InvoiceHistory {
	node, err := ihq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InvoiceHistory ID in the query.
// Returns a *NotSingularError when more than one InvoiceHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ihq *InvoiceHistoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ihq.Limit(2).IDs(setContextOp(ctx, ihq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invoicehistory.Label}
	default:
		err = &NotSingularError{invoicehistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ihq *InvoiceHistoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ihq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InvoiceHistories.
func (ihq *InvoiceHistoryQuery) All(ctx context.Context) ([]*InvoiceHistory, error) {
	ctx = setContextOp(ctx, ihq.ctx, "All")
	if err := ihq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*InvoiceHistory, *InvoiceHistoryQuery]()
	return withInterceptors[[]*InvoiceHistory](ctx, ihq, qr, ihq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ihq *InvoiceHistoryQuery) AllX(ctx context.Context) []*InvoiceHistory {
	nodes, err := ihq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InvoiceHistory IDs.
func (ihq *InvoiceHistoryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ihq.ctx.Unique == nil && ihq.path != nil {
		ihq.Unique(true)
	}
	ctx = setContextOp(ctx, ihq.ctx, "IDs")
	if err = ihq.Select(invoicehistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ihq *InvoiceHistoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ihq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ihq *InvoiceHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ihq.ctx, "Count")
	if err := ihq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ihq, querierCount[*InvoiceHistoryQuery](), ihq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ihq *InvoiceHistoryQuery) CountX(ctx context.Context) int {
	count, err := ihq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ihq *InvoiceHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ihq.ctx, "Exist")
	switch _, err := ihq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ihq *InvoiceHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ihq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InvoiceHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ihq *InvoiceHistoryQuery) Clone() *InvoiceHistoryQuery {
	if ihq == nil {
		return nil
	}
	return &InvoiceHistoryQuery{
		config:        ihq.config,
		ctx:           ihq.ctx.Clone(),
		order:         append([]invoicehistory.OrderOption{}, ihq.order...),
		inters:        append([]Interceptor{}, ihq.inters...),
		predicates:    append([]predicate.InvoiceHistory{}, ihq.predicates...),
		withInvoice:   ihq.withInvoice.Clone(),
		withPerson:    ihq.withPerson.Clone(),
		withOldStatus: ihq.withOldStatus.Clone(),
		withNewStatus: ihq.withNewStatus.Clone(),
		// clone intermediate query.
		sql:  ihq.sql.Clone(),
		path: ihq.path,
	}
}

// WithInvoice tells the query-builder to eager-load the nodes that are connected to
// the "invoice" edge. The optional arguments are used to configure the query builder of the edge.
func (ihq *InvoiceHistoryQuery) WithInvoice(opts ...func(*InvoiceQuery)) *InvoiceHistoryQuery {
	query := (&InvoiceClient{config: ihq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ihq.withInvoice = query
	return ihq
}

// WithPerson tells the query-builder to eager-load the nodes that are connected to
// the "person" edge. The optional arguments are used to configure the query builder of the edge.
func (ihq *InvoiceHistoryQuery) WithPerson(opts ...func(*PersonQuery)) *InvoiceHistoryQuery {
	query := (&PersonClient{config: ihq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ihq.withPerson = query
	return ihq
}

// WithOldStatus tells the query-builder to eager-load the nodes that are connected to
// the "old_status" edge. The optional arguments are used to configure the query builder of the edge.
func (ihq *InvoiceHistoryQuery) WithOldStatus(opts ...func(*InvoiceStatusCodeQuery)) *InvoiceHistoryQuery {
	query := (&InvoiceStatusCodeClient{config: ihq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ihq.withOldStatus = query
	return ihq
}

// WithNewStatus tells the query-builder to eager-load the nodes that are connected to
// the "new_status" edge. The optional arguments are used to configure the query builder of the edge.
func (ihq *InvoiceHistoryQuery) WithNewStatus(opts ...func(*InvoiceStatusCodeQuery)) *InvoiceHistoryQuery {
	query := (&InvoiceStatusCodeClient{config: ihq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ihq.withNewStatus = query
	return ihq
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
//	client.InvoiceHistory.Query().
//		GroupBy(invoicehistory.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ihq *InvoiceHistoryQuery) GroupBy(field string, fields ...string) *InvoiceHistoryGroupBy {
	ihq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InvoiceHistoryGroupBy{build: ihq}
	grbuild.flds = &ihq.ctx.Fields
	grbuild.label = invoicehistory.Label
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
//	client.InvoiceHistory.Query().
//		Select(invoicehistory.FieldCreatedAt).
//		Scan(ctx, &v)
func (ihq *InvoiceHistoryQuery) Select(fields ...string) *InvoiceHistorySelect {
	ihq.ctx.Fields = append(ihq.ctx.Fields, fields...)
	sbuild := &InvoiceHistorySelect{InvoiceHistoryQuery: ihq}
	sbuild.label = invoicehistory.Label
	sbuild.flds, sbuild.scan = &ihq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InvoiceHistorySelect configured with the given aggregations.
func (ihq *InvoiceHistoryQuery) Aggregate(fns ...AggregateFunc) *InvoiceHistorySelect {
	return ihq.Select().Aggregate(fns...)
}

func (ihq *InvoiceHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ihq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ihq); err != nil {
				return err
			}
		}
	}
	for _, f := range ihq.ctx.Fields {
		if !invoicehistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ihq.path != nil {
		prev, err := ihq.path(ctx)
		if err != nil {
			return err
		}
		ihq.sql = prev
	}
	return nil
}

func (ihq *InvoiceHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*InvoiceHistory, error) {
	var (
		nodes       = []*InvoiceHistory{}
		_spec       = ihq.querySpec()
		loadedTypes = [4]bool{
			ihq.withInvoice != nil,
			ihq.withPerson != nil,
			ihq.withOldStatus != nil,
			ihq.withNewStatus != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*InvoiceHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &InvoiceHistory{config: ihq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ihq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ihq.withInvoice; query != nil {
		if err := ihq.loadInvoice(ctx, query, nodes, nil,
			func(n *InvoiceHistory, e *Invoice) { n.Edges.Invoice = e }); err != nil {
			return nil, err
		}
	}
	if query := ihq.withPerson; query != nil {
		if err := ihq.loadPerson(ctx, query, nodes, nil,
			func(n *InvoiceHistory, e *Person) { n.Edges.Person = e }); err != nil {
			return nil, err
		}
	}
	if query := ihq.withOldStatus; query != nil {
		if err := ihq.loadOldStatus(ctx, query, nodes, nil,
			func(n *InvoiceHistory, e *InvoiceStatusCode) { n.Edges.OldStatus = e }); err != nil {
			return nil, err
		}
	}
	if query := ihq.withNewStatus; query != nil {
		if err := ihq.loadNewStatus(ctx, query, nodes, nil,
			func(n *InvoiceHistory, e *InvoiceStatusCode) { n.Edges.NewStatus = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ihq *InvoiceHistoryQuery) loadInvoice(ctx context.Context, query *InvoiceQuery, nodes []*InvoiceHistory, init func(*InvoiceHistory), assign func(*InvoiceHistory, *Invoice)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*InvoiceHistory)
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
func (ihq *InvoiceHistoryQuery) loadPerson(ctx context.Context, query *PersonQuery, nodes []*InvoiceHistory, init func(*InvoiceHistory), assign func(*InvoiceHistory, *Person)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*InvoiceHistory)
	for i := range nodes {
		fk := nodes[i].PersonID
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
			return fmt.Errorf(`unexpected foreign-key "person_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ihq *InvoiceHistoryQuery) loadOldStatus(ctx context.Context, query *InvoiceStatusCodeQuery, nodes []*InvoiceHistory, init func(*InvoiceHistory), assign func(*InvoiceHistory, *InvoiceStatusCode)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*InvoiceHistory)
	for i := range nodes {
		if nodes[i].OldStatusCode == nil {
			continue
		}
		fk := *nodes[i].OldStatusCode
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(invoicestatuscode.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "old_status_code" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ihq *InvoiceHistoryQuery) loadNewStatus(ctx context.Context, query *InvoiceStatusCodeQuery, nodes []*InvoiceHistory, init func(*InvoiceHistory), assign func(*InvoiceHistory, *InvoiceStatusCode)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*InvoiceHistory)
	for i := range nodes {
		if nodes[i].NewStatusCode == nil {
			continue
		}
		fk := *nodes[i].NewStatusCode
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(invoicestatuscode.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "new_status_code" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ihq *InvoiceHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ihq.querySpec()
	_spec.Node.Columns = ihq.ctx.Fields
	if len(ihq.ctx.Fields) > 0 {
		_spec.Unique = ihq.ctx.Unique != nil && *ihq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ihq.driver, _spec)
}

func (ihq *InvoiceHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(invoicehistory.Table, invoicehistory.Columns, sqlgraph.NewFieldSpec(invoicehistory.FieldID, field.TypeUUID))
	_spec.From = ihq.sql
	if unique := ihq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ihq.path != nil {
		_spec.Unique = true
	}
	if fields := ihq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoicehistory.FieldID)
		for i := range fields {
			if fields[i] != invoicehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ihq.withInvoice != nil {
			_spec.Node.AddColumnOnce(invoicehistory.FieldInvoiceID)
		}
		if ihq.withPerson != nil {
			_spec.Node.AddColumnOnce(invoicehistory.FieldPersonID)
		}
		if ihq.withOldStatus != nil {
			_spec.Node.AddColumnOnce(invoicehistory.FieldOldStatusCode)
		}
		if ihq.withNewStatus != nil {
			_spec.Node.AddColumnOnce(invoicehistory.FieldNewStatusCode)
		}
	}
	if ps := ihq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ihq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ihq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ihq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ihq *InvoiceHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ihq.driver.Dialect())
	t1 := builder.Table(invoicehistory.Table)
	columns := ihq.ctx.Fields
	if len(columns) == 0 {
		columns = invoicehistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ihq.sql != nil {
		selector = ihq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ihq.ctx.Unique != nil && *ihq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ihq.predicates {
		p(selector)
	}
	for _, p := range ihq.order {
		p(selector)
	}
	if offset := ihq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ihq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InvoiceHistoryGroupBy is the group-by builder for InvoiceHistory entities.
type InvoiceHistoryGroupBy struct {
	selector
	build *InvoiceHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ihgb *InvoiceHistoryGroupBy) Aggregate(fns ...AggregateFunc) *InvoiceHistoryGroupBy {
	ihgb.fns = append(ihgb.fns, fns...)
	return ihgb
}

// Scan applies the selector query and scans the result into the given value.
func (ihgb *InvoiceHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ihgb.build.ctx, "GroupBy")
	if err := ihgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceHistoryQuery, *InvoiceHistoryGroupBy](ctx, ihgb.build, ihgb, ihgb.build.inters, v)
}

func (ihgb *InvoiceHistoryGroupBy) sqlScan(ctx context.Context, root *InvoiceHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ihgb.fns))
	for _, fn := range ihgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ihgb.flds)+len(ihgb.fns))
		for _, f := range *ihgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ihgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ihgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InvoiceHistorySelect is the builder for selecting fields of InvoiceHistory entities.
type InvoiceHistorySelect struct {
	*InvoiceHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ihs *InvoiceHistorySelect) Aggregate(fns ...AggregateFunc) *InvoiceHistorySelect {
	ihs.fns = append(ihs.fns, fns...)
	return ihs
}

// Scan applies the selector query and scans the result into the given value.
func (ihs *InvoiceHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ihs.ctx, "Select")
	if err := ihs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceHistoryQuery, *InvoiceHistorySelect](ctx, ihs.InvoiceHistoryQuery, ihs, ihs.inters, v)
}

func (ihs *InvoiceHistorySelect) sqlScan(ctx context.Context, root *InvoiceHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ihs.fns))
	for _, fn := range ihs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ihs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ihs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
