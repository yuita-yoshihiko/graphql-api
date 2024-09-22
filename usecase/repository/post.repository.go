package repository

import (
	"context"
	"graphql-api/domain/models"
)

type PostRepository interface {
	Fetch(context.Context, int64) (*models.Post, error)
	FetchByUserID(context.Context, int64) ([]*models.Post, error)
	Create(context.Context, *models.Post) error
	Update(context.Context, *models.Post, []string) error
}