// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString},
		{Name: "display_name", Type: field.TypeString, Default: ""},
		{Name: "phone_number", Type: field.TypeString, Nullable: true},
		{Name: "email_address", Type: field.TypeString, Nullable: true},
		{Name: "login_times", Type: field.TypeInt32, Default: 0},
		{Name: "kyc_verify", Type: field.TypeBool, Default: false},
		{Name: "ga_verify", Type: field.TypeBool, Default: false},
		{Name: "signup_method", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
		{Name: "avatar", Type: field.TypeString, Default: ""},
		{Name: "region", Type: field.TypeString, Default: ""},
		{Name: "age", Type: field.TypeInt32, Default: 0},
		{Name: "gender", Type: field.TypeString, Default: ""},
		{Name: "birthday", Type: field.TypeString, Default: ""},
		{Name: "country", Type: field.TypeString, Default: ""},
		{Name: "province", Type: field.TypeString, Default: ""},
		{Name: "city", Type: field.TypeString, Default: ""},
		{Name: "career", Type: field.TypeString, Default: ""},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserFrozensColumns holds the columns for the "user_frozens" table.
	UserFrozensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "frozen_by", Type: field.TypeUUID},
		{Name: "frozen_cause", Type: field.TypeString},
		{Name: "start_at", Type: field.TypeInt64},
		{Name: "end_at", Type: field.TypeInt64, Default: 0},
		{Name: "status", Type: field.TypeString},
		{Name: "unfrozen_by", Type: field.TypeUUID},
	}
	// UserFrozensTable holds the schema information for the "user_frozens" table.
	UserFrozensTable = &schema.Table{
		Name:       "user_frozens",
		Columns:    UserFrozensColumns,
		PrimaryKey: []*schema.Column{UserFrozensColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userfrozen_id_status",
				Unique:  true,
				Columns: []*schema.Column{UserFrozensColumns[0], UserFrozensColumns[6]},
			},
		},
	}
	// UserProvidersColumns holds the columns for the "user_providers" table.
	UserProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "provider_id", Type: field.TypeUUID},
		{Name: "provider_user_id", Type: field.TypeString, Unique: true},
		{Name: "user_provider_info", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// UserProvidersTable holds the schema information for the "user_providers" table.
	UserProvidersTable = &schema.Table{
		Name:       "user_providers",
		Columns:    UserProvidersColumns,
		PrimaryKey: []*schema.Column{UserProvidersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userprovider_user_id_provider_id",
				Unique:  true,
				Columns: []*schema.Column{UserProvidersColumns[1], UserProvidersColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
		UserFrozensTable,
		UserProvidersTable,
	}
)

func init() {
}
