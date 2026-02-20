package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"server/models"
	"server/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/api/idtoken"
)

func GoogleAuth(c *gin.Context){
	var req models.GoogleRequest
	if err := c.ShouldBindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid Request"})
		return
	}
	payload,err:= idtoken.Validate(
		context.Background(),
		req.Token,
		os.Getenv("GOOGLE_CLIENT_ID"),
	)
	if err!=nil{
		log.Printf("Token validation error %v",err)
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid Google Token"})
		return
	}

	googleID := payload.Subject
	name,_ := payload.Claims["name"].(string)
	email,_ := payload.Claims["email"].(string)
	avatarURL,_ := payload.Claims["picture"].(string)

	if email == ""{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Email not found"})
		return
	}

	var userID int
	var userRole string

	err = config.DB.QueryRow("SELECT id,role FROM users WHERE email = ?",email).Scan(&userID,&userRole)

	if err == sql.ErrNoRows {
		//new user entry
		result , err := config.DB.Exec(
			`INSERT INTO users (google_id,display_name,email,avatar_url,last_login_at) VALUES (?,?,?,?,NOW())`,googleID,name,email,avatarURL,
		)
		if err != nil{
			log.Printf("Database insert error: %v",err)
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed To create User"})
			return
		}
		lastID, err := result.LastInsertId()
		if err != nil {
			log.Printf("Failed to get last insert ID: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID"})
			return
		}

		userID = int(lastID)
		userRole = "user"

	}else if err != nil {
		log.Printf("Database query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}else {
		//updating existing user
		_,err = config.DB.Exec(`UPDATE users SET last_login_at=NOW(),display_name = ? , avatar_url = ? , google_id = ? WHERE id = ?`,name,avatarURL,googleID,userID)

		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}
	}

	//JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":userID,
		"email":email,
		"role":userRole,
	})
	tokenString,err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.SetCookie(
		"auth_token",
		tokenString,
		0,
		"/",
		os.Getenv("COOKIE_DOMAIN"),
		true,
		true,
	)
	c.JSON(http.StatusOK,gin.H{
		"message":"Logged In Successfully",	
	})
}

func Logout(c *gin.Context) {
    c.SetCookie("auth_token", "", -1, "/", os.Getenv("COOKIE_DOMAIN"), false, true)
    c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}