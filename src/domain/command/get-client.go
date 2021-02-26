package domain

import (
	interfaces "webapi/src/domain/contract/interface"
	values "webapi/src/domain/contract/value"
)

// GetClientCommand - fc
type GetClientCommand struct {
	Repository interfaces.IClientRepository
}

// GetModelValidate command
func (g *GetClientCommand) GetModelValidate() interface{} {
	return &values.InputLogin{}
}

func (g *GetClientCommand) Execute(
	input values.RequestData,
) (result values.ResponseData, err *values.ResponseError) {
	args := input.Args.(*values.InputLogin)

	errs, client := g.Repository.Get(args)

	if errs != nil {
		return
	}

	result.Data = client
	result.StatusCode = 200

	return
}
