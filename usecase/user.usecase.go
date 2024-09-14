package usecase

import (
	"context"
	"graphql-api/domain/models/graphql"
	"graphql-api/usecase/converter"
	"graphql-api/usecase/repository"
)

type UserUsecase interface {
	Fetch(ctx context.Context, id int64) (*graphql.UserDetail, error)
}

type userUsecaseImpl struct {
	repository repository.UserRepository
	converter	converter.UserConverter
}

func NewUserUsecase(r repository.UserRepository,
	c converter.UserConverter,
	) UserUsecase {
	return &userUsecaseImpl{
		repository: r,
		converter: c,
	}
}

func (u *userUsecaseImpl) Fetch(ctx context.Context, id int64) (*graphql.UserDetail, error) {
	us, err := u.repository.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.converter.ConvertUserModelToGraphQLType(us)
}
