package domain

import (
	values "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"
)

// IClientRepository - interface
type IClientRepository interface {
	Get(params interface{}) (*[]entity.Client, *values.ResponseError)
}
