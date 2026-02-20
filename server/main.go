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
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	routes.Routes(r)

	r.Run(":5000")
}