package routes

import (
	"server/handlers"
	"server/middlewares"
	"server/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://bit-urls.vercel.app"},
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
		auth.GET("/me", middlewares.AuthMiddleware(), handlers.Me)
	}

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/check", handlers.UrlAvailabilityChecker)
		protected.POST("/url/add",middlewares.RateLimit(config.RDB), handlers.AddUrl)
		protected.GET("/urls", handlers.GetUserUrls)
		protected.DELETE("/url/:shortcode", handlers.DeleteUrl)
	}
	r.GET("/:shortcode", handlers.RedirectUrl)
}
