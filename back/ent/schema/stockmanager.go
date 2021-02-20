package schema

import (
	"entgo.io/ent"
)

// StockManager holds the schema definition for the StockManager entity.
type StockManager struct {
	ent.Schema
}

// Fields of the StockManager.
func (StockManager) Fields() []ent.Field {
	return nil
}

// Edges of the StockManager.
func (StockManager) Edges() []ent.Edge {
	return nil
}
