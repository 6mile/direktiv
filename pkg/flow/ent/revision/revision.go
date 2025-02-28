// Code generated by entc, DO NOT EDIT.

package revision

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the revision type in the database.
	Label = "revision"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// EdgeWorkflow holds the string denoting the workflow edge name in mutations.
	EdgeWorkflow = "workflow"
	// EdgeRefs holds the string denoting the refs edge name in mutations.
	EdgeRefs = "refs"
	// EdgeInstances holds the string denoting the instances edge name in mutations.
	EdgeInstances = "instances"
	// Table holds the table name of the revision in the database.
	Table = "revisions"
	// WorkflowTable is the table that holds the workflow relation/edge.
	WorkflowTable = "revisions"
	// WorkflowInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowInverseTable = "workflows"
	// WorkflowColumn is the table column denoting the workflow relation/edge.
	WorkflowColumn = "workflow_revisions"
	// RefsTable is the table that holds the refs relation/edge.
	RefsTable = "refs"
	// RefsInverseTable is the table name for the Ref entity.
	// It exists in this package in order to avoid circular dependency with the "ref" package.
	RefsInverseTable = "refs"
	// RefsColumn is the table column denoting the refs relation/edge.
	RefsColumn = "revision_refs"
	// InstancesTable is the table that holds the instances relation/edge.
	InstancesTable = "instances"
	// InstancesInverseTable is the table name for the Instance entity.
	// It exists in this package in order to avoid circular dependency with the "instance" package.
	InstancesInverseTable = "instances"
	// InstancesColumn is the table column denoting the instances relation/edge.
	InstancesColumn = "revision_instances"
)

// Columns holds all SQL columns for revision fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldHash,
	FieldSource,
	FieldMetadata,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "revisions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"workflow_revisions",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
