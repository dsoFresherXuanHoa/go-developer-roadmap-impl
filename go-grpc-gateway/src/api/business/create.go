package business

import (
	"context"
	"fmt"
	"go-grpc-gateway/src/api/entity"
	"go-grpc-gateway/src/proto/gRPC/accounts"

	"golang.org/x/crypto/bcrypt"
)

type CreateStorage interface {
	CreateAccount(ctx context.Context, account *entity.AccountCreatable) (*uint, error)
}

type createBusiness struct {
	storage CreateStorage
}

func NewCreateBusiness(storage CreateStorage) *createBusiness {
	return &createBusiness{storage: storage}
}

func (business *createBusiness) CreateAccountGrpc(ctx context.Context, req *accounts.AccountRequest) (*accounts.AccountResponse, error) {
	account := entity.AccountCreatable{Email: &req.Email, Password: &req.Password}
	hashPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(*account.Password), 5)
	hashPassword := string(hashPasswordBytes)
	account.Password = &hashPassword
	if _, err := business.storage.CreateAccount(ctx, &account); err != nil {
		fmt.Println("Error while create an account in business: " + err.Error())
		return nil, err
	} else {
		return &accounts.AccountResponse{Email: *account.Email}, nil
	}
}
