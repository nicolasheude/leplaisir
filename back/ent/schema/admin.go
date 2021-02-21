package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.String("Identifiant"),
		field.String("MotDePasse"),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return nil
}
