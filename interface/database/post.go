package database

import (
	"context"
	"graphql-api/domain/models"
	"graphql-api/infrastructure/db"
	"graphql-api/usecase/repository"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type postRepositoryImpl struct {
	db db.DBAdministrator
}

func NewPostRepository(db db.DBAdministrator) repository.PostRepository {
	return &postRepositoryImpl{db: db}
}

func (r *postRepositoryImpl) Fetch(ctx context.Context, id int64) (*models.Post, error) {
	post, err := models.Posts(
		models.PostWhere.ID.EQ(id),
	).One(ctx, r.db.GetDao(ctx))
	return post, r.db.Error(err)
}

func (r *postRepositoryImpl) FetchByUserID(ctx context.Context, userID int64) ([]*models.Post, error) {
	posts, err := models.Posts(
		models.PostWhere.UserID.EQ(userID),
	).All(ctx, r.db.GetDao(ctx))
	return posts, r.db.Error(err)
}

func (r *postRepositoryImpl) Create(ctx context.Context, m *models.Post) error {
	return r.db.Error(m.Insert(ctx, r.db.GetDao(ctx), boil.Infer()))
}

func (r *postRepositoryImpl) Update(ctx context.Context, m *models.Post, columnsToUpdate []string) error {
	_, err := m.Update(ctx, r.db.GetDao(ctx), boil.Whitelist(columnsToUpdate...))
	return r.db.Error(err)
}
