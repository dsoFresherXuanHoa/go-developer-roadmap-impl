package business

import (
	"context"

	"go-grpc-auth-services/src/module/entity"

	"gorm.io/gorm"
)

type RegisterStorage interface {
	SignUp(ctx context.Context, register *entity.RegisterCreatable, db *gorm.DB) (*uint, *uint, error)
}

type registerBusiness struct {
	storage RegisterStorage
}

func NewRegisterBusiness(storage RegisterStorage) *registerBusiness {
	return &registerBusiness{storage: storage}
}

func (business *registerBusiness) SignUp(ctx context.Context, register *entity.RegisterCreatable, db *gorm.DB) (*uint, *uint, error) {
	// TODO: Add Validate Here as soon as possible
	if employeeId, accountId, err := business.storage.SignUp(ctx, register, db); err != nil {
		return nil, nil, err
	} else {
		return employeeId, accountId, nil
	}
}
