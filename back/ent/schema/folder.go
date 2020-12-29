package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// Folder holds the schema definition for the Folder entity.
type Folder struct {
	ent.Schema
}

// Fields of the Folder.
func (Folder) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("parent"),
		field.String("path"),
		field.String("name"),
		field.Int("author"),
		field.Int("size").Optional().Default(0),
		field.Enum("status").Values("0", "10", "20").Default("0"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Folder.
func (Folder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mfiles", MFile.Type),
		edge.To("c", Folder.Type).From("p").Unique(),
	}
}
