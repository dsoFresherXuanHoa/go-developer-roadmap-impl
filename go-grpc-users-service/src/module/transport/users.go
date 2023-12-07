package transport

import (
	"context"
	"fmt"
	"go-grpc-users-services/src/config"
	"go-grpc-users-services/src/module/business"
	"go-grpc-users-services/src/module/storage"
	"go-grpc-users-services/src/proto/gRPC/users"

	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
	users.UnimplementedUserServiceServer
}

func NewGrpcServer(db *gorm.DB) *server {
	return &server{db: db}
}

func (s server) CreateUserGrpc(ctx context.Context, req *users.Request) (*users.Response, error) {
	fmt.Println("Call users services via gRPC!")
	// TODO: Look after this db instance
	db, _ := config.GetGormInstance()
	storage := storage.NewSQLStore(db)
	business := business.NewCreateBusiness(storage)
	if res, err := business.CreateUserGrpc(ctx, req); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
