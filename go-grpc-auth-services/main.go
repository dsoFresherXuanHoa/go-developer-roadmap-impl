package main

import (
	"fmt"
	"go-grpc-auth-services/src/config"
)

func main() {
	if db, err := config.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else {
		const port = 3000
		entities := []interface{}{}
		db.AutoMigrate(entities...)

		fmt.Println("All entity has been synced to db!")
		RouteConfig(db, port)
	}
}
