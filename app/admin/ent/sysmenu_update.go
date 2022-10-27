// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysmenu"
	"slash-admin/pkg/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysMenuUpdate is the builder for updating SysMenu entities.
type SysMenuUpdate struct {
	config
	hooks     []Hook
	mutation  *SysMenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysMenuUpdate builder.
func (smu *SysMenuUpdate) Where(ps ...predicate.SysMenu) *SysMenuUpdate {
	smu.mutation.Where(ps...)
	return smu
}

// SetMenuLevel sets the "menu_level" field.
func (smu *SysMenuUpdate) SetMenuLevel(u uint32) *SysMenuUpdate {
	smu.mutation.ResetMenuLevel()
	smu.mutation.SetMenuLevel(u)
	return smu
}

// AddMenuLevel adds u to the "menu_level" field.
func (smu *SysMenuUpdate) AddMenuLevel(u int32) *SysMenuUpdate {
	smu.mutation.AddMenuLevel(u)
	return smu
}

// SetMenuType sets the "menu_type" field.
func (smu *SysMenuUpdate) SetMenuType(u uint32) *SysMenuUpdate {
	smu.mutation.ResetMenuType()
	smu.mutation.SetMenuType(u)
	return smu
}

// AddMenuType adds u to the "menu_type" field.
func (smu *SysMenuUpdate) AddMenuType(u int32) *SysMenuUpdate {
	smu.mutation.AddMenuType(u)
	return smu
}

// SetParentID sets the "parent_id" field.
func (smu *SysMenuUpdate) SetParentID(u uint) *SysMenuUpdate {
	smu.mutation.ResetParentID()
	smu.mutation.SetParentID(u)
	return smu
}

// AddParentID adds u to the "parent_id" field.
func (smu *SysMenuUpdate) AddParentID(u int) *SysMenuUpdate {
	smu.mutation.AddParentID(u)
	return smu
}

// SetPath sets the "path" field.
func (smu *SysMenuUpdate) SetPath(s string) *SysMenuUpdate {
	smu.mutation.SetPath(s)
	return smu
}

// SetName sets the "name" field.
func (smu *SysMenuUpdate) SetName(s string) *SysMenuUpdate {
	smu.mutation.SetName(s)
	return smu
}

// SetRedirect sets the "redirect" field.
func (smu *SysMenuUpdate) SetRedirect(s string) *SysMenuUpdate {
	smu.mutation.SetRedirect(s)
	return smu
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableRedirect(s *string) *SysMenuUpdate {
	if s != nil {
		smu.SetRedirect(*s)
	}
	return smu
}

// ClearRedirect clears the value of the "redirect" field.
func (smu *SysMenuUpdate) ClearRedirect() *SysMenuUpdate {
	smu.mutation.ClearRedirect()
	return smu
}

// SetComponent sets the "component" field.
func (smu *SysMenuUpdate) SetComponent(s string) *SysMenuUpdate {
	smu.mutation.SetComponent(s)
	return smu
}

// SetOrderNo sets the "order_no" field.
func (smu *SysMenuUpdate) SetOrderNo(u uint32) *SysMenuUpdate {
	smu.mutation.ResetOrderNo()
	smu.mutation.SetOrderNo(u)
	return smu
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableOrderNo(u *uint32) *SysMenuUpdate {
	if u != nil {
		smu.SetOrderNo(*u)
	}
	return smu
}

// AddOrderNo adds u to the "order_no" field.
func (smu *SysMenuUpdate) AddOrderNo(u int32) *SysMenuUpdate {
	smu.mutation.AddOrderNo(u)
	return smu
}

// SetDisabled sets the "disabled" field.
func (smu *SysMenuUpdate) SetDisabled(b bool) *SysMenuUpdate {
	smu.mutation.SetDisabled(b)
	return smu
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableDisabled(b *bool) *SysMenuUpdate {
	if b != nil {
		smu.SetDisabled(*b)
	}
	return smu
}

// SetMeta sets the "meta" field.
func (smu *SysMenuUpdate) SetMeta(tm types.MenuMeta) *SysMenuUpdate {
	smu.mutation.SetMeta(tm)
	return smu
}

// SetNillableMeta sets the "meta" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableMeta(tm *types.MenuMeta) *SysMenuUpdate {
	if tm != nil {
		smu.SetMeta(*tm)
	}
	return smu
}

// ClearMeta clears the value of the "meta" field.
func (smu *SysMenuUpdate) ClearMeta() *SysMenuUpdate {
	smu.mutation.ClearMeta()
	return smu
}

// SetCreatedAt sets the "created_at" field.
func (smu *SysMenuUpdate) SetCreatedAt(t time.Time) *SysMenuUpdate {
	smu.mutation.SetCreatedAt(t)
	return smu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableCreatedAt(t *time.Time) *SysMenuUpdate {
	if t != nil {
		smu.SetCreatedAt(*t)
	}
	return smu
}

// SetUpdatedAt sets the "updated_at" field.
func (smu *SysMenuUpdate) SetUpdatedAt(t time.Time) *SysMenuUpdate {
	smu.mutation.SetUpdatedAt(t)
	return smu
}

// SetDeletedAt sets the "deleted_at" field.
func (smu *SysMenuUpdate) SetDeletedAt(t time.Time) *SysMenuUpdate {
	smu.mutation.SetDeletedAt(t)
	return smu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableDeletedAt(t *time.Time) *SysMenuUpdate {
	if t != nil {
		smu.SetDeletedAt(*t)
	}
	return smu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (smu *SysMenuUpdate) ClearDeletedAt() *SysMenuUpdate {
	smu.mutation.ClearDeletedAt()
	return smu
}

// Mutation returns the SysMenuMutation object of the builder.
func (smu *SysMenuUpdate) Mutation() *SysMenuMutation {
	return smu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (smu *SysMenuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	smu.defaults()
	if len(smu.hooks) == 0 {
		affected, err = smu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smu.mutation = mutation
			affected, err = smu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smu.hooks) - 1; i >= 0; i-- {
			if smu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (smu *SysMenuUpdate) SaveX(ctx context.Context) int {
	affected, err := smu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (smu *SysMenuUpdate) Exec(ctx context.Context) error {
	_, err := smu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smu *SysMenuUpdate) ExecX(ctx context.Context) {
	if err := smu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smu *SysMenuUpdate) defaults() {
	if _, ok := smu.mutation.UpdatedAt(); !ok {
		v := sysmenu.UpdateDefaultUpdatedAt()
		smu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (smu *SysMenuUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysMenuUpdate {
	smu.modifiers = append(smu.modifiers, modifiers...)
	return smu
}

func (smu *SysMenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysmenu.Table,
			Columns: sysmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysmenu.FieldID,
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
	if value, ok := smu.mutation.MenuLevel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldMenuLevel,
		})
	}
	if value, ok := smu.mutation.AddedMenuLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldMenuLevel,
		})
	}
	if value, ok := smu.mutation.MenuType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldMenuType,
		})
	}
	if value, ok := smu.mutation.AddedMenuType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldMenuType,
		})
	}
	if value, ok := smu.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: sysmenu.FieldParentID,
		})
	}
	if value, ok := smu.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: sysmenu.FieldParentID,
		})
	}
	if value, ok := smu.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldPath,
		})
	}
	if value, ok := smu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldName,
		})
	}
	if value, ok := smu.mutation.Redirect(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldRedirect,
		})
	}
	if smu.mutation.RedirectCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysmenu.FieldRedirect,
		})
	}
	if value, ok := smu.mutation.Component(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldComponent,
		})
	}
	if value, ok := smu.mutation.OrderNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldOrderNo,
		})
	}
	if value, ok := smu.mutation.AddedOrderNo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldOrderNo,
		})
	}
	if value, ok := smu.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldDisabled,
		})
	}
	if value, ok := smu.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldMeta,
		})
	}
	if smu.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysmenu.FieldMeta,
		})
	}
	if value, ok := smu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldCreatedAt,
		})
	}
	if value, ok := smu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldUpdatedAt,
		})
	}
	if value, ok := smu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldDeletedAt,
		})
	}
	if smu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysmenu.FieldDeletedAt,
		})
	}
	_spec.AddModifiers(smu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, smu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysmenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SysMenuUpdateOne is the builder for updating a single SysMenu entity.
type SysMenuUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysMenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetMenuLevel sets the "menu_level" field.
func (smuo *SysMenuUpdateOne) SetMenuLevel(u uint32) *SysMenuUpdateOne {
	smuo.mutation.ResetMenuLevel()
	smuo.mutation.SetMenuLevel(u)
	return smuo
}

// AddMenuLevel adds u to the "menu_level" field.
func (smuo *SysMenuUpdateOne) AddMenuLevel(u int32) *SysMenuUpdateOne {
	smuo.mutation.AddMenuLevel(u)
	return smuo
}

// SetMenuType sets the "menu_type" field.
func (smuo *SysMenuUpdateOne) SetMenuType(u uint32) *SysMenuUpdateOne {
	smuo.mutation.ResetMenuType()
	smuo.mutation.SetMenuType(u)
	return smuo
}

// AddMenuType adds u to the "menu_type" field.
func (smuo *SysMenuUpdateOne) AddMenuType(u int32) *SysMenuUpdateOne {
	smuo.mutation.AddMenuType(u)
	return smuo
}

// SetParentID sets the "parent_id" field.
func (smuo *SysMenuUpdateOne) SetParentID(u uint) *SysMenuUpdateOne {
	smuo.mutation.ResetParentID()
	smuo.mutation.SetParentID(u)
	return smuo
}

// AddParentID adds u to the "parent_id" field.
func (smuo *SysMenuUpdateOne) AddParentID(u int) *SysMenuUpdateOne {
	smuo.mutation.AddParentID(u)
	return smuo
}

// SetPath sets the "path" field.
func (smuo *SysMenuUpdateOne) SetPath(s string) *SysMenuUpdateOne {
	smuo.mutation.SetPath(s)
	return smuo
}

// SetName sets the "name" field.
func (smuo *SysMenuUpdateOne) SetName(s string) *SysMenuUpdateOne {
	smuo.mutation.SetName(s)
	return smuo
}

// SetRedirect sets the "redirect" field.
func (smuo *SysMenuUpdateOne) SetRedirect(s string) *SysMenuUpdateOne {
	smuo.mutation.SetRedirect(s)
	return smuo
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableRedirect(s *string) *SysMenuUpdateOne {
	if s != nil {
		smuo.SetRedirect(*s)
	}
	return smuo
}

// ClearRedirect clears the value of the "redirect" field.
func (smuo *SysMenuUpdateOne) ClearRedirect() *SysMenuUpdateOne {
	smuo.mutation.ClearRedirect()
	return smuo
}

// SetComponent sets the "component" field.
func (smuo *SysMenuUpdateOne) SetComponent(s string) *SysMenuUpdateOne {
	smuo.mutation.SetComponent(s)
	return smuo
}

// SetOrderNo sets the "order_no" field.
func (smuo *SysMenuUpdateOne) SetOrderNo(u uint32) *SysMenuUpdateOne {
	smuo.mutation.ResetOrderNo()
	smuo.mutation.SetOrderNo(u)
	return smuo
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableOrderNo(u *uint32) *SysMenuUpdateOne {
	if u != nil {
		smuo.SetOrderNo(*u)
	}
	return smuo
}

// AddOrderNo adds u to the "order_no" field.
func (smuo *SysMenuUpdateOne) AddOrderNo(u int32) *SysMenuUpdateOne {
	smuo.mutation.AddOrderNo(u)
	return smuo
}

// SetDisabled sets the "disabled" field.
func (smuo *SysMenuUpdateOne) SetDisabled(b bool) *SysMenuUpdateOne {
	smuo.mutation.SetDisabled(b)
	return smuo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableDisabled(b *bool) *SysMenuUpdateOne {
	if b != nil {
		smuo.SetDisabled(*b)
	}
	return smuo
}

// SetMeta sets the "meta" field.
func (smuo *SysMenuUpdateOne) SetMeta(tm types.MenuMeta) *SysMenuUpdateOne {
	smuo.mutation.SetMeta(tm)
	return smuo
}

// SetNillableMeta sets the "meta" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableMeta(tm *types.MenuMeta) *SysMenuUpdateOne {
	if tm != nil {
		smuo.SetMeta(*tm)
	}
	return smuo
}

// ClearMeta clears the value of the "meta" field.
func (smuo *SysMenuUpdateOne) ClearMeta() *SysMenuUpdateOne {
	smuo.mutation.ClearMeta()
	return smuo
}

// SetCreatedAt sets the "created_at" field.
func (smuo *SysMenuUpdateOne) SetCreatedAt(t time.Time) *SysMenuUpdateOne {
	smuo.mutation.SetCreatedAt(t)
	return smuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableCreatedAt(t *time.Time) *SysMenuUpdateOne {
	if t != nil {
		smuo.SetCreatedAt(*t)
	}
	return smuo
}

// SetUpdatedAt sets the "updated_at" field.
func (smuo *SysMenuUpdateOne) SetUpdatedAt(t time.Time) *SysMenuUpdateOne {
	smuo.mutation.SetUpdatedAt(t)
	return smuo
}

// SetDeletedAt sets the "deleted_at" field.
func (smuo *SysMenuUpdateOne) SetDeletedAt(t time.Time) *SysMenuUpdateOne {
	smuo.mutation.SetDeletedAt(t)
	return smuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableDeletedAt(t *time.Time) *SysMenuUpdateOne {
	if t != nil {
		smuo.SetDeletedAt(*t)
	}
	return smuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (smuo *SysMenuUpdateOne) ClearDeletedAt() *SysMenuUpdateOne {
	smuo.mutation.ClearDeletedAt()
	return smuo
}

// Mutation returns the SysMenuMutation object of the builder.
func (smuo *SysMenuUpdateOne) Mutation() *SysMenuMutation {
	return smuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (smuo *SysMenuUpdateOne) Select(field string, fields ...string) *SysMenuUpdateOne {
	smuo.fields = append([]string{field}, fields...)
	return smuo
}

// Save executes the query and returns the updated SysMenu entity.
func (smuo *SysMenuUpdateOne) Save(ctx context.Context) (*SysMenu, error) {
	var (
		err  error
		node *SysMenu
	)
	smuo.defaults()
	if len(smuo.hooks) == 0 {
		node, err = smuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smuo.mutation = mutation
			node, err = smuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(smuo.hooks) - 1; i >= 0; i-- {
			if smuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, smuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysMenu)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysMenuMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (smuo *SysMenuUpdateOne) SaveX(ctx context.Context) *SysMenu {
	node, err := smuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (smuo *SysMenuUpdateOne) Exec(ctx context.Context) error {
	_, err := smuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smuo *SysMenuUpdateOne) ExecX(ctx context.Context) {
	if err := smuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smuo *SysMenuUpdateOne) defaults() {
	if _, ok := smuo.mutation.UpdatedAt(); !ok {
		v := sysmenu.UpdateDefaultUpdatedAt()
		smuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (smuo *SysMenuUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysMenuUpdateOne {
	smuo.modifiers = append(smuo.modifiers, modifiers...)
	return smuo
}

func (smuo *SysMenuUpdateOne) sqlSave(ctx context.Context) (_node *SysMenu, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysmenu.Table,
			Columns: sysmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysmenu.FieldID,
			},
		},
	}
	id, ok := smuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysMenu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := smuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysmenu.FieldID)
		for _, f := range fields {
			if !sysmenu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysmenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := smuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smuo.mutation.MenuLevel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldMenuLevel,
		})
	}
	if value, ok := smuo.mutation.AddedMenuLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldMenuLevel,
		})
	}
	if value, ok := smuo.mutation.MenuType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldMenuType,
		})
	}
	if value, ok := smuo.mutation.AddedMenuType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldMenuType,
		})
	}
	if value, ok := smuo.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: sysmenu.FieldParentID,
		})
	}
	if value, ok := smuo.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: sysmenu.FieldParentID,
		})
	}
	if value, ok := smuo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldPath,
		})
	}
	if value, ok := smuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldName,
		})
	}
	if value, ok := smuo.mutation.Redirect(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldRedirect,
		})
	}
	if smuo.mutation.RedirectCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysmenu.FieldRedirect,
		})
	}
	if value, ok := smuo.mutation.Component(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldComponent,
		})
	}
	if value, ok := smuo.mutation.OrderNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldOrderNo,
		})
	}
	if value, ok := smuo.mutation.AddedOrderNo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysmenu.FieldOrderNo,
		})
	}
	if value, ok := smuo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldDisabled,
		})
	}
	if value, ok := smuo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldMeta,
		})
	}
	if smuo.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysmenu.FieldMeta,
		})
	}
	if value, ok := smuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldCreatedAt,
		})
	}
	if value, ok := smuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldUpdatedAt,
		})
	}
	if value, ok := smuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldDeletedAt,
		})
	}
	if smuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysmenu.FieldDeletedAt,
		})
	}
	_spec.AddModifiers(smuo.modifiers...)
	_node = &SysMenu{config: smuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, smuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysmenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
