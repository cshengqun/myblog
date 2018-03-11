package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "github.com/cshengqun/myblog/models"
	"strconv"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
)

func Index(c *gin.Context) {
	var page Page
	page.Size = 2
	page.Idx = 0
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

func ReadPage(c *gin.Context) {
	var page Page
	page.Size = 2
	page.Idx, _ = strconv.Atoi(c.Param("idx"))
	page.Idx--
	if page.Idx < 0 {
		page.Idx = 0
	}
	if err := page.Get(); err == nil {
		if len(page.Blogs) == 0 {
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusOK, "page.html", gin.H{
				"page": page,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
}

func GetCreateBlog(c *gin.Context) {
	c.HTML(http.StatusOK, "pageCreate.html", gin.H{

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

func UpdateBlog(c *gin.Context) {
	var blog Blog
	if err := c.ShouldBind(&blog); err == nil {
		if err = blog.Update(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"retCode": 0,
				"retMsg": "update successfully",
			})
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

func ReadBlog(c *gin.Context) {
	var blog Blog	
	var err error
	blog.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"retCode": -1,
			"retMsg": err.Error(),
		})
	}
	if err := blog.Read(); err == nil {
		content := template.HTML(blackfriday.Run([]byte(blog.Content)))
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
