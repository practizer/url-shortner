
package main

import(
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"

	"server/routes"
	"server/config"
)
var DB *sql.DB

func main(){

	config.InitDB()

	r := gin.Default()
	routes.Routes(r)

	r.Run(":5000")
}