// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/app/admin/ent/sysdictionary"
	"slash-admin/app/admin/ent/sysdictionarydetail"
	"slash-admin/pkg/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysDictionaryDetailUpdate is the builder for updating SysDictionaryDetail entities.
type SysDictionaryDetailUpdate struct {
	config
	hooks     []Hook
	mutation  *SysDictionaryDetailMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysDictionaryDetailUpdate builder.
func (sddu *SysDictionaryDetailUpdate) Where(ps ...predicate.SysDictionaryDetail) *SysDictionaryDetailUpdate {
	sddu.mutation.Where(ps...)
	return sddu
}

// SetTitle sets the "title" field.
func (sddu *SysDictionaryDetailUpdate) SetTitle(s string) *SysDictionaryDetailUpdate {
	sddu.mutation.SetTitle(s)
	return sddu
}

// SetKey sets the "key" field.
func (sddu *SysDictionaryDetailUpdate) SetKey(s string) *SysDictionaryDetailUpdate {
	sddu.mutation.SetKey(s)
	return sddu
}

// SetValue sets the "value" field.
func (sddu *SysDictionaryDetailUpdate) SetValue(s string) *SysDictionaryDetailUpdate {
	sddu.mutation.SetValue(s)
	return sddu
}

// SetStatus sets the "status" field.
func (sddu *SysDictionaryDetailUpdate) SetStatus(t types.Status) *SysDictionaryDetailUpdate {
	sddu.mutation.ResetStatus()
	sddu.mutation.SetStatus(t)
	return sddu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sddu *SysDictionaryDetailUpdate) SetNillableStatus(t *types.Status) *SysDictionaryDetailUpdate {
	if t != nil {
		sddu.SetStatus(*t)
	}
	return sddu
}

// AddStatus adds t to the "status" field.
func (sddu *SysDictionaryDetailUpdate) AddStatus(t types.Status) *SysDictionaryDetailUpdate {
	sddu.mutation.AddStatus(t)
	return sddu
}

// ClearStatus clears the value of the "status" field.
func (sddu *SysDictionaryDetailUpdate) ClearStatus() *SysDictionaryDetailUpdate {
	sddu.mutation.ClearStatus()
	return sddu
}

// SetDictionaryID sets the "dictionary_id" field.
func (sddu *SysDictionaryDetailUpdate) SetDictionaryID(u uint64) *SysDictionaryDetailUpdate {
	sddu.mutation.SetDictionaryID(u)
	return sddu
}

// SetNillableDictionaryID sets the "dictionary_id" field if the given value is not nil.
func (sddu *SysDictionaryDetailUpdate) SetNillableDictionaryID(u *uint64) *SysDictionaryDetailUpdate {
	if u != nil {
		sddu.SetDictionaryID(*u)
	}
	return sddu
}

// ClearDictionaryID clears the value of the "dictionary_id" field.
func (sddu *SysDictionaryDetailUpdate) ClearDictionaryID() *SysDictionaryDetailUpdate {
	sddu.mutation.ClearDictionaryID()
	return sddu
}

// SetRemark sets the "remark" field.
func (sddu *SysDictionaryDetailUpdate) SetRemark(s string) *SysDictionaryDetailUpdate {
	sddu.mutation.SetRemark(s)
	return sddu
}

// SetOrderNo sets the "order_no" field.
func (sddu *SysDictionaryDetailUpdate) SetOrderNo(u uint32) *SysDictionaryDetailUpdate {
	sddu.mutation.ResetOrderNo()
	sddu.mutation.SetOrderNo(u)
	return sddu
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (sddu *SysDictionaryDetailUpdate) SetNillableOrderNo(u *uint32) *SysDictionaryDetailUpdate {
	if u != nil {
		sddu.SetOrderNo(*u)
	}
	return sddu
}

// AddOrderNo adds u to the "order_no" field.
func (sddu *SysDictionaryDetailUpdate) AddOrderNo(u int32) *SysDictionaryDetailUpdate {
	sddu.mutation.AddOrderNo(u)
	return sddu
}

// SetCreatedAt sets the "created_at" field.
func (sddu *SysDictionaryDetailUpdate) SetCreatedAt(t time.Time) *SysDictionaryDetailUpdate {
	sddu.mutation.SetCreatedAt(t)
	return sddu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sddu *SysDictionaryDetailUpdate) SetNillableCreatedAt(t *time.Time) *SysDictionaryDetailUpdate {
	if t != nil {
		sddu.SetCreatedAt(*t)
	}
	return sddu
}

// SetUpdatedAt sets the "updated_at" field.
func (sddu *SysDictionaryDetailUpdate) SetUpdatedAt(t time.Time) *SysDictionaryDetailUpdate {
	sddu.mutation.SetUpdatedAt(t)
	return sddu
}

// SetDeletedAt sets the "deleted_at" field.
func (sddu *SysDictionaryDetailUpdate) SetDeletedAt(t time.Time) *SysDictionaryDetailUpdate {
	sddu.mutation.SetDeletedAt(t)
	return sddu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sddu *SysDictionaryDetailUpdate) SetNillableDeletedAt(t *time.Time) *SysDictionaryDetailUpdate {
	if t != nil {
		sddu.SetDeletedAt(*t)
	}
	return sddu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sddu *SysDictionaryDetailUpdate) ClearDeletedAt() *SysDictionaryDetailUpdate {
	sddu.mutation.ClearDeletedAt()
	return sddu
}

// SetParentID sets the "parent" edge to the SysDictionary entity by ID.
func (sddu *SysDictionaryDetailUpdate) SetParentID(id uint64) *SysDictionaryDetailUpdate {
	sddu.mutation.SetParentID(id)
	return sddu
}

// SetNillableParentID sets the "parent" edge to the SysDictionary entity by ID if the given value is not nil.
func (sddu *SysDictionaryDetailUpdate) SetNillableParentID(id *uint64) *SysDictionaryDetailUpdate {
	if id != nil {
		sddu = sddu.SetParentID(*id)
	}
	return sddu
}

// SetParent sets the "parent" edge to the SysDictionary entity.
func (sddu *SysDictionaryDetailUpdate) SetParent(s *SysDictionary) *SysDictionaryDetailUpdate {
	return sddu.SetParentID(s.ID)
}

// Mutation returns the SysDictionaryDetailMutation object of the builder.
func (sddu *SysDictionaryDetailUpdate) Mutation() *SysDictionaryDetailMutation {
	return sddu.mutation
}

// ClearParent clears the "parent" edge to the SysDictionary entity.
func (sddu *SysDictionaryDetailUpdate) ClearParent() *SysDictionaryDetailUpdate {
	sddu.mutation.ClearParent()
	return sddu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sddu *SysDictionaryDetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sddu.defaults()
	if len(sddu.hooks) == 0 {
		affected, err = sddu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictionaryDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sddu.mutation = mutation
			affected, err = sddu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sddu.hooks) - 1; i >= 0; i-- {
			if sddu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sddu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sddu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sddu *SysDictionaryDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := sddu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sddu *SysDictionaryDetailUpdate) Exec(ctx context.Context) error {
	_, err := sddu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sddu *SysDictionaryDetailUpdate) ExecX(ctx context.Context) {
	if err := sddu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sddu *SysDictionaryDetailUpdate) defaults() {
	if _, ok := sddu.mutation.UpdatedAt(); !ok {
		v := sysdictionarydetail.UpdateDefaultUpdatedAt()
		sddu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sddu *SysDictionaryDetailUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysDictionaryDetailUpdate {
	sddu.modifiers = append(sddu.modifiers, modifiers...)
	return sddu
}

func (sddu *SysDictionaryDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdictionarydetail.Table,
			Columns: sysdictionarydetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysdictionarydetail.FieldID,
			},
		},
	}
	if ps := sddu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sddu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldTitle,
		})
	}
	if value, ok := sddu.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldKey,
		})
	}
	if value, ok := sddu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldValue,
		})
	}
	if value, ok := sddu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysdictionarydetail.FieldStatus,
		})
	}
	if value, ok := sddu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysdictionarydetail.FieldStatus,
		})
	}
	if sddu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: sysdictionarydetail.FieldStatus,
		})
	}
	if value, ok := sddu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldRemark,
		})
	}
	if value, ok := sddu.mutation.OrderNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysdictionarydetail.FieldOrderNo,
		})
	}
	if value, ok := sddu.mutation.AddedOrderNo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysdictionarydetail.FieldOrderNo,
		})
	}
	if value, ok := sddu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionarydetail.FieldCreatedAt,
		})
	}
	if value, ok := sddu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionarydetail.FieldUpdatedAt,
		})
	}
	if value, ok := sddu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionarydetail.FieldDeletedAt,
		})
	}
	if sddu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysdictionarydetail.FieldDeletedAt,
		})
	}
	if sddu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictionarydetail.ParentTable,
			Columns: []string{sysdictionarydetail.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysdictionary.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sddu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictionarydetail.ParentTable,
			Columns: []string{sysdictionarydetail.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysdictionary.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(sddu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sddu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdictionarydetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SysDictionaryDetailUpdateOne is the builder for updating a single SysDictionaryDetail entity.
type SysDictionaryDetailUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysDictionaryDetailMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTitle sets the "title" field.
func (sdduo *SysDictionaryDetailUpdateOne) SetTitle(s string) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.SetTitle(s)
	return sdduo
}

// SetKey sets the "key" field.
func (sdduo *SysDictionaryDetailUpdateOne) SetKey(s string) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.SetKey(s)
	return sdduo
}

// SetValue sets the "value" field.
func (sdduo *SysDictionaryDetailUpdateOne) SetValue(s string) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.SetValue(s)
	return sdduo
}

// SetStatus sets the "status" field.
func (sdduo *SysDictionaryDetailUpdateOne) SetStatus(t types.Status) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.ResetStatus()
	sdduo.mutation.SetStatus(t)
	return sdduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sdduo *SysDictionaryDetailUpdateOne) SetNillableStatus(t *types.Status) *SysDictionaryDetailUpdateOne {
	if t != nil {
		sdduo.SetStatus(*t)
	}
	return sdduo
}

// AddStatus adds t to the "status" field.
func (sdduo *SysDictionaryDetailUpdateOne) AddStatus(t types.Status) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.AddStatus(t)
	return sdduo
}

// ClearStatus clears the value of the "status" field.
func (sdduo *SysDictionaryDetailUpdateOne) ClearStatus() *SysDictionaryDetailUpdateOne {
	sdduo.mutation.ClearStatus()
	return sdduo
}

// SetDictionaryID sets the "dictionary_id" field.
func (sdduo *SysDictionaryDetailUpdateOne) SetDictionaryID(u uint64) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.SetDictionaryID(u)
	return sdduo
}

// SetNillableDictionaryID sets the "dictionary_id" field if the given value is not nil.
func (sdduo *SysDictionaryDetailUpdateOne) SetNillableDictionaryID(u *uint64) *SysDictionaryDetailUpdateOne {
	if u != nil {
		sdduo.SetDictionaryID(*u)
	}
	return sdduo
}

// ClearDictionaryID clears the value of the "dictionary_id" field.
func (sdduo *SysDictionaryDetailUpdateOne) ClearDictionaryID() *SysDictionaryDetailUpdateOne {
	sdduo.mutation.ClearDictionaryID()
	return sdduo
}

// SetRemark sets the "remark" field.
func (sdduo *SysDictionaryDetailUpdateOne) SetRemark(s string) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.SetRemark(s)
	return sdduo
}

// SetOrderNo sets the "order_no" field.
func (sdduo *SysDictionaryDetailUpdateOne) SetOrderNo(u uint32) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.ResetOrderNo()
	sdduo.mutation.SetOrderNo(u)
	return sdduo
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (sdduo *SysDictionaryDetailUpdateOne) SetNillableOrderNo(u *uint32) *SysDictionaryDetailUpdateOne {
	if u != nil {
		sdduo.SetOrderNo(*u)
	}
	return sdduo
}

// AddOrderNo adds u to the "order_no" field.
func (sdduo *SysDictionaryDetailUpdateOne) AddOrderNo(u int32) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.AddOrderNo(u)
	return sdduo
}

// SetCreatedAt sets the "created_at" field.
func (sdduo *SysDictionaryDetailUpdateOne) SetCreatedAt(t time.Time) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.SetCreatedAt(t)
	return sdduo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sdduo *SysDictionaryDetailUpdateOne) SetNillableCreatedAt(t *time.Time) *SysDictionaryDetailUpdateOne {
	if t != nil {
		sdduo.SetCreatedAt(*t)
	}
	return sdduo
}

// SetUpdatedAt sets the "updated_at" field.
func (sdduo *SysDictionaryDetailUpdateOne) SetUpdatedAt(t time.Time) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.SetUpdatedAt(t)
	return sdduo
}

// SetDeletedAt sets the "deleted_at" field.
func (sdduo *SysDictionaryDetailUpdateOne) SetDeletedAt(t time.Time) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.SetDeletedAt(t)
	return sdduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdduo *SysDictionaryDetailUpdateOne) SetNillableDeletedAt(t *time.Time) *SysDictionaryDetailUpdateOne {
	if t != nil {
		sdduo.SetDeletedAt(*t)
	}
	return sdduo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sdduo *SysDictionaryDetailUpdateOne) ClearDeletedAt() *SysDictionaryDetailUpdateOne {
	sdduo.mutation.ClearDeletedAt()
	return sdduo
}

// SetParentID sets the "parent" edge to the SysDictionary entity by ID.
func (sdduo *SysDictionaryDetailUpdateOne) SetParentID(id uint64) *SysDictionaryDetailUpdateOne {
	sdduo.mutation.SetParentID(id)
	return sdduo
}

// SetNillableParentID sets the "parent" edge to the SysDictionary entity by ID if the given value is not nil.
func (sdduo *SysDictionaryDetailUpdateOne) SetNillableParentID(id *uint64) *SysDictionaryDetailUpdateOne {
	if id != nil {
		sdduo = sdduo.SetParentID(*id)
	}
	return sdduo
}

// SetParent sets the "parent" edge to the SysDictionary entity.
func (sdduo *SysDictionaryDetailUpdateOne) SetParent(s *SysDictionary) *SysDictionaryDetailUpdateOne {
	return sdduo.SetParentID(s.ID)
}

// Mutation returns the SysDictionaryDetailMutation object of the builder.
func (sdduo *SysDictionaryDetailUpdateOne) Mutation() *SysDictionaryDetailMutation {
	return sdduo.mutation
}

// ClearParent clears the "parent" edge to the SysDictionary entity.
func (sdduo *SysDictionaryDetailUpdateOne) ClearParent() *SysDictionaryDetailUpdateOne {
	sdduo.mutation.ClearParent()
	return sdduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sdduo *SysDictionaryDetailUpdateOne) Select(field string, fields ...string) *SysDictionaryDetailUpdateOne {
	sdduo.fields = append([]string{field}, fields...)
	return sdduo
}

// Save executes the query and returns the updated SysDictionaryDetail entity.
func (sdduo *SysDictionaryDetailUpdateOne) Save(ctx context.Context) (*SysDictionaryDetail, error) {
	var (
		err  error
		node *SysDictionaryDetail
	)
	sdduo.defaults()
	if len(sdduo.hooks) == 0 {
		node, err = sdduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictionaryDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdduo.mutation = mutation
			node, err = sdduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sdduo.hooks) - 1; i >= 0; i-- {
			if sdduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sdduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysDictionaryDetail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysDictionaryDetailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sdduo *SysDictionaryDetailUpdateOne) SaveX(ctx context.Context) *SysDictionaryDetail {
	node, err := sdduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sdduo *SysDictionaryDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := sdduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdduo *SysDictionaryDetailUpdateOne) ExecX(ctx context.Context) {
	if err := sdduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdduo *SysDictionaryDetailUpdateOne) defaults() {
	if _, ok := sdduo.mutation.UpdatedAt(); !ok {
		v := sysdictionarydetail.UpdateDefaultUpdatedAt()
		sdduo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sdduo *SysDictionaryDetailUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysDictionaryDetailUpdateOne {
	sdduo.modifiers = append(sdduo.modifiers, modifiers...)
	return sdduo
}

func (sdduo *SysDictionaryDetailUpdateOne) sqlSave(ctx context.Context) (_node *SysDictionaryDetail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdictionarydetail.Table,
			Columns: sysdictionarydetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sysdictionarydetail.FieldID,
			},
		},
	}
	id, ok := sdduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysDictionaryDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sdduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdictionarydetail.FieldID)
		for _, f := range fields {
			if !sysdictionarydetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysdictionarydetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sdduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdduo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldTitle,
		})
	}
	if value, ok := sdduo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldKey,
		})
	}
	if value, ok := sdduo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldValue,
		})
	}
	if value, ok := sdduo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysdictionarydetail.FieldStatus,
		})
	}
	if value, ok := sdduo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: sysdictionarydetail.FieldStatus,
		})
	}
	if sdduo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: sysdictionarydetail.FieldStatus,
		})
	}
	if value, ok := sdduo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictionarydetail.FieldRemark,
		})
	}
	if value, ok := sdduo.mutation.OrderNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysdictionarydetail.FieldOrderNo,
		})
	}
	if value, ok := sdduo.mutation.AddedOrderNo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sysdictionarydetail.FieldOrderNo,
		})
	}
	if value, ok := sdduo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionarydetail.FieldCreatedAt,
		})
	}
	if value, ok := sdduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionarydetail.FieldUpdatedAt,
		})
	}
	if value, ok := sdduo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictionarydetail.FieldDeletedAt,
		})
	}
	if sdduo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysdictionarydetail.FieldDeletedAt,
		})
	}
	if sdduo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictionarydetail.ParentTable,
			Columns: []string{sysdictionarydetail.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysdictionary.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdduo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictionarydetail.ParentTable,
			Columns: []string{sysdictionarydetail.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: sysdictionary.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(sdduo.modifiers...)
	_node = &SysDictionaryDetail{config: sdduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sdduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdictionarydetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
