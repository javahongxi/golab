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
	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimit())
	r.Use(middleware.CircuitBreaker())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	r.GET("/ping", controller.NewHealthController().Ping)
	r.GET("/health", controller.NewHealthController().Health)

	userRepo := model.NewUserRepository(model.DB)
	userCtrl := controller.NewUserController(userRepo)
	authCtrl := controller.NewAuthController(userRepo)

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/register", authCtrl.Register)
				auth.POST("/login", authCtrl.Login)
			}

			users := v1.Group("/users")
			{
				users.GET("", userCtrl.ListUsers)
				users.GET("/:id", userCtrl.GetUser)

				users.POST("", userCtrl.CreateUser)
				users.PUT("/:id", middleware.Auth(), userCtrl.UpdateUser)
				users.DELETE("/:id", middleware.Auth(), userCtrl.DeleteUser)
			}
		}
	}

	return r
}