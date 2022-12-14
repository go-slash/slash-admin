// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slash-admin/app/admin/ent/sysmenu"
	"slash-admin/app/admin/ent/sysrole"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysRoleCreate is the builder for creating a SysRole entity.
type SysRoleCreate struct {
	config
	mutation *SysRoleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (src *SysRoleCreate) SetName(s string) *SysRoleCreate {
	src.mutation.SetName(s)
	return src
}

// SetValue sets the "value" field.
func (src *SysRoleCreate) SetValue(s string) *SysRoleCreate {
	src.mutation.SetValue(s)
	return src
}

// SetDefaultRouter sets the "default_router" field.
func (src *SysRoleCreate) SetDefaultRouter(s string) *SysRoleCreate {
	src.mutation.SetDefaultRouter(s)
	return src
}

// SetNillableDefaultRouter sets the "default_router" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableDefaultRouter(s *string) *SysRoleCreate {
	if s != nil {
		src.SetDefaultRouter(*s)
	}
	return src
}

// SetStatus sets the "status" field.
func (src *SysRoleCreate) SetStatus(t types.Status) *SysRoleCreate {
	src.mutation.SetStatus(t)
	return src
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableStatus(t *types.Status) *SysRoleCreate {
	if t != nil {
		src.SetStatus(*t)
	}
	return src
}

// SetRemark sets the "remark" field.
func (src *SysRoleCreate) SetRemark(s string) *SysRoleCreate {
	src.mutation.SetRemark(s)
	return src
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableRemark(s *string) *SysRoleCreate {
	if s != nil {
		src.SetRemark(*s)
	}
	return src
}

// SetOrderNo sets the "order_no" field.
func (src *SysRoleCreate) SetOrderNo(u uint32) *SysRoleCreate {
	src.mutation.SetOrderNo(u)
	return src
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableOrderNo(u *uint32) *SysRoleCreate {
	if u != nil {
		src.SetOrderNo(*u)
	}
	return src
}

// SetCreatedAt sets the "created_at" field.
func (src *SysRoleCreate) SetCreatedAt(t time.Time) *SysRoleCreate {
	src.mutation.SetCreatedAt(t)
	return src
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableCreatedAt(t *time.Time) *SysRoleCreate {
	if t != nil {
		src.SetCreatedAt(*t)
	}
	return src
}

// SetUpdatedAt sets the "updated_at" field.
func (src *SysRoleCreate) SetUpdatedAt(t time.Time) *SysRoleCreate {
	src.mutation.SetUpdatedAt(t)
	return src
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableUpdatedAt(t *time.Time) *SysRoleCreate {
	if t != nil {
		src.SetUpdatedAt(*t)
	}
	return src
}

// SetDeletedAt sets the "deleted_at" field.
func (src *SysRoleCreate) SetDeletedAt(t time.Time) *SysRoleCreate {
	src.mutation.SetDeletedAt(t)
	return src
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableDeletedAt(t *time.Time) *SysRoleCreate {
	if t != nil {
		src.SetDeletedAt(*t)
	}
	return src
}

// SetID sets the "id" field.
func (src *SysRoleCreate) SetID(u uint64) *SysRoleCreate {
	src.mutation.SetID(u)
	return src
}

// AddMenuIDs adds the "menus" edge to the SysMenu entity by IDs.
func (src *SysRoleCreate) AddMenuIDs(ids ...uint64) *SysRoleCreate {
	src.mutation.AddMenuIDs(ids...)
	return src
}

// AddMenus adds the "menus" edges to the SysMenu entity.
func (src *SysRoleCreate) AddMenus(s ...*SysMenu) *SysRoleCreate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return src.AddMenuIDs(ids...)
}

// AddRoleIDs adds the "role" edge to the SysUser entity by IDs.
func (src *SysRoleCreate) AddRoleIDs(ids ...uint64) *SysRoleCreate {
	src.mutation.AddRoleIDs(ids...)
	return src
}

// AddRole adds the "role" edges to the SysUser entity.
func (src *SysRoleCreate) AddRole(s ...*SysUser) *SysRoleCreate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return src.AddRoleIDs(ids...)
}

// Mutation returns the SysRoleMutation object of the builder.
func (src *SysRoleCreate) Mutation() *SysRoleMutation {
	return src.mutation
}

// Save creates the SysRole in the database.
func (src *SysRoleCreate) Save(ctx context.Context) (*SysRole, error) {
	var (
		err  error
		node *SysRole
	)
	src.defaults()
	if len(src.hooks) == 0 {
		if err = src.check(); err != nil {
			return nil, err
		}
		node, err = src.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = src.check(); err != nil {
				return nil, err
			}
			src.mutation = mutation
			if node, err = src.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(src.hooks) - 1; i >= 0; i-- {
			if src.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = src.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, src.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysRole)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysRoleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (src *SysRoleCreate) SaveX(ctx context.Context) *SysRole {
	v, err := src.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (src *SysRoleCreate) Exec(ctx context.Context) error {
	_, err := src.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (src *SysRoleCreate) ExecX(ctx context.Context) {
	if err := src.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (src *SysRoleCreate) defaults() {
	if _, ok := src.mutation.DefaultRouter(); !ok {
		v := sysrole.DefaultDefaultRouter
		src.mutation.SetDefaultRouter(v)
	}
	if _, ok := src.mutation.Status(); !ok {
		v := sysrole.DefaultStatus
		src.mutation.SetStatus(v)
	}
	if _, ok := src.mutation.Remark(); !ok {
		v := sysrole.DefaultRemark
		src.mutation.SetRemark(v)
	}
	if _, ok := src.mutation.OrderNo(); !ok {
		v := sysrole.DefaultOrderNo
		src.mutation.SetOrderNo(v)
	}
	if _, ok := src.mutation.CreatedAt(); !ok {
		v := sysrole.DefaultCreatedAt()
		src.mutation.SetCreatedAt(v)
	}
	if _, ok := src.mutation.UpdatedAt(); !ok {
		v := sysrole.DefaultUpdatedAt()
		src.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (src *SysRoleCreate) check() error {
	if _, ok := src.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "SysRole.name"`)}
	}
	if _, ok := src.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "SysRole.value"`)}
	}
	if _, ok := src.mutation.DefaultRouter(); !ok {
		return &ValidationError{Name: "default_router", err: errors.New(`ent: missing required field "SysRole.default_router"`)}
	}
	if _, ok := src.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "SysRole.remark"`)}
	}
	if _, ok := src.mutation.OrderNo(); !ok {
		return &ValidationError{Name: "order_no", err: errors.New(`ent: missing required field "SysRole.order_no"`)}
	}
	if _, ok := src.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SysRole.created_at"`)}
	}
	if _, ok := src.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SysRole.updated_at"`)}
	}
	return nil
}

func (src *SysRoleCreate) sqlSave(ctx context.Context) (*SysRole, error) {
	_node, _spec := src.createSpec()
	if err := sqlgraph.CreateNode(ctx, src.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (src *SysRoleCreate) createSpec() (*SysRole, *sqlgraph.CreateSpec) {
	var (
		_node = &SysRole{config: src.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysrole.FieldID,
			},
		}
	)
	_spec.OnConflict = src.conflict
	if id, ok := src.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := src.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldName,
		})
		_node.Name = value
	}
	if value, ok := src.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := src.mutation.DefaultRouter(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldDefaultRouter,
		})
		_node.DefaultRouter = value
	}
	if value, ok := src.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := src.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := src.mutation.OrderNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysrole.FieldOrderNo,
		})
		_node.OrderNo = value
	}
	if value, ok := src.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := src.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := src.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := src.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.MenusTable,
			Columns: sysrole.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := src.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysrole.RoleTable,
			Columns: []string{sysrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysRole.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysRoleUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (src *SysRoleCreate) OnConflict(opts ...sql.ConflictOption) *SysRoleUpsertOne {
	src.conflict = opts
	return &SysRoleUpsertOne{
		create: src,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysRole.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (src *SysRoleCreate) OnConflictColumns(columns ...string) *SysRoleUpsertOne {
	src.conflict = append(src.conflict, sql.ConflictColumns(columns...))
	return &SysRoleUpsertOne{
		create: src,
	}
}

type (
	// SysRoleUpsertOne is the builder for "upsert"-ing
	//  one SysRole node.
	SysRoleUpsertOne struct {
		create *SysRoleCreate
	}

	// SysRoleUpsert is the "OnConflict" setter.
	SysRoleUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *SysRoleUpsert) SetName(v string) *SysRoleUpsert {
	u.Set(sysrole.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SysRoleUpsert) UpdateName() *SysRoleUpsert {
	u.SetExcluded(sysrole.FieldName)
	return u
}

// SetValue sets the "value" field.
func (u *SysRoleUpsert) SetValue(v string) *SysRoleUpsert {
	u.Set(sysrole.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysRoleUpsert) UpdateValue() *SysRoleUpsert {
	u.SetExcluded(sysrole.FieldValue)
	return u
}

// SetDefaultRouter sets the "default_router" field.
func (u *SysRoleUpsert) SetDefaultRouter(v string) *SysRoleUpsert {
	u.Set(sysrole.FieldDefaultRouter, v)
	return u
}

// UpdateDefaultRouter sets the "default_router" field to the value that was provided on create.
func (u *SysRoleUpsert) UpdateDefaultRouter() *SysRoleUpsert {
	u.SetExcluded(sysrole.FieldDefaultRouter)
	return u
}

// SetStatus sets the "status" field.
func (u *SysRoleUpsert) SetStatus(v types.Status) *SysRoleUpsert {
	u.Set(sysrole.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *SysRoleUpsert) UpdateStatus() *SysRoleUpsert {
	u.SetExcluded(sysrole.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *SysRoleUpsert) AddStatus(v types.Status) *SysRoleUpsert {
	u.Add(sysrole.FieldStatus, v)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *SysRoleUpsert) ClearStatus() *SysRoleUpsert {
	u.SetNull(sysrole.FieldStatus)
	return u
}

// SetRemark sets the "remark" field.
func (u *SysRoleUpsert) SetRemark(v string) *SysRoleUpsert {
	u.Set(sysrole.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *SysRoleUpsert) UpdateRemark() *SysRoleUpsert {
	u.SetExcluded(sysrole.FieldRemark)
	return u
}

// SetOrderNo sets the "order_no" field.
func (u *SysRoleUpsert) SetOrderNo(v uint32) *SysRoleUpsert {
	u.Set(sysrole.FieldOrderNo, v)
	return u
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *SysRoleUpsert) UpdateOrderNo() *SysRoleUpsert {
	u.SetExcluded(sysrole.FieldOrderNo)
	return u
}

// AddOrderNo adds v to the "order_no" field.
func (u *SysRoleUpsert) AddOrderNo(v uint32) *SysRoleUpsert {
	u.Add(sysrole.FieldOrderNo, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SysRoleUpsert) SetCreatedAt(v time.Time) *SysRoleUpsert {
	u.Set(sysrole.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SysRoleUpsert) UpdateCreatedAt() *SysRoleUpsert {
	u.SetExcluded(sysrole.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysRoleUpsert) SetUpdatedAt(v time.Time) *SysRoleUpsert {
	u.Set(sysrole.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysRoleUpsert) UpdateUpdatedAt() *SysRoleUpsert {
	u.SetExcluded(sysrole.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysRoleUpsert) SetDeletedAt(v time.Time) *SysRoleUpsert {
	u.Set(sysrole.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysRoleUpsert) UpdateDeletedAt() *SysRoleUpsert {
	u.SetExcluded(sysrole.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysRoleUpsert) ClearDeletedAt() *SysRoleUpsert {
	u.SetNull(sysrole.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysRole.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysrole.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysRoleUpsertOne) UpdateNewValues() *SysRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysrole.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysRole.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysRoleUpsertOne) Ignore() *SysRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysRoleUpsertOne) DoNothing() *SysRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysRoleCreate.OnConflict
// documentation for more info.
func (u *SysRoleUpsertOne) Update(set func(*SysRoleUpsert)) *SysRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysRoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *SysRoleUpsertOne) SetName(v string) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SysRoleUpsertOne) UpdateName() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateName()
	})
}

// SetValue sets the "value" field.
func (u *SysRoleUpsertOne) SetValue(v string) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysRoleUpsertOne) UpdateValue() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateValue()
	})
}

// SetDefaultRouter sets the "default_router" field.
func (u *SysRoleUpsertOne) SetDefaultRouter(v string) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetDefaultRouter(v)
	})
}

// UpdateDefaultRouter sets the "default_router" field to the value that was provided on create.
func (u *SysRoleUpsertOne) UpdateDefaultRouter() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateDefaultRouter()
	})
}

// SetStatus sets the "status" field.
func (u *SysRoleUpsertOne) SetStatus(v types.Status) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *SysRoleUpsertOne) AddStatus(v types.Status) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *SysRoleUpsertOne) UpdateStatus() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *SysRoleUpsertOne) ClearStatus() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.ClearStatus()
	})
}

// SetRemark sets the "remark" field.
func (u *SysRoleUpsertOne) SetRemark(v string) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *SysRoleUpsertOne) UpdateRemark() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateRemark()
	})
}

// SetOrderNo sets the "order_no" field.
func (u *SysRoleUpsertOne) SetOrderNo(v uint32) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetOrderNo(v)
	})
}

// AddOrderNo adds v to the "order_no" field.
func (u *SysRoleUpsertOne) AddOrderNo(v uint32) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.AddOrderNo(v)
	})
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *SysRoleUpsertOne) UpdateOrderNo() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateOrderNo()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SysRoleUpsertOne) SetCreatedAt(v time.Time) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SysRoleUpsertOne) UpdateCreatedAt() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysRoleUpsertOne) SetUpdatedAt(v time.Time) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysRoleUpsertOne) UpdateUpdatedAt() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysRoleUpsertOne) SetDeletedAt(v time.Time) *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysRoleUpsertOne) UpdateDeletedAt() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysRoleUpsertOne) ClearDeletedAt() *SysRoleUpsertOne {
	return u.Update(func(s *SysRoleUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *SysRoleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysRoleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysRoleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysRoleUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysRoleUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysRoleCreateBulk is the builder for creating many SysRole entities in bulk.
type SysRoleCreateBulk struct {
	config
	builders []*SysRoleCreate
	conflict []sql.ConflictOption
}

// Save creates the SysRole entities in the database.
func (srcb *SysRoleCreateBulk) Save(ctx context.Context) ([]*SysRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(srcb.builders))
	nodes := make([]*SysRole, len(srcb.builders))
	mutators := make([]Mutator, len(srcb.builders))
	for i := range srcb.builders {
		func(i int, root context.Context) {
			builder := srcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysRoleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, srcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = srcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, srcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, srcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (srcb *SysRoleCreateBulk) SaveX(ctx context.Context) []*SysRole {
	v, err := srcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (srcb *SysRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := srcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srcb *SysRoleCreateBulk) ExecX(ctx context.Context) {
	if err := srcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysRole.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysRoleUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (srcb *SysRoleCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysRoleUpsertBulk {
	srcb.conflict = opts
	return &SysRoleUpsertBulk{
		create: srcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysRole.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (srcb *SysRoleCreateBulk) OnConflictColumns(columns ...string) *SysRoleUpsertBulk {
	srcb.conflict = append(srcb.conflict, sql.ConflictColumns(columns...))
	return &SysRoleUpsertBulk{
		create: srcb,
	}
}

// SysRoleUpsertBulk is the builder for "upsert"-ing
// a bulk of SysRole nodes.
type SysRoleUpsertBulk struct {
	create *SysRoleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysRole.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysrole.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysRoleUpsertBulk) UpdateNewValues() *SysRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysrole.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysRole.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysRoleUpsertBulk) Ignore() *SysRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysRoleUpsertBulk) DoNothing() *SysRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysRoleCreateBulk.OnConflict
// documentation for more info.
func (u *SysRoleUpsertBulk) Update(set func(*SysRoleUpsert)) *SysRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysRoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *SysRoleUpsertBulk) SetName(v string) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SysRoleUpsertBulk) UpdateName() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateName()
	})
}

// SetValue sets the "value" field.
func (u *SysRoleUpsertBulk) SetValue(v string) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysRoleUpsertBulk) UpdateValue() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateValue()
	})
}

// SetDefaultRouter sets the "default_router" field.
func (u *SysRoleUpsertBulk) SetDefaultRouter(v string) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetDefaultRouter(v)
	})
}

// UpdateDefaultRouter sets the "default_router" field to the value that was provided on create.
func (u *SysRoleUpsertBulk) UpdateDefaultRouter() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateDefaultRouter()
	})
}

// SetStatus sets the "status" field.
func (u *SysRoleUpsertBulk) SetStatus(v types.Status) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *SysRoleUpsertBulk) AddStatus(v types.Status) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *SysRoleUpsertBulk) UpdateStatus() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *SysRoleUpsertBulk) ClearStatus() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.ClearStatus()
	})
}

// SetRemark sets the "remark" field.
func (u *SysRoleUpsertBulk) SetRemark(v string) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *SysRoleUpsertBulk) UpdateRemark() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateRemark()
	})
}

// SetOrderNo sets the "order_no" field.
func (u *SysRoleUpsertBulk) SetOrderNo(v uint32) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetOrderNo(v)
	})
}

// AddOrderNo adds v to the "order_no" field.
func (u *SysRoleUpsertBulk) AddOrderNo(v uint32) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.AddOrderNo(v)
	})
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *SysRoleUpsertBulk) UpdateOrderNo() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateOrderNo()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SysRoleUpsertBulk) SetCreatedAt(v time.Time) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SysRoleUpsertBulk) UpdateCreatedAt() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysRoleUpsertBulk) SetUpdatedAt(v time.Time) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysRoleUpsertBulk) UpdateUpdatedAt() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysRoleUpsertBulk) SetDeletedAt(v time.Time) *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysRoleUpsertBulk) UpdateDeletedAt() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysRoleUpsertBulk) ClearDeletedAt() *SysRoleUpsertBulk {
	return u.Update(func(s *SysRoleUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *SysRoleUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SysRoleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysRoleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysRoleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
