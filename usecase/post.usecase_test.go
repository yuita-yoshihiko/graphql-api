package usecase_test

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

func TestPostFetch(t *testing.T) {
	t.Helper()
	d := testhelper.LoadFixture(
		"../",
		"testdata/posts/fixtures/1",
	)
	dbAdministrator := db.NewDBAdministrator(d)
	uc := usecase.NewPostUsecase(
		database.NewPostRepository(dbAdministrator),
		converter.NewPostConverter(),
	)
	ctx := context.Background()

	tests := []struct {
		name    string
		args    int64
		want    *graphql.PostDetail
		wantErr bool
	}{
		{
			name: "正常に取得できる",
			args: 1,
			want: &graphql.PostDetail{
				ID: 1,
				User: &graphql.UserDetail{
					ID: 1,
				},
				Title:   "テスト",
				Content: "テスト",
			},
			wantErr: false,
		},
		{
			name:    "存在しないID",
			args:    100,
			want:    nil,
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

func TestPostCreate(t *testing.T) {
	t.Helper()
	d := testhelper.LoadFixture(
		"../",
		"testdata/posts/fixtures/1",
	)
	dbAdministrator := db.NewDBAdministrator(d)
	uc := usecase.NewPostUsecase(
		database.NewPostRepository(dbAdministrator),
		converter.NewPostConverter(),
	)
	ctx := context.Background()

	tests := []struct {
		name    string
		args    graphql.CreatePostInput
		want    *graphql.PostDetail
		wantErr bool
	}{
		{
			name: "正常に登録できる",
			args: graphql.CreatePostInput{
				UserID:  1,
				Title:   "テスト",
				Content: "テスト",
			},
			want: &graphql.PostDetail{
				ID: 10001,
				User: &graphql.UserDetail{
					ID: 1,
				},
				Title:   "テスト",
				Content: "テスト",
			},
			wantErr: false,
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

func TestPostUpdate(t *testing.T) {
	t.Helper()
	d := testhelper.LoadFixture(
		"../",
		"testdata/posts/fixtures/1",
	)
	dbAdministrator := db.NewDBAdministrator(d)
	uc := usecase.NewPostUsecase(
		database.NewPostRepository(dbAdministrator),
		converter.NewPostConverter(),
	)
	ctx := context.Background()

	tests := []struct {
		name    string
		args    graphql.UpdatePostInput
		want    *graphql.PostDetail
		wantErr bool
	}{
		{
			name: "正常に更新できる",
			args: graphql.UpdatePostInput{
				ID:      1,
				Title:   "更新テスト",
				Content: "更新テスト",
			},
			want: &graphql.PostDetail{
				ID: 1,
				User: &graphql.UserDetail{
					ID: 1,
				},
				Title:   "更新テスト",
				Content: "更新テスト",
			},
			wantErr: false,
		},
		{
			name: "データが存在しない場合エラーになる",
			args: graphql.UpdatePostInput{
				ID:      2,
				Title:   "更新テスト",
				Content: "更新テスト",
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
