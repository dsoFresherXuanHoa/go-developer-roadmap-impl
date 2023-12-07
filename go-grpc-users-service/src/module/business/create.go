package business

import (
	"context"
	"go-grpc-users-services/src/module/entity"
	"go-grpc-users-services/src/proto/gRPC/users"
)

type CreateStorage interface {
	CreateUser(ctx context.Context, user *entity.UserCreatable) (*uint, error)
	CreateUserGrpc(ctx context.Context, req *users.Request) (*users.Response, error)
}

type createBusiness struct {
	storage CreateStorage
}

func NewCreateBusiness(storage CreateStorage) *createBusiness {
	return &createBusiness{storage: storage}
}

func (business *createBusiness) CreateUser(ctx context.Context, user *entity.UserCreatable) (*uint, error) {
	if id, err := business.storage.CreateUser(ctx, user); err != nil {
		return nil, err
	} else {
		return id, nil
	}
}

func (business *createBusiness) CreateUserGrpc(ctx context.Context, req *users.Request) (*users.Response, error) {
	user := entity.UserCreatable{FullName: req.FullName}
	if id, err := business.storage.CreateUser(ctx, &user); err != nil {
		return nil, err
	} else {
		return &users.Response{UserId: int32(*id)}, nil
	}
}
