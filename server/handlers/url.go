package handlers

import (
	"database/sql"
	"net/http"
	"net/url"
	"server/config"
	"server/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UrlAvailabilityChecker(c *gin.Context) {

	var input models.Url

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Invalid request body",
		})
		return
	}

	var shortCode string

	err := config.DB.QueryRow("SELECT short_code FROM urls WHERE short_code = ?", input.ShortCode).Scan(&shortCode)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Short Code Already Exists",
		})
		return
	}

	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Short Code Available",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"status": false,
		"error":  "Database error",
	})
}

func AddUrl(c *gin.Context) {
	var input models.UrlRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	// Validate and normalize URL
	originalUrl := strings.TrimSpace(input.OriginalUrl)
	if originalUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL cannot be empty"})
		return
	}

	// Add scheme if missing
	if !strings.HasPrefix(originalUrl, "http://") && !strings.HasPrefix(originalUrl, "https://") {
		originalUrl = "https://" + originalUrl
	}

	// Validate URL format
	_, err := url.ParseRequestURI(originalUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}

	userIDvalue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDvalue.(int)

	_, err = config.DB.Exec("INSERT INTO urls (user_id,short_code,original_url,created_at,expires_at) VALUES (?,?,?,?,?)", userID, input.ShortCode, originalUrl, time.Now(), time.Now().Add(48*time.Hour))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "URL added Successfully",
	})
}
func DeleteUrl(c *gin.Context) {
	userIDvalue, exists := c.Get("user_id");
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDvalue.(int)
	shortcode := c.Param("shortcode")

	_, err := config.DB.Exec("DELETE FROM urls WHERE user_id = ? AND short_code = ?", userID, shortcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete URL"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "URL deleted successfully",
	})
}

func RedirectUrl(c *gin.Context){
	shortcode := c.Param("shortcode")

	var originalUrl string

	err:=config.DB.QueryRow(
		"SELECT original_url FROM urls WHERE short_code = ?",shortcode,
	).Scan(&originalUrl)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound,gin.H{"error":"URL not Found"})
		return
	}
	
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Database Error"});
		return
	}
	_,err = config.DB.Exec("UPDATE urls SET clicks = clicks+1 WHERE short_code=?",shortcode)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed To Update click count"})
		return
	}
	c.Redirect(http.StatusFound,originalUrl)
}

func GetUserUrls(c *gin.Context) {
	userIDvalue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDvalue.(int)

	rows, err := config.DB.Query(
		"SELECT id, user_id, short_code, original_url, clicks, created_at, expires_at FROM urls WHERE user_id = ? ORDER BY created_at DESC",
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URLs"})
		return
	}
	defer rows.Close()

	var urls []models.Url

	for rows.Next() {
		var url models.Url
		err := rows.Scan(&url.ID, &url.UserID, &url.ShortCode, &url.OriginalUrl, &url.Clicks, &url.CreatedAt, &url.ExpiresAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning URLs"})
			return
		}
		urls = append(urls, url)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over URLs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"urls":   urls,
	})
}
