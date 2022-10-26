// Code generated by ent, DO NOT EDIT.

package systoken

import (
	"slash-admin/app/admin/ent/predicate"
	"slash-admin/pkg/types"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v types.Status) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// ExpiredAt applies equality check predicate on the "expired_at" field. It's identical to ExpiredAtEQ.
func ExpiredAt(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiredAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUUID), v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUUID), v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUUID), v))
	})
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUUID), v))
	})
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUUID), v))
	})
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUUID), v))
	})
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUUID), v))
	})
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToken), v))
	})
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldToken), v...))
	})
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldToken), v...))
	})
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToken), v))
	})
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToken), v))
	})
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToken), v))
	})
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToken), v))
	})
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldToken), v))
	})
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldToken), v))
	})
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldToken), v))
	})
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldToken), v))
	})
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldToken), v))
	})
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSource), v))
	})
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSource), v...))
	})
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSource), v...))
	})
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSource), v))
	})
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSource), v))
	})
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSource), v))
	})
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSource), v))
	})
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSource), v))
	})
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSource), v))
	})
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSource), v))
	})
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSource), v))
	})
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSource), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v types.Status) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v types.Status) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...types.Status) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...types.Status) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v types.Status) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v types.Status) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v types.Status) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v types.Status) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatus)))
	})
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatus)))
	})
}

// ExpiredAtEQ applies the EQ predicate on the "expired_at" field.
func ExpiredAtEQ(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiredAt), v))
	})
}

// ExpiredAtNEQ applies the NEQ predicate on the "expired_at" field.
func ExpiredAtNEQ(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiredAt), v))
	})
}

// ExpiredAtIn applies the In predicate on the "expired_at" field.
func ExpiredAtIn(vs ...time.Time) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExpiredAt), v...))
	})
}

// ExpiredAtNotIn applies the NotIn predicate on the "expired_at" field.
func ExpiredAtNotIn(vs ...time.Time) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExpiredAt), v...))
	})
}

// ExpiredAtGT applies the GT predicate on the "expired_at" field.
func ExpiredAtGT(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiredAt), v))
	})
}

// ExpiredAtGTE applies the GTE predicate on the "expired_at" field.
func ExpiredAtGTE(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiredAt), v))
	})
}

// ExpiredAtLT applies the LT predicate on the "expired_at" field.
func ExpiredAtLT(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiredAt), v))
	})
}

// ExpiredAtLTE applies the LTE predicate on the "expired_at" field.
func ExpiredAtLTE(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiredAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysToken) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysToken) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
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
func Not(p predicate.SysToken) predicate.SysToken {
	return predicate.SysToken(func(s *sql.Selector) {
		p(s.Not())
	})
}
