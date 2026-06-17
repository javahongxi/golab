package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/javahongxi/golab/gin/controller"
	"github.com/javahongxi/golab/gin/middleware"
	"github.com/javahongxi/golab/gin/model"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	r.GET("/ping", controller.NewHealthController().Ping)
	r.GET("/health", controller.NewHealthController().Health)

	userRepo := model.NewUserRepository(model.DB)
	userCtrl := controller.NewUserController(userRepo)

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.POST("", userCtrl.CreateUser)
				users.GET("", userCtrl.ListUsers)
				users.GET("/:id", userCtrl.GetUser)
				users.PUT("/:id", userCtrl.UpdateUser)
				users.DELETE("/:id", userCtrl.DeleteUser)
			}
		}
	}

	return r
}