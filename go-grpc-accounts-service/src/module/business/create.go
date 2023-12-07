package business

import (
	"context"
	"fmt"

	"go-grpc-accounts-service/src/module/entity"
	"go-grpc-accounts-service/src/proto/gRPC/accounts"

	"golang.org/x/crypto/bcrypt"
)

type CreateStorage interface {
	CreateAccount(ctx context.Context, account *entity.AccountCreatable) (*uint, error)
	CreateAccountGrpc(ctx context.Context, req *accounts.Request) (*accounts.Response, error)
}

type createBusiness struct {
	storage CreateStorage
}

func NewCreateBusiness(storage CreateStorage) *createBusiness {
	return &createBusiness{storage: storage}
}

func (business *createBusiness) CreateAccount(ctx context.Context, account *entity.AccountCreatable) (*uint, error) {
	hashPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(*account.Password), 5)
	hashPassword := string(hashPasswordBytes)
	account.Password = &hashPassword
	if id, err := business.storage.CreateAccount(ctx, account); err != nil {
		return nil, err
	} else {
		return id, nil
	}
}

func (business *createBusiness) CreateAccountGrpc(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
	userId := uint(req.UserId)
	account := entity.AccountCreatable{Username: &req.Username, Password: &req.Password, UserId: &userId}
	hashPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(*account.Password), 5)
	hashPassword := string(hashPasswordBytes)
	account.Password = &hashPassword
	fmt.Println(account)
	if id, err := business.storage.CreateAccount(ctx, &account); err != nil {
		return nil, err
	} else {
		return &accounts.Response{AccountId: int32(*id)}, nil
	}
}
