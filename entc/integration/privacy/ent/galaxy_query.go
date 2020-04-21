// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/galaxy"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/planet"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// GalaxyQuery is the builder for querying Galaxy entities.
type GalaxyQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Galaxy
	// eager-loading edges.
	withPlanets *PlanetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (gq *GalaxyQuery) Where(ps ...predicate.Galaxy) *GalaxyQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GalaxyQuery) Limit(limit int) *GalaxyQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GalaxyQuery) Offset(offset int) *GalaxyQuery {
	gq.offset = &offset
	return gq
}

// Order adds an order step to the query.
func (gq *GalaxyQuery) Order(o ...Order) *GalaxyQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryPlanets chains the current query on the planets edge.
func (gq *GalaxyQuery) QueryPlanets() *PlanetQuery {
	query := &PlanetQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(galaxy.Table, galaxy.FieldID, gq.sqlQuery()),
			sqlgraph.To(planet.Table, planet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, galaxy.PlanetsTable, galaxy.PlanetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Galaxy entity in the query. Returns *NotFoundError when no galaxy was found.
func (gq *GalaxyQuery) First(ctx context.Context) (*Galaxy, error) {
	gas, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(gas) == 0 {
		return nil, &NotFoundError{galaxy.Label}
	}
	return gas[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GalaxyQuery) FirstX(ctx context.Context) *Galaxy {
	ga, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ga
}

// FirstID returns the first Galaxy id in the query. Returns *NotFoundError when no id was found.
func (gq *GalaxyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{galaxy.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (gq *GalaxyQuery) FirstXID(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Galaxy entity in the query, returns an error if not exactly one entity was returned.
func (gq *GalaxyQuery) Only(ctx context.Context) (*Galaxy, error) {
	gas, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(gas) {
	case 1:
		return gas[0], nil
	case 0:
		return nil, &NotFoundError{galaxy.Label}
	default:
		return nil, &NotSingularError{galaxy.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GalaxyQuery) OnlyX(ctx context.Context) *Galaxy {
	ga, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ga
}

// OnlyID returns the only Galaxy id in the query, returns an error if not exactly one id was returned.
func (gq *GalaxyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{galaxy.Label}
	default:
		err = &NotSingularError{galaxy.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (gq *GalaxyQuery) OnlyXID(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Galaxies.
func (gq *GalaxyQuery) All(ctx context.Context) ([]*Galaxy, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GalaxyQuery) AllX(ctx context.Context) []*Galaxy {
	gas, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return gas
}

// IDs executes the query and returns a list of Galaxy ids.
func (gq *GalaxyQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gq.Select(galaxy.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GalaxyQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GalaxyQuery) Count(ctx context.Context) (int, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GalaxyQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GalaxyQuery) Exist(ctx context.Context) (bool, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GalaxyQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GalaxyQuery) Clone() *GalaxyQuery {
	return &GalaxyQuery{
		config:     gq.config,
		limit:      gq.limit,
		offset:     gq.offset,
		order:      append([]Order{}, gq.order...),
		unique:     append([]string{}, gq.unique...),
		predicates: append([]predicate.Galaxy{}, gq.predicates...),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

//  WithPlanets tells the query-builder to eager-loads the nodes that are connected to
// the "planets" edge. The optional arguments used to configure the query builder of the edge.
func (gq *GalaxyQuery) WithPlanets(opts ...func(*PlanetQuery)) *GalaxyQuery {
	query := &PlanetQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withPlanets = query
	return gq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Galaxy.Query().
//		GroupBy(galaxy.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gq *GalaxyQuery) GroupBy(field string, fields ...string) *GalaxyGroupBy {
	group := &GalaxyGroupBy{config: gq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Galaxy.Query().
//		Select(galaxy.FieldName).
//		Scan(ctx, &v)
//
func (gq *GalaxyQuery) Select(field string, fields ...string) *GalaxySelect {
	selector := &GalaxySelect{config: gq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.sqlQuery(), nil
	}
	return selector
}

func (gq *GalaxyQuery) prepareQuery(ctx context.Context) error {
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	if err := galaxy.Policy.EvalQuery(ctx, gq); err != nil {
		return err
	}
	return nil
}

func (gq *GalaxyQuery) sqlAll(ctx context.Context) ([]*Galaxy, error) {
	var (
		nodes       = []*Galaxy{}
		_spec       = gq.querySpec()
		loadedTypes = [1]bool{
			gq.withPlanets != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Galaxy{config: gq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gq.withPlanets; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Galaxy)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Planet(func(s *sql.Selector) {
			s.Where(sql.InValues(galaxy.PlanetsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.galaxy_planets
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "galaxy_planets" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "galaxy_planets" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Planets = append(node.Edges.Planets, n)
		}
	}

	return nodes, nil
}

func (gq *GalaxyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GalaxyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (gq *GalaxyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   galaxy.Table,
			Columns: galaxy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: galaxy.FieldID,
			},
		},
		From:   gq.sql,
		Unique: true,
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GalaxyQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(galaxy.Table)
	selector := builder.Select(t1.Columns(galaxy.Columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(galaxy.Columns...)...)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GalaxyGroupBy is the builder for group-by Galaxy entities.
type GalaxyGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GalaxyGroupBy) Aggregate(fns ...Aggregate) *GalaxyGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scan the result into the given value.
func (ggb *GalaxyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ggb.path(ctx)
	if err != nil {
		return err
	}
	ggb.sql = query
	return ggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ggb *GalaxyGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ggb *GalaxyGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GalaxyGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ggb *GalaxyGroupBy) StringsX(ctx context.Context) []string {
	v, err := ggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ggb *GalaxyGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GalaxyGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ggb *GalaxyGroupBy) IntsX(ctx context.Context) []int {
	v, err := ggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ggb *GalaxyGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GalaxyGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ggb *GalaxyGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ggb *GalaxyGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GalaxyGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ggb *GalaxyGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ggb *GalaxyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ggb.sqlQuery().Query()
	if err := ggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ggb *GalaxyGroupBy) sqlQuery() *sql.Selector {
	selector := ggb.sql
	columns := make([]string, 0, len(ggb.fields)+len(ggb.fns))
	columns = append(columns, ggb.fields...)
	for _, fn := range ggb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ggb.fields...)
}

// GalaxySelect is the builder for select fields of Galaxy entities.
type GalaxySelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (gs *GalaxySelect) Scan(ctx context.Context, v interface{}) error {
	query, err := gs.path(ctx)
	if err != nil {
		return err
	}
	gs.sql = query
	return gs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gs *GalaxySelect) ScanX(ctx context.Context, v interface{}) {
	if err := gs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (gs *GalaxySelect) Strings(ctx context.Context) ([]string, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GalaxySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gs *GalaxySelect) StringsX(ctx context.Context) []string {
	v, err := gs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (gs *GalaxySelect) Ints(ctx context.Context) ([]int, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GalaxySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gs *GalaxySelect) IntsX(ctx context.Context) []int {
	v, err := gs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (gs *GalaxySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GalaxySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gs *GalaxySelect) Float64sX(ctx context.Context) []float64 {
	v, err := gs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (gs *GalaxySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GalaxySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gs *GalaxySelect) BoolsX(ctx context.Context) []bool {
	v, err := gs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gs *GalaxySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gs.sqlQuery().Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gs *GalaxySelect) sqlQuery() sql.Querier {
	selector := gs.sql
	selector.Select(selector.Columns(gs.fields...)...)
	return selector
}