package customvalidator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

// Custom error response
type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Validate method for echo's interface
func (cv *CustomValidator) Validate(i interface{}) []ErrorResponse {

	err := cv.Validator.Struct(i)

	if err != nil {
		return cv.validationErrorHandler(err)
	}

	return nil
}

func (cv *CustomValidator) HumanizeMessage(errResp []ErrorResponse) string {
	var messages []string
	for _, v := range errResp {
		message := fmt.Sprintf("%s: %s", v.Field, v.Message)
		messages = append(messages, message)
	}
	humanized := strings.Join(messages, ", ")
	return humanized
}

func (cv *CustomValidator) validationErrorHandler(err error) []ErrorResponse {
	var errors []ErrorResponse

	validationErrors := err.(validator.ValidationErrors)

	for _, e := range validationErrors {
		error := ErrorResponse{
			Field:   e.Field(),
			Message: cv.getErrorMsg(e),
		}
		errors = append(errors, error)
	}

	return errors
}

func (cv *CustomValidator) getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Minimum length is %s", fe.Param())
	case "max":
		return fmt.Sprintf("Maximum length is %s", fe.Param())
	default:
		return "Invalid value"
	}
}

func NewCustomValidator() *CustomValidator {
	v := validator.New(validator.WithRequiredStructEnabled())

	return &CustomValidator{
		Validator: v,
	}
}
