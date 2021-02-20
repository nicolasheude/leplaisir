// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"dechild/ent/predicate"
	"dechild/ent/stockmanager"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StockManagerQuery is the builder for querying StockManager entities.
type StockManagerQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.StockManager
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StockManagerQuery builder.
func (smq *StockManagerQuery) Where(ps ...predicate.StockManager) *StockManagerQuery {
	smq.predicates = append(smq.predicates, ps...)
	return smq
}

// Limit adds a limit step to the query.
func (smq *StockManagerQuery) Limit(limit int) *StockManagerQuery {
	smq.limit = &limit
	return smq
}

// Offset adds an offset step to the query.
func (smq *StockManagerQuery) Offset(offset int) *StockManagerQuery {
	smq.offset = &offset
	return smq
}

// Order adds an order step to the query.
func (smq *StockManagerQuery) Order(o ...OrderFunc) *StockManagerQuery {
	smq.order = append(smq.order, o...)
	return smq
}

// First returns the first StockManager entity from the query.
// Returns a *NotFoundError when no StockManager was found.
func (smq *StockManagerQuery) First(ctx context.Context) (*StockManager, error) {
	nodes, err := smq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{stockmanager.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smq *StockManagerQuery) FirstX(ctx context.Context) *StockManager {
	node, err := smq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StockManager ID from the query.
// Returns a *NotFoundError when no StockManager ID was found.
func (smq *StockManagerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = smq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{stockmanager.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smq *StockManagerQuery) FirstIDX(ctx context.Context) int {
	id, err := smq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StockManager entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one StockManager entity is not found.
// Returns a *NotFoundError when no StockManager entities are found.
func (smq *StockManagerQuery) Only(ctx context.Context) (*StockManager, error) {
	nodes, err := smq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{stockmanager.Label}
	default:
		return nil, &NotSingularError{stockmanager.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smq *StockManagerQuery) OnlyX(ctx context.Context) *StockManager {
	node, err := smq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StockManager ID in the query.
// Returns a *NotSingularError when exactly one StockManager ID is not found.
// Returns a *NotFoundError when no entities are found.
func (smq *StockManagerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = smq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{stockmanager.Label}
	default:
		err = &NotSingularError{stockmanager.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smq *StockManagerQuery) OnlyIDX(ctx context.Context) int {
	id, err := smq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StockManagers.
func (smq *StockManagerQuery) All(ctx context.Context) ([]*StockManager, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return smq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (smq *StockManagerQuery) AllX(ctx context.Context) []*StockManager {
	nodes, err := smq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StockManager IDs.
func (smq *StockManagerQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := smq.Select(stockmanager.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smq *StockManagerQuery) IDsX(ctx context.Context) []int {
	ids, err := smq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smq *StockManagerQuery) Count(ctx context.Context) (int, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return smq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (smq *StockManagerQuery) CountX(ctx context.Context) int {
	count, err := smq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smq *StockManagerQuery) Exist(ctx context.Context) (bool, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return smq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (smq *StockManagerQuery) ExistX(ctx context.Context) bool {
	exist, err := smq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StockManagerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smq *StockManagerQuery) Clone() *StockManagerQuery {
	if smq == nil {
		return nil
	}
	return &StockManagerQuery{
		config:     smq.config,
		limit:      smq.limit,
		offset:     smq.offset,
		order:      append([]OrderFunc{}, smq.order...),
		predicates: append([]predicate.StockManager{}, smq.predicates...),
		// clone intermediate query.
		sql:  smq.sql.Clone(),
		path: smq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Activite string `json:"Activite,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StockManager.Query().
//		GroupBy(stockmanager.FieldActivite).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (smq *StockManagerQuery) GroupBy(field string, fields ...string) *StockManagerGroupBy {
	group := &StockManagerGroupBy{config: smq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return smq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Activite string `json:"Activite,omitempty"`
//	}
//
//	client.StockManager.Query().
//		Select(stockmanager.FieldActivite).
//		Scan(ctx, &v)
//
func (smq *StockManagerQuery) Select(field string, fields ...string) *StockManagerSelect {
	smq.fields = append([]string{field}, fields...)
	return &StockManagerSelect{StockManagerQuery: smq}
}

func (smq *StockManagerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range smq.fields {
		if !stockmanager.ValidColumn(f) {
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

func (smq *StockManagerQuery) sqlAll(ctx context.Context) ([]*StockManager, error) {
	var (
		nodes = []*StockManager{}
		_spec = smq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &StockManager{config: smq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, smq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (smq *StockManagerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smq.querySpec()
	return sqlgraph.CountNodes(ctx, smq.driver, _spec)
}

func (smq *StockManagerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := smq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (smq *StockManagerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stockmanager.Table,
			Columns: stockmanager.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stockmanager.FieldID,
			},
		},
		From:   smq.sql,
		Unique: true,
	}
	if fields := smq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, stockmanager.FieldID)
		for i := range fields {
			if fields[i] != stockmanager.FieldID {
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
				ps[i](selector, stockmanager.ValidColumn)
			}
		}
	}
	return _spec
}

func (smq *StockManagerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smq.driver.Dialect())
	t1 := builder.Table(stockmanager.Table)
	selector := builder.Select(t1.Columns(stockmanager.Columns...)...).From(t1)
	if smq.sql != nil {
		selector = smq.sql
		selector.Select(selector.Columns(stockmanager.Columns...)...)
	}
	for _, p := range smq.predicates {
		p(selector)
	}
	for _, p := range smq.order {
		p(selector, stockmanager.ValidColumn)
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

// StockManagerGroupBy is the group-by builder for StockManager entities.
type StockManagerGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smgb *StockManagerGroupBy) Aggregate(fns ...AggregateFunc) *StockManagerGroupBy {
	smgb.fns = append(smgb.fns, fns...)
	return smgb
}

// Scan applies the group-by query and scans the result into the given value.
func (smgb *StockManagerGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := smgb.path(ctx)
	if err != nil {
		return err
	}
	smgb.sql = query
	return smgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (smgb *StockManagerGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := smgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *StockManagerGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("ent: StockManagerGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (smgb *StockManagerGroupBy) StringsX(ctx context.Context) []string {
	v, err := smgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *StockManagerGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = smgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stockmanager.Label}
	default:
		err = fmt.Errorf("ent: StockManagerGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (smgb *StockManagerGroupBy) StringX(ctx context.Context) string {
	v, err := smgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *StockManagerGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("ent: StockManagerGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (smgb *StockManagerGroupBy) IntsX(ctx context.Context) []int {
	v, err := smgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *StockManagerGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = smgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stockmanager.Label}
	default:
		err = fmt.Errorf("ent: StockManagerGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (smgb *StockManagerGroupBy) IntX(ctx context.Context) int {
	v, err := smgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *StockManagerGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("ent: StockManagerGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (smgb *StockManagerGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := smgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *StockManagerGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = smgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stockmanager.Label}
	default:
		err = fmt.Errorf("ent: StockManagerGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (smgb *StockManagerGroupBy) Float64X(ctx context.Context) float64 {
	v, err := smgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *StockManagerGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("ent: StockManagerGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (smgb *StockManagerGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := smgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *StockManagerGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = smgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stockmanager.Label}
	default:
		err = fmt.Errorf("ent: StockManagerGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (smgb *StockManagerGroupBy) BoolX(ctx context.Context) bool {
	v, err := smgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (smgb *StockManagerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range smgb.fields {
		if !stockmanager.ValidColumn(f) {
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

func (smgb *StockManagerGroupBy) sqlQuery() *sql.Selector {
	selector := smgb.sql
	columns := make([]string, 0, len(smgb.fields)+len(smgb.fns))
	columns = append(columns, smgb.fields...)
	for _, fn := range smgb.fns {
		columns = append(columns, fn(selector, stockmanager.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(smgb.fields...)
}

// StockManagerSelect is the builder for selecting fields of StockManager entities.
type StockManagerSelect struct {
	*StockManagerQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sms *StockManagerSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sms.prepareQuery(ctx); err != nil {
		return err
	}
	sms.sql = sms.StockManagerQuery.sqlQuery(ctx)
	return sms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sms *StockManagerSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sms *StockManagerSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("ent: StockManagerSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sms *StockManagerSelect) StringsX(ctx context.Context) []string {
	v, err := sms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sms *StockManagerSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stockmanager.Label}
	default:
		err = fmt.Errorf("ent: StockManagerSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sms *StockManagerSelect) StringX(ctx context.Context) string {
	v, err := sms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sms *StockManagerSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("ent: StockManagerSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sms *StockManagerSelect) IntsX(ctx context.Context) []int {
	v, err := sms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sms *StockManagerSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stockmanager.Label}
	default:
		err = fmt.Errorf("ent: StockManagerSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sms *StockManagerSelect) IntX(ctx context.Context) int {
	v, err := sms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sms *StockManagerSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("ent: StockManagerSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sms *StockManagerSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sms *StockManagerSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stockmanager.Label}
	default:
		err = fmt.Errorf("ent: StockManagerSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sms *StockManagerSelect) Float64X(ctx context.Context) float64 {
	v, err := sms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sms *StockManagerSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("ent: StockManagerSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sms *StockManagerSelect) BoolsX(ctx context.Context) []bool {
	v, err := sms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sms *StockManagerSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stockmanager.Label}
	default:
		err = fmt.Errorf("ent: StockManagerSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sms *StockManagerSelect) BoolX(ctx context.Context) bool {
	v, err := sms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sms *StockManagerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sms.sqlQuery().Query()
	if err := sms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sms *StockManagerSelect) sqlQuery() sql.Querier {
	selector := sms.sql
	selector.Select(selector.Columns(sms.fields...)...)
	return selector
}
