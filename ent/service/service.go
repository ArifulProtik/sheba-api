// Code generated by entc, DO NOT EDIT.

package service

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the service type in the database.
	Label = "service"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldInstrument holds the string denoting the instrument field in the database.
	FieldInstrument = "instrument"
	// FieldLocationid holds the string denoting the locationid field in the database.
	FieldLocationid = "locationid"
	// FieldCategoryid holds the string denoting the categoryid field in the database.
	FieldCategoryid = "categoryid"
	// FieldCost holds the string denoting the cost field in the database.
	FieldCost = "cost"
	// FieldAdditionalCost holds the string denoting the additional_cost field in the database.
	FieldAdditionalCost = "additional_cost"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the service in the database.
	Table = "services"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "services"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_service"
)

// Columns holds all SQL columns for service fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldInstrument,
	FieldLocationid,
	FieldCategoryid,
	FieldCost,
	FieldAdditionalCost,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "services"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_service",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultAdditionalCost holds the default value on creation for the "additional_cost" field.
	DefaultAdditionalCost float64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
