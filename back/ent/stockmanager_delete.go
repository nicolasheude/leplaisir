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

// StockManagerDelete is the builder for deleting a StockManager entity.
type StockManagerDelete struct {
	config
	hooks    []Hook
	mutation *StockManagerMutation
}

// Where adds a new predicate to the StockManagerDelete builder.
func (smd *StockManagerDelete) Where(ps ...predicate.StockManager) *StockManagerDelete {
	smd.mutation.predicates = append(smd.mutation.predicates, ps...)
	return smd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (smd *StockManagerDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(smd.hooks) == 0 {
		affected, err = smd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockManagerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smd.mutation = mutation
			affected, err = smd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smd.hooks) - 1; i >= 0; i-- {
			mut = smd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (smd *StockManagerDelete) ExecX(ctx context.Context) int {
	n, err := smd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (smd *StockManagerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: stockmanager.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stockmanager.FieldID,
			},
		},
	}
	if ps := smd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, smd.driver, _spec)
}

// StockManagerDeleteOne is the builder for deleting a single StockManager entity.
type StockManagerDeleteOne struct {
	smd *StockManagerDelete
}

// Exec executes the deletion query.
func (smdo *StockManagerDeleteOne) Exec(ctx context.Context) error {
	n, err := smdo.smd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{stockmanager.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (smdo *StockManagerDeleteOne) ExecX(ctx context.Context) {
	smdo.smd.ExecX(ctx)
}
