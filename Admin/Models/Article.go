package Models

import (
	"dhblog-admin/Databases"
	"fmt"
	"time"
)

type Article struct {
	ID      int    `gorm:"primaryKey;autoIncrement"`
	Title   string `gorm:"size:30;not null"`
	Content string `gorm:"type:text;not null"`
	Created int64  `gorm:"autoCreateTime"`
	Updated int64  `gorm:"autoCreateTime"`
	Viewnum int64  ``
}

// InsertArticle 插入文章
func InsertArticle(title string, content string) bool {
	article := &Article{
		Title:   title,
		Content: content,
		Created: time.Now().Unix(),
		Updated: time.Now().Unix(),
		Viewnum: 0,
	}
	result := Databases.DB.Create(&article)
	if result.Error == nil {
		return true
	}
	fmt.Println("插入数据失败", result.Error)
	return false
}

// SelectArticle 查询文章
func SelectArticle(start int, limit int) ([]Article, bool) {
	var article []Article
	result := Databases.DB.Offset(start).Limit(limit).Find(&article)
	for _, value := range article {
		fmt.Println("查询到的数据为", value)
	}
	fmt.Print("查询所有数据", result)
	return article, true
}

// UpdateArticle 更新文章数据
func UpdateArticle(id int, title string, content string) bool {
	article := &Article{
		Title:   title,
		Content: content,
	}

	res := Databases.DB.Find(&Article{ID: id}).Updates(&article)
	if res.Error == nil {
		return true
	}
	fmt.Println("文章更新出错", res.Error)
	return false
}

// UpdateArticleUpdated 更改文章更新时间
func UpdateArticleUpdated(id int) bool {
	res := Databases.DB.Find(&Article{ID: id}).Updates(&Article{Updated: time.Now().Unix()})
	if res.Error == nil {
		return true
	}
	fmt.Println("更改文章更新时间出错", res.Error)
	return false
}

// UpdateArticleViewNum 文章观看数 + 1
func UpdateArticleViewNum(id int) bool {
	var num int
	Databases.DB.Model(&Article{}).Where(&Article{ID: id}).Select("viewnum").Scan(&num)

	// 次数加一
	num++

	res := Databases.DB.Find(&Article{ID: id}).Update("viewnum", num)
	if res.Error == nil {
		return true
	}
	fmt.Println("文章观看数新增出错", res.Error)
	return false
}

// FindArticle 查询指定文章
func FindArticle(id int) (Article, bool) {
	var article Article
	result := Databases.DB.Where(&Article{ID: id}).Find(&article)
	if result.Error == nil {
		return article, true
	}
	fmt.Println("查询指定文章出错", result.Error)
	return article, false
}

// SelectArticleSize 查询总页数
func SelectArticleSize() int64 {
	var article Article
	var count int64
	Databases.DB.Model(&article).Count(&count)
	return count
}
