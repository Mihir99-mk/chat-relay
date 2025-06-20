package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AuthProvider holds the schema definition for the AuthProvider entity.
type AuthProvider struct {
	ent.Schema
}

func (AuthProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonAttribute{
			Id: "id_int",
		},
	}
}

func (AuthProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_auth_providers"},
	}
}

// Fields of the AuthProvider.
func (AuthProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StorageKey("name").MaxLen(64).NotEmpty(),
		field.String("display_name").StorageKey("display_name").Nillable(),
	}
}

// Edges of the AuthProvider.
func (AuthProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", AuthUser.Type),
	}
}
