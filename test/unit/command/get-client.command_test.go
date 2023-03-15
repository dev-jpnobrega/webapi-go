package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	command "webapi/src/domain/command"
	valueObject "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"

	fixures "webapi/test/fixures"
)

var clientsMock []entity.Client

func TestMain(m *testing.M) {
	client1 := entity.Client{ID: uuid.New(), Name: "JP", Age: 10}
	client2 := entity.Client{ID: uuid.New(), Name: "JP", Age: 10}
	client3 := entity.Client{ID: uuid.New(), Name: "JP", Age: 10}

	clientsMock = append(clientsMock, client1, client2, client3)

	m.Run()
}

func TestGetClientCommand(t *testing.T) {

	t.Run("Get clients without filters", func(t *testing.T) {
		data := new(valueObject.RequestData)
		args := new(valueObject.SearchArgs)

		commandExec := &command.GetClientCommand{
			Repository: &fixures.ClientRepository{
				ClientsMock: clientsMock,
			},
		}

		data.Args = args

		result, err := commandExec.Execute(*data)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.EqualValues(t, 200, result.StatusCode)
	})

	t.Run("Get clients with error", func(t *testing.T) {
		data := new(valueObject.RequestData)
		args := new(valueObject.SearchArgs)
		args.Name = "thorw"

		commandExec := &command.GetClientCommand{
			Repository: &fixures.ClientRepository{},
		}

		data.Args = args

		_, err := commandExec.Execute(*data)

		assert.NotNil(t, err)
		assert.EqualValues(t, 400, err.StatusCode)
		assert.Len(t, err.Errors, 1)
	})

	t.Run("GetClient validate model", func(t *testing.T) {
		data := new(valueObject.RequestData)
		args := new(valueObject.SearchArgs)

		commandExec := &command.GetClientCommand{
			Repository: &fixures.ClientRepository{},
		}

		data.Args = args

		validated := commandExec.GetModelValidate()

		assert.Equal(t, args, validated.Modal)
	})
}
