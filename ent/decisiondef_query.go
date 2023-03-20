// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/workflow/ent/decisiondef"
	"github.com/woocoos/workflow/ent/decisionreqdef"
	"github.com/woocoos/workflow/ent/predicate"
)

// DecisionDefQuery is the builder for querying DecisionDef entities.
type DecisionDefQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.DecisionDef
	withReqDef *DecisionReqDefQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*DecisionDef) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DecisionDefQuery builder.
func (ddq *DecisionDefQuery) Where(ps ...predicate.DecisionDef) *DecisionDefQuery {
	ddq.predicates = append(ddq.predicates, ps...)
	return ddq
}

// Limit the number of records to be returned by this query.
func (ddq *DecisionDefQuery) Limit(limit int) *DecisionDefQuery {
	ddq.ctx.Limit = &limit
	return ddq
}

// Offset to start from.
func (ddq *DecisionDefQuery) Offset(offset int) *DecisionDefQuery {
	ddq.ctx.Offset = &offset
	return ddq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ddq *DecisionDefQuery) Unique(unique bool) *DecisionDefQuery {
	ddq.ctx.Unique = &unique
	return ddq
}

// Order specifies how the records should be ordered.
func (ddq *DecisionDefQuery) Order(o ...OrderFunc) *DecisionDefQuery {
	ddq.order = append(ddq.order, o...)
	return ddq
}

// QueryReqDef chains the current query on the "req_def" edge.
func (ddq *DecisionDefQuery) QueryReqDef() *DecisionReqDefQuery {
	query := (&DecisionReqDefClient{config: ddq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ddq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ddq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(decisiondef.Table, decisiondef.FieldID, selector),
			sqlgraph.To(decisionreqdef.Table, decisionreqdef.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, decisiondef.ReqDefTable, decisiondef.ReqDefColumn),
		)
		fromU = sqlgraph.SetNeighbors(ddq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DecisionDef entity from the query.
// Returns a *NotFoundError when no DecisionDef was found.
func (ddq *DecisionDefQuery) First(ctx context.Context) (*DecisionDef, error) {
	nodes, err := ddq.Limit(1).All(setContextOp(ctx, ddq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{decisiondef.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ddq *DecisionDefQuery) FirstX(ctx context.Context) *DecisionDef {
	node, err := ddq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DecisionDef ID from the query.
// Returns a *NotFoundError when no DecisionDef ID was found.
func (ddq *DecisionDefQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ddq.Limit(1).IDs(setContextOp(ctx, ddq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{decisiondef.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ddq *DecisionDefQuery) FirstIDX(ctx context.Context) int {
	id, err := ddq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DecisionDef entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DecisionDef entity is found.
// Returns a *NotFoundError when no DecisionDef entities are found.
func (ddq *DecisionDefQuery) Only(ctx context.Context) (*DecisionDef, error) {
	nodes, err := ddq.Limit(2).All(setContextOp(ctx, ddq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{decisiondef.Label}
	default:
		return nil, &NotSingularError{decisiondef.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ddq *DecisionDefQuery) OnlyX(ctx context.Context) *DecisionDef {
	node, err := ddq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DecisionDef ID in the query.
// Returns a *NotSingularError when more than one DecisionDef ID is found.
// Returns a *NotFoundError when no entities are found.
func (ddq *DecisionDefQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ddq.Limit(2).IDs(setContextOp(ctx, ddq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{decisiondef.Label}
	default:
		err = &NotSingularError{decisiondef.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ddq *DecisionDefQuery) OnlyIDX(ctx context.Context) int {
	id, err := ddq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DecisionDefs.
func (ddq *DecisionDefQuery) All(ctx context.Context) ([]*DecisionDef, error) {
	ctx = setContextOp(ctx, ddq.ctx, "All")
	if err := ddq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DecisionDef, *DecisionDefQuery]()
	return withInterceptors[[]*DecisionDef](ctx, ddq, qr, ddq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ddq *DecisionDefQuery) AllX(ctx context.Context) []*DecisionDef {
	nodes, err := ddq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DecisionDef IDs.
func (ddq *DecisionDefQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ddq.ctx.Unique == nil && ddq.path != nil {
		ddq.Unique(true)
	}
	ctx = setContextOp(ctx, ddq.ctx, "IDs")
	if err = ddq.Select(decisiondef.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ddq *DecisionDefQuery) IDsX(ctx context.Context) []int {
	ids, err := ddq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ddq *DecisionDefQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ddq.ctx, "Count")
	if err := ddq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ddq, querierCount[*DecisionDefQuery](), ddq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ddq *DecisionDefQuery) CountX(ctx context.Context) int {
	count, err := ddq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ddq *DecisionDefQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ddq.ctx, "Exist")
	switch _, err := ddq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ddq *DecisionDefQuery) ExistX(ctx context.Context) bool {
	exist, err := ddq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DecisionDefQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ddq *DecisionDefQuery) Clone() *DecisionDefQuery {
	if ddq == nil {
		return nil
	}
	return &DecisionDefQuery{
		config:     ddq.config,
		ctx:        ddq.ctx.Clone(),
		order:      append([]OrderFunc{}, ddq.order...),
		inters:     append([]Interceptor{}, ddq.inters...),
		predicates: append([]predicate.DecisionDef{}, ddq.predicates...),
		withReqDef: ddq.withReqDef.Clone(),
		// clone intermediate query.
		sql:  ddq.sql.Clone(),
		path: ddq.path,
	}
}

// WithReqDef tells the query-builder to eager-load the nodes that are connected to
// the "req_def" edge. The optional arguments are used to configure the query builder of the edge.
func (ddq *DecisionDefQuery) WithReqDef(opts ...func(*DecisionReqDefQuery)) *DecisionDefQuery {
	query := (&DecisionReqDefClient{config: ddq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ddq.withReqDef = query
	return ddq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int `json:"created_by,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DecisionDef.Query().
//		GroupBy(decisiondef.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ddq *DecisionDefQuery) GroupBy(field string, fields ...string) *DecisionDefGroupBy {
	ddq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DecisionDefGroupBy{build: ddq}
	grbuild.flds = &ddq.ctx.Fields
	grbuild.label = decisiondef.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int `json:"created_by,omitempty"`
//	}
//
//	client.DecisionDef.Query().
//		Select(decisiondef.FieldCreatedBy).
//		Scan(ctx, &v)
func (ddq *DecisionDefQuery) Select(fields ...string) *DecisionDefSelect {
	ddq.ctx.Fields = append(ddq.ctx.Fields, fields...)
	sbuild := &DecisionDefSelect{DecisionDefQuery: ddq}
	sbuild.label = decisiondef.Label
	sbuild.flds, sbuild.scan = &ddq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DecisionDefSelect configured with the given aggregations.
func (ddq *DecisionDefQuery) Aggregate(fns ...AggregateFunc) *DecisionDefSelect {
	return ddq.Select().Aggregate(fns...)
}

func (ddq *DecisionDefQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ddq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ddq); err != nil {
				return err
			}
		}
	}
	for _, f := range ddq.ctx.Fields {
		if !decisiondef.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ddq.path != nil {
		prev, err := ddq.path(ctx)
		if err != nil {
			return err
		}
		ddq.sql = prev
	}
	return nil
}

func (ddq *DecisionDefQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DecisionDef, error) {
	var (
		nodes       = []*DecisionDef{}
		_spec       = ddq.querySpec()
		loadedTypes = [1]bool{
			ddq.withReqDef != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DecisionDef).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DecisionDef{config: ddq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ddq.modifiers) > 0 {
		_spec.Modifiers = ddq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ddq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ddq.withReqDef; query != nil {
		if err := ddq.loadReqDef(ctx, query, nodes, nil,
			func(n *DecisionDef, e *DecisionReqDef) { n.Edges.ReqDef = e }); err != nil {
			return nil, err
		}
	}
	for i := range ddq.loadTotal {
		if err := ddq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ddq *DecisionDefQuery) loadReqDef(ctx context.Context, query *DecisionReqDefQuery, nodes []*DecisionDef, init func(*DecisionDef), assign func(*DecisionDef, *DecisionReqDef)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DecisionDef)
	for i := range nodes {
		fk := nodes[i].ReqDefID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(decisionreqdef.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "req_def_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ddq *DecisionDefQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ddq.querySpec()
	if len(ddq.modifiers) > 0 {
		_spec.Modifiers = ddq.modifiers
	}
	_spec.Node.Columns = ddq.ctx.Fields
	if len(ddq.ctx.Fields) > 0 {
		_spec.Unique = ddq.ctx.Unique != nil && *ddq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ddq.driver, _spec)
}

func (ddq *DecisionDefQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(decisiondef.Table, decisiondef.Columns, sqlgraph.NewFieldSpec(decisiondef.FieldID, field.TypeInt))
	_spec.From = ddq.sql
	if unique := ddq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ddq.path != nil {
		_spec.Unique = true
	}
	if fields := ddq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, decisiondef.FieldID)
		for i := range fields {
			if fields[i] != decisiondef.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ddq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ddq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ddq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ddq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ddq *DecisionDefQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ddq.driver.Dialect())
	t1 := builder.Table(decisiondef.Table)
	columns := ddq.ctx.Fields
	if len(columns) == 0 {
		columns = decisiondef.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ddq.sql != nil {
		selector = ddq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ddq.ctx.Unique != nil && *ddq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ddq.predicates {
		p(selector)
	}
	for _, p := range ddq.order {
		p(selector)
	}
	if offset := ddq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ddq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DecisionDefGroupBy is the group-by builder for DecisionDef entities.
type DecisionDefGroupBy struct {
	selector
	build *DecisionDefQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ddgb *DecisionDefGroupBy) Aggregate(fns ...AggregateFunc) *DecisionDefGroupBy {
	ddgb.fns = append(ddgb.fns, fns...)
	return ddgb
}

// Scan applies the selector query and scans the result into the given value.
func (ddgb *DecisionDefGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ddgb.build.ctx, "GroupBy")
	if err := ddgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DecisionDefQuery, *DecisionDefGroupBy](ctx, ddgb.build, ddgb, ddgb.build.inters, v)
}

func (ddgb *DecisionDefGroupBy) sqlScan(ctx context.Context, root *DecisionDefQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ddgb.fns))
	for _, fn := range ddgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ddgb.flds)+len(ddgb.fns))
		for _, f := range *ddgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ddgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ddgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DecisionDefSelect is the builder for selecting fields of DecisionDef entities.
type DecisionDefSelect struct {
	*DecisionDefQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dds *DecisionDefSelect) Aggregate(fns ...AggregateFunc) *DecisionDefSelect {
	dds.fns = append(dds.fns, fns...)
	return dds
}

// Scan applies the selector query and scans the result into the given value.
func (dds *DecisionDefSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dds.ctx, "Select")
	if err := dds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DecisionDefQuery, *DecisionDefSelect](ctx, dds.DecisionDefQuery, dds, dds.inters, v)
}

func (dds *DecisionDefSelect) sqlScan(ctx context.Context, root *DecisionDefQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dds.fns))
	for _, fn := range dds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}