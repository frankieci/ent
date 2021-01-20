// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/ent/group"
	"github.com/facebook/ent/entc/integration/ent/groupinfo"
	"github.com/facebook/ent/entc/integration/ent/internal"
	"github.com/facebook/ent/entc/integration/ent/predicate"
	"github.com/facebook/ent/schema/field"
)

// GroupInfoUpdate is the builder for updating GroupInfo entities.
type GroupInfoUpdate struct {
	config
	hooks    []Hook
	mutation *GroupInfoMutation
}

// Where adds a new predicate for the GroupInfoUpdate builder.
func (giu *GroupInfoUpdate) Where(ps ...predicate.GroupInfo) *GroupInfoUpdate {
	giu.mutation.predicates = append(giu.mutation.predicates, ps...)
	return giu
}

// SetDesc sets the "desc" field.
func (giu *GroupInfoUpdate) SetDesc(s string) *GroupInfoUpdate {
	giu.mutation.SetDesc(s)
	return giu
}

// SetMaxUsers sets the "max_users" field.
func (giu *GroupInfoUpdate) SetMaxUsers(i int) *GroupInfoUpdate {
	giu.mutation.ResetMaxUsers()
	giu.mutation.SetMaxUsers(i)
	return giu
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (giu *GroupInfoUpdate) SetNillableMaxUsers(i *int) *GroupInfoUpdate {
	if i != nil {
		giu.SetMaxUsers(*i)
	}
	return giu
}

// AddMaxUsers adds i to the "max_users" field.
func (giu *GroupInfoUpdate) AddMaxUsers(i int) *GroupInfoUpdate {
	giu.mutation.AddMaxUsers(i)
	return giu
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (giu *GroupInfoUpdate) AddGroupIDs(ids ...int) *GroupInfoUpdate {
	giu.mutation.AddGroupIDs(ids...)
	return giu
}

// AddGroups adds the "groups" edges to the Group entity.
func (giu *GroupInfoUpdate) AddGroups(g ...*Group) *GroupInfoUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giu.AddGroupIDs(ids...)
}

// Mutation returns the GroupInfoMutation object of the builder.
func (giu *GroupInfoUpdate) Mutation() *GroupInfoMutation {
	return giu.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (giu *GroupInfoUpdate) ClearGroups() *GroupInfoUpdate {
	giu.mutation.ClearGroups()
	return giu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (giu *GroupInfoUpdate) RemoveGroupIDs(ids ...int) *GroupInfoUpdate {
	giu.mutation.RemoveGroupIDs(ids...)
	return giu
}

// RemoveGroups removes "groups" edges to Group entities.
func (giu *GroupInfoUpdate) RemoveGroups(g ...*Group) *GroupInfoUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (giu *GroupInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(giu.hooks) == 0 {
		affected, err = giu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			giu.mutation = mutation
			affected, err = giu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(giu.hooks) - 1; i >= 0; i-- {
			mut = giu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, giu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (giu *GroupInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := giu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (giu *GroupInfoUpdate) Exec(ctx context.Context) error {
	_, err := giu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (giu *GroupInfoUpdate) ExecX(ctx context.Context) {
	if err := giu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (giu *GroupInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupinfo.Table,
			Columns: groupinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupinfo.FieldID,
			},
		},
	}
	if ps := giu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := giu.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupinfo.FieldDesc,
		})
	}
	if value, ok := giu.mutation.MaxUsers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupinfo.FieldMaxUsers,
		})
	}
	if value, ok := giu.mutation.AddedMaxUsers(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupinfo.FieldMaxUsers,
		})
	}
	if giu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		edge.Schema = giu.schemaConfig.GroupInfo
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := giu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !giu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		edge.Schema = giu.schemaConfig.GroupInfo
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := giu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		edge.Schema = giu.schemaConfig.GroupInfo
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = giu.schemaConfig.GroupInfo
	ctx = internal.NewSchemaConfigContext(ctx, giu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, giu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupinfo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GroupInfoUpdateOne is the builder for updating a single GroupInfo entity.
type GroupInfoUpdateOne struct {
	config
	hooks    []Hook
	mutation *GroupInfoMutation
}

// SetDesc sets the "desc" field.
func (giuo *GroupInfoUpdateOne) SetDesc(s string) *GroupInfoUpdateOne {
	giuo.mutation.SetDesc(s)
	return giuo
}

// SetMaxUsers sets the "max_users" field.
func (giuo *GroupInfoUpdateOne) SetMaxUsers(i int) *GroupInfoUpdateOne {
	giuo.mutation.ResetMaxUsers()
	giuo.mutation.SetMaxUsers(i)
	return giuo
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (giuo *GroupInfoUpdateOne) SetNillableMaxUsers(i *int) *GroupInfoUpdateOne {
	if i != nil {
		giuo.SetMaxUsers(*i)
	}
	return giuo
}

// AddMaxUsers adds i to the "max_users" field.
func (giuo *GroupInfoUpdateOne) AddMaxUsers(i int) *GroupInfoUpdateOne {
	giuo.mutation.AddMaxUsers(i)
	return giuo
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (giuo *GroupInfoUpdateOne) AddGroupIDs(ids ...int) *GroupInfoUpdateOne {
	giuo.mutation.AddGroupIDs(ids...)
	return giuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (giuo *GroupInfoUpdateOne) AddGroups(g ...*Group) *GroupInfoUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giuo.AddGroupIDs(ids...)
}

// Mutation returns the GroupInfoMutation object of the builder.
func (giuo *GroupInfoUpdateOne) Mutation() *GroupInfoMutation {
	return giuo.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (giuo *GroupInfoUpdateOne) ClearGroups() *GroupInfoUpdateOne {
	giuo.mutation.ClearGroups()
	return giuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (giuo *GroupInfoUpdateOne) RemoveGroupIDs(ids ...int) *GroupInfoUpdateOne {
	giuo.mutation.RemoveGroupIDs(ids...)
	return giuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (giuo *GroupInfoUpdateOne) RemoveGroups(g ...*Group) *GroupInfoUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giuo.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the updated GroupInfo entity.
func (giuo *GroupInfoUpdateOne) Save(ctx context.Context) (*GroupInfo, error) {
	var (
		err  error
		node *GroupInfo
	)
	if len(giuo.hooks) == 0 {
		node, err = giuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			giuo.mutation = mutation
			node, err = giuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(giuo.hooks) - 1; i >= 0; i-- {
			mut = giuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, giuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (giuo *GroupInfoUpdateOne) SaveX(ctx context.Context) *GroupInfo {
	node, err := giuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (giuo *GroupInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := giuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (giuo *GroupInfoUpdateOne) ExecX(ctx context.Context) {
	if err := giuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (giuo *GroupInfoUpdateOne) sqlSave(ctx context.Context) (_node *GroupInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupinfo.Table,
			Columns: groupinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupinfo.FieldID,
			},
		},
	}
	id, ok := giuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GroupInfo.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := giuo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupinfo.FieldDesc,
		})
	}
	if value, ok := giuo.mutation.MaxUsers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupinfo.FieldMaxUsers,
		})
	}
	if value, ok := giuo.mutation.AddedMaxUsers(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupinfo.FieldMaxUsers,
		})
	}
	if giuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		edge.Schema = giuo.schemaConfig.GroupInfo
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := giuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !giuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		edge.Schema = giuo.schemaConfig.GroupInfo
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := giuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		edge.Schema = giuo.schemaConfig.GroupInfo
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = giuo.schemaConfig.GroupInfo
	ctx = internal.NewSchemaConfigContext(ctx, giuo.schemaConfig)
	_node = &GroupInfo{config: giuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, giuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupinfo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
