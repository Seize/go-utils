package validutil

import (
	validator "gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
)

// ValidateID is used to help to validate ID in a struct
func ValidateID(fl validator.FieldLevel) bool {
	if !bson.IsObjectIdHex(fl.Field().String()) {
		return false
	}

	return bson.ObjectIdHex(fl.Field().String()).Valid()
}
