package domain_test

import (
	"testing"

	command "webapi/src/domain/command"
	valueObject "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"
	fixures "webapi/test/fixures"

	"github.com/stretchr/testify/assert"
)

func TestPutClientCommand(t *testing.T) {
	t.Run("Customer updated successfully", func(t *testing.T) {
		data := new(valueObject.RequestData)
		args := new(entity.Client)

		args.ID = clientsMock[0].ID
		args.Name = "JP-updated"

		commandExec := &command.PutClientComannd{
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

	t.Run("Customer updated with error", func(t *testing.T) {
		data := new(valueObject.RequestData)
		args := new(entity.Client)
		args.Name = "thorw"

		commandExec := &command.PutClientComannd{
			Repository: &fixures.ClientRepository{
				ClientsMock: clientsMock,
			},
		}

		data.Args = args

		_, err := commandExec.Execute(*data)

		assert.NotNil(t, err)
		assert.EqualValues(t, 400, err.StatusCode)
		assert.Len(t, err.Errors, 2)
	})

	t.Run("PutClient validate model", func(t *testing.T) {
		data := new(valueObject.RequestData)
		args := new(valueObject.PutArgs)

		commandExec := &command.PutClientComannd{
			Repository: &fixures.ClientRepository{},
		}

		data.Args = args

		validated := commandExec.GetModelValidate()

		assert.Equal(t, args, validated.Modal)
	})
}
