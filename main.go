package main

import (
	"database/sql"
	"log"
	"mygram/config"
	"mygram/routes"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {

	var err error
	db, err = config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	r := gin.Default()
	api := r.Group("/api")
	routes.UserRoutes(api, db)
	routes.PhotoRoutes(api, db)
	routes.CommentRoutes(api, db)
	routes.SocialMediaRoutes(api, db)

	r.Run(":8181")
}
