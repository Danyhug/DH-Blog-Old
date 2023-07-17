package Databases

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"path/filepath"
)

// DB 数据库对象
var DB *gorm.DB

//var path, _ = os.Getwd()

func InitDB() {
	fullPath := filepath.Join("P:\\Code\\Vue\\dh-blog\\Admin", "db.sqlite")

	fmt.Println("数据库目录", fullPath)
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db = db.Debug()
	DB = db
}
