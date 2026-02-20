package routes

import (
	"server/handlers"
	"server/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := r.Group("/auth")
	{
		auth.POST("/login", handlers.GoogleAuth)
		auth.POST("/logout", handlers.Logout)
	}

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/check", handlers.UrlAvailabilityChecker)
		protected.POST("/url/add", handlers.AddUrl)
	}
	r.GET("/:shortcode", handlers.RedirectUrl)
}
