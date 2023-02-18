// Code generated by ent, DO NOT EDIT.

package screenshot

const (
	// Label holds the string label denoting the screenshot type in the database.
	Label = "screenshot"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldStoredPath holds the string denoting the stored_path field in the database.
	FieldStoredPath = "stored_path"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the screenshot in the database.
	Table = "screenshots"
)

// Columns holds all SQL columns for screenshot fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldURL,
	FieldStoredPath,
	FieldCreatedAt,
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