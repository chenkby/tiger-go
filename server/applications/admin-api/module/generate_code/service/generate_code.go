package service

import (
	"fmt"
	"html/template"
	"os"
	"tiger-go/tiger"
)

var GenerateCode = new(generateCode)

type generateCode struct {
}
type Column struct {
	Name             string `json:"name"`
	DefaultValue     string `json:"default_value"`
	IsNull           string `json:"is_null"`
	MaxLength        int64  `json:"max_length"`
	NumericPrecision int64  `json:"numeric_precision"`
	Type             string `json:"type"`
	ColumnKey        string `json:"column_key"`
	Extra            string `json:"extra"`
	Comment          string `json:"comment"`
}

func (s *generateCode) Model(dbName string, tableName string) (string, error) {
	_, columns := s.GetColumns(dbName, tableName)
	s.ParseTemplate("test", map[string]interface{}{
		"columns": columns,
	})
	return "", nil
}

func (*generateCode) GetColumns(dbName string, tableName string) (err error, columns []*Column) {
	fields := "COLUMN_NAME name,COLUMN_DEFAULT default_value,IS_NULLABLE is_null,CHARACTER_MAXIMUM_LENGTH max_length,NUMERIC_PRECISION numeric_precision, COLUMN_TYPE type,COLUMN_KEY column_key,EXTRA extra,COLUMN_COMMENT comment"
	err = tiger.Db().Raw(fmt.Sprintf("select %s from information_schema.COLUMNS where TABLE_SCHEMA='%s' AND TABLE_NAME='%s'", fields, dbName, tableName)).Scan(&columns).Error
	return err, columns
}

func (s *generateCode) ParseTemplate(filename string, data interface{}) {
	dir, _ := os.Getwd()
	t, err := template.ParseFiles(dir + "/resource/template/model/internal.go.tpl")

	if err != nil {
		p, _ := os.Getwd()
		fmt.Println("发生错误：", err, p)
	} else {
		err := t.Execute(os.Stdout, data)
		fmt.Println(err)
	}
}
