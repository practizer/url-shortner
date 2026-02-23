package handlers

import (
	"net/http"
	"net/url"
	"server/config"
	"server/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func UrlAvailabilityChecker(c *gin.Context) {

	var input struct {
		ShortCode string `json:"short_code"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Invalid request body",
		})
		return
	}

	if strings.TrimSpace(input.ShortCode) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Short code cannot be empty",
		})
		return
	}

	isAvailable := config.URLStore.CheckAvailability(input.ShortCode)

	if !isAvailable {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Short Code Already Exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Short Code Available",
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

	shortCode := strings.TrimSpace(input.ShortCode)
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short code cannot be empty"})
		return
	}

	// Create URL in storage
	err = config.URLStore.CreateURL(shortCode, originalUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "URL added Successfully",
	})
}

func DeleteUrl(c *gin.Context) {
	shortcode := c.Param("shortcode")

	err := config.URLStore.DeleteURL(shortcode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "URL deleted successfully",
	})
}

func RedirectUrl(c *gin.Context){
	shortcode := c.Param("shortcode")

	url := config.URLStore.GetURL(shortcode)
	if url == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Increment clicks
	config.URLStore.IncrementClicks(shortcode)

	c.Redirect(http.StatusFound, url.OriginalUrl)
}

func GetUserUrls(c *gin.Context) {
	// Get all URLs (no user concept in in-memory mode)
	urls := config.URLStore.GetAllURLs()

	if urls == nil {
		urls = []*models.Url{}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"urls":   urls,
	})
}
