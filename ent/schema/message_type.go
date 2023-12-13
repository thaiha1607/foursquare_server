package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// MessageType holds the schema definition for the MessageType entity.
type MessageType struct {
	ent.Schema
}

// Annotations of the MessageType.
func (MessageType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "message_type"},
	}
}

// Fields of the MessageType.
func (MessageType) Fields() []ent.Field {
	return nil
}

// Edges of the MessageType.
func (MessageType) Edges() []ent.Edge {
	return nil
}

// Mixin of the MessageType.
func (MessageType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TextCodeMixin{},
	}
}
