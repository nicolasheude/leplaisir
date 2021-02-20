package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ContactParents holds the schema definition for the ContactParents entity.
type ContactParents struct {
	ent.Schema
}

// Fields of the ContactParents.
func (ContactParents) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Positive(),
		field.String("Nom"),
		field.String("Prenom"),
		field.String("Sexe"),
		field.String("Adresse"),
		field.String("Ville"),
		field.String("CP"),
		field.String("Email"),
		field.String("Tel1"),
		field.String("Tel2").Optional(),
	}
}

// Edges of the ContactParents.
func (ContactParents) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("child", Form.Type).
			Ref("contactParents").
			Required(),
	}
}
