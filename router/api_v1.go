package router

import (
	"github.com/dkkundu/go_api/user"
	"github.com/gin-gonic/gin"
)

func RoutesVersion1(api *gin.RouterGroup) {
	{
		api.GET("/users", user.GetUsers)
		api.GET("/users/:id", user.GetUserByID)
		api.POST("/users", user.CreateUser)
		api.PUT("/users/:id", user.UpdateUser)
		api.DELETE("/users/:id", user.DeleteUser)
	}
}
