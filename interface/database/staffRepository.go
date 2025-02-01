package database

import (
	"context"
	"graphql-api/domain/models"
	"graphql-api/infrastructure/db"
	"graphql-api/usecase/repository"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type staffRepositoryImpl struct {
	db db.DBAdministrator
}

func NewStaffRepository(db db.DBAdministrator) repository.StaffRepository {
	return &staffRepositoryImpl{db: db}
}

func (r *staffRepositoryImpl) Fetch(ctx context.Context, id int64) (*models.Staff, error) {
	m, err := models.Staffs(
		models.StaffWhere.ID.EQ(id),
	).One(ctx, r.db.GetDao(ctx))
	return m, r.db.Error(err)
}

func (r *staffRepositoryImpl) Create(ctx context.Context, m *models.Staff) error {
	return r.db.Error(m.Insert(ctx, r.db.GetDao(ctx), boil.Infer()))
}

func (r *staffRepositoryImpl) Update(ctx context.Context, m *models.Staff, columnsToUpdate []string) error {
	_, err := m.Update(ctx, r.db.GetDao(ctx), boil.Whitelist(columnsToUpdate...))
	return r.db.Error(err)
}
