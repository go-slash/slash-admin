// Code generated by ent, DO NOT EDIT.

package systoken

import (
	"time"
)

const (
	// Label holds the string label denoting the systoken type in the database.
	Label = "sys_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldExpiredAt holds the string denoting the expired_at field in the database.
	FieldExpiredAt = "expired_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the systoken in the database.
	Table = "sys_token"
)

// Columns holds all SQL columns for systoken fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldToken,
	FieldSource,
	FieldStatus,
	FieldExpiredAt,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)