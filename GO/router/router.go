package router

import (
	"GO/controllers"
	"GO/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    userRoutes := r.Group("/users")
    {
        userRoutes.POST("/register", controllers.RegisterUser)
        userRoutes.POST("/login", controllers.LoginUser)
        userRoutes.PUT("/:userId", controllers.UpdateUser)
        userRoutes.DELETE("/:userId", controllers.DeleteUser)
    }

    photoRoutes := r.Group("/photos")
    {
        photoRoutes.Use(middlewares.AuthMiddleware())
        photoRoutes.POST("/", controllers.CreatePhoto)
        photoRoutes.GET("/", controllers.GetPhotos)
        photoRoutes.PUT("/:photoId", controllers.UpdatePhoto)
        photoRoutes.DELETE("/:photoId", controllers.DeletePhoto)
    }

    return r
}
