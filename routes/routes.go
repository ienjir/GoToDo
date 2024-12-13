package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/GoToDo/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/", controllers.CreateUser)
	}
}
