package usecase

import (
	"context"
	"graphql-api/domain/models/graphql"
	dbDataloader "graphql-api/interface/database/dataloader"
	"graphql-api/usecase/converter"
	"graphql-api/usecase/repository"
	"graphql-api/utils"

	"github.com/graph-gophers/dataloader/v7"
)

type CommentUsecase interface {
	Fetch(context.Context, int64) (*graphql.CommentDetail, error)
	FetchByPostID(context.Context, int64) ([]*graphql.CommentDetail, error)
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

func (u *commentUsecaseImpl) FetchByPostID(ctx context.Context, id int64) ([]*graphql.CommentDetail, error) {
	return dbDataloader.LoaderFor[dbDataloader.CommentDataloaderKey, dataloader.Interface[int64, []*graphql.CommentDetail]](
		ctx,
		dbDataloader.CDataloaderKey,
	).Load(ctx, id)()
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
	columns, err := u.converter.ConvertRawArgsToDBColumnNames(utils.GetRawArgs(ctx))
	if err != nil {
		return nil, err
	}
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
