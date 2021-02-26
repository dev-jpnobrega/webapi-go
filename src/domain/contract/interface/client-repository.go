package domain

import entity "webapi/src/domain/entity"

// IClientRepository - interface
type IClientRepository interface {
	Get(params interface{}) (error, []entity.Client)
}
