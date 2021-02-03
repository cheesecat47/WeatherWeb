package v1

import (
	"github.com/cheesecat47/webpractice/api/controller"
	"github.com/cheesecat47/webpractice/api/middleware"
	"github.com/gin-gonic/gin"
)

//InitRoutes func is for initializing api server routes
func InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	//regular group: non-authorization
	regular := api.Group("/public")
	{
		regular.POST("/login", controller.Login)
	}

	//restricted group: need authorization
	restricted := api.Group("/protected")
	restricted.Use(middleware.TokenAuthMiddleware())
	{
		restricted.POST("/user:id", controller.DeleteUser)
	}
	return router

}
