package utils

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

func GetRawArgs(ctx context.Context) map[string]interface{} {
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

func ConvertRawArgsToColumnNames(rawArgs map[string]interface{}, inputStruct interface{}, modelStruct interface{}) (fieldNames []string, err error) {
	if reflect.TypeOf(rawArgs["params"]).Kind() != reflect.Map {
		return fieldNames, errors.New("invalid type")
	}

	m := reflect.ValueOf(rawArgs["params"]).MapKeys()
	toType := reflect.TypeOf(modelStruct)

	for _, k := range m {
		str := k.String()
		if strings.HasPrefix(str, "<") {
			return nil, fmt.Errorf("invalid param found. Error: %s", str)
		}
		inputFieldName, err := getJsonFieldName(str, inputStruct)
		if err != nil {
			return nil, err
		}
		modelField, ok := toType.FieldByName(fieldConverter(inputFieldName))
		if !ok {
			return nil, fmt.Errorf("invalid field found. field: %s", inputFieldName)
		}
		fieldNames = append(fieldNames, modelField.Tag.Get("boil"))
	}
	if modelField, ok := toType.FieldByName("UpdatedAt"); ok {
		fieldNames = append(fieldNames, modelField.Tag.Get("boil"))
	}

	return fieldNames, nil
}

var fieldConvertMap map[string]string = map[string]string{
	"KidneyGfr":                  "KidneyGFR",
	"RiskKbn":                    "RiskKBN",
	"RiskKbnDiabetes":            "RiskKBNDiabetes",
	"RiskKbnDiabeticNephropathy": "RiskKBNDiabeticNephropathy",
	"HdlCholesterol":             "HDLCholesterol",
	"LdlCholesterol":             "LDLCholesterol",
	"GammaGt":                    "GammaGT",
}

func fieldConverter(s string) string {
	if v, ok := fieldConvertMap[s]; ok {
		return v
	}
	return s
}

func GetArgs(ctx context.Context, argName string) (interface{}, bool) {
	v, ok := graphql.GetFieldContext(ctx).Parent.Args[argName]
	fmt.Println(v, ok)
	if ok {
		return v, true
	}
	return nil, false
}
