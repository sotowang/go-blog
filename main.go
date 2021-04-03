package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/controller"
	"go_blog/dao/db"
)

func main() {
	router := gin.Default()
	dns := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	//加载静态文件
	router.Static("/static/","./static")

	//加载模板
	router.LoadHTMLGlob("views/*")
	router.GET("/",controller.IndexHandle)  //首页
	router.GET("/category/",controller.CategoryList) //分类列表页面
	_=router.Run(":8001")
}
