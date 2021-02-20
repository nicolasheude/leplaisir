// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"dechild/ent/contactparents"
	"dechild/ent/form"
	"dechild/ent/predicate"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FormQuery is the builder for querying Form entities.
type FormQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Form
	// eager-loading edges.
	withContactParents *ContactParentsQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FormQuery builder.
func (fq *FormQuery) Where(ps ...predicate.Form) *FormQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FormQuery) Limit(limit int) *FormQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FormQuery) Offset(offset int) *FormQuery {
	fq.offset = &offset
	return fq
}

// Order adds an order step to the query.
func (fq *FormQuery) Order(o ...OrderFunc) *FormQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryContactParents chains the current query on the "contactParents" edge.
func (fq *FormQuery) QueryContactParents() *ContactParentsQuery {
	query := &ContactParentsQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(form.Table, form.FieldID, selector),
			sqlgraph.To(contactparents.Table, contactparents.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, form.ContactParentsTable, form.ContactParentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Form entity from the query.
// Returns a *NotFoundError when no Form was found.
func (fq *FormQuery) First(ctx context.Context) (*Form, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{form.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FormQuery) FirstX(ctx context.Context) *Form {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Form ID from the query.
// Returns a *NotFoundError when no Form ID was found.
func (fq *FormQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{form.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FormQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Form entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Form entity is not found.
// Returns a *NotFoundError when no Form entities are found.
func (fq *FormQuery) Only(ctx context.Context) (*Form, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{form.Label}
	default:
		return nil, &NotSingularError{form.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FormQuery) OnlyX(ctx context.Context) *Form {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Form ID in the query.
// Returns a *NotSingularError when exactly one Form ID is not found.
// Returns a *NotFoundError when no entities are found.
func (fq *FormQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{form.Label}
	default:
		err = &NotSingularError{form.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FormQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Forms.
func (fq *FormQuery) All(ctx context.Context) ([]*Form, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FormQuery) AllX(ctx context.Context) []*Form {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Form IDs.
func (fq *FormQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fq.Select(form.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FormQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FormQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FormQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FormQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FormQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FormQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FormQuery) Clone() *FormQuery {
	if fq == nil {
		return nil
	}
	return &FormQuery{
		config:             fq.config,
		limit:              fq.limit,
		offset:             fq.offset,
		order:              append([]OrderFunc{}, fq.order...),
		predicates:         append([]predicate.Form{}, fq.predicates...),
		withContactParents: fq.withContactParents.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithContactParents tells the query-builder to eager-load the nodes that are connected to
// the "contactParents" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FormQuery) WithContactParents(opts ...func(*ContactParentsQuery)) *FormQuery {
	query := &ContactParentsQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withContactParents = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Email string `json:"Email,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Form.Query().
//		GroupBy(form.FieldEmail).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fq *FormQuery) GroupBy(field string, fields ...string) *FormGroupBy {
	group := &FormGroupBy{config: fq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Email string `json:"Email,omitempty"`
//	}
//
//	client.Form.Query().
//		Select(form.FieldEmail).
//		Scan(ctx, &v)
//
func (fq *FormQuery) Select(field string, fields ...string) *FormSelect {
	fq.fields = append([]string{field}, fields...)
	return &FormSelect{FormQuery: fq}
}

func (fq *FormQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fq.fields {
		if !form.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FormQuery) sqlAll(ctx context.Context) ([]*Form, error) {
	var (
		nodes       = []*Form{}
		withFKs     = fq.withFKs
		_spec       = fq.querySpec()
		loadedTypes = [1]bool{
			fq.withContactParents != nil,
		}
	)
	if fq.withContactParents != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, form.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Form{config: fq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fq.withContactParents; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Form)
		for i := range nodes {
			if fk := nodes[i].form_contact_parents; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(contactparents.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "form_contact_parents" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ContactParents = n
			}
		}
	}

	return nodes, nil
}

func (fq *FormQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FormQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (fq *FormQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   form.Table,
			Columns: form.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: form.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, form.FieldID)
		for i := range fields {
			if fields[i] != form.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, form.ValidColumn)
			}
		}
	}
	return _spec
}

func (fq *FormQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(form.Table)
	selector := builder.Select(t1.Columns(form.Columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(form.Columns...)...)
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector, form.ValidColumn)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FormGroupBy is the group-by builder for Form entities.
type FormGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FormGroupBy) Aggregate(fns ...AggregateFunc) *FormGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgb *FormGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fgb *FormGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FormGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FormGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fgb *FormGroupBy) StringsX(ctx context.Context) []string {
	v, err := fgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FormGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{form.Label}
	default:
		err = fmt.Errorf("ent: FormGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fgb *FormGroupBy) StringX(ctx context.Context) string {
	v, err := fgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FormGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FormGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fgb *FormGroupBy) IntsX(ctx context.Context) []int {
	v, err := fgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FormGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{form.Label}
	default:
		err = fmt.Errorf("ent: FormGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fgb *FormGroupBy) IntX(ctx context.Context) int {
	v, err := fgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FormGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FormGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fgb *FormGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FormGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{form.Label}
	default:
		err = fmt.Errorf("ent: FormGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fgb *FormGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FormGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FormGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fgb *FormGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FormGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{form.Label}
	default:
		err = fmt.Errorf("ent: FormGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fgb *FormGroupBy) BoolX(ctx context.Context) bool {
	v, err := fgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fgb *FormGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fgb.fields {
		if !form.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FormGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql
	columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
	columns = append(columns, fgb.fields...)
	for _, fn := range fgb.fns {
		columns = append(columns, fn(selector, form.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(fgb.fields...)
}

// FormSelect is the builder for selecting fields of Form entities.
type FormSelect struct {
	*FormQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FormSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	fs.sql = fs.FormQuery.sqlQuery(ctx)
	return fs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fs *FormSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fs *FormSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FormSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fs *FormSelect) StringsX(ctx context.Context) []string {
	v, err := fs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fs *FormSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{form.Label}
	default:
		err = fmt.Errorf("ent: FormSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fs *FormSelect) StringX(ctx context.Context) string {
	v, err := fs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fs *FormSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FormSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fs *FormSelect) IntsX(ctx context.Context) []int {
	v, err := fs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fs *FormSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{form.Label}
	default:
		err = fmt.Errorf("ent: FormSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fs *FormSelect) IntX(ctx context.Context) int {
	v, err := fs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fs *FormSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FormSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fs *FormSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fs *FormSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{form.Label}
	default:
		err = fmt.Errorf("ent: FormSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fs *FormSelect) Float64X(ctx context.Context) float64 {
	v, err := fs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fs *FormSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FormSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fs *FormSelect) BoolsX(ctx context.Context) []bool {
	v, err := fs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fs *FormSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{form.Label}
	default:
		err = fmt.Errorf("ent: FormSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fs *FormSelect) BoolX(ctx context.Context) bool {
	v, err := fs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fs *FormSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fs.sqlQuery().Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fs *FormSelect) sqlQuery() sql.Querier {
	selector := fs.sql
	selector.Select(selector.Columns(fs.fields...)...)
	return selector
}
