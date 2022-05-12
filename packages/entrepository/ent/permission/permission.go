// Code generated by entc, DO NOT EDIT.

package permission

import (
	"time"
)

const (
	// Label holds the string label denoting the permission type in the database.
	Label = "permission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// Table holds the table name of the permission in the database.
	Table = "permissions"
	// RoleTable is the table that holds the role relation/edge.
	RoleTable = "permissions"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "roles"
	// RoleColumn is the table column denoting the role relation/edge.
	RoleColumn = "role_id"
)

// Columns holds all SQL columns for permission fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldRoleID,
	FieldAction,
	FieldValue,
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
