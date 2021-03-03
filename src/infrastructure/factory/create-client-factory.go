package infrastructure

import (
	infrastructure "webapi/src/db"
	command "webapi/src/domain/command"
	interfaces "webapi/src/domain/contract/interface"
	repository "webapi/src/infrastructure/repository"
)

// CreateClientFactory - factory
func CreateClientFactory() interfaces.ICommand {
	return &command.CreateClientComannd{
		Repository: &repository.ClientRepository{&infrastructure.DB{}},
	}
}
