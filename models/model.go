package models

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
)

type Blog struct {
	Id      int `PK`
	Title   string
	Content string
	Ctime   time.Time
	Mtime   time.Time
}

func GetLink() beedb.Model {
	db, err := sql.Open("mymysql", "test/testop/testop")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}

func GetAll() (blogs []Blog) {
	db := GetLink()
	db.FindAll(&blogs)
	return
}

func GetBlog(id int) (blog Blog) {
	db := GetLink()
	db.Where("id=?", id).Find(&blog)
	return
}

func UpdateBlog(blog Blog) (id int64, err error) {
	db := GetLink()
	props := make(map[string]interface{}, 0)
	props["title"] = blog.Title
	props["content"] = blog.Content
	id, err = db.Update(props)
	return
}

func SaveBlog(blog Blog) (bg Blog) {
	db := GetLink()
	db.Save(&blog)
	return bg
}

func DelBlog(blog Blog) {
	db := GetLink()
	db.Delete(&blog)
	return
}
