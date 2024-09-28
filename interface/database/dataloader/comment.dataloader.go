package dataloader

import (
	"context"
	"errors"
	"time"

	"github.com/graph-gophers/dataloader/v7"
	"graphql-api/domain/models/graphql"
	"graphql-api/usecase/converter"
	"graphql-api/usecase/repository"
)

type CommentDataLoaderKey struct{}

var CommentDataloaderKey CommentDataLoaderKey = CommentDataLoaderKey{}

type (
	CommentDataloader     Dataloader[int64, *graphql.CommentDetail]
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

func (d *commentDataloaderImpl) NewInterface() dataloader.Interface[int64, *graphql.CommentDetail] {
	var opt dataloader.Option[int64, *graphql.CommentDetail]
	if d.waitTime > 0 {
		opt = dataloader.WithWait[int64, *graphql.CommentDetail](d.waitTime)
	}
	return dataloader.NewBatchedLoader(d.BatchFunc(), opt)
}

func (d *commentDataloaderImpl) BatchFunc() dataloader.BatchFunc[int64, *graphql.CommentDetail] {
	return func(ctx context.Context, keys []int64) []*dataloader.Result[*graphql.CommentDetail] {
		comments, err := d.repository.FetchMany(ctx, keys)
		if err != nil {
			return MakeErrorResults[int64, *graphql.CommentDetail](keys, err)
		}

		if len(comments) == 0 {
			var results []*dataloader.Result[*graphql.CommentDetail]
			for _, key := range keys {
				results = append(results, &dataloader.Result[*graphql.CommentDetail]{Data: &graphql.CommentDetail{ID: key}, Error: nil})
			}
			return results
		}

		commentMap := make(map[int64]*graphql.CommentDetail, len(comments))
		for _, comment := range comments {
			commentMap[comment.ID], err = d.converter.ConvertCommentModelToGraphQLType(comment)
			if err != nil {
				return MakeErrorResults[int64, *graphql.CommentDetail](keys, err)
			}
		}

		if len(keys) != len(commentMap) {
			return MakeErrorResults[int64, *graphql.CommentDetail](keys, errors.New("an error occured, length defferent from keys and results"))
		}
		return SortResults(keys, commentMap, false)
	}
}
