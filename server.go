package main

import (
	"os"

	"github.com/dkkundu/go_api/router"

	"github.com/gin-gonic/gin"
)

func init() {
	os.Setenv("IP", "0.0.0.0")
	os.Setenv("PORT", "8000")
}

func main() {
	server := gin.Default()
	// Initialize router groups
	v1 := server.Group("/api")
	{
		// Add ping, login, and logout routes
		v1.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		v1.POST("/login", func(c *gin.Context) {
			c.String(200, "login")
		})
		v1.POST("/logout", func(c *gin.Context) {
			c.String(200, "logout")
		})

		// Initialize API routes for version 1
		api := v1.Group("/v1")
		{
			router.RoutesVersion1(api)
		}
	}

	server.Run()
}
