package migration

import (
	"github.com/jinzhu/gorm"
	g "gopkg.in/gormigrate.v1"
)

// MigraCreateTableUser - Create table User
func MigraCreateTableUser() g.Migration {
	return g.Migration{
		ID: "20200302080000",
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("client").Error
		},
		Migrate: func(tx *gorm.DB) error {
			type Client struct {
				BaseModel
				CompanyID int
				Name      string
				Email     string
				Age       int
			}

			return tx.AutoMigrate(&Client{}).Error
		},
	}
}
