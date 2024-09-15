package database

import (
	"context"
	"graphql-api/domain/models"
	"graphql-api/infrastructure/db"
	"graphql-api/usecase/repository"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type userRepositoryImpl struct {
	db db.DBAdministrator
}

func NewUserRepository(db db.DBAdministrator) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Fetch(ctx context.Context, id int64) (*models.User, error) {
	user, err := models.Users(
		models.UserWhere.ID.EQ(id),
	).One(ctx, r.db.GetDao(ctx))
	return user, r.db.Error(err)
}

func (r *userRepositoryImpl) Create(ctx context.Context, m *models.User) error {
	return r.db.Error(m.Insert(ctx, r.db.GetDao(ctx), boil.Infer()))
}
