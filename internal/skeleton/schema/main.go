package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
)

const Template = `# ---------- Query and Mutation -----------
extend type Query {
  {{ .Upper }}(ID: ID!): {{ .Upper }}Detail!
}

extend type Mutation {
  Create{{ .Upper }}(params: {{ .Upper }}CreateInput!): {{ .Upper }}Detail!
  Update{{ .Upper }}(params: {{ .Upper }}UpdateInput!): {{ .Upper }}Detail!
}

# ---------- Response Type -----------

type {{ .Upper }}Detail {
	"""field"""
}

# ---------- Params Type -----------
input {{ .Upper }}CreateInput {
	"""field"""
}

input {{ .Upper }}UpdateInput {
	"""field"""
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
	filePath := filepath.Join("schema", fmt.Sprintf("%v.graphql", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, nil
	}
	log.Printf("schemaを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
