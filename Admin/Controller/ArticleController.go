package Controller

import (
	"dhblog-admin/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

// QuerySingleArticle 查询指定文章
func QuerySingleArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// 观看数 + 1
	Models.UpdateArticleViewNum(id)

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

// UploadImg 上传图片
func UploadImg(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	now := time.Now()

	year := now.Year()
	month := now.Month()
	day := now.Day()

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
	dst := fmt.Sprintf("./upload/img/%d%02d%02d/%s", year, month, day, fileName)
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, dst)
}

type UpdateData struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// UpdateArticle 更新文章
func UpdateArticle(c *gin.Context) {
	var reqData UpdateData
	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		fmt.Println("更新文章出错", err)
		return
	}
	// 更新文章
	success := Models.UpdateArticle(reqData.ID, reqData.Title, reqData.Content)
	if success {
		// 更新时间
		Models.UpdateArticleUpdated(reqData.ID)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "文章更新成功",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  "文章更新失败",
	})
}

// QueryPageSize 查询页数
func QueryPageSize(c *gin.Context) {
	data := Models.SelectArticleSize()
	c.JSON(200, gin.H{
		"code": 1,
		"size": data,
	})
}
