package main

import (
	"gin-restful-mongo/configs"
	"gin-restful-mongo/routes"
)

func main() {
	if client, err := configs.MongoDriverConfig(); err != nil {
		panic("Can't connect to database! Try again later!")
	} else {
		routes.RouterConfig(client)
	}
}
