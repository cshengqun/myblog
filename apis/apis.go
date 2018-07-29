package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "github.com/cshengqun/myblog/models"
	"strconv"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	. "github.com/cshengqun/myblog/env"
)

func Index(c *gin.Context) {
	var page Page
	page.Size = Env.Conf.PageSize
	page.Idx = 0
	page.PreIdx = page.Idx - 1
	page.NextIdx = page.Idx + 1
	Env.Report("myblog", "123", "456", "START_CALL", "{\"code\":\"1001\",\"value\":\"hello\"}", "{\"msg\":\"world\"}")
	if err := page.Get(); err == nil {
		c.HTML(http.StatusOK, "page.html", gin.H{
			"page": page,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
}

func Admin(c *gin.Context) {
	var page Page
	page.Size = Env.Conf.PageSize
	page.Idx = 0
	page.PreIdx = page.Idx - 1
	page.NextIdx = page.Idx + 1
	if err := page.Get(); err == nil {
		c.HTML(http.StatusOK, "admin.html", gin.H{
				"page":page,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
}

func ReadPage(c *gin.Context) {
	var page Page
	page.Size = Env.Conf.PageSize
	page.Idx, _ = strconv.Atoi(c.Param("idx"))
	if page.Idx < 0 {
		page.Idx = 0
	}
	page.PreIdx = page.Idx - 1
	page.NextIdx = page.Idx + 1
	if err := page.Get(); err == nil {
		if len(page.Blogs) == 0 {
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusOK, "page.html", gin.H{
				"page": page,
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
}

func GetCreateBlog(c *gin.Context) {
	c.HTML(http.StatusOK, "createBlog.html", gin.H{

	})
}

func PostCreateBlog(c *gin.Context) {
	var blog Blog
	if err := c.ShouldBind(&blog); err == nil {
		if err = blog.Create(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"retCode": 0,
				"retMsg": "create blog successfully",
			})
		} else {
			c.JSON(http.StatusServiceUnavailable, gin.H{
					"retCode": -1,
					"retMsg": err.Error(),
			})
		}
		
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
}

func PostUpdateBlog(c *gin.Context) {
	var blog Blog
	var err error
	blog.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})	
		return
	}
	if err := c.ShouldBind(&blog); err == nil {
		Env.Logger.Info("id:%d name:%s content:%s", blog.Id, blog.Name, blog.Content)	
		if err = blog.Update(); err == nil {
			/*
			c.JSON(http.StatusOK, gin.H{
				"retCode": 0,
				"retMsg": "update successfully",
			})
			*/
			GetReadBlog(c);	
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"retCode": -1,
				"retMsg": err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
}

func GetReadBlog(c *gin.Context) {
	var blog Blog	
	var err error
	blog.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
	if rows, err := blog.Read(); err == nil && rows == 1 {
		content := template.HTML(blackfriday.Run([]byte(blog.Content), blackfriday.WithExtensions(blackfriday.HardLineBreak)))
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"blog": blog,
			"content": content,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
}

func GetUpdateBlog(c *gin.Context) {
	var blog Blog
	var err error
	blog.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
	if cnt, err := blog.Read(); err == nil && cnt == 1 {
	//	content := template.HTML(blackfriday.Run([]byte(blog.Content)))
		c.HTML(http.StatusOK, "updateBlog.html", gin.H{
			"blog": blog,
			"content": blog.Content,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
}


func DeleteBlog(c *gin.Context) {
	var blog Blog
	var err error
	blog.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
	if err := blog.Delete();err == nil {
		c.JSON(http.StatusOK, gin.H{
			"retCode": 0,
			"retMsg": "delete successfully",
		})	
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
}
