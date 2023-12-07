package main

import (
	"fmt"
	"go-clean-architecture/src/configs"
	"go-clean-architecture/src/modules/users/entity"
)

func main() {
	if db, err := configs.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else {
		const port = 3000

		entities := []interface{}{&entity.User{}}
		db.AutoMigrate(entities...)
		fmt.Println("All entity has been synced to db!")
		RouteConfig(db, port)
	}
}
