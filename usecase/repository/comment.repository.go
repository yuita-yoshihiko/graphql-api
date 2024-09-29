package repository

import (
	"context"
	"graphql-api/domain/models"
)

type CommentRepository interface {
	Fetch(context.Context, int64) (*models.Comment, error)
	FetchByPostIDs(context.Context, []int64) ([]*models.Comment, error)
	Create(context.Context, *models.Comment) error
	Update(context.Context, *models.Comment, []string) error
}
