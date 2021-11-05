// Code generated by entc, DO NOT EDIT.

package userprovider

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the userprovider type in the database.
	Label = "user_provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldProviderID holds the string denoting the provider_id field in the database.
	FieldProviderID = "provider_id"
	// FieldProviderUserID holds the string denoting the provider_user_id field in the database.
	FieldProviderUserID = "provider_user_id"
	// FieldUserProviderInfo holds the string denoting the user_provider_info field in the database.
	FieldUserProviderInfo = "user_provider_info"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the userprovider in the database.
	Table = "user_providers"
)

// Columns holds all SQL columns for userprovider fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldProviderID,
	FieldProviderUserID,
	FieldUserProviderInfo,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
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
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() int64
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() int64
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() int64
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
