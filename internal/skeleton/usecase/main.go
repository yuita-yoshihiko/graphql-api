package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
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
	Create(context.Context, graphql.Create{{ .Upper }}Input) (*graphql.{{ .Upper }}Detail, error)
	Update(context.Context, graphql.Update{{ .Upper }}Input) (*graphql.{{ .Upper }}Detail, error)
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
	return u.converter.Convert{{ .Upper }}ModelToGraphQLType(m)
}

func (u *{{ .Lower }}UseCaseImpl) Create(ctx context.Context, input graphql.Create{{ .Upper }}Input) (*graphql.{{ .Upper }}Detail, error) {
	m, err := u.converter.Convert{{ .Upper }}GraphQLTypeToModel(input)
	if err != nil {
		return nil, err
	}
	if err := u.repository.Create(ctx, m); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, m.ID)
}

func (u *{{ .Lower }}UseCaseImpl) Update(ctx context.Context, input graphql.Update{{ .Upper }}Input) (*graphql.{{ .Upper }}Detail, error) {
	columns, err := u.converter.ConvertRawArgsToDBColumnNames(utils.GetRawArgs(ctx))
	if err != nil {
		return nil, err
	}
	m, err := u.converter.Convert{{ .Upper }}GraphQLTypeToModelUpdate(input)
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

type name string

func (n name) Upper() string {
	return strcase.ToCamel(string(n))
}

func (n name) Lower() string {
	return strcase.ToLowerCamel(string(n))
}

var list []name = []name{
	"staff",
}

func main() {
	log.Println("開始")
	for _, m := range list {
		if err := templateExport(m); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("完了")
}

func templateExport(m name) error {
	tpl, err := template.New("").Parse(Template)

	file, err := createGoFile(m.Lower())
	if err != nil {
		return err
	} else if file == nil {
		return nil
	}
	defer file.Close()
	return tpl.Execute(file, m)
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
