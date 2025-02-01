package main

import (
	"fmt"
	"graphql-api/utils"
	"log"
	"os"
	"path/filepath"
)

const Template = `package usecase

import (
	"context"
	"graphql-api/domain/models/graphql"
	"graphql-api/usecase/converter"
	"graphql-api/usecase/repository"
	"graphql-api/utils"
)

type {{ .Upper }}UseCase interface {
	Fetch(context.Context, int64) (*graphql.{{ .Upper }}Detail, error)
	Create(context.Context, graphql.{{ .Upper }}CreateInput) (*graphql.{{ .Upper }}Detail, error)
	Update(context.Context, graphql.{{ .Upper }}UpdateInput) (*graphql.{{ .Upper }}Detail, error)
}

type {{ .Lower }}UseCaseImpl struct {
	repository repository.{{ .Upper }}Repository
	converter  converter.{{ .Upper }}Converter
}

func New{{ .Upper }}UseCase(r repository.{{ .Upper }}Repository,
	c converter.{{ .Upper }}Converter,
) {{ .Upper }}UseCase {
	return &{{ .Lower }}UseCaseImpl{
		repository: r,
		converter:  c,
	}
}

func (u *{{ .Lower }}UseCaseImpl) Fetch(ctx context.Context, id int64) (*graphql.{{ .Upper }}Detail, error) {
	m, err := u.repository.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.converter.To{{ .Upper }}Detail(m)
}

func (u *{{ .Lower }}UseCaseImpl) Create(ctx context.Context, input graphql.{{ .Upper }}CreateInput) (*graphql.{{ .Upper }}Detail, error) {
	m, err := u.converter.To{{ .Upper }}ModelFromCreateInput(input)
	if err != nil {
		return nil, err
	}
	if err := u.repository.Create(ctx, m); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, m.ID)
}

func (u *{{ .Lower }}UseCaseImpl) Update(ctx context.Context, input graphql.{ .Upper }}Update{Input) (*graphql.{{ .Upper }}Detail, error) {
	columns, err := u.converter.ToDBColumnsFromGraphQLFields(utils.GetGraphQLFields(ctx))
	if err != nil {
		return nil, err
	}
	m, err := u.converter.To{{ .Upper }}ModelFromUpdateInput(input)
	if err != nil {
		return nil, err
	}
	m.ID = input.ID
	if err := u.repository.Update(ctx, m, columns); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, m.ID)
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
	filePath := filepath.Join("usecase", fmt.Sprintf("%v.usecase.go", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, nil
	}
	log.Printf("UseCaseを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
