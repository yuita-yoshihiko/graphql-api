package usecase

import (
	"context"
	"graphql-api/domain/models/graphql"
	"graphql-api/usecase/converter"
	"graphql-api/usecase/repository"
	"graphql-api/utils"
)

type PostUsecase interface {
	Fetch(context.Context, int64) (*graphql.PostDetail, error)
	FetchByUserID(context.Context, int64) ([]*graphql.PostDetail, error)
	Create(context.Context, graphql.CreatePostInput) (*graphql.PostDetail, error)
	Update(context.Context, graphql.UpdatePostInput) (*graphql.PostDetail, error)
}

type postUsecaseImpl struct {
	repository repository.PostRepository
	converter  converter.PostConverter
}

func NewPostUsecase(r repository.PostRepository,
	c converter.PostConverter,
) PostUsecase {
	return &postUsecaseImpl{
		repository: r,
		converter:  c,
	}
}

func (u *postUsecaseImpl) Fetch(ctx context.Context, id int64) (*graphql.PostDetail, error) {
	p, err := u.repository.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.converter.ConvertPostModelToGraphQLType(p)
}

func (u *postUsecaseImpl) FetchByUserID(ctx context.Context, userID int64) ([]*graphql.PostDetail, error) {
	ps, err := u.repository.FetchByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return u.converter.ConvertPostModelsToGraphQLTypes(ps)
}

func (u *postUsecaseImpl) Create(ctx context.Context, input graphql.CreatePostInput) (*graphql.PostDetail, error) {
	p, err := u.converter.ConvertPostGraphQLTypeToModel(input)
	if err != nil {
		return nil, err
	}
	if err := u.repository.Create(ctx, p); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, p.ID)
}

func (u *postUsecaseImpl) Update(ctx context.Context, input graphql.UpdatePostInput) (*graphql.PostDetail, error) {
	columns, err := u.converter.ConvertRawArgsToDBColumnNames(utils.GetGraphQLFields(ctx))
	if err != nil {
		return nil, err
	}
	p, err := u.converter.ConvertPostGraphQLTypeToModelUpdate(input)
	if err != nil {
		return nil, err
	}
	p.ID = input.ID
	if err := u.repository.Update(ctx, p, columns); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, p.ID)
}
