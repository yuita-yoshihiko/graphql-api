package usecase

import (
	"context"
	"graphql-api/domain/models/graphql"
	"graphql-api/usecase/converter"
	"graphql-api/usecase/repository"
	"graphql-api/utils"
)

type UserUsecase interface {
	Fetch(context.Context, int64) (*graphql.UserDetail, error)
	Create(context.Context, graphql.CreateUserInput) (*graphql.UserDetail, error)
	Update(context.Context, graphql.UpdateUserInput) (*graphql.UserDetail, error)
	Delete(context.Context, int64) (*graphql.UserDetail, error)
}

type userUsecaseImpl struct {
	repository repository.UserRepository
	converter  converter.UserConverter
}

func NewUserUsecase(r repository.UserRepository,
	c converter.UserConverter,
) UserUsecase {
	return &userUsecaseImpl{
		repository: r,
		converter:  c,
	}
}

func (u *userUsecaseImpl) Fetch(ctx context.Context, id int64) (*graphql.UserDetail, error) {
	us, err := u.repository.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.converter.ConvertUserModelToGraphQLType(us)
}

func (u *userUsecaseImpl) Create(ctx context.Context, input graphql.CreateUserInput) (*graphql.UserDetail, error) {
	us, err := u.converter.ConvertUserGraphQLTypeToModel(input)
	if err != nil {
		return nil, err
	}
	if err := u.repository.Create(ctx, us); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, us.ID)
}

func (u *userUsecaseImpl) Update(ctx context.Context, input graphql.UpdateUserInput) (*graphql.UserDetail, error) {
	columns, err := u.converter.ConvertRawArgsToDBColumnNames(utils.GetGraphQLFields(ctx))
	if err != nil {
		return nil, err
	}
	us, err := u.converter.ConvertUserGraphQLTypeToModelUpdate(input)
	if err != nil {
		return nil, err
	}
	us.ID = input.ID
	if err := u.repository.Update(ctx, us, columns); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, us.ID)
}

func (u *userUsecaseImpl) Delete(ctx context.Context, id int64) (*graphql.UserDetail, error) {
	us, err := u.repository.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := u.repository.Delete(ctx, us); err != nil {
		return nil, err
	}
	return u.converter.ConvertUserModelToGraphQLType(us)
}
