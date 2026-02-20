package handlers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

func urlAvailablityChecker(c *gin.Context){

	var input models.Url

	

	err:=config.DB.QueryRow("SELECT short_code FROM urls WHERE short_code = ?",input.ShortCode)

	if err==nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"error":"Short Code Already Exists",
		});
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"status":true,
		"message":"Short Code Available",
	})

}