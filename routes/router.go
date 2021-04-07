package routes

import (
	"blog/midware"
	"blog/api/v1"
	"blog/utils"
	"github.com/gin-gonic/gin"
	_"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(midware.JwtToken())
	{
		// User model routeinterface
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// Article model route interface
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		// Category model route interface
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// Upload file
		auth.POST("upload", v1.Upload)
	}

	rV1 := r.Group("api/v1")
	{
		// User model routeinterface
		rV1.POST("user/add", v1.AddUser)
		rV1.GET("users", v1.GetUsers)
		// Article model route interface
		rV1.GET("article/:id", v1.GetArticle)
		rV1.GET("articles", v1.GetArticles)
		rV1.GET("articles/catelist/:id", v1.GetCategoryArticles)
		// Category model route interface
		rV1.GET("categorys", v1.GetCategorys)
		// Login
		rV1.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}