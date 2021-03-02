package main

import (
	"log"
	migration "webapi/src/db/migrations"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
)

func main() {

	db, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost/workout?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	defer db.Close()

	createTableUser := migration.MigraCreateTableUser()

	mm := []*gormigrate.Migration{
		&createTableUser,
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, mm)

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
