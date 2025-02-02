package utils

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

func GetGraphQLFields(ctx context.Context) map[string]interface{} {
	return graphql.GetFieldContext(ctx).Field.ArgumentMap(graphql.GetOperationContext(ctx).Variables)
}

func getJsonFieldName(tag string, s interface{}) (string, error) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		return "", errors.New("invalid type")
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get("json"), ",")[0]
		if v == tag {
			return f.Name, nil
		}
	}
	return "", fmt.Errorf("cannot find tag: %s", tag)
}

func ConvertUpdateInputToDBColumnNames(m, g interface{}) []string {
	var columns []string
	mrt := reflect.TypeOf(m)
	grt := reflect.TypeOf(g)
	for i := 0; i < mrt.NumField(); i++ {
		mf := mrt.Field(i)
		if mf.Name == "UpdatedAt" {
			columns = append(columns, "updated_at")
		}
		for i := 0; i < grt.NumField(); i++ {
			gf := grt.Field(i)
			if mf.Name == gf.Name {
				columns = append(columns, mf.Tag.Get(`boil`))
				break
			}
		}
	}
	return columns
}
