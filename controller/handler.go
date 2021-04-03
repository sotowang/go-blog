package controller

import (
	"github.com/gin-gonic/gin"
	"go_blog/service"
	"net/http"
	"strconv"
)

//访问方面的控制器
func IndexHandle(c *gin.Context) {
	//从sevice取数据
	//1.加载文章数据
	articleRecordList, err := service.GetArticleRecordList(0, 5)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//加载分类数据
	categoryList, err := service.GetALLCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}

	//分配数据
	//var data map[string]interface{} = make(map[string]interface{})
	//data["article_list"] = articleRecordList
	//data["category_list"] = categoryList
	//c.HTML(http.StatusOK,"views/index.hmtl",data)

	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

//点击分类就进行分类
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err!=nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
	}

	//按分类id,获取分类下的文章列表
	articleRecordList ,err:= service.GetArticleListByCategoryId(int(categoryId),0,15)
	if err!=nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
	}

	//再次加载所有分类数据,用于分类去
	categoryList, err := service.GetALLCategoryList()
	if err!=nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
	}
	c.HTML(http.StatusOK,"views/index.html",gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}
