package helper

import "github.com/go-playground/validator/v10"

// Response struct
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta struct
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Success bool   `json:"success"`
}

// APIResponse function
func APIResponse(message string, code int, success bool, data interface{}) Response {
	metaData := Meta{
		Message: message,
		Code:    code,
		Success: success,
	}

	jsonResponse := Response{
		Meta: metaData,
		Data: data,
	}

	return jsonResponse
}

// FormatValidationError function
func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
