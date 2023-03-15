package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	command "webapi/src/domain/command"
	valueObject "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"

	fixures "webapi/test/fixures"
)

func TestCreateClientCommand(t *testing.T) {
	t.Run("Customer created successfully", func(t *testing.T) {
		data := new(valueObject.RequestData)
		client := new(entity.Client)

		client.CompanyID = 1
		client.Email = "test@test.com"
		client.Name = "JP"

		commandExec := &command.CreateClientComannd{
			Repository: &fixures.ClientRepository{},
		}

		data.Args = client

		result, err := commandExec.Execute(*data)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.EqualValues(t, 200, result.StatusCode)
	})

	t.Run("Client created with error", func(t *testing.T) {
		data := new(valueObject.RequestData)
		client := new(entity.Client)

		client.Age = 100
		client.Email = "test@test.com"
		client.Name = "JP"

		commandExec := &command.CreateClientComannd{
			Repository: &fixures.ClientRepository{},
		}

		data.Args = client

		_, err := commandExec.Execute(*data)

		assert.NotNil(t, err)
		assert.EqualValues(t, 400, err.StatusCode)
		assert.Len(t, err.Errors, 1)
	})

	t.Run("Client validate model", func(t *testing.T) {
		data := new(valueObject.RequestData)
		client := new(entity.Client)

		commandExec := &command.CreateClientComannd{
			Repository: &fixures.ClientRepository{},
		}

		data.Args = client

		validated := commandExec.GetModelValidate()

		assert.Equal(t, client, validated.Modal)
	})
}
