// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysmenu"
	"slash-admin/app/admin/ent/sysrole"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysMenuQuery is the builder for querying SysMenu entities.
type SysMenuQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.SysMenu
	withRoles    *SysRoleQuery
	withParent   *SysMenuQuery
	withChildren *SysMenuQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysMenuQuery builder.
func (smq *SysMenuQuery) Where(ps ...predicate.SysMenu) *SysMenuQuery {
	smq.predicates = append(smq.predicates, ps...)
	return smq
}

// Limit adds a limit step to the query.
func (smq *SysMenuQuery) Limit(limit int) *SysMenuQuery {
	smq.limit = &limit
	return smq
}

// Offset adds an offset step to the query.
func (smq *SysMenuQuery) Offset(offset int) *SysMenuQuery {
	smq.offset = &offset
	return smq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smq *SysMenuQuery) Unique(unique bool) *SysMenuQuery {
	smq.unique = &unique
	return smq
}

// Order adds an order step to the query.
func (smq *SysMenuQuery) Order(o ...OrderFunc) *SysMenuQuery {
	smq.order = append(smq.order, o...)
	return smq
}

// QueryRoles chains the current query on the "roles" edge.
func (smq *SysMenuQuery) QueryRoles() *SysRoleQuery {
	query := &SysRoleQuery{config: smq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysmenu.Table, sysmenu.FieldID, selector),
			sqlgraph.To(sysrole.Table, sysrole.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, sysmenu.RolesTable, sysmenu.RolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(smq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (smq *SysMenuQuery) QueryParent() *SysMenuQuery {
	query := &SysMenuQuery{config: smq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysmenu.Table, sysmenu.FieldID, selector),
			sqlgraph.To(sysmenu.Table, sysmenu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sysmenu.ParentTable, sysmenu.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(smq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (smq *SysMenuQuery) QueryChildren() *SysMenuQuery {
	query := &SysMenuQuery{config: smq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysmenu.Table, sysmenu.FieldID, selector),
			sqlgraph.To(sysmenu.Table, sysmenu.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sysmenu.ChildrenTable, sysmenu.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(smq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysMenu entity from the query.
// Returns a *NotFoundError when no SysMenu was found.
func (smq *SysMenuQuery) First(ctx context.Context) (*SysMenu, error) {
	nodes, err := smq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysmenu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smq *SysMenuQuery) FirstX(ctx context.Context) *SysMenu {
	node, err := smq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysMenu ID from the query.
// Returns a *NotFoundError when no SysMenu ID was found.
func (smq *SysMenuQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = smq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysmenu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smq *SysMenuQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := smq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysMenu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysMenu entity is found.
// Returns a *NotFoundError when no SysMenu entities are found.
func (smq *SysMenuQuery) Only(ctx context.Context) (*SysMenu, error) {
	nodes, err := smq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysmenu.Label}
	default:
		return nil, &NotSingularError{sysmenu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smq *SysMenuQuery) OnlyX(ctx context.Context) *SysMenu {
	node, err := smq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysMenu ID in the query.
// Returns a *NotSingularError when more than one SysMenu ID is found.
// Returns a *NotFoundError when no entities are found.
func (smq *SysMenuQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = smq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = &NotSingularError{sysmenu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smq *SysMenuQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := smq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysMenus.
func (smq *SysMenuQuery) All(ctx context.Context) ([]*SysMenu, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return smq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (smq *SysMenuQuery) AllX(ctx context.Context) []*SysMenu {
	nodes, err := smq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysMenu IDs.
func (smq *SysMenuQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := smq.Select(sysmenu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smq *SysMenuQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := smq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smq *SysMenuQuery) Count(ctx context.Context) (int, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return smq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (smq *SysMenuQuery) CountX(ctx context.Context) int {
	count, err := smq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smq *SysMenuQuery) Exist(ctx context.Context) (bool, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return smq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (smq *SysMenuQuery) ExistX(ctx context.Context) bool {
	exist, err := smq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysMenuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smq *SysMenuQuery) Clone() *SysMenuQuery {
	if smq == nil {
		return nil
	}
	return &SysMenuQuery{
		config:       smq.config,
		limit:        smq.limit,
		offset:       smq.offset,
		order:        append([]OrderFunc{}, smq.order...),
		predicates:   append([]predicate.SysMenu{}, smq.predicates...),
		withRoles:    smq.withRoles.Clone(),
		withParent:   smq.withParent.Clone(),
		withChildren: smq.withChildren.Clone(),
		// clone intermediate query.
		sql:    smq.sql.Clone(),
		path:   smq.path,
		unique: smq.unique,
	}
}

// WithRoles tells the query-builder to eager-load the nodes that are connected to
// the "roles" edge. The optional arguments are used to configure the query builder of the edge.
func (smq *SysMenuQuery) WithRoles(opts ...func(*SysRoleQuery)) *SysMenuQuery {
	query := &SysRoleQuery{config: smq.config}
	for _, opt := range opts {
		opt(query)
	}
	smq.withRoles = query
	return smq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (smq *SysMenuQuery) WithParent(opts ...func(*SysMenuQuery)) *SysMenuQuery {
	query := &SysMenuQuery{config: smq.config}
	for _, opt := range opts {
		opt(query)
	}
	smq.withParent = query
	return smq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (smq *SysMenuQuery) WithChildren(opts ...func(*SysMenuQuery)) *SysMenuQuery {
	query := &SysMenuQuery{config: smq.config}
	for _, opt := range opts {
		opt(query)
	}
	smq.withChildren = query
	return smq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ParentID uint64 `json:"parent_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysMenu.Query().
//		GroupBy(sysmenu.FieldParentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (smq *SysMenuQuery) GroupBy(field string, fields ...string) *SysMenuGroupBy {
	grbuild := &SysMenuGroupBy{config: smq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return smq.sqlQuery(ctx), nil
	}
	grbuild.label = sysmenu.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ParentID uint64 `json:"parent_id,omitempty"`
//	}
//
//	client.SysMenu.Query().
//		Select(sysmenu.FieldParentID).
//		Scan(ctx, &v)
func (smq *SysMenuQuery) Select(fields ...string) *SysMenuSelect {
	smq.fields = append(smq.fields, fields...)
	selbuild := &SysMenuSelect{SysMenuQuery: smq}
	selbuild.label = sysmenu.Label
	selbuild.flds, selbuild.scan = &smq.fields, selbuild.Scan
	return selbuild
}

func (smq *SysMenuQuery) prepareQuery(ctx context.Context) error {
	for _, f := range smq.fields {
		if !sysmenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if smq.path != nil {
		prev, err := smq.path(ctx)
		if err != nil {
			return err
		}
		smq.sql = prev
	}
	return nil
}

func (smq *SysMenuQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysMenu, error) {
	var (
		nodes       = []*SysMenu{}
		_spec       = smq.querySpec()
		loadedTypes = [3]bool{
			smq.withRoles != nil,
			smq.withParent != nil,
			smq.withChildren != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysMenu).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysMenu{config: smq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(smq.modifiers) > 0 {
		_spec.Modifiers = smq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, smq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := smq.withRoles; query != nil {
		if err := smq.loadRoles(ctx, query, nodes,
			func(n *SysMenu) { n.Edges.Roles = []*SysRole{} },
			func(n *SysMenu, e *SysRole) { n.Edges.Roles = append(n.Edges.Roles, e) }); err != nil {
			return nil, err
		}
	}
	if query := smq.withParent; query != nil {
		if err := smq.loadParent(ctx, query, nodes, nil,
			func(n *SysMenu, e *SysMenu) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := smq.withChildren; query != nil {
		if err := smq.loadChildren(ctx, query, nodes,
			func(n *SysMenu) { n.Edges.Children = []*SysMenu{} },
			func(n *SysMenu, e *SysMenu) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (smq *SysMenuQuery) loadRoles(ctx context.Context, query *SysRoleQuery, nodes []*SysMenu, init func(*SysMenu), assign func(*SysMenu, *SysRole)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint64]*SysMenu)
	nids := make(map[uint64]map[*SysMenu]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(sysmenu.RolesTable)
		s.Join(joinT).On(s.C(sysrole.FieldID), joinT.C(sysmenu.RolesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(sysmenu.RolesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(sysmenu.RolesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := uint64(values[0].(*sql.NullInt64).Int64)
			inValue := uint64(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*SysMenu]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "roles" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (smq *SysMenuQuery) loadParent(ctx context.Context, query *SysMenuQuery, nodes []*SysMenu, init func(*SysMenu), assign func(*SysMenu, *SysMenu)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*SysMenu)
	for i := range nodes {
		fk := nodes[i].ParentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(sysmenu.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (smq *SysMenuQuery) loadChildren(ctx context.Context, query *SysMenuQuery, nodes []*SysMenu, init func(*SysMenu), assign func(*SysMenu, *SysMenu)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*SysMenu)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.InValues(sysmenu.ChildrenColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ParentID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (smq *SysMenuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smq.querySpec()
	if len(smq.modifiers) > 0 {
		_spec.Modifiers = smq.modifiers
	}
	_spec.Node.Columns = smq.fields
	if len(smq.fields) > 0 {
		_spec.Unique = smq.unique != nil && *smq.unique
	}
	return sqlgraph.CountNodes(ctx, smq.driver, _spec)
}

func (smq *SysMenuQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := smq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (smq *SysMenuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysmenu.Table,
			Columns: sysmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysmenu.FieldID,
			},
		},
		From:   smq.sql,
		Unique: true,
	}
	if unique := smq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := smq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysmenu.FieldID)
		for i := range fields {
			if fields[i] != sysmenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := smq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := smq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := smq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (smq *SysMenuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smq.driver.Dialect())
	t1 := builder.Table(sysmenu.Table)
	columns := smq.fields
	if len(columns) == 0 {
		columns = sysmenu.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if smq.sql != nil {
		selector = smq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if smq.unique != nil && *smq.unique {
		selector.Distinct()
	}
	for _, m := range smq.modifiers {
		m(selector)
	}
	for _, p := range smq.predicates {
		p(selector)
	}
	for _, p := range smq.order {
		p(selector)
	}
	if offset := smq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (smq *SysMenuQuery) Modify(modifiers ...func(s *sql.Selector)) *SysMenuSelect {
	smq.modifiers = append(smq.modifiers, modifiers...)
	return smq.Select()
}

// SysMenuGroupBy is the group-by builder for SysMenu entities.
type SysMenuGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smgb *SysMenuGroupBy) Aggregate(fns ...AggregateFunc) *SysMenuGroupBy {
	smgb.fns = append(smgb.fns, fns...)
	return smgb
}

// Scan applies the group-by query and scans the result into the given value.
func (smgb *SysMenuGroupBy) Scan(ctx context.Context, v any) error {
	query, err := smgb.path(ctx)
	if err != nil {
		return err
	}
	smgb.sql = query
	return smgb.sqlScan(ctx, v)
}

func (smgb *SysMenuGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range smgb.fields {
		if !sysmenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := smgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (smgb *SysMenuGroupBy) sqlQuery() *sql.Selector {
	selector := smgb.sql.Select()
	aggregation := make([]string, 0, len(smgb.fns))
	for _, fn := range smgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(smgb.fields)+len(smgb.fns))
		for _, f := range smgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(smgb.fields...)...)
}

// SysMenuSelect is the builder for selecting fields of SysMenu entities.
type SysMenuSelect struct {
	*SysMenuQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sms *SysMenuSelect) Scan(ctx context.Context, v any) error {
	if err := sms.prepareQuery(ctx); err != nil {
		return err
	}
	sms.sql = sms.SysMenuQuery.sqlQuery(ctx)
	return sms.sqlScan(ctx, v)
}

func (sms *SysMenuSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := sms.sql.Query()
	if err := sms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sms *SysMenuSelect) Modify(modifiers ...func(s *sql.Selector)) *SysMenuSelect {
	sms.modifiers = append(sms.modifiers, modifiers...)
	return sms
}
