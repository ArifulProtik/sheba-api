package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("Location"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return nil
}
