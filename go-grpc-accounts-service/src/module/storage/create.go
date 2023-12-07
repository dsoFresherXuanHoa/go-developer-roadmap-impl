package storage

import (
	"context"
	"go-grpc-accounts-service/src/module/entity"
	"go-grpc-accounts-service/src/proto/gRPC/accounts"
)

func (s *sqlStorage) CreateAccount(ctx context.Context, account *entity.AccountCreatable) (*uint, error) {
	if err := s.db.Table(entity.AccountCreatable{}.GetTableName()).Create(&account).Error; err != nil {
		return nil, err
	}
	return &account.ID, nil
}

// TODO: It not needed! Remove as soon as possible!
func (s *sqlStorage) CreateAccountGrpc(ctx context.Context, req *accounts.Request) (*accounts.Response, error) {
	return nil, nil
}
