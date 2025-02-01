package repository

import (
	"context"
	"graphql-api/domain/models"
)

type StaffRepository interface {
	Fetch(context.Context, int64) (*models.Staff, error)
	Create(context.Context, *models.Staff) error
	Update(context.Context, *models.Staff, []string) error
}
