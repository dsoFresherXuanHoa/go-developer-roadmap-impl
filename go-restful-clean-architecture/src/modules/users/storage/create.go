package storage

import (
	"context"
	"go-clean-architecture/src/modules/users/entity"
)

func (s *sqlStorage) Create(ctx context.Context, user *entity.UserCreatable) error {
	if err := s.db.Table(entity.UserCreatable{}.GetTableName()).Create(&user).Error; err != nil {
		return err
	}
	return nil
}
