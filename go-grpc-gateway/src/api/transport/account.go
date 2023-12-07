package transport

import (
	"context"
	"fmt"
	"go-grpc-gateway/src/api/business"
	"go-grpc-gateway/src/api/storage"
	"go-grpc-gateway/src/config"
	"go-grpc-gateway/src/proto/gRPC/accounts"

	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
	accounts.UnimplementedAccountServiceServer
}

func NewGrpcServer(db *gorm.DB) *server {
	return &server{db: db}
}

func (s server) CreateAccountGrpc(ctx context.Context, req *accounts.AccountRequest) (*accounts.AccountResponse, error) {
	fmt.Println("Call account services via gRPC!")
	db, _ := config.GetGormInstance()
	storage := storage.NewSQLStore(db)
	business := business.NewCreateBusiness(storage)
	if res, err := business.CreateAccountGrpc(ctx, req); err != nil {
		fmt.Println("Error while create an account on GRPC Service call: " + err.Error())
		return nil, err
	} else {
		return res, nil
	}
}
