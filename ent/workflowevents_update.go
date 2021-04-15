// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/vorteil/direktiv/ent/predicate"
	"github.com/vorteil/direktiv/ent/workflow"
	"github.com/vorteil/direktiv/ent/workflowevents"
	"github.com/vorteil/direktiv/ent/workfloweventswait"
	"github.com/vorteil/direktiv/ent/workflowinstance"
)

// WorkflowEventsUpdate is the builder for updating WorkflowEvents entities.
type WorkflowEventsUpdate struct {
	config
	hooks    []Hook
	mutation *WorkflowEventsMutation
}

// Where adds a new predicate for the WorkflowEventsUpdate builder.
func (weu *WorkflowEventsUpdate) Where(ps ...predicate.WorkflowEvents) *WorkflowEventsUpdate {
	weu.mutation.predicates = append(weu.mutation.predicates, ps...)
	return weu
}

// SetEvents sets the "events" field.
func (weu *WorkflowEventsUpdate) SetEvents(m []map[string]interface{}) *WorkflowEventsUpdate {
	weu.mutation.SetEvents(m)
	return weu
}

// SetCorrelations sets the "correlations" field.
func (weu *WorkflowEventsUpdate) SetCorrelations(s []string) *WorkflowEventsUpdate {
	weu.mutation.SetCorrelations(s)
	return weu
}

// SetSignature sets the "signature" field.
func (weu *WorkflowEventsUpdate) SetSignature(b []byte) *WorkflowEventsUpdate {
	weu.mutation.SetSignature(b)
	return weu
}

// ClearSignature clears the value of the "signature" field.
func (weu *WorkflowEventsUpdate) ClearSignature() *WorkflowEventsUpdate {
	weu.mutation.ClearSignature()
	return weu
}

// SetCount sets the "count" field.
func (weu *WorkflowEventsUpdate) SetCount(i int) *WorkflowEventsUpdate {
	weu.mutation.ResetCount()
	weu.mutation.SetCount(i)
	return weu
}

// AddCount adds i to the "count" field.
func (weu *WorkflowEventsUpdate) AddCount(i int) *WorkflowEventsUpdate {
	weu.mutation.AddCount(i)
	return weu
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (weu *WorkflowEventsUpdate) SetWorkflowID(id uuid.UUID) *WorkflowEventsUpdate {
	weu.mutation.SetWorkflowID(id)
	return weu
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (weu *WorkflowEventsUpdate) SetWorkflow(w *Workflow) *WorkflowEventsUpdate {
	return weu.SetWorkflowID(w.ID)
}

// AddWfeventswaitIDs adds the "wfeventswait" edge to the WorkflowEventsWait entity by IDs.
func (weu *WorkflowEventsUpdate) AddWfeventswaitIDs(ids ...int) *WorkflowEventsUpdate {
	weu.mutation.AddWfeventswaitIDs(ids...)
	return weu
}

// AddWfeventswait adds the "wfeventswait" edges to the WorkflowEventsWait entity.
func (weu *WorkflowEventsUpdate) AddWfeventswait(w ...*WorkflowEventsWait) *WorkflowEventsUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return weu.AddWfeventswaitIDs(ids...)
}

// SetWorkflowinstanceID sets the "workflowinstance" edge to the WorkflowInstance entity by ID.
func (weu *WorkflowEventsUpdate) SetWorkflowinstanceID(id int) *WorkflowEventsUpdate {
	weu.mutation.SetWorkflowinstanceID(id)
	return weu
}

// SetNillableWorkflowinstanceID sets the "workflowinstance" edge to the WorkflowInstance entity by ID if the given value is not nil.
func (weu *WorkflowEventsUpdate) SetNillableWorkflowinstanceID(id *int) *WorkflowEventsUpdate {
	if id != nil {
		weu = weu.SetWorkflowinstanceID(*id)
	}
	return weu
}

// SetWorkflowinstance sets the "workflowinstance" edge to the WorkflowInstance entity.
func (weu *WorkflowEventsUpdate) SetWorkflowinstance(w *WorkflowInstance) *WorkflowEventsUpdate {
	return weu.SetWorkflowinstanceID(w.ID)
}

// Mutation returns the WorkflowEventsMutation object of the builder.
func (weu *WorkflowEventsUpdate) Mutation() *WorkflowEventsMutation {
	return weu.mutation
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (weu *WorkflowEventsUpdate) ClearWorkflow() *WorkflowEventsUpdate {
	weu.mutation.ClearWorkflow()
	return weu
}

// ClearWfeventswait clears all "wfeventswait" edges to the WorkflowEventsWait entity.
func (weu *WorkflowEventsUpdate) ClearWfeventswait() *WorkflowEventsUpdate {
	weu.mutation.ClearWfeventswait()
	return weu
}

// RemoveWfeventswaitIDs removes the "wfeventswait" edge to WorkflowEventsWait entities by IDs.
func (weu *WorkflowEventsUpdate) RemoveWfeventswaitIDs(ids ...int) *WorkflowEventsUpdate {
	weu.mutation.RemoveWfeventswaitIDs(ids...)
	return weu
}

// RemoveWfeventswait removes "wfeventswait" edges to WorkflowEventsWait entities.
func (weu *WorkflowEventsUpdate) RemoveWfeventswait(w ...*WorkflowEventsWait) *WorkflowEventsUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return weu.RemoveWfeventswaitIDs(ids...)
}

// ClearWorkflowinstance clears the "workflowinstance" edge to the WorkflowInstance entity.
func (weu *WorkflowEventsUpdate) ClearWorkflowinstance() *WorkflowEventsUpdate {
	weu.mutation.ClearWorkflowinstance()
	return weu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (weu *WorkflowEventsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(weu.hooks) == 0 {
		if err = weu.check(); err != nil {
			return 0, err
		}
		affected, err = weu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowEventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = weu.check(); err != nil {
				return 0, err
			}
			weu.mutation = mutation
			affected, err = weu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(weu.hooks) - 1; i >= 0; i-- {
			mut = weu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, weu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (weu *WorkflowEventsUpdate) SaveX(ctx context.Context) int {
	affected, err := weu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (weu *WorkflowEventsUpdate) Exec(ctx context.Context) error {
	_, err := weu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (weu *WorkflowEventsUpdate) ExecX(ctx context.Context) {
	if err := weu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (weu *WorkflowEventsUpdate) check() error {
	if _, ok := weu.mutation.WorkflowID(); weu.mutation.WorkflowCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workflow\"")
	}
	return nil
}

func (weu *WorkflowEventsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workflowevents.Table,
			Columns: workflowevents.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workflowevents.FieldID,
			},
		},
	}
	if ps := weu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := weu.mutation.Events(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workflowevents.FieldEvents,
		})
	}
	if value, ok := weu.mutation.Correlations(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workflowevents.FieldCorrelations,
		})
	}
	if value, ok := weu.mutation.Signature(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: workflowevents.FieldSignature,
		})
	}
	if weu.mutation.SignatureCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: workflowevents.FieldSignature,
		})
	}
	if value, ok := weu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workflowevents.FieldCount,
		})
	}
	if value, ok := weu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workflowevents.FieldCount,
		})
	}
	if weu.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflowevents.WorkflowTable,
			Columns: []string{workflowevents.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weu.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflowevents.WorkflowTable,
			Columns: []string{workflowevents.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if weu.mutation.WfeventswaitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowevents.WfeventswaitTable,
			Columns: []string{workflowevents.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workfloweventswait.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weu.mutation.RemovedWfeventswaitIDs(); len(nodes) > 0 && !weu.mutation.WfeventswaitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowevents.WfeventswaitTable,
			Columns: []string{workflowevents.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workfloweventswait.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weu.mutation.WfeventswaitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowevents.WfeventswaitTable,
			Columns: []string{workflowevents.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workfloweventswait.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if weu.mutation.WorkflowinstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflowevents.WorkflowinstanceTable,
			Columns: []string{workflowevents.WorkflowinstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowinstance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weu.mutation.WorkflowinstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflowevents.WorkflowinstanceTable,
			Columns: []string{workflowevents.WorkflowinstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowinstance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, weu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflowevents.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WorkflowEventsUpdateOne is the builder for updating a single WorkflowEvents entity.
type WorkflowEventsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkflowEventsMutation
}

// SetEvents sets the "events" field.
func (weuo *WorkflowEventsUpdateOne) SetEvents(m []map[string]interface{}) *WorkflowEventsUpdateOne {
	weuo.mutation.SetEvents(m)
	return weuo
}

// SetCorrelations sets the "correlations" field.
func (weuo *WorkflowEventsUpdateOne) SetCorrelations(s []string) *WorkflowEventsUpdateOne {
	weuo.mutation.SetCorrelations(s)
	return weuo
}

// SetSignature sets the "signature" field.
func (weuo *WorkflowEventsUpdateOne) SetSignature(b []byte) *WorkflowEventsUpdateOne {
	weuo.mutation.SetSignature(b)
	return weuo
}

// ClearSignature clears the value of the "signature" field.
func (weuo *WorkflowEventsUpdateOne) ClearSignature() *WorkflowEventsUpdateOne {
	weuo.mutation.ClearSignature()
	return weuo
}

// SetCount sets the "count" field.
func (weuo *WorkflowEventsUpdateOne) SetCount(i int) *WorkflowEventsUpdateOne {
	weuo.mutation.ResetCount()
	weuo.mutation.SetCount(i)
	return weuo
}

// AddCount adds i to the "count" field.
func (weuo *WorkflowEventsUpdateOne) AddCount(i int) *WorkflowEventsUpdateOne {
	weuo.mutation.AddCount(i)
	return weuo
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (weuo *WorkflowEventsUpdateOne) SetWorkflowID(id uuid.UUID) *WorkflowEventsUpdateOne {
	weuo.mutation.SetWorkflowID(id)
	return weuo
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (weuo *WorkflowEventsUpdateOne) SetWorkflow(w *Workflow) *WorkflowEventsUpdateOne {
	return weuo.SetWorkflowID(w.ID)
}

// AddWfeventswaitIDs adds the "wfeventswait" edge to the WorkflowEventsWait entity by IDs.
func (weuo *WorkflowEventsUpdateOne) AddWfeventswaitIDs(ids ...int) *WorkflowEventsUpdateOne {
	weuo.mutation.AddWfeventswaitIDs(ids...)
	return weuo
}

// AddWfeventswait adds the "wfeventswait" edges to the WorkflowEventsWait entity.
func (weuo *WorkflowEventsUpdateOne) AddWfeventswait(w ...*WorkflowEventsWait) *WorkflowEventsUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return weuo.AddWfeventswaitIDs(ids...)
}

// SetWorkflowinstanceID sets the "workflowinstance" edge to the WorkflowInstance entity by ID.
func (weuo *WorkflowEventsUpdateOne) SetWorkflowinstanceID(id int) *WorkflowEventsUpdateOne {
	weuo.mutation.SetWorkflowinstanceID(id)
	return weuo
}

// SetNillableWorkflowinstanceID sets the "workflowinstance" edge to the WorkflowInstance entity by ID if the given value is not nil.
func (weuo *WorkflowEventsUpdateOne) SetNillableWorkflowinstanceID(id *int) *WorkflowEventsUpdateOne {
	if id != nil {
		weuo = weuo.SetWorkflowinstanceID(*id)
	}
	return weuo
}

// SetWorkflowinstance sets the "workflowinstance" edge to the WorkflowInstance entity.
func (weuo *WorkflowEventsUpdateOne) SetWorkflowinstance(w *WorkflowInstance) *WorkflowEventsUpdateOne {
	return weuo.SetWorkflowinstanceID(w.ID)
}

// Mutation returns the WorkflowEventsMutation object of the builder.
func (weuo *WorkflowEventsUpdateOne) Mutation() *WorkflowEventsMutation {
	return weuo.mutation
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (weuo *WorkflowEventsUpdateOne) ClearWorkflow() *WorkflowEventsUpdateOne {
	weuo.mutation.ClearWorkflow()
	return weuo
}

// ClearWfeventswait clears all "wfeventswait" edges to the WorkflowEventsWait entity.
func (weuo *WorkflowEventsUpdateOne) ClearWfeventswait() *WorkflowEventsUpdateOne {
	weuo.mutation.ClearWfeventswait()
	return weuo
}

// RemoveWfeventswaitIDs removes the "wfeventswait" edge to WorkflowEventsWait entities by IDs.
func (weuo *WorkflowEventsUpdateOne) RemoveWfeventswaitIDs(ids ...int) *WorkflowEventsUpdateOne {
	weuo.mutation.RemoveWfeventswaitIDs(ids...)
	return weuo
}

// RemoveWfeventswait removes "wfeventswait" edges to WorkflowEventsWait entities.
func (weuo *WorkflowEventsUpdateOne) RemoveWfeventswait(w ...*WorkflowEventsWait) *WorkflowEventsUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return weuo.RemoveWfeventswaitIDs(ids...)
}

// ClearWorkflowinstance clears the "workflowinstance" edge to the WorkflowInstance entity.
func (weuo *WorkflowEventsUpdateOne) ClearWorkflowinstance() *WorkflowEventsUpdateOne {
	weuo.mutation.ClearWorkflowinstance()
	return weuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (weuo *WorkflowEventsUpdateOne) Select(field string, fields ...string) *WorkflowEventsUpdateOne {
	weuo.fields = append([]string{field}, fields...)
	return weuo
}

// Save executes the query and returns the updated WorkflowEvents entity.
func (weuo *WorkflowEventsUpdateOne) Save(ctx context.Context) (*WorkflowEvents, error) {
	var (
		err  error
		node *WorkflowEvents
	)
	if len(weuo.hooks) == 0 {
		if err = weuo.check(); err != nil {
			return nil, err
		}
		node, err = weuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowEventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = weuo.check(); err != nil {
				return nil, err
			}
			weuo.mutation = mutation
			node, err = weuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(weuo.hooks) - 1; i >= 0; i-- {
			mut = weuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, weuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (weuo *WorkflowEventsUpdateOne) SaveX(ctx context.Context) *WorkflowEvents {
	node, err := weuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (weuo *WorkflowEventsUpdateOne) Exec(ctx context.Context) error {
	_, err := weuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (weuo *WorkflowEventsUpdateOne) ExecX(ctx context.Context) {
	if err := weuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (weuo *WorkflowEventsUpdateOne) check() error {
	if _, ok := weuo.mutation.WorkflowID(); weuo.mutation.WorkflowCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workflow\"")
	}
	return nil
}

func (weuo *WorkflowEventsUpdateOne) sqlSave(ctx context.Context) (_node *WorkflowEvents, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workflowevents.Table,
			Columns: workflowevents.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workflowevents.FieldID,
			},
		},
	}
	id, ok := weuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WorkflowEvents.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := weuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflowevents.FieldID)
		for _, f := range fields {
			if !workflowevents.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workflowevents.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := weuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := weuo.mutation.Events(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workflowevents.FieldEvents,
		})
	}
	if value, ok := weuo.mutation.Correlations(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workflowevents.FieldCorrelations,
		})
	}
	if value, ok := weuo.mutation.Signature(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: workflowevents.FieldSignature,
		})
	}
	if weuo.mutation.SignatureCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: workflowevents.FieldSignature,
		})
	}
	if value, ok := weuo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workflowevents.FieldCount,
		})
	}
	if value, ok := weuo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workflowevents.FieldCount,
		})
	}
	if weuo.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflowevents.WorkflowTable,
			Columns: []string{workflowevents.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weuo.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflowevents.WorkflowTable,
			Columns: []string{workflowevents.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if weuo.mutation.WfeventswaitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowevents.WfeventswaitTable,
			Columns: []string{workflowevents.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workfloweventswait.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weuo.mutation.RemovedWfeventswaitIDs(); len(nodes) > 0 && !weuo.mutation.WfeventswaitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowevents.WfeventswaitTable,
			Columns: []string{workflowevents.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workfloweventswait.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weuo.mutation.WfeventswaitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowevents.WfeventswaitTable,
			Columns: []string{workflowevents.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workfloweventswait.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if weuo.mutation.WorkflowinstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflowevents.WorkflowinstanceTable,
			Columns: []string{workflowevents.WorkflowinstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowinstance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weuo.mutation.WorkflowinstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflowevents.WorkflowinstanceTable,
			Columns: []string{workflowevents.WorkflowinstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowinstance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkflowEvents{config: weuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, weuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflowevents.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
