package router

import (
	"github.com/gin-gonic/gin"
	"techtestify/internal/auth"
	"techtestify/internal/middleware"
	"techtestify/internal/quiz"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())

	protected.GET("/profile", func(c *gin.Context) {
		email := c.GetString("email")
		role := c.GetString("role")
		c.JSON(200, gin.H{"email": email, "role": role})
	})
	protected.GET("/tests", quiz.GetAllTests)
	admin := protected.Group("/admin")
	admin.Use(middleware.RequireRole("admin"))
	admin.POST("/tests", quiz.CreateTest)
	admin.POST("/tests/:id/questions", quiz.AddQuestion)
	admin.GET("/dashboard", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Admin!"})
	})

	return r
}
