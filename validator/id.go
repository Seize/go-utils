package validator

import (
	"encoding/hex"

	validator "gopkg.in/go-playground/validator.v9"
)

// ValidateID is used to help to validate ID in a struct
func ValidateID(fl validator.FieldLevel) bool {
	return isObjectIDHex(fl.Field().String())
}

// isObjectIDHex returns whether s is a valid hex representation of
// an ObjectId. See the ObjectIdHex function.
// @see https://godoc.org/labix.org/v2/mgo/bson#IsObjectIdHex
func isObjectIDHex(s string) bool {
	if len(s) != 24 {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil
}
