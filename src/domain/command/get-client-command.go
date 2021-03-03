package domain

import (
	interfaces "webapi/src/domain/contract/interface"
	values "webapi/src/domain/contract/value"
)

// GetClientCommand - Search clients by arguments
type GetClientCommand struct {
	Repository interfaces.IClientRepository
}

// GetModelValidate command
func (g *GetClientCommand) GetModelValidate() interface{} {
	return &values.SearchArgs{}
}

// Execute - Inicialize command operation
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
