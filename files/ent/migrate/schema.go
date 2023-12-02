// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "content_type", Type: field.TypeString},
		{Name: "download_count", Type: field.TypeInt, Nullable: true, Default: 0},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FilesTable,
	}
)

func init() {
}