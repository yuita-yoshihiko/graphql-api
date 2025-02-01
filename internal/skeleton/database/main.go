package main

import (
	"fmt"
	"graphql-api/utils"
	"log"
	"os"
	"path/filepath"
)

const Template = `package database

import (
	"context"
	"graphql-api/domain/models"
	"graphql-api/infrastructure/db"
	"graphql-api/usecase/repository"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type {{ .Lower }}RepositoryImpl struct {
	db db.DBAdministrator
}

func New{{ .Upper }}Repository(db db.DBAdministrator) repository.{{ .Upper }}Repository {
	return &{{ .Lower }}RepositoryImpl{db: db}
}

func (r *{{ .Lower }}RepositoryImpl) Fetch(ctx context.Context, id int64) (*models.{{ .Upper }}, error) {
	m, err := models.{{ .Upper }}s(
		models.{{ .Upper }}Where.ID.EQ(id),
	).One(ctx, r.db.GetDao(ctx))
	return m, r.db.Error(err)
}

func (r *{{ .Lower }}RepositoryImpl) Create(ctx context.Context, m *models.{{ .Upper }}) error {
	return r.db.Error(m.Insert(ctx, r.db.GetDao(ctx), boil.Infer()))
}

func (r *{{ .Lower }}RepositoryImpl) Update(ctx context.Context, m *models.{{ .Upper }}, columnsToUpdate []string) error {
	_, err := m.Update(ctx, r.db.GetDao(ctx), boil.Whitelist(columnsToUpdate...))
	return r.db.Error(err)
}

`

func main() {
	log.Println("開始")
	list := utils.GetNameList()
	for _, m := range list {
		err := utils.TemplateExport(m, func(name string) (*os.File, error) {
			return createGoFile(name)
		}, Template)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("完了")
}

func createGoFile(name string) (*os.File, error) {
	filePath := filepath.Join("interface", "database", fmt.Sprintf("%vRepository.go", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, nil
	}
	log.Printf("databaseを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
