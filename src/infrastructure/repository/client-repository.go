package infrastructure

import (
	db "webapi/src/db"
	entity "webapi/src/domain/entity"
)

// ClientRepository - Struct repository
type ClientRepository struct {
	Database db.IDB
}

// Get - by params
func (c *ClientRepository) Get(params interface{}) (error, []entity.Client) {

	return nil, []entity.Client{
		{
			ID:        "",
			Name:      "JPNobrega",
			Age:       21,
			CompanyID: 0,
		},
	}
}
