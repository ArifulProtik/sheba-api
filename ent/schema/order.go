package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("serviceid", uuid.UUID{}),
		field.UUID("providerid", uuid.UUID{}),
		field.Float("totalcost"),
		field.Strings("address"),
		field.Bool("is_declined").Default(false),
		field.Bool("payment_ok").Default(false),
		field.Bool("is_accepted").Default(false),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("order").Required().
			Unique(),
	}
}
