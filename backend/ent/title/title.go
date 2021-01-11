// Code generated by entc, DO NOT EDIT.

package title

const (
	// Label holds the string label denoting the title type in the database.
	Label = "title"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitlename holds the string denoting the titlename field in the database.
	FieldTitlename = "titlename"

	// EdgePersonal holds the string denoting the personal edge name in mutations.
	EdgePersonal = "personal"

	// Table holds the table name of the title in the database.
	Table = "titles"
	// PersonalTable is the table the holds the personal relation/edge.
	PersonalTable = "personals"
	// PersonalInverseTable is the table name for the Personal entity.
	// It exists in this package in order to avoid circular dependency with the "personal" package.
	PersonalInverseTable = "personals"
	// PersonalColumn is the table column denoting the personal relation/edge.
	PersonalColumn = "title_id"
)

// Columns holds all SQL columns for title fields.
var Columns = []string{
	FieldID,
	FieldTitlename,
}
