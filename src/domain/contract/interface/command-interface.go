package domain

import values "webapi/src/domain/contract/value"

// ICommand - Representation command interface
type ICommand interface {
	GetModelValidate() interface{}
	Execute(values.RequestData) (values.ResponseData, *values.ResponseError)
}
