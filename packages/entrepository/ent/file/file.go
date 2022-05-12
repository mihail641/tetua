// Code generated by entc, DO NOT EDIT.

package file

import (
	"time"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDisk holds the string denoting the disk field in the database.
	FieldDisk = "disk"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeUserAvatars holds the string denoting the user_avatars edge name in mutations.
	EdgeUserAvatars = "user_avatars"
	// Table holds the table name of the file in the database.
	Table = "files"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "files"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "featured_image_id"
	// UserAvatarsTable is the table that holds the user_avatars relation/edge.
	UserAvatarsTable = "users"
	// UserAvatarsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserAvatarsInverseTable = "users"
	// UserAvatarsColumn is the table column denoting the user_avatars relation/edge.
	UserAvatarsColumn = "avatar_image_id"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldDisk,
	FieldPath,
	FieldType,
	FieldSize,
	FieldUserID,
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
