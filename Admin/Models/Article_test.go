package Models

import (
	"dhblog-admin/Databases"
	"testing"
)

func TestCreate(t *testing.T) {
	Databases.InitDB()

	title := "哇哈哈"
	content := "我也是"
	success := InsertArticle(title, content)
	if !success {
		t.Error("错误")
	}
}

func TestQuery(t *testing.T) {
	Databases.InitDB()
	success := SelectArticle(0, 10)
	if !success {
		t.Error("错误")
	}
}
