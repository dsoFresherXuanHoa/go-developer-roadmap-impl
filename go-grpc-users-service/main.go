package main

import (
	"fmt"
	"go-grpc-users-services/src/config"
	"go-grpc-users-services/src/module/entity"
	"go-grpc-users-services/src/module/transport"
	"go-grpc-users-services/src/proto/gRPC/users"
	"net"

	"google.golang.org/grpc"
)

func main() {
	if db, err := config.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else {
		address := "0.0.0.0:50052"
		entities := []interface{}{&entity.User{}}

		db.AutoMigrate(entities...)
		fmt.Println("All entity has been synced to db!")

		if lis, err := net.Listen("tcp", address); err != nil {
			fmt.Println("Error while start gRPC server (users service) at: " + address)
		} else {
			fmt.Println("Server (users service) is running at: " + address)
			server := transport.NewGrpcServer(db)
			s := grpc.NewServer()
			users.RegisterUserServiceServer(s, server)
			s.Serve(lis)
		}
	}
}
