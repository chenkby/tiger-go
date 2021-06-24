package define

// 创建请求参数
type GenerateCodeCreateReq struct {
	DbName    string `json:"db_name" form:"db_name" default:"default"`
	TableName string `json:"table_name" form:"table_name"`
}
