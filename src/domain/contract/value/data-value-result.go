package domain

import (
	entity "webapi/src/domain/entity"

	"github.com/cheekybits/genny/generic"
)

type ResponseData struct {
	Data       interface{}
	StatusCode int
}

type ValidateModal struct {
	Modal generic.Type
}

type RequestData struct {
	Authorization string `header:"Authorization" form:"Authorization" query:"Authorization" validate:"required"`
	XAppToken     string `header:"X-App-Token" form:"X-App-Token" query:"X-App-Token" validate:"required"`
	UserInfo      entity.User
	Args          interface{}
}

type ErrorModel struct {
	Code    int
	Message string
}

type ResponseError struct {
	StatusCode int
	Errors     []ErrorModel
}

func (re *ResponseError) New(message string, code int, statusCode int) *ResponseError {
	re = new(ResponseError)
	re.StatusCode = statusCode
	re.Errors = append(re.Errors, ErrorModel{
		Code:    code,
		Message: message,
	})

	return re
}

func (re *ResponseError) Append(code int, message string) *ResponseError {
	re.Errors = append(re.Errors, ErrorModel{
		Code:    code,
		Message: message,
	})

	return re
}
