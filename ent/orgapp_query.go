// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/workflow/ent/orgapp"
	"github.com/woocoos/workflow/ent/predicate"

	"github.com/woocoos/workflow/ent/internal"
)

// OrgAppQuery is the builder for querying OrgApp entities.
type OrgAppQuery struct {
	config
	ctx        *QueryContext
	order      []orgapp.OrderOption
	inters     []Interceptor
	predicates []predicate.OrgApp
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*OrgApp) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrgAppQuery builder.
func (oaq *OrgAppQuery) Where(ps ...predicate.OrgApp) *OrgAppQuery {
	oaq.predicates = append(oaq.predicates, ps...)
	return oaq
}

// Limit the number of records to be returned by this query.
func (oaq *OrgAppQuery) Limit(limit int) *OrgAppQuery {
	oaq.ctx.Limit = &limit
	return oaq
}

// Offset to start from.
func (oaq *OrgAppQuery) Offset(offset int) *OrgAppQuery {
	oaq.ctx.Offset = &offset
	return oaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oaq *OrgAppQuery) Unique(unique bool) *OrgAppQuery {
	oaq.ctx.Unique = &unique
	return oaq
}

// Order specifies how the records should be ordered.
func (oaq *OrgAppQuery) Order(o ...orgapp.OrderOption) *OrgAppQuery {
	oaq.order = append(oaq.order, o...)
	return oaq
}

// First returns the first OrgApp entity from the query.
// Returns a *NotFoundError when no OrgApp was found.
func (oaq *OrgAppQuery) First(ctx context.Context) (*OrgApp, error) {
	nodes, err := oaq.Limit(1).All(setContextOp(ctx, oaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orgapp.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oaq *OrgAppQuery) FirstX(ctx context.Context) *OrgApp {
	node, err := oaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrgApp ID from the query.
// Returns a *NotFoundError when no OrgApp ID was found.
func (oaq *OrgAppQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oaq.Limit(1).IDs(setContextOp(ctx, oaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orgapp.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oaq *OrgAppQuery) FirstIDX(ctx context.Context) int {
	id, err := oaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrgApp entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrgApp entity is found.
// Returns a *NotFoundError when no OrgApp entities are found.
func (oaq *OrgAppQuery) Only(ctx context.Context) (*OrgApp, error) {
	nodes, err := oaq.Limit(2).All(setContextOp(ctx, oaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orgapp.Label}
	default:
		return nil, &NotSingularError{orgapp.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oaq *OrgAppQuery) OnlyX(ctx context.Context) *OrgApp {
	node, err := oaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrgApp ID in the query.
// Returns a *NotSingularError when more than one OrgApp ID is found.
// Returns a *NotFoundError when no entities are found.
func (oaq *OrgAppQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oaq.Limit(2).IDs(setContextOp(ctx, oaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orgapp.Label}
	default:
		err = &NotSingularError{orgapp.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oaq *OrgAppQuery) OnlyIDX(ctx context.Context) int {
	id, err := oaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrgApps.
func (oaq *OrgAppQuery) All(ctx context.Context) ([]*OrgApp, error) {
	ctx = setContextOp(ctx, oaq.ctx, "All")
	if err := oaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrgApp, *OrgAppQuery]()
	return withInterceptors[[]*OrgApp](ctx, oaq, qr, oaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oaq *OrgAppQuery) AllX(ctx context.Context) []*OrgApp {
	nodes, err := oaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrgApp IDs.
func (oaq *OrgAppQuery) IDs(ctx context.Context) (ids []int, err error) {
	if oaq.ctx.Unique == nil && oaq.path != nil {
		oaq.Unique(true)
	}
	ctx = setContextOp(ctx, oaq.ctx, "IDs")
	if err = oaq.Select(orgapp.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oaq *OrgAppQuery) IDsX(ctx context.Context) []int {
	ids, err := oaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oaq *OrgAppQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oaq.ctx, "Count")
	if err := oaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oaq, querierCount[*OrgAppQuery](), oaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oaq *OrgAppQuery) CountX(ctx context.Context) int {
	count, err := oaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oaq *OrgAppQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oaq.ctx, "Exist")
	switch _, err := oaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oaq *OrgAppQuery) ExistX(ctx context.Context) bool {
	exist, err := oaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrgAppQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oaq *OrgAppQuery) Clone() *OrgAppQuery {
	if oaq == nil {
		return nil
	}
	return &OrgAppQuery{
		config:     oaq.config,
		ctx:        oaq.ctx.Clone(),
		order:      append([]orgapp.OrderOption{}, oaq.order...),
		inters:     append([]Interceptor{}, oaq.inters...),
		predicates: append([]predicate.OrgApp{}, oaq.predicates...),
		// clone intermediate query.
		sql:  oaq.sql.Clone(),
		path: oaq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OrgID int `json:"org_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrgApp.Query().
//		GroupBy(orgapp.FieldOrgID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oaq *OrgAppQuery) GroupBy(field string, fields ...string) *OrgAppGroupBy {
	oaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrgAppGroupBy{build: oaq}
	grbuild.flds = &oaq.ctx.Fields
	grbuild.label = orgapp.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OrgID int `json:"org_id,omitempty"`
//	}
//
//	client.OrgApp.Query().
//		Select(orgapp.FieldOrgID).
//		Scan(ctx, &v)
func (oaq *OrgAppQuery) Select(fields ...string) *OrgAppSelect {
	oaq.ctx.Fields = append(oaq.ctx.Fields, fields...)
	sbuild := &OrgAppSelect{OrgAppQuery: oaq}
	sbuild.label = orgapp.Label
	sbuild.flds, sbuild.scan = &oaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrgAppSelect configured with the given aggregations.
func (oaq *OrgAppQuery) Aggregate(fns ...AggregateFunc) *OrgAppSelect {
	return oaq.Select().Aggregate(fns...)
}

func (oaq *OrgAppQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oaq); err != nil {
				return err
			}
		}
	}
	for _, f := range oaq.ctx.Fields {
		if !orgapp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oaq.path != nil {
		prev, err := oaq.path(ctx)
		if err != nil {
			return err
		}
		oaq.sql = prev
	}
	return nil
}

func (oaq *OrgAppQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrgApp, error) {
	var (
		nodes = []*OrgApp{}
		_spec = oaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrgApp).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrgApp{config: oaq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = oaq.schemaConfig.OrgApp
	ctx = internal.NewSchemaConfigContext(ctx, oaq.schemaConfig)
	if len(oaq.modifiers) > 0 {
		_spec.Modifiers = oaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range oaq.loadTotal {
		if err := oaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oaq *OrgAppQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oaq.querySpec()
	_spec.Node.Schema = oaq.schemaConfig.OrgApp
	ctx = internal.NewSchemaConfigContext(ctx, oaq.schemaConfig)
	if len(oaq.modifiers) > 0 {
		_spec.Modifiers = oaq.modifiers
	}
	_spec.Node.Columns = oaq.ctx.Fields
	if len(oaq.ctx.Fields) > 0 {
		_spec.Unique = oaq.ctx.Unique != nil && *oaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oaq.driver, _spec)
}

func (oaq *OrgAppQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orgapp.Table, orgapp.Columns, sqlgraph.NewFieldSpec(orgapp.FieldID, field.TypeInt))
	_spec.From = oaq.sql
	if unique := oaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oaq.path != nil {
		_spec.Unique = true
	}
	if fields := oaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgapp.FieldID)
		for i := range fields {
			if fields[i] != orgapp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oaq *OrgAppQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oaq.driver.Dialect())
	t1 := builder.Table(orgapp.Table)
	columns := oaq.ctx.Fields
	if len(columns) == 0 {
		columns = orgapp.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oaq.sql != nil {
		selector = oaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oaq.ctx.Unique != nil && *oaq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(oaq.schemaConfig.OrgApp)
	ctx = internal.NewSchemaConfigContext(ctx, oaq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range oaq.predicates {
		p(selector)
	}
	for _, p := range oaq.order {
		p(selector)
	}
	if offset := oaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrgAppGroupBy is the group-by builder for OrgApp entities.
type OrgAppGroupBy struct {
	selector
	build *OrgAppQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oagb *OrgAppGroupBy) Aggregate(fns ...AggregateFunc) *OrgAppGroupBy {
	oagb.fns = append(oagb.fns, fns...)
	return oagb
}

// Scan applies the selector query and scans the result into the given value.
func (oagb *OrgAppGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oagb.build.ctx, "GroupBy")
	if err := oagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgAppQuery, *OrgAppGroupBy](ctx, oagb.build, oagb, oagb.build.inters, v)
}

func (oagb *OrgAppGroupBy) sqlScan(ctx context.Context, root *OrgAppQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(oagb.fns))
	for _, fn := range oagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*oagb.flds)+len(oagb.fns))
		for _, f := range *oagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*oagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrgAppSelect is the builder for selecting fields of OrgApp entities.
type OrgAppSelect struct {
	*OrgAppQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oas *OrgAppSelect) Aggregate(fns ...AggregateFunc) *OrgAppSelect {
	oas.fns = append(oas.fns, fns...)
	return oas
}

// Scan applies the selector query and scans the result into the given value.
func (oas *OrgAppSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oas.ctx, "Select")
	if err := oas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgAppQuery, *OrgAppSelect](ctx, oas.OrgAppQuery, oas, oas.inters, v)
}

func (oas *OrgAppSelect) sqlScan(ctx context.Context, root *OrgAppQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oas.fns))
	for _, fn := range oas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
