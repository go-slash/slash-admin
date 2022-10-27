// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysdictionary"
	"slash-admin/pkg/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysDictionaryUpdate is the builder for updating SysDictionary entities.
type SysDictionaryUpdate struct {
	config
	hooks     []Hook
	mutation  *SysDictionaryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysDictionaryUpdate builder.
func (sdu *SysDictionaryUpdate) Where(ps ...predicate.SysDictionary) *SysDictionaryUpdate {
	sdu.mutation.Where(ps...)
	return sdu
}

// SetTitle sets the "title" field.
func (sdu *SysDictionaryUpdate) SetTitle(s string) *SysDictionaryUpdate {
	sdu.mutation.SetTitle(s)
	return sdu
}

// SetName sets the "name" field.
func (sdu *SysDictionaryUpdate) SetName(s string) *SysDictionaryUpdate {
	sdu.mutation.SetName(s)
	return sdu
}

// SetStatus sets the "status" field.
func (sdu *SysDictionaryUpdate) SetStatus(t types.Status) *SysDictionaryUpdate {
	sdu.mutation.ResetStatus()
	sdu.mutation.SetStatus(t)
	return sdu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sdu *SysDictionaryUpdate) SetNillableStatus(t *types.Status) *SysDictionaryUpdate {
	if t != nil {
		sdu.SetStatus(*t)
	}
	return sdu
}

// AddStatus adds t to the "status" field.
func (sdu *SysDictionaryUpdate) AddStatus(t types.Status) *SysDictionaryUpdate {
	sdu.mutation.AddStatus(t)
	return sdu
}

// ClearStatus clears the value of the "status" field.
func (sdu *SysDictionaryUpdate) ClearStatus() *SysDictionaryUpdate {
	sdu.mutation.ClearStatus()
	return sdu
}

// SetDesc sets the "desc" field.
func (sdu *SysDictionaryUpdate) SetDesc(s string) *SysDictionaryUpdate {
	sdu.mutation.SetDesc(s)
	return sdu
}

// SetCreatedAt sets the "created_at" field.
func (sdu *SysDictionaryUpdate) SetCreatedAt(t time.Time) *SysDictionaryUpdate {
	sdu.mutation.SetCreatedAt(t)
	return sdu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sdu *SysDictionaryUpdate) SetNillableCreatedAt(t *time.Time) *SysDictionaryUpdate {
	if t != nil {
		sdu.SetCreatedAt(*t)
	}
	return sdu
}

// SetUpdatedAt sets the "updated_at" field.
func (sdu *SysDictionaryUpdate) SetUpdatedAt(t time.Time) *SysDictionaryUpdate {
	sdu.mutation.SetUpdatedAt(t)
	return sdu
}

// SetDeletedAt sets the "deleted_at" field.
func (sdu *SysDictionaryUpdate) SetDeletedAt(t time.Time) *SysDictionaryUpdate {
	sdu.mutation.SetDeletedAt(t)
	return sdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdu *SysDictionaryUpdate) SetNillableDeletedAt(t *time.Time) *SysDictionaryUpdate {
	if t != nil {
		sdu.SetDeletedAt(*t)
	}
	return sdu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sdu *SysDictionaryUpdate) ClearDeletedAt() *SysDictionaryUpdate {
	sdu.mutation.ClearDeletedAt()
	return sdu
}

// Mutation returns the SysDictionaryMutation object of the builder.
func (sdu *SysDictionaryUpdate) Mutation() *SysDictionaryMutation {
	return sdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sdu *SysDictionaryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sdu.defaults()
	if len(sdu.hooks) == 0 {
		affected, err = sdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictionaryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdu.mutation = mutation
			affected, err = sdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdu.hooks) - 1; i >= 0; i-- {
			if sdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sdu *SysDictionaryUpdate) SaveX(ctx context.Context) int {
	affected, err := sdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sdu *SysDictionaryUpdate) Exec(ctx context.Context) error {
	_, err := sdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdu *SysDictionaryUpdate) ExecX(ctx context.Context) {
	if err := sdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdu *SysDictionaryUpdate) defaults() {
	if _, ok := sdu.mutation.UpdatedAt(); !ok {
		v := sysdictionary.UpdateDefaultUpdatedAt()
		sdu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sdu *SysDictionaryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysDictionaryUpdate {
	sdu.modifiers = append(sdu.modifiers, modifiers...)
	return sdu
}

func (sdu *SysDictionaryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdictionary.Table,
			Columns: sysdictionary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysdictionary.FieldID,
			},
		},
	}
	if ps := sdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionary.FieldTitle,
		})
	}
	if value, ok := sdu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionary.FieldName,
		})
	}
	if value, ok := sdu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysdictionary.FieldStatus,
		})
	}
	if value, ok := sdu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysdictionary.FieldStatus,
		})
	}
	if sdu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: sysdictionary.FieldStatus,
		})
	}
	if value, ok := sdu.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionary.FieldDesc,
		})
	}
	if value, ok := sdu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionary.FieldCreatedAt,
		})
	}
	if value, ok := sdu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionary.FieldUpdatedAt,
		})
	}
	if value, ok := sdu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionary.FieldDeletedAt,
		})
	}
	if sdu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysdictionary.FieldDeletedAt,
		})
	}
	_spec.AddModifiers(sdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdictionary.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SysDictionaryUpdateOne is the builder for updating a single SysDictionary entity.
type SysDictionaryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysDictionaryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTitle sets the "title" field.
func (sduo *SysDictionaryUpdateOne) SetTitle(s string) *SysDictionaryUpdateOne {
	sduo.mutation.SetTitle(s)
	return sduo
}

// SetName sets the "name" field.
func (sduo *SysDictionaryUpdateOne) SetName(s string) *SysDictionaryUpdateOne {
	sduo.mutation.SetName(s)
	return sduo
}

// SetStatus sets the "status" field.
func (sduo *SysDictionaryUpdateOne) SetStatus(t types.Status) *SysDictionaryUpdateOne {
	sduo.mutation.ResetStatus()
	sduo.mutation.SetStatus(t)
	return sduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sduo *SysDictionaryUpdateOne) SetNillableStatus(t *types.Status) *SysDictionaryUpdateOne {
	if t != nil {
		sduo.SetStatus(*t)
	}
	return sduo
}

// AddStatus adds t to the "status" field.
func (sduo *SysDictionaryUpdateOne) AddStatus(t types.Status) *SysDictionaryUpdateOne {
	sduo.mutation.AddStatus(t)
	return sduo
}

// ClearStatus clears the value of the "status" field.
func (sduo *SysDictionaryUpdateOne) ClearStatus() *SysDictionaryUpdateOne {
	sduo.mutation.ClearStatus()
	return sduo
}

// SetDesc sets the "desc" field.
func (sduo *SysDictionaryUpdateOne) SetDesc(s string) *SysDictionaryUpdateOne {
	sduo.mutation.SetDesc(s)
	return sduo
}

// SetCreatedAt sets the "created_at" field.
func (sduo *SysDictionaryUpdateOne) SetCreatedAt(t time.Time) *SysDictionaryUpdateOne {
	sduo.mutation.SetCreatedAt(t)
	return sduo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sduo *SysDictionaryUpdateOne) SetNillableCreatedAt(t *time.Time) *SysDictionaryUpdateOne {
	if t != nil {
		sduo.SetCreatedAt(*t)
	}
	return sduo
}

// SetUpdatedAt sets the "updated_at" field.
func (sduo *SysDictionaryUpdateOne) SetUpdatedAt(t time.Time) *SysDictionaryUpdateOne {
	sduo.mutation.SetUpdatedAt(t)
	return sduo
}

// SetDeletedAt sets the "deleted_at" field.
func (sduo *SysDictionaryUpdateOne) SetDeletedAt(t time.Time) *SysDictionaryUpdateOne {
	sduo.mutation.SetDeletedAt(t)
	return sduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sduo *SysDictionaryUpdateOne) SetNillableDeletedAt(t *time.Time) *SysDictionaryUpdateOne {
	if t != nil {
		sduo.SetDeletedAt(*t)
	}
	return sduo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sduo *SysDictionaryUpdateOne) ClearDeletedAt() *SysDictionaryUpdateOne {
	sduo.mutation.ClearDeletedAt()
	return sduo
}

// Mutation returns the SysDictionaryMutation object of the builder.
func (sduo *SysDictionaryUpdateOne) Mutation() *SysDictionaryMutation {
	return sduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sduo *SysDictionaryUpdateOne) Select(field string, fields ...string) *SysDictionaryUpdateOne {
	sduo.fields = append([]string{field}, fields...)
	return sduo
}

// Save executes the query and returns the updated SysDictionary entity.
func (sduo *SysDictionaryUpdateOne) Save(ctx context.Context) (*SysDictionary, error) {
	var (
		err  error
		node *SysDictionary
	)
	sduo.defaults()
	if len(sduo.hooks) == 0 {
		node, err = sduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictionaryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sduo.mutation = mutation
			node, err = sduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sduo.hooks) - 1; i >= 0; i-- {
			if sduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysDictionary)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysDictionaryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sduo *SysDictionaryUpdateOne) SaveX(ctx context.Context) *SysDictionary {
	node, err := sduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sduo *SysDictionaryUpdateOne) Exec(ctx context.Context) error {
	_, err := sduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sduo *SysDictionaryUpdateOne) ExecX(ctx context.Context) {
	if err := sduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sduo *SysDictionaryUpdateOne) defaults() {
	if _, ok := sduo.mutation.UpdatedAt(); !ok {
		v := sysdictionary.UpdateDefaultUpdatedAt()
		sduo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sduo *SysDictionaryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysDictionaryUpdateOne {
	sduo.modifiers = append(sduo.modifiers, modifiers...)
	return sduo
}

func (sduo *SysDictionaryUpdateOne) sqlSave(ctx context.Context) (_node *SysDictionary, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdictionary.Table,
			Columns: sysdictionary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysdictionary.FieldID,
			},
		},
	}
	id, ok := sduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysDictionary.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdictionary.FieldID)
		for _, f := range fields {
			if !sysdictionary.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysdictionary.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sduo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionary.FieldTitle,
		})
	}
	if value, ok := sduo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionary.FieldName,
		})
	}
	if value, ok := sduo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysdictionary.FieldStatus,
		})
	}
	if value, ok := sduo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysdictionary.FieldStatus,
		})
	}
	if sduo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: sysdictionary.FieldStatus,
		})
	}
	if value, ok := sduo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionary.FieldDesc,
		})
	}
	if value, ok := sduo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionary.FieldCreatedAt,
		})
	}
	if value, ok := sduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionary.FieldUpdatedAt,
		})
	}
	if value, ok := sduo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionary.FieldDeletedAt,
		})
	}
	if sduo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysdictionary.FieldDeletedAt,
		})
	}
	_spec.AddModifiers(sduo.modifiers...)
	_node = &SysDictionary{config: sduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdictionary.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
