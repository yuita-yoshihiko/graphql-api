package main

import (
	"fmt"
	"graphql-api/utils"
	"log"
	"os"
	"path/filepath"
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
	filePath := filepath.Join("schema", fmt.Sprintf("%v.graphql", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, nil
	}
	log.Printf("schemaを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
