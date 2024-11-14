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

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAddInterviewStatusMemoCreate(t *testing.T) {
	t.Helper()
	d := testhelper.LoadFixture(
		"../",
		"testdata/users/fixtures/1",
	)
	dbAdministrator := db.NewDBAdministrator(d)
	uc := usecase.NewUserUsecase(
		database.NewUserRepository(dbAdministrator),
		converter.NewUserConverter(),
	)
	ctx := context.Background()

	tests := []struct {
		name    string
		args    graphql.CreateUserInput
		want    *graphql.UserDetail
		wantErr bool
	}{
		{
			name: "正常に登録できる",
			args: graphql.CreateUserInput{
				Name: "テスト",
			},
			want: &graphql.UserDetail{
				Name: "テスト",
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
			opt := cmp.Options{
				cmpopts.IgnoreFields(
					graphql.UserDetail{},
					"ID",
					"Name",
				),
			}
			if diff := cmp.Diff(tt.want, got, opt); diff != "" {
				t.Errorf("diff(-got +want)\n%s", diff)
			}
		})
	}
}
