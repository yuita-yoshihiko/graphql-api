package dataloader

import (
	"context"
	"time"

	"graphql-api/domain/models/graphql"
	"graphql-api/usecase/converter"
	"graphql-api/usecase/repository"

	"github.com/graph-gophers/dataloader/v7"
)

var CDataloaderKey CommentDataloaderKey = CommentDataloaderKey{}

type (
	CommentDataloaderKey  struct{}
	CommentDataloader     Dataloader[int64, []*graphql.CommentDetail]
	commentDataloaderImpl struct {
		converter  converter.CommentConverter
		repository repository.CommentRepository
		waitTime   time.Duration
	}
)

func NewCommentDataLoader(c converter.CommentConverter, r repository.CommentRepository, wt time.Duration) CommentDataloader {
	return &commentDataloaderImpl{
		converter:  c,
		repository: r,
		waitTime:   wt,
	}
}

func (d *commentDataloaderImpl) NewInterface() dataloader.Interface[int64, []*graphql.CommentDetail] {
	var opt dataloader.Option[int64, []*graphql.CommentDetail]
	if d.waitTime > 0 {
		opt = dataloader.WithWait[int64, []*graphql.CommentDetail](d.waitTime)
	}
	return dataloader.NewBatchedLoader(d.BatchFunc(), opt)
}

func (d *commentDataloaderImpl) BatchFunc() dataloader.BatchFunc[int64, []*graphql.CommentDetail] {
	return func(ctx context.Context, keys []int64) []*dataloader.Result[[]*graphql.CommentDetail] {
		comments, err := d.repository.FetchByPostIDs(ctx, keys)
		if err != nil {
			return MakeErrorResults[int64, []*graphql.CommentDetail](keys, err)
		}

		commentMap := make(map[int64][]*graphql.CommentDetail, len(keys))
		for _, comment := range comments {
			commentDetail, err := d.converter.ConvertCommentModelToGraphQLType(comment)
			if err != nil {
				return MakeErrorResults[int64, []*graphql.CommentDetail](keys, err)
			}
			commentMap[comment.PostID] = append(commentMap[comment.PostID], commentDetail)
		}

		return SortResults[int64, []*graphql.CommentDetail](keys, commentMap, true)
	}
}

func LoadCommentByPostID(ctx context.Context, id int64) ([]*graphql.CommentDetail, error) {
	loader := DataloaderFromCtx[CommentDataloaderKey, dataloader.Interface[int64, []*graphql.CommentDetail]](ctx, CDataloaderKey)
	thunk := loader.Load(ctx, id)
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result, nil
}
