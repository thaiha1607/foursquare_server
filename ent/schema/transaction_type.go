package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// TransactionType holds the schema definition for the TransactionType entity.
type TransactionType struct {
	ent.Schema
}

// Annotations of the TransactionType.
func (TransactionType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "transaction_type"},
	}
}

// Fields of the TransactionType.
func (TransactionType) Fields() []ent.Field {
	return nil
}

// Edges of the TransactionType.
func (TransactionType) Edges() []ent.Edge {
	return nil
}

// Mixin of the TransactionType.
func (TransactionType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TextCodeMixin{},
	}
}
