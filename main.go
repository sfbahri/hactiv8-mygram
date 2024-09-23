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

	// Create tables
	// err = createTables(db)
	// if err != nil {
	// 	log.Fatal("Failed to create tables:", err)
	// }

	r := gin.Default()
	api := r.Group("/api")
	routes.UserRoutes(api, db)
	routes.PhotoRoutes(api, db)
	routes.CommentRoutes(api, db)
	routes.SocialMediaRoutes(api, db)

	r.Run(":8181")
}

// func createTables(db *sql.DB) error {
// 	queries := []string{
// 		`CREATE TABLE IF NOT EXISTS users (
// 			id SERIAL PRIMARY KEY,
// 			username VARCHAR(100) UNIQUE NOT NULL,
// 			email VARCHAR(100) UNIQUE NOT NULL,
// 			password VARCHAR(100) NOT NULL,
// 			age INT,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		)`,

// 		`CREATE TABLE IF NOT EXISTS photos (
// 			id SERIAL PRIMARY KEY,
// 			title VARCHAR(100) NOT NULL,
// 			caption TEXT,
// 			photo_url TEXT NOT NULL,
// 			user_id INT REFERENCES users(id),
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		)`,

// 		`CREATE TABLE IF NOT EXISTS comments (
// 			id SERIAL PRIMARY KEY,
// 			user_id INT REFERENCES users(id),
// 			photo_id INT REFERENCES photos(id),
// 			message TEXT NOT NULL,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		)`,

// 		`CREATE TABLE IF NOT EXISTS socialmedia (
// 			id SERIAL PRIMARY KEY,
// 			name VARCHAR(100) NOT NULL,
// 			social_media_url TEXT NOT NULL,
// 			user_id INT REFERENCES users(id),
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		)`,
// 	}

// 	for _, query := range queries {
// 		_, err := db.Exec(query)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
