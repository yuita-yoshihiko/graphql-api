package usecase

import (
	"context"
	"graphql-api/domain/models/graphql"
	"graphql-api/usecase/converter"
	"graphql-api/usecase/repository"
)

type CommentUsecase interface {
	Fetch(context.Context, int64) (*graphql.CommentDetail, error)
	Create(context.Context, graphql.CreateCommentInput) (*graphql.CommentDetail, error)
	Update(context.Context, graphql.UpdateCommentInput) (*graphql.CommentDetail, error)
}

type commentUsecaseImpl struct {
	repository repository.CommentRepository
	converter  converter.CommentConverter
}

func NewCommentUsecase(r repository.CommentRepository,
	c converter.CommentConverter,
) CommentUsecase {
	return &commentUsecaseImpl{
		repository: r,
		converter:  c,
	}
}

func (u *commentUsecaseImpl) Fetch(ctx context.Context, id int64) (*graphql.CommentDetail, error) {
	c, err := u.repository.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.converter.ConvertCommentModelToGraphQLType(c)
}

func (u *commentUsecaseImpl) Create(ctx context.Context, input graphql.CreateCommentInput) (*graphql.CommentDetail, error) {
	c, err := u.converter.ConvertCommentGraphQLTypeToModel(input)
	if err != nil {
		return nil, err
	}
	if err := u.repository.Create(ctx, c); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, c.ID)
}

func (u *commentUsecaseImpl) Update(ctx context.Context, input graphql.UpdateCommentInput) (*graphql.CommentDetail, error) {
	columns := u.converter.ConvertRawArgsToDBColumnNames()
	c, err := u.converter.ConvertCommentGraphQLTypeToModelUpdate(input)
	if err != nil {
		return nil, err
	}
	c.ID = input.ID
	if err := u.repository.Update(ctx, c, columns); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, c.ID)
}
