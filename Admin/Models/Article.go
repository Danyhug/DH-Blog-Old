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
}

// InsertArticle 插入文章
func InsertArticle(title string, content string) bool {
	article := &Article{
		Title:   title,
		Content: content,
		Created: time.Now().Unix(),
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
