package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// CommonAttribute is a mixin with dynamic ID field logic
type CommonAttribute struct {
	mixin.Schema
	Id string
}

func (ca CommonAttribute) Fields() []ent.Field {
	fields := []ent.Field{}

	switch ca.Id {
	case "id_uuid":
		fields = append(fields, field.String("id").
			StorageKey("id_uuid").
			DefaultFunc(func() string {
				return uuid.New().String()
			}).
			MaxLen(40).
			Unique().
			Immutable())
	case "id_int":
		fields = append(fields, field.Int("id").
			StorageKey("id_int").
			Unique().
			Positive().
			Immutable())
	default:
		panic("unsupported ID_TYPE: must be 'id_uuid' or 'id_int'")
	}

	// Add audit fields
	fields = append(fields,
		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.Time("deleted_at").
			Nillable().
			Optional(),

		field.String("created_by").
			MaxLen(100).
			Nillable().
			Optional(),

		field.String("updated_by").
			MaxLen(100).
			Nillable().
			Optional(),

		field.String("deleted_by").
			MaxLen(100).
			Nillable().
			Optional(),

		field.String("user_agent").
			Nillable().
			Optional(),

		field.String("ip_address").
			Nillable().
			Optional(),
	)

	return fields
}
