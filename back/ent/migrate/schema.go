// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContactParentsColumns holds the columns for the "contact_parents" table.
	ContactParentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "nom", Type: field.TypeString},
		{Name: "prenom", Type: field.TypeString},
		{Name: "sexe", Type: field.TypeString},
		{Name: "adresse", Type: field.TypeString},
		{Name: "ville", Type: field.TypeString},
		{Name: "cp", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "tel1", Type: field.TypeString},
		{Name: "tel2", Type: field.TypeString, Nullable: true},
	}
	// ContactParentsTable holds the schema information for the "contact_parents" table.
	ContactParentsTable = &schema.Table{
		Name:        "contact_parents",
		Columns:     ContactParentsColumns,
		PrimaryKey:  []*schema.Column{ContactParentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FormsColumns holds the columns for the "forms" table.
	FormsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString},
		{Name: "nom", Type: field.TypeString},
		{Name: "prenom", Type: field.TypeString},
		{Name: "anniversaire", Type: field.TypeString},
		{Name: "nationalite", Type: field.TypeString},
		{Name: "sexe", Type: field.TypeString},
		{Name: "niveau", Type: field.TypeString},
		{Name: "fft", Type: field.TypeString},
		{Name: "locationr", Type: field.TypeString},
		{Name: "repas", Type: field.TypeString},
		{Name: "formule", Type: field.TypeString},
		{Name: "semaine", Type: field.TypeString},
		{Name: "form_contact_parents", Type: field.TypeInt, Nullable: true},
	}
	// FormsTable holds the schema information for the "forms" table.
	FormsTable = &schema.Table{
		Name:       "forms",
		Columns:    FormsColumns,
		PrimaryKey: []*schema.Column{FormsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "forms_contact_parents_contactParents",
				Columns: []*schema.Column{FormsColumns[13]},

				RefColumns: []*schema.Column{ContactParentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StockManagersColumns holds the columns for the "stock_managers" table.
	StockManagersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// StockManagersTable holds the schema information for the "stock_managers" table.
	StockManagersTable = &schema.Table{
		Name:        "stock_managers",
		Columns:     StockManagersColumns,
		PrimaryKey:  []*schema.Column{StockManagersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContactParentsTable,
		FormsTable,
		StockManagersTable,
	}
)

func init() {
	FormsTable.ForeignKeys[0].RefTable = ContactParentsTable
}
