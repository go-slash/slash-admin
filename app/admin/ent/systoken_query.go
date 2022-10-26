// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/systoken"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysTokenQuery is the builder for querying SysToken entities.
type SysTokenQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysToken
	loadTotal  []func(context.Context, []*SysToken) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysTokenQuery builder.
func (stq *SysTokenQuery) Where(ps ...predicate.SysToken) *SysTokenQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit adds a limit step to the query.
func (stq *SysTokenQuery) Limit(limit int) *SysTokenQuery {
	stq.limit = &limit
	return stq
}

// Offset adds an offset step to the query.
func (stq *SysTokenQuery) Offset(offset int) *SysTokenQuery {
	stq.offset = &offset
	return stq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (stq *SysTokenQuery) Unique(unique bool) *SysTokenQuery {
	stq.unique = &unique
	return stq
}

// Order adds an order step to the query.
func (stq *SysTokenQuery) Order(o ...OrderFunc) *SysTokenQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// First returns the first SysToken entity from the query.
// Returns a *NotFoundError when no SysToken was found.
func (stq *SysTokenQuery) First(ctx context.Context) (*SysToken, error) {
	nodes, err := stq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{systoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *SysTokenQuery) FirstX(ctx context.Context) *SysToken {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysToken ID from the query.
// Returns a *NotFoundError when no SysToken ID was found.
func (stq *SysTokenQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = stq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{systoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *SysTokenQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysToken entity is found.
// Returns a *NotFoundError when no SysToken entities are found.
func (stq *SysTokenQuery) Only(ctx context.Context) (*SysToken, error) {
	nodes, err := stq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{systoken.Label}
	default:
		return nil, &NotSingularError{systoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *SysTokenQuery) OnlyX(ctx context.Context) *SysToken {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysToken ID in the query.
// Returns a *NotSingularError when more than one SysToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (stq *SysTokenQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = stq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{systoken.Label}
	default:
		err = &NotSingularError{systoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *SysTokenQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysTokens.
func (stq *SysTokenQuery) All(ctx context.Context) ([]*SysToken, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return stq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (stq *SysTokenQuery) AllX(ctx context.Context) []*SysToken {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysToken IDs.
func (stq *SysTokenQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := stq.Select(systoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *SysTokenQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *SysTokenQuery) Count(ctx context.Context) (int, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return stq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (stq *SysTokenQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *SysTokenQuery) Exist(ctx context.Context) (bool, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return stq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *SysTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *SysTokenQuery) Clone() *SysTokenQuery {
	if stq == nil {
		return nil
	}
	return &SysTokenQuery{
		config:     stq.config,
		limit:      stq.limit,
		offset:     stq.offset,
		order:      append([]OrderFunc{}, stq.order...),
		predicates: append([]predicate.SysToken{}, stq.predicates...),
		// clone intermediate query.
		sql:    stq.sql.Clone(),
		path:   stq.path,
		unique: stq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysToken.Query().
//		GroupBy(systoken.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (stq *SysTokenQuery) GroupBy(field string, fields ...string) *SysTokenGroupBy {
	grbuild := &SysTokenGroupBy{config: stq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return stq.sqlQuery(ctx), nil
	}
	grbuild.label = systoken.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid,omitempty"`
//	}
//
//	client.SysToken.Query().
//		Select(systoken.FieldUUID).
//		Scan(ctx, &v)
func (stq *SysTokenQuery) Select(fields ...string) *SysTokenSelect {
	stq.fields = append(stq.fields, fields...)
	selbuild := &SysTokenSelect{SysTokenQuery: stq}
	selbuild.label = systoken.Label
	selbuild.flds, selbuild.scan = &stq.fields, selbuild.Scan
	return selbuild
}

func (stq *SysTokenQuery) prepareQuery(ctx context.Context) error {
	for _, f := range stq.fields {
		if !systoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	return nil
}

func (stq *SysTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysToken, error) {
	var (
		nodes = []*SysToken{}
		_spec = stq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysToken{config: stq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(stq.modifiers) > 0 {
		_spec.Modifiers = stq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range stq.loadTotal {
		if err := stq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (stq *SysTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	if len(stq.modifiers) > 0 {
		_spec.Modifiers = stq.modifiers
	}
	_spec.Node.Columns = stq.fields
	if len(stq.fields) > 0 {
		_spec.Unique = stq.unique != nil && *stq.unique
	}
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *SysTokenQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := stq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (stq *SysTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   systoken.Table,
			Columns: systoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: systoken.FieldID,
			},
		},
		From:   stq.sql,
		Unique: true,
	}
	if unique := stq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := stq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systoken.FieldID)
		for i := range fields {
			if fields[i] != systoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (stq *SysTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(systoken.Table)
	columns := stq.fields
	if len(columns) == 0 {
		columns = systoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if stq.unique != nil && *stq.unique {
		selector.Distinct()
	}
	for _, m := range stq.modifiers {
		m(selector)
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector)
	}
	if offset := stq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (stq *SysTokenQuery) Modify(modifiers ...func(s *sql.Selector)) *SysTokenSelect {
	stq.modifiers = append(stq.modifiers, modifiers...)
	return stq.Select()
}

// SysTokenGroupBy is the group-by builder for SysToken entities.
type SysTokenGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *SysTokenGroupBy) Aggregate(fns ...AggregateFunc) *SysTokenGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the group-by query and scans the result into the given value.
func (stgb *SysTokenGroupBy) Scan(ctx context.Context, v any) error {
	query, err := stgb.path(ctx)
	if err != nil {
		return err
	}
	stgb.sql = query
	return stgb.sqlScan(ctx, v)
}

func (stgb *SysTokenGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range stgb.fields {
		if !systoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := stgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (stgb *SysTokenGroupBy) sqlQuery() *sql.Selector {
	selector := stgb.sql.Select()
	aggregation := make([]string, 0, len(stgb.fns))
	for _, fn := range stgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(stgb.fields)+len(stgb.fns))
		for _, f := range stgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(stgb.fields...)...)
}

// SysTokenSelect is the builder for selecting fields of SysToken entities.
type SysTokenSelect struct {
	*SysTokenQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sts *SysTokenSelect) Scan(ctx context.Context, v any) error {
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	sts.sql = sts.SysTokenQuery.sqlQuery(ctx)
	return sts.sqlScan(ctx, v)
}

func (sts *SysTokenSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := sts.sql.Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sts *SysTokenSelect) Modify(modifiers ...func(s *sql.Selector)) *SysTokenSelect {
	sts.modifiers = append(sts.modifiers, modifiers...)
	return sts
}