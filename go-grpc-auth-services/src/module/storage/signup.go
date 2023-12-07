package storage

import (
	"context"
	"fmt"

	"go-grpc-auth-services/src/module/entity"
	"go-grpc-auth-services/src/proto/gRPC/accounts"
	"go-grpc-auth-services/src/proto/gRPC/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
)

func (s *sqlStorage) SignUp(ctx context.Context, register *entity.RegisterCreatable, db *gorm.DB) (*uint, *uint, error) {
	userServiceHost := "localhost:50052"
	accountServiceHost := "localhost:50051"
	if cc, err := grpc.Dial(userServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		fmt.Println("Error while connect to user service at: " + userServiceHost)
		return nil, nil, err
	} else {
		// TODO: Save User
		defer cc.Close()
		req := users.Request{FullName: *register.FullName}
		client := users.NewUserServiceClient(cc)
		if userId, err := client.CreateUserGrpc(ctx, &req); err != nil {
			return nil, nil, err
		} else if cc, err := grpc.Dial(accountServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
			fmt.Println("Error while connect to account service at: " + accountServiceHost)
			return nil, nil, err
		} else {
			// TODO: Save user
			defer cc.Close()
			req := accounts.Request{Username: *register.Username, Password: *register.Password, UserId: userId.UserId}
			fmt.Println(req)
			client := accounts.NewAccountServiceClient(cc)
			if accountId, err := client.CreateAccountGrpc(ctx, &req); err != nil {
				return nil, nil, err
			} else {
				userId := uint(userId.UserId)
				accountId := uint(accountId.AccountId)
				return &userId, &accountId, nil
			}
		}
	}
}
