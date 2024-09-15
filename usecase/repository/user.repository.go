package repository

import (
	"context"
	"graphql-api/domain/models"
)

type UserRepository interface {
	Fetch(context.Context, int64) (*models.User, error)
	Create(context.Context, *models.User) error
}
