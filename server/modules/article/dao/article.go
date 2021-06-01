package dao

import (
	"mang/modules/article/dao/internal"
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

var Article = &articleDao{
	internal.Article,
}
