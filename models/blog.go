package models

import (
	db "github.com/cshengqun/myblog/database"
	"fmt"	
)

type Blog struct {
	Id int           `json:"id" form:"id"`
	Name string      `json:"name" form:"name" binding:"required"`
	Content string   `json:"content" form:"content" binding:"required"`
	CreatedAt string `json:"createdAt" form:"createdAt"`
}

func (blog *Blog) Create() (err error) {
	_, err = db.SqlDB.Exec("insert into blogs(name, content) values(?,?)", blog.Name, blog.Content)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (blog *Blog) Update() (err error) {
	_, err = db.SqlDB.Exec("update blogs set name=?, content=? where id=?", blog.Name, blog.Content, blog.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (blog *Blog) Read() (err error) {
	res, err := db.SqlDB.Query("select name, content, createdAt from blogs where id=?", blog.Id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Close()

	for res.Next() {
		err = res.Scan(&blog.Name, &blog.Content, &blog.CreatedAt)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return
}

func (blog *Blog) Delete() (err error) {
	_, err = db.SqlDB.Exec("delete from blogs where id=?", blog.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
