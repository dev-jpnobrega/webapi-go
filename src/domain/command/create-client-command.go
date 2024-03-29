package domain

import (
	interfaces "webapi/src/domain/contract/interface"
	values "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"

	"github.com/google/uuid"
)

// A CreateClientComannd represent business logic to create client
type CreateClientComannd struct {
	Repository interfaces.IClientRepository
}

func (c *CreateClientComannd) GetModelValidate() *values.ValidateModal {
	return &values.ValidateModal{
		Modal: &entity.Client{},
	}
}

func (c *CreateClientComannd) Execute(input values.RequestData) (
	result values.ResponseData, err *values.ResponseError,
) {
	client := input.Args.(*entity.Client)
	client.ID = uuid.New()

	_, err = c.Repository.Create(client)

	if err != nil {
		return
	}

	result.Data = client
	result.StatusCode = 200

	return
}
