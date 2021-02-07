package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// MFile holds the schema definition for the MFile entity.
type MFile struct {
	ent.Schema
}

// Fields of the MFile.
func (MFile) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("parent").Unique(),
		field.String("name"),
		field.Int("author"),
		field.Int("md5").Unique(),
		field.Int("size").Optional().Default(0),
		field.String("desc").Optional().Default(""),
		field.Enum("status").Values("0", "10", "20").Default("0"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the MFile.
func (MFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("folder", Folder.Type).Ref("mfiles").Unique(),
	}
}
