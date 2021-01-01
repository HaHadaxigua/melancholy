// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/HaHadaxigua/melancholy/ent/folder"
	"github.com/HaHadaxigua/melancholy/ent/mfile"
	"github.com/HaHadaxigua/melancholy/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// FolderQuery is the builder for querying Folder entities.
type FolderQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Folder
	// eager-loading edges.
	withMfiles *MFileQuery
	withP      *FolderQuery
	withC      *FolderQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (fq *FolderQuery) Where(ps ...predicate.Folder) *FolderQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FolderQuery) Limit(limit int) *FolderQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FolderQuery) Offset(offset int) *FolderQuery {
	fq.offset = &offset
	return fq
}

// Order adds an order step to the query.
func (fq *FolderQuery) Order(o ...OrderFunc) *FolderQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryMfiles chains the current query on the mfiles edge.
func (fq *FolderQuery) QueryMfiles() *MFileQuery {
	query := &MFileQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(folder.Table, folder.FieldID, selector),
			sqlgraph.To(mfile.Table, mfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, folder.MfilesTable, folder.MfilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryP chains the current query on the p edge.
func (fq *FolderQuery) QueryP() *FolderQuery {
	query := &FolderQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(folder.Table, folder.FieldID, selector),
			sqlgraph.To(folder.Table, folder.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, folder.PTable, folder.PColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryC chains the current query on the c edge.
func (fq *FolderQuery) QueryC() *FolderQuery {
	query := &FolderQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(folder.Table, folder.FieldID, selector),
			sqlgraph.To(folder.Table, folder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, folder.CTable, folder.CColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Folder entity in the query. Returns *NotFoundError when no folder was found.
func (fq *FolderQuery) First(ctx context.Context) (*Folder, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{folder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FolderQuery) FirstX(ctx context.Context) *Folder {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Folder id in the query. Returns *NotFoundError when no id was found.
func (fq *FolderQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{folder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FolderQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Folder entity in the query, returns an error if not exactly one entity was returned.
func (fq *FolderQuery) Only(ctx context.Context) (*Folder, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{folder.Label}
	default:
		return nil, &NotSingularError{folder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FolderQuery) OnlyX(ctx context.Context) *Folder {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only Folder id in the query, returns an error if not exactly one id was returned.
func (fq *FolderQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{folder.Label}
	default:
		err = &NotSingularError{folder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FolderQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Folders.
func (fq *FolderQuery) All(ctx context.Context) ([]*Folder, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FolderQuery) AllX(ctx context.Context) []*Folder {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Folder ids.
func (fq *FolderQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fq.Select(folder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FolderQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FolderQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FolderQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FolderQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FolderQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FolderQuery) Clone() *FolderQuery {
	if fq == nil {
		return nil
	}
	return &FolderQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]OrderFunc{}, fq.order...),
		unique:     append([]string{}, fq.unique...),
		predicates: append([]predicate.Folder{}, fq.predicates...),
		withMfiles: fq.withMfiles.Clone(),
		withP:      fq.withP.Clone(),
		withC:      fq.withC.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

//  WithMfiles tells the query-builder to eager-loads the nodes that are connected to
// the "mfiles" edge. The optional arguments used to configure the query builder of the edge.
func (fq *FolderQuery) WithMfiles(opts ...func(*MFileQuery)) *FolderQuery {
	query := &MFileQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withMfiles = query
	return fq
}

//  WithP tells the query-builder to eager-loads the nodes that are connected to
// the "p" edge. The optional arguments used to configure the query builder of the edge.
func (fq *FolderQuery) WithP(opts ...func(*FolderQuery)) *FolderQuery {
	query := &FolderQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withP = query
	return fq
}

//  WithC tells the query-builder to eager-loads the nodes that are connected to
// the "c" edge. The optional arguments used to configure the query builder of the edge.
func (fq *FolderQuery) WithC(opts ...func(*FolderQuery)) *FolderQuery {
	query := &FolderQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withC = query
	return fq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Parent int `json:"parent,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Folder.Query().
//		GroupBy(folder.FieldParent).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fq *FolderQuery) GroupBy(field string, fields ...string) *FolderGroupBy {
	group := &FolderGroupBy{config: fq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Parent int `json:"parent,omitempty"`
//	}
//
//	client.Folder.Query().
//		Select(folder.FieldParent).
//		Scan(ctx, &v)
//
func (fq *FolderQuery) Select(field string, fields ...string) *FolderSelect {
	selector := &FolderSelect{config: fq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(), nil
	}
	return selector
}

func (fq *FolderQuery) prepareQuery(ctx context.Context) error {
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FolderQuery) sqlAll(ctx context.Context) ([]*Folder, error) {
	var (
		nodes       = []*Folder{}
		withFKs     = fq.withFKs
		_spec       = fq.querySpec()
		loadedTypes = [3]bool{
			fq.withMfiles != nil,
			fq.withP != nil,
			fq.withC != nil,
		}
	)
	if fq.withP != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, folder.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Folder{config: fq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
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
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fq.withMfiles; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Folder)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Mfiles = []*MFile{}
		}
		query.withFKs = true
		query.Where(predicate.MFile(func(s *sql.Selector) {
			s.Where(sql.InValues(folder.MfilesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.folder_mfiles
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "folder_mfiles" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "folder_mfiles" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Mfiles = append(node.Edges.Mfiles, n)
		}
	}

	if query := fq.withP; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Folder)
		for i := range nodes {
			if fk := nodes[i].folder_c; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(folder.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "folder_c" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.P = n
			}
		}
	}

	if query := fq.withC; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Folder)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.C = []*Folder{}
		}
		query.withFKs = true
		query.Where(predicate.Folder(func(s *sql.Selector) {
			s.Where(sql.InValues(folder.CColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.folder_c
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "folder_c" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "folder_c" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.C = append(node.Edges.C, n)
		}
	}

	return nodes, nil
}

func (fq *FolderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FolderQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (fq *FolderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   folder.Table,
			Columns: folder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: folder.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
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
				ps[i](selector, folder.ValidColumn)
			}
		}
	}
	return _spec
}

func (fq *FolderQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(folder.Table)
	selector := builder.Select(t1.Columns(folder.Columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(folder.Columns...)...)
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector, folder.ValidColumn)
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

// FolderGroupBy is the builder for group-by Folder entities.
type FolderGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FolderGroupBy) Aggregate(fns ...AggregateFunc) *FolderGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scan the result into the given value.
func (fgb *FolderGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fgb *FolderGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (fgb *FolderGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FolderGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fgb *FolderGroupBy) StringsX(ctx context.Context) []string {
	v, err := fgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (fgb *FolderGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{folder.Label}
	default:
		err = fmt.Errorf("ent: FolderGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fgb *FolderGroupBy) StringX(ctx context.Context) string {
	v, err := fgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (fgb *FolderGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FolderGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fgb *FolderGroupBy) IntsX(ctx context.Context) []int {
	v, err := fgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (fgb *FolderGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{folder.Label}
	default:
		err = fmt.Errorf("ent: FolderGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fgb *FolderGroupBy) IntX(ctx context.Context) int {
	v, err := fgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (fgb *FolderGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FolderGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fgb *FolderGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (fgb *FolderGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{folder.Label}
	default:
		err = fmt.Errorf("ent: FolderGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fgb *FolderGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (fgb *FolderGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FolderGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fgb *FolderGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (fgb *FolderGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{folder.Label}
	default:
		err = fmt.Errorf("ent: FolderGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fgb *FolderGroupBy) BoolX(ctx context.Context) bool {
	v, err := fgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fgb *FolderGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fgb.fields {
		if !folder.ValidColumn(f) {
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

func (fgb *FolderGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql
	columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
	columns = append(columns, fgb.fields...)
	for _, fn := range fgb.fns {
		columns = append(columns, fn(selector, folder.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(fgb.fields...)
}

// FolderSelect is the builder for select fields of Folder entities.
type FolderSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (fs *FolderSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := fs.path(ctx)
	if err != nil {
		return err
	}
	fs.sql = query
	return fs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fs *FolderSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (fs *FolderSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FolderSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fs *FolderSelect) StringsX(ctx context.Context) []string {
	v, err := fs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (fs *FolderSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{folder.Label}
	default:
		err = fmt.Errorf("ent: FolderSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fs *FolderSelect) StringX(ctx context.Context) string {
	v, err := fs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (fs *FolderSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FolderSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fs *FolderSelect) IntsX(ctx context.Context) []int {
	v, err := fs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (fs *FolderSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{folder.Label}
	default:
		err = fmt.Errorf("ent: FolderSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fs *FolderSelect) IntX(ctx context.Context) int {
	v, err := fs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (fs *FolderSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FolderSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fs *FolderSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (fs *FolderSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{folder.Label}
	default:
		err = fmt.Errorf("ent: FolderSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fs *FolderSelect) Float64X(ctx context.Context) float64 {
	v, err := fs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (fs *FolderSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FolderSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fs *FolderSelect) BoolsX(ctx context.Context) []bool {
	v, err := fs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (fs *FolderSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{folder.Label}
	default:
		err = fmt.Errorf("ent: FolderSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fs *FolderSelect) BoolX(ctx context.Context) bool {
	v, err := fs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fs *FolderSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fs.fields {
		if !folder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := fs.sqlQuery().Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fs *FolderSelect) sqlQuery() sql.Querier {
	selector := fs.sql
	selector.Select(selector.Columns(fs.fields...)...)
	return selector
}