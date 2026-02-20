package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		tokenString,err := c.Cookie("auth_token")
		if err!=nil{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Authentication required"})
			c.Abort()
			return
		}
		token,err := jwt.Parse(tokenString,func(token *jwt.Token)(interface{},error){
			return []byte(os.Getenv("JWT_SECRET")),nil
		})

		if err!=nil || !token.Valid {
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid Token Claims"})
			c.Abort()
			return 
		}
		claims,ok := token.Claims.(jwt.MapClaims)
		if !ok{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid token claims"})
			c.Abort()
			return 
		}
		userIDFloat,ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid User ID"})
			c.Abort()
			return 
		}
		userID := int(userIDFloat)

		c.Set("user_id",userID)
		c.Next()
	}
}