package main

import (
	DB "webapi/src/db"
	helper "webapi/src/infrastructure/helper"
	router "webapi/src/infrastructure/router"
)

func main() {
	server := helper.CreateHTTPServer()

	database := &DB.DB{}
	database.Connect("postgres", "postgres://postgres:postgres@localhost/workout?sslmode=disable")

	router.Build(server)

	server.Logger.Fatal(server.Start(":8082"))
}
