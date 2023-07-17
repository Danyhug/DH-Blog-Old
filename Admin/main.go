package main

import (
	"dhblog-admin/Controller"
	"dhblog-admin/Databases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	Databases.InitDB()

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 配置跨域处理
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},                   // 允许的源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的 HTTP 方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		AllowCredentials: true,                                                // 允许携带身份验证凭证（如 Cookie）
	}))

	// 用户相关接口
	//user := router.Group('/user')
	//{
	//}

	// 文章相关接口
	article := router.Group("/article")
	{
		// 新增
		article.POST("add", Controller.CreateArticle)
		// 查询
		article.POST("query", Controller.QueryArticle)
		// 查询单用户
		article.GET(":id", Controller.QuerySingleArticle)
	}
	//
	//// 评论相关接口
	//
	//v1 := router.Group("/v1")
	//v1.GET("/", func(ctx *gin.Context) { ctx.String(202, "ok") })
	//
	err := router.Run("0.0.0.0:2233")
	if err != nil {
		return
	}
}
