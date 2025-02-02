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

const TestTemplate = `package usecase_test

import (
	"context"
	"graphql-api/domain/models/graphql"
	"graphql-api/infrastructure/db"
	"graphql-api/interface/database"
	"graphql-api/usecase"
	"graphql-api/usecase/converter"
	"graphql-api/utils/testhelper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test{{ .Upper }}Fetch(t *testing.T) {
	t.Helper()
	d := testhelper.LoadFixture(
		"../",
		"testdata/{{ .Lower }}s/fixtures/1",
	)
	dbAdministrator := db.NewDBAdministrator(d)
	uc := usecase.New{{ .Upper }}Usecase(
		database.New{{ .Upper }}Repository(dbAdministrator),
		converter.New{{ .Upper }}Converter(),
		// その他必要に応じて追加
	)
	ctx := context.Background()

	tests := []struct {
		name    string
		args    int64
		want    *graphql.{{ .Upper }}Detail
		wantErr bool
	}{
		{
			name: "正常に取得できる",
			args: 1,
			want: &graphql.{{ .Upper }}Detail{
				ID: 1,
			},
		},
		{
			name:    "存在しないIDの場合エラーが返る",
			args:    100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uc.Fetch(ctx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test{{ .Upper }}Create(t *testing.T) {
	t.Helper()
	d := testhelper.LoadFixture(
		"../",
		"testdata/{{ .Lower }}s/fixtures/1",
	)
	dbAdministrator := db.NewDBAdministrator(d)
	uc := usecase.New{{ .Upper }}Usecase(
		database.New{{ .Upper }}Repository(dbAdministrator),
		converter.New{{ .Upper }}Converter(),
		// その他必要に応じて追加
	)
	ctx := context.Background()

	tests := []struct {
		name    string
		args    graphql.Create{{ .Upper }}Input
		want    *graphql.{{ .Upper }}Detail
	}{
		{
			name: "正常に登録できる",
			args: graphql.Create{{ .Upper }}Input{
			},
			want: &graphql.{{ .Upper }}Detail{
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uc.Create(ctx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test{{ .Upper }}Update(t *testing.T) {
	t.Helper()
	d := testhelper.LoadFixture(
		"../",
		"testdata/{{ .Lower }}s/fixtures/1",
	)
	dbAdministrator := db.NewDBAdministrator(d)
	uc := usecase.New{{ .Upper }}Usecase(
		database.New{{ .Upper }}Repository(dbAdministrator),
		converter.New{{ .Upper }}Converter(),
	)
	ctx := context.Background()

	tests := []struct {
		name    string
		args    graphql.Update{{ .Upper }}Input
		want    *graphql.{{ .Upper }}Detail
		wantErr bool
	}{
		{
			name: "正常に更新できる",
			args: graphql.Update{{ .Upper }}Input{
			},
			want: &graphql.{{ .Upper }}Detail{
			},
			wantErr: false,
		},
		{
			name: "データが存在しない場合エラーになる",
			args: graphql.Update{{ .Upper }}Input{
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uc.Update(ctx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}
			testhelper.AssertResponse(t, tt.want, got)
		})
	}
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
		err = utils.TemplateExport(m, func(name string) (*os.File, error) {
			return createTestGoFile(name)
		}, TestTemplate)
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

func createTestGoFile(name string) (*os.File, error) {
	filePath := filepath.Join("usecase", fmt.Sprintf("%v.usecase_test.go", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, nil
	}
	log.Printf("UseCaseTestを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
