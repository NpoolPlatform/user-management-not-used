// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldEmailAddress holds the string denoting the email_address field in the database.
	FieldEmailAddress = "email_address"
	// FieldSignupMethod holds the string denoting the signup_method field in the database.
	FieldSignupMethod = "signup_method"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldRegion holds the string denoting the region field in the database.
	FieldRegion = "region"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "province"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldCareer holds the string denoting the career field in the database.
	FieldCareer = "career"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldStreetAddress1 holds the string denoting the street_address1 field in the database.
	FieldStreetAddress1 = "street_address1"
	// FieldStreetAddress2 holds the string denoting the street_address2 field in the database.
	FieldStreetAddress2 = "street_address2"
	// FieldCompony holds the string denoting the compony field in the database.
	FieldCompony = "compony"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldSalt,
	FieldDisplayName,
	FieldPhoneNumber,
	FieldEmailAddress,
	FieldSignupMethod,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
	FieldAvatar,
	FieldRegion,
	FieldAge,
	FieldGender,
	FieldBirthday,
	FieldCountry,
	FieldProvince,
	FieldCity,
	FieldCareer,
	FieldFirstName,
	FieldLastName,
	FieldStreetAddress1,
	FieldStreetAddress2,
	FieldCompony,
	FieldPostalCode,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDisplayName holds the default value on creation for the "display_name" field.
	DefaultDisplayName string
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultAvatar holds the default value on creation for the "avatar" field.
	DefaultAvatar string
	// DefaultRegion holds the default value on creation for the "region" field.
	DefaultRegion string
	// DefaultAge holds the default value on creation for the "age" field.
	DefaultAge uint32
	// DefaultGender holds the default value on creation for the "gender" field.
	DefaultGender string
	// DefaultBirthday holds the default value on creation for the "birthday" field.
	DefaultBirthday string
	// DefaultCountry holds the default value on creation for the "country" field.
	DefaultCountry string
	// DefaultProvince holds the default value on creation for the "province" field.
	DefaultProvince string
	// DefaultCity holds the default value on creation for the "city" field.
	DefaultCity string
	// DefaultCareer holds the default value on creation for the "career" field.
	DefaultCareer string
	// DefaultFirstName holds the default value on creation for the "first_name" field.
	DefaultFirstName string
	// DefaultLastName holds the default value on creation for the "last_name" field.
	DefaultLastName string
	// DefaultStreetAddress1 holds the default value on creation for the "street_address1" field.
	DefaultStreetAddress1 string
	// DefaultStreetAddress2 holds the default value on creation for the "street_address2" field.
	DefaultStreetAddress2 string
	// DefaultCompony holds the default value on creation for the "compony" field.
	DefaultCompony string
	// DefaultPostalCode holds the default value on creation for the "postal_code" field.
	DefaultPostalCode string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
