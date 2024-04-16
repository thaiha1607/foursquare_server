// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/address"
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/personaddress"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// PersonAddressQuery is the builder for querying PersonAddress entities.
type PersonAddressQuery struct {
	config
	ctx           *QueryContext
	order         []personaddress.OrderOption
	inters        []Interceptor
	predicates    []predicate.PersonAddress
	withPersons   *PersonQuery
	withAddresses *AddressQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PersonAddressQuery builder.
func (paq *PersonAddressQuery) Where(ps ...predicate.PersonAddress) *PersonAddressQuery {
	paq.predicates = append(paq.predicates, ps...)
	return paq
}

// Limit the number of records to be returned by this query.
func (paq *PersonAddressQuery) Limit(limit int) *PersonAddressQuery {
	paq.ctx.Limit = &limit
	return paq
}

// Offset to start from.
func (paq *PersonAddressQuery) Offset(offset int) *PersonAddressQuery {
	paq.ctx.Offset = &offset
	return paq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (paq *PersonAddressQuery) Unique(unique bool) *PersonAddressQuery {
	paq.ctx.Unique = &unique
	return paq
}

// Order specifies how the records should be ordered.
func (paq *PersonAddressQuery) Order(o ...personaddress.OrderOption) *PersonAddressQuery {
	paq.order = append(paq.order, o...)
	return paq
}

// QueryPersons chains the current query on the "persons" edge.
func (paq *PersonAddressQuery) QueryPersons() *PersonQuery {
	query := (&PersonClient{config: paq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(personaddress.Table, personaddress.PersonsColumn, selector),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, personaddress.PersonsTable, personaddress.PersonsColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAddresses chains the current query on the "addresses" edge.
func (paq *PersonAddressQuery) QueryAddresses() *AddressQuery {
	query := (&AddressClient{config: paq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(personaddress.Table, personaddress.AddressesColumn, selector),
			sqlgraph.To(address.Table, address.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, personaddress.AddressesTable, personaddress.AddressesColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PersonAddress entity from the query.
// Returns a *NotFoundError when no PersonAddress was found.
func (paq *PersonAddressQuery) First(ctx context.Context) (*PersonAddress, error) {
	nodes, err := paq.Limit(1).All(setContextOp(ctx, paq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{personaddress.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (paq *PersonAddressQuery) FirstX(ctx context.Context) *PersonAddress {
	node, err := paq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single PersonAddress entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PersonAddress entity is found.
// Returns a *NotFoundError when no PersonAddress entities are found.
func (paq *PersonAddressQuery) Only(ctx context.Context) (*PersonAddress, error) {
	nodes, err := paq.Limit(2).All(setContextOp(ctx, paq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{personaddress.Label}
	default:
		return nil, &NotSingularError{personaddress.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (paq *PersonAddressQuery) OnlyX(ctx context.Context) *PersonAddress {
	node, err := paq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of PersonAddresses.
func (paq *PersonAddressQuery) All(ctx context.Context) ([]*PersonAddress, error) {
	ctx = setContextOp(ctx, paq.ctx, "All")
	if err := paq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PersonAddress, *PersonAddressQuery]()
	return withInterceptors[[]*PersonAddress](ctx, paq, qr, paq.inters)
}

// AllX is like All, but panics if an error occurs.
func (paq *PersonAddressQuery) AllX(ctx context.Context) []*PersonAddress {
	nodes, err := paq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (paq *PersonAddressQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, paq.ctx, "Count")
	if err := paq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, paq, querierCount[*PersonAddressQuery](), paq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (paq *PersonAddressQuery) CountX(ctx context.Context) int {
	count, err := paq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (paq *PersonAddressQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, paq.ctx, "Exist")
	switch _, err := paq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (paq *PersonAddressQuery) ExistX(ctx context.Context) bool {
	exist, err := paq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PersonAddressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (paq *PersonAddressQuery) Clone() *PersonAddressQuery {
	if paq == nil {
		return nil
	}
	return &PersonAddressQuery{
		config:        paq.config,
		ctx:           paq.ctx.Clone(),
		order:         append([]personaddress.OrderOption{}, paq.order...),
		inters:        append([]Interceptor{}, paq.inters...),
		predicates:    append([]predicate.PersonAddress{}, paq.predicates...),
		withPersons:   paq.withPersons.Clone(),
		withAddresses: paq.withAddresses.Clone(),
		// clone intermediate query.
		sql:  paq.sql.Clone(),
		path: paq.path,
	}
}

// WithPersons tells the query-builder to eager-load the nodes that are connected to
// the "persons" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *PersonAddressQuery) WithPersons(opts ...func(*PersonQuery)) *PersonAddressQuery {
	query := (&PersonClient{config: paq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	paq.withPersons = query
	return paq
}

// WithAddresses tells the query-builder to eager-load the nodes that are connected to
// the "addresses" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *PersonAddressQuery) WithAddresses(opts ...func(*AddressQuery)) *PersonAddressQuery {
	query := (&AddressClient{config: paq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	paq.withAddresses = query
	return paq
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
//	client.PersonAddress.Query().
//		GroupBy(personaddress.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (paq *PersonAddressQuery) GroupBy(field string, fields ...string) *PersonAddressGroupBy {
	paq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PersonAddressGroupBy{build: paq}
	grbuild.flds = &paq.ctx.Fields
	grbuild.label = personaddress.Label
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
//	client.PersonAddress.Query().
//		Select(personaddress.FieldCreatedAt).
//		Scan(ctx, &v)
func (paq *PersonAddressQuery) Select(fields ...string) *PersonAddressSelect {
	paq.ctx.Fields = append(paq.ctx.Fields, fields...)
	sbuild := &PersonAddressSelect{PersonAddressQuery: paq}
	sbuild.label = personaddress.Label
	sbuild.flds, sbuild.scan = &paq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PersonAddressSelect configured with the given aggregations.
func (paq *PersonAddressQuery) Aggregate(fns ...AggregateFunc) *PersonAddressSelect {
	return paq.Select().Aggregate(fns...)
}

func (paq *PersonAddressQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range paq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, paq); err != nil {
				return err
			}
		}
	}
	for _, f := range paq.ctx.Fields {
		if !personaddress.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if paq.path != nil {
		prev, err := paq.path(ctx)
		if err != nil {
			return err
		}
		paq.sql = prev
	}
	return nil
}

func (paq *PersonAddressQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PersonAddress, error) {
	var (
		nodes       = []*PersonAddress{}
		_spec       = paq.querySpec()
		loadedTypes = [2]bool{
			paq.withPersons != nil,
			paq.withAddresses != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PersonAddress).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PersonAddress{config: paq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, paq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := paq.withPersons; query != nil {
		if err := paq.loadPersons(ctx, query, nodes, nil,
			func(n *PersonAddress, e *Person) { n.Edges.Persons = e }); err != nil {
			return nil, err
		}
	}
	if query := paq.withAddresses; query != nil {
		if err := paq.loadAddresses(ctx, query, nodes, nil,
			func(n *PersonAddress, e *Address) { n.Edges.Addresses = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (paq *PersonAddressQuery) loadPersons(ctx context.Context, query *PersonQuery, nodes []*PersonAddress, init func(*PersonAddress), assign func(*PersonAddress, *Person)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*PersonAddress)
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
func (paq *PersonAddressQuery) loadAddresses(ctx context.Context, query *AddressQuery, nodes []*PersonAddress, init func(*PersonAddress), assign func(*PersonAddress, *Address)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*PersonAddress)
	for i := range nodes {
		fk := nodes[i].AddressID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(address.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "address_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (paq *PersonAddressQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := paq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, paq.driver, _spec)
}

func (paq *PersonAddressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(personaddress.Table, personaddress.Columns, nil)
	_spec.From = paq.sql
	if unique := paq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if paq.path != nil {
		_spec.Unique = true
	}
	if fields := paq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if paq.withPersons != nil {
			_spec.Node.AddColumnOnce(personaddress.FieldPersonID)
		}
		if paq.withAddresses != nil {
			_spec.Node.AddColumnOnce(personaddress.FieldAddressID)
		}
	}
	if ps := paq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := paq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := paq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := paq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (paq *PersonAddressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(paq.driver.Dialect())
	t1 := builder.Table(personaddress.Table)
	columns := paq.ctx.Fields
	if len(columns) == 0 {
		columns = personaddress.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if paq.sql != nil {
		selector = paq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if paq.ctx.Unique != nil && *paq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range paq.predicates {
		p(selector)
	}
	for _, p := range paq.order {
		p(selector)
	}
	if offset := paq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := paq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PersonAddressGroupBy is the group-by builder for PersonAddress entities.
type PersonAddressGroupBy struct {
	selector
	build *PersonAddressQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pagb *PersonAddressGroupBy) Aggregate(fns ...AggregateFunc) *PersonAddressGroupBy {
	pagb.fns = append(pagb.fns, fns...)
	return pagb
}

// Scan applies the selector query and scans the result into the given value.
func (pagb *PersonAddressGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pagb.build.ctx, "GroupBy")
	if err := pagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PersonAddressQuery, *PersonAddressGroupBy](ctx, pagb.build, pagb, pagb.build.inters, v)
}

func (pagb *PersonAddressGroupBy) sqlScan(ctx context.Context, root *PersonAddressQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pagb.fns))
	for _, fn := range pagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pagb.flds)+len(pagb.fns))
		for _, f := range *pagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PersonAddressSelect is the builder for selecting fields of PersonAddress entities.
type PersonAddressSelect struct {
	*PersonAddressQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pas *PersonAddressSelect) Aggregate(fns ...AggregateFunc) *PersonAddressSelect {
	pas.fns = append(pas.fns, fns...)
	return pas
}

// Scan applies the selector query and scans the result into the given value.
func (pas *PersonAddressSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pas.ctx, "Select")
	if err := pas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PersonAddressQuery, *PersonAddressSelect](ctx, pas.PersonAddressQuery, pas, pas.inters, v)
}

func (pas *PersonAddressSelect) sqlScan(ctx context.Context, root *PersonAddressQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pas.fns))
	for _, fn := range pas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}