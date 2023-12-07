package main

import (
	"fmt"
	"go-grpc-accounts-service/src/config"
	"go-grpc-accounts-service/src/module/entity"
	"go-grpc-accounts-service/src/module/transport"
	"go-grpc-accounts-service/src/proto/gRPC/accounts"
	"net"

	"google.golang.org/grpc"
)

func main() {
	if db, err := config.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else {
		address := "0.0.0.0:50051"
		entities := []interface{}{&entity.Account{}}

		db.AutoMigrate(entities...)
		fmt.Println("All entity has been synced to db!")

		if lis, err := net.Listen("tcp", address); err != nil {
			fmt.Println("Error while start gRPC server (account service) at: " + address)
		} else {
			fmt.Println("Server (account service) is running at: " + address)
			s := grpc.NewServer()
			server := transport.NewGrpcServer(db)
			accounts.RegisterAccountServiceServer(s, server)
			s.Serve(lis)
		}
	}

}
