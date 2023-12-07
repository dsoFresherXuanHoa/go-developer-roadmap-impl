package main

import (
	"context"
	"fmt"
	"go-grpc-users-services/src/proto/gRPC/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	userServiceHost := "localhost:50052"
	if cc, err := grpc.Dial(userServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		fmt.Println("Error while connect to user service at: " + userServiceHost)
	} else {
		// TODO: Save User
		defer cc.Close()
		fullName := "Xuan Hoa Le"
		req := users.Request{FullName: fullName}
		client := users.NewUserServiceClient(cc)
		if _, err := client.CreateUserGrpc(context.Background(), &req); err != nil {
			fmt.Println("Error at User Service: " + err.Error())
		}
	}
}
