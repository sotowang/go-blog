package service

import (
	"fmt"
	"go_blog/dao/db"
	"go_blog/model"
)

// 获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleInfoList, err := db.GetAricleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	continueDemo()
	// 2.获取文章对应的分类（多个）
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	// 返回页面，做聚合
	// 遍历所有文章
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{ // 根据当前文章，生成结构体
			ArticleInfo: *article,
		}
		// 文章取出分类id
		categoryId := article.CategoryId
		// 遍历分类列表
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList,articleRecord)
	}
	return
}
func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		fmt.Println(111111111111)
	}
}
// 根据多个文章的id，获取多个分类id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABLE:
	// 遍历文章，得到每个文章
	for _, article := range articleInfoList {
		// 从当前文章取出分类id
		categoryId := article.CategoryId
		// 去重，防止重复
		for _, id := range ids {
			// 看当前id是否存在
			if id == categoryId {
				continue LABLE //结束当前article的执行,继续下一个article的执行,
			}
		}
		ids = append(ids,categoryId)
	}
	return
}

// 	根据分类id，获取该类文章和他们对应的分类信息
func GetArticleListByCategoryId(categoryId,pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId,pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	// 2.获取文章对应的分类
	categoryList, err := db.GetCategoryById(int64(categoryId))
	if err != nil {
		return
	}
	// 返回页面，做聚合
	// 遍历所有文章
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{ // 根据当前文章，生成结构体
			ArticleInfo: *article,
			Category : *categoryList,
		}
		articleRecordList = append(articleRecordList,articleRecord)
	}
	return
}
