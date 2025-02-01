package main

import (
	"fmt"
	"graphql-api/utils"
	"log"
	"os"
	"path/filepath"
	"time"
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
	filePath := filepath.Join("infrastructure", "db", "migrations", fmt.Sprintf("%s-create_%vs.sql", time.Now().Format("20060102150405"), name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, nil
	}
	log.Printf("migrationを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
