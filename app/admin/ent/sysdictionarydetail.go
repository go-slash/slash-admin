// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"slash-admin/app/admin/ent/sysdictionarydetail"
	"slash-admin/pkg/types"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// SysDictionaryDetail is the model entity for the SysDictionaryDetail schema.
type SysDictionaryDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// the title shown in the UI
	Title string `json:"title,omitempty"`
	// key
	Key string `json:"key,omitempty"`
	// value
	Value string `json:"value,omitempty"`
	// 0=禁用 1=开启
	Status types.Status `json:"status,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 排序编号
	OrderNo uint32 `json:"order_no,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysDictionaryDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysdictionarydetail.FieldID, sysdictionarydetail.FieldOrderNo:
			values[i] = new(sql.NullInt64)
		case sysdictionarydetail.FieldTitle, sysdictionarydetail.FieldKey, sysdictionarydetail.FieldValue, sysdictionarydetail.FieldRemark:
			values[i] = new(sql.NullString)
		case sysdictionarydetail.FieldCreatedAt, sysdictionarydetail.FieldUpdatedAt, sysdictionarydetail.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case sysdictionarydetail.FieldStatus:
			values[i] = new(types.Status)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysDictionaryDetail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysDictionaryDetail fields.
func (sdd *SysDictionaryDetail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysdictionarydetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sdd.ID = uint64(value.Int64)
		case sysdictionarydetail.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				sdd.Title = value.String
			}
		case sysdictionarydetail.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				sdd.Key = value.String
			}
		case sysdictionarydetail.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				sdd.Value = value.String
			}
		case sysdictionarydetail.FieldStatus:
			if value, ok := values[i].(*types.Status); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil {
				sdd.Status = *value
			}
		case sysdictionarydetail.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				sdd.Remark = value.String
			}
		case sysdictionarydetail.FieldOrderNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_no", values[i])
			} else if value.Valid {
				sdd.OrderNo = uint32(value.Int64)
			}
		case sysdictionarydetail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sdd.CreatedAt = value.Time
			}
		case sysdictionarydetail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sdd.UpdatedAt = value.Time
			}
		case sysdictionarydetail.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sdd.DeletedAt = new(time.Time)
				*sdd.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysDictionaryDetail.
// Note that you need to call SysDictionaryDetail.Unwrap() before calling this method if this SysDictionaryDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (sdd *SysDictionaryDetail) Update() *SysDictionaryDetailUpdateOne {
	return (&SysDictionaryDetailClient{config: sdd.config}).UpdateOne(sdd)
}

// Unwrap unwraps the SysDictionaryDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sdd *SysDictionaryDetail) Unwrap() *SysDictionaryDetail {
	_tx, ok := sdd.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysDictionaryDetail is not a transactional entity")
	}
	sdd.config.driver = _tx.drv
	return sdd
}

// String implements the fmt.Stringer.
func (sdd *SysDictionaryDetail) String() string {
	var builder strings.Builder
	builder.WriteString("SysDictionaryDetail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sdd.ID))
	builder.WriteString("title=")
	builder.WriteString(sdd.Title)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(sdd.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(sdd.Value)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", sdd.Status))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(sdd.Remark)
	builder.WriteString(", ")
	builder.WriteString("order_no=")
	builder.WriteString(fmt.Sprintf("%v", sdd.OrderNo))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sdd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sdd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := sdd.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// SysDictionaryDetails is a parsable slice of SysDictionaryDetail.
type SysDictionaryDetails []*SysDictionaryDetail

func (sdd SysDictionaryDetails) config(cfg config) {
	for _i := range sdd {
		sdd[_i].config = cfg
	}
}
