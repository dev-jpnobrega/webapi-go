package domain

// ResponseData ...
type ResponseData struct {
	Data       interface{}
	StatusCode int
}

// RequestData ...
type RequestData struct {
	Authorization string
	XAppToken     string
	UserInfo      interface{}
	Args          interface{}
}

// InputLogin - parameters login user
type InputLogin struct {
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
	Pass  string `json:"pass" form:"pass" query:"pass" validate:"required"`
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
