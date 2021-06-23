package generate_code

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
	Name      string `json:"name"`
	Default   string `json:"default"`
	IsNull    string `json:"is_null"`
	MaxLength int64  `json:"max_length"`
	Precision int64  `json:"precision"`
	Type      string `json:"type"`
	Key       string `json:"key"`
	Extra     string `json:"extra"`
	Comment   string `json:"comment"`
}

func (*generateCode) Model(dbName string, tableName string) {

}

func (*generateCode) GetColumns(dbName string, tableName string) (err error, columns []*Column) {
	fields := "COLUMN_NAME name,COLUMN_DEFAULT default,IS_NULLABLE is_null,CHARACTER_MAXIMUM_LENGTH max_length,NUMERIC_PRECISION precision, COLUMN_TYPE type,COLUMN_KEY key,EXTRA extra,COLUMN_COMMENT comment"
	err = tiger.Db(dbName).Table(tableName).Select(fields).Scan(&columns).Error
	return err, columns
}

func parseTemplate(filename string, data interface{}) {
	t,_:=template.ParseFiles("aa.tpl")
	err :=t.Execute(os.Stdout, data)
	fmt.Println(err)
}
