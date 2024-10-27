package validator

import (
	"GoAuth/src/pkg/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Validate validates the given struct and returns a custom ValidationError if there are validation errors.
func Validate(req interface{}) map[string]any {
	err := validator.New(validator.WithRequiredStructEnabled()).Struct(req)
	if err != nil {
		errorMap := make(map[string]any)
		for _, err := range err.(validator.ValidationErrors) {
			errorMap[utils.CamelToSnake(err.Field())] = fmt.Sprintf("Validation failed on %s: '%s'", utils.CamelToSnake(err.Field()), err.Tag())
		}

		return errorMap
	}

	return nil
}
