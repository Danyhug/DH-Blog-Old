package Controller

import (
	"dhblog-admin/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CreateData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// CreateArticle 创建新文章
func CreateArticle(c *gin.Context) {
	var reqData CreateData
	c.ShouldBindJSON(&reqData)

	status := Models.InsertArticle(reqData.Title, reqData.Content)
	if status {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "文章已成功添加",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": -1,
		"msg":  "文章添加失败",
	})
}

type QueryData struct {
	Start int `json:"start"`
	Limit int `json:"limit"`
}

// QueryArticle 查询文章
func QueryArticle(c *gin.Context) {
	var reqData QueryData
	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		fmt.Println("出错", err)
	}
	start := reqData.Start
	limit := reqData.Limit
	fmt.Println("测试", start, limit)

	data, status := Models.SelectArticle(start, limit)
	if status {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "文章查询成功",
			"data": data,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "文章查询失败",
		"data": "",
	})
}

// 查询指定文章
func QuerySingleArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, status := Models.FindArticle(id)
	if status {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "该文章获取成功",
			"data": data,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": -1,
		"msg":  "该文章获取失败",
		"data": "",
	})
}
