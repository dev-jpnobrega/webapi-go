package domain

import (
	values "webapi/src/domain/contract/value"
)

// ICommand - Representation command interface
type ICommand interface {
	GetModelValidate() *values.ValidateModal
	Execute(values.RequestData) (values.ResponseData, *values.ResponseError)
}
