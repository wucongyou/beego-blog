package models

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
)

type Blog struct {
	Id      int
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
	orm := GetLink()
	orm.OrderBy("mtime desc").FindAll(&blogs)
	return
}

func GetBlog(id int) (blog Blog) {
	orm := GetLink()
	orm.Where("id=?", id).Find(&blog)
	return
}

func UpdateBlog(blog Blog) (id int64, err error) {
	orm := GetLink()
	props := make(map[string]interface{}, 0)
	props["title"] = blog.Title
	props["content"] = blog.Content
	orm.SetTable("blog").Where("id=?", blog.Id).Update(props)
	return
}

func SaveBlog(blog Blog) (bg Blog) {
	orm := GetLink()
	orm.Save(&blog)
	return bg
}

func DelBlog(blog Blog) {
	orm := GetLink()
	orm.Delete(&blog)
	orm.SetTable("blog").Where("id=?", blog.Id).DeleteRow()
	return
}
