package main

import (
	"jwt-auth/config"
	"jwt-auth/handler"
	"jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.AutoMigrateDataBase()

	r := gin.Default()

	apiGroup := r.Group("/api")

	authHandler := handler.NewAuthHandler()
	apiGroup.POST("/register", authHandler.RegisterHandler)
	apiGroup.POST("/login", authHandler.LoginHandler)

	protected := r.Group("/api/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/user", authHandler.CurrentUser)

	r.Run(":5000")

}
