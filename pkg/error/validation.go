package error

import (
	"github.com/go-playground/validator/v10"
)

// ErrValidation is a struct to represent error validation
type ErrValidation struct {
	Message string                 `json:"message"`
	Errors  []errValidationMessage `json:"errors"`
}

// errValidationMessage is a struct to represent error validation message
type errValidationMessage struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
}

// newErrValidationMessage is a function to compose error validation message
func newErrValidationMessage(err validator.ValidationErrors) (result []errValidationMessage) {
	for _, fieldErr := range err {
		result = append(result, errValidationMessage{
			Field: fieldErr.Field(),
			Tag:   fieldErr.Tag(),
			Value: fieldErr.Value(),
		})
	}
	return
}

// ResponseValidation is a function to response error validation
func newErrValidation(err validator.ValidationErrors) *ErrValidation {
	return &ErrValidation{
		Message: string(ErrValidationMsg),
		Errors:  newErrValidationMessage(err),
	}
}
