package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

// ValidateInputJSON validates the input JSON in the gin context and returns a list of
// ErrorMsg objects containing the field name and error message if the validation fails.
// If the validation succeeds, it returns nil.
// The function takes a pointer to the gin context and an error as parameters.
// The error parameter should be the error returned by the validator ValidationErrors.As()
// function. If the error is not a ValidationErrors, the function returns nil.
// The function is useful for returning JSON error responses to the client when the
// input JSON is invalid.
func ValidateInputJSON(c *gin.Context, err error) []ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}
		return out
	}
	return nil
}
