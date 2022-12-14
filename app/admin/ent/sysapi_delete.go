// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysapi"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysApiDelete is the builder for deleting a SysApi entity.
type SysApiDelete struct {
	config
	hooks    []Hook
	mutation *SysApiMutation
}

// Where appends a list predicates to the SysApiDelete builder.
func (sad *SysApiDelete) Where(ps ...predicate.SysApi) *SysApiDelete {
	sad.mutation.Where(ps...)
	return sad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sad *SysApiDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sad.hooks) == 0 {
		affected, err = sad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysApiMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sad.mutation = mutation
			affected, err = sad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sad.hooks) - 1; i >= 0; i-- {
			if sad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sad *SysApiDelete) ExecX(ctx context.Context) int {
	n, err := sad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sad *SysApiDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sysapi.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysapi.FieldID,
			},
		},
	}
	if ps := sad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SysApiDeleteOne is the builder for deleting a single SysApi entity.
type SysApiDeleteOne struct {
	sad *SysApiDelete
}

// Exec executes the deletion query.
func (sado *SysApiDeleteOne) Exec(ctx context.Context) error {
	n, err := sado.sad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysapi.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sado *SysApiDeleteOne) ExecX(ctx context.Context) {
	sado.sad.ExecX(ctx)
}
