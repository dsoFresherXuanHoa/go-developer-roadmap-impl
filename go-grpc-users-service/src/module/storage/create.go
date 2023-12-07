package storage

import (
	"context"
	"go-grpc-users-services/src/module/entity"
	"go-grpc-users-services/src/proto/gRPC/users"
)

func (s *sqlStorage) CreateUser(ctx context.Context, user *entity.UserCreatable) (*uint, error) {
	if err := s.db.Table(entity.UserCreatable{}.GetTableName()).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user.ID, nil
}

func (s *sqlStorage) CreateUserGrpc(ctx context.Context, req *users.Request) (*users.Response, error) {
	return nil, nil
}
