package main
import (
	"github.com/gin-gonic/gin"
	. "github.com/cshengqun/myblog/apis"
	. "github.com/cshengqun/myblog/env"
)

func initRouter() *gin.Engine {
//	gin.DisableConsoleColor()
//	gin.DefaultWriter = Env.Logger
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*")
	router.Static("/static/", "../static/")
	router.GET("/", Index)
	router.GET("/blog/read/:id", GetReadBlog)
	router.GET("/page/read/:idx", ReadPage)

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		Env.Conf.Account: Env.Conf.Password,
	}))
	authorized.GET("/", Admin)
	authorized.GET("/blog/create", GetCreateBlog)
	authorized.POST("/blog/create", PostCreateBlog)
	authorized.GET("/blog/update/:id", GetUpdateBlog)
	authorized.POST("/blog/update/:id", PostUpdateBlog)
	authorized.GET("/blog/delete/:id", DeleteBlog)

	return router
}

func main() {
	router := initRouter()
	router.Run(Env.Conf.Ip + ":" + Env.Conf.Port)
}
