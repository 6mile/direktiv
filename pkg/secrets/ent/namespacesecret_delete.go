// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/secrets/ent/namespacesecret"
	"github.com/direktiv/direktiv/pkg/secrets/ent/predicate"
)

// NamespaceSecretDelete is the builder for deleting a NamespaceSecret entity.
type NamespaceSecretDelete struct {
	config
	hooks    []Hook
	mutation *NamespaceSecretMutation
}

// Where appends a list predicates to the NamespaceSecretDelete builder.
func (nsd *NamespaceSecretDelete) Where(ps ...predicate.NamespaceSecret) *NamespaceSecretDelete {
	nsd.mutation.Where(ps...)
	return nsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nsd *NamespaceSecretDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nsd.hooks) == 0 {
		affected, err = nsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NamespaceSecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nsd.mutation = mutation
			affected, err = nsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nsd.hooks) - 1; i >= 0; i-- {
			if nsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsd *NamespaceSecretDelete) ExecX(ctx context.Context) int {
	n, err := nsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nsd *NamespaceSecretDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: namespacesecret.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: namespacesecret.FieldID,
			},
		},
	}
	if ps := nsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, nsd.driver, _spec)
}

// NamespaceSecretDeleteOne is the builder for deleting a single NamespaceSecret entity.
type NamespaceSecretDeleteOne struct {
	nsd *NamespaceSecretDelete
}

// Exec executes the deletion query.
func (nsdo *NamespaceSecretDeleteOne) Exec(ctx context.Context) error {
	n, err := nsdo.nsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{namespacesecret.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nsdo *NamespaceSecretDeleteOne) ExecX(ctx context.Context) {
	nsdo.nsd.ExecX(ctx)
}
