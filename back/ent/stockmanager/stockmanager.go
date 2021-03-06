// Code generated by entc, DO NOT EDIT.

package stockmanager

const (
	// Label holds the string label denoting the stockmanager type in the database.
	Label = "stock_manager"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActivite holds the string denoting the activite field in the database.
	FieldActivite = "activite"
	// FieldSemaineA holds the string denoting the semainea field in the database.
	FieldSemaineA = "semainea"
	// FieldSemaineB holds the string denoting the semaineb field in the database.
	FieldSemaineB = "semaineb"
	// FieldSemaineC holds the string denoting the semainec field in the database.
	FieldSemaineC = "semainec"
	// FieldSemaineD holds the string denoting the semained field in the database.
	FieldSemaineD = "semained"
	// FieldSemaineE holds the string denoting the semainee field in the database.
	FieldSemaineE = "semainee"
	// FieldSemaineF holds the string denoting the semainef field in the database.
	FieldSemaineF = "semainef"
	// FieldSemaineG holds the string denoting the semaineg field in the database.
	FieldSemaineG = "semaineg"
	// FieldSemaineH holds the string denoting the semaineh field in the database.
	FieldSemaineH = "semaineh"

	// Table holds the table name of the stockmanager in the database.
	Table = "stock_managers"
)

// Columns holds all SQL columns for stockmanager fields.
var Columns = []string{
	FieldID,
	FieldActivite,
	FieldSemaineA,
	FieldSemaineB,
	FieldSemaineC,
	FieldSemaineD,
	FieldSemaineE,
	FieldSemaineF,
	FieldSemaineG,
	FieldSemaineH,
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
	// SemaineAValidator is a validator for the "SemaineA" field. It is called by the builders before save.
	SemaineAValidator func(int) error
	// SemaineBValidator is a validator for the "SemaineB" field. It is called by the builders before save.
	SemaineBValidator func(int) error
	// SemaineCValidator is a validator for the "SemaineC" field. It is called by the builders before save.
	SemaineCValidator func(int) error
	// SemaineDValidator is a validator for the "SemaineD" field. It is called by the builders before save.
	SemaineDValidator func(int) error
	// SemaineEValidator is a validator for the "SemaineE" field. It is called by the builders before save.
	SemaineEValidator func(int) error
	// SemaineFValidator is a validator for the "SemaineF" field. It is called by the builders before save.
	SemaineFValidator func(int) error
	// SemaineGValidator is a validator for the "SemaineG" field. It is called by the builders before save.
	SemaineGValidator func(int) error
	// SemaineHValidator is a validator for the "SemaineH" field. It is called by the builders before save.
	SemaineHValidator func(int) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)
