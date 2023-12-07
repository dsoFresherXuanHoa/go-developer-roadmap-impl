package business

import (
	"context"
	"errors"
	"go-clean-architecture/src/modules/users/entity"
)

type CreateStorage interface {
	Create(ctx context.Context, user *entity.UserCreatable) error
}

type createBusiness struct {
	storage CreateStorage
}

func NewCreateBusiness(storage CreateStorage) *createBusiness {
	return &createBusiness{storage: storage}
}

func (business *createBusiness) Create(ctx context.Context, user *entity.UserCreatable) error {
	if user.Username == nil {
		return errors.New("username can't be blank")
	} else if err := business.storage.Create(ctx, user); err != nil {
		return err
	}
	return nil
}
