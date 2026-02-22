package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimit(rdb *redis.Client) gin.HandlerFunc {
	return func (c *gin.Context){
		limit := 5;
		userIDvalue, exists := c.Get("user_id");
		if !exists {
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Unauthorized"})
			return 
		}
		userID := userIDvalue.(int)

		today := time.Now().UTC().Format("2006-01-02")
		key := fmt.Sprintf("ratelimit:%d:%s", userID, today)
		ctx := context.Background()

		count,err := rdb.Incr(ctx,key).Result()
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Server Error"})
			c.Abort()
			return 
		}
		if count == 1 {
			now := time.Now().UTC()
			midnight := time.Date(now.Year(),now.Month(),now.Day()+1,0,0,0,0,time.UTC)
			ttl := time.Until(midnight)
			rdb.Expire(ctx,key,ttl)
		}

		remaining := int64(limit) - count
		if remaining < 0 {
			remaining = 0
		}
		
		// Informational headers
		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", limit))
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))

		if count > int64(limit){
			c.JSON(http.StatusTooManyRequests,gin.H{
				"error":"Daily limit reached",
			})
			c.Abort()
			return 
		}
		c.Next()
	}
}