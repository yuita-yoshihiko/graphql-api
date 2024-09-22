package database

import (
	"context"
	"graphql-api/domain/models"
	"graphql-api/infrastructure/db"
	"graphql-api/usecase/repository"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type commentRepositoryImpl struct {
	db db.DBAdministrator
}

func NewCommentRepository(db db.DBAdministrator) repository.CommentRepository {
	return &commentRepositoryImpl{db: db}
}

func (r *commentRepositoryImpl) Fetch(ctx context.Context, id int64) (*models.Comment, error) {
	comment, err := models.Comments(
		models.CommentWhere.ID.EQ(id),
	).One(ctx, r.db.GetDao(ctx))
	return comment, r.db.Error(err)
}

func (r *commentRepositoryImpl) FetchByPostID(ctx context.Context, postID int64) ([]*models.Comment, error) {
	comments, err := models.Comments(
		models.CommentWhere.PostID.EQ(postID),
	).All(ctx, r.db.GetDao(ctx))
	return comments, r.db.Error(err)
}

func (r *commentRepositoryImpl) Create(ctx context.Context, m *models.Comment) error {
	return r.db.Error(m.Insert(ctx, r.db.GetDao(ctx), boil.Infer()))
}

func (r *commentRepositoryImpl) Update(ctx context.Context, m *models.Comment, columnsToUpdate []string) error {
	_, err := m.Update(ctx, r.db.GetDao(ctx), boil.Whitelist(columnsToUpdate...))
	return r.db.Error(err)
}
