package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").NotEmpty(),
		field.Strings("instrument").Optional(),
		field.Int("locationid"),
		field.Int("categoryid"),
		field.Float("cost"),
		field.Float("additional_cost").Default(0).Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("service").Unique().Required(),
	}
}
