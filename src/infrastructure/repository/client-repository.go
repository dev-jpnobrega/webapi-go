package infrastructure

import (
	"log"
	db "webapi/src/db"
	values "webapi/src/domain/contract/value"
	entity "webapi/src/domain/entity"
)

// A ClientRepository represent client database
type ClientRepository struct {
	Database db.IDB
}

func (c *ClientRepository) Get(params interface{}) (*[]entity.Client, *values.ResponseError) {
	clients := &[]entity.Client{}

	c.Database.GetModel(&entity.Client{}).Find(&clients)

	return clients, nil
}

func (c *ClientRepository) Create(client *entity.Client) (*entity.Client, *values.ResponseError) {
	clientResult := c.Database.GetModel(&entity.Client{}).Create(client).Value

	log.Println(clientResult)

	return client, nil
}

func (c *ClientRepository) Update(client *entity.Client) (bool, *values.ResponseError) {
	result := c.Database.GetModel(&entity.Client{}).Updates(client).Error

	log.Println(result)

	return true, nil
}
