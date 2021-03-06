package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("username"),
		field.String("password"),
		field.String("salt"),
		field.String("display_name").Default(""),
		field.String("phone_number").Optional(),
		field.String("email_address").Optional(),
		field.UUID("app_id", uuid.UUID{}).
			Default(uuid.New),
		// Signup by user or import from other app
		field.String("signup_method"),
		field.String("avatar").Default(""),
		field.String("region").Default(""),
		field.Uint32("age").Default(0),
		field.String("gender").Default(""),
		field.String("birthday").Default(""),
		field.String("country").Default(""),
		field.String("province").Default(""),
		field.String("city").Default(""),
		field.String("career").Default(""),
		field.String("first_name").Default(""),
		field.String("last_name").Default(""),
		field.String("street_address1").Default(""),
		field.String("street_address2").Default(""),
		field.String("compony").Default(""),
		field.String("postal_code").Default(""),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "app_id"),
		index.Fields("email_address", "app_id"),
		index.Fields("phone_number", "app_id"),
	}
}
