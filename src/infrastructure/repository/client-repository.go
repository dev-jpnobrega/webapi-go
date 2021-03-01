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
func (c *ClientRepository) Get(params interface{}) ([]entity.Client, *values.ResponseError) {

	return []entity.Client{
			{
				ID:        "",
				Name:      "JPNobrega",
				Age:       21,
				CompanyID: 0,
			},
		},
		nil
}
