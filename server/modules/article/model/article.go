package model

import (
	"fmt"
	"gorm.io/gorm"
	"mang/modules/article/model/internal"
)

type Article internal.Article

type ArticleCommand struct {

}

func (m *Article) BeforeCreate(tx *gorm.DB) error{
	fmt.Println("我是测试的，我在使用前被调用了")
	return nil
}
