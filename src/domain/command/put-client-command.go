package domain

import (
	interfaces "webapi/src/domain/contract/interface"
	values "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"
)

// A PutClientComannd represent business logic to edit client
type PutClientComannd struct {
	Repository interfaces.IClientRepository
}

func (c *PutClientComannd) GetModelValidate() *values.ValidateModal {
	return &values.ValidateModal{
		Modal: &values.PutArgs{},
	}
}

func (c *PutClientComannd) Execute(input values.RequestData) (
	result values.ResponseData, err *values.ResponseError,
) {
	client := input.Args.(*entity.Client)

	_, err = c.Repository.Update(client)

	if err != nil {
		return
	}

	result.Data = client
	result.StatusCode = 200

	return
}
