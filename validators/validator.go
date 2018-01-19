package validators

import (
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

// Dbunique Validator
func Dbunique(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	return true
}
