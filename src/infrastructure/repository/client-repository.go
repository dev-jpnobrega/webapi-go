package infrastructure

import (
	db "webapi/src/db"
	values "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"

	"github.com/labstack/gommon/log"
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

// Create - by entity.Client
func (c *ClientRepository) Create(client *entity.Client) (*entity.Client, *values.ResponseError) {

	clientRestul := c.Database.GetModel(&entity.Client{}).Create(client).Value

	log.Info(clientRestul)

	return client, nil
}

// Update - by client props
func (c *ClientRepository) Update(client *entity.Client) (bool, *values.ResponseError) {

	result := c.Database.GetModel(&entity.Client{}).Updates(client).Error

	log.Info(result)

	return true, nil
}
