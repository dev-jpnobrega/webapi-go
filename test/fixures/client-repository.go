package domain_test

import (
	values "webapi/src/domain/contract/value"

	entity "webapi/src/domain/entity"
)

// A ClientRepository represent client database
type ClientRepository struct {
	ClientsMock []entity.Client
}

func (c *ClientRepository) Get(params interface{}) (clientsResult *[]entity.Client, err *values.ResponseError) {
	args := params.(*values.SearchArgs)

	if args.Name == "thorw" {
		err := err.New("Falied getClients", 101, 400)

		return nil, err
	}

	return &c.ClientsMock, nil
}

func (c *ClientRepository) Create(client *entity.Client) (clientResult *entity.Client, err *values.ResponseError) {

	if client.Age == 100 {
		err := err.New("Falied create client", 100, 400)

		return nil, err
	}

	return client, nil
}

func (c *ClientRepository) Update(client *entity.Client) (rs bool, err *values.ResponseError) {
	if client.Name == "thorw" {
		err := err.New("Falied getClients", 201, 400)
		err.Append(202, "user.notFound")

		return false, err
	}

	return true, nil
}
