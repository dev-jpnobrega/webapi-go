package domain

import (
	values "webapi/src/domain/contract/value"
)

// A ICommand represent command interface
type ICommand interface {
	GetModelValidate() *values.ValidateModal
	Execute(values.RequestData) (values.ResponseData, *values.ResponseError)
}
