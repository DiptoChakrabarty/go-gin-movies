package validators

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

func ValidateMovie(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Potter")
}
