package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	graphql1 "graphql-api/domain/models/graphql"
	graphql2 "graphql-api/infrastructure/graphql"
	"graphql-api/interface/database/dataloader"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, params graphql1.CreatePostInput) (*graphql1.PostDetail, error) {
	return r.PostUsecase.Create(ctx, params)
}

// UpdatePost is the resolver for the UpdatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, params graphql1.UpdatePostInput) (*graphql1.PostDetail, error) {
	return r.PostUsecase.Update(ctx, params)
}

// Comments is the resolver for the comments field.
func (r *postDetailResolver) Comments(ctx context.Context, obj *graphql1.PostDetail) ([]*graphql1.CommentDetail, error) {
	return dataloader.LoadCommentByPostID(ctx, obj.ID)
}

// Post is the resolver for the Post field.
func (r *queryResolver) Post(ctx context.Context, id int64) (*graphql1.PostDetail, error) {
	return r.PostUsecase.Fetch(ctx, id)
}

// PostDetail returns graphql2.PostDetailResolver implementation.
func (r *Resolver) PostDetail() graphql2.PostDetailResolver { return &postDetailResolver{r} }

type postDetailResolver struct{ *Resolver }
