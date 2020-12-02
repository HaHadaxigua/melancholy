package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"time"
)

// ExitLog holds the schema definition for the ExitLog entity.
type ExitLog struct {
	ent.Schema
}

// Fields of the ExitLog.
func (ExitLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("user_id"),
		field.String("token").Unique(),
		field.Time("date").Default(time.Now).Immutable(),
	}
}

// Edges of the ExitLog.
func (ExitLog) Edges() []ent.Edge {
	return nil
}
