package internal

import (
	"gorm.io/gorm"
	"tiger-go/tiger"
)

type ArticleDao struct {
	*gorm.DB
	Table      string
	PrimaryKey string
	Columns    columns
}

var Article = ArticleDao{
	tiger.Db().Table("article").Session(&gorm.Session{}),
	"article",
	"article_id",
	columns{
		ArticleId: "article_id",
	},
}

type columns struct {
	ArticleId string
}
