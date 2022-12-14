// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysdictionary"
	"slash-admin/app/admin/ent/sysdictionarydetail"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysDictionaryQuery is the builder for querying SysDictionary entities.
type SysDictionaryQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.SysDictionary
	withDetails *SysDictionaryDetailQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysDictionaryQuery builder.
func (sdq *SysDictionaryQuery) Where(ps ...predicate.SysDictionary) *SysDictionaryQuery {
	sdq.predicates = append(sdq.predicates, ps...)
	return sdq
}

// Limit adds a limit step to the query.
func (sdq *SysDictionaryQuery) Limit(limit int) *SysDictionaryQuery {
	sdq.limit = &limit
	return sdq
}

// Offset adds an offset step to the query.
func (sdq *SysDictionaryQuery) Offset(offset int) *SysDictionaryQuery {
	sdq.offset = &offset
	return sdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sdq *SysDictionaryQuery) Unique(unique bool) *SysDictionaryQuery {
	sdq.unique = &unique
	return sdq
}

// Order adds an order step to the query.
func (sdq *SysDictionaryQuery) Order(o ...OrderFunc) *SysDictionaryQuery {
	sdq.order = append(sdq.order, o...)
	return sdq
}

// QueryDetails chains the current query on the "details" edge.
func (sdq *SysDictionaryQuery) QueryDetails() *SysDictionaryDetailQuery {
	query := &SysDictionaryDetailQuery{config: sdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysdictionary.Table, sysdictionary.FieldID, selector),
			sqlgraph.To(sysdictionarydetail.Table, sysdictionarydetail.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sysdictionary.DetailsTable, sysdictionary.DetailsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysDictionary entity from the query.
// Returns a *NotFoundError when no SysDictionary was found.
func (sdq *SysDictionaryQuery) First(ctx context.Context) (*SysDictionary, error) {
	nodes, err := sdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysdictionary.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sdq *SysDictionaryQuery) FirstX(ctx context.Context) *SysDictionary {
	node, err := sdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysDictionary ID from the query.
// Returns a *NotFoundError when no SysDictionary ID was found.
func (sdq *SysDictionaryQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = sdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysdictionary.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sdq *SysDictionaryQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := sdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysDictionary entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysDictionary entity is found.
// Returns a *NotFoundError when no SysDictionary entities are found.
func (sdq *SysDictionaryQuery) Only(ctx context.Context) (*SysDictionary, error) {
	nodes, err := sdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysdictionary.Label}
	default:
		return nil, &NotSingularError{sysdictionary.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sdq *SysDictionaryQuery) OnlyX(ctx context.Context) *SysDictionary {
	node, err := sdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysDictionary ID in the query.
// Returns a *NotSingularError when more than one SysDictionary ID is found.
// Returns a *NotFoundError when no entities are found.
func (sdq *SysDictionaryQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = sdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysdictionary.Label}
	default:
		err = &NotSingularError{sysdictionary.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sdq *SysDictionaryQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := sdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysDictionaries.
func (sdq *SysDictionaryQuery) All(ctx context.Context) ([]*SysDictionary, error) {
	if err := sdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sdq *SysDictionaryQuery) AllX(ctx context.Context) []*SysDictionary {
	nodes, err := sdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysDictionary IDs.
func (sdq *SysDictionaryQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := sdq.Select(sysdictionary.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sdq *SysDictionaryQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := sdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sdq *SysDictionaryQuery) Count(ctx context.Context) (int, error) {
	if err := sdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sdq *SysDictionaryQuery) CountX(ctx context.Context) int {
	count, err := sdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sdq *SysDictionaryQuery) Exist(ctx context.Context) (bool, error) {
	if err := sdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sdq *SysDictionaryQuery) ExistX(ctx context.Context) bool {
	exist, err := sdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysDictionaryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sdq *SysDictionaryQuery) Clone() *SysDictionaryQuery {
	if sdq == nil {
		return nil
	}
	return &SysDictionaryQuery{
		config:      sdq.config,
		limit:       sdq.limit,
		offset:      sdq.offset,
		order:       append([]OrderFunc{}, sdq.order...),
		predicates:  append([]predicate.SysDictionary{}, sdq.predicates...),
		withDetails: sdq.withDetails.Clone(),
		// clone intermediate query.
		sql:    sdq.sql.Clone(),
		path:   sdq.path,
		unique: sdq.unique,
	}
}

// WithDetails tells the query-builder to eager-load the nodes that are connected to
// the "details" edge. The optional arguments are used to configure the query builder of the edge.
func (sdq *SysDictionaryQuery) WithDetails(opts ...func(*SysDictionaryDetailQuery)) *SysDictionaryQuery {
	query := &SysDictionaryDetailQuery{config: sdq.config}
	for _, opt := range opts {
		opt(query)
	}
	sdq.withDetails = query
	return sdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysDictionary.Query().
//		GroupBy(sysdictionary.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sdq *SysDictionaryQuery) GroupBy(field string, fields ...string) *SysDictionaryGroupBy {
	grbuild := &SysDictionaryGroupBy{config: sdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sdq.sqlQuery(ctx), nil
	}
	grbuild.label = sysdictionary.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.SysDictionary.Query().
//		Select(sysdictionary.FieldTitle).
//		Scan(ctx, &v)
func (sdq *SysDictionaryQuery) Select(fields ...string) *SysDictionarySelect {
	sdq.fields = append(sdq.fields, fields...)
	selbuild := &SysDictionarySelect{SysDictionaryQuery: sdq}
	selbuild.label = sysdictionary.Label
	selbuild.flds, selbuild.scan = &sdq.fields, selbuild.Scan
	return selbuild
}

func (sdq *SysDictionaryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sdq.fields {
		if !sysdictionary.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sdq.path != nil {
		prev, err := sdq.path(ctx)
		if err != nil {
			return err
		}
		sdq.sql = prev
	}
	return nil
}

func (sdq *SysDictionaryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysDictionary, error) {
	var (
		nodes       = []*SysDictionary{}
		_spec       = sdq.querySpec()
		loadedTypes = [1]bool{
			sdq.withDetails != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysDictionary).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysDictionary{config: sdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sdq.modifiers) > 0 {
		_spec.Modifiers = sdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sdq.withDetails; query != nil {
		if err := sdq.loadDetails(ctx, query, nodes,
			func(n *SysDictionary) { n.Edges.Details = []*SysDictionaryDetail{} },
			func(n *SysDictionary, e *SysDictionaryDetail) { n.Edges.Details = append(n.Edges.Details, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sdq *SysDictionaryQuery) loadDetails(ctx context.Context, query *SysDictionaryDetailQuery, nodes []*SysDictionary, init func(*SysDictionary), assign func(*SysDictionary, *SysDictionaryDetail)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*SysDictionary)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.SysDictionaryDetail(func(s *sql.Selector) {
		s.Where(sql.InValues(sysdictionary.DetailsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DictionaryID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dictionary_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sdq *SysDictionaryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sdq.querySpec()
	if len(sdq.modifiers) > 0 {
		_spec.Modifiers = sdq.modifiers
	}
	_spec.Node.Columns = sdq.fields
	if len(sdq.fields) > 0 {
		_spec.Unique = sdq.unique != nil && *sdq.unique
	}
	return sqlgraph.CountNodes(ctx, sdq.driver, _spec)
}

func (sdq *SysDictionaryQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sdq *SysDictionaryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdictionary.Table,
			Columns: sysdictionary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysdictionary.FieldID,
			},
		},
		From:   sdq.sql,
		Unique: true,
	}
	if unique := sdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdictionary.FieldID)
		for i := range fields {
			if fields[i] != sysdictionary.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sdq *SysDictionaryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sdq.driver.Dialect())
	t1 := builder.Table(sysdictionary.Table)
	columns := sdq.fields
	if len(columns) == 0 {
		columns = sysdictionary.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sdq.sql != nil {
		selector = sdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sdq.unique != nil && *sdq.unique {
		selector.Distinct()
	}
	for _, m := range sdq.modifiers {
		m(selector)
	}
	for _, p := range sdq.predicates {
		p(selector)
	}
	for _, p := range sdq.order {
		p(selector)
	}
	if offset := sdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sdq *SysDictionaryQuery) Modify(modifiers ...func(s *sql.Selector)) *SysDictionarySelect {
	sdq.modifiers = append(sdq.modifiers, modifiers...)
	return sdq.Select()
}

// SysDictionaryGroupBy is the group-by builder for SysDictionary entities.
type SysDictionaryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sdgb *SysDictionaryGroupBy) Aggregate(fns ...AggregateFunc) *SysDictionaryGroupBy {
	sdgb.fns = append(sdgb.fns, fns...)
	return sdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sdgb *SysDictionaryGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sdgb.path(ctx)
	if err != nil {
		return err
	}
	sdgb.sql = query
	return sdgb.sqlScan(ctx, v)
}

func (sdgb *SysDictionaryGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sdgb.fields {
		if !sysdictionary.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sdgb *SysDictionaryGroupBy) sqlQuery() *sql.Selector {
	selector := sdgb.sql.Select()
	aggregation := make([]string, 0, len(sdgb.fns))
	for _, fn := range sdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sdgb.fields)+len(sdgb.fns))
		for _, f := range sdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sdgb.fields...)...)
}

// SysDictionarySelect is the builder for selecting fields of SysDictionary entities.
type SysDictionarySelect struct {
	*SysDictionaryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sds *SysDictionarySelect) Scan(ctx context.Context, v any) error {
	if err := sds.prepareQuery(ctx); err != nil {
		return err
	}
	sds.sql = sds.SysDictionaryQuery.sqlQuery(ctx)
	return sds.sqlScan(ctx, v)
}

func (sds *SysDictionarySelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := sds.sql.Query()
	if err := sds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sds *SysDictionarySelect) Modify(modifiers ...func(s *sql.Selector)) *SysDictionarySelect {
	sds.modifiers = append(sds.modifiers, modifiers...)
	return sds
}
