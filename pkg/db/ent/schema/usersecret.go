package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// UserSecret holds the schema definition for the UserSecret entity.
type UserSecret struct {
	ent.Schema
}

// Fields of the UserSecret.
func (UserSecret) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("app_id", uuid.UUID{}),
		field.String("secret"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the UserSecret.
func (UserSecret) Edges() []ent.Edge {
	return nil
}

func (UserSecret) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "app_id"),
	}
}
