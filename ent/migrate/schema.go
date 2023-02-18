// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ScreenshotsColumns holds the columns for the "screenshots" table.
	ScreenshotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "status", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "stored_path", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ScreenshotsTable holds the schema information for the "screenshots" table.
	ScreenshotsTable = &schema.Table{
		Name:       "screenshots",
		Columns:    ScreenshotsColumns,
		PrimaryKey: []*schema.Column{ScreenshotsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "screenshot_url",
				Unique:  false,
				Columns: []*schema.Column{ScreenshotsColumns[2]},
			},
			{
				Name:    "screenshot_created_at",
				Unique:  false,
				Columns: []*schema.Column{ScreenshotsColumns[4]},
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "token", Type: field.TypeString},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "token_token",
				Unique:  true,
				Columns: []*schema.Column{TokensColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ScreenshotsTable,
		TokensTable,
	}
)

func init() {
}