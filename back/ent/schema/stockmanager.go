package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// StockManager holds the schema definition for the StockManager entity.
type StockManager struct {
	ent.Schema
}

// Fields of the StockManager.
func (StockManager) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.String("Activite"),
		field.Int("SemaineA").Min(-1),
		field.Int("SemaineB").Min(-1),
		field.Int("SemaineC").Min(-1),
		field.Int("SemaineD").Min(-1),
		field.Int("SemaineE").Min(-1),
		field.Int("SemaineF").Min(-1),
		field.Int("SemaineG").Min(-1),
		field.Int("SemaineH").Min(-1),
	}
}

// Edges of the StockManager.
func (StockManager) Edges() []ent.Edge {
	return nil
}
