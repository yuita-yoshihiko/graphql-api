package main

import (
	"fmt"
	"graphql-api/utils"
	"log"
	"os"
	"path/filepath"
)

const Template = `package repository

import (
	"context"
	"graphql-api/domain/models"
)

type {{ .Upper }}Repository interface {
	Fetch(context.Context, int64) (*models.{{ .Upper }}, error)
	Create(context.Context, *models.{{ .Upper }}) error
	Update(context.Context, *models.{{ .Upper }}, []string) error
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
	filePath := filepath.Join("usecase", "repository", fmt.Sprintf("%v.repository.go", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, nil
	}
	log.Printf("Repositoryを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
