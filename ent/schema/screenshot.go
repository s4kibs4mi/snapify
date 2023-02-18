package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/snapify/models"
)

// Screenshot holds the schema definition for the Screenshot entity.
type Screenshot struct {
	ent.Schema
}

// Fields of the Screenshot.
func (Screenshot) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("status").GoType(models.Status("")),
		field.String("url"),
		field.String("stored_path").Nillable().Optional(),
		field.Time("created_at"),
	}
}

// Edges of the Screenshot.
func (Screenshot) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Screenshot) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("url"),
		index.Fields("created_at"),
	}
}
