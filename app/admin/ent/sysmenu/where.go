// Code generated by ent, DO NOT EDIT.

package sysmenu

import (
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/pkg/types"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// MenuLevel applies equality check predicate on the "menu_level" field. It's identical to MenuLevelEQ.
func MenuLevel(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMenuLevel), v))
	})
}

// MenuType applies equality check predicate on the "menu_type" field. It's identical to MenuTypeEQ.
func MenuType(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMenuType), v))
	})
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v uint) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentID), v))
	})
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Redirect applies equality check predicate on the "redirect" field. It's identical to RedirectEQ.
func Redirect(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRedirect), v))
	})
}

// Component applies equality check predicate on the "component" field. It's identical to ComponentEQ.
func Component(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComponent), v))
	})
}

// OrderNo applies equality check predicate on the "order_no" field. It's identical to OrderNoEQ.
func OrderNo(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderNo), v))
	})
}

// Disabled applies equality check predicate on the "disabled" field. It's identical to DisabledEQ.
func Disabled(v bool) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisabled), v))
	})
}

// Meta applies equality check predicate on the "meta" field. It's identical to MetaEQ.
func Meta(v types.MenuMeta) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMeta), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// MenuLevelEQ applies the EQ predicate on the "menu_level" field.
func MenuLevelEQ(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMenuLevel), v))
	})
}

// MenuLevelNEQ applies the NEQ predicate on the "menu_level" field.
func MenuLevelNEQ(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMenuLevel), v))
	})
}

// MenuLevelIn applies the In predicate on the "menu_level" field.
func MenuLevelIn(vs ...uint32) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMenuLevel), v...))
	})
}

// MenuLevelNotIn applies the NotIn predicate on the "menu_level" field.
func MenuLevelNotIn(vs ...uint32) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMenuLevel), v...))
	})
}

// MenuLevelGT applies the GT predicate on the "menu_level" field.
func MenuLevelGT(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMenuLevel), v))
	})
}

// MenuLevelGTE applies the GTE predicate on the "menu_level" field.
func MenuLevelGTE(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMenuLevel), v))
	})
}

// MenuLevelLT applies the LT predicate on the "menu_level" field.
func MenuLevelLT(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMenuLevel), v))
	})
}

// MenuLevelLTE applies the LTE predicate on the "menu_level" field.
func MenuLevelLTE(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMenuLevel), v))
	})
}

// MenuTypeEQ applies the EQ predicate on the "menu_type" field.
func MenuTypeEQ(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMenuType), v))
	})
}

// MenuTypeNEQ applies the NEQ predicate on the "menu_type" field.
func MenuTypeNEQ(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMenuType), v))
	})
}

// MenuTypeIn applies the In predicate on the "menu_type" field.
func MenuTypeIn(vs ...uint32) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMenuType), v...))
	})
}

// MenuTypeNotIn applies the NotIn predicate on the "menu_type" field.
func MenuTypeNotIn(vs ...uint32) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMenuType), v...))
	})
}

// MenuTypeGT applies the GT predicate on the "menu_type" field.
func MenuTypeGT(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMenuType), v))
	})
}

// MenuTypeGTE applies the GTE predicate on the "menu_type" field.
func MenuTypeGTE(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMenuType), v))
	})
}

// MenuTypeLT applies the LT predicate on the "menu_type" field.
func MenuTypeLT(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMenuType), v))
	})
}

// MenuTypeLTE applies the LTE predicate on the "menu_type" field.
func MenuTypeLTE(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMenuType), v))
	})
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v uint) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentID), v))
	})
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v uint) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParentID), v))
	})
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...uint) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldParentID), v...))
	})
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...uint) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldParentID), v...))
	})
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v uint) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldParentID), v))
	})
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v uint) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldParentID), v))
	})
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v uint) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldParentID), v))
	})
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v uint) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldParentID), v))
	})
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPath), v))
	})
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPath), v...))
	})
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPath), v...))
	})
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPath), v))
	})
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPath), v))
	})
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPath), v))
	})
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPath), v))
	})
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPath), v))
	})
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPath), v))
	})
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPath), v))
	})
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPath), v))
	})
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPath), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// RedirectEQ applies the EQ predicate on the "redirect" field.
func RedirectEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRedirect), v))
	})
}

// RedirectNEQ applies the NEQ predicate on the "redirect" field.
func RedirectNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRedirect), v))
	})
}

// RedirectIn applies the In predicate on the "redirect" field.
func RedirectIn(vs ...string) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRedirect), v...))
	})
}

// RedirectNotIn applies the NotIn predicate on the "redirect" field.
func RedirectNotIn(vs ...string) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRedirect), v...))
	})
}

// RedirectGT applies the GT predicate on the "redirect" field.
func RedirectGT(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRedirect), v))
	})
}

// RedirectGTE applies the GTE predicate on the "redirect" field.
func RedirectGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRedirect), v))
	})
}

// RedirectLT applies the LT predicate on the "redirect" field.
func RedirectLT(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRedirect), v))
	})
}

// RedirectLTE applies the LTE predicate on the "redirect" field.
func RedirectLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRedirect), v))
	})
}

// RedirectContains applies the Contains predicate on the "redirect" field.
func RedirectContains(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRedirect), v))
	})
}

// RedirectHasPrefix applies the HasPrefix predicate on the "redirect" field.
func RedirectHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRedirect), v))
	})
}

// RedirectHasSuffix applies the HasSuffix predicate on the "redirect" field.
func RedirectHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRedirect), v))
	})
}

// RedirectIsNil applies the IsNil predicate on the "redirect" field.
func RedirectIsNil() predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRedirect)))
	})
}

// RedirectNotNil applies the NotNil predicate on the "redirect" field.
func RedirectNotNil() predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRedirect)))
	})
}

// RedirectEqualFold applies the EqualFold predicate on the "redirect" field.
func RedirectEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRedirect), v))
	})
}

// RedirectContainsFold applies the ContainsFold predicate on the "redirect" field.
func RedirectContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRedirect), v))
	})
}

// ComponentEQ applies the EQ predicate on the "component" field.
func ComponentEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComponent), v))
	})
}

// ComponentNEQ applies the NEQ predicate on the "component" field.
func ComponentNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldComponent), v))
	})
}

// ComponentIn applies the In predicate on the "component" field.
func ComponentIn(vs ...string) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldComponent), v...))
	})
}

// ComponentNotIn applies the NotIn predicate on the "component" field.
func ComponentNotIn(vs ...string) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldComponent), v...))
	})
}

// ComponentGT applies the GT predicate on the "component" field.
func ComponentGT(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldComponent), v))
	})
}

// ComponentGTE applies the GTE predicate on the "component" field.
func ComponentGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldComponent), v))
	})
}

// ComponentLT applies the LT predicate on the "component" field.
func ComponentLT(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldComponent), v))
	})
}

// ComponentLTE applies the LTE predicate on the "component" field.
func ComponentLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldComponent), v))
	})
}

// ComponentContains applies the Contains predicate on the "component" field.
func ComponentContains(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldComponent), v))
	})
}

// ComponentHasPrefix applies the HasPrefix predicate on the "component" field.
func ComponentHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldComponent), v))
	})
}

// ComponentHasSuffix applies the HasSuffix predicate on the "component" field.
func ComponentHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldComponent), v))
	})
}

// ComponentEqualFold applies the EqualFold predicate on the "component" field.
func ComponentEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldComponent), v))
	})
}

// ComponentContainsFold applies the ContainsFold predicate on the "component" field.
func ComponentContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldComponent), v))
	})
}

// OrderNoEQ applies the EQ predicate on the "order_no" field.
func OrderNoEQ(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderNo), v))
	})
}

// OrderNoNEQ applies the NEQ predicate on the "order_no" field.
func OrderNoNEQ(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderNo), v))
	})
}

// OrderNoIn applies the In predicate on the "order_no" field.
func OrderNoIn(vs ...uint32) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderNo), v...))
	})
}

// OrderNoNotIn applies the NotIn predicate on the "order_no" field.
func OrderNoNotIn(vs ...uint32) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderNo), v...))
	})
}

// OrderNoGT applies the GT predicate on the "order_no" field.
func OrderNoGT(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderNo), v))
	})
}

// OrderNoGTE applies the GTE predicate on the "order_no" field.
func OrderNoGTE(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderNo), v))
	})
}

// OrderNoLT applies the LT predicate on the "order_no" field.
func OrderNoLT(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderNo), v))
	})
}

// OrderNoLTE applies the LTE predicate on the "order_no" field.
func OrderNoLTE(v uint32) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderNo), v))
	})
}

// DisabledEQ applies the EQ predicate on the "disabled" field.
func DisabledEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisabled), v))
	})
}

// DisabledNEQ applies the NEQ predicate on the "disabled" field.
func DisabledNEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisabled), v))
	})
}

// MetaEQ applies the EQ predicate on the "meta" field.
func MetaEQ(v types.MenuMeta) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMeta), v))
	})
}

// MetaNEQ applies the NEQ predicate on the "meta" field.
func MetaNEQ(v types.MenuMeta) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMeta), v))
	})
}

// MetaIn applies the In predicate on the "meta" field.
func MetaIn(vs ...types.MenuMeta) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMeta), v...))
	})
}

// MetaNotIn applies the NotIn predicate on the "meta" field.
func MetaNotIn(vs ...types.MenuMeta) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMeta), v...))
	})
}

// MetaGT applies the GT predicate on the "meta" field.
func MetaGT(v types.MenuMeta) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMeta), v))
	})
}

// MetaGTE applies the GTE predicate on the "meta" field.
func MetaGTE(v types.MenuMeta) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMeta), v))
	})
}

// MetaLT applies the LT predicate on the "meta" field.
func MetaLT(v types.MenuMeta) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMeta), v))
	})
}

// MetaLTE applies the LTE predicate on the "meta" field.
func MetaLTE(v types.MenuMeta) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMeta), v))
	})
}

// MetaIsNil applies the IsNil predicate on the "meta" field.
func MetaIsNil() predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMeta)))
	})
}

// MetaNotNil applies the NotNil predicate on the "meta" field.
func MetaNotNil() predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMeta)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysMenu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		p(s.Not())
	})
}
