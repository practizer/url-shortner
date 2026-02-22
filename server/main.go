
package main

import(
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"server/routes"
	"server/config"
)
var DB *sql.DB

func main(){
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config.InitDB()
	config.InitRedis()

	r := gin.Default()
	routes.Routes(r)

	r.Run(":5000")
}