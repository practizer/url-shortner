package routes

import(
	"github.com/gin-gonic/gin"
	"server/handlers"
)

func Routes(r *gin.Engine){

	auth := r.Group("/auth")
	{
		auth.POST("/login",handlers.GoogleAuth)
		auth.POST("/logout",handlers.Logout)
	}
}