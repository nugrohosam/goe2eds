package validations

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// SHOULD_BE_INTEGER ..
const SHOULD_BE_INTEGER "should-be-integer";

// SHOULD_BE_STRING ..
const SHOULD_BE_STRING "should-be-string";

// ShouldBeType ..
type ShouldBeType struct {
	Key      string
	Function func(validator.FieldLevel) bool
}

// ValidateShouldBeInteger ...
var ValidateShouldBeInteger = ShouldBeType{
	Key:      SHOULD_BE_INTEGER,
	Function: ShouldBeIntegerFunc,
}

// ValidateShouldBeString ...
var ValidateShouldBeString = ShouldBeType{
	Key:      SHOULD_BE_STRING,
	Function: ShouldBeStringFunc,
}

// ShouldBeIntegerFunc ..
func ShouldBeIntegerFunc(field validator.FieldLevel) bool {

	if field.Field().Kind().String() != reflect.Float64.String() {
		return false
	}

	return true
}

// ShouldBeStringFunc ..
func ShouldBeStringFunc(field validator.FieldLevel) bool {

	if field.Field().Kind().String() != reflect.String.String() {
		return false
	}

	return true
}
