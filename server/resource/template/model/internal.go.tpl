package internal

type Article struct {
{{- range .columns}}
	{{.Name}} int    `gorm:"primaryKey" json:"article_id"`
{{- end}}
	Title     string `json:"title"`
	Content   string `json:"content"`
	Template  int    `json:"template"`
	ViewNum   int    `json:"view_num"`
	Status    int    `json:"status"`
	Rank      int    `gorm:"default:1" json:"rank"`
	AllowDel  int    `json:"allow_del"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}
