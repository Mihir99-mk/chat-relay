package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AuthUser holds the schema definition for the AuthUser entity.
type AuthUser struct {
	ent.Schema
	CommonAttribute
}

func (AuthUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonAttribute{
			Id: "id_uuid",
		},
	}
}

func (AuthUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_auth_users"},
	}
}

// Fields of the AuthUser.
func (AuthUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("provider_id"), // e.g., "google", "slack", "github"

		field.String("provider_user_id"). // Slack user ID, Google sub, etc.
							MaxLen(255).
							NotEmpty(),

		field.String("team_id"). // Optional, Slack workspace/team
						MaxLen(255).
						Nillable().
						Optional(),

		field.String("name").
			MaxLen(255).
			Nillable().
			Optional(),

		field.String("real_name").
			MaxLen(255).
			Nillable().
			Optional(),

		field.String("email").
			MaxLen(255).
			Nillable().
			Optional(),

		field.Text("access_token").
			NotEmpty(),

		field.Text("refresh_token").
			Nillable().
			Optional(),

		field.String("token_type").
			MaxLen(64).
			Nillable().
			Optional(),

		field.Time("expires_at").
			Nillable().
			Optional(),

		field.Text("scope").
			Nillable().
			Optional(),

		field.JSON("raw_profile", map[string]interface{}{}).
			Optional(),
	}
}

// Edges of the AuthUser.
func (AuthUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider", AuthProvider.Type).
			Ref("users").
			Field("provider_id").
			Required().
			Unique().
			Comment("Foreign key to auth_providers(id)"),
	}
}
