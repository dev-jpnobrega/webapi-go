package infrastructure

import (
	infrastructure "webapi/src/db"
	command "webapi/src/domain/command"
	interfaces "webapi/src/domain/contract/interface"
	repository "webapi/src/infrastructure/repository"
)

func PutClientFactory() interfaces.ICommand {
	return &command.PutClientComannd{
		Repository: &repository.ClientRepository{
			Database: &infrastructure.DB{},
		},
	}
}
