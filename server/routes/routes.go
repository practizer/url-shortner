package routes

import (
	"server/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://localhost:3000", "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public routes - no authentication required
	r.POST("/check", handlers.UrlAvailabilityChecker)
	r.POST("/url/add", handlers.AddUrl)
	r.GET("/urls", handlers.GetUserUrls)
	r.DELETE("/url/:shortcode", handlers.DeleteUrl)
	r.GET("/:shortcode", handlers.RedirectUrl)
}
