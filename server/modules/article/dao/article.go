package dao

import (
	"gorm.io/gorm"
	"tiger-go/modules/article/dao/internal"
	"tiger-go/tiger"
	"tiger-go/tiger/db"
)

type articleDao struct {
	internal.ArticleDao
}

//var (
//	// Address is globally public accessible object for table address operations.
//	Article = &articleDao{
//		internal.Article,
//	}
//)

type dao struct {
	*gorm.DB
	PrimaryKey string
}

//func newDao() *gorm.DB {
//	dao := &dao{
//		tiger.Db().Table("article"),
//	}
//	//return dao
//}

var Article = &db.DB{
	DB:         tiger.Db().Table("article").Session(&gorm.Session{}),
	PrimaryKey: "article_id",
}


//var Article = &articleDao{
//	internal.Article,
//}
