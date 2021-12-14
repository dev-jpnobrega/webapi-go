package domain

import (
	entity "webapi/src/domain/entity"

	"github.com/cheekybits/genny/generic"
)

// ResponseData ...
type ResponseData struct {
	Data       interface{}
	StatusCode int
}

//ValidateModal -s
type ValidateModal struct {
	Modal generic.Type
}

// RequestData ...
type RequestData struct {
	Authorization string `header:"Authorization" form:"Authorization" query:"Authorization" validate:"required"`
	XAppToken     string `header:"X-App-Token" form:"X-App-Token" query:"X-App-Token" validate:"required"`
	UserInfo      entity.User
	Args          interface{}
}

// ErrorModel - ER
type ErrorModel struct {
	Code    int
	Message string
}

// ResponseError  FC
type ResponseError struct {
	StatusCode int
	Errors     []ErrorModel
}

// New - rF
func (re *ResponseError) New(message string, code int, statusCode int) *ResponseError {
	re = new(ResponseError)
	re.StatusCode = statusCode
	re.Errors = append(re.Errors, ErrorModel{
		Code:    code,
		Message: message,
	})

	return re
}

// Append - FC
func (re *ResponseError) Append(code int, message string) *ResponseError {
	re.Errors = append(re.Errors, ErrorModel{
		Code:    code,
		Message: message,
	})

	return re
}
