package routes

import (
	"post_blog/config"
	"post_blog/controllers"
	"post_blog/repositories"
	"post_blog/services"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.RouterGroup) {
	repo := repositories.NewPostRepo(config.DB)
	service := services.NewPostService(repo)
	ctrl := controllers.NewPostController(service)

	post := router.Group("/posts")
	{
		post.POST("", ctrl.Create)
		post.GET("", ctrl.List)
		post.GET("/:id", ctrl.Get)
		post.PATCH("/:id", ctrl.Update)
		post.DELETE("/:id", ctrl.Delete)
	}
}
