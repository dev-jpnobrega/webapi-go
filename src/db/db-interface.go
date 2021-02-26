package db

import (
	"github.com/jinzhu/gorm"
)

// IDB - Database interface
type IDB interface {
	Connect(driver string, connectionString string) (*gorm.DB, error)
	Get() *gorm.DB
	GetModel(model interface{}) *gorm.DB
}
