package domain

import (
	interfaces "webapi/src/domain/contract/interface"
	values "webapi/src/domain/contract/value"
)

// A GetClientCommand represent business logic to search clients
type GetClientCommand struct {
	Repository interfaces.IClientRepository
}

func (c *GetClientCommand) GetModelValidate() *values.ValidateModal {
	return &values.ValidateModal{
		Modal: &values.SearchArgs{},
	}
}

func (g *GetClientCommand) Execute(input values.RequestData) (
	result values.ResponseData, err *values.ResponseError,
) {
	args := input.Args.(*values.SearchArgs)

	clients, err := g.Repository.Get(args)

	if err != nil {
		return
	}

	result.Data = clients
	result.StatusCode = 200

	return
}
