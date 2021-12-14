package domain

import (
	interfaces "webapi/src/domain/contract/interface"
	values "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"
)

// PutClientComannd - PUT client command
type PutClientComannd struct {
	Repository interfaces.IClientRepository
}

// GetModelValidate command
func (c *PutClientComannd) GetModelValidate() *values.ValidateModal {
	return &values.ValidateModal{
		Modal: &values.PutArgs{},
	}
}

// Execute - Create client execute
func (c *PutClientComannd) Execute(input values.RequestData) (
	result values.ResponseData, err *values.ResponseError,
) {
	client := input.Args.(*entity.Client)
	// client.ID = uuid.New()

	_, err = c.Repository.Update(client)

	if err != nil {
		return
	}

	result.Data = client
	result.StatusCode = 200

	return
}
