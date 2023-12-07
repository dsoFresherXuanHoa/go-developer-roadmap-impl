package storage

import (
	"context"
	"fmt"
	"go-grpc-gateway/src/api/entity"
)

func (s *sqlStorage) CreateAccount(ctx context.Context, account *entity.AccountCreatable) (*uint, error) {
	if err := s.db.Table(entity.AccountCreatable{}.GetTableName()).Create(&account).Error; err != nil {
		fmt.Println("Error while create an account in storage: " + err.Error())
		return nil, err
	}
	return &account.ID, nil
}
