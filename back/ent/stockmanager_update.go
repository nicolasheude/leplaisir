// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"dechild/ent/predicate"
	"dechild/ent/stockmanager"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StockManagerUpdate is the builder for updating StockManager entities.
type StockManagerUpdate struct {
	config
	hooks    []Hook
	mutation *StockManagerMutation
}

// Where adds a new predicate for the StockManagerUpdate builder.
func (smu *StockManagerUpdate) Where(ps ...predicate.StockManager) *StockManagerUpdate {
	smu.mutation.predicates = append(smu.mutation.predicates, ps...)
	return smu
}

// Mutation returns the StockManagerMutation object of the builder.
func (smu *StockManagerUpdate) Mutation() *StockManagerMutation {
	return smu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (smu *StockManagerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(smu.hooks) == 0 {
		affected, err = smu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockManagerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smu.mutation = mutation
			affected, err = smu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smu.hooks) - 1; i >= 0; i-- {
			mut = smu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (smu *StockManagerUpdate) SaveX(ctx context.Context) int {
	affected, err := smu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (smu *StockManagerUpdate) Exec(ctx context.Context) error {
	_, err := smu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smu *StockManagerUpdate) ExecX(ctx context.Context) {
	if err := smu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (smu *StockManagerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stockmanager.Table,
			Columns: stockmanager.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stockmanager.FieldID,
			},
		},
	}
	if ps := smu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, smu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stockmanager.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StockManagerUpdateOne is the builder for updating a single StockManager entity.
type StockManagerUpdateOne struct {
	config
	hooks    []Hook
	mutation *StockManagerMutation
}

// Mutation returns the StockManagerMutation object of the builder.
func (smuo *StockManagerUpdateOne) Mutation() *StockManagerMutation {
	return smuo.mutation
}

// Save executes the query and returns the updated StockManager entity.
func (smuo *StockManagerUpdateOne) Save(ctx context.Context) (*StockManager, error) {
	var (
		err  error
		node *StockManager
	)
	if len(smuo.hooks) == 0 {
		node, err = smuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockManagerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smuo.mutation = mutation
			node, err = smuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(smuo.hooks) - 1; i >= 0; i-- {
			mut = smuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (smuo *StockManagerUpdateOne) SaveX(ctx context.Context) *StockManager {
	node, err := smuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (smuo *StockManagerUpdateOne) Exec(ctx context.Context) error {
	_, err := smuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smuo *StockManagerUpdateOne) ExecX(ctx context.Context) {
	if err := smuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (smuo *StockManagerUpdateOne) sqlSave(ctx context.Context) (_node *StockManager, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stockmanager.Table,
			Columns: stockmanager.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stockmanager.FieldID,
			},
		},
	}
	id, ok := smuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing StockManager.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := smuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &StockManager{config: smuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, smuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stockmanager.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
