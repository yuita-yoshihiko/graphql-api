package utils

import (
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

// name 型の定義
type Name string

// CamelCase に変換
func (n Name) Upper() string {
	return strcase.ToCamel(string(n))
}

// lowerCamelCase に変換
func (n Name) Lower() string {
	return strcase.ToLowerCamel(string(n))
}

// 環境変数 SKELETON_NAMES を取得して Name のスライスを返す
func GetNameList() []Name {
	p := os.Getenv("SKELETON_NAMES")
	if p == "" {
		log.Fatal("環境変数 SKELETON_NAMES が設定されていません")
		return nil
	}
	names := strings.Split(p, ",")
	var list []Name
	for _, n := range names {
		n = strings.TrimSpace(n)
		if n != "" {
			list = append(list, Name(n))
		}
	}
	return list
}

func TemplateExport(m Name, createFileFunc func(string) (*os.File, error), tmpl string) error {
	tpl, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}

	file, err := createFileFunc(m.Lower())
	if err != nil {
		return err
	} else if file == nil {
		return nil
	}
	defer file.Close()

	return tpl.Execute(file, m)
}
