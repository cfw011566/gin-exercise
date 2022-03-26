package validators

import (
	"strings"

	"github.com/go-playground/validator"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	if strings.Contains(field.Field().String(), "Cool") {
		return true
	}
	if strings.Contains(field.Field().String(), "Title") {
		return true
	}
	return false
}
