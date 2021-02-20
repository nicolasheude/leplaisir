// Code generated by entc, DO NOT EDIT.

package contactparents

const (
	// Label holds the string label denoting the contactparents type in the database.
	Label = "contact_parents"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNom holds the string denoting the nom field in the database.
	FieldNom = "nom"
	// FieldPrenom holds the string denoting the prenom field in the database.
	FieldPrenom = "prenom"
	// FieldSexe holds the string denoting the sexe field in the database.
	FieldSexe = "sexe"
	// FieldAdresse holds the string denoting the adresse field in the database.
	FieldAdresse = "adresse"
	// FieldVille holds the string denoting the ville field in the database.
	FieldVille = "ville"
	// FieldCP holds the string denoting the cp field in the database.
	FieldCP = "cp"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldTel1 holds the string denoting the tel1 field in the database.
	FieldTel1 = "tel1"
	// FieldTel2 holds the string denoting the tel2 field in the database.
	FieldTel2 = "tel2"

	// EdgeChild holds the string denoting the child edge name in mutations.
	EdgeChild = "child"

	// Table holds the table name of the contactparents in the database.
	Table = "contact_parents"
	// ChildTable is the table the holds the child relation/edge.
	ChildTable = "forms"
	// ChildInverseTable is the table name for the Form entity.
	// It exists in this package in order to avoid circular dependency with the "form" package.
	ChildInverseTable = "forms"
	// ChildColumn is the table column denoting the child relation/edge.
	ChildColumn = "form_contact_parents"
)

// Columns holds all SQL columns for contactparents fields.
var Columns = []string{
	FieldID,
	FieldNom,
	FieldPrenom,
	FieldSexe,
	FieldAdresse,
	FieldVille,
	FieldCP,
	FieldEmail,
	FieldTel1,
	FieldTel2,
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
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)
