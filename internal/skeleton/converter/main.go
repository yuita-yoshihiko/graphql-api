package main

import (
	"fmt"
	"graphql-api/utils"
	"log"
	"os"
	"path/filepath"
)

const Template = `package converter

import (
	"graphql-api/domain/models"
	"graphql-api/domain/models/graphql"
	"graphql-api/utils"
)

type {{ .Upper }}Converter interface {
	To{{ .Upper }}ModelFromCreateInput(graphql.{{ .Upper }}CreateInput) (*models.{{ .Upper }}, error)
	To{{ .Upper }}ModelFromUpdateInput(graphql.{{ .Upper }}UpdateInput) (*models.{{ .Upper }}, error)
	To{{ .Upper }}Detail(*models.{{ .Upper }}) (*graphql.{{ .Upper }}Detail, error)
	To{{ .Upper }}Details([]*models.{{ .Upper }}) ([]*graphql.{{ .Upper }}Detail, error)
	ToDBColumnsFromGraphQLFields(map[string]any) ([]string, error)
}

type {{ .Lower }}ConverterImpl struct {
}

func New{{ .Upper }}Converter() {{ .Upper }}Converter {
	return &{{ .Lower }}ConverterImpl{}
}

func (c {{ .Lower }}ConverterImpl) To{{ .Upper }}ModelFromCreateInput(input graphql.{{ .Upper }}CreateInput) (*models.{{ .Upper }}, error) {
	return nil, nil
}

func (c {{ .Lower }}ConverterImpl) To{{ .Upper }}ModelFromUpdateInput(input graphql.{{ .Upper }}UpdateInput) (*models.{{ .Upper }}, error) {
	return nil, nil
}

func (c {{ .Lower }}ConverterImpl) To{{ .Upper }}Detail(m *models.{{ .Upper }}) (*graphql.{{ .Upper }}Detail, error) {
	if m == nil {
		return nil, nil
	}
	return &graphql.{{ .Upper }}Detail{}, nil
}

func (c {{ .Lower }}ConverterImpl) To{{ .Upper }}Details(ms []*models.{{ .Upper}}) ([]*graphql.{{ .Upper }}Detail, error) {
	details := []*graphql.{{ .Upper }}Detail{}
	for _, m := range ms {
		detail, err := c.To{{ .Upper }}Detail(m)
		if err != nil {
			return nil, err
		} else if detail == nil {
			continue
		}
		details = append(details, detail)
	}
	return details, nil
}

func (c {{ .Lower }}ConverterImpl) ToDBColumnsFromGraphQLFields(rawArgs map[string]any) ([]string, error) {
	i := graphql.{{ .Upper }}Detail{}
	m := models.{{ .Upper }}{}
	return utils.ConvertRawArgsToColumnNames(rawArgs, i, m)
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
	filePath := filepath.Join("usecase", "converter", fmt.Sprintf("%v.converter.go", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, nil
	}
	log.Printf("Converterを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
