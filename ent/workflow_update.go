// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vorteil/direktiv/ent/namespace"
	"github.com/vorteil/direktiv/ent/predicate"
	"github.com/vorteil/direktiv/ent/workflow"
	"github.com/vorteil/direktiv/ent/workflowevents"
	"github.com/vorteil/direktiv/ent/workflowinstance"
)

// WorkflowUpdate is the builder for updating Workflow entities.
type WorkflowUpdate struct {
	config
	hooks    []Hook
	mutation *WorkflowMutation
}

// Where adds a new predicate for the WorkflowUpdate builder.
func (wu *WorkflowUpdate) Where(ps ...predicate.Workflow) *WorkflowUpdate {
	wu.mutation.predicates = append(wu.mutation.predicates, ps...)
	return wu
}

// SetName sets the "name" field.
func (wu *WorkflowUpdate) SetName(s string) *WorkflowUpdate {
	wu.mutation.SetName(s)
	return wu
}

// SetDescription sets the "description" field.
func (wu *WorkflowUpdate) SetDescription(s string) *WorkflowUpdate {
	wu.mutation.SetDescription(s)
	return wu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableDescription(s *string) *WorkflowUpdate {
	if s != nil {
		wu.SetDescription(*s)
	}
	return wu
}

// ClearDescription clears the value of the "description" field.
func (wu *WorkflowUpdate) ClearDescription() *WorkflowUpdate {
	wu.mutation.ClearDescription()
	return wu
}

// SetActive sets the "active" field.
func (wu *WorkflowUpdate) SetActive(b bool) *WorkflowUpdate {
	wu.mutation.SetActive(b)
	return wu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableActive(b *bool) *WorkflowUpdate {
	if b != nil {
		wu.SetActive(*b)
	}
	return wu
}

// SetRevision sets the "revision" field.
func (wu *WorkflowUpdate) SetRevision(i int) *WorkflowUpdate {
	wu.mutation.ResetRevision()
	wu.mutation.SetRevision(i)
	return wu
}

// SetNillableRevision sets the "revision" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableRevision(i *int) *WorkflowUpdate {
	if i != nil {
		wu.SetRevision(*i)
	}
	return wu
}

// AddRevision adds i to the "revision" field.
func (wu *WorkflowUpdate) AddRevision(i int) *WorkflowUpdate {
	wu.mutation.AddRevision(i)
	return wu
}

// SetWorkflow sets the "workflow" field.
func (wu *WorkflowUpdate) SetWorkflow(b []byte) *WorkflowUpdate {
	wu.mutation.SetWorkflow(b)
	return wu
}

// SetLogToEvents sets the "logToEvents" field.
func (wu *WorkflowUpdate) SetLogToEvents(s string) *WorkflowUpdate {
	wu.mutation.SetLogToEvents(s)
	return wu
}

// SetNillableLogToEvents sets the "logToEvents" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableLogToEvents(s *string) *WorkflowUpdate {
	if s != nil {
		wu.SetLogToEvents(*s)
	}
	return wu
}

// ClearLogToEvents clears the value of the "logToEvents" field.
func (wu *WorkflowUpdate) ClearLogToEvents() *WorkflowUpdate {
	wu.mutation.ClearLogToEvents()
	return wu
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (wu *WorkflowUpdate) SetNamespaceID(id string) *WorkflowUpdate {
	wu.mutation.SetNamespaceID(id)
	return wu
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (wu *WorkflowUpdate) SetNamespace(n *Namespace) *WorkflowUpdate {
	return wu.SetNamespaceID(n.ID)
}

// AddInstanceIDs adds the "instances" edge to the WorkflowInstance entity by IDs.
func (wu *WorkflowUpdate) AddInstanceIDs(ids ...int) *WorkflowUpdate {
	wu.mutation.AddInstanceIDs(ids...)
	return wu
}

// AddInstances adds the "instances" edges to the WorkflowInstance entity.
func (wu *WorkflowUpdate) AddInstances(w ...*WorkflowInstance) *WorkflowUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddInstanceIDs(ids...)
}

// AddWfeventIDs adds the "wfevents" edge to the WorkflowEvents entity by IDs.
func (wu *WorkflowUpdate) AddWfeventIDs(ids ...int) *WorkflowUpdate {
	wu.mutation.AddWfeventIDs(ids...)
	return wu
}

// AddWfevents adds the "wfevents" edges to the WorkflowEvents entity.
func (wu *WorkflowUpdate) AddWfevents(w ...*WorkflowEvents) *WorkflowUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddWfeventIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wu *WorkflowUpdate) Mutation() *WorkflowMutation {
	return wu.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (wu *WorkflowUpdate) ClearNamespace() *WorkflowUpdate {
	wu.mutation.ClearNamespace()
	return wu
}

// ClearInstances clears all "instances" edges to the WorkflowInstance entity.
func (wu *WorkflowUpdate) ClearInstances() *WorkflowUpdate {
	wu.mutation.ClearInstances()
	return wu
}

// RemoveInstanceIDs removes the "instances" edge to WorkflowInstance entities by IDs.
func (wu *WorkflowUpdate) RemoveInstanceIDs(ids ...int) *WorkflowUpdate {
	wu.mutation.RemoveInstanceIDs(ids...)
	return wu
}

// RemoveInstances removes "instances" edges to WorkflowInstance entities.
func (wu *WorkflowUpdate) RemoveInstances(w ...*WorkflowInstance) *WorkflowUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveInstanceIDs(ids...)
}

// ClearWfevents clears all "wfevents" edges to the WorkflowEvents entity.
func (wu *WorkflowUpdate) ClearWfevents() *WorkflowUpdate {
	wu.mutation.ClearWfevents()
	return wu
}

// RemoveWfeventIDs removes the "wfevents" edge to WorkflowEvents entities by IDs.
func (wu *WorkflowUpdate) RemoveWfeventIDs(ids ...int) *WorkflowUpdate {
	wu.mutation.RemoveWfeventIDs(ids...)
	return wu
}

// RemoveWfevents removes "wfevents" edges to WorkflowEvents entities.
func (wu *WorkflowUpdate) RemoveWfevents(w ...*WorkflowEvents) *WorkflowUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveWfeventIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkflowUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkflowUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkflowUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkflowUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkflowUpdate) check() error {
	if v, ok := wu.mutation.Name(); ok {
		if err := workflow.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := wu.mutation.Description(); ok {
		if err := workflow.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	if _, ok := wu.mutation.NamespaceID(); wu.mutation.NamespaceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"namespace\"")
	}
	return nil
}

func (wu *WorkflowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workflow.Table,
			Columns: workflow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workflow.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflow.FieldName,
		})
	}
	if value, ok := wu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflow.FieldDescription,
		})
	}
	if wu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workflow.FieldDescription,
		})
	}
	if value, ok := wu.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: workflow.FieldActive,
		})
	}
	if value, ok := wu.mutation.Revision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workflow.FieldRevision,
		})
	}
	if value, ok := wu.mutation.AddedRevision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workflow.FieldRevision,
		})
	}
	if value, ok := wu.mutation.Workflow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: workflow.FieldWorkflow,
		})
	}
	if value, ok := wu.mutation.LogToEvents(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflow.FieldLogToEvents,
		})
	}
	if wu.mutation.LogToEventsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workflow.FieldLogToEvents,
		})
	}
	if wu.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.NamespaceTable,
			Columns: []string{workflow.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: namespace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.NamespaceTable,
			Columns: []string{workflow.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.InstancesTable,
			Columns: []string{workflow.InstancesColumn},
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
	if nodes := wu.mutation.RemovedInstancesIDs(); len(nodes) > 0 && !wu.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.InstancesTable,
			Columns: []string{workflow.InstancesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.InstancesTable,
			Columns: []string{workflow.InstancesColumn},
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
	if wu.mutation.WfeventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WfeventsTable,
			Columns: []string{workflow.WfeventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedWfeventsIDs(); len(nodes) > 0 && !wu.mutation.WfeventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WfeventsTable,
			Columns: []string{workflow.WfeventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.WfeventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WfeventsTable,
			Columns: []string{workflow.WfeventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflow.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WorkflowUpdateOne is the builder for updating a single Workflow entity.
type WorkflowUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkflowMutation
}

// SetName sets the "name" field.
func (wuo *WorkflowUpdateOne) SetName(s string) *WorkflowUpdateOne {
	wuo.mutation.SetName(s)
	return wuo
}

// SetDescription sets the "description" field.
func (wuo *WorkflowUpdateOne) SetDescription(s string) *WorkflowUpdateOne {
	wuo.mutation.SetDescription(s)
	return wuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableDescription(s *string) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetDescription(*s)
	}
	return wuo
}

// ClearDescription clears the value of the "description" field.
func (wuo *WorkflowUpdateOne) ClearDescription() *WorkflowUpdateOne {
	wuo.mutation.ClearDescription()
	return wuo
}

// SetActive sets the "active" field.
func (wuo *WorkflowUpdateOne) SetActive(b bool) *WorkflowUpdateOne {
	wuo.mutation.SetActive(b)
	return wuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableActive(b *bool) *WorkflowUpdateOne {
	if b != nil {
		wuo.SetActive(*b)
	}
	return wuo
}

// SetRevision sets the "revision" field.
func (wuo *WorkflowUpdateOne) SetRevision(i int) *WorkflowUpdateOne {
	wuo.mutation.ResetRevision()
	wuo.mutation.SetRevision(i)
	return wuo
}

// SetNillableRevision sets the "revision" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableRevision(i *int) *WorkflowUpdateOne {
	if i != nil {
		wuo.SetRevision(*i)
	}
	return wuo
}

// AddRevision adds i to the "revision" field.
func (wuo *WorkflowUpdateOne) AddRevision(i int) *WorkflowUpdateOne {
	wuo.mutation.AddRevision(i)
	return wuo
}

// SetWorkflow sets the "workflow" field.
func (wuo *WorkflowUpdateOne) SetWorkflow(b []byte) *WorkflowUpdateOne {
	wuo.mutation.SetWorkflow(b)
	return wuo
}

// SetLogToEvents sets the "logToEvents" field.
func (wuo *WorkflowUpdateOne) SetLogToEvents(s string) *WorkflowUpdateOne {
	wuo.mutation.SetLogToEvents(s)
	return wuo
}

// SetNillableLogToEvents sets the "logToEvents" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableLogToEvents(s *string) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetLogToEvents(*s)
	}
	return wuo
}

// ClearLogToEvents clears the value of the "logToEvents" field.
func (wuo *WorkflowUpdateOne) ClearLogToEvents() *WorkflowUpdateOne {
	wuo.mutation.ClearLogToEvents()
	return wuo
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (wuo *WorkflowUpdateOne) SetNamespaceID(id string) *WorkflowUpdateOne {
	wuo.mutation.SetNamespaceID(id)
	return wuo
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (wuo *WorkflowUpdateOne) SetNamespace(n *Namespace) *WorkflowUpdateOne {
	return wuo.SetNamespaceID(n.ID)
}

// AddInstanceIDs adds the "instances" edge to the WorkflowInstance entity by IDs.
func (wuo *WorkflowUpdateOne) AddInstanceIDs(ids ...int) *WorkflowUpdateOne {
	wuo.mutation.AddInstanceIDs(ids...)
	return wuo
}

// AddInstances adds the "instances" edges to the WorkflowInstance entity.
func (wuo *WorkflowUpdateOne) AddInstances(w ...*WorkflowInstance) *WorkflowUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddInstanceIDs(ids...)
}

// AddWfeventIDs adds the "wfevents" edge to the WorkflowEvents entity by IDs.
func (wuo *WorkflowUpdateOne) AddWfeventIDs(ids ...int) *WorkflowUpdateOne {
	wuo.mutation.AddWfeventIDs(ids...)
	return wuo
}

// AddWfevents adds the "wfevents" edges to the WorkflowEvents entity.
func (wuo *WorkflowUpdateOne) AddWfevents(w ...*WorkflowEvents) *WorkflowUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddWfeventIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wuo *WorkflowUpdateOne) Mutation() *WorkflowMutation {
	return wuo.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (wuo *WorkflowUpdateOne) ClearNamespace() *WorkflowUpdateOne {
	wuo.mutation.ClearNamespace()
	return wuo
}

// ClearInstances clears all "instances" edges to the WorkflowInstance entity.
func (wuo *WorkflowUpdateOne) ClearInstances() *WorkflowUpdateOne {
	wuo.mutation.ClearInstances()
	return wuo
}

// RemoveInstanceIDs removes the "instances" edge to WorkflowInstance entities by IDs.
func (wuo *WorkflowUpdateOne) RemoveInstanceIDs(ids ...int) *WorkflowUpdateOne {
	wuo.mutation.RemoveInstanceIDs(ids...)
	return wuo
}

// RemoveInstances removes "instances" edges to WorkflowInstance entities.
func (wuo *WorkflowUpdateOne) RemoveInstances(w ...*WorkflowInstance) *WorkflowUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveInstanceIDs(ids...)
}

// ClearWfevents clears all "wfevents" edges to the WorkflowEvents entity.
func (wuo *WorkflowUpdateOne) ClearWfevents() *WorkflowUpdateOne {
	wuo.mutation.ClearWfevents()
	return wuo
}

// RemoveWfeventIDs removes the "wfevents" edge to WorkflowEvents entities by IDs.
func (wuo *WorkflowUpdateOne) RemoveWfeventIDs(ids ...int) *WorkflowUpdateOne {
	wuo.mutation.RemoveWfeventIDs(ids...)
	return wuo
}

// RemoveWfevents removes "wfevents" edges to WorkflowEvents entities.
func (wuo *WorkflowUpdateOne) RemoveWfevents(w ...*WorkflowEvents) *WorkflowUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveWfeventIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkflowUpdateOne) Select(field string, fields ...string) *WorkflowUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Workflow entity.
func (wuo *WorkflowUpdateOne) Save(ctx context.Context) (*Workflow, error) {
	var (
		err  error
		node *Workflow
	)
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) SaveX(ctx context.Context) *Workflow {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkflowUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkflowUpdateOne) check() error {
	if v, ok := wuo.mutation.Name(); ok {
		if err := workflow.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := wuo.mutation.Description(); ok {
		if err := workflow.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	if _, ok := wuo.mutation.NamespaceID(); wuo.mutation.NamespaceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"namespace\"")
	}
	return nil
}

func (wuo *WorkflowUpdateOne) sqlSave(ctx context.Context) (_node *Workflow, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workflow.Table,
			Columns: workflow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workflow.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Workflow.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflow.FieldID)
		for _, f := range fields {
			if !workflow.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workflow.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflow.FieldName,
		})
	}
	if value, ok := wuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflow.FieldDescription,
		})
	}
	if wuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workflow.FieldDescription,
		})
	}
	if value, ok := wuo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: workflow.FieldActive,
		})
	}
	if value, ok := wuo.mutation.Revision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workflow.FieldRevision,
		})
	}
	if value, ok := wuo.mutation.AddedRevision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workflow.FieldRevision,
		})
	}
	if value, ok := wuo.mutation.Workflow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: workflow.FieldWorkflow,
		})
	}
	if value, ok := wuo.mutation.LogToEvents(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflow.FieldLogToEvents,
		})
	}
	if wuo.mutation.LogToEventsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workflow.FieldLogToEvents,
		})
	}
	if wuo.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.NamespaceTable,
			Columns: []string{workflow.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: namespace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.NamespaceTable,
			Columns: []string{workflow.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.InstancesTable,
			Columns: []string{workflow.InstancesColumn},
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
	if nodes := wuo.mutation.RemovedInstancesIDs(); len(nodes) > 0 && !wuo.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.InstancesTable,
			Columns: []string{workflow.InstancesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.InstancesTable,
			Columns: []string{workflow.InstancesColumn},
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
	if wuo.mutation.WfeventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WfeventsTable,
			Columns: []string{workflow.WfeventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedWfeventsIDs(); len(nodes) > 0 && !wuo.mutation.WfeventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WfeventsTable,
			Columns: []string{workflow.WfeventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.WfeventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WfeventsTable,
			Columns: []string{workflow.WfeventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Workflow{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflow.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
