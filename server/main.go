
package main

import(
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"server/routes"
)

func main(){
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	log.Println("Starting URL Shortener Server (In-Memory Mode)")
	
	r := gin.Default()
	routes.Routes(r)

	r.Run(":5000")
}