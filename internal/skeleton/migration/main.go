package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
)

const Template = `
-- +migrate Up
CREATE TABLE IF NOT EXISTS {{ .Lower }}s(
    id bigserial not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now(),
    primary key(id)
);
comment on table {{ .Lower }}s is 'テーブルの説明';
comment on column {{ .Lower }}s. is 'カラムの説明';
comment on column {{ .Lower }}s. is 'カラムの説明';

-- +migrate Down
DROP TABLE IF EXISTS {{ .Lower }}s;

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
	filePath := filepath.Join("infrastructure", "db", "migrations", fmt.Sprintf("%s-create_%vs.sql", time.Now().Format("20060102150405"), name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, nil
	}
	log.Printf("migrationを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
