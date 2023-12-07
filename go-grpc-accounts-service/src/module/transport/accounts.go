package transport

import (
	"context"
	"fmt"
	"go-grpc-accounts-service/src/config"
	"go-grpc-accounts-service/src/module/business"
	"go-grpc-accounts-service/src/module/storage"
	"go-grpc-accounts-service/src/proto/gRPC/accounts"

	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
	accounts.UnimplementedAccountServiceServer
}

func NewGrpcServer(db *gorm.DB) *server {
	return &server{db: db}
}

func (s server) CreateAccountGrpc(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
	fmt.Println("Call account services via gRPC!")
	// TODO: Look after this db instance
	db, _ := config.GetGormInstance()
	storage := storage.NewSQLStore(db)
	business := business.NewCreateBusiness(storage)
	if res, err := business.CreateAccountGrpc(ctx, req); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
