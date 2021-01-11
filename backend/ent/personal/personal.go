// Code generated by entc, DO NOT EDIT.

package personal

const (
	// Label holds the string label denoting the personal type in the database.
	Label = "personal"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgeTitle holds the string denoting the title edge name in mutations.
	EdgeTitle = "title"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "department"
	// EdgeGender holds the string denoting the gender edge name in mutations.
	EdgeGender = "gender"

	// Table holds the table name of the personal in the database.
	Table = "personals"
	// TitleTable is the table the holds the title relation/edge.
	TitleTable = "personals"
	// TitleInverseTable is the table name for the Title entity.
	// It exists in this package in order to avoid circular dependency with the "title" package.
	TitleInverseTable = "titles"
	// TitleColumn is the table column denoting the title relation/edge.
	TitleColumn = "title_id"
	// DepartmentTable is the table the holds the department relation/edge.
	DepartmentTable = "personals"
	// DepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentInverseTable = "departments"
	// DepartmentColumn is the table column denoting the department relation/edge.
	DepartmentColumn = "department_id"
	// GenderTable is the table the holds the gender relation/edge.
	GenderTable = "personals"
	// GenderInverseTable is the table name for the Gender entity.
	// It exists in this package in order to avoid circular dependency with the "gender" package.
	GenderInverseTable = "genders"
	// GenderColumn is the table column denoting the gender relation/edge.
	GenderColumn = "gender_id"
)

// Columns holds all SQL columns for personal fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Personal type.
var ForeignKeys = []string{
	"department_id",
	"gender_id",
	"title_id",
}
