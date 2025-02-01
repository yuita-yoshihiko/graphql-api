package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
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
	if err != nil {
		return err
	}
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
	filePath := filepath.Join("usecase", "repository", fmt.Sprintf("%v.repository.go", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, nil
	}
	log.Printf("Repositoryを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
