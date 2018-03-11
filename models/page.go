package models

import (
	db "github.com/cshengqun/myblog/database"
	"fmt"
)

type Page struct {
	Size int
	Idx int
	Blogs []Blog
}

func (page *Page) Get() (err error) {
	res, err := db.SqlDB.Query("SELECT * FROM blogs ORDER BY createdAt DESC LIMIT ?,?", page.Idx * page.Size, page.Size)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Close()

	for res.Next() {
		var blog Blog
		err = res.Scan(&blog.Id, &blog.Name, &blog.Content, &blog.CreatedAt)
		if err != nil {
			panic(err)
		}
		page.Blogs = append(page.Blogs, blog)
	}
	return
}
