package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Form holds the schema definition for the Form entity.
type Form struct {
	ent.Schema
}

// Fields of the Form.
func (Form) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.String("Email"),
		field.String("Nom"),
		field.String("Prenom"),
		field.String("Anniversaire"),
		field.String("Nationalite"),
		field.String("Sexe"),
		field.String("Niveau"),
		field.String("FFT"),
		field.String("LocationR"),
		field.String("Repas"),
		field.String("Formule"),
		field.String("Semaine"),
	}
}

// Edges of the Form.
func (Form) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contactParents", ContactParents.Type).
			Unique(),
	}
}
