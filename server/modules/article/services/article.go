package services

import (
	"tiger-go/modules/article/dao"
)

var Article = new(ArticleService)

type ArticleService struct{}

//func (s *articleService) List() {
//	fmt.Println("service.article.list")
//}

//func (s *articleService) Create(ctx context.Context, req *define.ArticleCreateReq) {
//
//	fmt.Printf("req2:%p", req, &req)
//	//err := dao.Article.Insert(req)
//	//
//	//fmt.Println("err", err, req)
//
//	//err := dao.Article.Delete(nil, "article_id=11")
//	//fmt.Println("删除：", err)
//	//result := ([]*model.Article)(nil)
//	//err := dao.Article.Select("article_id,content").Order("article_id desc").Limit(5).All(&result)
//	//fmt.Println("query error:", err, result)
//	//if err == nil {
//	//	fmt.Println(result[0])
//	//}
//	//err = dao.Article.DB.Transaction(func(tx *gorm.DB) error {
//	//	return nil
//	//})
//}

//func (s *articleService) Update(ctx context.Context, req *define.ArticleUpdateReq) (*define.ArticleUpdateRep, error) {
//	return nil, nil
//}
//
//func (s *articleService) View(ctx context.Context, req *define.ArticleViewReq) (*define.ArticleViewRep, error) {
//
//	var result *define.ArticleViewRep
//	err := dao.Article.First(&result).Error
//	if err != nil {
//		return nil, err
//	} else {
//		return result, nil
//	}
//}

func (s *ArticleService) FindOne(dest interface{}, condition interface{}) error {

	err := dao.Article.Find(dest, condition).Error
	return err
}
