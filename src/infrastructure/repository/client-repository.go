package infrastructure

import (
	db "webapi/src/db"
	values "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"
)

// ClientRepository - Struct repository
type ClientRepository struct {
	Database db.IDB
}

// Get - by params
func (c *ClientRepository) Get(params interface{}) (*[]entity.Client, *values.ResponseError) {

	clients := &[]entity.Client{}

	c.Database.GetModel(&entity.Client{}).Find(&clients)

	return clients, nil
}
