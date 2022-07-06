package domain

import (
	values "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"
)

// A IClientRepository represent repository client interface
type IClientRepository interface {
	Get(params interface{}) (*[]entity.Client, *values.ResponseError)
	Create(client *entity.Client) (*entity.Client, *values.ResponseError)
	Update(client *entity.Client) (bool, *values.ResponseError)
}
