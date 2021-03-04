package infrastructure

import (
	infrastructure "webapi/src/db"
	command "webapi/src/domain/command"
	interfaces "webapi/src/domain/contract/interface"
	repository "webapi/src/infrastructure/repository"
)

// GetClientFactory - factory
func GetClientFactory() interfaces.ICommand {
	return &command.GetClientCommand{
		Repository: &repository.ClientRepository{
			Database: &infrastructure.DB{},
		},
	}
}
