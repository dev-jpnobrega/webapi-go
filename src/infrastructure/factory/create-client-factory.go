package infrastructure

import (
	infrastructure "webapi/src/db"
	command "webapi/src/domain/command"
	interfaces "webapi/src/domain/contract/interface"
	repository "webapi/src/infrastructure/repository"
)

func CreateClientFactory() interfaces.ICommand {
	return &command.CreateClientComannd{
		Repository: &repository.ClientRepository{
			Database: &infrastructure.DB{},
		},
	}
}
